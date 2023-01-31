// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: importduplicate/importduplicate.proto

package importduplicate

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_sortkeys "github.com/cosmos/gogoproto/sortkeys"
	proto1 "importduplicate/proto"
	sortkeys "importduplicate/sortkeys"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MapAndSortKeys struct {
	Key                  *sortkeys.Object `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	KeyValue             map[int32]string `protobuf:"bytes,2,rep,name=keyValue,proto3" json:"keyValue,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Value                *proto1.Subject  `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MapAndSortKeys) Reset()         { *m = MapAndSortKeys{} }
func (m *MapAndSortKeys) String() string { return proto.CompactTextString(m) }
func (*MapAndSortKeys) ProtoMessage()    {}
func (*MapAndSortKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_95811f4491d73b7d, []int{0}
}
func (m *MapAndSortKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapAndSortKeys.Unmarshal(m, b)
}
func (m *MapAndSortKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapAndSortKeys.Marshal(b, m, deterministic)
}
func (m *MapAndSortKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapAndSortKeys.Merge(m, src)
}
func (m *MapAndSortKeys) XXX_Size() int {
	return xxx_messageInfo_MapAndSortKeys.Size(m)
}
func (m *MapAndSortKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_MapAndSortKeys.DiscardUnknown(m)
}

var xxx_messageInfo_MapAndSortKeys proto.InternalMessageInfo

func (m *MapAndSortKeys) GetKey() *sortkeys.Object {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MapAndSortKeys) GetKeyValue() map[int32]string {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func (m *MapAndSortKeys) GetValue() *proto1.Subject {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*MapAndSortKeys)(nil), "importduplicate.MapAndSortKeys")
	proto.RegisterMapType((map[int32]string)(nil), "importduplicate.MapAndSortKeys.KeyValueEntry")
}

func init() {
	proto.RegisterFile("importduplicate/importduplicate.proto", fileDescriptor_95811f4491d73b7d)
}

var fileDescriptor_95811f4491d73b7d = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x49, 0x29, 0x2d, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c, 0x49, 0xd5, 0x47, 0xe3, 0xeb, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa3, 0x09, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0xe5,
	0xf4, 0x41, 0x2c, 0x88, 0x32, 0x29, 0x75, 0x74, 0xd3, 0x8a, 0xf3, 0x8b, 0x4a, 0xb2, 0x53, 0x2b,
	0x8b, 0xc1, 0x8c, 0xc4, 0xa4, 0x1c, 0xa8, 0x79, 0x52, 0x8a, 0xe8, 0x0a, 0x21, 0x46, 0x81, 0x49,
	0x88, 0x12, 0xa5, 0x26, 0x26, 0x2e, 0x3e, 0xdf, 0xc4, 0x02, 0xc7, 0xbc, 0x94, 0xe0, 0xfc, 0xa2,
	0x12, 0xef, 0xd4, 0xca, 0x62, 0x21, 0x23, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0x6e, 0x23, 0x05, 0x3d, 0x74, 0xa7, 0xc2, 0x2c, 0xd3, 0xf3, 0x4f, 0xca, 0x4a, 0x4d, 0x2e,
	0x09, 0x02, 0x29, 0x16, 0xf2, 0xe4, 0xe2, 0xc8, 0x4e, 0xad, 0x0c, 0x4b, 0xcc, 0x29, 0x4d, 0x95,
	0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0xd2, 0xc5, 0xd0, 0x88, 0x6a, 0x8d, 0x9e, 0x37, 0x54, 0xbd,
	0x6b, 0x5e, 0x49, 0x51, 0x65, 0x10, 0x5c, 0xbb, 0x90, 0x09, 0x17, 0x6b, 0x19, 0xd8, 0x1c, 0x66,
	0xb0, 0x03, 0xe4, 0xf4, 0xb0, 0x86, 0x95, 0x5e, 0x70, 0x29, 0xc4, 0x7a, 0x88, 0x62, 0x29, 0x6b,
	0x2e, 0x5e, 0x14, 0x03, 0x85, 0x04, 0x10, 0xbe, 0x60, 0x85, 0xb8, 0x51, 0x04, 0x66, 0x30, 0x93,
	0x02, 0xa3, 0x06, 0x27, 0x54, 0xa3, 0x15, 0x93, 0x05, 0xa3, 0x93, 0xe2, 0x87, 0x87, 0x72, 0x8c,
	0x3f, 0x1e, 0xca, 0x31, 0xae, 0x78, 0x24, 0xc7, 0xb8, 0xe3, 0x91, 0x1c, 0x63, 0x14, 0x7a, 0x4c,
	0x24, 0xb1, 0x81, 0x6d, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x4d, 0xcb, 0x45, 0xca,
	0x01, 0x00, 0x00,
}

func (this *MapAndSortKeys) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MapAndSortKeys)
	if !ok {
		that2, ok := that.(MapAndSortKeys)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Key.Equal(that1.Key) {
		return false
	}
	if len(this.KeyValue) != len(that1.KeyValue) {
		return false
	}
	for i := range this.KeyValue {
		if this.KeyValue[i] != that1.KeyValue[i] {
			return false
		}
	}
	if !this.Value.Equal(that1.Value) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MapAndSortKeys) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&importduplicate.MapAndSortKeys{")
	if this.Key != nil {
		s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	}
	keysForKeyValue := make([]int32, 0, len(this.KeyValue))
	for k := range this.KeyValue {
		keysForKeyValue = append(keysForKeyValue, k)
	}
	github_com_cosmos_gogoproto_sortkeys.Int32s(keysForKeyValue)
	mapStringForKeyValue := "map[int32]string{"
	for _, k := range keysForKeyValue {
		mapStringForKeyValue += fmt.Sprintf("%#v: %#v,", k, this.KeyValue[k])
	}
	mapStringForKeyValue += "}"
	if this.KeyValue != nil {
		s = append(s, "KeyValue: "+mapStringForKeyValue+",\n")
	}
	if this.Value != nil {
		s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringImportduplicate(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedMapAndSortKeys(r randyImportduplicate, easy bool) *MapAndSortKeys {
	this := &MapAndSortKeys{}
	if r.Intn(5) != 0 {
		this.Key = sortkeys.NewPopulatedObject(r, easy)
	}
	if r.Intn(5) != 0 {
		v1 := r.Intn(10)
		this.KeyValue = make(map[int32]string)
		for i := 0; i < v1; i++ {
			this.KeyValue[int32(r.Int31())] = randStringImportduplicate(r)
		}
	}
	if r.Intn(5) != 0 {
		this.Value = proto1.NewPopulatedSubject(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedImportduplicate(r, 4)
	}
	return this
}

type randyImportduplicate interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneImportduplicate(r randyImportduplicate) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringImportduplicate(r randyImportduplicate) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneImportduplicate(r)
	}
	return string(tmps)
}
func randUnrecognizedImportduplicate(r randyImportduplicate, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldImportduplicate(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldImportduplicate(dAtA []byte, r randyImportduplicate, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateImportduplicate(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateImportduplicate(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateImportduplicate(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateImportduplicate(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateImportduplicate(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateImportduplicate(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateImportduplicate(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
