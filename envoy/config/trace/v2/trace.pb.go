// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v2/trace.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/altipla-consulting/envoy-api/envoy/api/v2/core"
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

// The tracing configuration specifies global
// settings for the HTTP tracer used by Envoy. The configuration is defined by
// the :ref:`Bootstrap <envoy_api_msg_config.bootstrap.v2.Bootstrap>` :ref:`tracing
// <envoy_api_field_config.bootstrap.v2.Bootstrap.tracing>` field. Envoy may support other tracers
// in the future, but right now the HTTP tracer is the only one supported.
type Tracing struct {
	// Provides configuration for the HTTP tracer.
	Http                 *Tracing_Http `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Tracing) Reset()         { *m = Tracing{} }
func (m *Tracing) String() string { return proto.CompactTextString(m) }
func (*Tracing) ProtoMessage()    {}
func (*Tracing) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_940b649faf0ffb9f, []int{0}
}
func (m *Tracing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing.Unmarshal(m, b)
}
func (m *Tracing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing.Marshal(b, m, deterministic)
}
func (dst *Tracing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing.Merge(dst, src)
}
func (m *Tracing) XXX_Size() int {
	return xxx_messageInfo_Tracing.Size(m)
}
func (m *Tracing) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing proto.InternalMessageInfo

func (m *Tracing) GetHttp() *Tracing_Http {
	if m != nil {
		return m.Http
	}
	return nil
}

type Tracing_Http struct {
	// The name of the HTTP trace driver to instantiate. The name must match a
	// supported HTTP trace driver. Built-in trace drivers:
	//
	// - *envoy.lightstep*
	// - *envoy.zipkin*
	// - *envoy.dynamic.ot*
	// - *envoy.tracers.datadog*
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Trace driver specific configuration which depends on the driver being instantiated.
	// See the trace drivers for examples:
	//
	// - :ref:`LightstepConfig <envoy_api_msg_config.trace.v2.LightstepConfig>`
	// - :ref:`ZipkinConfig <envoy_api_msg_config.trace.v2.ZipkinConfig>`
	// - :ref:`DynamicOtConfig <envoy_api_msg_config.trace.v2.DynamicOtConfig>`
	// - :ref:`DatadogConfig <envoy_api_msg_config.trace.v2.DatadogConfig>`
	//
	// Types that are valid to be assigned to ConfigType:
	//	*Tracing_Http_Config
	//	*Tracing_Http_TypedConfig
	ConfigType           isTracing_Http_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Tracing_Http) Reset()         { *m = Tracing_Http{} }
func (m *Tracing_Http) String() string { return proto.CompactTextString(m) }
func (*Tracing_Http) ProtoMessage()    {}
func (*Tracing_Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_940b649faf0ffb9f, []int{0, 0}
}
func (m *Tracing_Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing_Http.Unmarshal(m, b)
}
func (m *Tracing_Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing_Http.Marshal(b, m, deterministic)
}
func (dst *Tracing_Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing_Http.Merge(dst, src)
}
func (m *Tracing_Http) XXX_Size() int {
	return xxx_messageInfo_Tracing_Http.Size(m)
}
func (m *Tracing_Http) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracing_Http.DiscardUnknown(m)
}

var xxx_messageInfo_Tracing_Http proto.InternalMessageInfo

func (m *Tracing_Http) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTracing_Http_ConfigType interface {
	isTracing_Http_ConfigType()
}

type Tracing_Http_Config struct {
	Config *types.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type Tracing_Http_TypedConfig struct {
	TypedConfig *types.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*Tracing_Http_Config) isTracing_Http_ConfigType() {}

func (*Tracing_Http_TypedConfig) isTracing_Http_ConfigType() {}

func (m *Tracing_Http) GetConfigType() isTracing_Http_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *Tracing_Http) GetConfig() *types.Struct {
	if x, ok := m.GetConfigType().(*Tracing_Http_Config); ok {
		return x.Config
	}
	return nil
}

func (m *Tracing_Http) GetTypedConfig() *types.Any {
	if x, ok := m.GetConfigType().(*Tracing_Http_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Tracing_Http) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Tracing_Http_OneofMarshaler, _Tracing_Http_OneofUnmarshaler, _Tracing_Http_OneofSizer, []interface{}{
		(*Tracing_Http_Config)(nil),
		(*Tracing_Http_TypedConfig)(nil),
	}
}

func _Tracing_Http_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Tracing_Http)
	// config_type
	switch x := m.ConfigType.(type) {
	case *Tracing_Http_Config:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *Tracing_Http_TypedConfig:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TypedConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Tracing_Http.ConfigType has unexpected type %T", x)
	}
	return nil
}

func _Tracing_Http_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Tracing_Http)
	switch tag {
	case 2: // config_type.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types.Struct)
		err := b.DecodeMessage(msg)
		m.ConfigType = &Tracing_Http_Config{msg}
		return true, err
	case 3: // config_type.typed_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types.Any)
		err := b.DecodeMessage(msg)
		m.ConfigType = &Tracing_Http_TypedConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Tracing_Http_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Tracing_Http)
	// config_type
	switch x := m.ConfigType.(type) {
	case *Tracing_Http_Config:
		s := proto.Size(x.Config)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tracing_Http_TypedConfig:
		s := proto.Size(x.TypedConfig)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Configuration for the LightStep tracer.
type LightstepConfig struct {
	// The cluster manager cluster that hosts the LightStep collectors.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	// File containing the access token to the `LightStep
	// <http://lightstep.com/>`_ API.
	AccessTokenFile      string   `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile,proto3" json:"access_token_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LightstepConfig) Reset()         { *m = LightstepConfig{} }
func (m *LightstepConfig) String() string { return proto.CompactTextString(m) }
func (*LightstepConfig) ProtoMessage()    {}
func (*LightstepConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_940b649faf0ffb9f, []int{1}
}
func (m *LightstepConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightstepConfig.Unmarshal(m, b)
}
func (m *LightstepConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightstepConfig.Marshal(b, m, deterministic)
}
func (dst *LightstepConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightstepConfig.Merge(dst, src)
}
func (m *LightstepConfig) XXX_Size() int {
	return xxx_messageInfo_LightstepConfig.Size(m)
}
func (m *LightstepConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LightstepConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LightstepConfig proto.InternalMessageInfo

func (m *LightstepConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *LightstepConfig) GetAccessTokenFile() string {
	if m != nil {
		return m.AccessTokenFile
	}
	return ""
}

type ZipkinConfig struct {
	// The cluster manager cluster that hosts the Zipkin collectors. Note that the
	// Zipkin cluster must be defined in the :ref:`Bootstrap static cluster
	// resources <envoy_api_field_config.bootstrap.v2.Bootstrap.StaticResources.clusters>`.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	// The API endpoint of the Zipkin service where the spans will be sent. When
	// using a standard Zipkin installation, the API endpoint is typically
	// /api/v1/spans, which is the default value.
	CollectorEndpoint string `protobuf:"bytes,2,opt,name=collector_endpoint,json=collectorEndpoint,proto3" json:"collector_endpoint,omitempty"`
	// Determines whether a 128bit trace id will be used when creating a new
	// trace instance. The default value is false, which will result in a 64 bit trace id being used.
	TraceId_128Bit bool `protobuf:"varint,3,opt,name=trace_id_128bit,json=traceId128bit,proto3" json:"trace_id_128bit,omitempty"`
	// Determines whether client and server spans will shared the same span id.
	// The default value is true.
	SharedSpanContext    *types.BoolValue `protobuf:"bytes,4,opt,name=shared_span_context,json=sharedSpanContext,proto3" json:"shared_span_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ZipkinConfig) Reset()         { *m = ZipkinConfig{} }
func (m *ZipkinConfig) String() string { return proto.CompactTextString(m) }
func (*ZipkinConfig) ProtoMessage()    {}
func (*ZipkinConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_940b649faf0ffb9f, []int{2}
}
func (m *ZipkinConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZipkinConfig.Unmarshal(m, b)
}
func (m *ZipkinConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZipkinConfig.Marshal(b, m, deterministic)
}
func (dst *ZipkinConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZipkinConfig.Merge(dst, src)
}
func (m *ZipkinConfig) XXX_Size() int {
	return xxx_messageInfo_ZipkinConfig.Size(m)
}
func (m *ZipkinConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ZipkinConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ZipkinConfig proto.InternalMessageInfo

func (m *ZipkinConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *ZipkinConfig) GetCollectorEndpoint() string {
	if m != nil {
		return m.CollectorEndpoint
	}
	return ""
}

func (m *ZipkinConfig) GetTraceId_128Bit() bool {
	if m != nil {
		return m.TraceId_128Bit
	}
	return false
}

func (m *ZipkinConfig) GetSharedSpanContext() *types.BoolValue {
	if m != nil {
		return m.SharedSpanContext
	}
	return nil
}

// DynamicOtConfig is used to dynamically load a tracer from a shared library
// that implements the `OpenTracing dynamic loading API
// <https://github.com/opentracing/opentracing-cpp>`_.
type DynamicOtConfig struct {
	// Dynamic library implementing the `OpenTracing API
	// <https://github.com/opentracing/opentracing-cpp>`_.
	Library string `protobuf:"bytes,1,opt,name=library,proto3" json:"library,omitempty"`
	// The configuration to use when creating a tracer from the given dynamic
	// library.
	Config               *types.Struct `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DynamicOtConfig) Reset()         { *m = DynamicOtConfig{} }
func (m *DynamicOtConfig) String() string { return proto.CompactTextString(m) }
func (*DynamicOtConfig) ProtoMessage()    {}
func (*DynamicOtConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_940b649faf0ffb9f, []int{3}
}
func (m *DynamicOtConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicOtConfig.Unmarshal(m, b)
}
func (m *DynamicOtConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicOtConfig.Marshal(b, m, deterministic)
}
func (dst *DynamicOtConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicOtConfig.Merge(dst, src)
}
func (m *DynamicOtConfig) XXX_Size() int {
	return xxx_messageInfo_DynamicOtConfig.Size(m)
}
func (m *DynamicOtConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicOtConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicOtConfig proto.InternalMessageInfo

func (m *DynamicOtConfig) GetLibrary() string {
	if m != nil {
		return m.Library
	}
	return ""
}

func (m *DynamicOtConfig) GetConfig() *types.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// Configuration for the Datadog tracer.
type DatadogConfig struct {
	// The cluster to use for submitting traces to the Datadog agent.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	// The name used for the service when traces are generated by envoy.
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatadogConfig) Reset()         { *m = DatadogConfig{} }
func (m *DatadogConfig) String() string { return proto.CompactTextString(m) }
func (*DatadogConfig) ProtoMessage()    {}
func (*DatadogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_940b649faf0ffb9f, []int{4}
}
func (m *DatadogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatadogConfig.Unmarshal(m, b)
}
func (m *DatadogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatadogConfig.Marshal(b, m, deterministic)
}
func (dst *DatadogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatadogConfig.Merge(dst, src)
}
func (m *DatadogConfig) XXX_Size() int {
	return xxx_messageInfo_DatadogConfig.Size(m)
}
func (m *DatadogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DatadogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DatadogConfig proto.InternalMessageInfo

func (m *DatadogConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *DatadogConfig) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

// Configuration structure.
type TraceServiceConfig struct {
	// The upstream gRPC cluster that hosts the metrics service.
	GrpcService          *core.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TraceServiceConfig) Reset()         { *m = TraceServiceConfig{} }
func (m *TraceServiceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceServiceConfig) ProtoMessage()    {}
func (*TraceServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_940b649faf0ffb9f, []int{5}
}
func (m *TraceServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceServiceConfig.Unmarshal(m, b)
}
func (m *TraceServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceServiceConfig.Marshal(b, m, deterministic)
}
func (dst *TraceServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceServiceConfig.Merge(dst, src)
}
func (m *TraceServiceConfig) XXX_Size() int {
	return xxx_messageInfo_TraceServiceConfig.Size(m)
}
func (m *TraceServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceServiceConfig proto.InternalMessageInfo

func (m *TraceServiceConfig) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*Tracing)(nil), "envoy.config.trace.v2.Tracing")
	proto.RegisterType((*Tracing_Http)(nil), "envoy.config.trace.v2.Tracing.Http")
	proto.RegisterType((*LightstepConfig)(nil), "envoy.config.trace.v2.LightstepConfig")
	proto.RegisterType((*ZipkinConfig)(nil), "envoy.config.trace.v2.ZipkinConfig")
	proto.RegisterType((*DynamicOtConfig)(nil), "envoy.config.trace.v2.DynamicOtConfig")
	proto.RegisterType((*DatadogConfig)(nil), "envoy.config.trace.v2.DatadogConfig")
	proto.RegisterType((*TraceServiceConfig)(nil), "envoy.config.trace.v2.TraceServiceConfig")
}

func init() {
	proto.RegisterFile("envoy/config/trace/v2/trace.proto", fileDescriptor_trace_940b649faf0ffb9f)
}

var fileDescriptor_trace_940b649faf0ffb9f = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x6e, 0xb6, 0xa1, 0xa5, 0xde, 0xad, 0xb6, 0x35, 0xa0, 0x96, 0x15, 0x54, 0x25, 0x45, 0x88,
	0x03, 0x4a, 0xd4, 0x45, 0x40, 0x39, 0xb2, 0x2d, 0xb0, 0x20, 0x7e, 0xa4, 0x6c, 0xc5, 0xa1, 0x97,
	0xc8, 0xeb, 0xcc, 0x66, 0xad, 0x66, 0x6d, 0xcb, 0xf1, 0x06, 0x72, 0xe3, 0xcc, 0x63, 0xf0, 0x18,
	0x9c, 0x78, 0x1b, 0xc4, 0x9d, 0x07, 0x40, 0xb1, 0xbd, 0xb4, 0x34, 0x1c, 0x90, 0x7a, 0xb3, 0xe7,
	0xfb, 0x3e, 0xcf, 0x7c, 0x93, 0x99, 0xa0, 0x3b, 0xc0, 0x4b, 0x51, 0x45, 0x54, 0xf0, 0x09, 0xcb,
	0x22, 0xad, 0x08, 0x85, 0xa8, 0xec, 0xdb, 0x43, 0x28, 0x95, 0xd0, 0x02, 0xdf, 0x30, 0x94, 0xd0,
	0x52, 0x42, 0x8b, 0x94, 0xfd, 0xde, 0x5d, 0xab, 0x24, 0x92, 0xd5, 0x02, 0x2a, 0x14, 0x44, 0x99,
	0x92, 0x34, 0x29, 0x40, 0x95, 0x6c, 0x21, 0xee, 0xdd, 0xcc, 0x84, 0xc8, 0x72, 0x88, 0xcc, 0x6d,
	0x3c, 0x9f, 0x44, 0x84, 0x57, 0x0e, 0xba, 0x75, 0x11, 0x2a, 0xb4, 0x9a, 0x53, 0xed, 0xd0, 0x9d,
	0x8b, 0xe8, 0x47, 0x45, 0xa4, 0x04, 0x55, 0x38, 0x7c, 0xab, 0x24, 0x39, 0x4b, 0x89, 0x86, 0x68,
	0x71, 0xb0, 0x40, 0xf0, 0xc3, 0x43, 0xab, 0xc7, 0x8a, 0x50, 0xc6, 0x33, 0xfc, 0x04, 0xf9, 0x53,
	0xad, 0xe5, 0xb6, 0xb7, 0xeb, 0xdd, 0x6f, 0xf7, 0xf7, 0xc2, 0x7f, 0x3a, 0x09, 0x1d, 0x3b, 0x1c,
	0x6a, 0x2d, 0x63, 0x23, 0xe8, 0x7d, 0xf5, 0x90, 0x5f, 0x5f, 0xf1, 0x6d, 0xe4, 0x73, 0x32, 0x03,
	0xf3, 0xc2, 0xda, 0x60, 0xed, 0xdb, 0xcf, 0xef, 0xcb, 0xbe, 0x6a, 0xed, 0x7a, 0xb1, 0x09, 0xe3,
	0x7d, 0xb4, 0x62, 0x5f, 0xdb, 0x6e, 0x99, 0x14, 0x5b, 0xa1, 0x2d, 0x3b, 0x5c, 0x94, 0x1d, 0x8e,
	0x8c, 0xa9, 0xe1, 0x52, 0xec, 0x88, 0xf8, 0x29, 0xea, 0xe8, 0x4a, 0x42, 0x9a, 0x38, 0xe1, 0xb2,
	0x11, 0x5e, 0x6f, 0x08, 0x9f, 0xf1, 0x6a, 0xb8, 0x14, 0xb7, 0x0d, 0xf7, 0xd0, 0x50, 0x07, 0xeb,
	0xa8, 0x6d, 0x45, 0x49, 0x1d, 0x0d, 0x3e, 0x7b, 0xa8, 0xfb, 0x86, 0x65, 0x53, 0x5d, 0x68, 0x90,
	0x96, 0x82, 0x1f, 0xa3, 0x4d, 0x2a, 0xf2, 0x1c, 0xa8, 0x16, 0x2a, 0xa1, 0xf9, 0xbc, 0xd0, 0xa0,
	0x9a, 0xc5, 0x6f, 0xfc, 0xe1, 0x1c, 0x5a, 0x0a, 0x7e, 0x84, 0x36, 0x09, 0xa5, 0x50, 0x14, 0x89,
	0x16, 0xa7, 0xc0, 0x93, 0x09, 0xcb, 0xc1, 0x78, 0xfa, 0x4b, 0xd7, 0xb5, 0x9c, 0xe3, 0x9a, 0xf2,
	0x82, 0xe5, 0x10, 0xfc, 0xf2, 0x50, 0xe7, 0x84, 0xc9, 0x53, 0xc6, 0x2f, 0x99, 0xff, 0x00, 0xe1,
	0x33, 0x1d, 0xf0, 0x54, 0x0a, 0xc6, 0x75, 0xb3, 0x80, 0xb3, 0xc7, 0x9f, 0x3b, 0x0e, 0xbe, 0x87,
	0xba, 0xe6, 0x4b, 0x26, 0x2c, 0x4d, 0xf6, 0xfb, 0x07, 0x63, 0xa6, 0x4d, 0x4b, 0xaf, 0xc6, 0xeb,
	0x26, 0xfc, 0x2a, 0xb5, 0x41, 0xfc, 0x1a, 0x5d, 0x2b, 0xa6, 0x44, 0x41, 0x9a, 0x14, 0x92, 0xf0,
	0xba, 0xfb, 0x1a, 0x3e, 0xe9, 0x6d, 0xdf, 0xb4, 0xbf, 0xd7, 0x68, 0xff, 0x40, 0x88, 0xfc, 0x03,
	0xc9, 0xe7, 0x10, 0x6f, 0x5a, 0xd9, 0x48, 0x92, 0xda, 0x64, 0x2d, 0x0a, 0x32, 0xd4, 0x3d, 0xaa,
	0x38, 0x99, 0x31, 0xfa, 0x5e, 0x3b, 0xe3, 0x7b, 0x68, 0x35, 0x67, 0x63, 0x45, 0x54, 0xd5, 0xb4,
	0xbb, 0x40, 0x70, 0xf4, 0x9f, 0xe3, 0xb2, 0x18, 0x96, 0x60, 0x8e, 0xd6, 0x8f, 0x88, 0x26, 0xa9,
	0xc8, 0x2e, 0xd9, 0xdf, 0x07, 0xa8, 0xe3, 0x16, 0x33, 0x31, 0xf3, 0xdc, 0xe8, 0x6c, 0xdb, 0xc1,
	0xef, 0xc8, 0x0c, 0x02, 0x8a, 0x70, 0xbd, 0x14, 0x30, 0xb2, 0x31, 0x97, 0xfb, 0x2d, 0xea, 0x9c,
	0xdf, 0x70, 0xb7, 0x55, 0x3b, 0x6e, 0xab, 0x88, 0x64, 0xf5, 0x32, 0xd5, 0x3f, 0x82, 0xf0, 0xa5,
	0x92, 0xd4, 0x69, 0x07, 0xa8, 0xce, 0x71, 0xe5, 0x8b, 0xd7, 0xda, 0xf0, 0xe2, 0x76, 0x76, 0x0e,
	0xf0, 0x4f, 0x5a, 0x65, 0x7f, 0xbc, 0x62, 0xac, 0x3f, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x82,
	0x3b, 0x7a, 0x90, 0x89, 0x04, 0x00, 0x00,
}
