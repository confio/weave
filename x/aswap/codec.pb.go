// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/aswap/codec.proto

package aswap

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_iov_one_weave "github.com/iov-one/weave"
	weave "github.com/iov-one/weave"
	coin "github.com/iov-one/weave/coin"
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

// Swap is designed to hold some coins for atomic swap, locked by preimage_hash
type Swap struct {
	// metadata is used for schema versioning support
	Metadata *weave.Metadata                  `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Sender   github_com_iov_one_weave.Address `protobuf:"bytes,2,opt,name=sender,proto3,casttype=github.com/iov-one/weave.Address" json:"sender,omitempty"`
	// preimage hash is a hash of user's secret preimage that would serve as a "password"
	// to unlock the coins in the swap should the transaction be a success
	PreimageHash github_com_iov_one_weave.Condition `protobuf:"bytes,3,opt,name=preimage_hash,json=preimageHash,proto3,casttype=github.com/iov-one/weave.Condition" json:"preimage_hash,omitempty"`
	Recipient    github_com_iov_one_weave.Address   `protobuf:"bytes,4,opt,name=recipient,proto3,casttype=github.com/iov-one/weave.Address" json:"recipient,omitempty"`
	// If unreleased before timeout, swap will return coins to sender.
	// Timeout represents wall clock time as read from the block header. Timeout
	// is represented using POSIX time format.
	// Expiration time is inclusive meaning that the swap expires as soon as
	// the current time is equal or greater than timeout value.
	// nonexpired: [created, timeout)
	// expired: [timeout, infinity)
	Timeout github_com_iov_one_weave.UnixTime `protobuf:"varint,5,opt,name=timeout,proto3,casttype=github.com/iov-one/weave.UnixTime" json:"timeout,omitempty"`
	// max length 128 characters
	Memo string `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *Swap) Reset()         { *m = Swap{} }
func (m *Swap) String() string { return proto.CompactTextString(m) }
func (*Swap) ProtoMessage()    {}
func (*Swap) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad79b700d8686a3f, []int{0}
}
func (m *Swap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Swap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Swap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Swap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Swap.Merge(m, src)
}
func (m *Swap) XXX_Size() int {
	return m.Size()
}
func (m *Swap) XXX_DiscardUnknown() {
	xxx_messageInfo_Swap.DiscardUnknown(m)
}

var xxx_messageInfo_Swap proto.InternalMessageInfo

func (m *Swap) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Swap) GetSender() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Swap) GetPreimageHash() github_com_iov_one_weave.Condition {
	if m != nil {
		return m.PreimageHash
	}
	return nil
}

func (m *Swap) GetRecipient() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *Swap) GetTimeout() github_com_iov_one_weave.UnixTime {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Swap) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

// CreateSwapMsg creates a Swap with some coins.
type CreateSwapMsg struct {
	Metadata *weave.Metadata                  `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Sender   github_com_iov_one_weave.Address `protobuf:"bytes,2,opt,name=sender,proto3,casttype=github.com/iov-one/weave.Address" json:"sender,omitempty"`
	// sha256 hash of preimage
	PreimageHash []byte                           `protobuf:"bytes,3,opt,name=preimage_hash,json=preimageHash,proto3" json:"preimage_hash,omitempty"`
	Recipient    github_com_iov_one_weave.Address `protobuf:"bytes,4,opt,name=recipient,proto3,casttype=github.com/iov-one/weave.Address" json:"recipient,omitempty"`
	// amount may contain multiple token types
	Amount []*coin.Coin `protobuf:"bytes,5,rep,name=amount,proto3" json:"amount,omitempty"`
	// Timeout represents wall clock time.
	Timeout github_com_iov_one_weave.UnixTime `protobuf:"varint,6,opt,name=timeout,proto3,casttype=github.com/iov-one/weave.UnixTime" json:"timeout,omitempty"`
	// max length 128 character
	Memo string `protobuf:"bytes,7,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *CreateSwapMsg) Reset()         { *m = CreateSwapMsg{} }
func (m *CreateSwapMsg) String() string { return proto.CompactTextString(m) }
func (*CreateSwapMsg) ProtoMessage()    {}
func (*CreateSwapMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad79b700d8686a3f, []int{1}
}
func (m *CreateSwapMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateSwapMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateSwapMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateSwapMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSwapMsg.Merge(m, src)
}
func (m *CreateSwapMsg) XXX_Size() int {
	return m.Size()
}
func (m *CreateSwapMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSwapMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSwapMsg proto.InternalMessageInfo

func (m *CreateSwapMsg) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CreateSwapMsg) GetSender() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *CreateSwapMsg) GetPreimageHash() []byte {
	if m != nil {
		return m.PreimageHash
	}
	return nil
}

