// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/mint/events.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// EventMint is emitted when new coins are minted by the minter
type EventMint struct {
	BondedRatio      cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=bondedRatio,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"bondedRatio"`
	Inflation        cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=inflation,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation"`
	AnnualProvisions cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=annualProvisions,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"annualProvisions"`
	Amount           cosmossdk_io_math.Int       `protobuf:"bytes,4,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
}

func (m *EventMint) Reset()         { *m = EventMint{} }
func (m *EventMint) String() string { return proto.CompactTextString(m) }
func (*EventMint) ProtoMessage()    {}
func (*EventMint) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ae1e817e75710b8, []int{0}
}
func (m *EventMint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMint.Merge(m, src)
}
func (m *EventMint) XXX_Size() int {
	return m.Size()
}
func (m *EventMint) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMint.DiscardUnknown(m)
}

var xxx_messageInfo_EventMint proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventMint)(nil), "modules.mint.EventMint")
}

func init() { proto.RegisterFile("modules/mint/events.proto", fileDescriptor_0ae1e817e75710b8) }

var fileDescriptor_0ae1e817e75710b8 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd1, 0x41, 0x4b, 0x02, 0x41,
	0x14, 0xc0, 0xf1, 0x5d, 0x0b, 0xc1, 0xa9, 0x43, 0x2c, 0x05, 0xab, 0xc1, 0x18, 0x1d, 0x22, 0x88,
	0x76, 0x88, 0xbe, 0x40, 0x98, 0x1d, 0x84, 0xa2, 0xb0, 0x5b, 0x10, 0x31, 0xee, 0x4e, 0xeb, 0x90,
	0xf3, 0x9e, 0x38, 0x6f, 0x25, 0xbf, 0x45, 0x1f, 0xa6, 0x0f, 0xe1, 0x51, 0x82, 0x20, 0x3a, 0x48,
	0xe8, 0x17, 0x89, 0x75, 0x36, 0x12, 0xbc, 0x79, 0x7b, 0xc3, 0x7f, 0xde, 0xef, 0xf2, 0x58, 0xd5,
	0x60, 0x92, 0xf5, 0x94, 0x15, 0x46, 0x03, 0x09, 0x35, 0x54, 0x40, 0x36, 0xea, 0x0f, 0x90, 0x30,
	0xd8, 0x2e, 0x52, 0x94, 0xa7, 0xda, 0x6e, 0x8a, 0x29, 0x2e, 0x82, 0xc8, 0x27, 0xf7, 0xa7, 0x56,
	0x8d, 0xd1, 0x1a, 0xb4, 0x4f, 0x2e, 0xb8, 0x87, 0x4b, 0x87, 0x9f, 0x25, 0x56, 0xb9, 0xca, 0xbd,
	0x1b, 0x0d, 0x14, 0xdc, 0xb3, 0xad, 0x0e, 0x42, 0xa2, 0x92, 0xb6, 0x24, 0x8d, 0xa1, 0x7f, 0xe0,
	0x1f, 0x57, 0x1a, 0x67, 0xe3, 0x69, 0xdd, 0xfb, 0x9e, 0xd6, 0xf7, 0xdd, 0xa2, 0x4d, 0x5e, 0x22,
	0x8d, 0xc2, 0x48, 0xea, 0x46, 0xd7, 0x2a, 0x95, 0xf1, 0xa8, 0xa9, 0xe2, 0x8f, 0xf7, 0x53, 0x56,
	0xb8, 0x4d, 0x15, 0xb7, 0x97, 0x95, 0xe0, 0x96, 0x55, 0x34, 0x3c, 0xf7, 0xf2, 0x19, 0xc2, 0xd2,
	0xba, 0xe4, 0xbf, 0x11, 0x3c, 0xb2, 0x1d, 0x09, 0x90, 0xc9, 0xde, 0xdd, 0x00, 0x87, 0xda, 0x6a,
	0x04, 0x1b, 0x6e, 0xac, 0xeb, 0xae, 0x50, 0xc1, 0x25, 0x2b, 0x4b, 0x83, 0x19, 0x50, 0xb8, 0xb9,
	0x40, 0x4f, 0x0a, 0x74, 0x6f, 0x15, 0x6d, 0x01, 0x2d, 0x71, 0x2d, 0xa0, 0x76, 0xb1, 0xda, 0xb8,
	0x18, 0xcf, 0xb8, 0x3f, 0x99, 0x71, 0xff, 0x67, 0xc6, 0xfd, 0xb7, 0x39, 0xf7, 0x26, 0x73, 0xee,
	0x7d, 0xcd, 0xb9, 0xf7, 0x70, 0x94, 0x6a, 0xea, 0x66, 0x9d, 0x28, 0x46, 0x23, 0x74, 0x0a, 0x9a,
	0x94, 0xf8, 0xbb, 0xee, 0xab, 0xbb, 0x2f, 0x8d, 0xfa, 0xca, 0x76, 0xca, 0x8b, 0x03, 0x9d, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x83, 0x04, 0x43, 0xac, 0xfc, 0x01, 0x00, 0x00,
}

func (m *EventMint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.BondedRatio.Size()
		i -= size
		if _, err := m.BondedRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *EventMint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BondedRatio.Size()
	n += 1 + l + sovEvents(uint64(l))
	l = m.Inflation.Size()
	n += 1 + l + sovEvents(uint64(l))
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovEvents(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventMint) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventMint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondedRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BondedRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
