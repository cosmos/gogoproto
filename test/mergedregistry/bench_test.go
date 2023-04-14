package mergedregistry_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func BenchmarkDefaultGlobalMergedFileDescriptors(b *testing.B) {
	gf := protoregistry.GlobalFiles
	afd := proto.AllFileDescriptors()
	wantSize := gf.NumFiles() + len(afd)

	// Note the global and app counts here.
	// This benchmark is interesting for allocations,
	// but with only probably 11 file descriptors, we don't see much concurrency help.
	b.Run(fmt.Sprintf("%d global, %d app", gf.NumFiles(), len(afd)), func(b *testing.B) {
		b.Run("sequential", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				fds, err := proto.MergedGlobalFileDescriptors()
				if err != nil {
					b.Fatal(err)
				}

				if len(fds.File) != wantSize {
					b.Fatalf("expected %d FDs, got %d", wantSize, len(fds.File))
				}
			}
		})

		b.Run("concurrent", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				fds, err := proto.ConcurrentMergedFileDescriptors(gf, afd)
				if err != nil {
					b.Fatal(err)
				}

				if len(fds.File) != wantSize {
					b.Fatalf("expected %d FDs, got %d", wantSize, len(fds.File))
				}
			}
		})
	})
}
