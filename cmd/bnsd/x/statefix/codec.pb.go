// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cmd/bnsd/x/statefix/codec.proto

package statefix

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	weave "github.com/iov-one/weave"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ExecutedFix struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	FixID    string          `protobuf:"bytes,2,opt,name=fix_id,json=fixId,proto3" json:"fix_id,omitempty"`
}

func (m *ExecutedFix) Reset()         { *m = ExecutedFix{} }
func (m *ExecutedFix) String() string { return proto.CompactTextString(m) }
func (*ExecutedFix) ProtoMessage()    {}
func (*ExecutedFix) Descriptor() ([]byte, []int) {
	return fileDescriptor_28f7a42f85283a60, []int{0}
}
func (m *ExecutedFix) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecutedFix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExecutedFix.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExecutedFix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutedFix.Merge(m, src)
}
func (m *ExecutedFix) XXX_Size() int {
	return m.Size()
}
func (m *ExecutedFix) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutedFix.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutedFix proto.InternalMessageInfo

func (m *ExecutedFix) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ExecutedFix) GetFixID() string {
	if m != nil {
		return m.FixID
	}
	return ""
}

type ExecuteFixMsg struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	FixID    string          `protobuf:"bytes,2,opt,name=fix_id,json=fixId,proto3" json:"fix_id,omitempty"`
}

func (m *ExecuteFixMsg) Reset()         { *m = ExecuteFixMsg{} }
func (m *ExecuteFixMsg) String() string { return proto.CompactTextString(m) }
func (*ExecuteFixMsg) ProtoMessage()    {}
func (*ExecuteFixMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_28f7a42f85283a60, []int{1}
}
func (m *ExecuteFixMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecuteFixMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExecuteFixMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExecuteFixMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteFixMsg.Merge(m, src)
}
func (m *ExecuteFixMsg) XXX_Size() int {
	return m.Size()
}
func (m *ExecuteFixMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteFixMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteFixMsg proto.InternalMessageInfo

func (m *ExecuteFixMsg) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ExecuteFixMsg) GetFixID() string {
	if m != nil {
		return m.FixID
	}
	return ""
}

func init() {
	proto.RegisterType((*ExecutedFix)(nil), "statefix.ExecutedFix")
	proto.RegisterType((*ExecuteFixMsg)(nil), "statefix.ExecuteFixMsg")
}

func init() { proto.RegisterFile("cmd/bnsd/x/statefix/codec.proto", fileDescriptor_28f7a42f85283a60) }

var fileDescriptor_28f7a42f85283a60 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xce, 0x4d, 0xd1,
	0x4f, 0xca, 0x2b, 0x4e, 0xd1, 0xaf, 0xd0, 0x2f, 0x2e, 0x49, 0x2c, 0x49, 0x4d, 0xcb, 0xac, 0xd0,
	0x4f, 0xce, 0x4f, 0x49, 0x4d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x89, 0x4a,
	0x71, 0x23, 0x09, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99, 0xfa, 0x20, 0x16, 0x44, 0x54,
	0x29, 0x86, 0x8b, 0xdb, 0xb5, 0x22, 0x35, 0xb9, 0xb4, 0x24, 0x35, 0xc5, 0x2d, 0xb3, 0x42, 0x48,
	0x9b, 0x8b, 0x23, 0x37, 0xb5, 0x24, 0x31, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xdb, 0x88, 0x5f, 0xaf, 0x3c, 0x35, 0xb1, 0x2c, 0x55, 0xcf, 0x17, 0x2a, 0x1c, 0x04, 0x57, 0x20,
	0xa4, 0xc0, 0xc5, 0x96, 0x96, 0x59, 0x11, 0x9f, 0x99, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0xe9,
	0xc4, 0xf9, 0xe8, 0x9e, 0x3c, 0xab, 0x5b, 0x66, 0x85, 0xa7, 0x4b, 0x10, 0x6b, 0x5a, 0x66, 0x85,
	0x67, 0x8a, 0x52, 0x1c, 0x17, 0x2f, 0xd4, 0x74, 0xb7, 0xcc, 0x0a, 0xdf, 0xe2, 0x74, 0x2a, 0x9b,
	0xef, 0x24, 0x71, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e,
	0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0xef,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x73, 0xdf, 0xc8, 0xdf, 0x2e, 0x01, 0x00, 0x00,
}

func (m *ExecutedFix) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecutedFix) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.FixID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.FixID)))
		i += copy(dAtA[i:], m.FixID)
	}
	return i, nil
}

func (m *ExecuteFixMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecuteFixMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n2, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.FixID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.FixID)))
		i += copy(dAtA[i:], m.FixID)
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
func (m *ExecutedFix) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.FixID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *ExecuteFixMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.FixID)
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
func (m *ExecutedFix) Unmarshal(dAtA []byte) error {
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
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExecutedFix: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecutedFix: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &weave.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixID", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FixID = string(dAtA[iNdEx:postIndex])
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
			if (iNdEx + skippy) < 0 {
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
func (m *ExecuteFixMsg) Unmarshal(dAtA []byte) error {
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
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExecuteFixMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecuteFixMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &weave.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixID", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FixID = string(dAtA[iNdEx:postIndex])
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
			if (iNdEx + skippy) < 0 {
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
			if length < 0 {
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
			if iNdEx < 0 {
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
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCodec
				}
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
