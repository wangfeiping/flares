// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sealedmonsters/v1beta/reveal.proto

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

type MsgReveal struct {
	Id           string                                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	SolutionHash string                                        `protobuf:"bytes,3,opt,name=solutionHash,proto3" json:"solutionHash,omitempty"`
	Solution     string                                        `protobuf:"bytes,4,opt,name=solution,proto3" json:"solution,omitempty"`
	Scavenger    string                                        `protobuf:"bytes,5,opt,name=scavenger,proto3" json:"scavenger,omitempty"`
}

func (m *MsgReveal) Reset()         { *m = MsgReveal{} }
func (m *MsgReveal) String() string { return proto.CompactTextString(m) }
func (*MsgReveal) ProtoMessage()    {}
func (*MsgReveal) Descriptor() ([]byte, []int) {
	return fileDescriptor_efc875fdea187d7e, []int{0}
}
func (m *MsgReveal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReveal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReveal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReveal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReveal.Merge(m, src)
}
func (m *MsgReveal) XXX_Size() int {
	return m.Size()
}
func (m *MsgReveal) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReveal.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReveal proto.InternalMessageInfo

func (m *MsgReveal) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgReveal) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgReveal) GetSolutionHash() string {
	if m != nil {
		return m.SolutionHash
	}
	return ""
}

func (m *MsgReveal) GetSolution() string {
	if m != nil {
		return m.Solution
	}
	return ""
}

func (m *MsgReveal) GetScavenger() string {
	if m != nil {
		return m.Scavenger
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgReveal)(nil), "flares.sealedmonsters.v1beta1.MsgReveal")
}

func init() {
	proto.RegisterFile("sealedmonsters/v1beta/reveal.proto", fileDescriptor_efc875fdea187d7e)
}

var fileDescriptor_efc875fdea187d7e = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0x3d, 0x4f, 0x02, 0x31,
	0x18, 0xa6, 0xf8, 0x49, 0x43, 0x1c, 0x1a, 0x87, 0x0b, 0xd1, 0x4a, 0x98, 0x58, 0xb8, 0x86, 0xe8,
	0x1f, 0x80, 0xc9, 0xc4, 0xe8, 0xc0, 0xe8, 0x56, 0xda, 0x97, 0xd2, 0x78, 0x5c, 0x49, 0xdf, 0x72,
	0xea, 0xbf, 0xf0, 0x57, 0x19, 0x47, 0x46, 0x27, 0x63, 0xee, 0xfe, 0x85, 0x93, 0xb1, 0x27, 0x2a,
	0x4e, 0xed, 0xf3, 0xd1, 0xa7, 0xef, 0xfb, 0xd0, 0x1e, 0x82, 0xcc, 0x40, 0x2f, 0x5c, 0x8e, 0x01,
	0x3c, 0x8a, 0x62, 0x38, 0x85, 0x20, 0x85, 0x87, 0x02, 0x64, 0x96, 0x2e, 0xbd, 0x0b, 0x8e, 0x9d,
	0xce, 0x32, 0xe9, 0x01, 0xd3, 0x6d, 0x6b, 0x5a, 0x5b, 0x87, 0x9d, 0x63, 0xe3, 0x8c, 0x8b, 0x4e,
	0xf1, 0x75, 0xab, 0x1f, 0xf5, 0x9e, 0x09, 0x6d, 0x5d, 0xa3, 0x99, 0xc4, 0x20, 0x76, 0x44, 0x9b,
	0x56, 0x27, 0xa4, 0x4b, 0xfa, 0xad, 0x49, 0xd3, 0x6a, 0x76, 0x45, 0x0f, 0x94, 0x07, 0x19, 0x9c,
	0x4f, 0x9a, 0x5d, 0xd2, 0x6f, 0x8f, 0x87, 0x1f, 0x6f, 0x67, 0x03, 0x63, 0xc3, 0x7c, 0x35, 0x4d,
	0x95, 0x5b, 0x08, 0xe5, 0x70, 0xe1, 0xf0, 0xfb, 0x18, 0xa0, 0xbe, 0x13, 0xe1, 0x71, 0x09, 0x98,
	0x8e, 0x94, 0x1a, 0x69, 0xed, 0x01, 0x71, 0xb2, 0x49, 0x60, 0x3d, 0xda, 0x46, 0x97, 0xad, 0x82,
	0x75, 0xf9, 0xa5, 0xc4, 0x79, 0xb2, 0x13, 0xbf, 0xd9, 0xe2, 0x58, 0x87, 0x1e, 0x6e, 0x70, 0xb2,
	0x1b, 0xf5, 0x1f, 0xcc, 0x4e, 0x68, 0x0b, 0x95, 0x2c, 0x20, 0x37, 0xe0, 0x93, 0xbd, 0x28, 0xfe,
	0x12, 0xe3, 0x9b, 0x97, 0x92, 0x93, 0x75, 0xc9, 0xc9, 0x7b, 0xc9, 0xc9, 0x53, 0xc5, 0x1b, 0xeb,
	0x8a, 0x37, 0x5e, 0x2b, 0xde, 0xb8, 0xbd, 0xf8, 0x33, 0xef, 0xbd, 0xcc, 0xcd, 0x0c, 0xec, 0xd2,
	0xe6, 0x46, 0xd4, 0x75, 0x89, 0x07, 0xf1, 0xaf, 0xdb, 0xb8, 0xc1, 0x74, 0x3f, 0xf6, 0x73, 0xfe,
	0x19, 0x00, 0x00, 0xff, 0xff, 0x62, 0x55, 0x78, 0x8f, 0x7a, 0x01, 0x00, 0x00,
}

