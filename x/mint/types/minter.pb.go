// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/mint/v1/minter.proto

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

// Minter represents the minting state.
type Minter struct {
	// inflation defines the current annual inflation rate
	Inflation cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflation"`
	// annualProvisions defines the current annual expected provisions
	AnnualProvisions cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=annualProvisions,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"annualProvisions"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fa832c0616ce37e, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

type WeightedAddress struct {
	Address string                      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Weight  cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=weight,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"weight"`
}

func (m *WeightedAddress) Reset()         { *m = WeightedAddress{} }
func (m *WeightedAddress) String() string { return proto.CompactTextString(m) }
func (*WeightedAddress) ProtoMessage()    {}
func (*WeightedAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fa832c0616ce37e, []int{1}
}
func (m *WeightedAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WeightedAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WeightedAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WeightedAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedAddress.Merge(m, src)
}
func (m *WeightedAddress) XXX_Size() int {
	return m.Size()
}
func (m *WeightedAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedAddress.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedAddress proto.InternalMessageInfo

func (m *WeightedAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type DistributionProportions struct {
	// staking defines the proportion of the minted minted_denom that is to be
	// allocated as staking rewards.
	Staking cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=staking,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"staking"`
	// fundedAddresses defines the proportion of the minted minted_denom that is
	// to the set of funded addresses.
	FundedAddresses cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=fundedAddresses,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"fundedAddresses"`
	// communityPool defines the proportion of the minted minted_denom that is
	// to be allocated to the community pool.
	CommunityPool cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=communityPool,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"communityPool"`
}

func (m *DistributionProportions) Reset()         { *m = DistributionProportions{} }
func (m *DistributionProportions) String() string { return proto.CompactTextString(m) }
func (*DistributionProportions) ProtoMessage()    {}
func (*DistributionProportions) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fa832c0616ce37e, []int{2}
}
func (m *DistributionProportions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributionProportions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributionProportions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributionProportions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionProportions.Merge(m, src)
}
func (m *DistributionProportions) XXX_Size() int {
	return m.Size()
}
func (m *DistributionProportions) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionProportions.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionProportions proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Minter)(nil), "modules.mint.v1.Minter")
	proto.RegisterType((*WeightedAddress)(nil), "modules.mint.v1.WeightedAddress")
	proto.RegisterType((*DistributionProportions)(nil), "modules.mint.v1.DistributionProportions")
}

func init() { proto.RegisterFile("modules/mint/v1/minter.proto", fileDescriptor_7fa832c0616ce37e) }

