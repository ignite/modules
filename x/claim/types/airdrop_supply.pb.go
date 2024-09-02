// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/claim/v1/airdrop_supply.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

type AirdropSupply struct {
	Supply types.Coin `protobuf:"bytes,1,opt,name=supply,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"supply"`
}

func (m *AirdropSupply) Reset()         { *m = AirdropSupply{} }
func (m *AirdropSupply) String() string { return proto.CompactTextString(m) }
func (*AirdropSupply) ProtoMessage()    {}
func (*AirdropSupply) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd44ab68141c502d, []int{0}
}
func (m *AirdropSupply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AirdropSupply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AirdropSupply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AirdropSupply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AirdropSupply.Merge(m, src)
}
func (m *AirdropSupply) XXX_Size() int {
	return m.Size()
}
func (m *AirdropSupply) XXX_DiscardUnknown() {
	xxx_messageInfo_AirdropSupply.DiscardUnknown(m)
}

var xxx_messageInfo_AirdropSupply proto.InternalMessageInfo

func (m *AirdropSupply) GetSupply() types.Coin {
	if m != nil {
		return m.Supply
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*AirdropSupply)(nil), "modules.claim.v1.AirdropSupply")
}

func init() {
	proto.RegisterFile("modules/claim/v1/airdrop_supply.proto", fileDescriptor_fd44ab68141c502d)
}

var fileDescriptor_fd44ab68141c502d = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0xce, 0x49, 0xcc, 0xcc, 0xd5, 0x2f, 0x33, 0xd4, 0x4f, 0xcc, 0x2c,
	0x4a, 0x29, 0xca, 0x2f, 0x88, 0x2f, 0x2e, 0x2d, 0x28, 0xc8, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x80, 0x2a, 0xd3, 0x03, 0x2b, 0xd3, 0x2b, 0x33, 0x94, 0x12, 0x4c, 0xcc, 0xcd,
	0xcc, 0xcb, 0xd7, 0x07, 0x93, 0x10, 0x45, 0x52, 0x72, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xfa,
	0x49, 0x89, 0xc5, 0xa9, 0xfa, 0x65, 0x86, 0x49, 0xa9, 0x25, 0x89, 0x86, 0xfa, 0xc9, 0xf9, 0x99,
	0x79, 0x50, 0x79, 0x49, 0x88, 0x7c, 0x3c, 0x98, 0xa7, 0x0f, 0xe1, 0x40, 0xa5, 0x44, 0xd2, 0xf3,
	0xd3, 0xf3, 0x21, 0xe2, 0x20, 0x16, 0x44, 0x54, 0xa9, 0x8b, 0x91, 0x8b, 0xd7, 0x11, 0xe2, 0x9c,
	0x60, 0xb0, 0x6b, 0x84, 0x2a, 0xb9, 0xd8, 0x20, 0xee, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36,
	0x92, 0xd4, 0x83, 0x1a, 0x03, 0xb2, 0x53, 0x0f, 0x6a, 0xa7, 0x9e, 0x73, 0x7e, 0x66, 0x9e, 0x93,
	0xdb, 0x89, 0x7b, 0xf2, 0x0c, 0xab, 0xee, 0xcb, 0x6b, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0x42, 0xed, 0x84, 0x52, 0xba, 0xc5, 0x29, 0xd9, 0xfa, 0x25, 0x95, 0x05, 0xa9,
	0xc5, 0x60, 0x0d, 0xc5, 0xb3, 0x9e, 0x6f, 0xd0, 0xe2, 0xc9, 0x49, 0x4d, 0x4f, 0x4c, 0xae, 0x8c,
	0x07, 0xb9, 0xba, 0x78, 0xc5, 0xf3, 0x0d, 0x5a, 0x8c, 0x41, 0x50, 0x0b, 0x9d, 0x1c, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x1d, 0xc9, 0x86, 0xcc, 0xf4, 0xbc, 0xcc,
	0x92, 0x54, 0x7d, 0x58, 0xa8, 0x56, 0x40, 0xc3, 0x15, 0x6c, 0x4d, 0x12, 0x1b, 0xd8, 0x5b, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xf1, 0xb1, 0x14, 0x75, 0x01, 0x00, 0x00,
}

func (m *AirdropSupply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AirdropSupply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AirdropSupply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Supply.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAirdropSupply(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAirdropSupply(dAtA []byte, offset int, v uint64) int {
	offset -= sovAirdropSupply(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AirdropSupply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Supply.Size()
	n += 1 + l + sovAirdropSupply(uint64(l))
	return n
}

func sovAirdropSupply(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAirdropSupply(x uint64) (n int) {
	return sovAirdropSupply(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AirdropSupply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAirdropSupply
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
			return fmt.Errorf("proto: AirdropSupply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AirdropSupply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAirdropSupply
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
				return ErrInvalidLengthAirdropSupply
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAirdropSupply
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Supply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAirdropSupply(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAirdropSupply
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
func skipAirdropSupply(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAirdropSupply
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
					return 0, ErrIntOverflowAirdropSupply
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
					return 0, ErrIntOverflowAirdropSupply
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
				return 0, ErrInvalidLengthAirdropSupply
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAirdropSupply
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAirdropSupply
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAirdropSupply        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAirdropSupply          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAirdropSupply = fmt.Errorf("proto: unexpected end of group")
)
