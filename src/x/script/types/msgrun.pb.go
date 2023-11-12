// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blit/script/msgrun.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/x/authz"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// MsgRun runs a script at a specific address
type MsgRun struct {
	CallerAddress    string `protobuf:"bytes,2,opt,name=caller_address,json=callerAddress,proto3" json:"caller_address,omitempty"`
	ScriptAddress    string `protobuf:"bytes,3,opt,name=script_address,json=scriptAddress,proto3" json:"script_address,omitempty"`
	ExtraCode        string `protobuf:"bytes,4,opt,name=extra_code,json=extraCode,proto3" json:"extra_code,omitempty"`
	FunctionName     string `protobuf:"bytes,5,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	Kwargs           string `protobuf:"bytes,6,opt,name=kwargs,proto3" json:"kwargs,omitempty"`
	Grantee          string `protobuf:"bytes,7,opt,name=grantee,proto3" json:"grantee,omitempty"`
	AttachedMessages string `protobuf:"bytes,1,opt,name=attached_messages,json=attachedMessages,proto3" json:"attached_messages,omitempty"`
}

func (m *MsgRun) Reset()         { *m = MsgRun{} }
func (m *MsgRun) String() string { return proto.CompactTextString(m) }
func (*MsgRun) ProtoMessage()    {}
func (*MsgRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_960b83671d1b01be, []int{0}
}
func (m *MsgRun) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRun.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRun.Merge(m, src)
}
func (m *MsgRun) XXX_Size() int {
	return m.Size()
}
func (m *MsgRun) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRun.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRun proto.InternalMessageInfo

func (m *MsgRun) GetCallerAddress() string {
	if m != nil {
		return m.CallerAddress
	}
	return ""
}

func (m *MsgRun) GetScriptAddress() string {
	if m != nil {
		return m.ScriptAddress
	}
	return ""
}

func (m *MsgRun) GetExtraCode() string {
	if m != nil {
		return m.ExtraCode
	}
	return ""
}

func (m *MsgRun) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *MsgRun) GetKwargs() string {
	if m != nil {
		return m.Kwargs
	}
	return ""
}

func (m *MsgRun) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *MsgRun) GetAttachedMessages() string {
	if m != nil {
		return m.AttachedMessages
	}
	return ""
}

type MsgRunResponse struct {
	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (m *MsgRunResponse) Reset()         { *m = MsgRunResponse{} }
func (m *MsgRunResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRunResponse) ProtoMessage()    {}
func (*MsgRunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_960b83671d1b01be, []int{1}
}
func (m *MsgRunResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRunResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRunResponse.Merge(m, src)
}
func (m *MsgRunResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRunResponse proto.InternalMessageInfo

func (m *MsgRunResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgRun)(nil), "blit.script.MsgRun")
	proto.RegisterType((*MsgRunResponse)(nil), "blit.script.MsgRunResponse")
}

func init() { proto.RegisterFile("blit/script/msgrun.proto", fileDescriptor_960b83671d1b01be) }

var fileDescriptor_960b83671d1b01be = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x91, 0x41, 0x4f, 0xf2, 0x30,
	0x18, 0xc7, 0xd9, 0xfb, 0xea, 0x90, 0x2a, 0x44, 0xa7, 0xd1, 0x42, 0xe2, 0x42, 0x30, 0x26, 0x26,
	0x2a, 0x0d, 0xf1, 0x13, 0xa8, 0x67, 0x3c, 0x70, 0xf4, 0xb2, 0x74, 0xdb, 0x43, 0x21, 0x6e, 0xed,
	0xd2, 0x76, 0x08, 0x7e, 0x0a, 0x3f, 0x96, 0x47, 0x8e, 0x1e, 0x0d, 0x7c, 0x0a, 0x6f, 0x66, 0x6d,
	0xc7, 0x69, 0xfd, 0xff, 0x7f, 0xbf, 0x2e, 0x79, 0xfa, 0x20, 0x1c, 0x67, 0x73, 0x4d, 0x54, 0x22,
	0xe7, 0x85, 0x26, 0xb9, 0x62, 0xb2, 0xe4, 0xc3, 0x42, 0x0a, 0x2d, 0x82, 0xc3, 0x8a, 0x0c, 0x2d,
	0xe9, 0x75, 0x99, 0x10, 0x2c, 0x03, 0x62, 0x50, 0x5c, 0x4e, 0x09, 0xe5, 0x2b, 0xeb, 0xf5, 0xba,
	0x89, 0x50, 0xb9, 0x50, 0x91, 0x49, 0xc4, 0x06, 0x87, 0x2e, 0x6c, 0xaa, 0xfe, 0x4b, 0x16, 0xa3,
	0xea, 0xe3, 0xc0, 0x19, 0x13, 0x4c, 0xd8, 0x0b, 0xd5, 0xc9, 0xb5, 0x7d, 0xa7, 0xd3, 0x52, 0xcf,
	0x3e, 0xc8, 0x62, 0x14, 0x83, 0xa6, 0x23, 0x9b, 0xac, 0x31, 0xf8, 0xf5, 0x90, 0x3f, 0x56, 0x6c,
	0x52, 0xf2, 0xe0, 0x1a, 0x75, 0x12, 0x9a, 0x65, 0x20, 0x23, 0x9a, 0xa6, 0x12, 0x94, 0xc2, 0xff,
	0xfa, 0xde, 0x4d, 0x6b, 0xd2, 0xb6, 0xed, 0xa3, 0x2d, 0x2b, 0xcd, 0x8e, 0xb0, 0xd3, 0xfe, 0x5b,
	0xcd, 0xb6, 0xb5, 0x76, 0x89, 0x10, 0x2c, 0xb5, 0xa4, 0x51, 0x22, 0x52, 0xc0, 0x7b, 0x46, 0x69,
	0x99, 0xe6, 0x59, 0xa4, 0x10, 0x5c, 0xa1, 0xf6, 0xb4, 0xe4, 0x89, 0x9e, 0x0b, 0x1e, 0x71, 0x9a,
	0x03, 0xde, 0x37, 0xc6, 0x51, 0x5d, 0xbe, 0xd0, 0x1c, 0x82, 0x73, 0xe4, 0xbf, 0xbd, 0x53, 0xc9,
	0x14, 0xf6, 0x0d, 0x75, 0x29, 0xc0, 0xa8, 0xc9, 0x24, 0xe5, 0x1a, 0x00, 0x37, 0x0d, 0xa8, 0x63,
	0x70, 0x8b, 0x4e, 0xa8, 0xd6, 0x34, 0x99, 0x41, 0x1a, 0xe5, 0xa0, 0x14, 0x65, 0xa0, 0xb0, 0x67,
	0x9c, 0xe3, 0x1a, 0x8c, 0x5d, 0x3f, 0xb8, 0x43, 0x1d, 0x3b, 0xfa, 0x04, 0x54, 0x21, 0xb8, 0x82,
	0xa0, 0x87, 0x0e, 0xa4, 0x3b, 0xbb, 0x5b, 0xbb, 0xfc, 0x74, 0xff, 0xb5, 0x09, 0xbd, 0xf5, 0x26,
	0xf4, 0x7e, 0x36, 0xa1, 0xf7, 0xb9, 0x0d, 0x1b, 0xeb, 0x6d, 0xd8, 0xf8, 0xde, 0x86, 0x8d, 0xd7,
	0x53, 0xb3, 0xf1, 0x65, 0xbd, 0x73, 0xbd, 0x2a, 0x40, 0xc5, 0xbe, 0x79, 0xdf, 0x87, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x32, 0x58, 0x3a, 0xfe, 0x0f, 0x02, 0x00, 0x00,
}

func (m *MsgRun) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRun) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRun) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintMsgrun(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Kwargs) > 0 {
		i -= len(m.Kwargs)
		copy(dAtA[i:], m.Kwargs)
		i = encodeVarintMsgrun(dAtA, i, uint64(len(m.Kwargs)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.FunctionName) > 0 {
		i -= len(m.FunctionName)
		copy(dAtA[i:], m.FunctionName)
		i = encodeVarintMsgrun(dAtA, i, uint64(len(m.FunctionName)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ExtraCode) > 0 {
		i -= len(m.ExtraCode)
		copy(dAtA[i:], m.ExtraCode)
		i = encodeVarintMsgrun(dAtA, i, uint64(len(m.ExtraCode)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ScriptAddress) > 0 {
		i -= len(m.ScriptAddress)
		copy(dAtA[i:], m.ScriptAddress)
		i = encodeVarintMsgrun(dAtA, i, uint64(len(m.ScriptAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CallerAddress) > 0 {
		i -= len(m.CallerAddress)
		copy(dAtA[i:], m.CallerAddress)
		i = encodeVarintMsgrun(dAtA, i, uint64(len(m.CallerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AttachedMessages) > 0 {
		i -= len(m.AttachedMessages)
		copy(dAtA[i:], m.AttachedMessages)
		i = encodeVarintMsgrun(dAtA, i, uint64(len(m.AttachedMessages)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRunResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRunResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRunResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Response) > 0 {
		i -= len(m.Response)
		copy(dAtA[i:], m.Response)
		i = encodeVarintMsgrun(dAtA, i, uint64(len(m.Response)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgrun(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgrun(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRun) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AttachedMessages)
	if l > 0 {
		n += 1 + l + sovMsgrun(uint64(l))
	}
	l = len(m.CallerAddress)
	if l > 0 {
		n += 1 + l + sovMsgrun(uint64(l))
	}
	l = len(m.ScriptAddress)
	if l > 0 {
		n += 1 + l + sovMsgrun(uint64(l))
	}
	l = len(m.ExtraCode)
	if l > 0 {
		n += 1 + l + sovMsgrun(uint64(l))
	}
	l = len(m.FunctionName)
	if l > 0 {
		n += 1 + l + sovMsgrun(uint64(l))
	}
	l = len(m.Kwargs)
	if l > 0 {
		n += 1 + l + sovMsgrun(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovMsgrun(uint64(l))
	}
	return n
}

func (m *MsgRunResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Response)
	if l > 0 {
		n += 1 + l + sovMsgrun(uint64(l))
	}
	return n
}

func sovMsgrun(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgrun(x uint64) (n int) {
	return sovMsgrun(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRun) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgrun
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
			return fmt.Errorf("proto: MsgRun: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRun: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachedMessages", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgrun
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
				return ErrInvalidLengthMsgrun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgrun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttachedMessages = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgrun
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
				return ErrInvalidLengthMsgrun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgrun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScriptAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgrun
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
				return ErrInvalidLengthMsgrun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgrun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ScriptAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgrun
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
				return ErrInvalidLengthMsgrun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgrun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtraCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FunctionName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgrun
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
				return ErrInvalidLengthMsgrun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgrun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FunctionName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kwargs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgrun
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
				return ErrInvalidLengthMsgrun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgrun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kwargs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgrun
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
				return ErrInvalidLengthMsgrun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgrun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgrun(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgrun
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
func (m *MsgRunResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgrun
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
			return fmt.Errorf("proto: MsgRunResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRunResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgrun
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
				return ErrInvalidLengthMsgrun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgrun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Response = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgrun(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgrun
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
func skipMsgrun(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgrun
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
					return 0, ErrIntOverflowMsgrun
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
					return 0, ErrIntOverflowMsgrun
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
				return 0, ErrInvalidLengthMsgrun
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgrun
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgrun
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgrun        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgrun          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgrun = fmt.Errorf("proto: unexpected end of group")
)
