// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validator/validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

type Validator struct {
	// the account address of validator owner
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// description of the validator
	Description *Description `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// the consensus public key of the tendermint validator
	PubKey *types.Any `protobuf:"bytes,3,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	// validator consensus power
	Power int32 `protobuf:"varint,4,opt,name=power,proto3" json:"power,omitempty"`
	// has the validator been removed from validator set
	Jailed bool `protobuf:"varint,5,opt,name=jailed,proto3" json:"jailed,omitempty"`
	// the reason of validator jailing
	JailedReason string `protobuf:"bytes,6,opt,name=jailedReason,proto3" json:"jailedReason,omitempty"`
}

func (m *Validator) Reset()      { *m = Validator{} }
func (*Validator) ProtoMessage() {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_e972accd7c1f0747, []int{0}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Validator)(nil), "zigbeealliance.distributedcomplianceledger.validator.Validator")
}

func init() { proto.RegisterFile("validator/validator.proto", fileDescriptor_e972accd7c1f0747) }

var fileDescriptor_e972accd7c1f0747 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x8e, 0xd4, 0x30,
	0x10, 0x86, 0xe3, 0x83, 0x44, 0x9c, 0xef, 0xaa, 0x28, 0x42, 0xde, 0x43, 0xca, 0x46, 0x57, 0xa5,
	0x89, 0x2d, 0x01, 0x15, 0xdd, 0xae, 0x10, 0x05, 0x34, 0x28, 0x48, 0x14, 0x34, 0x28, 0x89, 0x8d,
	0x31, 0xca, 0x66, 0x22, 0xdb, 0xe1, 0x08, 0x4f, 0x40, 0x49, 0x49, 0xb9, 0x0f, 0xc1, 0x43, 0x9c,
	0xa8, 0x56, 0x54, 0x94, 0x68, 0xb7, 0xe1, 0x31, 0xd0, 0xc6, 0xd9, 0xcd, 0xd2, 0xd2, 0xfd, 0x33,
	0xbf, 0xe7, 0xd7, 0xcc, 0x27, 0xe3, 0xd9, 0xc7, 0xa2, 0x56, 0xbc, 0xb0, 0xa0, 0xd9, 0x51, 0xd1,
	0x56, 0x83, 0x85, 0xf0, 0xf1, 0x67, 0x25, 0x4b, 0x21, 0x8a, 0xba, 0x56, 0x45, 0x53, 0x09, 0xca,
	0x95, 0xb1, 0x5a, 0x95, 0x9d, 0x15, 0xbc, 0x82, 0x55, 0xeb, 0xba, 0xb5, 0xe0, 0x52, 0x68, 0x7a,
	0x9c, 0xbd, 0x9a, 0x49, 0x00, 0x59, 0x0b, 0x36, 0x64, 0x94, 0xdd, 0x3b, 0x56, 0x34, 0xbd, 0x0b,
	0xbc, 0x8a, 0x24, 0x48, 0x18, 0x24, 0xdb, 0xab, 0xb1, 0x3b, 0xab, 0xc0, 0xac, 0xc0, 0xbc, 0x75,
	0x86, 0x2b, 0x46, 0xeb, 0xc1, 0xb4, 0x1c, 0x17, 0xa6, 0xd2, 0xaa, 0xb5, 0x0a, 0x1a, 0x67, 0x5e,
	0xdf, 0x9e, 0xe1, 0xf3, 0xd7, 0x07, 0x3f, 0xa4, 0xd8, 0x87, 0x9b, 0x46, 0x68, 0x82, 0x12, 0x94,
	0x9e, 0x2f, 0xc9, 0xcf, 0xef, 0x59, 0x34, 0x66, 0x2d, 0x38, 0xd7, 0xc2, 0x98, 0x57, 0x56, 0xab,
	0x46, 0xe6, 0xee, 0x59, 0x58, 0xe1, 0x8b, 0x93, 0x48, 0x72, 0x96, 0xa0, 0xf4, 0xe2, 0xe1, 0x82,
	0xfe, 0xcf, 0xc9, 0xf4, 0xe9, 0x14, 0x94, 0x9f, 0xa6, 0x86, 0xcf, 0x70, 0xd0, 0x76, 0xe5, 0x0b,
	0xd1, 0x93, 0x3b, 0x43, 0x7e, 0x44, 0x1d, 0x1c, 0x7a, 0x80, 0x43, 0x17, 0x4d, 0xbf, 0x24, 0x3f,
	0xa6, 0x5d, 0x2b, 0xdd, 0xb7, 0x16, 0xe8, 0xcb, 0x61, 0x2a, 0x1f, 0xa7, 0xc3, 0x08, 0xfb, 0x2d,
	0xdc, 0x08, 0x4d, 0xee, 0x26, 0x28, 0xf5, 0x73, 0x57, 0x84, 0xf7, 0x71, 0xf0, 0xa1, 0x50, 0xb5,
	0xe0, 0xc4, 0x4f, 0x50, 0x7a, 0x2f, 0x1f, 0xab, 0xf0, 0x1a, 0x5f, 0x3a, 0x95, 0x8b, 0xc2, 0x40,
	0x43, 0x82, 0x3d, 0x91, 0xfc, 0x9f, 0xde, 0x93, 0xcb, 0x2f, 0xeb, 0xb9, 0xf7, 0x6d, 0x3d, 0xf7,
	0xfe, 0xac, 0xe7, 0xde, 0x92, 0xdf, 0x6e, 0x63, 0xb4, 0xd9, 0xc6, 0xe8, 0xf7, 0x36, 0x46, 0x5f,
	0x77, 0xb1, 0xb7, 0xd9, 0xc5, 0xde, 0xaf, 0x5d, 0xec, 0xbd, 0x79, 0x2e, 0x95, 0x7d, 0xdf, 0x95,
	0xb4, 0x82, 0x15, 0x73, 0x6c, 0xb2, 0x03, 0x1c, 0x76, 0x02, 0x27, 0x9b, 0xe8, 0x64, 0x0e, 0x0f,
	0xfb, 0x34, 0xfd, 0x27, 0x66, 0xfb, 0x56, 0x98, 0x32, 0x18, 0xae, 0x7e, 0xf4, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x58, 0x0e, 0xb0, 0x72, 0x73, 0x02, 0x00, 0x00,
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.JailedReason) > 0 {
		i -= len(m.JailedReason)
		copy(dAtA[i:], m.JailedReason)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.JailedReason)))
		i--
		dAtA[i] = 0x32
	}
	if m.Jailed {
		i--
		if m.Jailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Power != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x20
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Description != nil {
		{
			size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.Description != nil {
		l = m.Description.Size()
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovValidator(uint64(m.Power))
	}
	if m.Jailed {
		n += 2
	}
	l = len(m.JailedReason)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Description == nil {
				m.Description = &Description{}
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &types.Any{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Jailed = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailedReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JailedReason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
