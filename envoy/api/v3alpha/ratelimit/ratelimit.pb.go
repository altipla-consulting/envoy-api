// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/ratelimit/ratelimit.proto

package envoy_api_v3alpha_ratelimit

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A RateLimitDescriptor is a list of hierarchical entries that are used by the service to
// determine the final rate limit key and overall allowed limit. Here are some examples of how
// they might be used for the domain "envoy".
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["remote_address": "10.0.0.1"]
//
// What it does: Limits all unauthenticated traffic for the IP address 10.0.0.1. The
// configuration supplies a default limit for the *remote_address* key. If there is a desire to
// raise the limit for 10.0.0.1 or block it entirely it can be specified directly in the
// configuration.
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["path": "/foo/bar"]
//
// What it does: Limits all unauthenticated traffic globally for a specific path (or prefix if
// configured that way in the service).
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["path": "/foo/bar"], ["remote_address": "10.0.0.1"]
//
// What it does: Limits unauthenticated traffic to a specific path for a specific IP address.
// Like (1) we can raise/block specific IP addresses if we want with an override configuration.
//
// .. code-block:: cpp
//
//   ["authenticated": "true"], ["client_id": "foo"]
//
// What it does: Limits all traffic for an authenticated client "foo"
//
// .. code-block:: cpp
//
//   ["authenticated": "true"], ["client_id": "foo"], ["path": "/foo/bar"]
//
// What it does: Limits traffic to a specific path for an authenticated client "foo"
//
// The idea behind the API is that (1)/(2)/(3) and (4)/(5) can be sent in 1 request if desired.
// This enables building complex application scenarios with a generic backend.
type RateLimitDescriptor struct {
	// Descriptor entries.
	Entries              []*RateLimitDescriptor_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitDescriptor) Reset()         { *m = RateLimitDescriptor{} }
func (m *RateLimitDescriptor) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor) ProtoMessage()    {}
func (*RateLimitDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fa87ac9676b2311, []int{0}
}

func (m *RateLimitDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitDescriptor.Unmarshal(m, b)
}
func (m *RateLimitDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitDescriptor.Marshal(b, m, deterministic)
}
func (m *RateLimitDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor.Merge(m, src)
}
func (m *RateLimitDescriptor) XXX_Size() int {
	return xxx_messageInfo_RateLimitDescriptor.Size(m)
}
func (m *RateLimitDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor proto.InternalMessageInfo

func (m *RateLimitDescriptor) GetEntries() []*RateLimitDescriptor_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type RateLimitDescriptor_Entry struct {
	// Descriptor key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Descriptor value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitDescriptor_Entry) Reset()         { *m = RateLimitDescriptor_Entry{} }
func (m *RateLimitDescriptor_Entry) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor_Entry) ProtoMessage()    {}
func (*RateLimitDescriptor_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fa87ac9676b2311, []int{0, 0}
}

func (m *RateLimitDescriptor_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Unmarshal(m, b)
}
func (m *RateLimitDescriptor_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Marshal(b, m, deterministic)
}
func (m *RateLimitDescriptor_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor_Entry.Merge(m, src)
}
func (m *RateLimitDescriptor_Entry) XXX_Size() int {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Size(m)
}
func (m *RateLimitDescriptor_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor_Entry proto.InternalMessageInfo

func (m *RateLimitDescriptor_Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RateLimitDescriptor_Entry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*RateLimitDescriptor)(nil), "envoy.api.v3alpha.ratelimit.RateLimitDescriptor")
	proto.RegisterType((*RateLimitDescriptor_Entry)(nil), "envoy.api.v3alpha.ratelimit.RateLimitDescriptor.Entry")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/ratelimit/ratelimit.proto", fileDescriptor_3fa87ac9676b2311)
}

var fileDescriptor_3fa87ac9676b2311 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0x4e, 0xcc, 0x29, 0xc8, 0x48, 0xd4, 0x2f, 0x4a,
	0x2c, 0x49, 0xcd, 0xc9, 0xcc, 0xcd, 0x2c, 0x41, 0xb0, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xa4, 0xc1, 0x8a, 0xf5, 0x12, 0x0b, 0x32, 0xf5, 0xa0, 0x8a, 0xf5, 0xe0, 0x4a, 0xa4, 0xc4, 0xcb,
	0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0x2e, 0xa5, 0x2d, 0x8c, 0x5c,
	0xc2, 0x41, 0x89, 0x25, 0xa9, 0x3e, 0x20, 0x65, 0x2e, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25,
	0xf9, 0x45, 0x42, 0x51, 0x5c, 0xec, 0xa9, 0x79, 0x25, 0x45, 0x99, 0xa9, 0xc5, 0x12, 0x8c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0x66, 0x7a, 0x78, 0xcc, 0xd7, 0xc3, 0x62, 0x84, 0x9e, 0x6b, 0x5e, 0x49,
	0x51, 0xa5, 0x13, 0xc7, 0x2f, 0x27, 0xd6, 0x49, 0x8c, 0x4c, 0x1c, 0x8c, 0x41, 0x30, 0x03, 0xa5,
	0x1c, 0xb9, 0x58, 0xc1, 0x72, 0x42, 0x92, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x4e, 0xec, 0xbf, 0x9c, 0x58, 0x8a, 0x98, 0x14, 0x18, 0x83, 0x40, 0x62, 0x42, 0xb2,
	0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0xa8, 0x92, 0x10, 0x51, 0x27, 0x5b, 0x2e,
	0xcd, 0xcc, 0x7c, 0x88, 0x8b, 0x0a, 0x8a, 0xf2, 0x2b, 0x2a, 0xf1, 0x39, 0xce, 0x89, 0x2f, 0x08,
	0xc6, 0x0c, 0x00, 0xf9, 0x39, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x79, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf2, 0x0e, 0xd5, 0xc2, 0x61, 0x01, 0x00, 0x00,
}
