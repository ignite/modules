// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package fundraisingv1

import (
	_ "cosmossdk.io/api/amino"
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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
	md_VestingQueue              protoreflect.MessageDescriptor
	fd_VestingQueue_auction_id   protoreflect.FieldDescriptor
	fd_VestingQueue_auctioneer   protoreflect.FieldDescriptor
	fd_VestingQueue_paying_coin  protoreflect.FieldDescriptor
	fd_VestingQueue_release_time protoreflect.FieldDescriptor
	fd_VestingQueue_released     protoreflect.FieldDescriptor
)

func init() {
	file_modules_fundraising_v1_vesting_queue_proto_init()
	md_VestingQueue = File_modules_fundraising_v1_vesting_queue_proto.Messages().ByName("VestingQueue")
	fd_VestingQueue_auction_id = md_VestingQueue.Fields().ByName("auction_id")
	fd_VestingQueue_auctioneer = md_VestingQueue.Fields().ByName("auctioneer")
	fd_VestingQueue_paying_coin = md_VestingQueue.Fields().ByName("paying_coin")
	fd_VestingQueue_release_time = md_VestingQueue.Fields().ByName("release_time")
	fd_VestingQueue_released = md_VestingQueue.Fields().ByName("released")
}

var _ protoreflect.Message = (*fastReflection_VestingQueue)(nil)

type fastReflection_VestingQueue VestingQueue

func (x *VestingQueue) ProtoReflect() protoreflect.Message {
	return (*fastReflection_VestingQueue)(x)
}

