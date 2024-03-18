package test

import (
	types "github.com/cosmos/gogoproto/types/any"
	"github.com/cosmos/gogoproto/types/any/internal"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
)

func NewTestInterfaceRegistry() internal.InterfaceRegistry {
	registry := internal.NewInterfaceRegistry()
	RegisterInterfaces(registry)
	return registry
}

func RegisterInterfaces(registry internal.InterfaceRegistry) {
	registry.RegisterInterface("Animal", (*Animal)(nil))
	registry.RegisterImplementations(
		(*Animal)(nil),
		&Dog{},
		&Cat{},
	)
	registry.RegisterImplementations(
		(*HasAnimalI)(nil),
		&HasAnimal{},
	)
	registry.RegisterImplementations(
		(*HasHasAnimalI)(nil),
		&HasHasAnimal{},
	)
}

func TestAnyPackUnpack(t *testing.T) {
	registry := internal.NewInterfaceRegistry()

	spot := &Dog{Name: "Spot"}
	var animal Animal

	// with cache
	any, err := types.NewAnyWithCacheWithValue(spot)
	require.NoError(t, err)
	require.Equal(t, spot, any.GetCachedValue())
	err = registry.UnpackAny(any, &animal)
	require.NoError(t, err)
	require.Equal(t, spot, animal)
}

type TestI interface {
	DoSomething()
}

// A struct that has the same typeURL as Dog, but is actually another
// concrete type.
type FakeDog struct{}

var (
	_ proto.Message = &FakeDog{}
	_ Animal        = &FakeDog{}
)

// dummy implementation of proto.Message and Animal
func (dog FakeDog) Reset()                  {}
func (dog FakeDog) String() string          { return "fakedog" }
func (dog FakeDog) ProtoMessage()           {}
func (dog FakeDog) XXX_MessageName() string { return proto.MessageName(&Dog{}) }
func (dog FakeDog) Greet() string           { return "fakedog" }

func TestRegister(t *testing.T) {
	registry := NewTestInterfaceRegistry()
	registry.RegisterInterface("Animal", (*Animal)(nil))
	registry.RegisterInterface("TestI", (*TestI)(nil))

	// Happy path.
	require.NotPanics(t, func() {
		registry.RegisterImplementations((*Animal)(nil), &Dog{})
	})

	// Dog doesn't implement TestI
	require.Panics(t, func() {
		registry.RegisterImplementations((*TestI)(nil), &Dog{})
	})

	// nil proto message
	require.Panics(t, func() {
		registry.RegisterImplementations((*TestI)(nil), nil)
	})

	// Not an interface.
	require.Panics(t, func() {
		registry.RegisterInterface("not_an_interface", (*Dog)(nil))
	})

	// Duplicate registration with same concrete type.
	require.NotPanics(t, func() {
		registry.RegisterImplementations((*Animal)(nil), &Dog{})
	})

	// Duplicate registration with different concrete type on same typeURL.
	require.PanicsWithError(
		t,
		"concrete type *test.Dog has already been registered under typeURL /tests/dog, cannot register *test.FakeDog under same typeURL. "+
			"This usually means that there are conflicting modules registering different concrete types for a same interface implementation",
		func() {
			registry.RegisterImplementations((*Animal)(nil), &FakeDog{})
		},
	)
}

func TestUnpackInterfaces(t *testing.T) {
	registry := NewTestInterfaceRegistry()

	spot := &Dog{Name: "Spot"}
	any, err := types.NewAnyWithCacheWithValue(spot)
	require.NoError(t, err)

	hasAny := HasAnimal{
		Animal: any,
		X:      1,
	}
	bz, err := hasAny.Marshal()
	require.NoError(t, err)

	var hasAny2 HasAnimal
	err = hasAny2.Unmarshal(bz)
	require.NoError(t, err)

	err = types.UnpackInterfaces(hasAny2, registry)
	require.NoError(t, err)

	require.Equal(t, spot, hasAny2.Animal.GetCachedValue())
}

func TestNested(t *testing.T) {
	registry := NewTestInterfaceRegistry()

	spot := &Dog{Name: "Spot"}
	any, err := types.NewAnyWithCacheWithValue(spot)
	require.NoError(t, err)

	ha := &HasAnimal{Animal: any}
	any2, err := types.NewAnyWithCacheWithValue(ha)
	require.NoError(t, err)

	hha := &HasHasAnimal{HasAnimal: any2}
	any3, err := types.NewAnyWithCacheWithValue(hha)
	require.NoError(t, err)

	hhha := HasHasHasAnimal{HasHasAnimal: any3}

	// marshal
	bz, err := hhha.Marshal()
	require.NoError(t, err)

	// unmarshal
	var hhha2 HasHasHasAnimal
	err = hhha2.Unmarshal(bz)
	require.NoError(t, err)
	err = types.UnpackInterfaces(hhha2, registry)
	require.NoError(t, err)

	require.Equal(t, spot, hhha2.TheHasHasAnimal().TheHasAnimal().TheAnimal())
}

func TestAny_ProtoJSON(t *testing.T) {
	spot := &Dog{Name: "Spot"}
	any, err := types.NewAnyWithCacheWithValue(spot)
	require.NoError(t, err)

	jm := &jsonpb.Marshaler{}
	json, err := jm.MarshalToString(any)
	require.NoError(t, err)
	require.Equal(t, "{\"@type\":\"/Dog\",\"name\":\"Spot\"}", json)

	registry := NewTestInterfaceRegistry()
	jum := &jsonpb.Unmarshaler{}
	var any2 types.Any
	err = jum.Unmarshal(strings.NewReader(json), &any2)
	require.NoError(t, err)
	var animal Animal
	err = registry.UnpackAny(&any2, &animal)
	require.NoError(t, err)
	require.Equal(t, spot, animal)

	ha := &HasAnimal{
		Animal: any,
	}
	err = ha.UnpackInterfaces(types.ProtoJSONPacker{JSONPBMarshaler: jm})
	require.NoError(t, err)
	json, err = jm.MarshalToString(ha)
	require.NoError(t, err)
	require.Equal(t, "{\"animal\":{\"@type\":\"/Dog\",\"name\":\"Spot\"}}", json)

	require.NoError(t, err)
	var ha2 HasAnimal
	err = jum.Unmarshal(strings.NewReader(json), &ha2)
	require.NoError(t, err)
	err = ha2.UnpackInterfaces(registry)
	require.NoError(t, err)
	require.Equal(t, spot, ha2.Animal.GetCachedValue())
}
