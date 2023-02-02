// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enumprefix/enumprefix.proto

package enumprefix

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	test "github.com/cosmos/gogoproto/test"
	math "math"
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

type MyMessage struct {
	TheField             test.TheTestEnum `protobuf:"varint,1,opt,name=TheField,enum=test.TheTestEnum" json:"TheField"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MyMessage) Reset()         { *m = MyMessage{} }
func (m *MyMessage) String() string { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()    {}
func (*MyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_30184dad340c663b, []int{0}
}
func (m *MyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyMessage.Unmarshal(m, b)
}
func (m *MyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyMessage.Marshal(b, m, deterministic)
}
func (m *MyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyMessage.Merge(m, src)
}
func (m *MyMessage) XXX_Size() int {
	return xxx_messageInfo_MyMessage.Size(m)
}
func (m *MyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MyMessage proto.InternalMessageInfo

func (m *MyMessage) GetTheField() test.TheTestEnum {
	if m != nil {
		return m.TheField
	}
	return test.A
}

func init() {
	proto.RegisterType((*MyMessage)(nil), "enumprefix.MyMessage")
}

func init() { proto.RegisterFile("enumprefix/enumprefix.proto", fileDescriptor_30184dad340c663b) }

var fileDescriptor_30184dad340c663b = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xcd, 0x2b, 0xcd,
	0x2d, 0x28, 0x4a, 0x4d, 0xcb, 0xac, 0xd0, 0x47, 0x30, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xb8, 0x10, 0x22, 0x52, 0x42, 0x25, 0xa9, 0xc5, 0x25, 0xfa, 0x25, 0x19, 0xa9, 0x20, 0x1a, 0x22,
	0x2f, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x66, 0xea, 0x83, 0x58, 0x10, 0x51, 0x25, 0x07, 0x2e,
	0x4e, 0xdf, 0x4a, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x21, 0x63, 0x2e, 0x8e, 0x90, 0x8c,
	0x54, 0xb7, 0xcc, 0xd4, 0x9c, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x41, 0x3d, 0xb0,
	0x09, 0x21, 0x19, 0xa9, 0x21, 0xa9, 0xc5, 0x25, 0xae, 0x79, 0xa5, 0xb9, 0x4e, 0x2c, 0x27, 0xee,
	0xc9, 0x33, 0x04, 0xc1, 0x15, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0x45, 0xf7, 0x91, 0xd7, 0x95,
	0x00, 0x00, 0x00,
}
