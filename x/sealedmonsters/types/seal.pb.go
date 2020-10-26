// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sealedmonsters/v1beta/seal.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type MsgSeal struct {
	Id                    string                                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator               github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	SolutionHash          string                                        `protobuf:"bytes,3,opt,name=solutionHash,proto3" json:"solutionHash,omitempty"`
	SolutionScavengerHash string                                        `protobuf:"bytes,4,opt,name=solutionScavengerHash,proto3" json:"solutionScavengerHash,omitempty"`
	Scavenger             string                                        `protobuf:"bytes,5,opt,name=scavenger,proto3" json:"scavenger,omitempty"`
}

func (m *MsgSeal) Reset()         { *m = MsgSeal{} }
func (m *MsgSeal) String() string { return proto.CompactTextString(m) }
func (*MsgSeal) ProtoMessage()    {}
func (*MsgSeal) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b0dee66a7c80092, []int{0}
}
func (m *MsgSeal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSeal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSeal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSeal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSeal.Merge(m, src)
}
func (m *MsgSeal) XXX_Size() int {
	return m.Size()
}
func (m *MsgSeal) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSeal.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSeal proto.InternalMessageInfo

func (m *MsgSeal) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgSeal) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgSeal) GetSolutionHash() string {
	if m != nil {
		return m.SolutionHash
	}
	return ""
}

func (m *MsgSeal) GetSolutionScavengerHash() string {
	if m != nil {
		return m.SolutionScavengerHash
	}
	return ""
}

func (m *MsgSeal) GetScavenger() string {
	if m != nil {
		return m.Scavenger
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgSeal)(nil), "flares.sealedmonsters.v1beta1.MsgSeal")
}

func init() { proto.RegisterFile("sealedmonsters/v1beta/seal.proto", fileDescriptor_0b0dee66a7c80092) }

var fileDescriptor_0b0dee66a7c80092 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0x42, 0x31,
	0x14, 0xc6, 0x29, 0xfe, 0x21, 0x34, 0xc4, 0xa1, 0xd1, 0xe4, 0xc6, 0x68, 0x25, 0x4c, 0x2c, 0xdc,
	0x86, 0xc8, 0x0b, 0xc0, 0x64, 0x62, 0x74, 0x80, 0xcd, 0xad, 0xb4, 0x87, 0xd2, 0x78, 0xb9, 0x25,
	0x3d, 0x05, 0xf5, 0x2d, 0x7c, 0x2c, 0x47, 0x46, 0x27, 0x62, 0xe0, 0x2d, 0x9c, 0x0c, 0xbd, 0xde,
	0x28, 0xc6, 0xa9, 0xed, 0xaf, 0xbf, 0xf6, 0xe4, 0xfb, 0x68, 0x13, 0x41, 0x66, 0xa0, 0x67, 0x2e,
	0xc7, 0x00, 0x1e, 0xc5, 0xb2, 0x3b, 0x86, 0x20, 0xc5, 0x8e, 0xa6, 0x73, 0xef, 0x82, 0x63, 0x97,
	0x93, 0x4c, 0x7a, 0xc0, 0x74, 0x5f, 0x4c, 0x0b, 0xb1, 0x7b, 0x7e, 0x6a, 0x9c, 0x71, 0xd1, 0x14,
	0xbb, 0x5d, 0xf1, 0xa8, 0xb5, 0x26, 0xb4, 0x76, 0x87, 0x66, 0x04, 0x32, 0x63, 0x27, 0xb4, 0x6a,
	0x75, 0x42, 0x9a, 0xa4, 0x5d, 0x1f, 0x56, 0xad, 0x66, 0xb7, 0xb4, 0xa6, 0x3c, 0xc8, 0xe0, 0x7c,
	0x52, 0x6d, 0x92, 0x76, 0x63, 0xd0, 0xfd, 0x5c, 0x5f, 0x75, 0x8c, 0x0d, 0xd3, 0xc5, 0x38, 0x55,
	0x6e, 0x26, 0x94, 0xc3, 0x99, 0xc3, 0xef, 0xa5, 0x83, 0xfa, 0x51, 0x84, 0x97, 0x39, 0x60, 0xda,
	0x57, 0xaa, 0xaf, 0xb5, 0x07, 0xc4, 0x61, 0xf9, 0x03, 0x6b, 0xd1, 0x06, 0xba, 0x6c, 0x11, 0xac,
	0xcb, 0x6f, 0x24, 0x4e, 0x93, 0x83, 0x38, 0x66, 0x8f, 0xb1, 0x1e, 0x3d, 0x2b, 0xcf, 0x23, 0x25,
	0x97, 0x90, 0x1b, 0xf0, 0x51, 0x3e, 0x8c, 0xf2, 0xff, 0x97, 0xec, 0x82, 0xd6, 0xb1, 0x04, 0xc9,
	0x51, 0x34, 0x7f, 0xc0, 0xe0, 0xfe, 0x6d, 0xc3, 0xc9, 0x6a, 0xc3, 0xc9, 0xc7, 0x86, 0x93, 0xd7,
	0x2d, 0xaf, 0xac, 0xb6, 0xbc, 0xf2, 0xbe, 0xe5, 0x95, 0x87, 0xde, 0xaf, 0x24, 0x4f, 0x32, 0x37,
	0x13, 0xb0, 0x73, 0x9b, 0x1b, 0x51, 0xd4, 0x28, 0x9e, 0xc5, 0x9f, 0xc6, 0x63, 0xb6, 0xf1, 0x71,
	0xec, 0xed, 0xfa, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x27, 0x6d, 0x68, 0x07, 0x90, 0x01, 0x00, 0x00,
}

