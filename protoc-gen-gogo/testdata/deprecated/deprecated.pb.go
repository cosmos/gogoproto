// Code generated by protoc-gen-gogo. DO NOT EDIT.
// deprecated/deprecated.proto is a deprecated file.

// package deprecated contains only deprecated messages and services.

package deprecated

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// DeprecatedEnum contains deprecated values.
type DeprecatedEnum int32 // Deprecated: Do not use.
const (
	// DEPRECATED is the iota value of this enum.
	DeprecatedEnum_DEPRECATED DeprecatedEnum = 0 // Deprecated: Do not use.
)

var DeprecatedEnum_name = map[int32]string{
	0: "DEPRECATED",
}

var DeprecatedEnum_value = map[string]int32{
	"DEPRECATED": 0,
}

func (x DeprecatedEnum) String() string {
	return proto.EnumName(DeprecatedEnum_name, int32(x))
}

func (DeprecatedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f64ba265cd7eae3f, []int{0}
}

// DeprecatedRequest is a request to DeprecatedCall.
//
// Deprecated: Do not use.
type DeprecatedRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeprecatedRequest) Reset()         { *m = DeprecatedRequest{} }
func (m *DeprecatedRequest) String() string { return proto.CompactTextString(m) }
func (*DeprecatedRequest) ProtoMessage()    {}
func (*DeprecatedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f64ba265cd7eae3f, []int{0}
}
func (m *DeprecatedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeprecatedRequest.Unmarshal(m, b)
}
func (m *DeprecatedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeprecatedRequest.Marshal(b, m, deterministic)
}
func (m *DeprecatedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeprecatedRequest.Merge(m, src)
}
func (m *DeprecatedRequest) XXX_Size() int {
	return xxx_messageInfo_DeprecatedRequest.Size(m)
}
func (m *DeprecatedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeprecatedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeprecatedRequest proto.InternalMessageInfo

// Deprecated: Do not use.
type DeprecatedResponse struct {
	// DeprecatedField contains a DeprecatedEnum.
	DeprecatedField DeprecatedEnum `protobuf:"varint,1,opt,name=deprecated_field,json=deprecatedField,proto3,enum=deprecated.DeprecatedEnum" json:"deprecated_field,omitempty"` // Deprecated: Do not use.
	// DeprecatedOneof contains a deprecated field.
	//
	// Types that are valid to be assigned to DeprecatedOneof:
	//	*DeprecatedResponse_DeprecatedOneofField
	DeprecatedOneof      isDeprecatedResponse_DeprecatedOneof `protobuf_oneof:"deprecated_oneof"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *DeprecatedResponse) Reset()         { *m = DeprecatedResponse{} }
func (m *DeprecatedResponse) String() string { return proto.CompactTextString(m) }
func (*DeprecatedResponse) ProtoMessage()    {}
func (*DeprecatedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f64ba265cd7eae3f, []int{1}
}
func (m *DeprecatedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeprecatedResponse.Unmarshal(m, b)
}
func (m *DeprecatedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeprecatedResponse.Marshal(b, m, deterministic)
}
func (m *DeprecatedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeprecatedResponse.Merge(m, src)
}
func (m *DeprecatedResponse) XXX_Size() int {
	return xxx_messageInfo_DeprecatedResponse.Size(m)
}
func (m *DeprecatedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeprecatedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeprecatedResponse proto.InternalMessageInfo

type isDeprecatedResponse_DeprecatedOneof interface {
	isDeprecatedResponse_DeprecatedOneof()
}

type DeprecatedResponse_DeprecatedOneofField struct {
	DeprecatedOneofField string `protobuf:"bytes,2,opt,name=deprecated_oneof_field,json=deprecatedOneofField,proto3,oneof" json:"deprecated_oneof_field,omitempty"`
}

func (*DeprecatedResponse_DeprecatedOneofField) isDeprecatedResponse_DeprecatedOneof() {}

func (m *DeprecatedResponse) GetDeprecatedOneof() isDeprecatedResponse_DeprecatedOneof {
	if m != nil {
		return m.DeprecatedOneof
	}
	return nil
}

// Deprecated: Do not use.
func (m *DeprecatedResponse) GetDeprecatedField() DeprecatedEnum {
	if m != nil {
		return m.DeprecatedField
	}
	return DeprecatedEnum_DEPRECATED
}

// Deprecated: Do not use.
func (m *DeprecatedResponse) GetDeprecatedOneofField() string {
	if x, ok := m.GetDeprecatedOneof().(*DeprecatedResponse_DeprecatedOneofField); ok {
		return x.DeprecatedOneofField
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DeprecatedResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DeprecatedResponse_DeprecatedOneofField)(nil),
	}
}

func init() {
	proto.RegisterEnum("deprecated.DeprecatedEnum", DeprecatedEnum_name, DeprecatedEnum_value)
	proto.RegisterType((*DeprecatedRequest)(nil), "deprecated.DeprecatedRequest")
	proto.RegisterType((*DeprecatedResponse)(nil), "deprecated.DeprecatedResponse")
}

func init() { proto.RegisterFile("deprecated/deprecated.proto", fileDescriptor_f64ba265cd7eae3f) }

var fileDescriptor_f64ba265cd7eae3f = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xfb, 0x40,
	0x10, 0xc7, 0x3b, 0xfb, 0x83, 0x1f, 0xba, 0x87, 0x5a, 0x17, 0xd1, 0x10, 0x51, 0x4a, 0x4e, 0x41,
	0x68, 0x02, 0xf5, 0xd6, 0x8b, 0x98, 0x26, 0xa2, 0x27, 0x25, 0xf6, 0xe4, 0x45, 0xd2, 0x64, 0x1a,
	0x03, 0x49, 0x36, 0x66, 0x37, 0x3e, 0x83, 0xef, 0xe3, 0xc5, 0xc7, 0x93, 0x5d, 0x8b, 0xbb, 0x05,
	0xbd, 0x0c, 0xbb, 0x33, 0xdf, 0xcf, 0xfc, 0xa5, 0xa7, 0x05, 0x76, 0x3d, 0xe6, 0x99, 0xc4, 0x22,
	0x34, 0xcf, 0xa0, 0xeb, 0xb9, 0xe4, 0x8c, 0x1a, 0x8f, 0x77, 0x42, 0x0f, 0xe3, 0x9f, 0x5f, 0x8a,
	0xaf, 0x03, 0x0a, 0xb9, 0x20, 0x0e, 0x78, 0x1f, 0x40, 0x99, 0x1d, 0x11, 0x1d, 0x6f, 0x05, 0xb2,
	0x3b, 0x3a, 0x31, 0xf4, 0xf3, 0xa6, 0xc2, 0xba, 0x70, 0x60, 0x0a, 0xfe, 0x78, 0xee, 0x06, 0x56,
	0x21, 0x43, 0x26, 0xed, 0xd0, 0x44, 0xc4, 0x81, 0xf4, 0xc0, 0x84, 0x6f, 0x14, 0xc6, 0x16, 0xf4,
	0xd8, 0x4a, 0xc5, 0x5b, 0xe4, 0x9b, 0x6d, 0x42, 0x32, 0x05, 0x7f, 0x5f, 0x41, 0xb7, 0xa3, 0xf4,
	0xc8, 0x68, 0xee, 0x95, 0x44, 0xb3, 0xaa, 0xc3, 0x88, 0xed, 0xb4, 0xa2, 0xf9, 0x0b, 0x9f, 0x8e,
	0x77, 0x4b, 0x33, 0x46, 0x69, 0x9c, 0x3c, 0xa4, 0xc9, 0xf2, 0x7a, 0x95, 0xc4, 0x93, 0x91, 0x4b,
	0xf6, 0xc0, 0x25, 0x0e, 0xcc, 0x5b, 0x7b, 0xf0, 0x47, 0xec, 0xdf, 0xaa, 0x1c, 0xd9, 0xca, 0xc6,
	0x97, 0x59, 0x5d, 0xb3, 0xb3, 0xdf, 0xa7, 0xda, 0x6e, 0xca, 0x3d, 0xff, 0x2b, 0xfc, 0xbd, 0x2e,
	0xef, 0xdf, 0x3b, 0x01, 0x57, 0x99, 0x28, 0x7e, 0xba, 0x2a, 0x2b, 0xf9, 0x32, 0xac, 0x83, 0x9c,
	0x37, 0x61, 0xce, 0x45, 0xc3, 0x45, 0x58, 0xf2, 0x92, 0xeb, 0x9b, 0x84, 0xda, 0xe6, 0xb3, 0x12,
	0xdb, 0x99, 0x72, 0x86, 0x12, 0x85, 0x2c, 0x32, 0x99, 0x59, 0xd7, 0xfb, 0x04, 0x58, 0xff, 0xd7,
	0xba, 0xcb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x60, 0x6e, 0x9d, 0xe0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeprecatedServiceClient is the client API for DeprecatedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Do not use.
type DeprecatedServiceClient interface {
	// DeprecatedCall takes a DeprecatedRequest and returns a DeprecatedResponse.
	DeprecatedCall(ctx context.Context, in *DeprecatedRequest, opts ...grpc.CallOption) (*DeprecatedResponse, error)
}

type deprecatedServiceClient struct {
	cc grpc1.ClientConn
}

// Deprecated: Do not use.
func NewDeprecatedServiceClient(cc grpc1.ClientConn) DeprecatedServiceClient {
	return &deprecatedServiceClient{cc}
}

// Deprecated: Do not use.
func (c *deprecatedServiceClient) DeprecatedCall(ctx context.Context, in *DeprecatedRequest, opts ...grpc.CallOption) (*DeprecatedResponse, error) {
	out := new(DeprecatedResponse)
	err := c.cc.Invoke(ctx, "/deprecated.DeprecatedService/DeprecatedCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeprecatedServiceServer is the server API for DeprecatedService service.
//
// Deprecated: Do not use.
type DeprecatedServiceServer interface {
	// DeprecatedCall takes a DeprecatedRequest and returns a DeprecatedResponse.
	DeprecatedCall(context.Context, *DeprecatedRequest) (*DeprecatedResponse, error)
}

// Deprecated: Do not use.
// UnimplementedDeprecatedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeprecatedServiceServer struct {
}

func (*UnimplementedDeprecatedServiceServer) DeprecatedCall(ctx context.Context, req *DeprecatedRequest) (*DeprecatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeprecatedCall not implemented")
}

// Deprecated: Do not use.
func RegisterDeprecatedServiceServer(s grpc1.Server, srv DeprecatedServiceServer) {
	s.RegisterService(&_DeprecatedService_serviceDesc, srv)
}

func _DeprecatedService_DeprecatedCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeprecatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeprecatedServiceServer).DeprecatedCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deprecated.DeprecatedService/DeprecatedCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeprecatedServiceServer).DeprecatedCall(ctx, req.(*DeprecatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeprecatedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deprecated.DeprecatedService",
	HandlerType: (*DeprecatedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeprecatedCall",
			Handler:    _DeprecatedService_DeprecatedCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deprecated/deprecated.proto",
}
