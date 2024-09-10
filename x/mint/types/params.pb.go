// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/mint/v1/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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
	// mintDenom defines the type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mintDenom,proto3" json:"mintDenom,omitempty"`
	// inflationRateChange defines the maximum annual change in inflation rate
	InflationRateChange cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=inflationRateChange,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflationRateChange"`
	// inflationMax defines the maximum inflation rate
	InflationMax cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=inflationMax,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflationMax"`
	// inflationMin defines the minimum inflation rate
	InflationMin cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=inflationMin,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"inflationMin"`
	// goalBonded defines the goal of percent bonded atoms
	GoalBonded cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=goalBonded,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"goalBonded"`
	// blocksPerYear defines the expected blocks per year
	BlocksPerYear uint64 `protobuf:"varint,6,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty"`
	// distributionProportions defines the proportion of the minted denom
	DistributionProportions DistributionProportions `protobuf:"bytes,7,opt,name=distributionProportions,proto3" json:"distributionProportions"`
	// fundedAddresses defines the list of funded addresses
	FundedAddresses []WeightedAddress `protobuf:"bytes,8,rep,name=fundedAddresses,proto3" json:"fundedAddresses"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b999d29ac522126, []int{0}
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

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func (m *Params) GetDistributionProportions() DistributionProportions {
	if m != nil {
		return m.DistributionProportions
	}
	return DistributionProportions{}
}

func (m *Params) GetFundedAddresses() []WeightedAddress {
	if m != nil {
		return m.FundedAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "modules.mint.v1.Params")
}

func init() { proto.RegisterFile("modules/mint/v1/params.proto", fileDescriptor_6b999d29ac522126) }

var fileDescriptor_6b999d29ac522126 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0x86, 0x63, 0x9a, 0x06, 0x7a, 0x05, 0x45, 0x18, 0x10, 0x26, 0x54, 0xae, 0xc5, 0x50, 0x59,
	0x95, 0xb0, 0xd5, 0x22, 0x31, 0x30, 0x41, 0xc8, 0x08, 0x52, 0xe4, 0x01, 0x44, 0x97, 0xe8, 0x62,
	0x7f, 0x3d, 0x9f, 0x9a, 0xbb, 0xcf, 0xba, 0xbb, 0x54, 0xcd, 0x5f, 0x60, 0xe2, 0x27, 0x30, 0x32,
	0x76, 0xe0, 0x47, 0x74, 0xac, 0x98, 0x10, 0x43, 0x85, 0x92, 0xa1, 0x7f, 0x03, 0x5d, 0xce, 0x14,
	0x9a, 0x96, 0x29, 0x2c, 0xf6, 0xdd, 0xfb, 0x7e, 0xf7, 0x7c, 0xef, 0xf0, 0x92, 0x0d, 0x81, 0xc5,
	0x78, 0x04, 0x3a, 0x15, 0x5c, 0x9a, 0xf4, 0x70, 0x27, 0xad, 0xa8, 0xa2, 0x42, 0x27, 0x95, 0x42,
	0x83, 0x7e, 0xbb, 0x76, 0x13, 0xeb, 0x26, 0x87, 0x3b, 0x9d, 0xbb, 0x54, 0x70, 0x89, 0xe9, 0xfc,
	0xeb, 0x66, 0x3a, 0x8f, 0x72, 0xd4, 0x02, 0xf5, 0x60, 0x7e, 0x4b, 0xdd, 0xa5, 0xb6, 0xee, 0x33,
	0x64, 0xe8, 0x74, 0x7b, 0xaa, 0xd5, 0x2b, 0x2b, 0xed, 0x1f, 0x94, 0x73, 0x9f, 0x7c, 0x5e, 0x25,
	0xad, 0xfe, 0x3c, 0x83, 0xbf, 0x41, 0xd6, 0xac, 0xd5, 0x03, 0x89, 0x22, 0xf0, 0x22, 0x2f, 0x5e,
	0xcb, 0xfe, 0x08, 0x7e, 0x49, 0xee, 0x71, 0xb9, 0x3f, 0xa2, 0x86, 0xa3, 0xcc, 0xa8, 0x81, 0xd7,
	0x25, 0x95, 0x0c, 0x82, 0x1b, 0x76, 0xae, 0xfb, 0xfc, 0xe4, 0x6c, 0xb3, 0xf1, 0xe3, 0x6c, 0xf3,
	0xb1, 0xcb, 0xa3, 0x8b, 0x83, 0x84, 0x63, 0x2a, 0xa8, 0x29, 0x93, 0x37, 0xc0, 0x68, 0x3e, 0xe9,
	0x41, 0xfe, 0xed, 0xeb, 0x53, 0x52, 0xc7, 0xed, 0x41, 0xfe, 0xe5, 0xfc, 0x78, 0xdb, 0xcb, 0xae,
	0x43, 0xfa, 0x7b, 0xe4, 0xf6, 0x85, 0xfc, 0x96, 0x1e, 0x05, 0x2b, 0x4b, 0xad, 0xb8, 0xc4, 0xba,
	0xcc, 0xe6, 0x32, 0x68, 0xfe, 0x2f, 0x36, 0x97, 0xfe, 0x3b, 0x42, 0x18, 0xd2, 0x51, 0x17, 0x65,
	0x01, 0x45, 0xb0, 0xba, 0x14, 0xf9, 0x2f, 0x92, 0xbf, 0x45, 0xda, 0xc3, 0x11, 0xe6, 0x07, 0x7a,
	0x50, 0x81, 0x1a, 0x4c, 0x80, 0xaa, 0xa0, 0x15, 0x79, 0x71, 0x33, 0xbb, 0xe3, 0xe4, 0x3e, 0xa8,
	0x0f, 0x40, 0x95, 0x5f, 0x92, 0x87, 0x05, 0xd7, 0x46, 0xf1, 0xe1, 0xd8, 0x46, 0xea, 0x2b, 0xac,
	0x50, 0xd9, 0x93, 0x0e, 0x6e, 0x46, 0x5e, 0xbc, 0xbe, 0x1b, 0x27, 0x0b, 0xfd, 0x4a, 0x7a, 0xd7,
	0xcf, 0x77, 0x9b, 0x36, 0x76, 0xf6, 0x2f, 0x9c, 0xdf, 0x27, 0xed, 0xfd, 0xb1, 0xcd, 0xf6, 0xaa,
	0x28, 0x14, 0x68, 0x0d, 0x3a, 0xb8, 0x15, 0xad, 0xc4, 0xeb, 0xbb, 0xd1, 0x95, 0x0d, 0xef, 0x81,
	0xb3, 0xd2, 0x5c, 0x4c, 0xd6, 0xe4, 0xc5, 0xe7, 0x2f, 0x3a, 0x1f, 0xcf, 0x8f, 0xb7, 0x1f, 0xfc,
	0x6e, 0xea, 0x91, 0xeb, 0xaa, 0xeb, 0x65, 0xf7, 0xe5, 0xc9, 0x34, 0xf4, 0x4e, 0xa7, 0xa1, 0xf7,
	0x73, 0x1a, 0x7a, 0x9f, 0x66, 0x61, 0xe3, 0x74, 0x16, 0x36, 0xbe, 0xcf, 0xc2, 0xc6, 0xde, 0x16,
	0xe3, 0xa6, 0x1c, 0x0f, 0x93, 0x1c, 0x45, 0xca, 0x99, 0xe4, 0x06, 0xd2, 0x05, 0x84, 0x99, 0x54,
	0xa0, 0x87, 0xad, 0x79, 0xd7, 0x9f, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x95, 0xc6, 0x29, 0x61,
	0x7e, 0x03, 0x00, 0x00,
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
	if len(m.FundedAddresses) > 0 {
		for iNdEx := len(m.FundedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FundedAddresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	{
		size, err := m.DistributionProportions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.BlocksPerYear != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.GoalBonded.Size()
		i -= size
		if _, err := m.GoalBonded.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.InflationMin.Size()
		i -= size
		if _, err := m.InflationMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InflationMax.Size()
		i -= size
		if _, err := m.InflationMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.InflationRateChange.Size()
		i -= size
		if _, err := m.InflationRateChange.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
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
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.InflationRateChange.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.InflationMax.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.InflationMin.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.GoalBonded.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.BlocksPerYear != 0 {
		n += 1 + sovParams(uint64(m.BlocksPerYear))
	}
	l = m.DistributionProportions.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.FundedAddresses) > 0 {
		for _, e := range m.FundedAddresses {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRateChange", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRateChange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoalBonded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GoalBonded.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionProportions", wireType)
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
			if err := m.DistributionProportions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundedAddresses", wireType)
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
			m.FundedAddresses = append(m.FundedAddresses, WeightedAddress{})
			if err := m.FundedAddresses[len(m.FundedAddresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
