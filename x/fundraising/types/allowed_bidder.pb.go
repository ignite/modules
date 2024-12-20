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
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x3f, 0x4f, 0x3a, 0x31,
	0x18, 0xc7, 0xaf, 0xbf, 0x9f, 0x21, 0xa1, 0xf1, 0x4f, 0x72, 0x41, 0x3d, 0x49, 0x3c, 0x88, 0x13,
	0x91, 0x70, 0x95, 0xb8, 0xb9, 0x71, 0x8b, 0x61, 0x14, 0x37, 0x97, 0x4b, 0xa1, 0x67, 0x69, 0xa4,
	0x2d, 0xb9, 0xf6, 0x10, 0x57, 0x27, 0x47, 0x5f, 0x82, 0x2f, 0xc1, 0x81, 0x97, 0xe0, 0xc0, 0x48,
	0x98, 0x8c, 0x03, 0x31, 0xdc, 0xe0, 0xdb, 0x30, 0xb4, 0x35, 0xd1, 0xad, 0xcf, 0xf7, 0xf9, 0x7e,
	0x9f, 0x4f, 0xdb, 0x07, 0x36, 0xb9, 0x24, 0xf9, 0x28, 0x55, 0xe8, 0x36, 0x17, 0x24, 0xc3, 0x4c,
	0x31, 0x41, 0xd1, 0xa4, 0x8d, 0xf0, 0x68, 0x24, 0xef, 0x53, 0x92, 0xf4, 0x19, 0x21, 0x69, 0x16,
	0x8d, 0x33, 0xa9, 0xa5, 0x7f, 0xe0, 0xcc, 0xd1, 0x2f, 0x73, 0x34, 0x69, 0x57, 0x0f, 0x07, 0x52,
	0x71, 0xa9, 0x10, 0x57, 0x26, 0xcb, 0x15, 0xb5, 0x81, 0xea, 0x91, 0x6d, 0x24, 0xa6, 0x42, 0xb6,
	0x70, 0xad, 0x0a, 0x95, 0x54, 0x5a, 0x7d, 0x73, 0xb2, 0xea, 0xc9, 0x1b, 0x80, 0x3b, 0x1d, 0x8b,
	0x8e, 0x0d, 0xd9, 0x3f, 0x86, 0x10, 0xe7, 0x03, 0xcd, 0xa4, 0x48, 0x18, 0x09, 0x40, 0x1d, 0x34,
	0xb6, 0x7a, 0x65, 0xa7, 0x74, 0x89, 0x7f, 0x06, 0x4b, 0xf6, 0x8a, 0xc1, 0xbf, 0x3a, 0x68, 0x94,
	0xe3, 0x60, 0x39, 0x6b, 0x55, 0x1c, 0xa8, 0x43, 0x48, 0x96, 0x2a, 0x75, 0xad, 0x33, 0x26, 0x68,
	0xcf, 0xf9, 0xfc, 0x2b, 0xb8, 0xcb, 0xf1, 0x74, 0xf3, 0xb0, 0x04, 0x73, 0x99, 0x0b, 0x1d, 0xfc,
	0x37, 0xc9, 0xe6, 0x7c, 0x55, 0xf3, 0x3e, 0x56, 0xb5, 0x7d, 0x9b, 0x56, 0xe4, 0x2e, 0x62, 0x12,
	0x71, 0xac, 0x87, 0x51, 0x57, 0xe8, 0xe5, 0xac, 0x05, 0xdd, 0xd8, 0xae, 0xd0, 0xbd, 0x6d, 0x8e,
	0xa7, 0x31, 0x23, 0x1d, 0x33, 0xe0, 0x62, 0xef, 0xe9, 0xa5, 0xe6, 0x3d, 0x7e, 0xbd, 0x9e, 0x3a,
	0x46, 0x7c, 0x39, 0x5f, 0x87, 0x60, 0xb1, 0x0e, 0xc1, 0xe7, 0x3a, 0x04, 0xcf, 0x45, 0xe8, 0x2d,
	0x8a, 0xd0, 0x7b, 0x2f, 0x42, 0xef, 0xa6, 0x45, 0x99, 0x1e, 0xe6, 0xfd, 0x68, 0x20, 0x39, 0x62,
	0x54, 0x30, 0x9d, 0xa2, 0x9f, 0x0d, 0x4c, 0xff, 0xec, 0x40, 0x3f, 0x8c, 0x53, 0xd5, 0x2f, 0x99,
	0x6f, 0x39, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x56, 0x7f, 0x1f, 0xa7, 0x01, 0x00, 0x00,
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
