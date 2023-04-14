package proto

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/slices"
	protov2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/descriptorpb"
)

// MergedFileDescriptors returns a single FileDescriptorSet containing all the
// file descriptors registered with the given globalFiles and appFiles.
//
// In contrast to MergedFileDescriptorsWithValidation,
// MergedFileDescriptors does not validate import paths
func MergedFileDescriptors(globalFiles *protoregistry.Files, appFiles map[string][]byte) (*descriptorpb.FileDescriptorSet, error) {
	return mergedFileDescriptors(globalFiles, appFiles, false)
}

// MergedFileDescriptorsWithValidation returns a single FileDescriptorSet containing all the
// file descriptors registered with the given globalFiles and appFiles.
//
// If there are any incorrect import paths that do not match
// the fully qualified package name, or if there is a common file descriptor
// that differs accross globalFiles and appFiles, an error is returned.
func MergedFileDescriptorsWithValidation(globalFiles *protoregistry.Files, appFiles map[string][]byte) (*descriptorpb.FileDescriptorSet, error) {
	return mergedFileDescriptors(globalFiles, appFiles, true)
}

// ConcurrentMergedFileDescriptors is an alternate implementation of MergedFileDescriptors
// that spreads work across all available CPU cores.
//
// This is intended as an temporary API to benchmark and test both implementations,
// before only offering the concurrent version.
func ConcurrentMergedFileDescriptors(globalFiles *protoregistry.Files, appFiles map[string][]byte) (*descriptorpb.FileDescriptorSet, error) {
	return concurrentMergeFileDescriptors(globalFiles, appFiles, false)
}

// ConcurrentMergedFileDescriptorsWithValidation is
// an alternate implementation of MergedFileDescriptorsWithValidation
// that spreads work across all available CPU cores.
//
// This is intended as an temporary API to benchmark and test both implementations,
// before only offering the concurrent version.
func ConcurrentMergedFileDescriptorsWithValidation(globalFiles *protoregistry.Files, appFiles map[string][]byte) (*descriptorpb.FileDescriptorSet, error) {
	return concurrentMergeFileDescriptors(globalFiles, appFiles, true)
}

// MergedGlobalFileDescriptors calls MergedFileDescriptors
// with [protoregistry.GlobalFiles] and all files
// registered through [RegisterFile].
func MergedGlobalFileDescriptors() (*descriptorpb.FileDescriptorSet, error) {
	return MergedFileDescriptors(protoregistry.GlobalFiles, protoFiles)
}

// MergedGlobalFileDescriptorsWithValidation calls MergedFileDescriptorsWithValidation
// with [protoregistry.GlobalFiles] and all files
// registered through [RegisterFile].
func MergedGlobalFileDescriptorsWithValidation() (*descriptorpb.FileDescriptorSet, error) {
	return MergedFileDescriptorsWithValidation(protoregistry.GlobalFiles, protoFiles)
}

