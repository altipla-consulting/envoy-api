// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto

package envoy_config_filter_http_dynamic_forward_proxy_v2alpha

import (
	fmt "fmt"
	v2alpha "github.com/altipla-consulting/envoy-api/envoy/config/common/dynamic_forward_proxy/v2alpha"
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

// Configuration for the dynamic forward proxy HTTP filter. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
type FilterConfig struct {
	// The DNS cache configuration that the filter will attach to. Note this configuration must
	// match that of associated :ref:`dynamic forward proxy cluster configuration
	// <envoy_api_field_config.cluster.dynamic_forward_proxy.v2alpha.ClusterConfig.dns_cache_config>`.
	DnsCacheConfig       *v2alpha.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a2356b260c47da, []int{0}
}

func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetDnsCacheConfig() *v2alpha.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

// Per route Configuration for the dynamic forward proxy HTTP filter.
type PerRouteConfig struct {
	// Indicates that before DNS lookup, the host header will be swapped with
	// this value. If not set or empty, the original host header value
	// will be used and no rewrite will happen.
	HostRewrite          string   `protobuf:"bytes,1,opt,name=host_rewrite,json=hostRewrite,proto3" json:"host_rewrite,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PerRouteConfig) Reset()         { *m = PerRouteConfig{} }
func (m *PerRouteConfig) String() string { return proto.CompactTextString(m) }
func (*PerRouteConfig) ProtoMessage()    {}
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a2356b260c47da, []int{1}
}

func (m *PerRouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerRouteConfig.Unmarshal(m, b)
}
func (m *PerRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerRouteConfig.Marshal(b, m, deterministic)
}
func (m *PerRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerRouteConfig.Merge(m, src)
}
func (m *PerRouteConfig) XXX_Size() int {
	return xxx_messageInfo_PerRouteConfig.Size(m)
}
func (m *PerRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PerRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PerRouteConfig proto.InternalMessageInfo

func (m *PerRouteConfig) GetHostRewrite() string {
	if m != nil {
		return m.HostRewrite
	}
	return ""
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.config.filter.http.dynamic_forward_proxy.v2alpha.FilterConfig")
	proto.RegisterType((*PerRouteConfig)(nil), "envoy.config.filter.http.dynamic_forward_proxy.v2alpha.PerRouteConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto", fileDescriptor_85a2356b260c47da)
}

var fileDescriptor_85a2356b260c47da = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0x89, 0xa0, 0x68, 0x36, 0xc6, 0xe8, 0xc5, 0xb1, 0x93, 0xee, 0xe4, 0x29, 0x81, 0x0d,
	0xbc, 0xaf, 0x2b, 0x3b, 0x97, 0xfe, 0x81, 0x92, 0xa5, 0xa9, 0x0d, 0xb4, 0x79, 0x25, 0xcd, 0x3a,
	0xfb, 0x03, 0xbc, 0xf8, 0x93, 0x3d, 0x49, 0xf2, 0xa6, 0x50, 0x18, 0x0a, 0xbb, 0x85, 0xf7, 0x91,
	0xef, 0x7d, 0x3c, 0x9a, 0x29, 0xd3, 0xc3, 0xc0, 0x25, 0x98, 0x52, 0xbf, 0xf1, 0x52, 0xd7, 0x4e,
	0x59, 0x5e, 0x39, 0xd7, 0xf2, 0x62, 0x30, 0xa2, 0xd1, 0x32, 0x2f, 0xc1, 0x9e, 0x84, 0x2d, 0xf2,
	0xd6, 0xc2, 0xfb, 0xc0, 0xfb, 0xb5, 0xa8, 0xdb, 0x4a, 0x5c, 0xa6, 0xac, 0xb5, 0xe0, 0x20, 0x7a,
	0x0d, 0x4e, 0x86, 0x4e, 0x86, 0x4e, 0xe6, 0x9d, 0xec, 0xf2, 0xaf, 0xb3, 0x73, 0xb9, 0x1d, 0xb5,
	0x48, 0x68, 0x1a, 0x30, 0xff, 0x65, 0x98, 0x2e, 0x97, 0x42, 0x56, 0x0a, 0x57, 0x2f, 0x1f, 0x7b,
	0x51, 0xeb, 0x42, 0x38, 0xc5, 0x7f, 0x1e, 0x08, 0x56, 0x1f, 0x84, 0x4e, 0xf7, 0xa1, 0x64, 0x17,
	0xf4, 0xd1, 0x91, 0xce, 0x7f, 0x3f, 0xe7, 0xb8, 0x72, 0x41, 0x9e, 0xc8, 0xcb, 0x64, 0xbd, 0x65,
	0xa3, 0x7e, 0xec, 0xf8, 0x3b, 0x9d, 0x25, 0xa6, 0xdb, 0x79, 0x13, 0xca, 0xe3, 0xfb, 0xaf, 0xf8,
	0xf6, 0x93, 0xdc, 0xcc, 0x49, 0x36, 0x2b, 0x46, 0x64, 0xb5, 0xa1, 0xb3, 0x54, 0xd9, 0x0c, 0x8e,
	0xee, 0x3c, 0x89, 0x9e, 0xe9, 0xb4, 0x82, 0xce, 0xe5, 0x56, 0x9d, 0xac, 0x76, 0x2a, 0x44, 0x3c,
	0x64, 0x13, 0x3f, 0xcb, 0x70, 0x14, 0x1f, 0x68, 0xa2, 0x01, 0xab, 0x70, 0xed, 0x75, 0x07, 0x8e,
	0x17, 0x09, 0xe2, 0x3d, 0xd2, 0xd4, 0xc3, 0xd4, 0x9f, 0x27, 0x25, 0x87, 0xbb, 0x70, 0xa7, 0xcd,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xe7, 0x30, 0x35, 0x11, 0x02, 0x00, 0x00,
}
