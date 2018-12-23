// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/capture.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/altipla-consulting/envoy-api/envoy/api/v2/core"
import types "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Connection properties.
type Connection struct {
	// Global unique connection ID for Envoy session. Matches connection IDs used
	// in Envoy logs.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Local address.
	LocalAddress *core.Address `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	// Remote address.
	RemoteAddress        *core.Address `protobuf:"bytes,3,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_capture_057b03d552469814, []int{0}
}
func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (dst *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(dst, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Connection) GetLocalAddress() *core.Address {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *Connection) GetRemoteAddress() *core.Address {
	if m != nil {
		return m.RemoteAddress
	}
	return nil
}

// Event in a capture trace.
type Event struct {
	// Timestamp for event.
	Timestamp *types.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Read or write with content as bytes string.
	//
	// Types that are valid to be assigned to EventSelector:
	//	*Event_Read_
	//	*Event_Write_
	EventSelector        isEvent_EventSelector `protobuf_oneof:"event_selector"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_capture_057b03d552469814, []int{1}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type isEvent_EventSelector interface {
	isEvent_EventSelector()
}

type Event_Read_ struct {
	Read *Event_Read `protobuf:"bytes,2,opt,name=read,proto3,oneof"`
}

type Event_Write_ struct {
	Write *Event_Write `protobuf:"bytes,3,opt,name=write,proto3,oneof"`
}

func (*Event_Read_) isEvent_EventSelector() {}

func (*Event_Write_) isEvent_EventSelector() {}

func (m *Event) GetEventSelector() isEvent_EventSelector {
	if m != nil {
		return m.EventSelector
	}
	return nil
}

func (m *Event) GetRead() *Event_Read {
	if x, ok := m.GetEventSelector().(*Event_Read_); ok {
		return x.Read
	}
	return nil
}

func (m *Event) GetWrite() *Event_Write {
	if x, ok := m.GetEventSelector().(*Event_Write_); ok {
		return x.Write
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_Read_)(nil),
		(*Event_Write_)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// event_selector
	switch x := m.EventSelector.(type) {
	case *Event_Read_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Read); err != nil {
			return err
		}
	case *Event_Write_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Write); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.EventSelector has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 2: // event_selector.read
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Event_Read)
		err := b.DecodeMessage(msg)
		m.EventSelector = &Event_Read_{msg}
		return true, err
	case 3: // event_selector.write
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Event_Write)
		err := b.DecodeMessage(msg)
		m.EventSelector = &Event_Write_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// event_selector
	switch x := m.EventSelector.(type) {
	case *Event_Read_:
		s := proto.Size(x.Read)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Write_:
		s := proto.Size(x.Write)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Data read by Envoy from the transport socket.
type Event_Read struct {
	// Binary data read.
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_Read) Reset()         { *m = Event_Read{} }
func (m *Event_Read) String() string { return proto.CompactTextString(m) }
func (*Event_Read) ProtoMessage()    {}
func (*Event_Read) Descriptor() ([]byte, []int) {
	return fileDescriptor_capture_057b03d552469814, []int{1, 0}
}
func (m *Event_Read) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_Read.Unmarshal(m, b)
}
func (m *Event_Read) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_Read.Marshal(b, m, deterministic)
}
func (dst *Event_Read) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_Read.Merge(dst, src)
}
func (m *Event_Read) XXX_Size() int {
	return xxx_messageInfo_Event_Read.Size(m)
}
func (m *Event_Read) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_Read.DiscardUnknown(m)
}

var xxx_messageInfo_Event_Read proto.InternalMessageInfo

func (m *Event_Read) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Data written by Envoy to the transport socket.
type Event_Write struct {
	// Binary data written.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Stream was half closed after this write.
	EndStream            bool     `protobuf:"varint,2,opt,name=end_stream,json=endStream,proto3" json:"end_stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_Write) Reset()         { *m = Event_Write{} }
func (m *Event_Write) String() string { return proto.CompactTextString(m) }
func (*Event_Write) ProtoMessage()    {}
func (*Event_Write) Descriptor() ([]byte, []int) {
	return fileDescriptor_capture_057b03d552469814, []int{1, 1}
}
func (m *Event_Write) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_Write.Unmarshal(m, b)
}
func (m *Event_Write) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_Write.Marshal(b, m, deterministic)
}
func (dst *Event_Write) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_Write.Merge(dst, src)
}
func (m *Event_Write) XXX_Size() int {
	return xxx_messageInfo_Event_Write.Size(m)
}
func (m *Event_Write) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_Write.DiscardUnknown(m)
}

var xxx_messageInfo_Event_Write proto.InternalMessageInfo

func (m *Event_Write) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Event_Write) GetEndStream() bool {
	if m != nil {
		return m.EndStream
	}
	return false
}

