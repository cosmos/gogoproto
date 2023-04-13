package proto

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"strings"

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
// If there are any incorrect import paths that do not match
// the fully qualified package name, or if there is a common file descriptor
// that differs accross globalFiles and appFiles, an error is returned.
func MergedFileDescriptors(globalFiles *protoregistry.Files, appFiles map[string][]byte) (*descriptorpb.FileDescriptorSet, error) {
	return mergedFileDescriptors(globalFiles, appFiles)
}

// MergedGlobalFileDescriptors calls MergedGlobalFileDescriptors
// with [protoregistry.GlobalFiles] and all files
// registered through [RegisterFile].
func MergedGlobalFileDescriptors() (*descriptorpb.FileDescriptorSet, error) {
	return MergedFileDescriptors(protoregistry.GlobalFiles, protoFiles)
}

func mergedFileDescriptors(globalFiles *protoregistry.Files, appFiles map[string][]byte) (*descriptorpb.FileDescriptorSet, error) {
	const debug = true // TODO: remove this constant

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
