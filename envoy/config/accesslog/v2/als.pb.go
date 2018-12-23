// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/accesslog/v2/als.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/altipla-consulting/envoy-api/envoy/api/v2/core"
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

// Configuration for the built-in *envoy.http_grpc_access_log*
// :ref:`AccessLog <envoy_api_msg_config.filter.accesslog.v2.AccessLog>`. This configuration will
// populate :ref:`StreamAccessLogsMessage.http_logs
// <envoy_api_field_service.accesslog.v2.StreamAccessLogsMessage.http_logs>`.
type HttpGrpcAccessLogConfig struct {
	CommonConfig *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// Additional request headers to log in :ref:`HTTPRequestProperties.request_headers
	// <envoy_api_field_data.accesslog.v2.HTTPRequestProperties.request_headers>`.
	AdditionalRequestHeadersToLog []string `protobuf:"bytes,2,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	// Additional response headers to log in :ref:`HTTPResponseProperties.response_headers
	// <envoy_api_field_data.accesslog.v2.HTTPResponseProperties.response_headers>`.
	AdditionalResponseHeadersToLog []string `protobuf:"bytes,3,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	// Additional response trailers to log in :ref:`HTTPResponseProperties.response_trailers
	// <envoy_api_field_data.accesslog.v2.HTTPResponseProperties.response_trailers>`.
	AdditionalResponseTrailersToLog []string `protobuf:"bytes,4,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *HttpGrpcAccessLogConfig) Reset()         { *m = HttpGrpcAccessLogConfig{} }
func (m *HttpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*HttpGrpcAccessLogConfig) ProtoMessage()    {}
func (*HttpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_als_58af27603741cadb, []int{0}
}
func (m *HttpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *HttpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (dst *HttpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpGrpcAccessLogConfig.Merge(dst, src)
}
func (m *HttpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_HttpGrpcAccessLogConfig.Size(m)
}
func (m *HttpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *HttpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

func (m *HttpGrpcAccessLogConfig) GetAdditionalResponseTrailersToLog() []string {
	if m != nil {
		return m.AdditionalResponseTrailersToLog
	}
	return nil
}

// Configuration for the built-in *envoy.tcp_grpc_access_log* type. This configuration will
// populate *StreamAccessLogsMessage.tcp_logs*.
// [#not-implemented-hide:]
type TcpGrpcAccessLogConfig struct {
	CommonConfig         *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TcpGrpcAccessLogConfig) Reset()         { *m = TcpGrpcAccessLogConfig{} }
func (m *TcpGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*TcpGrpcAccessLogConfig) ProtoMessage()    {}
func (*TcpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_als_58af27603741cadb, []int{1}
}
func (m *TcpGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *TcpGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (dst *TcpGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpGrpcAccessLogConfig.Merge(dst, src)
}
func (m *TcpGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_TcpGrpcAccessLogConfig.Size(m)
}
func (m *TcpGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TcpGrpcAccessLogConfig proto.InternalMessageInfo

func (m *TcpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

// Common configuration for gRPC access logs.
type CommonGrpcAccessLogConfig struct {
	// The friendly name of the access log to be returned in :ref:`StreamAccessLogsMessage.Identifier
	// <envoy_api_msg_service.accesslog.v2.StreamAccessLogsMessage.Identifier>`. This allows the
	// access log server to differentiate between different access logs coming from the same Envoy.
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// The gRPC service for the access log service.
	GrpcService          *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CommonGrpcAccessLogConfig) Reset()         { *m = CommonGrpcAccessLogConfig{} }