func mergedFileDescriptors(globalFiles *protoregistry.Files, appFiles map[string][]byte, debug bool) (*descriptorpb.FileDescriptorSet, error) {
	fds := &descriptorpb.FileDescriptorSet{
		// Pre-size the Files since we are going to copy them
		// when we range over globalFiles.
		File: make([]*descriptorpb.FileDescriptorProto, 0, globalFiles.NumFiles()),
	}

	// While combing through the file descriptors, we'll also log any errors
	// we encounter -- only if debug is true. Otherwise, we will skip the work
	// to check import path or file descriptor differences.
	var (
		checkImportErr []string
		diffErr        []string
	)

	// Add protoregistry file descriptors to our final file descriptor set.
	globalFiles.RangeFiles(func(fileDescriptor protoreflect.FileDescriptor) bool {
		if debug {
			fd := protodesc.ToFileDescriptorProto(fileDescriptor)
			if err := CheckImportPath(fd.GetName(), fd.GetPackage()); err != nil {
				checkImportErr = append(checkImportErr, err.Error())
			}
		}

		fds.File = append(fds.File, protodesc.ToFileDescriptorProto(fileDescriptor))

		return true
	})

	// Reuse a single gzip reader throughout the loop,
	// so we don't have to repeatedly allocate new readers.
	gzr := new(gzip.Reader)

	// Also reuse a single byte buffer for each gzip read.
	buf := new(bytes.Buffer)

	// Load gogo proto file descriptors.
	// Normal usage would go through the AllFileDescriptors method,
	// which returns a copy of the package-level map.
	//
	// In tests especially, this method can be part of a hot call stack.
	// Because we are in the same package and we know what we're doing,
	// we can read from the raw map.
	for _, compressedBz := range appFiles {
		if err := gzr.Reset(bytes.NewReader(compressedBz)); err != nil {
			return nil, err
		}

		buf.Reset()
		if _, err := buf.ReadFrom(gzr); err != nil {
			return nil, err
		}

		fd := &descriptorpb.FileDescriptorProto{}
		if err := protov2.Unmarshal(buf.Bytes(), fd); err != nil {
			return nil, err
		}

		if debug {
			err := CheckImportPath(fd.GetName(), fd.GetPackage())
			if err != nil {
				checkImportErr = append(checkImportErr, err.Error())
			}
		}

		// If it's not in the protoregistry file descriptors, add it.
		protoregFd, err := globalFiles.FindFileByPath(*fd.Name)
		// If we already loaded gogo's file descriptor, compare that the 2
		// are strictly equal, or else log a warning.
		if err != nil {
			// If we have a NotFound error, we add this file descriptor to the
			// final file descriptor set.
			if errors.Is(err, protoregistry.NotFound) {
				fds.File = append(fds.File, fd)
			} else {
				return nil, err
			}
		} else {
			// If there's a mismatch, we log a warning. If there was no
			// mismatch, then we do nothing, and take the protoregistry file
			// descriptor as the correct one.
			if debug && !protov2.Equal(protodesc.ToFileDescriptorProto(protoregFd), fd) {
				diff := cmp.Diff(protodesc.ToFileDescriptorProto(protoregFd), fd, protocmp.Transform())
				diffErr = append(diffErr, fmt.Sprintf("Mismatch in %s:\n%s", *fd.Name, diff))
			}
		}
	}

	if debug {
		errStr := new(bytes.Buffer)
		if len(checkImportErr) > 0 {
			fmt.Fprintf(errStr, "Got %d file descriptor import path errors:\n\t%s\n", len(checkImportErr), strings.Join(checkImportErr, "\n\t"))
		}
		if len(diffErr) > 0 {
			fmt.Fprintf(errStr, "Got %d file descriptor mismatches. Make sure gogoproto and protoregistry use the same .proto files. '-' lines are from protoregistry, '+' lines from gogo's registry.\n\n\t%s\n", len(diffErr), strings.Join(diffErr, "\n\t"))
		}
		if errStr.Len() > 0 {
			return nil, fmt.Errorf(errStr.String())
		}
	}

	slices.SortFunc(fds.File, func(x, y *descriptorpb.FileDescriptorProto) bool {
		return *x.Name < *y.Name
	})

	return fds, nil
}

// MergedRegistry returns a *protoregistry.Files that acts as a single registry
// which contains all the file descriptors registered with both gogoproto and
// protoregistry (the latter taking precendence if there's a mismatch).
func MergedRegistry() (*protoregistry.Files, error) {
	fds, err := MergedGlobalFileDescriptors()
	if err != nil {
		return nil, err
	}

	return protodesc.NewFiles(fds)
}

// CheckImportPath checks that the import path of the given file descriptor
// matches its fully qualified package name. To mimic gogo's old behavior, the
// fdPackage string can be empty.
//
// Example:
// Proto file "google/protobuf/descriptor.proto" should be imported
// from OS path ./google/protobuf/descriptor.proto, relatively to a protoc
// path folder (-I flag).
func CheckImportPath(fdName, fdPackage string) error {
	expectedPrefix := strings.ReplaceAll(fdPackage, ".", "/") + "/"
	if !strings.HasPrefix(fdName, expectedPrefix) {
		return fmt.Errorf("file name %s does not start with expected %s; please make sure your folder structure matches the proto files fully-qualified names", fdName, expectedPrefix)
	}

	return nil
}