// Sequence of read/write events that constitute a captured trace on a socket.
// Multiple Trace messages might be emitted for a given connection ID, with the
// sink (e.g. file set, network) responsible for later reassembly.
type Trace struct {
	// Connection properties.
	Connection *Connection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	// Sequence of observed events.
	Events               []*Event `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_capture_057b03d552469814, []int{2}
}
func (m *Trace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trace.Unmarshal(m, b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
}
func (dst *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(dst, src)
}
func (m *Trace) XXX_Size() int {
	return xxx_messageInfo_Trace.Size(m)
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *Trace) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Connection)(nil), "envoy.data.tap.v2alpha.Connection")
	proto.RegisterType((*Event)(nil), "envoy.data.tap.v2alpha.Event")
	proto.RegisterType((*Event_Read)(nil), "envoy.data.tap.v2alpha.Event.Read")
	proto.RegisterType((*Event_Write)(nil), "envoy.data.tap.v2alpha.Event.Write")
	proto.RegisterType((*Trace)(nil), "envoy.data.tap.v2alpha.Trace")
}

func init() {
	proto.RegisterFile("envoy/data/tap/v2alpha/capture.proto", fileDescriptor_capture_057b03d552469814)
}

var fileDescriptor_capture_057b03d552469814 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x6b, 0xdc, 0x30,
	0x10, 0xc5, 0x63, 0xc7, 0x0e, 0xcd, 0x24, 0x59, 0x8a, 0x0e, 0x25, 0x18, 0x42, 0x82, 0xdb, 0xc3,
	0x9e, 0x24, 0x70, 0x29, 0x84, 0xf6, 0x50, 0xb2, 0xa5, 0x90, 0xb3, 0x1a, 0x28, 0xf4, 0xb2, 0x4c,
	0xac, 0x69, 0x6a, 0xb0, 0x2d, 0x21, 0x2b, 0x2e, 0xbd, 0xf6, 0x93, 0x94, 0x7e, 0xd2, 0xa2, 0x3f,
	0xd9, 0xed, 0x61, 0xbb, 0xbd, 0x79, 0xac, 0xdf, 0x7b, 0xf3, 0x66, 0x06, 0x5e, 0xd1, 0x38, 0xeb,
	0x1f, 0x42, 0xa1, 0x43, 0xe1, 0xd0, 0x88, 0xb9, 0xc1, 0xde, 0x7c, 0x43, 0xd1, 0xa2, 0x71, 0x8f,
	0x96, 0xb8, 0xb1, 0xda, 0x69, 0xf6, 0x22, 0x50, 0xdc, 0x53, 0xdc, 0xa1, 0xe1, 0x89, 0xaa, 0x2e,
	0xa3, 0x1a, 0x4d, 0x27, 0xe6, 0x46, 0xb4, 0xda, 0x92, 0x40, 0xa5, 0x2c, 0x4d, 0x53, 0x14, 0x56,
	0x97, 0x0f, 0x5a, 0x3f, 0xf4, 0x24, 0x42, 0x75, 0xff, 0xf8, 0x55, 0xb8, 0x6e, 0xa0, 0xc9, 0xe1,
	0x60, 0x22, 0x50, 0xff, 0xca, 0x00, 0x3e, 0xe8, 0x71, 0xa4, 0xd6, 0x75, 0x7a, 0x64, 0x0b, 0xc8,
	0x3b, 0x75, 0x9e, 0x5d, 0x65, 0xcb, 0x42, 0xe6, 0x9d, 0x62, 0xef, 0xe1, 0xac, 0xd7, 0x2d, 0xf6,
	0xeb, 0x64, 0x7b, 0x9e, 0x5f, 0x65, 0xcb, 0x93, 0xa6, 0xe2, 0x31, 0x10, 0x9a, 0x8e, 0xcf, 0x0d,
	0xf7, 0x8d, 0xf9, 0x4d, 0x24, 0xe4, 0x69, 0x10, 0xa4, 0x8a, 0xdd, 0xc0, 0xc2, 0xd2, 0xa0, 0x1d,
	0x6d, 0x1c, 0x0e, 0xff, 0xeb, 0x70, 0x16, 0x15, 0xa9, 0xac, 0x7f, 0xe7, 0x50, 0x7e, 0x9c, 0x69,
	0x74, 0xec, 0x1a, 0x8e, 0x37, 0xf9, 0x43, 0x48, 0xef, 0x13, 0x27, 0xe4, 0x4f, 0x13, 0xf2, 0xbb,
	0x27, 0x42, 0x6e, 0x61, 0x76, 0x0d, 0x85, 0x25, 0x54, 0x29, 0x7e, 0xcd, 0x77, 0xef, 0x93, 0x87,
	0x36, 0x5c, 0x12, 0xaa, 0xdb, 0x03, 0x19, 0x14, 0xec, 0x1d, 0x94, 0xdf, 0x6d, 0xe7, 0x28, 0xe5,
	0x7e, 0xb9, 0x5f, 0xfa, 0xd9, 0xa3, 0xb7, 0x07, 0x32, 0x6a, 0xaa, 0x0a, 0x0a, 0x6f, 0xc6, 0x18,
	0x14, 0x5e, 0x10, 0x32, 0x9f, 0xca, 0xf0, 0x5d, 0xbd, 0x85, 0x32, 0xd0, 0xbb, 0x1e, 0xd9, 0x05,
	0x00, 0x8d, 0x6a, 0x3d, 0x39, 0x4b, 0x38, 0x84, 0xd4, 0xcf, 0xe4, 0x31, 0x8d, 0xea, 0x53, 0xf8,
	0xb1, 0x7a, 0x0e, 0x0b, 0xf2, 0xfd, 0xd6, 0x13, 0xf5, 0xd4, 0x3a, 0x6d, 0xeb, 0x9f, 0x19, 0x94,
	0x77, 0x16, 0x5b, 0x62, 0x2b, 0x80, 0x76, 0x73, 0xd0, 0xb4, 0xa5, 0x7f, 0x0e, 0xbc, 0x3d, 0xbd,
	0xfc, 0x4b, 0xc5, 0xde, 0xc0, 0x51, 0xf0, 0xf7, 0xf7, 0x3e, 0x5c, 0x9e, 0x34, 0x17, 0x7b, 0xa7,
	0x96, 0x09, 0x5e, 0x15, 0x5f, 0xf2, 0xb9, 0xb9, 0x3f, 0x0a, 0xa7, 0x78, 0xfd, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xb6, 0x5a, 0xf3, 0xab, 0xdb, 0x02, 0x00, 0x00,
}
