// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/gov/sample_test.proto

package gov

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

// ProposalOptions is a sum type of all possible messages that
// may be dispatched via a governance proposal.
//
// For the test case, we only refer to package-internal messages
// and handlers, but an application can reference messages from any package.
type ProposalOptions struct {
	// Types that are valid to be assigned to Option:
	//	*ProposalOptions_Text
	//	*ProposalOptions_Electorate
	//	*ProposalOptions_Rule
	Option isProposalOptions_Option `protobuf_oneof:"option"`
}

func (m *ProposalOptions) Reset()         { *m = ProposalOptions{} }
func (m *ProposalOptions) String() string { return proto.CompactTextString(m) }
func (*ProposalOptions) ProtoMessage()    {}
func (*ProposalOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3ad181f69ab09f1, []int{0}
}
func (m *ProposalOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalOptions.Merge(m, src)
}
func (m *ProposalOptions) XXX_Size() int {
	return m.Size()
}
func (m *ProposalOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalOptions proto.InternalMessageInfo

type isProposalOptions_Option interface {
	isProposalOptions_Option()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ProposalOptions_Text struct {
	Text *TextResolutionMsg `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}
type ProposalOptions_Electorate struct {
	Electorate *UpdateElectorateMsg `protobuf:"bytes,2,opt,name=electorate,proto3,oneof"`
}
type ProposalOptions_Rule struct {
	Rule *UpdateElectionRuleMsg `protobuf:"bytes,3,opt,name=rule,proto3,oneof"`
}

func (*ProposalOptions_Text) isProposalOptions_Option()       {}
func (*ProposalOptions_Electorate) isProposalOptions_Option() {}
func (*ProposalOptions_Rule) isProposalOptions_Option()       {}

func (m *ProposalOptions) GetOption() isProposalOptions_Option {
	if m != nil {
		return m.Option
	}
	return nil
}

func (m *ProposalOptions) GetText() *TextResolutionMsg {
	if x, ok := m.GetOption().(*ProposalOptions_Text); ok {
		return x.Text
	}
	return nil
}

func (m *ProposalOptions) GetElectorate() *UpdateElectorateMsg {
	if x, ok := m.GetOption().(*ProposalOptions_Electorate); ok {
		return x.Electorate
	}
	return nil
}

func (m *ProposalOptions) GetRule() *UpdateElectionRuleMsg {
	if x, ok := m.GetOption().(*ProposalOptions_Rule); ok {
		return x.Rule
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ProposalOptions) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ProposalOptions_OneofMarshaler, _ProposalOptions_OneofUnmarshaler, _ProposalOptions_OneofSizer, []interface{}{
		(*ProposalOptions_Text)(nil),
		(*ProposalOptions_Electorate)(nil),
		(*ProposalOptions_Rule)(nil),
	}
}

func _ProposalOptions_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ProposalOptions)
	// option
	switch x := m.Option.(type) {
	case *ProposalOptions_Text:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Text); err != nil {
			return err
		}
	case *ProposalOptions_Electorate:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Electorate); err != nil {
			return err
		}
	case *ProposalOptions_Rule:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rule); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ProposalOptions.Option has unexpected type %T", x)
	}
	return nil
}

func _ProposalOptions_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ProposalOptions)
	switch tag {
	case 1: // option.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TextResolutionMsg)
		err := b.DecodeMessage(msg)
		m.Option = &ProposalOptions_Text{msg}
		return true, err
	case 2: // option.electorate
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateElectorateMsg)
		err := b.DecodeMessage(msg)
		m.Option = &ProposalOptions_Electorate{msg}
		return true, err
	case 3: // option.rule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateElectionRuleMsg)
		err := b.DecodeMessage(msg)
		m.Option = &ProposalOptions_Rule{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ProposalOptions_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ProposalOptions)
	// option
	switch x := m.Option.(type) {
	case *ProposalOptions_Text:
		s := proto.Size(x.Text)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProposalOptions_Electorate:
		s := proto.Size(x.Electorate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ProposalOptions_Rule:
		s := proto.Size(x.Rule)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// AppCreateProposalMsg is byte compatible with CreateProposalMsg
// assuming we use decodeProposalOptions.
//
// It is here to demonstrate a clean app-level usage of proposals.
type AppCreateProposalMsg struct {
	Metadata *weave.Metadata        `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Base     *CreateProposalMsgBase `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Options  *ProposalOptions       `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
}

func (m *AppCreateProposalMsg) Reset()         { *m = AppCreateProposalMsg{} }
func (m *AppCreateProposalMsg) String() string { return proto.CompactTextString(m) }
func (*AppCreateProposalMsg) ProtoMessage()    {}
func (*AppCreateProposalMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3ad181f69ab09f1, []int{1}
}
func (m *AppCreateProposalMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppCreateProposalMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppCreateProposalMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppCreateProposalMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppCreateProposalMsg.Merge(m, src)
}
func (m *AppCreateProposalMsg) XXX_Size() int {
	return m.Size()
}
func (m *AppCreateProposalMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_AppCreateProposalMsg.DiscardUnknown(m)
}

var xxx_messageInfo_AppCreateProposalMsg proto.InternalMessageInfo

func (m *AppCreateProposalMsg) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *AppCreateProposalMsg) GetBase() *CreateProposalMsgBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *AppCreateProposalMsg) GetOptions() *ProposalOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

