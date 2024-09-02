// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/fundraising/v1/allowed_bidder.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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

// AllowedBidder defines an allowed bidder for the auction.
type AllowedBidder struct {
	// auction_id specifies the id of the auction
	AuctionId uint64 `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	// bidder specifies the bech32-encoded address that bids for the auction
	Bidder string `protobuf:"bytes,2,opt,name=bidder,proto3" json:"bidder,omitempty"`
	// max_bid_amount specifies the maximum bid amount that the bidder can bid
	MaxBidAmount cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=max_bid_amount,json=maxBidAmount,proto3,customtype=cosmossdk.io/math.Int" json:"max_bid_amount"`
}

func (m *AllowedBidder) Reset()         { *m = AllowedBidder{} }
func (m *AllowedBidder) String() string { return proto.CompactTextString(m) }
func (*AllowedBidder) ProtoMessage()    {}
func (*AllowedBidder) Descriptor() ([]byte, []int) {
	return fileDescriptor_d06892713683ad0f, []int{0}
}
func (m *AllowedBidder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedBidder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedBidder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedBidder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedBidder.Merge(m, src)
}
func (m *AllowedBidder) XXX_Size() int {
	return m.Size()
}
func (m *AllowedBidder) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedBidder.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedBidder proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AllowedBidder)(nil), "modules.fundraising.v1.AllowedBidder")
}

func init() {
	proto.RegisterFile("modules/fundraising/v1/allowed_bidder.proto", fileDescriptor_d06892713683ad0f)
}

var fileDescriptor_d06892713683ad0f = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xce, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0x2b, 0xcd, 0x4b, 0x29, 0x4a, 0xcc, 0x2c, 0xce, 0xcc, 0x4b, 0xd7,
	0x2f, 0x33, 0xd4, 0x4f, 0xcc, 0xc9, 0xc9, 0x2f, 0x4f, 0x4d, 0x89, 0x4f, 0xca, 0x4c, 0x49, 0x49,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x83, 0x2a, 0xd6, 0x43, 0x52, 0xac, 0x57,
	0x66, 0x28, 0x25, 0x9e, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0xac, 0x9f, 0x5b, 0x0c, 0xd6, 0x9b, 0x5b,
	0x9c, 0x0e, 0xd1, 0x20, 0x25, 0x09, 0x91, 0x88, 0x07, 0xf3, 0xf4, 0x21, 0x1c, 0xa8, 0x94, 0x48,
	0x7a, 0x7e, 0x7a, 0x3e, 0x44, 0x1c, 0xc4, 0x82, 0x88, 0x2a, 0xad, 0x62, 0xe4, 0xe2, 0x75, 0x84,
	0x58, 0xed, 0x04, 0xb6, 0x59, 0x48, 0x96, 0x8b, 0x2b, 0xb1, 0x34, 0xb9, 0x24, 0x33, 0x3f, 0x2f,
	0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x25, 0x88, 0x13, 0x2a, 0xe2, 0x99, 0x22, 0x24,
	0xc6, 0xc5, 0x06, 0x71, 0xa2, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x94, 0x27, 0x14, 0xc8,
	0xc5, 0x97, 0x9b, 0x58, 0x01, 0x72, 0x7e, 0x7c, 0x62, 0x6e, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0x33,
	0x48, 0xde, 0x49, 0xfb, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0x45, 0x21, 0x8e, 0x29, 0x4e,
	0xc9, 0xd6, 0xcb, 0xcc, 0xd7, 0xcf, 0x4d, 0x2c, 0xc9, 0xd0, 0xf3, 0xcc, 0x2b, 0xb9, 0xb4, 0x45,
	0x97, 0x0b, 0xea, 0x4a, 0xcf, 0xbc, 0x92, 0x20, 0x9e, 0xdc, 0xc4, 0x0a, 0xa7, 0xcc, 0x14, 0x47,
	0xb0, 0x01, 0x56, 0xfc, 0x1d, 0x0b, 0xe4, 0x19, 0x9a, 0x9e, 0x6f, 0xd0, 0x82, 0xda, 0xe1, 0xe4,
	0x7e, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c,
	0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xba, 0xe9, 0x99, 0x25, 0x19,
	0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x99, 0xe9, 0x79, 0x99, 0x25, 0xa9, 0xfa, 0xb0, 0x70,
	0xae, 0x40, 0x09, 0xe9, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xe7, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x6c, 0x17, 0x2a, 0x12, 0x8d, 0x01, 0x00, 0x00,
}

func (m *AllowedBidder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedBidder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedBidder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxBidAmount.Size()
		i -= size
		if _, err := m.MaxBidAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAllowedBidder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintAllowedBidder(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x12
	}
	if m.AuctionId != 0 {
		i = encodeVarintAllowedBidder(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAllowedBidder(dAtA []byte, offset int, v uint64) int {
	offset -= sovAllowedBidder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AllowedBidder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionId != 0 {
		n += 1 + sovAllowedBidder(uint64(m.AuctionId))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovAllowedBidder(uint64(l))
	}
	l = m.MaxBidAmount.Size()
	n += 1 + l + sovAllowedBidder(uint64(l))
	return n
}

func sovAllowedBidder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAllowedBidder(x uint64) (n int) {
	return sovAllowedBidder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AllowedBidder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllowedBidder
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
			return fmt.Errorf("proto: AllowedBidder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedBidder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllowedBidder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllowedBidder
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
				return ErrInvalidLengthAllowedBidder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllowedBidder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBidAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllowedBidder
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
				return ErrInvalidLengthAllowedBidder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllowedBidder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxBidAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAllowedBidder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAllowedBidder
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
func skipAllowedBidder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAllowedBidder
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
					return 0, ErrIntOverflowAllowedBidder
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
					return 0, ErrIntOverflowAllowedBidder
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
				return 0, ErrInvalidLengthAllowedBidder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAllowedBidder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAllowedBidder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAllowedBidder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAllowedBidder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAllowedBidder = fmt.Errorf("proto: unexpected end of group")
)
