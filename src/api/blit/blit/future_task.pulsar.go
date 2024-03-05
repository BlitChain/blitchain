// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package blit

import (
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_FutureTask              protoreflect.MessageDescriptor
	fd_FutureTask_index        protoreflect.FieldDescriptor
	fd_FutureTask_scheduled_on protoreflect.FieldDescriptor
	fd_FutureTask_task_id      protoreflect.FieldDescriptor
	fd_FutureTask_status       protoreflect.FieldDescriptor
	fd_FutureTask_gas_price    protoreflect.FieldDescriptor
)

func init() {
	file_blit_blit_future_task_proto_init()
	md_FutureTask = File_blit_blit_future_task_proto.Messages().ByName("FutureTask")
	fd_FutureTask_index = md_FutureTask.Fields().ByName("index")
	fd_FutureTask_scheduled_on = md_FutureTask.Fields().ByName("scheduled_on")
	fd_FutureTask_task_id = md_FutureTask.Fields().ByName("task_id")
	fd_FutureTask_status = md_FutureTask.Fields().ByName("status")
	fd_FutureTask_gas_price = md_FutureTask.Fields().ByName("gas_price")
}

var _ protoreflect.Message = (*fastReflection_FutureTask)(nil)

type fastReflection_FutureTask FutureTask

func (x *FutureTask) ProtoReflect() protoreflect.Message {
	return (*fastReflection_FutureTask)(x)
}

