// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/transport_socket/tap/v2alpha/tap.proto

package v2

import (
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v2/core"
	v2alpha "github.com/altipla-consulting/envoy-api/envoy/config/common/tap/v2alpha"
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

// Configuration for tap transport socket. This wraps another transport socket, providing the
// ability to interpose and record in plain text any traffic that is surfaced to Envoy.
type Tap struct {
	// Common configuration for the tap transport socket.
	CommonConfig *v2alpha.CommonExtensionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// The underlying transport socket being wrapped.
	TransportSocket      *core.TransportSocket `protobuf:"bytes,2,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Tap) Reset()         { *m = Tap{} }
func (m *Tap) String() string { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()    {}
func (*Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_07cb8c0b42756e40, []int{0}
}

func (m *Tap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tap.Unmarshal(m, b)
}
func (m *Tap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tap.Marshal(b, m, deterministic)
}
func (m *Tap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tap.Merge(m, src)
}
func (m *Tap) XXX_Size() int {
	return xxx_messageInfo_Tap.Size(m)
}
func (m *Tap) XXX_DiscardUnknown() {
	xxx_messageInfo_Tap.DiscardUnknown(m)
}

var xxx_messageInfo_Tap proto.InternalMessageInfo

func (m *Tap) GetCommonConfig() *v2alpha.CommonExtensionConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *Tap) GetTransportSocket() *core.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

func init() {
	proto.RegisterType((*Tap)(nil), "envoy.config.transport_socket.tap.v2alpha.Tap")
}

func init() {
	proto.RegisterFile("envoy/config/transport_socket/tap/v2alpha/tap.proto", fileDescriptor_07cb8c0b42756e40)
}

var fileDescriptor_07cb8c0b42756e40 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0x88, 0x52, 0xa2, 0x62, 0xc9, 0xc5, 0x52, 0x3c, 0x48, 0x4f, 0x0a, 0x32, 0x0b,
	0x29, 0xe8, 0x3d, 0xc5, 0x83, 0xb7, 0xa2, 0x01, 0xc1, 0x4b, 0x99, 0xc6, 0x55, 0x17, 0xdb, 0x9d,
	0x65, 0xb3, 0x2c, 0xed, 0x2b, 0xf8, 0x5e, 0xbe, 0x94, 0x27, 0xc9, 0x4e, 0x0a, 0x4d, 0x4e, 0xbd,
	0x0d, 0x33, 0xff, 0xff, 0xfd, 0x33, 0x93, 0x4e, 0x95, 0x09, 0xb4, 0x95, 0x15, 0x99, 0x0f, 0xfd,
	0x29, 0xbd, 0x43, 0x53, 0x5b, 0x72, 0x7e, 0x51, 0x53, 0xf5, 0xad, 0xbc, 0xf4, 0x68, 0x65, 0xc8,
	0x71, 0x65, 0xbf, 0xb0, 0xa9, 0xc1, 0x3a, 0xf2, 0x94, 0xdd, 0x46, 0x13, 0xb0, 0x09, 0xfa, 0x26,
	0x68, 0x84, 0xad, 0x69, 0x7c, 0xd7, 0xe1, 0x57, 0xb4, 0x5e, 0x93, 0xe9, 0x50, 0xb9, 0xc5, 0xe0,
	0xf1, 0x15, 0xab, 0xd1, 0x6a, 0x19, 0x72, 0x59, 0x91, 0x53, 0x72, 0x89, 0xb5, 0x6a, 0xa7, 0x97,
	0x01, 0x57, 0xfa, 0x1d, 0xbd, 0x92, 0xbb, 0x82, 0x07, 0x93, 0x5f, 0x91, 0x1e, 0x95, 0x68, 0x33,
	0x95, 0x9e, 0x33, 0x6e, 0xc1, 0x79, 0x23, 0x71, 0x2d, 0x6e, 0x4e, 0xf3, 0x7b, 0xe8, 0xec, 0xdb,
	0x26, 0xee, 0x6d, 0x09, 0xb3, 0xd8, 0x7a, 0xdc, 0x78, 0x65, 0x6a, 0x4d, 0x66, 0x16, 0x85, 0xc5,
	0xe0, 0xaf, 0x38, 0xfe, 0x11, 0xc9, 0x50, 0x3c, 0x9f, 0xb1, 0x87, 0xfb, 0xd9, 0x6b, 0x3a, 0xec,
	0xdf, 0x3c, 0x4a, 0x62, 0xd2, 0xa4, 0x4d, 0x42, 0xab, 0x21, 0xe4, 0xd0, 0x1c, 0x00, 0xe5, 0x4e,
	0xfa, 0x12, 0x95, 0x7b, 0xd4, 0x0b, 0xdf, 0x1b, 0x3d, 0xa5, 0x0f, 0x9a, 0x18, 0x61, 0x1d, 0x6d,
	0xb6, 0x70, 0xf0, 0x9f, 0x8b, 0x41, 0x89, 0x76, 0xde, 0x3c, 0x63, 0x2e, 0xde, 0x92, 0x90, 0x2f,
	0x4f, 0xe2, 0x67, 0xa6, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x29, 0x27, 0x66, 0x3f, 0xe0, 0x01,
	0x00, 0x00,
}
