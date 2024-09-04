// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/fundraising/v1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	// auction_creation_fee specifies the fee for auction creation.
	// this prevents from spamming attack and it is collected in the community
	// pool
	AuctionCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=auction_creation_fee,json=auctionCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"auction_creation_fee"`
	// place_bid_fee specifies the fee for placing a bid for an auction.
	// this prevents from spamming attack and it is collected in the community
	// pool
	PlaceBidFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=place_bid_fee,json=placeBidFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"place_bid_fee"`
	// extended_period specifies the extended period that determines how long
	// the extended auction round lasts
	ExtendedPeriod uint32 `protobuf:"varint,3,opt,name=extended_period,json=extendedPeriod,proto3" json:"extended_period,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3acf29630e7f59, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAuctionCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.AuctionCreationFee
	}
	return nil
}

func (m *Params) GetPlaceBidFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PlaceBidFee
	}
	return nil
}

func (m *Params) GetExtendedPeriod() uint32 {
	if m != nil {
		return m.ExtendedPeriod
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "modules.fundraising.v1.Params")
}

func init() {
	proto.RegisterFile("modules/fundraising/v1/params.proto", fileDescriptor_9b3acf29630e7f59)
}

var fileDescriptor_9b3acf29630e7f59 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xb1, 0x6e, 0xe2, 0x40,
	0x10, 0xf5, 0x82, 0x44, 0x61, 0x8e, 0x3b, 0x9d, 0x85, 0x4e, 0x40, 0x61, 0xd0, 0x5d, 0x71, 0x1c,
	0x12, 0x5e, 0xf9, 0xae, 0xbb, 0x12, 0x24, 0xd2, 0x22, 0xca, 0x34, 0xd6, 0xda, 0x3b, 0x38, 0xab,
	0xe0, 0x5d, 0xcb, 0x6b, 0x23, 0xf8, 0x80, 0x34, 0xa9, 0x22, 0xa5, 0x4b, 0x95, 0x32, 0x4a, 0xc5,
	0x67, 0x50, 0x52, 0xa6, 0x4a, 0x22, 0x28, 0xc8, 0x37, 0xa4, 0x8a, 0xbc, 0x5e, 0x24, 0xf8, 0x81,
	0x34, 0xf6, 0xcc, 0x9b, 0xb7, 0xf3, 0xde, 0xee, 0x8c, 0xf9, 0x2b, 0x12, 0x34, 0x9b, 0x81, 0xc4,
	0xd3, 0x8c, 0xd3, 0x84, 0x30, 0xc9, 0x78, 0x88, 0xe7, 0x2e, 0x8e, 0x49, 0x42, 0x22, 0xe9, 0xc4,
	0x89, 0x48, 0x85, 0xf5, 0x43, 0x93, 0x9c, 0x23, 0x92, 0x33, 0x77, 0x5b, 0xdf, 0x49, 0xc4, 0xb8,
	0xc0, 0xea, 0x5b, 0x50, 0x5b, 0x76, 0x20, 0x64, 0x24, 0x24, 0xf6, 0x89, 0x04, 0x3c, 0x77, 0x7d,
	0x48, 0x89, 0x8b, 0x03, 0xc1, 0xb8, 0xae, 0x37, 0x8b, 0xba, 0xa7, 0x32, 0x5c, 0x24, 0xba, 0x54,
	0x0f, 0x45, 0x28, 0x0a, 0x3c, 0x8f, 0x0a, 0xf4, 0xe7, 0x7b, 0xc9, 0xac, 0x8c, 0x95, 0x19, 0xeb,
	0x16, 0x99, 0x75, 0x92, 0x05, 0x29, 0x13, 0xdc, 0x0b, 0x12, 0x20, 0x2a, 0x98, 0x02, 0x34, 0x50,
	0xa7, 0xdc, 0xad, 0xfe, 0x6d, 0x3a, 0xba, 0x5d, 0xae, 0xed, 0x68, 0x6d, 0x67, 0x28, 0x18, 0x1f,
	0x8c, 0xd6, 0xcf, 0x6d, 0xe3, 0xf1, 0xa5, 0xdd, 0x0d, 0x59, 0x7a, 0x91, 0xf9, 0x4e, 0x20, 0x22,
	0xad, 0xad, 0x7f, 0x7d, 0x49, 0x2f, 0x71, 0xba, 0x8c, 0x41, 0xaa, 0x03, 0xf2, 0x6e, 0xbf, 0xea,
	0x7d, 0x99, 0x41, 0x48, 0x82, 0xa5, 0x97, 0xbb, 0x97, 0x0f, 0xfb, 0x55, 0x0f, 0x4d, 0x2c, 0x2d,
	0x3f, 0xd4, 0xea, 0x23, 0x00, 0xeb, 0x0a, 0x99, 0xb5, 0x78, 0x46, 0x02, 0xf0, 0x7c, 0x46, 0x95,
	0x9d, 0xd2, 0x67, 0xd9, 0xa9, 0x2a, 0xdd, 0x01, 0xa3, 0xb9, 0x8f, 0xdf, 0xe6, 0x37, 0x58, 0xa4,
	0xc0, 0x29, 0x50, 0x2f, 0x86, 0x84, 0x09, 0xda, 0x28, 0x77, 0x50, 0xb7, 0x36, 0xf9, 0x7a, 0x80,
	0xc7, 0x0a, 0xfd, 0xff, 0xe7, 0xed, 0xbe, 0x8d, 0xae, 0xf7, 0xab, 0x5e, 0xe7, 0x78, 0xe6, 0x8b,
	0x93, 0x0d, 0x28, 0x5e, 0x7c, 0x70, 0xb6, 0xde, 0xda, 0x68, 0xb3, 0xb5, 0xd1, 0xeb, 0xd6, 0x46,
	0x37, 0x3b, 0xdb, 0xd8, 0xec, 0x6c, 0xe3, 0x69, 0x67, 0x1b, 0xe7, 0xfd, 0x23, 0xeb, 0x2c, 0xe4,
	0x2c, 0x05, 0x7c, 0xd8, 0xa4, 0xd3, 0x4e, 0xea, 0x16, 0x7e, 0x45, 0x0d, 0xf3, 0xdf, 0x47, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x9e, 0x7f, 0xd4, 0xc7, 0x6f, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.AuctionCreationFee) != len(that1.AuctionCreationFee) {
		return false
	}
	for i := range this.AuctionCreationFee {
		if !this.AuctionCreationFee[i].Equal(&that1.AuctionCreationFee[i]) {
			return false
		}
	}
	if len(this.PlaceBidFee) != len(that1.PlaceBidFee) {
		return false
	}
	for i := range this.PlaceBidFee {
		if !this.PlaceBidFee[i].Equal(&that1.PlaceBidFee[i]) {
			return false
		}
	}
	if this.ExtendedPeriod != that1.ExtendedPeriod {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExtendedPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ExtendedPeriod))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PlaceBidFee) > 0 {
		for iNdEx := len(m.PlaceBidFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PlaceBidFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.AuctionCreationFee) > 0 {
		for iNdEx := len(m.AuctionCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AuctionCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AuctionCreationFee) > 0 {
		for _, e := range m.AuctionCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.PlaceBidFee) > 0 {
		for _, e := range m.PlaceBidFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.ExtendedPeriod != 0 {
		n += 1 + sovParams(uint64(m.ExtendedPeriod))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuctionCreationFee = append(m.AuctionCreationFee, types.Coin{})
			if err := m.AuctionCreationFee[len(m.AuctionCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlaceBidFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlaceBidFee = append(m.PlaceBidFee, types.Coin{})
			if err := m.PlaceBidFee[len(m.PlaceBidFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtendedPeriod", wireType)
			}
			m.ExtendedPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExtendedPeriod |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
