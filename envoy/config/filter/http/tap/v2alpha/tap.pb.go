// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/tap/v2alpha/tap.proto

package envoy_config_filter_http_tap_v2alpha

import (
	fmt "fmt"
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

// Top level configuration for the tap filter.
type Tap struct {
	// Common configuration for the HTTP tap filter.
	CommonConfig         *v2alpha.CommonExtensionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Tap) Reset()         { *m = Tap{} }
func (m *Tap) String() string { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()    {}
func (*Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee77d938a401b876, []int{0}
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

func init() {
	proto.RegisterType((*Tap)(nil), "envoy.config.filter.http.tap.v2alpha.Tap")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/tap/v2alpha/tap.proto", fileDescriptor_ee77d938a401b876)
}

var fileDescriptor_ee77d938a401b876 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0x2f, 0x49, 0x2c, 0xd0, 0x2f, 0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48,
	0x04, 0xb1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0xc0, 0xea, 0xf5, 0x20, 0xea, 0xf5,
	0x20, 0xea, 0xf5, 0x40, 0xea, 0xf5, 0x40, 0x6a, 0xa0, 0xea, 0xa5, 0x74, 0x50, 0x4c, 0x4d, 0xce,
	0xcf, 0xcd, 0xcd, 0xcf, 0x43, 0x31, 0x10, 0x22, 0x04, 0x31, 0x53, 0x4a, 0xbc, 0x2c, 0x31, 0x27,
	0x33, 0x25, 0xb1, 0x24, 0x55, 0x1f, 0xc6, 0x80, 0x48, 0x28, 0xe5, 0x70, 0x31, 0x87, 0x24, 0x16,
	0x08, 0xa5, 0x72, 0xf1, 0x42, 0xd4, 0xc7, 0x43, 0x0c, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36,
	0x32, 0xd3, 0x43, 0x71, 0x0b, 0xd4, 0x48, 0x24, 0x67, 0xe8, 0x39, 0x83, 0x85, 0x5c, 0x2b, 0x4a,
	0x52, 0xf3, 0x8a, 0x33, 0xf3, 0xf3, 0x9c, 0xc1, 0x0a, 0x9d, 0x38, 0x7e, 0x39, 0xb1, 0x76, 0x31,
	0x32, 0x09, 0x30, 0x06, 0xf1, 0x40, 0xf4, 0x40, 0xc5, 0x1d, 0xb8, 0x8c, 0x32, 0xf3, 0x21, 0x66,
	0x16, 0x14, 0xe5, 0x57, 0x54, 0xea, 0x11, 0xe3, 0x55, 0x27, 0x8e, 0x90, 0xc4, 0x82, 0x00, 0x90,
	0x6b, 0x03, 0x18, 0x93, 0xd8, 0xc0, 0xce, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x5b,
	0x3d, 0x01, 0x55, 0x01, 0x00, 0x00,
}
