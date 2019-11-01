// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/metrics/v3alpha/metrics_service.proto

package envoy_config_metrics_v3alpha

import (
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v3alpha/core"
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

// Metrics Service is configured as a built-in *envoy.metrics_service* :ref:`StatsSink
// <envoy_api_msg_config.metrics.v3alpha.StatsSink>`. This opaque configuration will be used to
// create Metrics Service.
type MetricsServiceConfig struct {
	// The upstream gRPC cluster that hosts the metrics service.
	GrpcService          *core.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetricsServiceConfig) Reset()         { *m = MetricsServiceConfig{} }
func (m *MetricsServiceConfig) String() string { return proto.CompactTextString(m) }
func (*MetricsServiceConfig) ProtoMessage()    {}
func (*MetricsServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_b66f37e11b5cdc04, []int{0}
}

func (m *MetricsServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsServiceConfig.Unmarshal(m, b)
}
func (m *MetricsServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsServiceConfig.Marshal(b, m, deterministic)
}
func (m *MetricsServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsServiceConfig.Merge(m, src)
}
func (m *MetricsServiceConfig) XXX_Size() int {
	return xxx_messageInfo_MetricsServiceConfig.Size(m)
}
func (m *MetricsServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsServiceConfig proto.InternalMessageInfo

func (m *MetricsServiceConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricsServiceConfig)(nil), "envoy.config.metrics.v3alpha.MetricsServiceConfig")
}

func init() {
	proto.RegisterFile("envoy/config/metrics/v3alpha/metrics_service.proto", fileDescriptor_b66f37e11b5cdc04)
}

var fileDescriptor_b66f37e11b5cdc04 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0xcf, 0x4d, 0x2d, 0x29, 0xca, 0x4c, 0x2e,
	0xd6, 0x2f, 0x33, 0x4e, 0xcc, 0x29, 0xc8, 0x48, 0x84, 0xf1, 0xe3, 0x8b, 0x53, 0x8b, 0xca, 0x32,
	0x93, 0x53, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x64, 0xc0, 0x7a, 0xf4, 0x20, 0x7a, 0xf4,
	0xa0, 0x6a, 0xf4, 0xa0, 0x7a, 0xa4, 0x34, 0x21, 0x26, 0x26, 0x16, 0x64, 0xc2, 0x8d, 0x49, 0xce,
	0x2f, 0x4a, 0xd5, 0x4f, 0x2f, 0x2a, 0x48, 0x46, 0x35, 0x48, 0x4a, 0xbc, 0x2c, 0x31, 0x27, 0x33,
	0x25, 0xb1, 0x24, 0x55, 0x1f, 0xc6, 0x80, 0x48, 0x28, 0x65, 0x70, 0x89, 0xf8, 0x42, 0x8c, 0x0d,
	0x86, 0x68, 0x70, 0x06, 0xdb, 0x25, 0x14, 0xc0, 0xc5, 0x83, 0x6c, 0x8c, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0xb7, 0x91, 0xb2, 0x1e, 0xc4, 0x41, 0x89, 0x05, 0x99, 0x30, 0x57, 0xe8, 0x81, 0xac, 0xd4,
	0x73, 0x2f, 0x2a, 0x48, 0x86, 0x1a, 0xe0, 0xc4, 0xf1, 0xcb, 0x89, 0xb5, 0x8b, 0x91, 0x49, 0x80,
	0x31, 0x88, 0x3b, 0x1d, 0x49, 0xd8, 0x99, 0x4b, 0x2b, 0x33, 0x1f, 0xa2, 0xbf, 0xa0, 0x28, 0xbf,
	0xa2, 0x52, 0x0f, 0x9f, 0xdf, 0x9c, 0x84, 0x51, 0x5d, 0x15, 0x00, 0x72, 0x6c, 0x00, 0x63, 0x12,
	0x1b, 0xd8, 0xd5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xee, 0x4c, 0x58, 0x4d, 0x01,
	0x00, 0x00,
}
