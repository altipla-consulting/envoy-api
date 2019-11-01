// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/mongo_proxy/v3alpha/mongo_proxy.proto

package envoy_config_filter_network_mongo_proxy_v3alpha

import (
	fmt "fmt"
	v3alpha "github.com/altipla-consulting/envoy-api/envoy/config/filter/fault/v3alpha"
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
	Delay *v3alpha.FaultDelay `protobuf:"bytes,3,opt,name=delay,proto3" json:"delay,omitempty"`
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
	return fileDescriptor_2b8f056578a6890a, []int{0}
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

func (m *MongoProxy) GetDelay() *v3alpha.FaultDelay {
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
	proto.RegisterType((*MongoProxy)(nil), "envoy.config.filter.network.mongo_proxy.v3alpha.MongoProxy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/mongo_proxy/v3alpha/mongo_proxy.proto", fileDescriptor_2b8f056578a6890a)
}

var fileDescriptor_2b8f056578a6890a = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0xe5, 0x52, 0xfe, 0xd4, 0x1d, 0x90, 0x82, 0x10, 0x51, 0x25, 0xa4, 0x88, 0x29, 0x4b,
	0x6d, 0xa9, 0x9d, 0x19, 0x08, 0x15, 0x13, 0x95, 0xa2, 0xbc, 0x40, 0x74, 0x24, 0x4e, 0xb0, 0x70,
	0x7c, 0x91, 0x6b, 0x42, 0xf3, 0x92, 0x3c, 0x10, 0x13, 0x72, 0x1c, 0xa0, 0x43, 0x17, 0x36, 0xdf,
	0x7d, 0x77, 0xbf, 0xef, 0xfc, 0xd1, 0x07, 0xa1, 0x3b, 0xec, 0x79, 0x81, 0xba, 0x92, 0x35, 0xaf,
	0xa4, 0xb2, 0xc2, 0x70, 0x2d, 0xec, 0x07, 0x9a, 0x37, 0xde, 0xa0, 0xae, 0x31, 0x6f, 0x0d, 0xee,
	0x7b, 0xde, 0xad, 0x41, 0xb5, 0xaf, 0x70, 0xd8, 0x63, 0xad, 0x41, 0x8b, 0x01, 0x1f, 0x10, 0xcc,
	0x23, 0x98, 0x47, 0xb0, 0x11, 0xc1, 0x0e, 0xc7, 0x47, 0xc4, 0x62, 0x79, 0xcc, 0xb3, 0x82, 0x77,
	0x65, 0x7f, 0x5d, 0x86, 0xca, 0xf3, 0x17, 0x37, 0x1d, 0x28, 0x59, 0x82, 0x15, 0xfc, 0xe7, 0xe1,
	0x85, 0xbb, 0x4f, 0x42, 0xe9, 0xd6, 0xf1, 0x53, 0x87, 0x0f, 0x62, 0x3a, 0xdf, 0x59, 0xb0, 0x79,
	0x6b, 0x44, 0x25, 0xf7, 0x21, 0x89, 0x48, 0x3c, 0x4b, 0xce, 0xbf, 0x92, 0xa9, 0x99, 0x44, 0x24,
	0xa3, 0x4e, 0x4b, 0x07, 0x29, 0xb8, 0xa5, 0x14, 0x8a, 0x42, 0xec, 0x76, 0xb9, 0xc2, 0x3a, 0x9c,
	0xb8, 0xc1, 0x6c, 0xe6, 0x3b, 0xcf, 0x58, 0x07, 0x8f, 0xf4, 0xb4, 0x14, 0x0a, 0xfa, 0xf0, 0x24,
	0x22, 0xf1, 0x7c, 0xb5, 0x64, 0xc7, 0x3e, 0xe8, 0x2f, 0x1c, 0xef, 0x65, 0x4f, 0xae, 0xda, 0xb8,
	0xa5, 0xcc, 0xef, 0x06, 0x2b, 0x7a, 0x2d, 0x1a, 0x69, 0xf3, 0xb2, 0xd7, 0xd0, 0xc8, 0x22, 0x6f,
	0x84, 0x85, 0x12, 0x2c, 0x84, 0xd3, 0x88, 0xc4, 0x17, 0xd9, 0x95, 0x13, 0x37, 0x5e, 0xdb, 0x8e,
	0x52, 0x92, 0xd1, 0x7b, 0x89, 0xde, 0xcd, 0x27, 0xf6, 0xcf, 0x64, 0x93, 0xcb, 0xbf, 0x38, 0x52,
	0x17, 0x51, 0x4a, 0x5e, 0xce, 0x86, 0xac, 0xd6, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xe8,
	0xe1, 0xaa, 0xe9, 0x01, 0x00, 0x00,
}