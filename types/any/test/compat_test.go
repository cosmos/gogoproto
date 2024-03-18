package test

import (
	types "github.com/cosmos/gogoproto/types/any"

	"testing"

	"github.com/stretchr/testify/suite"
	amino "github.com/tendermint/go-amino"
)

type TypeWithInterface struct {
	Animal Animal `json:"animal"`
	X      int64  `json:"x,omitempty"`
}

type Suite struct {
	suite.Suite
	cdc  *amino.Codec
	a    TypeWithInterface
	b    HasAnimal
	spot *Dog
}

func (s *Suite) SetupTest() {
	s.cdc = amino.NewCodec()
	s.cdc.RegisterInterface((*Animal)(nil), nil)
	s.cdc.RegisterConcrete(&Dog{}, "test/Dog", nil)

	s.spot = &Dog{Size_: "small", Name: "Spot"}
	s.a = TypeWithInterface{Animal: s.spot}

	any, err := types.NewAnyWithCacheWithValue(s.spot)
	s.Require().NoError(err)
	s.b = HasAnimal{Animal: any}
}

func (s *Suite) TestAminoBinary() {
	bz, err := s.cdc.MarshalBinaryBare(s.a)
	s.Require().NoError(err)

	// expect plain amino marshal to fail
	_, err = s.cdc.MarshalBinaryBare(s.b)
	s.Require().Error(err)

	// expect unpack interfaces before amino marshal to succeed
	err = types.UnpackInterfaces(s.b, types.AminoPacker{Cdc: s.cdc})
	s.Require().NoError(err)
	bz2, err := s.cdc.MarshalBinaryBare(s.b)
	s.Require().NoError(err)
	s.Require().Equal(bz, bz2)

	var c HasAnimal
	err = s.cdc.UnmarshalBinaryBare(bz, &c)
	s.Require().NoError(err)
	err = types.UnpackInterfaces(c, types.AminoUnpacker{Cdc: s.cdc})
	s.Require().NoError(err)
	s.Require().Equal(s.spot, c.Animal.GetCachedValue())
}

func (s *Suite) TestAminoJSON() {
	bz, err := s.cdc.MarshalJSON(s.a)
	s.Require().NoError(err)

	// expect plain amino marshal to fail
	_, err = s.cdc.MarshalJSON(s.b)
	s.Require().Error(err)

	// expect unpack interfaces before amino marshal to succeed
	err = types.UnpackInterfaces(s.b, types.AminoJSONPacker{Cdc: s.cdc})
	s.Require().NoError(err)
	bz2, err := s.cdc.MarshalJSON(s.b)
	s.Require().NoError(err)
	s.Require().Equal(string(bz), string(bz2))

	var c HasAnimal
	err = s.cdc.UnmarshalJSON(bz, &c)
	s.Require().NoError(err)
	err = types.UnpackInterfaces(c, types.AminoJSONUnpacker{Cdc: s.cdc})
	s.Require().NoError(err)
	s.Require().Equal(s.spot, c.Animal.GetCachedValue())
}

func (s *Suite) TestNested() {
	s.cdc.RegisterInterface((*HasAnimalI)(nil), nil)
	s.cdc.RegisterInterface((*HasHasAnimalI)(nil), nil)
	s.cdc.RegisterConcrete(&HasAnimal{}, "test/HasAnimal", nil)
	s.cdc.RegisterConcrete(&HasHasAnimal{}, "test/HasHasAnimal", nil)
	s.cdc.RegisterConcrete(&HasHasHasAnimal{}, "test/HasHasHasAnimal", nil)

	any, err := types.NewAnyWithCacheWithValue(&s.b)
	s.Require().NoError(err)
	hha := HasHasAnimal{HasAnimal: any}
	any2, err := types.NewAnyWithCacheWithValue(&hha)
	s.Require().NoError(err)
	hhha := HasHasHasAnimal{HasHasAnimal: any2}

	// marshal
	err = types.UnpackInterfaces(hhha, types.AminoPacker{Cdc: s.cdc})
	s.Require().NoError(err)
	bz, err := s.cdc.MarshalBinaryBare(hhha)
	s.Require().NoError(err)

	// unmarshal
	var hhha2 HasHasHasAnimal
	err = s.cdc.UnmarshalBinaryBare(bz, &hhha2)
	s.Require().NoError(err)
	err = types.UnpackInterfaces(hhha2, types.AminoUnpacker{Cdc: s.cdc})
	s.Require().NoError(err)

	s.Require().Equal(s.spot, hhha2.TheHasHasAnimal().TheHasAnimal().TheAnimal())

	// json marshal
	err = types.UnpackInterfaces(hhha, types.AminoJSONPacker{Cdc: s.cdc})
	s.Require().NoError(err)
	jsonBz, err := s.cdc.MarshalJSON(hhha)
	s.Require().NoError(err)

	// json unmarshal
	var hhha3 HasHasHasAnimal
	err = s.cdc.UnmarshalJSON(jsonBz, &hhha3)
	s.Require().NoError(err)
	err = types.UnpackInterfaces(hhha3, types.AminoJSONUnpacker{Cdc: s.cdc})
	s.Require().NoError(err)

	s.Require().Equal(s.spot, hhha3.TheHasHasAnimal().TheHasAnimal().TheAnimal())
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}
