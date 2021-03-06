// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue438.proto

package issue438

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/gogoproto/types"
	math "math"
	math_bits "math/bits"
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

type Types struct {
	Any                  *types.Any           `protobuf:"bytes,1,opt,name=any,proto3" json:"any,omitempty"`
	Api                  *types.Api           `protobuf:"bytes,2,opt,name=api,proto3" json:"api,omitempty"`
	Met                  *types.Method        `protobuf:"bytes,3,opt,name=met,proto3" json:"met,omitempty"`
	Mx                   *types.Mixin         `protobuf:"bytes,4,opt,name=mx,proto3" json:"mx,omitempty"`
	Dur                  *types.Duration      `protobuf:"bytes,5,opt,name=dur,proto3" json:"dur,omitempty"`
	Em                   *types.Empty         `protobuf:"bytes,6,opt,name=em,proto3" json:"em,omitempty"`
	Fm                   *types.FieldMask     `protobuf:"bytes,7,opt,name=fm,proto3" json:"fm,omitempty"`
	Sc                   *types.SourceContext `protobuf:"bytes,8,opt,name=sc,proto3" json:"sc,omitempty"`
	St                   *types.Struct        `protobuf:"bytes,9,opt,name=st,proto3" json:"st,omitempty"`
	Val                  *types.Value         `protobuf:"bytes,10,opt,name=val,proto3" json:"val,omitempty"`
	Nlval                types.NullValue      `protobuf:"varint,11,opt,name=nlval,proto3,enum=google.protobuf.NullValue" json:"nlval,omitempty"`
	Stval                *types.StringValue   `protobuf:"bytes,12,opt,name=stval,proto3" json:"stval,omitempty"`
	Bval                 *types.BoolValue     `protobuf:"bytes,13,opt,name=bval,proto3" json:"bval,omitempty"`
	Strval               *types.Struct        `protobuf:"bytes,14,opt,name=strval,proto3" json:"strval,omitempty"`
	Lstv                 *types.ListValue     `protobuf:"bytes,15,opt,name=lstv,proto3" json:"lstv,omitempty"`
	Ts                   *types.Timestamp     `protobuf:"bytes,16,opt,name=ts,proto3" json:"ts,omitempty"`
	T                    *types.Type          `protobuf:"bytes,17,opt,name=t,proto3" json:"t,omitempty"`
	F                    *types.Field         `protobuf:"bytes,18,opt,name=f,proto3" json:"f,omitempty"`
	En                   *types.Enum          `protobuf:"bytes,19,opt,name=en,proto3" json:"en,omitempty"`
	Enval                *types.EnumValue     `protobuf:"bytes,20,opt,name=enval,proto3" json:"enval,omitempty"`
	Opt                  *types.Option        `protobuf:"bytes,21,opt,name=opt,proto3" json:"opt,omitempty"`
	Dbl                  *types.DoubleValue   `protobuf:"bytes,22,opt,name=dbl,proto3" json:"dbl,omitempty"`
	Flt                  *types.FloatValue    `protobuf:"bytes,23,opt,name=flt,proto3" json:"flt,omitempty"`
	I64                  *types.Int64Value    `protobuf:"bytes,24,opt,name=i64,proto3" json:"i64,omitempty"`
	U64                  *types.UInt64Value   `protobuf:"bytes,25,opt,name=u64,proto3" json:"u64,omitempty"`
	I32                  *types.Int32Value    `protobuf:"bytes,26,opt,name=i32,proto3" json:"i32,omitempty"`
	U32                  *types.UInt32Value   `protobuf:"bytes,27,opt,name=u32,proto3" json:"u32,omitempty"`
	Bool                 *types.BoolValue     `protobuf:"bytes,28,opt,name=bool,proto3" json:"bool,omitempty"`
	Str                  *types.StringValue   `protobuf:"bytes,29,opt,name=str,proto3" json:"str,omitempty"`
	Bytes                *types.BytesValue    `protobuf:"bytes,30,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Types) Reset()         { *m = Types{} }
func (m *Types) String() string { return proto.CompactTextString(m) }
func (*Types) ProtoMessage()    {}
func (*Types) Descriptor() ([]byte, []int) {
	return fileDescriptor_43147f0c8dedbac4, []int{0}
}
func (m *Types) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Types.Unmarshal(m, b)
}
func (m *Types) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Types.Marshal(b, m, deterministic)
}
func (m *Types) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Types.Merge(m, src)
}
func (m *Types) XXX_Size() int {
	return xxx_messageInfo_Types.Size(m)
}
func (m *Types) XXX_DiscardUnknown() {
	xxx_messageInfo_Types.DiscardUnknown(m)
}

var xxx_messageInfo_Types proto.InternalMessageInfo

func (m *Types) GetAny() *types.Any {
	if m != nil {
		return m.Any
	}
	return nil
}

func (m *Types) GetApi() *types.Api {
	if m != nil {
		return m.Api
	}
	return nil
}

func (m *Types) GetMet() *types.Method {
	if m != nil {
		return m.Met
	}
	return nil
}

func (m *Types) GetMx() *types.Mixin {
	if m != nil {
		return m.Mx
	}
	return nil
}

func (m *Types) GetDur() *types.Duration {
	if m != nil {
		return m.Dur
	}
	return nil
}

func (m *Types) GetEm() *types.Empty {
	if m != nil {
		return m.Em
	}
	return nil
}

func (m *Types) GetFm() *types.FieldMask {
	if m != nil {
		return m.Fm
	}
	return nil
}

func (m *Types) GetSc() *types.SourceContext {
	if m != nil {
		return m.Sc
	}
	return nil
}

func (m *Types) GetSt() *types.Struct {
	if m != nil {
		return m.St
	}
	return nil
}

func (m *Types) GetVal() *types.Value {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *Types) GetNlval() types.NullValue {
	if m != nil {
		return m.Nlval
	}
	return types.NullValue_NULL_VALUE
}

func (m *Types) GetStval() *types.StringValue {
	if m != nil {
		return m.Stval
	}
	return nil
}

func (m *Types) GetBval() *types.BoolValue {
	if m != nil {
		return m.Bval
	}
	return nil
}

func (m *Types) GetStrval() *types.Struct {
	if m != nil {
		return m.Strval
	}
	return nil
}

func (m *Types) GetLstv() *types.ListValue {
	if m != nil {
		return m.Lstv
	}
	return nil
}

func (m *Types) GetTs() *types.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *Types) GetT() *types.Type {
	if m != nil {
		return m.T
	}
	return nil
}

func (m *Types) GetF() *types.Field {
	if m != nil {
		return m.F
	}
	return nil
}

func (m *Types) GetEn() *types.Enum {
	if m != nil {
		return m.En
	}
	return nil
}

func (m *Types) GetEnval() *types.EnumValue {
	if m != nil {
		return m.Enval
	}
	return nil
}

func (m *Types) GetOpt() *types.Option {
	if m != nil {
		return m.Opt
	}
	return nil
}

func (m *Types) GetDbl() *types.DoubleValue {
	if m != nil {
		return m.Dbl
	}
	return nil
}

func (m *Types) GetFlt() *types.FloatValue {
	if m != nil {
		return m.Flt
	}
	return nil
}

func (m *Types) GetI64() *types.Int64Value {
	if m != nil {
		return m.I64
	}
	return nil
}

func (m *Types) GetU64() *types.UInt64Value {
	if m != nil {
		return m.U64
	}
	return nil
}

func (m *Types) GetI32() *types.Int32Value {
	if m != nil {
		return m.I32
	}
	return nil
}

func (m *Types) GetU32() *types.UInt32Value {
	if m != nil {
		return m.U32
	}
	return nil
}

func (m *Types) GetBool() *types.BoolValue {
	if m != nil {
		return m.Bool
	}
	return nil
}

func (m *Types) GetStr() *types.StringValue {
	if m != nil {
		return m.Str
	}
	return nil
}

func (m *Types) GetBytes() *types.BytesValue {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func init() {
	proto.RegisterType((*Types)(nil), "issue438.Types")
}

func init() { proto.RegisterFile("issue438.proto", fileDescriptor_43147f0c8dedbac4) }

var fileDescriptor_43147f0c8dedbac4 = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd3, 0xef, 0x6e, 0x13, 0x39,
	0x10, 0x00, 0xf0, 0xdb, 0xdd, 0xa6, 0xd7, 0x73, 0xef, 0x7a, 0x60, 0xfa, 0xc7, 0x4d, 0x4b, 0xa8,
	0xa0, 0x94, 0x02, 0x62, 0x53, 0x92, 0x28, 0xe2, 0x2b, 0xa5, 0x45, 0x42, 0xa2, 0x20, 0xb9, 0x85,
	0xaf, 0xd5, 0x6e, 0xe2, 0xa4, 0x56, 0xed, 0xb5, 0xb5, 0xb6, 0x4b, 0xf2, 0x36, 0x3c, 0x12, 0x8f,
	0x80, 0xe0, 0x45, 0x90, 0xbd, 0x0e, 0x42, 0xd9, 0xba, 0xea, 0xb7, 0x44, 0xf3, 0x9b, 0xf1, 0x78,
	0xc6, 0x0b, 0x56, 0xa8, 0x52, 0x86, 0xf4, 0xba, 0xaf, 0x52, 0x59, 0x0a, 0x2d, 0xe0, 0xd2, 0xec,
	0x7f, 0xf3, 0x60, 0x4c, 0xf5, 0x85, 0xc9, 0xd3, 0x81, 0xe0, 0xed, 0x81, 0x50, 0x5c, 0xa8, 0xf6,
	0x58, 0x8c, 0x85, 0x63, 0x73, 0xbf, 0xaa, 0xdc, 0xe6, 0xe6, 0x58, 0x88, 0x31, 0x23, 0x6d, 0xf7,
	0x2f, 0x37, 0xa3, 0x76, 0x56, 0x4c, 0x83, 0x21, 0x49, 0x7d, 0xa8, 0x35, 0x1f, 0x1a, 0x9a, 0x32,
	0xd3, 0x54, 0x14, 0x3e, 0xbe, 0x35, 0x1f, 0x27, 0x5c, 0xea, 0x59, 0xdd, 0x9d, 0xf9, 0xe0, 0x88,
	0x12, 0x36, 0x3c, 0xe7, 0x99, 0xba, 0xf4, 0x62, 0x77, 0x5e, 0x28, 0x61, 0xca, 0x01, 0x39, 0x1f,
	0x88, 0x42, 0x93, 0x89, 0xf6, 0x6a, 0xbb, 0xa6, 0x74, 0x69, 0x06, 0xb3, 0xe8, 0x83, 0xf9, 0xa8,
	0xa6, 0x9c, 0x28, 0x9d, 0x71, 0xe9, 0x41, 0xb3, 0x06, 0xa6, 0x92, 0x84, 0xee, 0xf7, 0xa5, 0xcc,
	0xa4, 0x24, 0xa5, 0xaa, 0xe2, 0x0f, 0xbf, 0x01, 0xd0, 0x38, 0x9b, 0x4a, 0xa2, 0xe0, 0x1e, 0x48,
	0xb2, 0x62, 0x8a, 0xa2, 0x9d, 0x68, 0x7f, 0xb9, 0xb3, 0x9a, 0x56, 0x79, 0xe9, 0x2c, 0x2f, 0x7d,
	0x5d, 0x4c, 0xb1, 0x05, 0xce, 0x49, 0x8a, 0xe2, 0x90, 0x93, 0x14, 0x5b, 0x00, 0x9f, 0x82, 0x84,
	0x13, 0x8d, 0x12, 0xe7, 0x36, 0x6a, 0xee, 0x84, 0xe8, 0x0b, 0x31, 0xc4, 0xd6, 0xc0, 0x3d, 0x10,
	0xf3, 0x09, 0x5a, 0x70, 0x72, 0xbd, 0x2e, 0xe9, 0x84, 0x16, 0x38, 0xe6, 0x13, 0xf8, 0x1c, 0x24,
	0x43, 0x53, 0xa2, 0x86, 0x83, 0x9b, 0x35, 0x78, 0xe4, 0x57, 0x87, 0xad, 0xb2, 0x45, 0x09, 0x47,
	0x8b, 0x81, 0xa2, 0xc7, 0x76, 0x8d, 0x38, 0x26, 0x1c, 0x3e, 0x03, 0xf1, 0x88, 0xa3, 0xbf, 0x9d,
	0x6b, 0xd6, 0xdc, 0x5b, 0xbb, 0xd1, 0x93, 0x4c, 0x5d, 0xe2, 0x78, 0xc4, 0x61, 0x0a, 0x62, 0x35,
	0x40, 0x4b, 0xce, 0xb6, 0x6a, 0xf6, 0xd4, 0xed, 0xf6, 0x4d, 0xb5, 0x5a, 0x1c, 0xab, 0x01, 0x7c,
	0x02, 0x62, 0xa5, 0xd1, 0x3f, 0x81, 0x11, 0x9c, 0xba, 0x2d, 0xe3, 0x58, 0x69, 0xb8, 0x0f, 0x92,
	0xab, 0x8c, 0x21, 0x10, 0xe8, 0xf6, 0x73, 0xc6, 0x0c, 0xc1, 0x96, 0xc0, 0x03, 0xd0, 0x28, 0x98,
	0xb5, 0xcb, 0x3b, 0xd1, 0xfe, 0xca, 0x35, 0x1d, 0x7f, 0x30, 0x8c, 0x55, 0xbe, 0x82, 0xb0, 0x03,
	0x1a, 0x4a, 0xdb, 0x8c, 0x7f, 0x5d, 0xf5, 0xed, 0xeb, 0xfa, 0xa0, 0xc5, 0xd8, 0xe7, 0x38, 0x0a,
	0x53, 0xb0, 0x90, 0xdb, 0x94, 0xff, 0x02, 0x63, 0x39, 0x14, 0xc2, 0x1f, 0xe2, 0x1c, 0x6c, 0x83,
	0x45, 0xa5, 0x4b, 0x9b, 0xb1, 0x72, 0xf3, 0x65, 0x3d, 0xb3, 0x07, 0x30, 0xa5, 0xaf, 0xd0, 0xff,
	0x81, 0x03, 0xde, 0x53, 0xa5, 0xfd, 0x01, 0xd6, 0xd9, 0x2d, 0x69, 0x85, 0xee, 0x04, 0xf4, 0xd9,
	0xec, 0x8b, 0xc0, 0xb1, 0x56, 0xf0, 0x11, 0x88, 0x34, 0xba, 0xeb, 0xe8, 0x5a, 0x9d, 0x4e, 0x25,
	0xc1, 0x91, 0x86, 0xbb, 0x20, 0x1a, 0x21, 0x18, 0x98, 0xb7, 0xdb, 0x3a, 0x8e, 0x46, 0xf0, 0x31,
	0x88, 0x49, 0x81, 0xee, 0x05, 0x6a, 0x1d, 0x17, 0x86, 0xe3, 0x98, 0x14, 0x76, 0x29, 0xa4, 0xb0,
	0xb7, 0x5f, 0x0d, 0x34, 0x68, 0xa5, 0x1f, 0xb0, 0x83, 0xf6, 0xeb, 0x10, 0x52, 0xa3, 0xb5, 0xc0,
	0xb4, 0x3e, 0xca, 0xea, 0x21, 0x0b, 0xa9, 0x61, 0x0a, 0x92, 0x61, 0xce, 0xd0, 0x7a, 0x60, 0x7b,
	0x47, 0xc2, 0xe4, 0x8c, 0xf8, 0x17, 0x32, 0xcc, 0x19, 0x7c, 0x01, 0x92, 0x11, 0xd3, 0x68, 0xc3,
	0xf9, 0xad, 0xfa, 0xdd, 0x98, 0xc8, 0xfc, 0x68, 0xad, 0xb3, 0x9c, 0xf6, 0x7b, 0x08, 0x05, 0xf8,
	0xbb, 0x42, 0xf7, 0x7b, 0x9e, 0xd3, 0x7e, 0xcf, 0x76, 0x63, 0xfa, 0x3d, 0xb4, 0x19, 0xe8, 0xe6,
	0xd3, 0x9f, 0xde, 0xf4, 0x7b, 0xae, 0x7c, 0xb7, 0x83, 0x9a, 0xe1, 0xf2, 0xdd, 0xce, 0xac, 0x7c,
	0xb7, 0xe3, 0xca, 0x77, 0x3b, 0x68, 0xeb, 0x86, 0xf2, 0xbf, 0xbd, 0x71, 0x7e, 0x21, 0x17, 0x82,
	0xa1, 0xed, 0x5b, 0x3c, 0x54, 0x21, 0xec, 0xbb, 0x4b, 0x94, 0x2e, 0xd1, 0xfd, 0x5b, 0x7c, 0x0a,
	0x16, 0xc2, 0x97, 0xa0, 0x91, 0x4f, 0x35, 0x51, 0xa8, 0x15, 0xb8, 0xc0, 0xa1, 0x8d, 0xfa, 0xd5,
	0x3a, 0x79, 0xb8, 0xf4, 0xfd, 0x47, 0xeb, 0xaf, 0xaf, 0x3f, 0x5b, 0x51, 0xbe, 0xe8, 0x54, 0xf7,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0xc4, 0x8e, 0xb7, 0xe7, 0x06, 0x00, 0x00,
}

func (m *Types) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Any != nil {
		l = m.Any.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Api != nil {
		l = m.Api.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Met != nil {
		l = m.Met.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Mx != nil {
		l = m.Mx.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Dur != nil {
		l = m.Dur.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Em != nil {
		l = m.Em.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Fm != nil {
		l = m.Fm.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Sc != nil {
		l = m.Sc.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.St != nil {
		l = m.St.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Val != nil {
		l = m.Val.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Nlval != 0 {
		n += 1 + sovIssue438(uint64(m.Nlval))
	}
	if m.Stval != nil {
		l = m.Stval.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Bval != nil {
		l = m.Bval.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Strval != nil {
		l = m.Strval.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Lstv != nil {
		l = m.Lstv.ProtoSize()
		n += 1 + l + sovIssue438(uint64(l))
	}
	if m.Ts != nil {
		l = m.Ts.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.T != nil {
		l = m.T.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.F != nil {
		l = m.F.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.En != nil {
		l = m.En.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Enval != nil {
		l = m.Enval.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Opt != nil {
		l = m.Opt.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Dbl != nil {
		l = m.Dbl.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Flt != nil {
		l = m.Flt.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.I64 != nil {
		l = m.I64.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.U64 != nil {
		l = m.U64.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.I32 != nil {
		l = m.I32.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.U32 != nil {
		l = m.U32.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Bool != nil {
		l = m.Bool.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Str != nil {
		l = m.Str.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.Bytes != nil {
		l = m.Bytes.ProtoSize()
		n += 2 + l + sovIssue438(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIssue438(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIssue438(x uint64) (n int) {
	return sovIssue438(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
