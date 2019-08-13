// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/dynamic_forward_proxy/v2alpha/cluster.proto

package v2alpha

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

// Configuration for the dynamic forward proxy cluster. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
type ClusterConfig struct {
	// The DNS cache configuration that the cluster will attach to. Note this configuration must
	// match that of associated :ref:`dynamic forward proxy HTTP filter configuration
	// <envoy_api_field_config.filter.http.dynamic_forward_proxy.v2alpha.FilterConfig.dns_cache_config>`.
	DnsCacheConfig       *v2alpha.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClusterConfig) Reset()         { *m = ClusterConfig{} }
func (m *ClusterConfig) String() string { return proto.CompactTextString(m) }
func (*ClusterConfig) ProtoMessage()    {}
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeb9d327e132c11, []int{0}
}

func (m *ClusterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterConfig.Unmarshal(m, b)
}
func (m *ClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterConfig.Marshal(b, m, deterministic)
}
func (m *ClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterConfig.Merge(m, src)
}
func (m *ClusterConfig) XXX_Size() int {
	return xxx_messageInfo_ClusterConfig.Size(m)
}
func (m *ClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterConfig proto.InternalMessageInfo

func (m *ClusterConfig) GetDnsCacheConfig() *v2alpha.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterConfig)(nil), "envoy.config.cluster.dynamic_forward_proxy.v2alpha.ClusterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/dynamic_forward_proxy/v2alpha/cluster.proto", fileDescriptor_faeb9d327e132c11)
}

var fileDescriptor_faeb9d327e132c11 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d,
	0xd2, 0x4f, 0xa9, 0xcc, 0x4b, 0xcc, 0xcd, 0x4c, 0x8e, 0x4f, 0xcb, 0x2f, 0x2a, 0x4f, 0x2c, 0x4a,
	0x89, 0x2f, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x2f, 0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48, 0x84, 0xa9,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x02, 0x9b, 0xa0, 0x07, 0x31, 0x41, 0x0f, 0x26,
	0x87, 0xd5, 0x04, 0x3d, 0xa8, 0x09, 0x52, 0x8e, 0xa8, 0xb6, 0xe6, 0xe7, 0xe6, 0xe6, 0xe7, 0x11,
	0xb0, 0x34, 0x25, 0xaf, 0x38, 0x3e, 0x39, 0x31, 0x39, 0x23, 0x15, 0x62, 0xad, 0x94, 0x78, 0x59,
	0x62, 0x4e, 0x66, 0x4a, 0x62, 0x49, 0xaa, 0x3e, 0x8c, 0x01, 0x91, 0x50, 0x6a, 0x63, 0xe4, 0xe2,
	0x75, 0x86, 0xb8, 0xc2, 0x19, 0x6c, 0xbe, 0x50, 0x29, 0x97, 0x00, 0x5c, 0x77, 0x3c, 0xc4, 0x4e,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x47, 0x3d, 0x54, 0xc7, 0x83, 0x1d, 0x82, 0xdf, 0xed,
	0x7a, 0x2e, 0x79, 0xc5, 0xce, 0x20, 0x93, 0x20, 0x86, 0x3b, 0x71, 0xfc, 0x72, 0x62, 0xed, 0x62,
	0x64, 0x12, 0x60, 0x0c, 0xe2, 0x4b, 0x41, 0x95, 0xc9, 0xe3, 0x72, 0xc8, 0xcc, 0x87, 0x58, 0x00,
	0x31, 0x81, 0xf4, 0x80, 0x72, 0x92, 0x77, 0x81, 0x48, 0xbb, 0x41, 0x64, 0x03, 0x40, 0x92, 0x50,
	0xcf, 0x05, 0x80, 0x7c, 0x1b, 0xc0, 0x18, 0xc5, 0x0e, 0x55, 0x9b, 0xc4, 0x06, 0xf6, 0xbf, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x32, 0xa2, 0x23, 0xd3, 0x01, 0x00, 0x00,
}