var fileDescriptor_7fa832c0616ce37e = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x93, 0x5e, 0x68, 0xe9, 0xc0, 0xa5, 0x97, 0xa1, 0x70, 0x63, 0x95, 0x54, 0xba, 0x10,
	0x37, 0x4d, 0xa8, 0xbe, 0x80, 0x96, 0x6e, 0x44, 0xc5, 0x52, 0x17, 0x05, 0x45, 0x24, 0x4d, 0xa6,
	0xd3, 0xa1, 0xc9, 0x9c, 0x32, 0x33, 0xa9, 0xf6, 0x2d, 0xfa, 0x00, 0x3e, 0x46, 0x17, 0x3e, 0x42,
	0x97, 0xa5, 0x2b, 0x71, 0x51, 0xa4, 0x7d, 0x11, 0xc9, 0x9f, 0x2a, 0xea, 0x2e, 0xae, 0xe6, 0xcc,
	0x39, 0x1f, 0xbf, 0x99, 0xf3, 0xf1, 0xa1, 0xbd, 0x00, 0xbc, 0xd0, 0x27, 0xd2, 0x0e, 0x18, 0x57,
	0xf6, 0xb8, 0x11, 0x9f, 0x44, 0x58, 0x23, 0x01, 0x0a, 0x70, 0x29, 0x9d, 0x5a, 0x51, 0xd7, 0x1a,
	0x37, 0x2a, 0x3b, 0x2e, 0xc8, 0x00, 0xe4, 0x7d, 0x3c, 0xb6, 0x93, 0x4b, 0xa2, 0xad, 0x94, 0x29,
	0x50, 0x48, 0xfa, 0x51, 0x95, 0x74, 0x6b, 0xcf, 0x3a, 0xca, 0x5f, 0xc6, 0x48, 0x7c, 0x85, 0x8a,
	0x8c, 0xf7, 0x7d, 0x47, 0x31, 0xe0, 0x86, 0xbe, 0xaf, 0x1f, 0x16, 0x9b, 0x8d, 0xf9, 0xaa, 0xaa,
	0xbd, 0xae, 0xaa, 0xbb, 0x09, 0x49, 0x7a, 0x43, 0x8b, 0x81, 0x1d, 0x38, 0x6a, 0x60, 0x5d, 0x10,
	0xea, 0xb8, 0x93, 0x16, 0x71, 0x97, 0xb3, 0x3a, 0x4a, 0x1f, 0x6a, 0x11, 0xb7, 0xf3, 0xc9, 0xc0,
	0x77, 0xe8, 0x9f, 0xc3, 0x79, 0xe8, 0xf8, 0x6d, 0x01, 0x63, 0x26, 0x19, 0x70, 0x69, 0xe4, 0xb2,
	0x72, 0x7f, 0xa0, 0x6a, 0x53, 0x1d, 0x95, 0xba, 0x84, 0xd1, 0x81, 0x22, 0xde, 0xa9, 0xe7, 0x09,
	0x22, 0x25, 0x3e, 0x42, 0x05, 0x27, 0x29, 0xd3, 0x0d, 0x8c, 0xe5, 0xac, 0x5e, 0x4e, 0x31, 0xa9,
	0xe8, 0x5a, 0x09, 0xc6, 0x69, 0x67, 0x2b, 0xc4, 0x67, 0x28, 0xff, 0x10, 0x63, 0xb2, 0x7f, 0x2e,
	0x05, 0xd4, 0x9e, 0x72, 0xe8, 0x7f, 0x8b, 0x49, 0x25, 0x58, 0x2f, 0x8c, 0x2c, 0x68, 0x0b, 0x18,
	0x81, 0x88, 0x2a, 0x89, 0xcf, 0x51, 0x41, 0x2a, 0x67, 0xc8, 0x38, 0xcd, 0x6e, 0xee, 0x96, 0x80,
	0x6f, 0x51, 0xa9, 0x1f, 0x72, 0xef, 0x63, 0x71, 0xf2, 0x0b, 0x67, 0xbf, 0x93, 0x70, 0x17, 0xfd,
	0x75, 0x21, 0x08, 0x42, 0xce, 0xd4, 0xa4, 0x0d, 0xe0, 0x1b, 0x7f, 0xb2, 0xa2, 0xbf, 0x72, 0x9a,
	0x27, 0xf3, 0xb5, 0xa9, 0x2f, 0xd6, 0xa6, 0xfe, 0xb6, 0x36, 0xf5, 0xe9, 0xc6, 0xd4, 0x16, 0x1b,
	0x53, 0x7b, 0xd9, 0x98, 0xda, 0xcd, 0x01, 0x65, 0x6a, 0x10, 0xf6, 0x2c, 0x17, 0x02, 0x9b, 0x51,
	0xce, 0x14, 0xb1, 0xb7, 0xc1, 0x7f, 0x4c, 0xa2, 0xaf, 0x26, 0x23, 0x22, 0x7b, 0xf9, 0x38, 0xb5,
	0xc7, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x32, 0x1a, 0x61, 0xce, 0x17, 0x03, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *WeightedAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WeightedAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WeightedAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMinter(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DistributionProportions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributionProportions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistributionProportions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CommunityPool.Size()
		i -= size
		if _, err := m.CommunityPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.FundedAddresses.Size()
		i -= size
		if _, err := m.FundedAddresses.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Staking.Size()
		i -= size
		if _, err := m.Staking.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMinter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMinter(dAtA []byte, offset int, v uint64) int {
	offset -= sovMinter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovMinter(uint64(l))
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovMinter(uint64(l))
	return n
}

func (m *WeightedAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMinter(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovMinter(uint64(l))
	return n
}

func (m *DistributionProportions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Staking.Size()
	n += 1 + l + sovMinter(uint64(l))
	l = m.FundedAddresses.Size()
	n += 1 + l + sovMinter(uint64(l))
	l = m.CommunityPool.Size()
	n += 1 + l + sovMinter(uint64(l))
	return n
}

func sovMinter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMinter(x uint64) (n int) {
	return sovMinter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinter
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinter
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
func (m *WeightedAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinter
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
			return fmt.Errorf("proto: WeightedAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WeightedAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinter
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
func (m *DistributionProportions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinter
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
			return fmt.Errorf("proto: DistributionProportions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributionProportions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staking", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Staking.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundedAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FundedAddresses.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinter
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
				return ErrInvalidLengthMinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommunityPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinter
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
func skipMinter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMinter
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
					return 0, ErrIntOverflowMinter
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
					return 0, ErrIntOverflowMinter
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
				return 0, ErrInvalidLengthMinter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMinter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMinter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMinter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMinter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMinter = fmt.Errorf("proto: unexpected end of group")
)