func (x *VestingQueue) slowProtoReflect() protoreflect.Message {
	mi := &file_modules_fundraising_v1_vesting_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_VestingQueue_messageType fastReflection_VestingQueue_messageType
var _ protoreflect.MessageType = fastReflection_VestingQueue_messageType{}

type fastReflection_VestingQueue_messageType struct{}

func (x fastReflection_VestingQueue_messageType) Zero() protoreflect.Message {
	return (*fastReflection_VestingQueue)(nil)
}
func (x fastReflection_VestingQueue_messageType) New() protoreflect.Message {
	return new(fastReflection_VestingQueue)
}
func (x fastReflection_VestingQueue_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_VestingQueue
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_VestingQueue) Descriptor() protoreflect.MessageDescriptor {
	return md_VestingQueue
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_VestingQueue) Type() protoreflect.MessageType {
	return _fastReflection_VestingQueue_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_VestingQueue) New() protoreflect.Message {
	return new(fastReflection_VestingQueue)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_VestingQueue) Interface() protoreflect.ProtoMessage {
	return (*VestingQueue)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_VestingQueue) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.AuctionId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.AuctionId)
		if !f(fd_VestingQueue_auction_id, value) {
			return
		}
	}
	if x.Auctioneer != "" {
		value := protoreflect.ValueOfString(x.Auctioneer)
		if !f(fd_VestingQueue_auctioneer, value) {
			return
		}
	}
	if x.PayingCoin != nil {
		value := protoreflect.ValueOfMessage(x.PayingCoin.ProtoReflect())
		if !f(fd_VestingQueue_paying_coin, value) {
			return
		}
	}
	if x.ReleaseTime != nil {
		value := protoreflect.ValueOfMessage(x.ReleaseTime.ProtoReflect())
		if !f(fd_VestingQueue_release_time, value) {
			return
		}
	}
	if x.Released != false {
		value := protoreflect.ValueOfBool(x.Released)
		if !f(fd_VestingQueue_released, value) {
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
func (x *fastReflection_VestingQueue) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "modules.fundraising.v1.VestingQueue.auction_id":
		return x.AuctionId != uint64(0)
	case "modules.fundraising.v1.VestingQueue.auctioneer":
		return x.Auctioneer != ""
	case "modules.fundraising.v1.VestingQueue.paying_coin":
		return x.PayingCoin != nil
	case "modules.fundraising.v1.VestingQueue.release_time":
		return x.ReleaseTime != nil
	case "modules.fundraising.v1.VestingQueue.released":
		return x.Released != false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: modules.fundraising.v1.VestingQueue"))
		}
		panic(fmt.Errorf("message modules.fundraising.v1.VestingQueue does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VestingQueue) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "modules.fundraising.v1.VestingQueue.auction_id":
		x.AuctionId = uint64(0)
	case "modules.fundraising.v1.VestingQueue.auctioneer":
		x.Auctioneer = ""
	case "modules.fundraising.v1.VestingQueue.paying_coin":
		x.PayingCoin = nil
	case "modules.fundraising.v1.VestingQueue.release_time":
		x.ReleaseTime = nil
	case "modules.fundraising.v1.VestingQueue.released":
		x.Released = false
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: modules.fundraising.v1.VestingQueue"))
		}
		panic(fmt.Errorf("message modules.fundraising.v1.VestingQueue does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_VestingQueue) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "modules.fundraising.v1.VestingQueue.auction_id":
		value := x.AuctionId
		return protoreflect.ValueOfUint64(value)
	case "modules.fundraising.v1.VestingQueue.auctioneer":
		value := x.Auctioneer
		return protoreflect.ValueOfString(value)
	case "modules.fundraising.v1.VestingQueue.paying_coin":
		value := x.PayingCoin
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "modules.fundraising.v1.VestingQueue.release_time":
		value := x.ReleaseTime
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "modules.fundraising.v1.VestingQueue.released":
		value := x.Released
		return protoreflect.ValueOfBool(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: modules.fundraising.v1.VestingQueue"))
		}
		panic(fmt.Errorf("message modules.fundraising.v1.VestingQueue does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_VestingQueue) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "modules.fundraising.v1.VestingQueue.auction_id":
		x.AuctionId = value.Uint()
	case "modules.fundraising.v1.VestingQueue.auctioneer":
		x.Auctioneer = value.Interface().(string)
	case "modules.fundraising.v1.VestingQueue.paying_coin":
		x.PayingCoin = value.Message().Interface().(*v1beta1.Coin)
	case "modules.fundraising.v1.VestingQueue.release_time":
		x.ReleaseTime = value.Message().Interface().(*timestamppb.Timestamp)
	case "modules.fundraising.v1.VestingQueue.released":
		x.Released = value.Bool()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: modules.fundraising.v1.VestingQueue"))
		}
		panic(fmt.Errorf("message modules.fundraising.v1.VestingQueue does not contain field %s", fd.FullName()))
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
func (x *fastReflection_VestingQueue) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "modules.fundraising.v1.VestingQueue.paying_coin":
		if x.PayingCoin == nil {
			x.PayingCoin = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.PayingCoin.ProtoReflect())
	case "modules.fundraising.v1.VestingQueue.release_time":
		if x.ReleaseTime == nil {
			x.ReleaseTime = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.ReleaseTime.ProtoReflect())
	case "modules.fundraising.v1.VestingQueue.auction_id":
		panic(fmt.Errorf("field auction_id of message modules.fundraising.v1.VestingQueue is not mutable"))
	case "modules.fundraising.v1.VestingQueue.auctioneer":
		panic(fmt.Errorf("field auctioneer of message modules.fundraising.v1.VestingQueue is not mutable"))
	case "modules.fundraising.v1.VestingQueue.released":
		panic(fmt.Errorf("field released of message modules.fundraising.v1.VestingQueue is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: modules.fundraising.v1.VestingQueue"))
		}
		panic(fmt.Errorf("message modules.fundraising.v1.VestingQueue does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_VestingQueue) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "modules.fundraising.v1.VestingQueue.auction_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "modules.fundraising.v1.VestingQueue.auctioneer":
		return protoreflect.ValueOfString("")
	case "modules.fundraising.v1.VestingQueue.paying_coin":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "modules.fundraising.v1.VestingQueue.release_time":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "modules.fundraising.v1.VestingQueue.released":
		return protoreflect.ValueOfBool(false)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: modules.fundraising.v1.VestingQueue"))
		}
		panic(fmt.Errorf("message modules.fundraising.v1.VestingQueue does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_VestingQueue) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in modules.fundraising.v1.VestingQueue", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_VestingQueue) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VestingQueue) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_VestingQueue) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_VestingQueue) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*VestingQueue)
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
		if x.AuctionId != 0 {
			n += 1 + runtime.Sov(uint64(x.AuctionId))
		}
		l = len(x.Auctioneer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.PayingCoin != nil {
			l = options.Size(x.PayingCoin)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.ReleaseTime != nil {
			l = options.Size(x.ReleaseTime)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Released {
			n += 2
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
		x := input.Message.Interface().(*VestingQueue)
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
		if x.Released {
			i--
			if x.Released {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x28
		}
		if x.ReleaseTime != nil {
			encoded, err := options.Marshal(x.ReleaseTime)
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
			dAtA[i] = 0x22
		}
		if x.PayingCoin != nil {
			encoded, err := options.Marshal(x.PayingCoin)
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
			dAtA[i] = 0x1a
		}
		if len(x.Auctioneer) > 0 {
			i -= len(x.Auctioneer)
			copy(dAtA[i:], x.Auctioneer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Auctioneer)))
			i--
			dAtA[i] = 0x12
		}
		if x.AuctionId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.AuctionId))
			i--
			dAtA[i] = 0x8
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
		x := input.Message.Interface().(*VestingQueue)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: VestingQueue: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: VestingQueue: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
				}
				x.AuctionId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.AuctionId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Auctioneer", wireType)
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
				x.Auctioneer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PayingCoin", wireType)
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
				if x.PayingCoin == nil {
					x.PayingCoin = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PayingCoin); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ReleaseTime", wireType)
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
				if x.ReleaseTime == nil {
					x.ReleaseTime = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ReleaseTime); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Released", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Released = bool(v != 0)
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
// source: modules/fundraising/v1/vesting_queue.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// VestingQueue defines the vesting queue.
type VestingQueue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// auction_id specifies the id of the auction
	AuctionId uint64 `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	// auctioneer specifies the bech32-encoded address that creates the auction
	Auctioneer string `protobuf:"bytes,2,opt,name=auctioneer,proto3" json:"auctioneer,omitempty"`
	// paying_coin specifies the paying amount of coin
	PayingCoin *v1beta1.Coin `protobuf:"bytes,3,opt,name=paying_coin,json=payingCoin,proto3" json:"paying_coin,omitempty"`
	// release_time specifies the timestamp of the vesting schedule
	ReleaseTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=release_time,json=releaseTime,proto3" json:"release_time,omitempty"`
	// released specifies the status of distribution
	Released bool `protobuf:"varint,5,opt,name=released,proto3" json:"released,omitempty"`
}

func (x *VestingQueue) Reset() {
	*x = VestingQueue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_fundraising_v1_vesting_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VestingQueue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VestingQueue) ProtoMessage() {}

// Deprecated: Use VestingQueue.ProtoReflect.Descriptor instead.
func (*VestingQueue) Descriptor() ([]byte, []int) {
	return file_modules_fundraising_v1_vesting_queue_proto_rawDescGZIP(), []int{0}
}

func (x *VestingQueue) GetAuctionId() uint64 {
	if x != nil {
		return x.AuctionId
	}
	return 0
}

func (x *VestingQueue) GetAuctioneer() string {
	if x != nil {
		return x.Auctioneer
	}
	return ""
}

func (x *VestingQueue) GetPayingCoin() *v1beta1.Coin {
	if x != nil {
		return x.PayingCoin
	}
	return nil
}

func (x *VestingQueue) GetReleaseTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReleaseTime
	}
	return nil
}

