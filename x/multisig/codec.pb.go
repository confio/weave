// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/multisig/codec.proto

package multisig

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_iov_one_weave "github.com/iov-one/weave"
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

type Contract struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Participants defines a list of all signatures that are allowed to sign the
	// contract.
	Participants []*Participant `protobuf:"bytes,2,rep,name=participants,proto3" json:"participants,omitempty"`
	// Activation threshold defines the minimal weight value that must be
	// provided from participants in order to activate the contract. Weight is
	// computed as the sum of weights of all participating signatures.
	ActivationThreshold Weight `protobuf:"varint,3,opt,name=activation_threshold,json=activationThreshold,proto3,casttype=Weight" json:"activation_threshold,omitempty"`
	// Admin threshold defines the minimal weight value that must be provided
	// from participants in order to administrate the contract. Weight is
	// computed as the sum of weights of all participating signatures.
	AdminThreshold Weight `protobuf:"varint,4,opt,name=admin_threshold,json=adminThreshold,proto3,casttype=Weight" json:"admin_threshold,omitempty"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5080d98b87cf9a7, []int{0}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}
func (m *Contract) XXX_Size() int {
	return m.Size()
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Contract) GetParticipants() []*Participant {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *Contract) GetActivationThreshold() Weight {
	if m != nil {
		return m.ActivationThreshold
	}
	return 0
}

func (m *Contract) GetAdminThreshold() Weight {
	if m != nil {
		return m.AdminThreshold
	}
	return 0
}

// Participant clubs together a signature with a weight. The greater the weight
// the greater the power of a signature.
type Participant struct {
	Signature github_com_iov_one_weave.Address `protobuf:"bytes,1,opt,name=signature,proto3,casttype=github.com/iov-one/weave.Address" json:"signature,omitempty"`
	Weight    Weight                           `protobuf:"varint,2,opt,name=weight,proto3,casttype=Weight" json:"weight,omitempty"`
}

func (m *Participant) Reset()         { *m = Participant{} }
func (m *Participant) String() string { return proto.CompactTextString(m) }
func (*Participant) ProtoMessage()    {}
func (*Participant) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5080d98b87cf9a7, []int{1}
}
func (m *Participant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Participant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Participant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Participant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Participant.Merge(m, src)
}
func (m *Participant) XXX_Size() int {
	return m.Size()
}
func (m *Participant) XXX_DiscardUnknown() {
	xxx_messageInfo_Participant.DiscardUnknown(m)
}

var xxx_messageInfo_Participant proto.InternalMessageInfo

func (m *Participant) GetSignature() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Participant) GetWeight() Weight {
	if m != nil {
		return m.Weight
	}
	return 0
}

type CreateContractMsg struct {
	Metadata            *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Participants        []*Participant  `protobuf:"bytes,2,rep,name=participants,proto3" json:"participants,omitempty"`
	ActivationThreshold Weight          `protobuf:"varint,3,opt,name=activation_threshold,json=activationThreshold,proto3,casttype=Weight" json:"activation_threshold,omitempty"`
	AdminThreshold      Weight          `protobuf:"varint,4,opt,name=admin_threshold,json=adminThreshold,proto3,casttype=Weight" json:"admin_threshold,omitempty"`
}

func (m *CreateContractMsg) Reset()         { *m = CreateContractMsg{} }
func (m *CreateContractMsg) String() string { return proto.CompactTextString(m) }
func (*CreateContractMsg) ProtoMessage()    {}
func (*CreateContractMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5080d98b87cf9a7, []int{2}
}
func (m *CreateContractMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateContractMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateContractMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateContractMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateContractMsg.Merge(m, src)
}
func (m *CreateContractMsg) XXX_Size() int {
	return m.Size()
}
func (m *CreateContractMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateContractMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CreateContractMsg proto.InternalMessageInfo

func (m *CreateContractMsg) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CreateContractMsg) GetParticipants() []*Participant {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *CreateContractMsg) GetActivationThreshold() Weight {
	if m != nil {
		return m.ActivationThreshold
	}
	return 0
}

func (m *CreateContractMsg) GetAdminThreshold() Weight {
	if m != nil {
		return m.AdminThreshold
	}
	return 0
}

type UpdateContractMsg struct {
	Metadata            *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ContractID          []byte          `protobuf:"bytes,2,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	Participants        []*Participant  `protobuf:"bytes,3,rep,name=participants,proto3" json:"participants,omitempty"`
	ActivationThreshold Weight          `protobuf:"varint,4,opt,name=activation_threshold,json=activationThreshold,proto3,casttype=Weight" json:"activation_threshold,omitempty"`
	AdminThreshold      Weight          `protobuf:"varint,5,opt,name=admin_threshold,json=adminThreshold,proto3,casttype=Weight" json:"admin_threshold,omitempty"`
}

