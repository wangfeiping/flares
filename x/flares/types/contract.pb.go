// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flares/v1beta/contract.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type MsgContract struct {
	Id             string                                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator        github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Key            string                                        `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Receiver       string                                        `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Accept         string                                        `protobuf:"bytes,5,opt,name=accept,proto3" json:"accept,omitempty"`
	DurationHeight int32                                         `protobuf:"varint,6,opt,name=durationHeight,proto3" json:"durationHeight,omitempty"`
	Bottom         string                                        `protobuf:"bytes,7,opt,name=bottom,proto3" json:"bottom,omitempty"`
	Height         uint64                                        `protobuf:"fixed64,8,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *MsgContract) IsAuctions() bool { return m.DurationHeight > 0 }

func (m *MsgContract) Reset()         { *m = MsgContract{} }
func (m *MsgContract) String() string { return proto.CompactTextString(m) }
func (*MsgContract) ProtoMessage()    {}
func (*MsgContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdae3fe79527cbf4, []int{0}
}
func (m *MsgContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgContract.Merge(m, src)
}
func (m *MsgContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgContract proto.InternalMessageInfo

func (m *MsgContract) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgContract) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgContract) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MsgContract) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgContract) GetAccept() string {
	if m != nil {
		return m.Accept
	}
	return ""
}

func (m *MsgContract) GetDurationHeight() int32 {
	if m != nil {
		return m.DurationHeight
	}
	return 0
}

func (m *MsgContract) GetBottom() string {
	if m != nil {
		return m.Bottom
	}
	return ""
}

func (m *MsgContract) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgContract)(nil), "flares.flares.v1beta1.MsgContract")
}

func init() { proto.RegisterFile("flares/v1beta/contract.proto", fileDescriptor_bdae3fe79527cbf4) }

var fileDescriptor_bdae3fe79527cbf4 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0xcd, 0x4a, 0xc3, 0x40,
	0x18, 0xec, 0xa6, 0x36, 0xad, 0xab, 0x14, 0x59, 0x54, 0x96, 0x22, 0x31, 0x78, 0x90, 0x1c, 0x6c,
	0x42, 0xf1, 0x09, 0x5a, 0x41, 0x04, 0xf1, 0x92, 0xa3, 0xb7, 0xcd, 0x66, 0xbb, 0x5d, 0x6a, 0xb3,
	0x61, 0xf7, 0x6b, 0xb5, 0x6f, 0xe0, 0xd1, 0xc7, 0xf2, 0xd8, 0xa3, 0x27, 0x91, 0xe6, 0x2d, 0x3c,
	0x49, 0x7e, 0x2a, 0xe2, 0xe9, 0x9b, 0x99, 0x6f, 0x66, 0xe0, 0xdb, 0xc5, 0x67, 0xd3, 0x27, 0x66,
	0x84, 0x8d, 0x56, 0xa3, 0x44, 0x00, 0x8b, 0xb8, 0xce, 0xc0, 0x30, 0x0e, 0x61, 0x6e, 0x34, 0x68,
	0x72, 0x52, 0x6f, 0xc3, 0x66, 0xd4, 0xa6, 0xd1, 0xe0, 0x58, 0x6a, 0xa9, 0x2b, 0x47, 0x54, 0xa2,
	0xda, 0x7c, 0xf1, 0xea, 0xe0, 0x83, 0x07, 0x2b, 0x6f, 0x9a, 0x0a, 0xd2, 0xc7, 0x8e, 0x4a, 0x29,
	0xf2, 0x51, 0xb0, 0x1f, 0x3b, 0x2a, 0x25, 0xf7, 0xb8, 0xcb, 0x8d, 0x60, 0xa0, 0x0d, 0x75, 0x7c,
	0x14, 0x1c, 0x4e, 0x46, 0xdf, 0x9f, 0xe7, 0x43, 0xa9, 0x60, 0xb6, 0x4c, 0x42, 0xae, 0x17, 0x11,
	0xd7, 0x76, 0xa1, 0x6d, 0x33, 0x86, 0x36, 0x9d, 0x47, 0xb0, 0xce, 0x85, 0x0d, 0xc7, 0x9c, 0x8f,
	0xd3, 0xd4, 0x08, 0x6b, 0xe3, 0x5d, 0x03, 0x39, 0xc2, 0xed, 0xb9, 0x58, 0xd3, 0x76, 0xd5, 0x5e,
	0x42, 0x32, 0xc0, 0x3d, 0x23, 0xb8, 0x50, 0x2b, 0x61, 0xe8, 0x5e, 0x25, 0xff, 0x72, 0x72, 0x8a,
	0x5d, 0xc6, 0xb9, 0xc8, 0x81, 0x76, 0xaa, 0x4d, 0xc3, 0xc8, 0x25, 0xee, 0xa7, 0x4b, 0xc3, 0x40,
	0xe9, 0xec, 0x4e, 0x28, 0x39, 0x03, 0xea, 0xfa, 0x28, 0xe8, 0xc4, 0xff, 0xd4, 0x32, 0x9f, 0x68,
	0x00, 0xbd, 0xa0, 0xdd, 0x3a, 0x5f, 0xb3, 0x52, 0x9f, 0xd5, 0xb9, 0x9e, 0x8f, 0x02, 0x37, 0x6e,
	0xd8, 0xe4, 0xf6, 0x7d, 0xeb, 0xa1, 0xcd, 0xd6, 0x43, 0x5f, 0x5b, 0x0f, 0xbd, 0x15, 0x5e, 0x6b,
	0x53, 0x78, 0xad, 0x8f, 0xc2, 0x6b, 0x3d, 0x5e, 0xfd, 0xb9, 0xf7, 0x99, 0x65, 0x72, 0x2a, 0x54,
	0xae, 0x32, 0x19, 0x35, 0xdf, 0xf0, 0xb2, 0x03, 0xd5, 0xe5, 0x89, 0x5b, 0xbd, 0xec, 0xf5, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x3a, 0xdf, 0x9d, 0xa6, 0x01, 0x00, 0x00,
}

func (m *MsgContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Height))
		i--
		dAtA[i] = 0x41
	}
	if len(m.Bottom) > 0 {
		i -= len(m.Bottom)
		copy(dAtA[i:], m.Bottom)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Bottom)))
		i--
		dAtA[i] = 0x3a
	}
	if m.DurationHeight != 0 {
		i = encodeVarintContract(dAtA, i, uint64(m.DurationHeight))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Accept) > 0 {
		i -= len(m.Accept)
		copy(dAtA[i:], m.Accept)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Accept)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	l = len(m.Accept)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	if m.DurationHeight != 0 {
		n += 1 + sovContract(uint64(m.DurationHeight))
	}
	l = len(m.Bottom)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	if m.Height != 0 {
		n += 9
	}
	return n
}

func sovContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContract(x uint64) (n int) {
	return sovContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
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
			return fmt.Errorf("proto: MsgContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
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
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
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
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accept", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accept = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationHeight", wireType)
			}
			m.DurationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DurationHeight |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bottom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bottom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Height = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContract
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthContract
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
func skipContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
				return 0, ErrInvalidLengthContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContract = fmt.Errorf("proto: unexpected end of group")
)
