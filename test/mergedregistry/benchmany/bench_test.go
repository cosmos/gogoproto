package benchmany_test

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	protov2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// This benchmark creates dynamic file descriptors so that we can quantify performance
// of larger FD sets.
//
// Importing jhump/protoreflect brings in default google types,
// which breaks the primary mergedregistry tests;
// hence the standalone package.
func BenchmarkMergedFileDescriptors(b *testing.B) {
	bcs := []struct {
		globalSize int
		appSize    int
	}{
		{globalSize: 32, appSize: 0},
		{globalSize: 32, appSize: 32},
		{globalSize: 64, appSize: 64},
		{globalSize: 128, appSize: 128},
	}

	for _, bc := range bcs {
		globalSize := bc.globalSize
		appSize := bc.appSize
		wantSize := globalSize + appSize

		b.Run(fmt.Sprintf("%d global, %d app", globalSize, appSize), func(b *testing.B) {
			b.ReportAllocs()

			gf := new(protoregistry.Files)

			for i := 0; i < globalSize; i++ {
				gf.RegisterFile(generateFile(i).UnwrapFile())
			}

			apps := generateAppFiles(appSize, globalSize)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				fds, err := proto.MergedFileDescriptors(gf, apps)
				if err != nil {
					b.Fatal(err)
				}

				if len(fds.File) != wantSize {
					b.Fatalf("expected %d FDs, got %d", wantSize, len(fds.File))
				}
			}
		})
	}
}

func generateFile(n int) *desc.FileDescriptor {
	file := builder.NewFile(fmt.Sprintf("/example/foo%d/bar%d.proto", n, n))

	msg := builder.NewMessage(fmt.Sprintf("Msg%d", n)).
		AddField(builder.NewField("id", builder.FieldTypeInt64())).
		AddField(builder.NewField("name", builder.FieldTypeString()))
	file.AddMessage(msg)

	fd, err := file.Build()
	if err != nil {
		panic(err)
	}

	return fd
}

func generateAppFiles(size, offset int) map[string][]byte {
	out := make(map[string][]byte, size)

	buf := new(bytes.Buffer)
	for i := 0; i < size; i++ {
		fd := generateFile(i + offset).AsFileDescriptorProto()

		bz, err := protov2.Marshal(fd)
		if err != nil {
			panic(err)
		}

		buf.Reset()
		gzw := gzip.NewWriter(buf)
		if _, err := gzw.Write(bz); err != nil {
			panic(err)
		}
		if err := gzw.Close(); err != nil {
			panic(err)
		}

		name := strconv.Itoa(i) // Map key doesn't matter for merging the app files.
		out[name] = bytes.Clone(buf.Bytes())
	}

	return out
}