func (m *CommonGrpcAccessLogConfig) String() string { return proto.CompactTextString(m) }
func (*CommonGrpcAccessLogConfig) ProtoMessage()    {}
func (*CommonGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_als_58af27603741cadb, []int{2}
}
func (m *CommonGrpcAccessLogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Unmarshal(m, b)
}
func (m *CommonGrpcAccessLogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Marshal(b, m, deterministic)
}
func (dst *CommonGrpcAccessLogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonGrpcAccessLogConfig.Merge(dst, src)
}
func (m *CommonGrpcAccessLogConfig) XXX_Size() int {
	return xxx_messageInfo_CommonGrpcAccessLogConfig.Size(m)
}
func (m *CommonGrpcAccessLogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonGrpcAccessLogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonGrpcAccessLogConfig proto.InternalMessageInfo

func (m *CommonGrpcAccessLogConfig) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *CommonGrpcAccessLogConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.HttpGrpcAccessLogConfig")
	proto.RegisterType((*TcpGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.TcpGrpcAccessLogConfig")
	proto.RegisterType((*CommonGrpcAccessLogConfig)(nil), "envoy.config.accesslog.v2.CommonGrpcAccessLogConfig")
}

func init() {
	proto.RegisterFile("envoy/config/accesslog/v2/als.proto", fileDescriptor_als_58af27603741cadb)
}

var fileDescriptor_als_58af27603741cadb = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xbd, 0x4e, 0xeb, 0x30,
	0x14, 0xc7, 0x95, 0xb4, 0xf7, 0xa3, 0x6e, 0xaf, 0x74, 0x95, 0xe1, 0xf6, 0x43, 0xba, 0xbd, 0xbd,
	0x69, 0x87, 0x4e, 0x89, 0x14, 0x78, 0x01, 0xda, 0x81, 0x0a, 0x15, 0x86, 0xd0, 0x89, 0x25, 0x32,
	0x8e, 0x31, 0x96, 0x9c, 0x9c, 0x60, 0x9b, 0x48, 0x4c, 0xec, 0x4c, 0x3c, 0x0f, 0x13, 0x0f, 0xc3,
	0xc2, 0x5b, 0xa0, 0xd8, 0x05, 0x52, 0x68, 0x67, 0x36, 0x4b, 0xe7, 0x77, 0x7e, 0xff, 0xbf, 0xac,
	0x83, 0xc6, 0x34, 0x2f, 0xe1, 0x26, 0x24, 0x90, 0x5f, 0x70, 0x16, 0x62, 0x42, 0xa8, 0x52, 0x02,
	0x58, 0x58, 0x46, 0x21, 0x16, 0x2a, 0x28, 0x24, 0x68, 0xf0, 0xfa, 0x06, 0x0a, 0x2c, 0x14, 0xbc,
	0x41, 0x41, 0x19, 0x0d, 0x26, 0x76, 0x1f, 0x17, 0xbc, 0x5a, 0x21, 0x20, 0x69, 0xc8, 0x64, 0x41,
	0x12, 0x45, 0x65, 0xc9, 0x09, 0xb5, 0x82, 0x41, 0xb7, 0xc4, 0x82, 0xa7, 0x58, 0xd3, 0xf0, 0xf5,
	0x61, 0x07, 0xfe, 0x93, 0x8b, 0xba, 0x0b, 0xad, 0x8b, 0x43, 0x59, 0x90, 0x03, 0xe3, 0x5d, 0x02,
	0x9b, 0x9b, 0x1c, 0x8f, 0xa2, 0x5f, 0x04, 0xb2, 0x0c, 0xf2, 0xc4, 0x06, 0xf7, 0x9c, 0x91, 0x33,
	0x6d, 0x47, 0xfb, 0xc1, 0xce, 0x36, 0xc1, 0xdc, 0xf0, 0x5b, 0x64, 0x33, 0xf4, 0xf0, 0xfc, 0xd8,
	0xf8, 0x76, 0xe7, 0xb8, 0xbf, 0x9d, 0xb8, 0x63, 0xb5, 0xeb, 0x98, 0x05, 0xfa, 0x8f, 0xd3, 0x94,
	0x6b, 0x0e, 0x39, 0x16, 0x89, 0xa4, 0x57, 0xd7, 0x54, 0xe9, 0xe4, 0x92, 0xe2, 0x94, 0x4a, 0x95,
	0x68, 0x48, 0x04, 0xb0, 0x9e, 0x3b, 0x6a, 0x4c, 0x5b, 0xf1, 0xdf, 0x77, 0x30, 0xb6, 0xdc, 0xc2,
	0x62, 0x2b, 0x58, 0x02, 0xf3, 0x8e, 0x90, 0xbf, 0x61, 0x52, 0x05, 0xe4, 0x8a, 0x7e, 0x54, 0x35,
	0x8c, 0x6a, 0x58, 0x57, 0x59, 0x70, 0xc3, 0xb5, 0x44, 0xe3, 0x6d, 0x2e, 0x2d, 0x31, 0x17, 0x35,
	0x59, 0xd3, 0xc8, 0xfe, 0x7d, 0x96, 0xad, 0xd6, 0xa0, 0xb1, 0xf9, 0xb7, 0xe8, 0xcf, 0x8a, 0x7c,
	0xe1, 0x27, 0xfb, 0xf7, 0x0e, 0xea, 0xef, 0xdc, 0xf3, 0x26, 0xe8, 0xa7, 0x00, 0x96, 0xe4, 0x38,
	0xa3, 0x26, 0xbf, 0x35, 0x6b, 0x55, 0xa6, 0xa6, 0x74, 0x47, 0x4e, 0xfc, 0x43, 0x00, 0x3b, 0xc1,
	0x19, 0xf5, 0x8e, 0x51, 0xa7, 0x7e, 0x5a, 0x3d, 0xd7, 0x34, 0x1d, 0xae, 0x9b, 0xe2, 0x82, 0x57,
	0xe5, 0xaa, 0x0b, 0x0c, 0xaa, 0x8c, 0x53, 0x4b, 0x6d, 0x74, 0x6a, 0xb3, 0xda, 0xa0, 0x79, 0xe6,
	0x96, 0xd1, 0xf9, 0x77, 0x73, 0x87, 0x7b, 0x2f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xe9, 0x66,
	0xc4, 0x08, 0x03, 0x00, 0x00,
}