func (x *FutureTask) slowProtoReflect() protoreflect.Message {
	mi := &file_blit_blit_future_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_FutureTask_messageType fastReflection_FutureTask_messageType
var _ protoreflect.MessageType = fastReflection_FutureTask_messageType{}

type fastReflection_FutureTask_messageType struct{}

func (x fastReflection_FutureTask_messageType) Zero() protoreflect.Message {
	return (*fastReflection_FutureTask)(nil)
}
func (x fastReflection_FutureTask_messageType) New() protoreflect.Message {
	return new(fastReflection_FutureTask)
}
func (x fastReflection_FutureTask_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_FutureTask
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_FutureTask) Descriptor() protoreflect.MessageDescriptor {
	return md_FutureTask
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_FutureTask) Type() protoreflect.MessageType {
	return _fastReflection_FutureTask_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_FutureTask) New() protoreflect.Message {
	return new(fastReflection_FutureTask)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_FutureTask) Interface() protoreflect.ProtoMessage {
	return (*FutureTask)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_FutureTask) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Index != "" {
		value := protoreflect.ValueOfString(x.Index)
		if !f(fd_FutureTask_index, value) {
			return
		}
	}
	if x.ScheduledOn != nil {
		value := protoreflect.ValueOfMessage(x.ScheduledOn.ProtoReflect())
		if !f(fd_FutureTask_scheduled_on, value) {
			return
		}
	}
	if x.TaskId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TaskId)
		if !f(fd_FutureTask_task_id, value) {
			return
		}
	}
	if x.Status != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Status))
		if !f(fd_FutureTask_status, value) {
			return
		}
	}
	if x.GasPrice != nil {
		value := protoreflect.ValueOfMessage(x.GasPrice.ProtoReflect())
		if !f(fd_FutureTask_gas_price, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_FutureTask) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "blit.blit.FutureTask.index":
		return x.Index != ""
	case "blit.blit.FutureTask.scheduled_on":
		return x.ScheduledOn != nil
	case "blit.blit.FutureTask.task_id":
		return x.TaskId != uint64(0)
	case "blit.blit.FutureTask.status":
		return x.Status != 0
	case "blit.blit.FutureTask.gas_price":
		return x.GasPrice != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: blit.blit.FutureTask"))
		}
		panic(fmt.Errorf("message blit.blit.FutureTask does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FutureTask) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "blit.blit.FutureTask.index":
		x.Index = ""
	case "blit.blit.FutureTask.scheduled_on":
		x.ScheduledOn = nil
	case "blit.blit.FutureTask.task_id":
		x.TaskId = uint64(0)
	case "blit.blit.FutureTask.status":
		x.Status = 0
	case "blit.blit.FutureTask.gas_price":
		x.GasPrice = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: blit.blit.FutureTask"))
		}
		panic(fmt.Errorf("message blit.blit.FutureTask does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_FutureTask) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "blit.blit.FutureTask.index":
		value := x.Index
		return protoreflect.ValueOfString(value)
	case "blit.blit.FutureTask.scheduled_on":
		value := x.ScheduledOn
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "blit.blit.FutureTask.task_id":
		value := x.TaskId
		return protoreflect.ValueOfUint64(value)
	case "blit.blit.FutureTask.status":
		value := x.Status
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "blit.blit.FutureTask.gas_price":
		value := x.GasPrice
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: blit.blit.FutureTask"))
		}
		panic(fmt.Errorf("message blit.blit.FutureTask does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FutureTask) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "blit.blit.FutureTask.index":
		x.Index = value.Interface().(string)
	case "blit.blit.FutureTask.scheduled_on":
		x.ScheduledOn = value.Message().Interface().(*timestamppb.Timestamp)
	case "blit.blit.FutureTask.task_id":
		x.TaskId = value.Uint()
	case "blit.blit.FutureTask.status":
		x.Status = (FutureTaskStatus)(value.Enum())
	case "blit.blit.FutureTask.gas_price":
		x.GasPrice = value.Message().Interface().(*v1beta1.DecCoin)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: blit.blit.FutureTask"))
		}
		panic(fmt.Errorf("message blit.blit.FutureTask does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FutureTask) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "blit.blit.FutureTask.scheduled_on":
		if x.ScheduledOn == nil {
			x.ScheduledOn = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.ScheduledOn.ProtoReflect())
	case "blit.blit.FutureTask.gas_price":
		if x.GasPrice == nil {
			x.GasPrice = new(v1beta1.DecCoin)
		}
		return protoreflect.ValueOfMessage(x.GasPrice.ProtoReflect())
	case "blit.blit.FutureTask.index":
		panic(fmt.Errorf("field index of message blit.blit.FutureTask is not mutable"))
	case "blit.blit.FutureTask.task_id":
		panic(fmt.Errorf("field task_id of message blit.blit.FutureTask is not mutable"))
	case "blit.blit.FutureTask.status":
		panic(fmt.Errorf("field status of message blit.blit.FutureTask is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: blit.blit.FutureTask"))
		}
		panic(fmt.Errorf("message blit.blit.FutureTask does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_FutureTask) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "blit.blit.FutureTask.index":
		return protoreflect.ValueOfString("")
	case "blit.blit.FutureTask.scheduled_on":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "blit.blit.FutureTask.task_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "blit.blit.FutureTask.status":
		return protoreflect.ValueOfEnum(0)
	case "blit.blit.FutureTask.gas_price":
		m := new(v1beta1.DecCoin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: blit.blit.FutureTask"))
		}
		panic(fmt.Errorf("message blit.blit.FutureTask does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_FutureTask) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in blit.blit.FutureTask", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_FutureTask) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FutureTask) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_FutureTask) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_FutureTask) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*FutureTask)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Index)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.ScheduledOn != nil {
			l = options.Size(x.ScheduledOn)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.TaskId != 0 {
			n += 1 + runtime.Sov(uint64(x.TaskId))
		}
		if x.Status != 0 {
			n += 1 + runtime.Sov(uint64(x.Status))
		}
		if x.GasPrice != nil {
			l = options.Size(x.GasPrice)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*FutureTask)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.GasPrice != nil {
			encoded, err := options.Marshal(x.GasPrice)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if x.Status != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Status))
			i--
			dAtA[i] = 0x20
		}
		if x.TaskId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TaskId))
			i--
			dAtA[i] = 0x18
		}
		if x.ScheduledOn != nil {
			encoded, err := options.Marshal(x.ScheduledOn)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Index) > 0 {
			i -= len(x.Index)
			copy(dAtA[i:], x.Index)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Index)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*FutureTask)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: FutureTask: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: FutureTask: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Index = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ScheduledOn", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.ScheduledOn == nil {
					x.ScheduledOn = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ScheduledOn); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TaskId", wireType)
				}
				x.TaskId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TaskId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
				}
				x.Status = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Status |= FutureTaskStatus(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.GasPrice == nil {
					x.GasPrice = &v1beta1.DecCoin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.GasPrice); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: blit/blit/future_task.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FutureTaskStatus int32

const (
	FutureTaskStatus_NONE    FutureTaskStatus = 0
	FutureTaskStatus_PENDING FutureTaskStatus = 1
	FutureTaskStatus_POOL    FutureTaskStatus = 2
)

// Enum value maps for FutureTaskStatus.
var (
	FutureTaskStatus_name = map[int32]string{
		0: "NONE",
		1: "PENDING",
		2: "POOL",
	}
	FutureTaskStatus_value = map[string]int32{
		"NONE":    0,
		"PENDING": 1,
		"POOL":    2,
	}
)

func (x FutureTaskStatus) Enum() *FutureTaskStatus {
	p := new(FutureTaskStatus)
	*p = x
	return p
}

func (x FutureTaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FutureTaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_blit_blit_future_task_proto_enumTypes[0].Descriptor()
}

func (FutureTaskStatus) Type() protoreflect.EnumType {
	return &file_blit_blit_future_task_proto_enumTypes[0]
}

func (x FutureTaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FutureTaskStatus.Descriptor instead.
func (FutureTaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_blit_blit_future_task_proto_rawDescGZIP(), []int{0}
}

type FutureTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index       string                 `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	ScheduledOn *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=scheduled_on,json=scheduledOn,proto3" json:"scheduled_on,omitempty"`
	TaskId      uint64                 `protobuf:"varint,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Status      FutureTaskStatus       `protobuf:"varint,4,opt,name=status,proto3,enum=blit.blit.FutureTaskStatus" json:"status,omitempty"`
	GasPrice    *v1beta1.DecCoin       `protobuf:"bytes,5,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
}

func (x *FutureTask) Reset() {
	*x = FutureTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blit_blit_future_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FutureTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FutureTask) ProtoMessage() {}

// Deprecated: Use FutureTask.ProtoReflect.Descriptor instead.
func (*FutureTask) Descriptor() ([]byte, []int) {
	return file_blit_blit_future_task_proto_rawDescGZIP(), []int{0}
}

func (x *FutureTask) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *FutureTask) GetScheduledOn() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledOn
	}
	return nil
}