func (m *CreateSwapMsg) GetRecipient() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *CreateSwapMsg) GetAmount() []*coin.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *CreateSwapMsg) GetTimeout() github_com_iov_one_weave.UnixTime {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *CreateSwapMsg) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

// ReleaseSwapMsg releases the tokens to the recepient.
// This operation is authorized by preimage, which is sent raw and then hashed on the backend.
type ReleaseSwapMsg struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	EscrowId []byte          `protobuf:"bytes,2,opt,name=escrow_id,json=escrowId,proto3" json:"escrow_id,omitempty"`
	// raw preimage to unlock swap
	Preimage []byte `protobuf:"bytes,3,opt,name=preimage,proto3" json:"preimage,omitempty"`
}

func (m *ReleaseSwapMsg) Reset()         { *m = ReleaseSwapMsg{} }
func (m *ReleaseSwapMsg) String() string { return proto.CompactTextString(m) }
func (*ReleaseSwapMsg) ProtoMessage()    {}
func (*ReleaseSwapMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad79b700d8686a3f, []int{2}
}
func (m *ReleaseSwapMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReleaseSwapMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReleaseSwapMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReleaseSwapMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseSwapMsg.Merge(m, src)
}
func (m *ReleaseSwapMsg) XXX_Size() int {
	return m.Size()
}
func (m *ReleaseSwapMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseSwapMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseSwapMsg proto.InternalMessageInfo

func (m *ReleaseSwapMsg) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ReleaseSwapMsg) GetEscrowId() []byte {
	if m != nil {
		return m.EscrowId
	}
	return nil
}

func (m *ReleaseSwapMsg) GetPreimage() []byte {
	if m != nil {
		return m.Preimage
	}
	return nil
}

// ReturnSwapMsg releases the tokens to the sender.
// This operation only works if the Swap is expired.
type ReturnSwapMsg struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	EscrowId []byte          `protobuf:"bytes,2,opt,name=escrow_id,json=escrowId,proto3" json:"escrow_id,omitempty"`
}

func (m *ReturnSwapMsg) Reset()         { *m = ReturnSwapMsg{} }
func (m *ReturnSwapMsg) String() string { return proto.CompactTextString(m) }
func (*ReturnSwapMsg) ProtoMessage()    {}
func (*ReturnSwapMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad79b700d8686a3f, []int{3}
}
func (m *ReturnSwapMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReturnSwapMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReturnSwapMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReturnSwapMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnSwapMsg.Merge(m, src)
}
func (m *ReturnSwapMsg) XXX_Size() int {
	return m.Size()
}
func (m *ReturnSwapMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnSwapMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnSwapMsg proto.InternalMessageInfo

func (m *ReturnSwapMsg) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ReturnSwapMsg) GetEscrowId() []byte {
	if m != nil {
		return m.EscrowId
	}
	return nil
}

func init() {
	proto.RegisterType((*Swap)(nil), "aswap.Swap")
	proto.RegisterType((*CreateSwapMsg)(nil), "aswap.CreateSwapMsg")
	proto.RegisterType((*ReleaseSwapMsg)(nil), "aswap.ReleaseSwapMsg")
	proto.RegisterType((*ReturnSwapMsg)(nil), "aswap.ReturnSwapMsg")
}

func init() { proto.RegisterFile("x/aswap/codec.proto", fileDescriptor_ad79b700d8686a3f) }

