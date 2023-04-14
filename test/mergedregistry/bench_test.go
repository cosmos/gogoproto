package mergedregistry_test

import (
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func BenchmarkDefaultGlobalMergedFileDescriptors(b *testing.B) {
	b.ReportAllocs()

	wantSize := protoregistry.GlobalFiles.NumFiles() + len(proto.AllFileDescriptors())

	for i := 0; i < b.N; i++ {
		fds, err := proto.MergedGlobalFileDescriptors()
		if err != nil {
			b.Fatal(err)
		}

		if len(fds.File) != wantSize {
			b.Fatalf("expected %d FDs, got %d", wantSize, len(fds.File))
		}
	}
}