func (x *FutureTask) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *FutureTask) GetStatus() FutureTaskStatus {
	if x != nil {
		return x.Status
	}
	return FutureTaskStatus_NONE
}

func (x *FutureTask) GetGasPrice() *v1beta1.DecCoin {
	if x != nil {
		return x.GasPrice
	}
	return nil
}

var File_blit_blit_future_task_proto protoreflect.FileDescriptor

var file_blit_blit_future_task_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x62, 0x6c, 0x69, 0x74, 0x2f, 0x62, 0x6c, 0x69, 0x74, 0x2f, 0x66, 0x75, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62,
	0x6c, 0x69, 0x74, 0x2e, 0x62, 0x6c, 0x69, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf4, 0x01, 0x0a, 0x0a, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x47, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01,
	0x52, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62, 0x6c, 0x69, 0x74, 0x2e, 0x62, 0x6c,
	0x69, 0x74, 0x2e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x09, 0x67,
	0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x08, 0x67, 0x61,
	0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x2a, 0x33, 0x0a, 0x10, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x4f, 0x4c, 0x10, 0x02, 0x42, 0x81, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6c, 0x69, 0x74, 0x2e, 0x62, 0x6c, 0x69, 0x74, 0x42, 0x0f, 0x46,
	0x75, 0x74, 0x75, 0x72, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x1a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x62, 0x6c, 0x69, 0x74, 0x2f, 0x62, 0x6c, 0x69, 0x74, 0xa2, 0x02, 0x03, 0x42,
	0x42, 0x58, 0xaa, 0x02, 0x09, 0x42, 0x6c, 0x69, 0x74, 0x2e, 0x42, 0x6c, 0x69, 0x74, 0xca, 0x02,
	0x09, 0x42, 0x6c, 0x69, 0x74, 0x5c, 0x42, 0x6c, 0x69, 0x74, 0xe2, 0x02, 0x15, 0x42, 0x6c, 0x69,
	0x74, 0x5c, 0x42, 0x6c, 0x69, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0a, 0x42, 0x6c, 0x69, 0x74, 0x3a, 0x3a, 0x42, 0x6c, 0x69, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blit_blit_future_task_proto_rawDescOnce sync.Once
	file_blit_blit_future_task_proto_rawDescData = file_blit_blit_future_task_proto_rawDesc
)

func file_blit_blit_future_task_proto_rawDescGZIP() []byte {
	file_blit_blit_future_task_proto_rawDescOnce.Do(func() {
		file_blit_blit_future_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_blit_blit_future_task_proto_rawDescData)
	})
	return file_blit_blit_future_task_proto_rawDescData
}

var file_blit_blit_future_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_blit_blit_future_task_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_blit_blit_future_task_proto_goTypes = []interface{}{
	(FutureTaskStatus)(0),         // 0: blit.blit.FutureTaskStatus
	(*FutureTask)(nil),            // 1: blit.blit.FutureTask
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*v1beta1.DecCoin)(nil),       // 3: cosmos.base.v1beta1.DecCoin
}
var file_blit_blit_future_task_proto_depIdxs = []int32{
	2, // 0: blit.blit.FutureTask.scheduled_on:type_name -> google.protobuf.Timestamp
	0, // 1: blit.blit.FutureTask.status:type_name -> blit.blit.FutureTaskStatus
	3, // 2: blit.blit.FutureTask.gas_price:type_name -> cosmos.base.v1beta1.DecCoin
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_blit_blit_future_task_proto_init() }
func file_blit_blit_future_task_proto_init() {
	if File_blit_blit_future_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blit_blit_future_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FutureTask); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blit_blit_future_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blit_blit_future_task_proto_goTypes,
		DependencyIndexes: file_blit_blit_future_task_proto_depIdxs,
		EnumInfos:         file_blit_blit_future_task_proto_enumTypes,
		MessageInfos:      file_blit_blit_future_task_proto_msgTypes,
	}.Build()
	File_blit_blit_future_task_proto = out.File
	file_blit_blit_future_task_proto_rawDesc = nil
	file_blit_blit_future_task_proto_goTypes = nil
	file_blit_blit_future_task_proto_depIdxs = nil
}
