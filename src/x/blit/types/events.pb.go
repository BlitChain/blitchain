// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blit/blit/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/x/group"
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

// EventCreateTask is an event emitted when a task is created.
type EventCreateTask struct {
	// task_id is the unique ID of the task.
	TaskId uint64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (m *EventCreateTask) Reset()         { *m = EventCreateTask{} }
func (m *EventCreateTask) String() string { return proto.CompactTextString(m) }
func (*EventCreateTask) ProtoMessage()    {}
func (*EventCreateTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_62faec710b0b30a3, []int{0}
}
func (m *EventCreateTask) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateTask.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateTask.Merge(m, src)
}
func (m *EventCreateTask) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateTask) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateTask.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateTask proto.InternalMessageInfo

func (m *EventCreateTask) GetTaskId() uint64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func init() {
	proto.RegisterType((*EventCreateTask)(nil), "blit.blit.EventCreateTask")
}

func init() { proto.RegisterFile("blit/blit/events.proto", fileDescriptor_62faec710b0b30a3) }

var fileDescriptor_62faec710b0b30a3 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xca, 0xc9, 0x2c,
	0xd1, 0x07, 0x13, 0xa9, 0x65, 0xa9, 0x79, 0x25, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x9c, 0x20, 0x21, 0x3d, 0x10, 0x21, 0x25, 0x99, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0x1c, 0x0f, 0x96,
	0xd0, 0x87, 0x70, 0x20, 0xaa, 0xa4, 0xa4, 0x21, 0x3c, 0xfd, 0xf4, 0xa2, 0xfc, 0xd2, 0x02, 0xfd,
	0x32, 0x43, 0xfd, 0x92, 0xca, 0x82, 0x54, 0xa8, 0xa4, 0x92, 0x16, 0x17, 0xbf, 0x2b, 0xc8, 0x48,
	0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0xd4, 0x90, 0xc4, 0xe2, 0x6c, 0x21, 0x71, 0x2e, 0xf6, 0x92, 0xc4,
	0xe2, 0xec, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x36, 0x10, 0xd7, 0x33,
	0xc5, 0x49, 0xfb, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x04, 0xc1, 0x6e,
	0xab, 0x80, 0x38, 0x11, 0x6c, 0x7c, 0x12, 0x1b, 0xd8, 0x7c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x64, 0x0f, 0x4b, 0x09, 0xbc, 0x00, 0x00, 0x00,
}

func (m *EventCreateTask) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateTask) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateTask) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TaskId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.TaskId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreateTask) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TaskId != 0 {
		n += 1 + sovEvents(uint64(m.TaskId))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreateTask) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateTask: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateTask: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskId", wireType)
			}
			m.TaskId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TaskId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
