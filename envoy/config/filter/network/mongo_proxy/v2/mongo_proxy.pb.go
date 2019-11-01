// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/mongo_proxy/v2/mongo_proxy.proto

package envoy_config_filter_network_mongo_proxy_v2

import (
	fmt "fmt"
	v2 "github.com/altipla-consulting/envoy-api/envoy/config/filter/fault/v2"
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

type MongoProxy struct {
	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_mongo_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The optional path to use for writing Mongo access logs. If not access log
	// path is specified no access logs will be written. Note that access log is
	// also gated :ref:`runtime <config_network_filters_mongo_proxy_runtime>`.
	AccessLog string `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// Inject a fixed delay before proxying a Mongo operation. Delays are
	// applied to the following MongoDB operations: Query, Insert, GetMore,
	// and KillCursors. Once an active delay is in progress, all incoming
	// data up until the timer event fires will be a part of the delay.
	Delay *v2.FaultDelay `protobuf:"bytes,3,opt,name=delay,proto3" json:"delay,omitempty"`
	// Flag to specify whether :ref:`dynamic metadata
	// <config_network_filters_mongo_proxy_dynamic_metadata>` should be emitted. Defaults to false.
	EmitDynamicMetadata  bool     `protobuf:"varint,4,opt,name=emit_dynamic_metadata,json=emitDynamicMetadata,proto3" json:"emit_dynamic_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MongoProxy) Reset()         { *m = MongoProxy{} }
func (m *MongoProxy) String() string { return proto.CompactTextString(m) }
func (*MongoProxy) ProtoMessage()    {}
func (*MongoProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d590dd12f767c61, []int{0}
}

func (m *MongoProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongoProxy.Unmarshal(m, b)
}
func (m *MongoProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongoProxy.Marshal(b, m, deterministic)
}
func (m *MongoProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongoProxy.Merge(m, src)
}
func (m *MongoProxy) XXX_Size() int {
	return xxx_messageInfo_MongoProxy.Size(m)
}
func (m *MongoProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_MongoProxy.DiscardUnknown(m)
}

var xxx_messageInfo_MongoProxy proto.InternalMessageInfo

func (m *MongoProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *MongoProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func (m *MongoProxy) GetDelay() *v2.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *MongoProxy) GetEmitDynamicMetadata() bool {
	if m != nil {
		return m.EmitDynamicMetadata
	}
	return false
}

func init() {
	proto.RegisterType((*MongoProxy)(nil), "envoy.config.filter.network.mongo_proxy.v2.MongoProxy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/mongo_proxy/v2/mongo_proxy.proto", fileDescriptor_4d590dd12f767c61)
}

var fileDescriptor_4d590dd12f767c61 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0xc9, 0x38, 0xfe, 0x4c, 0x66, 0x21, 0x54, 0xc4, 0x32, 0x20, 0x14, 0x57, 0xc1, 0x45,
	0x02, 0x75, 0xe3, 0x42, 0x5c, 0x94, 0xc1, 0x95, 0x85, 0xd2, 0x17, 0x28, 0xb1, 0x4d, 0x4b, 0xb0,
	0xed, 0x2d, 0x69, 0x8c, 0xd3, 0x47, 0xf4, 0x95, 0x5c, 0x49, 0x7a, 0x47, 0x9c, 0xc5, 0x2c, 0xdc,
	0xdd, 0xe4, 0x3b, 0xe7, 0xdc, 0xe4, 0xd0, 0x27, 0xd5, 0x3b, 0x98, 0x44, 0x09, 0x7d, 0xad, 0x1b,
	0x51, 0xeb, 0xd6, 0x2a, 0x23, 0x7a, 0x65, 0x3f, 0xc1, 0xbc, 0x8b, 0x0e, 0xfa, 0x06, 0x8a, 0xc1,
	0xc0, 0x6e, 0x12, 0x2e, 0x3e, 0x3c, 0xf2, 0xc1, 0x80, 0x85, 0xe0, 0x7e, 0x76, 0x73, 0x74, 0x73,
	0x74, 0xf3, 0xbd, 0x9b, 0x1f, 0xca, 0x5d, 0xbc, 0x61, 0xc7, 0x36, 0xd5, 0xf2, 0xa3, 0xb5, 0x3e,
	0x7b, 0x1e, 0x30, 0x75, 0x73, 0xe3, 0x64, 0xab, 0x2b, 0x69, 0x95, 0xf8, 0x1d, 0x10, 0xdc, 0x7d,
	0x11, 0x4a, 0x53, 0x9f, 0x9a, 0xf9, 0xd0, 0x80, 0xd1, 0xf5, 0x68, 0xa5, 0x2d, 0x06, 0xa3, 0x6a,
	0xbd, 0x0b, 0x49, 0x44, 0xd8, 0x2a, 0x39, 0xff, 0x4e, 0x96, 0x66, 0x11, 0x91, 0x9c, 0x7a, 0x96,
	0xcd, 0x28, 0xb8, 0xa5, 0x54, 0x96, 0xa5, 0x1a, 0xc7, 0xa2, 0x85, 0x26, 0x5c, 0x78, 0x61, 0xbe,
	0xc2, 0x9b, 0x57, 0x68, 0x82, 0x67, 0x7a, 0x5a, 0xa9, 0x56, 0x4e, 0xe1, 0x49, 0x44, 0xd8, 0x3a,
	0x66, 0xfc, 0xd8, 0xb7, 0xf0, 0x85, 0x2e, 0xe6, 0x2f, 0x7e, 0xd8, 0x7a, 0x7d, 0x8e, 0xb6, 0x20,
	0xa6, 0xd7, 0xaa, 0xd3, 0xb6, 0xa8, 0xa6, 0x5e, 0x76, 0xba, 0x2c, 0x3a, 0x65, 0x65, 0x25, 0xad,
	0x0c, 0x97, 0x11, 0x61, 0x17, 0xf9, 0x95, 0x87, 0x5b, 0x64, 0xe9, 0x1e, 0x25, 0x29, 0x7d, 0xd4,
	0x80, 0x8b, 0xb0, 0xa2, 0xff, 0x57, 0x99, 0x5c, 0xfe, 0x95, 0x90, 0xf9, 0x62, 0x32, 0xf2, 0x76,
	0x36, 0x37, 0xf4, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x52, 0x62, 0x06, 0xd0, 0x01, 0x00,
	0x00,
}