type descriptorErrorCollector struct {
	validate bool

	// Close the quit channel to request the collection goroutine to stop.
	quit chan struct{}

	// The done channel will be closed once the collection goroutine has finished.
	done chan struct{}

	ProcessErrCh chan error
	ImportErrCh  chan error
	DiffCh       chan string

	// Set at the end of collect().
	err error
}

func newDescriptorErrorCollector(chanSize int, validate bool) *descriptorErrorCollector {
	c := &descriptorErrorCollector{
		validate: validate,

		quit: make(chan struct{}),
		done: make(chan struct{}),

		ProcessErrCh: make(chan error, chanSize),
		ImportErrCh:  make(chan error, chanSize),
		DiffCh:       make(chan string, chanSize),
	}
	go c.collect()
	return c
}

// collect runs in its own goroutine,
// collecting process errors and import path and file descriptor differences.
// If any of those occur, it assigns to c.err.
// Stop the goroutine by closing c.quit.
// The goroutine closes c.done when it returns.
func (c *descriptorErrorCollector) collect() {
	defer close(c.done)

	var processErrs []error
	var importErrMsgs, diffs []string
LOOP:
	for {
		select {
		case <-c.quit:
			break LOOP

		case err := <-c.ProcessErrCh:
			// Always accept process errors.
			processErrs = append(processErrs, err)

		case err := <-c.ImportErrCh:
			if !c.validate {
				panic(fmt.Errorf("BUG: import error received when validate=false: %w", err))
			}
			importErrMsgs = append(importErrMsgs, err.Error())

		case diff := <-c.DiffCh:
			if !c.validate {
				panic(fmt.Errorf("BUG: diff received when validate=false: %s", diff))
			}
			diffs = append(diffs, diff)
		}
	}

	if len(processErrs) == 0 && len(importErrMsgs) == 0 && len(diffs) == 0 {
		// No errors received. Stop here so we don't assign to c.err.
		return
	}

	var buf bytes.Buffer
	for err := range processErrs {
		fmt.Fprintf(&buf, "Failure during processing: %v\n", err)
	}
	if len(importErrMsgs) > 0 {
		fmt.Fprintf(&buf, "Got %d file descriptor import path errors:\n\t%s\n", len(importErrMsgs), strings.Join(importErrMsgs, "\n\t"))
	}
	if len(diffs) > 0 {
		fmt.Fprintf(&buf, "Got %d file descriptor mismatches. Make sure gogoproto and protoregistry use the same .proto files. '-' lines are from protoregistry, '+' lines from gogo's registry.\n\n\t%s\n", len(diffs), strings.Join(diffs, "\n\t"))
	}

	c.err = errors.New(buf.String())
}

type descriptorProcessor struct {
	processWG    sync.WaitGroup
	globalFileCh chan protoreflect.FileDescriptor
	appFileCh    chan []byte

	fdWG sync.WaitGroup
	fdCh chan *descriptorpb.FileDescriptorProto
	fds  []*descriptorpb.FileDescriptorProto
}

