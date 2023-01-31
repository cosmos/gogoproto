// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: indeximport_issue72/indeximport.proto

package indeximport_issue72

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	index "indeximport_issue72/index"
	io "io"
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

type IndexQueries struct {
	Queries              []*index.IndexQuery `protobuf:"bytes,1,rep,name=Queries" json:"Queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IndexQueries) Reset()         { *m = IndexQueries{} }
func (m *IndexQueries) String() string { return proto.CompactTextString(m) }
func (*IndexQueries) ProtoMessage()    {}
func (*IndexQueries) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d4c6a51f17d285, []int{0}
}
func (m *IndexQueries) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexQueries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexQueries.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexQueries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexQueries.Merge(m, src)
}
func (m *IndexQueries) XXX_Size() int {
	return m.Size()
}
func (m *IndexQueries) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexQueries.DiscardUnknown(m)
}

var xxx_messageInfo_IndexQueries proto.InternalMessageInfo

func (m *IndexQueries) GetQueries() []*index.IndexQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

func init() {
	proto.RegisterType((*IndexQueries)(nil), "indeximport_issue72.IndexQueries")
}

func init() {
	proto.RegisterFile("indeximport_issue72/indeximport.proto", fileDescriptor_07d4c6a51f17d285)
}

var fileDescriptor_07d4c6a51f17d285 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0xcc, 0x4b, 0x49,
	0xad, 0xc8, 0xcc, 0x2d, 0xc8, 0x2f, 0x2a, 0x89, 0xcf, 0x2c, 0x2e, 0x2e, 0x4d, 0x35, 0x37, 0xd2,
	0x47, 0x12, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc6, 0xa2, 0x4c, 0x0a, 0xb7, 0x5e,
	0x08, 0x09, 0xd1, 0x2b, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x66, 0xea, 0x83, 0x58, 0x10, 0x51,
	0x25, 0x7f, 0x2e, 0x1e, 0x4f, 0x90, 0xa2, 0xc0, 0xd2, 0xd4, 0xa2, 0xcc, 0xd4, 0x62, 0x21, 0x7b,
	0x2e, 0x76, 0x28, 0x53, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x55, 0x0f, 0x8b, 0xf1, 0x10,
	0x31, 0x3d, 0xb8, 0xce, 0xca, 0x20, 0x98, 0x2e, 0x27, 0x89, 0x1f, 0x0f, 0xe5, 0x18, 0x57, 0x3c,
	0x92, 0x63, 0xdc, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xba, 0x5e, 0xd7, 0xe4, 0x00, 0x00,
	0x00,
}

func (this *IndexQueries) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IndexQueries)
	if !ok {
		that2, ok := that.(IndexQueries)
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
	if len(this.Queries) != len(that1.Queries) {
		return false
	}
	for i := range this.Queries {
		if !this.Queries[i].Equal(that1.Queries[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *IndexQueries) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexQueries) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexQueries) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Queries) > 0 {
		for iNdEx := len(m.Queries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Queries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIndeximport(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintIndeximport(dAtA []byte, offset int, v uint64) int {
	offset -= sovIndeximport(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedIndexQueries(r randyIndeximport, easy bool) *IndexQueries {
	this := &IndexQueries{}
	if r.Intn(5) != 0 {
		v1 := r.Intn(5)
		this.Queries = make([]*index.IndexQuery, v1)
		for i := 0; i < v1; i++ {
			this.Queries[i] = index.NewPopulatedIndexQuery(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedIndeximport(r, 2)
	}
	return this
}

type randyIndeximport interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneIndeximport(r randyIndeximport) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringIndeximport(r randyIndeximport) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneIndeximport(r)
	}
	return string(tmps)
}
func randUnrecognizedIndeximport(r randyIndeximport, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldIndeximport(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldIndeximport(dAtA []byte, r randyIndeximport, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateIndeximport(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateIndeximport(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateIndeximport(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateIndeximport(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateIndeximport(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateIndeximport(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateIndeximport(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *IndexQueries) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, e := range m.Queries {
			l = e.Size()
			n += 1 + l + sovIndeximport(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIndeximport(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIndeximport(x uint64) (n int) {
	return sovIndeximport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexQueries) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndeximport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IndexQueries: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexQueries: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Queries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndeximport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIndeximport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndeximport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Queries = append(m.Queries, &index.IndexQuery{})
			if err := m.Queries[len(m.Queries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIndeximport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIndeximport
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIndeximport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIndeximport
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIndeximport
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIndeximport
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthIndeximport
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIndeximport
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIndeximport
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIndeximport        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIndeximport          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIndeximport = fmt.Errorf("proto: unexpected end of group")
)