func (m *UpdateContractMsg) Reset()         { *m = UpdateContractMsg{} }
func (m *UpdateContractMsg) String() string { return proto.CompactTextString(m) }
func (*UpdateContractMsg) ProtoMessage()    {}
func (*UpdateContractMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5080d98b87cf9a7, []int{3}
}
func (m *UpdateContractMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateContractMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateContractMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateContractMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateContractMsg.Merge(m, src)
}
func (m *UpdateContractMsg) XXX_Size() int {
	return m.Size()
}
func (m *UpdateContractMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateContractMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateContractMsg proto.InternalMessageInfo

func (m *UpdateContractMsg) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *UpdateContractMsg) GetContractID() []byte {
	if m != nil {
		return m.ContractID
	}
	return nil
}

func (m *UpdateContractMsg) GetParticipants() []*Participant {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *UpdateContractMsg) GetActivationThreshold() Weight {
	if m != nil {
		return m.ActivationThreshold
	}
	return 0
}

func (m *UpdateContractMsg) GetAdminThreshold() Weight {
	if m != nil {
		return m.AdminThreshold
	}
	return 0
}

func init() {
	proto.RegisterType((*Contract)(nil), "multisig.Contract")
	proto.RegisterType((*Participant)(nil), "multisig.Participant")
	proto.RegisterType((*CreateContractMsg)(nil), "multisig.CreateContractMsg")
	proto.RegisterType((*UpdateContractMsg)(nil), "multisig.UpdateContractMsg")
}

func init() { proto.RegisterFile("x/multisig/codec.proto", fileDescriptor_e5080d98b87cf9a7) }