var fileDescriptor_ad79b700d8686a3f = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9b, 0xb6, 0x9b, 0x6d, 0x67, 0xb7, 0x0a, 0xe3, 0x25, 0x54, 0x48, 0x63, 0x5c, 0x25,
	0x22, 0x9b, 0xc0, 0x7a, 0x15, 0xc4, 0xf6, 0xa2, 0xc8, 0x5e, 0x46, 0x3d, 0x78, 0x5a, 0xa6, 0x99,
	0x67, 0x3a, 0xe0, 0xcc, 0x84, 0x99, 0x49, 0xbb, 0x7f, 0x86, 0x7f, 0x96, 0xc7, 0xbd, 0xe9, 0x41,
	0x8a, 0xb4, 0xff, 0x45, 0x4f, 0xd2, 0x69, 0xbb, 0x46, 0xb0, 0x82, 0x3f, 0x6f, 0xef, 0x7d, 0xf3,
	0xfd, 0xbe, 0xf7, 0xf8, 0x90, 0x41, 0xb7, 0x2e, 0x33, 0x6a, 0x66, 0xb4, 0xcc, 0x72, 0xc5, 0x20,
	0x4f, 0x4b, 0xad, 0xac, 0xc2, 0x07, 0x4e, 0xea, 0x9f, 0x14, 0xdc, 0x4e, 0xaa, 0x71, 0x9a, 0x2b,
	0x91, 0x71, 0x35, 0x3d, 0x55, 0x12, 0xb2, 0x19, 0xd0, 0x29, 0xd4, 0xcd, 0xfd, 0x07, 0x3f, 0x71,
	0x71, 0xf9, 0x9d, 0xf5, 0xb4, 0x66, 0x2d, 0x54, 0xa1, 0x32, 0x27, 0x8f, 0xab, 0xb7, 0xae, 0x73,
	0x8d, 0xab, 0x36, 0xf6, 0xf8, 0x63, 0x13, 0xb5, 0x5f, 0xce, 0x68, 0x89, 0x1f, 0xa2, 0x8e, 0x00,
	0x4b, 0x19, 0xb5, 0x34, 0xf0, 0x22, 0x2f, 0x39, 0x3a, 0xbb, 0x99, 0xba, 0x15, 0xe9, 0xf9, 0x56,
	0x26, 0xd7, 0x06, 0xfc, 0x18, 0xf9, 0x06, 0x24, 0x03, 0x1d, 0x34, 0x23, 0x2f, 0x39, 0x1e, 0x9e,
	0xac, 0xe6, 0x83, 0x68, 0xdf, 0x8d, 0xe9, 0x53, 0xc6, 0x34, 0x18, 0x43, 0xb6, 0x19, 0xfc, 0x02,
	0xf5, 0x4a, 0x0d, 0x5c, 0xd0, 0x02, 0x2e, 0x26, 0xd4, 0x4c, 0x82, 0x96, 0x1b, 0x72, 0x7f, 0x35,
	0x1f, 0xc4, 0x7b, 0x87, 0x8c, 0x94, 0x64, 0xdc, 0x72, 0x25, 0xc9, 0xf1, 0x2e, 0xfc, 0x8c, 0x9a,
	0x09, 0x1e, 0xa2, 0xae, 0x86, 0x9c, 0x97, 0x1c, 0xa4, 0x0d, 0xda, 0xbf, 0x70, 0xcd, 0xb7, 0x18,
	0x7e, 0x82, 0x0e, 0x2d, 0x17, 0xa0, 0x2a, 0x1b, 0x1c, 0x44, 0x5e, 0xd2, 0x1a, 0xde, 0x5b, 0xcd,
	0x07, 0x77, 0xf6, 0x4e, 0x78, 0x2d, 0xf9, 0xe5, 0x2b, 0x2e, 0x80, 0xec, 0x52, 0x18, 0xa3, 0xb6,
	0x00, 0xa1, 0x02, 0x3f, 0xf2, 0x92, 0x2e, 0x71, 0x75, 0xfc, 0xb9, 0x89, 0x7a, 0x23, 0x0d, 0xd4,
	0xc2, 0x9a, 0xef, 0xb9, 0x29, 0xfe, 0x27, 0xe2, 0xbb, 0x3f, 0x44, 0xfc, 0x0f, 0xd0, 0xc5, 0xc8,
	0xa7, 0x42, 0x55, 0x72, 0x4d, 0xae, 0x95, 0x1c, 0x9d, 0xa1, 0x74, 0xfd, 0x47, 0xa6, 0x23, 0xc5,
	0x25, 0xd9, 0x7e, 0xa9, 0xe3, 0xf5, 0xff, 0x08, 0xef, 0x61, 0x0d, 0xef, 0x14, 0xdd, 0x20, 0xf0,
	0x0e, 0xa8, 0xf9, 0x3d, 0xbc, 0xb7, 0x51, 0x17, 0x4c, 0xae, 0xd5, 0xec, 0x82, 0xb3, 0x0d, 0x61,
	0xd2, 0xd9, 0x08, 0xcf, 0x19, 0xee, 0xa3, 0xce, 0x0e, 0xd4, 0x16, 0xdc, 0x75, 0x1f, 0xbf, 0x41,
	0x3d, 0x02, 0xb6, 0xd2, 0xf2, 0xaf, 0xaf, 0x1d, 0x06, 0x1f, 0x16, 0xa1, 0x77, 0xb5, 0x08, 0xbd,
	0x2f, 0x8b, 0xd0, 0x7b, 0xbf, 0x0c, 0x1b, 0x57, 0xcb, 0xb0, 0xf1, 0x69, 0x19, 0x36, 0xc6, 0xbe,
	0x7b, 0xac, 0x8f, 0xbe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xc5, 0x7a, 0x23, 0x4a, 0x04, 0x00,
	0x00,
}

