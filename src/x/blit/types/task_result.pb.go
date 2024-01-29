// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blit/blit/task_result.proto

package types

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

type TaskResult struct {
	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ExecutedOn string `protobuf:"bytes,2,opt,name=executed_on,json=executedOn,proto3" json:"executed_on,omitempty"`
	Creator    string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *TaskResult) Reset()         { *m = TaskResult{} }
func (m *TaskResult) String() string { return proto.CompactTextString(m) }
func (*TaskResult) ProtoMessage()    {}
func (*TaskResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_43bf6f182762327c, []int{0}
}
func (m *TaskResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResult.Merge(m, src)
}
func (m *TaskResult) XXX_Size() int {
	return m.Size()
}
func (m *TaskResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResult.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResult proto.InternalMessageInfo

func (m *TaskResult) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskResult) GetExecutedOn() string {
	if m != nil {
		return m.ExecutedOn
	}
	return ""
}

func (m *TaskResult) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*TaskResult)(nil), "blit.blit.TaskResult")
}

func init() { proto.RegisterFile("blit/blit/task_result.proto", fileDescriptor_43bf6f182762327c) }

var fileDescriptor_43bf6f182762327c = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xca, 0xc9, 0x2c,
	0xd1, 0x07, 0x13, 0x25, 0x89, 0xc5, 0xd9, 0xf1, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0x9c, 0x20, 0x71, 0x3d, 0x10, 0xa1, 0x14, 0xce, 0xc5, 0x15, 0x92,
	0x58, 0x9c, 0x1d, 0x04, 0x96, 0x16, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x09, 0x62, 0xca, 0x4c, 0x11, 0x92, 0xe7, 0xe2, 0x4e, 0xad, 0x48, 0x4d, 0x2e, 0x2d, 0x49,
	0x4d, 0x89, 0xcf, 0xcf, 0x93, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x82, 0x09, 0xf9, 0xe7,
	0x09, 0x49, 0x70, 0xb1, 0x27, 0x17, 0xa5, 0x26, 0x96, 0xe4, 0x17, 0x49, 0x30, 0x83, 0x25, 0x61,
	0x5c, 0x27, 0xed, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x12, 0x04, 0xbb,
	0xaa, 0x02, 0xea, 0xb8, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xbb, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x55, 0xd8, 0x0b, 0x9b, 0xb6, 0x00, 0x00, 0x00,
}

func (m *TaskResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTaskResult(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ExecutedOn) > 0 {
		i -= len(m.ExecutedOn)
		copy(dAtA[i:], m.ExecutedOn)
		i = encodeVarintTaskResult(dAtA, i, uint64(len(m.ExecutedOn)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTaskResult(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTaskResult(dAtA []byte, offset int, v uint64) int {
	offset -= sovTaskResult(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TaskResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTaskResult(uint64(m.Id))
	}
	l = len(m.ExecutedOn)
	if l > 0 {
		n += 1 + l + sovTaskResult(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTaskResult(uint64(l))
	}
	return n
}

func sovTaskResult(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTaskResult(x uint64) (n int) {
	return sovTaskResult(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TaskResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTaskResult
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
			return fmt.Errorf("proto: TaskResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutedOn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskResult
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
				return ErrInvalidLengthTaskResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaskResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutedOn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskResult
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
				return ErrInvalidLengthTaskResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaskResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTaskResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTaskResult
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
func skipTaskResult(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTaskResult
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
					return 0, ErrIntOverflowTaskResult
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
					return 0, ErrIntOverflowTaskResult
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
				return 0, ErrInvalidLengthTaskResult
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTaskResult
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTaskResult
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTaskResult        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTaskResult          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTaskResult = fmt.Errorf("proto: unexpected end of group")
)