func (p *descriptorProcessor) process(globalFiles *protoregistry.Files, ec *descriptorErrorCollector, validate bool) {
	defer p.processWG.Done()

	// Read the global files to exhaustion first.
	for fileDescriptor := range p.globalFileCh {
		fd := protodesc.ToFileDescriptorProto(fileDescriptor)
		if validate {
			if err := CheckImportPath(fd.GetName(), fd.GetPackage()); err != nil {
				// Track the import error but don't stop processing.
				ec.ImportErrCh <- err
			}
		}

		// Collect all the file descriptors in the collectFDs goroutine.
		p.fdCh <- fd
	}

	// Now handle all the app files.

	// Reuse a single gzip reader throughout the loop,
	// so we don't have to repeatedly allocate new readers.
	gzr := new(gzip.Reader)

	// Also reuse a single byte buffer for each gzip read.
	buf := new(bytes.Buffer)

	for compressedBz := range p.appFileCh {
		if err := gzr.Reset(bytes.NewReader(compressedBz)); err != nil {
			ec.ProcessErrCh <- err
			continue
		}

		buf.Reset()
		if _, err := buf.ReadFrom(gzr); err != nil {
			ec.ProcessErrCh <- err
			continue
		}

		fd := &descriptorpb.FileDescriptorProto{}
		if err := protov2.Unmarshal(buf.Bytes(), fd); err != nil {
			ec.ProcessErrCh <- err
			continue
		}

		if validate {
			// Ensure the import path on the app file is good.
			if err := CheckImportPath(fd.GetName(), fd.GetPackage()); err != nil {
				ec.ImportErrCh <- err
				// Don't break the loop here, continue to check for a file descriptor diff.
			}
		}

		// If the app FD is not in protoregistry, we need to track it.
		protoregFd, err := globalFiles.FindFileByPath(*fd.Name)
		if err != nil {
			if !errors.Is(err, protoregistry.NotFound) {
				// Non-nil error, and it wasn't a not found error.
				ec.ProcessErrCh <- err
				continue
			}
			// Otherwise it was a not found error, so add it.
			// At this point we can't validate.
			p.fdCh <- fd
			continue
		}

		if validate && !protov2.Equal(protodesc.ToFileDescriptorProto(protoregFd), fd) {
			diff := cmp.Diff(protodesc.ToFileDescriptorProto(protoregFd), fd, protocmp.Transform())
			ec.DiffCh <- fmt.Sprintf("Mismatch in %s:\n%s", *fd.Name, diff)
		}
	}
}

func (p *descriptorProcessor) collectFDs() {
	defer p.fdWG.Done()

	for fd := range p.fdCh {
		p.fds = append(p.fds, fd)
	}

	slices.SortFunc(p.fds, func(x, y *descriptorpb.FileDescriptorProto) bool {
		return *x.Name < *y.Name
	})
}

func concurrentMergeFileDescriptors(globalFiles *protoregistry.Files, appFiles map[string][]byte, validate bool) (*descriptorpb.FileDescriptorSet, error) {
	nProcs := runtime.GOMAXPROCS(0)

	ec := newDescriptorErrorCollector(nProcs, validate)

	p := &descriptorProcessor{
		globalFileCh: make(chan protoreflect.FileDescriptor, nProcs),
		appFileCh:    make(chan []byte, nProcs),

		fdCh: make(chan *descriptorpb.FileDescriptorProto, nProcs),
		fds:  make([]*descriptorpb.FileDescriptorProto, 0, globalFiles.NumFiles()),
	}

	// Start one goroutine per available processor.
	p.processWG.Add(nProcs)
	for i := 0; i < nProcs; i++ {
		go p.process(globalFiles, ec, validate)
	}

	p.fdWG.Add(1)
	go p.collectFDs()

	// Now synchronously iterate through globalFiles.
	globalFiles.RangeFiles(func(fileDescriptor protoreflect.FileDescriptor) bool {
		p.globalFileCh <- fileDescriptor
		return true
	})
	close(p.globalFileCh)

	// Same for appFiles.
	for _, bz := range appFiles {
		p.appFileCh <- bz
	}
	close(p.appFileCh)

	// File channels for processor have been closed.
	// Wait for the goroutine to complete.
	p.processWG.Wait()

	// Now close the FD channel since the processors are done,
	// and no more values will be sent.
	close(p.fdCh)

	// Wait until FD collection is complete.
	p.fdWG.Wait()

	// Stop the error collector, and if it found an error, return it.
	close(ec.quit)
	<-ec.done
	if ec.err != nil {
		return nil, ec.err
	}

	// Otherwise success.
	return &descriptorpb.FileDescriptorSet{
		File: p.fds,
	}, nil
}
