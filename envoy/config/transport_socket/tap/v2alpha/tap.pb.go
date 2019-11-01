// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/transport_socket/tap/v2alpha/tap.proto

package envoy_config_transport_socket_tap_v2alpha

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
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x49, 0x45, 0x29, 0x51, 0xb1, 0xec, 0xc5, 0x52, 0x3c, 0x48, 0x4f, 0x0a, 0x32, 0x81,
	0x2d, 0xe8, 0x7d, 0x4b, 0xef, 0x45, 0x17, 0x3c, 0x96, 0xe9, 0x1a, 0x35, 0xd8, 0x66, 0x42, 0x36,
	0x84, 0xf6, 0x15, 0x7c, 0x2f, 0x5f, 0xca, 0x93, 0x6c, 0x66, 0x0b, 0xdd, 0x3d, 0x79, 0x1b, 0x66,
	0xfe, 0xff, 0xfb, 0x67, 0x46, 0xce, 0xb4, 0x8d, 0xb4, 0x57, 0x15, 0xd9, 0x77, 0xf3, 0xa1, 0x82,
	0x47, 0x5b, 0x3b, 0xf2, 0x61, 0x55, 0x53, 0xf5, 0xa5, 0x83, 0x0a, 0xe8, 0x54, 0xcc, 0x71, 0xe3,
	0x3e, 0xb1, 0xa9, 0xc1, 0x79, 0x0a, 0x94, 0xdd, 0x27, 0x13, 0xb0, 0x09, 0xfa, 0x26, 0x68, 0x84,
	0xad, 0x69, 0x72, 0xc3, 0x7c, 0x74, 0x46, 0xc5, 0x5c, 0x55, 0xe4, 0xb5, 0x5a, 0x63, 0xad, 0x19,
	0x34, 0x79, 0xe8, 0xa4, 0x57, 0xb4, 0xdd, 0x92, 0xed, 0x64, 0x72, 0xab, 0x55, 0x5f, 0x47, 0xdc,
	0x98, 0x37, 0x0c, 0x5a, 0x1d, 0x0a, 0x1e, 0x4c, 0x7f, 0x84, 0x3c, 0x29, 0xd1, 0x65, 0x5a, 0x5e,
	0xb2, 0x61, 0xc5, 0xc4, 0xb1, 0xb8, 0x15, 0x77, 0xe7, 0xf9, 0x23, 0x74, 0xf6, 0x6d, 0x99, 0x47,
	0x5b, 0xc2, 0x3c, 0xb5, 0x16, 0xbb, 0xa0, 0x6d, 0x6d, 0xc8, 0xce, 0x93, 0xb0, 0x18, 0xfe, 0x16,
	0xa7, 0xdf, 0x62, 0x30, 0x12, 0xcf, 0x17, 0xec, 0xe1, 0x7e, 0xf6, 0x2a, 0x47, 0xfd, 0x9b, 0xc7,
	0x83, 0x94, 0x34, 0x6d, 0x93, 0xd0, 0x19, 0x88, 0x39, 0x34, 0xe7, 0x42, 0x79, 0x90, 0xbe, 0x24,
	0xe5, 0x11, 0xf5, 0x2a, 0xf4, 0x46, 0x0b, 0xf9, 0x64, 0x88, 0x11, 0xce, 0xd3, 0x6e, 0x0f, 0xff,
	0xfe, 0x73, 0x31, 0x2c, 0xd1, 0x2d, 0x9b, 0x67, 0x2c, 0xc5, 0xfa, 0x2c, 0x7d, 0x65, 0xf6, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x1d, 0xc8, 0x81, 0xf3, 0xdc, 0x01, 0x00, 0x00,
}