func (m *MsgReveal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReveal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReveal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Scavenger) > 0 {
		i -= len(m.Scavenger)
		copy(dAtA[i:], m.Scavenger)
		i = encodeVarintReveal(dAtA, i, uint64(len(m.Scavenger)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Solution) > 0 {
		i -= len(m.Solution)
		copy(dAtA[i:], m.Solution)
		i = encodeVarintReveal(dAtA, i, uint64(len(m.Solution)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SolutionHash) > 0 {
		i -= len(m.SolutionHash)
		copy(dAtA[i:], m.SolutionHash)
		i = encodeVarintReveal(dAtA, i, uint64(len(m.SolutionHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintReveal(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintReveal(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintReveal(dAtA []byte, offset int, v uint64) int {
	offset -= sovReveal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgReveal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovReveal(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovReveal(uint64(l))
	}
	l = len(m.SolutionHash)
	if l > 0 {
		n += 1 + l + sovReveal(uint64(l))
	}
	l = len(m.Solution)
	if l > 0 {
		n += 1 + l + sovReveal(uint64(l))
	}
	l = len(m.Scavenger)
	if l > 0 {
		n += 1 + l + sovReveal(uint64(l))
	}
	return n
}

func sovReveal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReveal(x uint64) (n int) {
	return sovReveal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgReveal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReveal
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
			return fmt.Errorf("proto: MsgReveal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReveal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReveal
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
				return ErrInvalidLengthReveal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReveal
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
					return ErrIntOverflowReveal
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
				return ErrInvalidLengthReveal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReveal
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
					return ErrIntOverflowReveal
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
				return ErrInvalidLengthReveal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReveal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SolutionHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Solution", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReveal
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
				return ErrInvalidLengthReveal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReveal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Solution = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scavenger", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReveal
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
				return ErrInvalidLengthReveal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReveal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scavenger = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReveal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReveal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReveal
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
func skipReveal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReveal
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
					return 0, ErrIntOverflowReveal
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
					return 0, ErrIntOverflowReveal
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
				return 0, ErrInvalidLengthReveal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReveal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReveal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReveal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReveal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReveal = fmt.Errorf("proto: unexpected end of group")
)