func (m *MsgSeal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSeal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSeal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Scavenger) > 0 {
		i -= len(m.Scavenger)
		copy(dAtA[i:], m.Scavenger)
		i = encodeVarintSeal(dAtA, i, uint64(len(m.Scavenger)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SolutionScavengerHash) > 0 {
		i -= len(m.SolutionScavengerHash)
		copy(dAtA[i:], m.SolutionScavengerHash)
		i = encodeVarintSeal(dAtA, i, uint64(len(m.SolutionScavengerHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SolutionHash) > 0 {
		i -= len(m.SolutionHash)
		copy(dAtA[i:], m.SolutionHash)
		i = encodeVarintSeal(dAtA, i, uint64(len(m.SolutionHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSeal(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintSeal(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSeal(dAtA []byte, offset int, v uint64) int {
	offset -= sovSeal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSeal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSeal(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSeal(uint64(l))
	}
	l = len(m.SolutionHash)
	if l > 0 {
		n += 1 + l + sovSeal(uint64(l))
	}
	l = len(m.SolutionScavengerHash)
	if l > 0 {
		n += 1 + l + sovSeal(uint64(l))
	}
	l = len(m.Scavenger)
	if l > 0 {
		n += 1 + l + sovSeal(uint64(l))
	}
	return n
}

func sovSeal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSeal(x uint64) (n int) {
	return sovSeal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSeal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSeal
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
			return fmt.Errorf("proto: MsgSeal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSeal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeal
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
				return ErrInvalidLengthSeal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSeal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSeal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSeal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SolutionHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeal
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
				return ErrInvalidLengthSeal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSeal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SolutionHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SolutionScavengerHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeal
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
				return ErrInvalidLengthSeal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSeal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SolutionScavengerHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scavenger", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeal
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
				return ErrInvalidLengthSeal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSeal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scavenger = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSeal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSeal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSeal
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
func skipSeal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSeal
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
					return 0, ErrIntOverflowSeal
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
					return 0, ErrIntOverflowSeal
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
				return 0, ErrInvalidLengthSeal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSeal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSeal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSeal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSeal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSeal = fmt.Errorf("proto: unexpected end of group")
)