func (m *Swap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Swap) MarshalTo(dAtA []byte) (int, error) {
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
	if len(m.Sender) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Sender)))
		i += copy(dAtA[i:], m.Sender)
	}
	if len(m.PreimageHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.PreimageHash)))
		i += copy(dAtA[i:], m.PreimageHash)
	}
	if len(m.Recipient) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Recipient)))
		i += copy(dAtA[i:], m.Recipient)
	}
	if m.Timeout != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Timeout))
	}
	if len(m.Memo) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Memo)))
		i += copy(dAtA[i:], m.Memo)
	}
	return i, nil
}

func (m *CreateSwapMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateSwapMsg) MarshalTo(dAtA []byte) (int, error) {
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
	if len(m.Sender) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Sender)))
		i += copy(dAtA[i:], m.Sender)
	}
	if len(m.PreimageHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.PreimageHash)))
		i += copy(dAtA[i:], m.PreimageHash)
	}
	if len(m.Recipient) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Recipient)))
		i += copy(dAtA[i:], m.Recipient)
	}
	if len(m.Amount) > 0 {
		for _, msg := range m.Amount {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Timeout != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Timeout))
	}
	if len(m.Memo) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Memo)))
		i += copy(dAtA[i:], m.Memo)
	}
	return i, nil
}

func (m *ReleaseSwapMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReleaseSwapMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n3, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.EscrowId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.EscrowId)))
		i += copy(dAtA[i:], m.EscrowId)
	}
	if len(m.Preimage) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Preimage)))
		i += copy(dAtA[i:], m.Preimage)
	}
	return i, nil
}

func (m *ReturnSwapMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReturnSwapMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n4, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.EscrowId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.EscrowId)))
		i += copy(dAtA[i:], m.EscrowId)
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
func (m *Swap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.PreimageHash)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovCodec(uint64(m.Timeout))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *CreateSwapMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.PreimageHash)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	if m.Timeout != 0 {
		n += 1 + sovCodec(uint64(m.Timeout))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *ReleaseSwapMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.EscrowId)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Preimage)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *ReturnSwapMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.EscrowId)
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
func (m *Swap) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Swap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Swap: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreimageHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreimageHash = append(m.PreimageHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PreimageHash == nil {
				m.PreimageHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = append(m.Recipient[:0], dAtA[iNdEx:postIndex]...)
			if m.Recipient == nil {
				m.Recipient = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= github_com_iov_one_weave.UnixTime(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
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
			m.Memo = string(dAtA[iNdEx:postIndex])
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
func (m *CreateSwapMsg) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CreateSwapMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateSwapMsg: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreimageHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreimageHash = append(m.PreimageHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PreimageHash == nil {
				m.PreimageHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = append(m.Recipient[:0], dAtA[iNdEx:postIndex]...)
			if m.Recipient == nil {
				m.Recipient = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = append(m.Amount, &coin.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= github_com_iov_one_weave.UnixTime(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
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
			m.Memo = string(dAtA[iNdEx:postIndex])
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
func (m *ReleaseSwapMsg) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ReleaseSwapMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReleaseSwapMsg: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EscrowId = append(m.EscrowId[:0], dAtA[iNdEx:postIndex]...)
			if m.EscrowId == nil {
				m.EscrowId = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Preimage", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Preimage = append(m.Preimage[:0], dAtA[iNdEx:postIndex]...)
			if m.Preimage == nil {
				m.Preimage = []byte{}
			}
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
func (m *ReturnSwapMsg) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ReturnSwapMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReturnSwapMsg: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EscrowId = append(m.EscrowId[:0], dAtA[iNdEx:postIndex]...)
			if m.EscrowId == nil {
				m.EscrowId = []byte{}
			}
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
