// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nameservice/v1beta/whois.proto

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

type MsgWhois struct {
	Id      string                                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Value   string                                        `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Owner   string                                        `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Price   string                                        `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (m *MsgWhois) Reset()         { *m = MsgWhois{} }
func (m *MsgWhois) String() string { return proto.CompactTextString(m) }
func (*MsgWhois) ProtoMessage()    {}
func (*MsgWhois) Descriptor() ([]byte, []int) {
	return fileDescriptor_04c88ce213062881, []int{0}
}
func (m *MsgWhois) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWhois) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWhois.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWhois) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWhois.Merge(m, src)
}
func (m *MsgWhois) XXX_Size() int {
	return m.Size()
}
func (m *MsgWhois) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWhois.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWhois proto.InternalMessageInfo

func (m *MsgWhois) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgWhois) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgWhois) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *MsgWhois) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgWhois) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgWhois)(nil), "flares.nameservice.v1beta1.MsgWhois")
}

func init() { proto.RegisterFile("nameservice/v1beta/whois.proto", fileDescriptor_04c88ce213062881) }

var fileDescriptor_04c88ce213062881 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x4e, 0xeb, 0x30,
	0x18, 0xc5, 0xe3, 0xdc, 0x5b, 0xfe, 0x44, 0x88, 0x21, 0xea, 0x10, 0x75, 0x30, 0x15, 0x53, 0x97,
	0xc6, 0x8a, 0x78, 0x82, 0x76, 0xad, 0x58, 0xba, 0x20, 0xb1, 0x39, 0xce, 0x57, 0xd7, 0xa2, 0x89,
	0x23, 0x7f, 0x6e, 0x02, 0x6f, 0xc1, 0x6b, 0xf0, 0x26, 0x8c, 0x1d, 0x99, 0x10, 0x4a, 0xde, 0x82,
	0x09, 0xc5, 0x06, 0x29, 0x93, 0x7d, 0x7e, 0x3a, 0x3e, 0x3a, 0x3e, 0x11, 0xad, 0x78, 0x09, 0x08,
	0xa6, 0x51, 0x02, 0x58, 0x93, 0xe5, 0x60, 0x39, 0x6b, 0xf7, 0x5a, 0x61, 0x5a, 0x1b, 0x6d, 0x75,
	0x3c, 0xdb, 0x1d, 0xb8, 0x01, 0x4c, 0x47, 0xb6, 0xd4, 0xdb, 0xb2, 0xd9, 0x54, 0x6a, 0xa9, 0x9d,
	0x8d, 0x0d, 0x37, 0xff, 0xe2, 0xf6, 0x8d, 0x44, 0x17, 0xf7, 0x28, 0x1f, 0x86, 0x90, 0xf8, 0x3a,
	0x0a, 0x55, 0x91, 0x90, 0x39, 0x59, 0x5c, 0x6e, 0x43, 0x55, 0xc4, 0x9b, 0xe8, 0x5c, 0x18, 0xe0,
	0x56, 0x9b, 0x24, 0x9c, 0x93, 0xc5, 0xd5, 0x3a, 0xfb, 0xfe, 0xbc, 0x59, 0x4a, 0x65, 0xf7, 0xc7,
	0x3c, 0x15, 0xba, 0x64, 0x42, 0x63, 0xa9, 0xf1, 0xf7, 0x58, 0x62, 0xf1, 0xc4, 0xec, 0x4b, 0x0d,
	0x98, 0xae, 0x84, 0x58, 0x15, 0x85, 0x01, 0xc4, 0xed, 0x5f, 0x42, 0x3c, 0x8d, 0x26, 0x0d, 0x3f,
	0x1c, 0x21, 0xf9, 0xe7, 0xf2, 0xbd, 0x18, 0xa8, 0x6e, 0x2b, 0x30, 0xc9, 0x7f, 0x4f, 0x9d, 0x18,
	0x68, 0x6d, 0x94, 0x80, 0x64, 0xe2, 0xa9, 0x13, 0xeb, 0xcd, 0x7b, 0x47, 0xc9, 0xa9, 0xa3, 0xe4,
	0xab, 0xa3, 0xe4, 0xb5, 0xa7, 0xc1, 0xa9, 0xa7, 0xc1, 0x47, 0x4f, 0x83, 0xc7, 0x6c, 0xd4, 0xa9,
	0xe5, 0x95, 0xdc, 0x81, 0xaa, 0x55, 0x25, 0x99, 0x9f, 0x83, 0x3d, 0xb3, 0xf1, 0x6e, 0xae, 0x62,
	0x7e, 0xe6, 0xfe, 0x7f, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x99, 0x4d, 0x8b, 0x53, 0x01,
	0x00, 0x00,
}

func (m *MsgWhois) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWhois) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWhois) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhois(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhois(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgWhois) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	return n
}

func sovWhois(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhois(x uint64) (n int) {
	return sovWhois(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgWhois) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhois
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
			return fmt.Errorf("proto: MsgWhois: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWhois: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
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
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
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
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhois(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWhois
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWhois
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
func skipWhois(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhois
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
					return 0, ErrIntOverflowWhois
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
					return 0, ErrIntOverflowWhois
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
				return 0, ErrInvalidLengthWhois
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhois
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhois
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhois        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhois          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhois = fmt.Errorf("proto: unexpected end of group")
)
