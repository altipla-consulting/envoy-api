// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/redis_proxy/v2/redis_proxy.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RedisProxy struct {
	// The prefix to use when emitting :ref:`statistics <config_network_filters_redis_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Name of cluster from cluster manager. See the :ref:`configuration section
	// <arch_overview_redis_configuration>` of the architecture overview for recommendations on
	// configuring the backing cluster.
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Network settings for the connection pool to the upstream cluster.
	Settings             *RedisProxy_ConnPoolSettings `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RedisProxy) Reset()         { *m = RedisProxy{} }
func (m *RedisProxy) String() string { return proto.CompactTextString(m) }
func (*RedisProxy) ProtoMessage()    {}
func (*RedisProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_redis_proxy_af67d9f0306fe2f9, []int{0}
}
func (m *RedisProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy.Unmarshal(m, b)
}
func (m *RedisProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy.Marshal(b, m, deterministic)
}
func (dst *RedisProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy.Merge(dst, src)
}
func (m *RedisProxy) XXX_Size() int {
	return xxx_messageInfo_RedisProxy.Size(m)
}
func (m *RedisProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy proto.InternalMessageInfo

func (m *RedisProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RedisProxy) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *RedisProxy) GetSettings() *RedisProxy_ConnPoolSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

// Redis connection pool settings.
type RedisProxy_ConnPoolSettings struct {
	// Per-operation timeout in milliseconds. The timer starts when the first
	// command of a pipeline is written to the backend connection. Each response received from Redis
	// resets the timer since it signifies that the next command is being processed by the backend.
	// The only exception to this behavior is when a connection to a backend is not yet established.
	// In that case, the connect timeout on the cluster will govern the timeout until the connection
	// is ready.
	OpTimeout            *types.Duration `protobuf:"bytes,1,opt,name=op_timeout,json=opTimeout,proto3" json:"op_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RedisProxy_ConnPoolSettings) Reset()         { *m = RedisProxy_ConnPoolSettings{} }
func (m *RedisProxy_ConnPoolSettings) String() string { return proto.CompactTextString(m) }
func (*RedisProxy_ConnPoolSettings) ProtoMessage()    {}
func (*RedisProxy_ConnPoolSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_redis_proxy_af67d9f0306fe2f9, []int{0, 0}
}
func (m *RedisProxy_ConnPoolSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_ConnPoolSettings.Unmarshal(m, b)
}
func (m *RedisProxy_ConnPoolSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_ConnPoolSettings.Marshal(b, m, deterministic)
}
func (dst *RedisProxy_ConnPoolSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_ConnPoolSettings.Merge(dst, src)
}
func (m *RedisProxy_ConnPoolSettings) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_ConnPoolSettings.Size(m)
}
func (m *RedisProxy_ConnPoolSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_ConnPoolSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_ConnPoolSettings proto.InternalMessageInfo

func (m *RedisProxy_ConnPoolSettings) GetOpTimeout() *types.Duration {
	if m != nil {
		return m.OpTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*RedisProxy)(nil), "envoy.config.filter.network.redis_proxy.v2.RedisProxy")
	proto.RegisterType((*RedisProxy_ConnPoolSettings)(nil), "envoy.config.filter.network.redis_proxy.v2.RedisProxy.ConnPoolSettings")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/redis_proxy/v2/redis_proxy.proto", fileDescriptor_redis_proxy_af67d9f0306fe2f9)
}

var fileDescriptor_redis_proxy_af67d9f0306fe2f9 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x99, 0xb4, 0x6a, 0x3b, 0x05, 0x29, 0x41, 0xb0, 0x76, 0xa1, 0x45, 0x37, 0xa5, 0x8b,
	0x19, 0x88, 0x5b, 0x57, 0x51, 0xd0, 0x65, 0x89, 0xae, 0x44, 0x28, 0x69, 0x3b, 0x19, 0x06, 0xe3,
	0xbc, 0x30, 0x79, 0x89, 0xed, 0x15, 0x3c, 0x81, 0x67, 0x10, 0x4f, 0xe0, 0xca, 0x9b, 0xb8, 0xf6,
	0x16, 0x92, 0x99, 0x44, 0x4b, 0x57, 0xee, 0x5e, 0xf2, 0xbf, 0xef, 0xff, 0xe7, 0xfd, 0xf4, 0x42,
	0xe8, 0x12, 0xd6, 0x7c, 0x01, 0x3a, 0x51, 0x92, 0x27, 0x2a, 0x45, 0x61, 0xb8, 0x16, 0xf8, 0x0c,
	0xe6, 0x91, 0x1b, 0xb1, 0x54, 0xf9, 0x2c, 0x33, 0xb0, 0x5a, 0xf3, 0x32, 0xd8, 0xfc, 0x64, 0x99,
	0x01, 0x04, 0x7f, 0x62, 0x69, 0xe6, 0x68, 0xe6, 0x68, 0x56, 0xd3, 0x6c, 0x73, 0xbd, 0x0c, 0x86,
	0xc7, 0x12, 0x40, 0xa6, 0x82, 0x5b, 0x72, 0x5e, 0x24, 0x7c, 0x59, 0x98, 0x18, 0x15, 0x68, 0xe7,
	0x35, 0x3c, 0x2c, 0xe3, 0x54, 0x2d, 0x63, 0x14, 0xbc, 0x19, 0x6a, 0xe1, 0x40, 0x82, 0x04, 0x3b,
	0xf2, 0x6a, 0x72, 0x7f, 0x4f, 0xdf, 0x3d, 0x4a, 0xa3, 0x2a, 0x61, 0x5a, 0x05, 0xf8, 0x13, 0xda,
	0xcb, 0x31, 0xc6, 0x59, 0x66, 0x44, 0xa2, 0x56, 0x03, 0x32, 0x22, 0xe3, 0x6e, 0xd8, 0xfd, 0xf8,
	0xfe, 0x6c, 0xb5, 0x8d, 0x37, 0x22, 0x11, 0xad, 0xd4, 0xa9, 0x15, 0xfd, 0x33, 0xba, 0xb7, 0x48,
	0x8b, 0x1c, 0x85, 0x19, 0x78, 0xdb, 0x7b, 0x8d, 0xe2, 0x03, 0xed, 0xe4, 0x02, 0x51, 0x69, 0x99,
	0x0f, 0x5a, 0x23, 0x32, 0xee, 0x05, 0xd7, 0xec, 0xff, 0xd7, 0xb2, 0xbf, 0xa7, 0xb1, 0x4b, 0xd0,
	0x7a, 0x0a, 0x90, 0xde, 0xd6, 0x76, 0x21, 0xad, 0xe2, 0x76, 0x5e, 0x88, 0xd7, 0x27, 0xd1, 0x6f,
	0xc8, 0xf0, 0x81, 0xf6, 0xb7, 0x37, 0xfd, 0x1b, 0x4a, 0x21, 0x9b, 0xa1, 0x7a, 0x12, 0x50, 0xa0,
	0x3d, 0xaa, 0x17, 0x1c, 0x31, 0x57, 0x24, 0x6b, 0x8a, 0x64, 0x57, 0x75, 0x91, 0xe1, 0xfe, 0xeb,
	0xd7, 0x09, 0xb1, 0xe6, 0x6f, 0xc4, 0xeb, 0x90, 0xa8, 0x0b, 0xd9, 0x9d, 0x63, 0xc3, 0xf6, 0xbd,
	0x57, 0x06, 0xf3, 0x5d, 0xcb, 0x9c, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x12, 0xda, 0x31,
	0xf6, 0x01, 0x00, 0x00,
}
