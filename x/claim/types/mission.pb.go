// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/claim/v1/mission.proto

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

type Mission struct {
	MissionId   uint64                      `protobuf:"varint,1,opt,name=mission_id,json=missionId,proto3" json:"mission_id,omitempty"`
	Description string                      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight      cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=weight,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"weight"`
}

func (m *Mission) Reset()         { *m = Mission{} }
func (m *Mission) String() string { return proto.CompactTextString(m) }
func (*Mission) ProtoMessage()    {}
func (*Mission) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd128c86affad, []int{0}
}
func (m *Mission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mission.Merge(m, src)
}
func (m *Mission) XXX_Size() int {
	return m.Size()
}
func (m *Mission) XXX_DiscardUnknown() {
	xxx_messageInfo_Mission.DiscardUnknown(m)
}

var xxx_messageInfo_Mission proto.InternalMessageInfo

func (m *Mission) GetMissionId() uint64 {
	if m != nil {
		return m.MissionId
	}
	return 0
}

func (m *Mission) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Mission)(nil), "modules.claim.v1.Mission")
}

func init() { proto.RegisterFile("modules/claim/v1/mission.proto", fileDescriptor_b88fd128c86affad) }

var fileDescriptor_b88fd128c86affad = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0xce, 0x49, 0xcc, 0xcc, 0xd5, 0x2f, 0x33, 0xd4, 0xcf, 0xcd, 0x2c,
	0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0xca, 0xeb, 0x81,
	0xe5, 0xf5, 0xca, 0x0c, 0xa5, 0x24, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xe3, 0xc1, 0xf2, 0xfa,
	0x10, 0x0e, 0x44, 0xb1, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x44, 0x1c, 0xc4, 0x82, 0x88, 0x2a,
	0x4d, 0x65, 0xe4, 0x62, 0xf7, 0x85, 0x18, 0x2a, 0x24, 0xcb, 0xc5, 0x05, 0x35, 0x3f, 0x3e, 0x33,
	0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x25, 0x88, 0x13, 0x2a, 0xe2, 0x99, 0x22, 0xa4, 0xc0, 0xc5,
	0x9d, 0x92, 0x5a, 0x9c, 0x5c, 0x94, 0x59, 0x50, 0x92, 0x99, 0x9f, 0x27, 0xc1, 0xa4, 0xc0, 0xa8,
	0xc1, 0x19, 0x84, 0x2c, 0x24, 0xe4, 0xc9, 0xc5, 0x56, 0x9e, 0x9a, 0x99, 0x9e, 0x51, 0x22, 0xc1,
	0x0c, 0x92, 0x74, 0x32, 0x3c, 0x71, 0x4f, 0x9e, 0xe1, 0xd6, 0x3d, 0x79, 0x69, 0x88, 0x43, 0x8a,
	0x53, 0xb2, 0xf5, 0x32, 0xf3, 0xf5, 0x73, 0x13, 0x4b, 0x32, 0xf4, 0x7c, 0x52, 0xd3, 0x13, 0x93,
	0x2b, 0x5d, 0x52, 0x93, 0x2f, 0x6d, 0xd1, 0xe5, 0x82, 0xba, 0xd3, 0x25, 0x35, 0x39, 0x08, 0x6a,
	0x80, 0x93, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xa9, 0xa7, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x67, 0xa6, 0xe7, 0x65, 0x96, 0xa4, 0xea,
	0xc3, 0x82, 0xa9, 0x02, 0x1a, 0x50, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x1f, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x91, 0x0c, 0x5e, 0x46, 0x01, 0x00, 0x00,
}

func (m *Mission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintMission(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintMission(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if m.MissionId != 0 {
		i = encodeVarintMission(dAtA, i, uint64(m.MissionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMission(dAtA []byte, offset int, v uint64) int {
	offset -= sovMission(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MissionId != 0 {
		n += 1 + sovMission(uint64(m.MissionId))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMission(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovMission(uint64(l))
	return n
}

func sovMission(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMission(x uint64) (n int) {
	return sovMission(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMission
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
			return fmt.Errorf("proto: Mission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissionId", wireType)
			}
			m.MissionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMission
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
				return ErrInvalidLengthMission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMission
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
				return ErrInvalidLengthMission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMission
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
			skippy, err := skipMission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMission
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
func skipMission(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMission
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
					return 0, ErrIntOverflowMission
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
					return 0, ErrIntOverflowMission
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
				return 0, ErrInvalidLengthMission
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMission
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMission
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMission        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMission          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMission = fmt.Errorf("proto: unexpected end of group")
)
