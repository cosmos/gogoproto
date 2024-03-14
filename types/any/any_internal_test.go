package types

import (
	"github.com/cosmos/gogoproto/test/testdata"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnyPackUnpack(t *testing.T) {
	registry := NewInterfaceRegistry()
	registry.RegisterInterface("Animal", (*testdata.HasAnimal)(nil))
	registry.RegisterImplementations(
		(*testdata.HasAnimal)(nil),
		&testdata.Dog{},
	)

	spot := &testdata.Dog{Name: "Spot"}
	var animal testdata.HasAnimal

	// with cache
	any, err := NewAnyWithValue(spot)
	require.NoError(t, err)
	require.Equal(t, spot, any.GetCachedValue())
	err = registry.UnpackAny(any, &animal)
	require.NoError(t, err)
	require.Equal(t, spot, animal)

	// without cache
	any.cachedValue = nil
	err = registry.UnpackAny(any, &animal)
	require.NoError(t, err)
	require.Equal(t, spot, animal)
}

func TestString(t *testing.T) {
	require := require.New(t)
	spot := &testdata.Dog{Name: "Spot"}
	any, err := NewAnyWithValue(spot)
	require.NoError(err)

	require.Equal("&Any{TypeUrl:/tests/dog,Value:[10 4 83 112 111 116],XXX_unrecognized:[]}", any.String())
	require.Equal(`&Any{TypeUrl: "/tests/dog",
  Value: []byte{0xa, 0x4, 0x53, 0x70, 0x6f, 0x74}
}`, any.GoString())
}