func (x *VestingQueue) GetReleased() bool {
	if x != nil {
		return x.Released
	}
	return false
}

var File_modules_fundraising_v1_vesting_queue_proto protoreflect.FileDescriptor

var file_modules_fundraising_v1_vesting_queue_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x72, 0x61,
	0x69, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x72, 0x61, 0x69, 0x73, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x02, 0x0a, 0x0c, 0x56, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0a, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2,
	0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x65, 0x65, 0x72, 0x12, 0x81, 0x01, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x6f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x43, 0x6f, 0x69, 0x6e, 0x42, 0x45, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x9a, 0xe7, 0xb0, 0x2a, 0x0c, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79,
	0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x79,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x47, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90,
	0xdf, 0x1f, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x42, 0xe0, 0x01, 0x0a,
	0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x66, 0x75, 0x6e,
	0x64, 0x72, 0x61, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x56, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x35, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x72,
	0x61, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x75, 0x6e, 0x64, 0x72, 0x61,
	0x69, 0x73, 0x69, 0x6e, 0x67, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x46, 0x58, 0xaa, 0x02, 0x16,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x72, 0x61, 0x69, 0x73,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x5c, 0x46, 0x75, 0x6e, 0x64, 0x72, 0x61, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x22, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x5c, 0x46, 0x75, 0x6e, 0x64, 0x72, 0x61,
	0x69, 0x73, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x3a, 0x3a,
	0x46, 0x75, 0x6e, 0x64, 0x72, 0x61, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_fundraising_v1_vesting_queue_proto_rawDescOnce sync.Once
	file_modules_fundraising_v1_vesting_queue_proto_rawDescData = file_modules_fundraising_v1_vesting_queue_proto_rawDesc
)

func file_modules_fundraising_v1_vesting_queue_proto_rawDescGZIP() []byte {
	file_modules_fundraising_v1_vesting_queue_proto_rawDescOnce.Do(func() {
		file_modules_fundraising_v1_vesting_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_fundraising_v1_vesting_queue_proto_rawDescData)
	})
	return file_modules_fundraising_v1_vesting_queue_proto_rawDescData
}

var file_modules_fundraising_v1_vesting_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_modules_fundraising_v1_vesting_queue_proto_goTypes = []interface{}{
	(*VestingQueue)(nil),          // 0: modules.fundraising.v1.VestingQueue
	(*v1beta1.Coin)(nil),          // 1: cosmos.base.v1beta1.Coin
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_modules_fundraising_v1_vesting_queue_proto_depIdxs = []int32{
	1, // 0: modules.fundraising.v1.VestingQueue.paying_coin:type_name -> cosmos.base.v1beta1.Coin
	2, // 1: modules.fundraising.v1.VestingQueue.release_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_modules_fundraising_v1_vesting_queue_proto_init() }
func file_modules_fundraising_v1_vesting_queue_proto_init() {
	if File_modules_fundraising_v1_vesting_queue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_fundraising_v1_vesting_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VestingQueue); i {
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
			RawDescriptor: file_modules_fundraising_v1_vesting_queue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_modules_fundraising_v1_vesting_queue_proto_goTypes,
		DependencyIndexes: file_modules_fundraising_v1_vesting_queue_proto_depIdxs,
		MessageInfos:      file_modules_fundraising_v1_vesting_queue_proto_msgTypes,
	}.Build()
	File_modules_fundraising_v1_vesting_queue_proto = out.File
	file_modules_fundraising_v1_vesting_queue_proto_rawDesc = nil
	file_modules_fundraising_v1_vesting_queue_proto_goTypes = nil
	file_modules_fundraising_v1_vesting_queue_proto_depIdxs = nil
}