var fileDescriptor_e5080d98b87cf9a7 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x93, 0xcf, 0x6a, 0xea, 0x40,
	0x14, 0x87, 0x1d, 0xf5, 0x8a, 0x77, 0xe2, 0x55, 0xcc, 0xfd, 0x43, 0x70, 0x11, 0x43, 0x70, 0x21,
	0x5c, 0x4c, 0x40, 0x57, 0x5d, 0x74, 0xd1, 0xd8, 0x8d, 0x0b, 0xa1, 0x84, 0x96, 0x2e, 0x65, 0xcc,
	0x4c, 0x93, 0x01, 0x93, 0x09, 0xc9, 0x44, 0xfb, 0x18, 0x7d, 0x84, 0x3e, 0x4e, 0x97, 0x2e, 0xbb,
	0x92, 0x12, 0xa1, 0xdb, 0xee, 0x5d, 0x15, 0xa6, 0x46, 0x53, 0xa4, 0x14, 0xea, 0xae, 0xbb, 0xcc,
	0x99, 0xef, 0x97, 0x73, 0xce, 0x07, 0x03, 0xff, 0xdd, 0x9a, 0x7e, 0x32, 0xe3, 0x34, 0xa6, 0xae,
	0xe9, 0x30, 0x4c, 0x1c, 0x23, 0x8c, 0x18, 0x67, 0x72, 0x35, 0xab, 0xb6, 0x3a, 0x2e, 0xe5, 0x5e,
	0x32, 0x35, 0x1c, 0xe6, 0x9b, 0x94, 0xcd, 0x7b, 0x2c, 0x20, 0xe6, 0x82, 0xa0, 0x39, 0xc9, 0xf3,
	0xad, 0x5e, 0x8e, 0x72, 0x99, 0xcb, 0x4c, 0x51, 0x9e, 0x26, 0x37, 0xe2, 0x24, 0x0e, 0xe2, 0xeb,
	0x0d, 0xd7, 0x9f, 0x01, 0xac, 0x0e, 0x59, 0xc0, 0x23, 0xe4, 0x70, 0xf9, 0x3f, 0xac, 0xfa, 0x84,
	0x23, 0x8c, 0x38, 0x52, 0x80, 0x06, 0xba, 0x52, 0xbf, 0x61, 0x88, 0x0e, 0xc6, 0x78, 0x5b, 0xb6,
	0x77, 0x80, 0x7c, 0x02, 0x6b, 0x21, 0x8a, 0x38, 0x75, 0x68, 0x88, 0x02, 0x1e, 0x2b, 0x45, 0xad,
	0xd4, 0x95, 0xfa, 0x7f, 0x8d, 0x6c, 0x5e, 0xe3, 0x62, 0x7f, 0x6b, 0xbf, 0x43, 0xe5, 0x53, 0xf8,
	0x07, 0x39, 0x9c, 0xce, 0x11, 0xa7, 0x2c, 0x98, 0x70, 0x2f, 0x22, 0xb1, 0xc7, 0x66, 0x58, 0x29,
	0x69, 0xa0, 0xfb, 0xcb, 0x82, 0x9b, 0x55, 0xbb, 0x72, 0x4d, 0xa8, 0xeb, 0x71, 0xfb, 0xf7, 0x9e,
	0xbb, 0xcc, 0x30, 0x79, 0x00, 0x1b, 0x08, 0xfb, 0x34, 0x9f, 0x2c, 0x1f, 0x24, 0xeb, 0x02, 0xd9,
	0x85, 0xf4, 0x04, 0x4a, 0xb9, 0x81, 0x64, 0x0b, 0xfe, 0x8c, 0xa9, 0x1b, 0x20, 0x9e, 0x44, 0x44,
	0xec, 0x5a, 0xb3, 0x3a, 0x9b, 0x55, 0x5b, 0xfb, 0xc8, 0xb1, 0x71, 0x86, 0x71, 0x44, 0xe2, 0xd8,
	0xde, 0xc7, 0x64, 0x1d, 0x56, 0x16, 0xa2, 0x99, 0x52, 0x3c, 0x68, 0xbf, 0xbd, 0xd1, 0x5f, 0x00,
	0x6c, 0x0e, 0x23, 0x82, 0x38, 0xc9, 0x2c, 0x8f, 0x63, 0xf7, 0x5b, 0x8b, 0xbe, 0x2f, 0xc2, 0xe6,
	0x55, 0x88, 0x8f, 0xd9, 0xd8, 0x84, 0x92, 0xb3, 0xcd, 0x4e, 0x28, 0x16, 0x76, 0x6b, 0x56, 0x3d,
	0x5d, 0xb5, 0x61, 0xf6, 0xcb, 0xd1, 0xb9, 0x0d, 0x33, 0x64, 0x84, 0x0f, 0x14, 0x95, 0x8e, 0x57,
	0x54, 0xfe, 0xb2, 0xa2, 0x1f, 0x9f, 0x29, 0xb2, 0x94, 0x87, 0x54, 0x05, 0xcb, 0x54, 0x05, 0x4f,
	0xa9, 0x0a, 0xee, 0xd6, 0x6a, 0x61, 0xb9, 0x56, 0x0b, 0x8f, 0x6b, 0xb5, 0x30, 0xad, 0x88, 0x57,
	0x39, 0x78, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x71, 0x30, 0x7c, 0x0e, 0x04, 0x00, 0x00,
}

