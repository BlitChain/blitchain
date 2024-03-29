// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blit/blit/timespec.proto

package blit_blit

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type TimeSpec struct {
	// Types that are valid to be assigned to TimeType:
	//
	//	*TimeSpec_BlockHeight
	//	*TimeSpec_Timestamp
	TimeType isTimeSpec_TimeType `protobuf_oneof:"time_type"`
}

func (m *TimeSpec) Reset()         { *m = TimeSpec{} }
func (m *TimeSpec) String() string { return proto.CompactTextString(m) }
func (*TimeSpec) ProtoMessage()    {}
func (*TimeSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad611c629229fe2e, []int{0}
}
func (m *TimeSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimeSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimeSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimeSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSpec.Merge(m, src)
}
func (m *TimeSpec) XXX_Size() int {
	return m.Size()
}
func (m *TimeSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSpec proto.InternalMessageInfo

type isTimeSpec_TimeType interface {
	isTimeSpec_TimeType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type TimeSpec_BlockHeight struct {
	BlockHeight int64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3,oneof" json:"block_height,omitempty"`
}
type TimeSpec_Timestamp struct {
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3,oneof" json:"timestamp,omitempty"`
}

func (*TimeSpec_BlockHeight) isTimeSpec_TimeType() {}
func (*TimeSpec_Timestamp) isTimeSpec_TimeType()   {}

func (m *TimeSpec) GetTimeType() isTimeSpec_TimeType {
	if m != nil {
		return m.TimeType
	}
	return nil
}

func (m *TimeSpec) GetBlockHeight() int64 {
	if x, ok := m.GetTimeType().(*TimeSpec_BlockHeight); ok {
		return x.BlockHeight
	}
	return 0
}

func (m *TimeSpec) GetTimestamp() string {
	if x, ok := m.GetTimeType().(*TimeSpec_Timestamp); ok {
		return x.Timestamp
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TimeSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TimeSpec_BlockHeight)(nil),
		(*TimeSpec_Timestamp)(nil),
	}
}

func init() {
	proto.RegisterType((*TimeSpec)(nil), "blit.blit.TimeSpec")
}

func init() { proto.RegisterFile("blit/blit/timespec.proto", fileDescriptor_ad611c629229fe2e) }

var fileDescriptor_ad611c629229fe2e = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xca, 0xc9, 0x2c,
	0xd1, 0x07, 0x13, 0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x05, 0xa9, 0xc9, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x9c, 0x20, 0x41, 0x3d, 0x10, 0xa1, 0x14, 0xc3, 0xc5, 0x11, 0x92, 0x99, 0x9b, 0x1a,
	0x5c, 0x90, 0x9a, 0x2c, 0xa4, 0xcc, 0xc5, 0x93, 0x94, 0x93, 0x9f, 0x9c, 0x1d, 0x9f, 0x91, 0x9a,
	0x99, 0x9e, 0x51, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xec, 0xc1, 0x10, 0xc4, 0x0d, 0x16, 0xf5,
	0x00, 0x0b, 0x0a, 0xc9, 0x71, 0x71, 0x82, 0x4d, 0x2b, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x52, 0x60,
	0xd4, 0xe0, 0xf4, 0x60, 0x08, 0x42, 0x08, 0x39, 0x71, 0x43, 0xe4, 0xe3, 0x4b, 0x2a, 0x0b, 0x52,
	0x9d, 0x24, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0xec, 0x12,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x52, 0x47, 0x0f, 0xa5, 0x00, 0x00, 0x00,
}

func (m *TimeSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimeSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeType != nil {
		{
			size := m.TimeType.Size()
			i -= size
			if _, err := m.TimeType.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *TimeSpec_BlockHeight) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimeSpec_BlockHeight) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintTimespec(dAtA, i, uint64(m.BlockHeight))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}
func (m *TimeSpec_Timestamp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimeSpec_Timestamp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Timestamp)
	copy(dAtA[i:], m.Timestamp)
	i = encodeVarintTimespec(dAtA, i, uint64(len(m.Timestamp)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func encodeVarintTimespec(dAtA []byte, offset int, v uint64) int {
	offset -= sovTimespec(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TimeSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TimeType != nil {
		n += m.TimeType.Size()
	}
	return n
}

func (m *TimeSpec_BlockHeight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovTimespec(uint64(m.BlockHeight))
	return n
}
func (m *TimeSpec_Timestamp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Timestamp)
	n += 1 + l + sovTimespec(uint64(l))
	return n
}

func sovTimespec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTimespec(x uint64) (n int) {
	return sovTimespec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TimeSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimespec
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
			return fmt.Errorf("proto: TimeSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimespec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TimeType = &TimeSpec_BlockHeight{v}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimespec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimespec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTimespec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeType = &TimeSpec_Timestamp{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimespec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTimespec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTimespec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimespec
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
					return 0, ErrIntOverflowTimespec
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
					return 0, ErrIntOverflowTimespec
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
				return 0, ErrInvalidLengthTimespec
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTimespec
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTimespec
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTimespec        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimespec          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTimespec = fmt.Errorf("proto: unexpected end of group")
)