// AppProposal is byte compatible with Proposal
// assuming we use decodeProposalOptions.
//
// It is here to demonstrate a clean app-level usage of proposals.
type AppProposal struct {
	Metadata *weave.Metadata  `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Common   *ProposalCommon  `protobuf:"bytes,2,opt,name=common,proto3" json:"common,omitempty"`
	Options  *ProposalOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
}

func (m *AppProposal) Reset()         { *m = AppProposal{} }
func (m *AppProposal) String() string { return proto.CompactTextString(m) }
func (*AppProposal) ProtoMessage()    {}
func (*AppProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3ad181f69ab09f1, []int{2}
}
func (m *AppProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppProposal.Merge(m, src)
}
func (m *AppProposal) XXX_Size() int {
	return m.Size()
}
func (m *AppProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AppProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AppProposal proto.InternalMessageInfo

func (m *AppProposal) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *AppProposal) GetCommon() *ProposalCommon {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *AppProposal) GetOptions() *ProposalOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposalOptions)(nil), "gov.ProposalOptions")
	proto.RegisterType((*AppCreateProposalMsg)(nil), "gov.AppCreateProposalMsg")
	proto.RegisterType((*AppProposal)(nil), "gov.AppProposal")
}

func init() { proto.RegisterFile("x/gov/sample_test.proto", fileDescriptor_a3ad181f69ab09f1) }

var fileDescriptor_a3ad181f69ab09f1 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xe3, 0xdb, 0xaa, 0xb7, 0x72, 0x87, 0xea, 0xe6, 0x56, 0xf7, 0x46, 0x1d, 0x22, 0x54,
	0x31, 0x20, 0x95, 0x3a, 0x08, 0x36, 0xb6, 0xb6, 0x42, 0x62, 0xa9, 0x40, 0x11, 0xcc, 0xc8, 0x49,
	0x0f, 0x26, 0x52, 0x92, 0x63, 0xc5, 0x4e, 0xe8, 0x63, 0xb0, 0xf1, 0x00, 0x3c, 0x04, 0xaf, 0xc0,
	0xd8, 0x91, 0x11, 0xb5, 0x2f, 0x82, 0x70, 0x92, 0xaa, 0x94, 0xa9, 0x5b, 0xec, 0xff, 0xfb, 0x72,
	0x7e, 0x5b, 0xa6, 0xff, 0x17, 0x9e, 0xc0, 0xc2, 0x53, 0x3c, 0x91, 0x31, 0xdc, 0x69, 0x50, 0x9a,
	0xc9, 0x0c, 0x35, 0xda, 0x0d, 0x81, 0x45, 0x7f, 0x24, 0x22, 0xfd, 0x90, 0x07, 0x2c, 0xc4, 0xc4,
	0x13, 0x28, 0xd0, 0x33, 0x59, 0x90, 0xdf, 0x9b, 0x95, 0x59, 0x98, 0xaf, 0xd2, 0xe9, 0x1f, 0x6e,
	0xe1, 0x11, 0x16, 0x23, 0x4c, 0xc1, 0x7b, 0x04, 0x5e, 0x80, 0x17, 0xe2, 0x1c, 0xc2, 0x8a, 0xfa,
	0x53, 0x8e, 0xdc, 0xda, 0x1a, 0xbc, 0x12, 0xda, 0xbd, 0xce, 0x50, 0xa2, 0xe2, 0xf1, 0x95, 0xd4,
	0x11, 0xa6, 0xca, 0x3e, 0xa6, 0x4d, 0x0d, 0x0b, 0xed, 0x90, 0x03, 0x72, 0xd4, 0x39, 0xfd, 0xc7,
	0x04, 0x16, 0xec, 0x06, 0x16, 0xda, 0x07, 0x85, 0x71, 0xfe, 0xc5, 0xcc, 0x94, 0xb8, 0xb4, 0x7c,
	0x43, 0xd9, 0xe7, 0x94, 0x42, 0x0c, 0xa1, 0xc6, 0x8c, 0x6b, 0x70, 0x7e, 0x19, 0xc7, 0x31, 0xce,
	0xad, 0x9c, 0x73, 0x0d, 0x17, 0x9b, 0xb0, 0xb4, 0xb6, 0x68, 0xfb, 0x84, 0x36, 0xb3, 0x3c, 0x06,
	0xa7, 0x61, 0xac, 0xfe, 0xae, 0x15, 0x61, 0xea, 0xe7, 0x71, 0xe5, 0x19, 0x72, 0xd2, 0xa6, 0x2d,
	0x34, 0x35, 0x07, 0x2f, 0x84, 0xf6, 0xc6, 0x52, 0x4e, 0x33, 0xe0, 0x1a, 0xea, 0x23, 0xcc, 0x94,
	0xb0, 0x87, 0xb4, 0x9d, 0x80, 0xe6, 0x73, 0xae, 0x79, 0x75, 0x84, 0x2e, 0x33, 0x77, 0xc1, 0x66,
	0xd5, 0xb6, 0xbf, 0x01, 0x6c, 0x46, 0x9b, 0x01, 0x57, 0x75, 0xef, 0xb2, 0xc1, 0x8f, 0x5f, 0x4e,
	0xb8, 0x02, 0xdf, 0x70, 0x36, 0xa3, 0xbf, 0xcb, 0xf9, 0xaa, 0x2a, 0xdd, 0x33, 0xca, 0xce, 0x15,
	0xfa, 0x35, 0x34, 0x78, 0x26, 0xb4, 0x33, 0x96, 0xb2, 0xce, 0xf7, 0x2b, 0x37, 0xa4, 0xad, 0x10,
	0x93, 0x04, 0xd3, 0xaa, 0xde, 0xdf, 0x6f, 0xb3, 0xa6, 0x26, 0xf2, 0x2b, 0x64, 0xdf, 0x66, 0x13,
	0xe7, 0x6d, 0xe5, 0x92, 0xe5, 0xca, 0x25, 0x1f, 0x2b, 0x97, 0x3c, 0xad, 0x5d, 0x6b, 0xb9, 0x76,
	0xad, 0xf7, 0xb5, 0x6b, 0x05, 0x2d, 0xf3, 0x34, 0xce, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbd,
	0x2b, 0x18, 0xbe, 0xa2, 0x02, 0x00, 0x00,
}

func (m *ProposalOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Option != nil {
		nn1, err := m.Option.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *ProposalOptions_Text) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Text != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSampleTest(dAtA, i, uint64(m.Text.Size()))
		n2, err := m.Text.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *ProposalOptions_Electorate) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Electorate != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSampleTest(dAtA, i, uint64(m.Electorate.Size()))
		n3, err := m.Electorate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *ProposalOptions_Rule) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Rule != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSampleTest(dAtA, i, uint64(m.Rule.Size()))
		n4, err := m.Rule.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *AppCreateProposalMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppCreateProposalMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSampleTest(dAtA, i, uint64(m.Metadata.Size()))
		n5, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Base != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSampleTest(dAtA, i, uint64(m.Base.Size()))
		n6, err := m.Base.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.Options != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSampleTest(dAtA, i, uint64(m.Options.Size()))
		n7, err := m.Options.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func (m *AppProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppProposal) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSampleTest(dAtA, i, uint64(m.Metadata.Size()))
		n8, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.Common != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSampleTest(dAtA, i, uint64(m.Common.Size()))
		n9, err := m.Common.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.Options != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSampleTest(dAtA, i, uint64(m.Options.Size()))
		n10, err := m.Options.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	return i, nil
}

func encodeVarintSampleTest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ProposalOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Option != nil {
		n += m.Option.Size()
	}
	return n
}

func (m *ProposalOptions_Text) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Text != nil {
		l = m.Text.Size()
		n += 1 + l + sovSampleTest(uint64(l))
	}
	return n
}
func (m *ProposalOptions_Electorate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Electorate != nil {
		l = m.Electorate.Size()
		n += 1 + l + sovSampleTest(uint64(l))
	}
	return n
}
func (m *ProposalOptions_Rule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Rule != nil {
		l = m.Rule.Size()
		n += 1 + l + sovSampleTest(uint64(l))
	}
	return n
}
func (m *AppCreateProposalMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovSampleTest(uint64(l))
	}
	if m.Base != nil {
		l = m.Base.Size()
		n += 1 + l + sovSampleTest(uint64(l))
	}
	if m.Options != nil {
		l = m.Options.Size()
		n += 1 + l + sovSampleTest(uint64(l))
	}
	return n
}

func (m *AppProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovSampleTest(uint64(l))
	}
	if m.Common != nil {
		l = m.Common.Size()
		n += 1 + l + sovSampleTest(uint64(l))
	}
	if m.Options != nil {
		l = m.Options.Size()
		n += 1 + l + sovSampleTest(uint64(l))
	}
	return n
}

func sovSampleTest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSampleTest(x uint64) (n int) {
	return sovSampleTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposalOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSampleTest
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
			return fmt.Errorf("proto: ProposalOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSampleTest
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
				return ErrInvalidLengthSampleTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSampleTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TextResolutionMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Option = &ProposalOptions_Text{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Electorate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSampleTest
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
				return ErrInvalidLengthSampleTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSampleTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &UpdateElectorateMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Option = &ProposalOptions_Electorate{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rule", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSampleTest
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
				return ErrInvalidLengthSampleTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSampleTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &UpdateElectionRuleMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Option = &ProposalOptions_Rule{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSampleTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSampleTest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSampleTest
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
func (m *AppCreateProposalMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSampleTest
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
			return fmt.Errorf("proto: AppCreateProposalMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppCreateProposalMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSampleTest
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
				return ErrInvalidLengthSampleTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSampleTest
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
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSampleTest
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
				return ErrInvalidLengthSampleTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSampleTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Base == nil {
				m.Base = &CreateProposalMsgBase{}
			}
			if err := m.Base.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSampleTest
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
				return ErrInvalidLengthSampleTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSampleTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Options == nil {
				m.Options = &ProposalOptions{}
			}
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSampleTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSampleTest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSampleTest
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
func (m *AppProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSampleTest
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
			return fmt.Errorf("proto: AppProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSampleTest
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
				return ErrInvalidLengthSampleTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSampleTest
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
				return fmt.Errorf("proto: wrong wireType = %d for field Common", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSampleTest
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
				return ErrInvalidLengthSampleTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSampleTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Common == nil {
				m.Common = &ProposalCommon{}
			}
			if err := m.Common.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSampleTest
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
				return ErrInvalidLengthSampleTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSampleTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Options == nil {
				m.Options = &ProposalOptions{}
			}
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSampleTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSampleTest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSampleTest
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
func skipSampleTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSampleTest
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
					return 0, ErrIntOverflowSampleTest
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
					return 0, ErrIntOverflowSampleTest
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
				return 0, ErrInvalidLengthSampleTest
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSampleTest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSampleTest
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
				next, err := skipSampleTest(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSampleTest
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
	ErrInvalidLengthSampleTest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSampleTest   = fmt.Errorf("proto: integer overflow")
)