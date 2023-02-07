package mergedregistry_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
)

func TestMergedRegistry(t *testing.T) {
	reg, err := proto.MergedRegistry()
	if err != nil {
		t.Error(err)
	}
	// There are 11 .proto files in `google/protobuf` directory.
	if reg.NumFiles() != 11 {
		t.Error(fmt.Errorf("expected 11 files, got %d", reg.NumFiles()))
	}
}
