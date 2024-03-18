package test

import (
	"github.com/cosmos/gogoproto/types/any"
	"github.com/cosmos/gogoproto/types/any/internal"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
)

// We implement a minimal proto.Message interface
func (d *Dog) XXX_MessageName() string { return "tests/dog" }

var (
	_ Animal        = (*Dog)(nil)
	_ proto.Message = (*Dog)(nil)
)

func TestAnimalPackUnpack(t *testing.T) {
	registry := internal.NewInterfaceRegistry()
	registry.RegisterInterface("Animal", (*Animal)(nil))
	registry.RegisterImplementations(
		(*Animal)(nil),
		&Dog{},
	)

	spot := &Dog{Name: "Spot"}
	var animal Animal

	// with cache
	any, err := types.NewAnyWithCacheWithValue(spot)
	require.NoError(t, err)
	require.Equal(t, spot, any.GetCachedValue())
	err = registry.UnpackAny(any, &animal)
	require.NoError(t, err)
	require.Equal(t, spot, animal)

	// without cache
	any.ResetCachedValue()
	err = registry.UnpackAny(any, &animal)
	require.NoError(t, err)
	require.Equal(t, spot, animal)
}

func TestString(t *testing.T) {
	require := require.New(t)
	spot := &Dog{Name: "Spot"}
	any, err := types.NewAnyWithCacheWithValue(spot)
	require.NoError(err)

	require.Equal("&Any{TypeUrl:/tests/dog,Value:[18 4 83 112 111 116],XXX_unrecognized:[]}", any.String())
	require.Equal(`&Any{TypeUrl: "/tests/dog",
  Value: []byte{0x12, 0x4, 0x53, 0x70, 0x6f, 0x74}
}`, any.GoString())
}
