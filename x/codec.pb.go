// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/codec.proto

package x

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Coin can hold any amount between -1 billion and +1 billion
// at steps of 10^-9. It is a fixed-point decimal
// representation and uses integers to avoid rounding
// associated with floats.
//
// Every code has a denomination, which is just a
//
// If you want anything more complex, you should write your
// own type, possibly borrowing from this code.
type Coin struct {
	// Whole coins, -10^15 < integer < 10^15
	Whole int64 `protobuf:"varint,1,opt,name=whole,proto3" json:"whole,omitempty"`
	// Billionth of coins. 0 <= abs(fractional) < 10^9
	// If fractional != 0, must have same sign as integer
	Fractional int64 `protobuf:"varint,2,opt,name=fractional,proto3" json:"fractional,omitempty"`
	// Ticker is 3-4 upper-case letters and
	// all Coins of the same currency can be combined
	Ticker               string   `protobuf:"bytes,3,opt,name=ticker,proto3" json:"ticker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coin) Reset()         { *m = Coin{} }
func (m *Coin) String() string { return proto.CompactTextString(m) }
func (*Coin) ProtoMessage()    {}
func (*Coin) Descriptor() ([]byte, []int) {
	return fileDescriptor_codec_2c204a1b663c1bb6, []int{0}
}
func (m *Coin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Coin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Coin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Coin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coin.Merge(dst, src)
}
func (m *Coin) XXX_Size() int {
	return m.Size()
}
func (m *Coin) XXX_DiscardUnknown() {
	xxx_messageInfo_Coin.DiscardUnknown(m)
}

var xxx_messageInfo_Coin proto.InternalMessageInfo

func (m *Coin) GetWhole() int64 {
	if m != nil {
		return m.Whole
	}
	return 0
}

func (m *Coin) GetFractional() int64 {
	if m != nil {
		return m.Fractional
	}
	return 0
}

func (m *Coin) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func init() {
	proto.RegisterType((*Coin)(nil), "x.Coin")
}
func (m *Coin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Coin) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Whole != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Whole))
	}
	if m.Fractional != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Fractional))
	}
	if len(m.Ticker) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Ticker)))
		i += copy(dAtA[i:], m.Ticker)
	}
	return i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Coin) Size() (n int) {
	var l int
	_ = l
	if m.Whole != 0 {
		n += 1 + sovCodec(uint64(m.Whole))
	}
	if m.Fractional != 0 {
		n += 1 + sovCodec(uint64(m.Fractional))
	}
	l = len(m.Ticker)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Coin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Coin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Coin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Whole", wireType)
			}
			m.Whole = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Whole |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fractional", wireType)
			}
			m.Fractional = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fractional |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ticker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCodec
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCodec(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("x/codec.proto", fileDescriptor_codec_2c204a1b663c1bb6) }

var fileDescriptor_codec_2c204a1b663c1bb6 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xad, 0xd0, 0x4f, 0xce,
	0x4f, 0x49, 0x4d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xac, 0x50, 0x0a, 0xe1, 0x62,
	0x71, 0xce, 0xcf, 0xcc, 0x13, 0x12, 0xe1, 0x62, 0x2d, 0xcf, 0xc8, 0xcf, 0x49, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0x60, 0x0e, 0x82, 0x70, 0x84, 0xe4, 0xb8, 0xb8, 0xd2, 0x8a, 0x12, 0x93, 0x4b, 0x32,
	0xf3, 0xf3, 0x12, 0x73, 0x24, 0x98, 0xc0, 0x52, 0x48, 0x22, 0x42, 0x62, 0x5c, 0x6c, 0x25, 0x99,
	0xc9, 0xd9, 0xa9, 0x45, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x93, 0xc0, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0x43,
	0x12, 0x1b, 0xd8, 0x46, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xc1, 0xfa, 0x9f, 0x82,
	0x00, 0x00, 0x00,
}
