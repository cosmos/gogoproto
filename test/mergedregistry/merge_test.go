package mergedregistry_test

import (
	"testing"

	"github.com/cosmos/gogoproto/proto"
)

func TestMergedRegistry(t *testing.T) {
	_, err := proto.MergedRegistry()
	if err != nil {
		t.Error(err)
	}
}
