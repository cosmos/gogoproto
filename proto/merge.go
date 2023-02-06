package proto

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
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
// file descriptors registered with both gogoproto and protoregistry. When
// merging it also performs the following checks:
// - check that all file descriptors' import paths are correct (i.e. match
// their fully-qualified package name).
// - check that if both gogo and protoregistry import the same file descriptor,
// that these are identical under proto.Equal.
func MergedFileDescriptors() (*descriptorpb.FileDescriptorSet, error) {
	fds := &descriptorpb.FileDescriptorSet{}

	// load gogo proto file descriptors
	gogoFds := AllFileDescriptors()
	gogoFdsMap := map[string]*descriptorpb.FileDescriptorProto{}
	for _, compressedBz := range gogoFds {
		rdr, err := gzip.NewReader(bytes.NewReader(compressedBz))
		if err != nil {
			return nil, err
		}

		bz, err := io.ReadAll(rdr)
		if err != nil {
			return nil, err
		}

		fd := &descriptorpb.FileDescriptorProto{}
		err = protov2.Unmarshal(bz, fd)
		if err != nil {
			return nil, err
		}

		err = CheckImportPath(*fd.Name, *fd.Package)
		if err != nil {
			return nil, err
		}

		fds.File = append(fds.File, fd)
		gogoFdsMap[*fd.Name] = fd
	}

	// load any protoregistry file descriptors not in gogo
	protoregistry.GlobalFiles.RangeFiles(func(fileDescriptor protoreflect.FileDescriptor) bool {
		fd := protodesc.ToFileDescriptorProto(fileDescriptor)
		if err := CheckImportPath(*fd.Name, *fd.Package); err != nil {
			panic(err)
		}

		gogoFd, found := gogoFdsMap[fileDescriptor.Path()]
		// If we already loaded gogo's file descriptor, compare that the 2
		// are strictly equal.
		if found {
			if !protov2.Equal(gogoFd, fd) {
				diff := cmp.Diff(fd, gogoFd, protocmp.Transform())
				panic(fmt.Errorf("got different file descriptors for %s; %s", *fd.Name, diff))
			}
		} else {
			fds.File = append(fds.File, protodesc.ToFileDescriptorProto(fileDescriptor))
		}

		return true
	})

	slices.SortFunc(fds.File, func(x, y *descriptorpb.FileDescriptorProto) bool {
		return *x.Name < *y.Name
	})

	return fds, nil
}

// MergedRegistry returns a *protoregistry.Files that acts as a single registry
// which contains all the file descriptors registered with both gogoproto and
// protoregistry.
func MergedRegistry() (*protoregistry.Files, error) {
	fds, err := MergedFileDescriptors()
	if err != nil {
		return nil, err
	}

	return protodesc.NewFiles(fds)
}

// CheckImportPath checks that the import path of the given file descriptor
// matches its fully qualified package name.
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