func (m *Contract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contract) MarshalTo(dAtA []byte) (int, error) {
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
	if len(m.Participants) > 0 {
		for _, msg := range m.Participants {
			dAtA[i] = 0x12
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.ActivationThreshold != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.ActivationThreshold))
	}
	if m.AdminThreshold != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.AdminThreshold))
	}
	return i, nil
}

func (m *Participant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Participant) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Signature)))
		i += copy(dAtA[i:], m.Signature)
	}
	if m.Weight != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Weight))
	}
	return i, nil
}

func (m *CreateContractMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateContractMsg) MarshalTo(dAtA []byte) (int, error) {
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
	if len(m.Participants) > 0 {
		for _, msg := range m.Participants {
			dAtA[i] = 0x12
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.ActivationThreshold != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.ActivationThreshold))
	}
	if m.AdminThreshold != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.AdminThreshold))
	}
	return i, nil
}

func (m *UpdateContractMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateContractMsg) MarshalTo(dAtA []byte) (int, error) {
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
	if len(m.ContractID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.ContractID)))
		i += copy(dAtA[i:], m.ContractID)
	}
	if len(m.Participants) > 0 {
		for _, msg := range m.Participants {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.ActivationThreshold != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.ActivationThreshold))
	}
	if m.AdminThreshold != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.AdminThreshold))
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
func (m *Contract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Participants) > 0 {
		for _, e := range m.Participants {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	if m.ActivationThreshold != 0 {
		n += 1 + sovCodec(uint64(m.ActivationThreshold))
	}
	if m.AdminThreshold != 0 {
		n += 1 + sovCodec(uint64(m.AdminThreshold))
	}
	return n
}

func (m *Participant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Weight != 0 {
		n += 1 + sovCodec(uint64(m.Weight))
	}
	return n
}

func (m *CreateContractMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Participants) > 0 {
		for _, e := range m.Participants {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	if m.ActivationThreshold != 0 {
		n += 1 + sovCodec(uint64(m.ActivationThreshold))
	}
	if m.AdminThreshold != 0 {
		n += 1 + sovCodec(uint64(m.AdminThreshold))
	}
	return n
}

func (m *UpdateContractMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.ContractID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Participants) > 0 {
		for _, e := range m.Participants {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	if m.ActivationThreshold != 0 {
		n += 1 + sovCodec(uint64(m.ActivationThreshold))
	}
	if m.AdminThreshold != 0 {
		n += 1 + sovCodec(uint64(m.AdminThreshold))
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
func (m *Contract) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Contract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contract: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
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
			m.Participants = append(m.Participants, &Participant{})
			if err := m.Participants[len(m.Participants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivationThreshold", wireType)
			}
			m.ActivationThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActivationThreshold |= Weight(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminThreshold", wireType)
			}
			m.AdminThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdminThreshold |= Weight(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *Participant) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Participant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Participant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
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
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= Weight(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *CreateContractMsg) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CreateContractMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateContractMsg: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
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
			m.Participants = append(m.Participants, &Participant{})
			if err := m.Participants[len(m.Participants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivationThreshold", wireType)
			}
			m.ActivationThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActivationThreshold |= Weight(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminThreshold", wireType)
			}
			m.AdminThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdminThreshold |= Weight(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *UpdateContractMsg) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UpdateContractMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateContractMsg: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field ContractID", wireType)
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
			m.ContractID = append(m.ContractID[:0], dAtA[iNdEx:postIndex]...)
			if m.ContractID == nil {
				m.ContractID = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
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
			m.Participants = append(m.Participants, &Participant{})
			if err := m.Participants[len(m.Participants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivationThreshold", wireType)
			}
			m.ActivationThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActivationThreshold |= Weight(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminThreshold", wireType)
			}
			m.AdminThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdminThreshold |= Weight(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
