// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flares/v1beta/contract.proto

package types

import (
	encoding_binary "encoding/binary"
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

type MsgContract struct {
	Id             string                                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator        github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Module         string                                        `protobuf:"bytes,3,opt,name=module,proto3" json:"module,omitempty"`
	Key            string                                        `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Receiver       string                                        `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Accept         string                                        `protobuf:"bytes,6,opt,name=accept,proto3" json:"accept,omitempty"`
	DurationHeight int32                                         `protobuf:"varint,7,opt,name=durationHeight,proto3" json:"durationHeight,omitempty"`
	Bottom         string                                        `protobuf:"bytes,8,opt,name=bottom,proto3" json:"bottom,omitempty"`
	Height         uint64                                        `protobuf:"fixed64,9,opt,name=height,proto3" json:"height,omitempty"`
	Code           uint32                                        `protobuf:"varint,10,opt,name=code,proto3" json:"code,omitempty"`
	Result         string                                        `protobuf:"bytes,11,opt,name=result,proto3" json:"result,omitempty"`
}

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

func (m *MsgContract) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
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

func (m *MsgContract) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *MsgContract) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgContract)(nil), "flares.flares.v1beta1.MsgContract")
}

func init() { proto.RegisterFile("flares/v1beta/contract.proto", fileDescriptor_bdae3fe79527cbf4) }

var fileDescriptor_bdae3fe79527cbf4 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0x41, 0x6b, 0xdb, 0x30,
	0x18, 0x8d, 0x9c, 0xc4, 0x49, 0x94, 0x2d, 0x0c, 0xb1, 0x0d, 0x11, 0x86, 0x67, 0x76, 0x18, 0x3e,
	0x2c, 0x36, 0x61, 0xbf, 0x20, 0x19, 0x8c, 0x42, 0xe9, 0xc5, 0xc7, 0xde, 0x6c, 0xe9, 0x8b, 0x23,
	0x12, 0x5b, 0x46, 0x92, 0xd3, 0xe6, 0x5f, 0xf4, 0x67, 0xf5, 0xd0, 0x43, 0x8e, 0x3d, 0x95, 0x92,
	0xfc, 0x8b, 0x9e, 0x8a, 0x2d, 0xa7, 0x94, 0x9e, 0xbe, 0xf7, 0x9e, 0xde, 0x7b, 0x82, 0xef, 0xc3,
	0x3f, 0x56, 0xdb, 0x44, 0x81, 0x8e, 0x76, 0xf3, 0x14, 0x4c, 0x12, 0x31, 0x59, 0x18, 0x95, 0x30,
	0x13, 0x96, 0x4a, 0x1a, 0x49, 0xbe, 0xd9, 0xd7, 0xb0, 0x1d, 0xd6, 0x34, 0x9f, 0x7e, 0xcd, 0x64,
	0x26, 0x1b, 0x47, 0x54, 0x23, 0x6b, 0xfe, 0xf5, 0xe0, 0xe0, 0xf1, 0x95, 0xce, 0xfe, 0xb5, 0x15,
	0x64, 0x82, 0x1d, 0xc1, 0x29, 0xf2, 0x51, 0x30, 0x8a, 0x1d, 0xc1, 0xc9, 0x25, 0x1e, 0x30, 0x05,
	0x89, 0x91, 0x8a, 0x3a, 0x3e, 0x0a, 0x3e, 0x2d, 0xe7, 0x2f, 0x4f, 0x3f, 0x67, 0x99, 0x30, 0xeb,
	0x2a, 0x0d, 0x99, 0xcc, 0x23, 0x26, 0x75, 0x2e, 0x75, 0x3b, 0x66, 0x9a, 0x6f, 0x22, 0xb3, 0x2f,
	0x41, 0x87, 0x0b, 0xc6, 0x16, 0x9c, 0x2b, 0xd0, 0x3a, 0x3e, 0x37, 0x90, 0xef, 0xd8, 0xcd, 0x25,
	0xaf, 0xb6, 0x40, 0xbb, 0xcd, 0x07, 0x2d, 0x23, 0x5f, 0x70, 0x77, 0x03, 0x7b, 0xda, 0x6b, 0xc4,
	0x1a, 0x92, 0x29, 0x1e, 0x2a, 0x60, 0x20, 0x76, 0xa0, 0x68, 0xbf, 0x91, 0xdf, 0x78, 0xdd, 0x92,
	0x30, 0x06, 0xa5, 0xa1, 0xae, 0x6d, 0xb1, 0x8c, 0xfc, 0xc6, 0x13, 0x5e, 0xa9, 0xc4, 0x08, 0x59,
	0x5c, 0x80, 0xc8, 0xd6, 0x86, 0x0e, 0x7c, 0x14, 0xf4, 0xe3, 0x0f, 0x6a, 0x9d, 0x4f, 0xa5, 0x31,
	0x32, 0xa7, 0x43, 0x9b, 0xb7, 0xac, 0xd6, 0xd7, 0x36, 0x37, 0xf2, 0x51, 0xe0, 0xc6, 0x2d, 0x23,
	0x04, 0xf7, 0x98, 0xe4, 0x40, 0xb1, 0x8f, 0x82, 0xcf, 0x71, 0x83, 0x6b, 0xaf, 0x02, 0x5d, 0x6d,
	0x0d, 0x1d, 0xdb, 0x0e, 0xcb, 0x96, 0xff, 0xef, 0x8f, 0x1e, 0x3a, 0x1c, 0x3d, 0xf4, 0x7c, 0xf4,
	0xd0, 0xdd, 0xc9, 0xeb, 0x1c, 0x4e, 0x5e, 0xe7, 0xf1, 0xe4, 0x75, 0xae, 0xff, 0xbc, 0xdb, 0xd9,
	0x4d, 0x52, 0x64, 0x2b, 0x10, 0xa5, 0x28, 0xb2, 0xa8, 0x3d, 0xe5, 0xed, 0x19, 0x34, 0xdb, 0x4b,
	0xdd, 0xe6, 0x3a, 0x7f, 0x5f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xec, 0x4e, 0x9c, 0xea, 0x01,
	0x00, 0x00,
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
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x5a
	}
	if m.Code != 0 {
		i = encodeVarintContract(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x50
	}
	if m.Height != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Height))
		i--
		dAtA[i] = 0x49
	}
	if len(m.Bottom) > 0 {
		i -= len(m.Bottom)
		copy(dAtA[i:], m.Bottom)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Bottom)))
		i--
		dAtA[i] = 0x42
	}
	if m.DurationHeight != 0 {
		i = encodeVarintContract(dAtA, i, uint64(m.DurationHeight))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Accept) > 0 {
		i -= len(m.Accept)
		copy(dAtA[i:], m.Accept)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Accept)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintContract(dAtA, i, uint64(len(m.Module)))
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
	l = len(m.Module)
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
	if m.Code != 0 {
		n += 1 + sovContract(uint64(m.Code))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
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
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
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
		case 5:
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
		case 6:
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
		case 7:
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
		case 8:
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
		case 9:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Height = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
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
			m.Result = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
