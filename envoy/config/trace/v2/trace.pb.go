// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v2/trace.proto

package envoy_config_trace_v2

import (
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v2/core"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	types "github.com/gogo/protobuf/types"
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

// Available Zipkin collector endpoint versions.
type ZipkinConfig_CollectorEndpointVersion int32

const (
	// Zipkin API v1, JSON over HTTP.
	// [#comment: The default implementation of Zipkin client before this field is added was only v1
	// and the way user configure this was by not explicitly specifying the version. Consequently,
	// before this is added, the corresponding Zipkin collector expected to receive v1 payload.
	// Hence the motivation of adding HTTP_JSON_V1 as the default is to avoid a breaking change when
	// user upgrading Envoy with this change. Furthermore, we also immediately deprecate this field,
	// since in Zipkin realm this v1 version is considered to be not preferable anymore.]
	ZipkinConfig_HTTP_JSON_V1 ZipkinConfig_CollectorEndpointVersion = 0 // Deprecated: Do not use.
	// Zipkin API v2, JSON over HTTP.
	ZipkinConfig_HTTP_JSON ZipkinConfig_CollectorEndpointVersion = 1
	// Zipkin API v2, protobuf over HTTP.
	ZipkinConfig_HTTP_PROTO ZipkinConfig_CollectorEndpointVersion = 2
	// [#not-implemented-hide:]
	ZipkinConfig_GRPC ZipkinConfig_CollectorEndpointVersion = 3
)

var ZipkinConfig_CollectorEndpointVersion_name = map[int32]string{
	0: "HTTP_JSON_V1",
	1: "HTTP_JSON",
	2: "HTTP_PROTO",
	3: "GRPC",
}

var ZipkinConfig_CollectorEndpointVersion_value = map[string]int32{
	"HTTP_JSON_V1": 0,
	"HTTP_JSON":    1,
	"HTTP_PROTO":   2,
	"GRPC":         3,
}

func (x ZipkinConfig_CollectorEndpointVersion) String() string {
	return proto.EnumName(ZipkinConfig_CollectorEndpointVersion_name, int32(x))
}

func (ZipkinConfig_CollectorEndpointVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{2, 0}
}

type OpenCensusConfig_TraceContext int32

const (
	// No-op default, no trace context is utilized.
	OpenCensusConfig_NONE OpenCensusConfig_TraceContext = 0
	// W3C Trace-Context format "traceparent:" header.
	OpenCensusConfig_TRACE_CONTEXT OpenCensusConfig_TraceContext = 1
	// Binary "grpc-trace-bin:" header.
	OpenCensusConfig_GRPC_TRACE_BIN OpenCensusConfig_TraceContext = 2
	// "X-Cloud-Trace-Context:" header.
	OpenCensusConfig_CLOUD_TRACE_CONTEXT OpenCensusConfig_TraceContext = 3
	// X-B3-* headers.
	OpenCensusConfig_B3 OpenCensusConfig_TraceContext = 4
)

var OpenCensusConfig_TraceContext_name = map[int32]string{
	0: "NONE",
	1: "TRACE_CONTEXT",
	2: "GRPC_TRACE_BIN",
	3: "CLOUD_TRACE_CONTEXT",
	4: "B3",
}

var OpenCensusConfig_TraceContext_value = map[string]int32{
	"NONE":                0,
	"TRACE_CONTEXT":       1,
	"GRPC_TRACE_BIN":      2,
	"CLOUD_TRACE_CONTEXT": 3,
	"B3":                  4,
}

func (x OpenCensusConfig_TraceContext) String() string {
	return proto.EnumName(OpenCensusConfig_TraceContext_name, int32(x))
}

func (OpenCensusConfig_TraceContext) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{5, 0}
}

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
	return fileDescriptor_0785d24fc8ab55c7, []int{0}
}

func (m *Tracing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing.Unmarshal(m, b)
}
func (m *Tracing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing.Marshal(b, m, deterministic)
}
func (m *Tracing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing.Merge(m, src)
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
	// - *envoy.tracers.opencensus*
	// - *envoy.tracers.xray*
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Trace driver specific configuration which depends on the driver being instantiated.
	// See the trace drivers for examples:
	//
	// - :ref:`LightstepConfig <envoy_api_msg_config.trace.v2.LightstepConfig>`
	// - :ref:`ZipkinConfig <envoy_api_msg_config.trace.v2.ZipkinConfig>`
	// - :ref:`DynamicOtConfig <envoy_api_msg_config.trace.v2.DynamicOtConfig>`
	// - :ref:`DatadogConfig <envoy_api_msg_config.trace.v2.DatadogConfig>`
	// - :ref:`OpenCensusConfig <envoy_api_msg_config.trace.v2.OpenCensusConfig>`
	// [#comment: TODO(marco) when XRay is implemented, uncomment the following; - :ref:`XRayConfig
	// <envoy_api_msg_config.trace.v2.XRayConfig>`]
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
	return fileDescriptor_0785d24fc8ab55c7, []int{0, 0}
}

func (m *Tracing_Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracing_Http.Unmarshal(m, b)
}
func (m *Tracing_Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracing_Http.Marshal(b, m, deterministic)
}
func (m *Tracing_Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracing_Http.Merge(m, src)
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

// Deprecated: Do not use.
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

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Tracing_Http) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Tracing_Http_Config)(nil),
		(*Tracing_Http_TypedConfig)(nil),
	}
}

// Configuration for the LightStep tracer.
type LightstepConfig struct {
	// The cluster manager cluster that hosts the LightStep collectors.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	// File containing the access token to the `LightStep
	// <https://lightstep.com/>`_ API.
	AccessTokenFile      string   `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile,proto3" json:"access_token_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LightstepConfig) Reset()         { *m = LightstepConfig{} }
func (m *LightstepConfig) String() string { return proto.CompactTextString(m) }
func (*LightstepConfig) ProtoMessage()    {}
func (*LightstepConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{1}
}

func (m *LightstepConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightstepConfig.Unmarshal(m, b)
}
func (m *LightstepConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightstepConfig.Marshal(b, m, deterministic)
}
func (m *LightstepConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightstepConfig.Merge(m, src)
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

// Configuration for the Zipkin tracer.
// [#next-free-field: 6]
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
	// Determines whether client and server spans will share the same span context.
	// The default value is true.
	SharedSpanContext *types.BoolValue `protobuf:"bytes,4,opt,name=shared_span_context,json=sharedSpanContext,proto3" json:"shared_span_context,omitempty"`
	// Determines the selected collector endpoint version. By default, the ``HTTP_JSON_V1`` will be
	// used.
	CollectorEndpointVersion ZipkinConfig_CollectorEndpointVersion `protobuf:"varint,5,opt,name=collector_endpoint_version,json=collectorEndpointVersion,proto3,enum=envoy.config.trace.v2.ZipkinConfig_CollectorEndpointVersion" json:"collector_endpoint_version,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                              `json:"-"`
	XXX_unrecognized         []byte                                `json:"-"`
	XXX_sizecache            int32                                 `json:"-"`
}

func (m *ZipkinConfig) Reset()         { *m = ZipkinConfig{} }
func (m *ZipkinConfig) String() string { return proto.CompactTextString(m) }
func (*ZipkinConfig) ProtoMessage()    {}
func (*ZipkinConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{2}
}

func (m *ZipkinConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZipkinConfig.Unmarshal(m, b)
}
func (m *ZipkinConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZipkinConfig.Marshal(b, m, deterministic)
}
func (m *ZipkinConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZipkinConfig.Merge(m, src)
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

func (m *ZipkinConfig) GetCollectorEndpointVersion() ZipkinConfig_CollectorEndpointVersion {
	if m != nil {
		return m.CollectorEndpointVersion
	}
	return ZipkinConfig_HTTP_JSON_V1
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
	return fileDescriptor_0785d24fc8ab55c7, []int{3}
}

func (m *DynamicOtConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicOtConfig.Unmarshal(m, b)
}
func (m *DynamicOtConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicOtConfig.Marshal(b, m, deterministic)
}
func (m *DynamicOtConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicOtConfig.Merge(m, src)
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
	return fileDescriptor_0785d24fc8ab55c7, []int{4}
}

func (m *DatadogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatadogConfig.Unmarshal(m, b)
}
func (m *DatadogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatadogConfig.Marshal(b, m, deterministic)
}
func (m *DatadogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatadogConfig.Merge(m, src)
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

// Configuration for the OpenCensus tracer.
// [#next-free-field: 13]
type OpenCensusConfig struct {
	// Configures tracing, e.g. the sampler, max number of annotations, etc.
	TraceConfig *v1.TraceConfig `protobuf:"bytes,1,opt,name=trace_config,json=traceConfig,proto3" json:"trace_config,omitempty"`
	// Enables the stdout exporter if set to true. This is intended for debugging
	// purposes.
	StdoutExporterEnabled bool `protobuf:"varint,2,opt,name=stdout_exporter_enabled,json=stdoutExporterEnabled,proto3" json:"stdout_exporter_enabled,omitempty"`
	// Enables the Stackdriver exporter if set to true. The project_id must also
	// be set.
	StackdriverExporterEnabled bool `protobuf:"varint,3,opt,name=stackdriver_exporter_enabled,json=stackdriverExporterEnabled,proto3" json:"stackdriver_exporter_enabled,omitempty"`
	// The Cloud project_id to use for Stackdriver tracing.
	StackdriverProjectId string `protobuf:"bytes,4,opt,name=stackdriver_project_id,json=stackdriverProjectId,proto3" json:"stackdriver_project_id,omitempty"`
	// (optional) By default, the Stackdriver exporter will connect to production
	// Stackdriver. If stackdriver_address is non-empty, it will instead connect
	// to this address, which is in the gRPC format:
	// https://github.com/grpc/grpc/blob/master/doc/naming.md
	StackdriverAddress string `protobuf:"bytes,10,opt,name=stackdriver_address,json=stackdriverAddress,proto3" json:"stackdriver_address,omitempty"`
	// Enables the Zipkin exporter if set to true. The url and service name must
	// also be set.
	ZipkinExporterEnabled bool `protobuf:"varint,5,opt,name=zipkin_exporter_enabled,json=zipkinExporterEnabled,proto3" json:"zipkin_exporter_enabled,omitempty"`
	// The URL to Zipkin, e.g. "http://127.0.0.1:9411/api/v2/spans"
	ZipkinUrl string `protobuf:"bytes,6,opt,name=zipkin_url,json=zipkinUrl,proto3" json:"zipkin_url,omitempty"`
	// Enables the OpenCensus Agent exporter if set to true. The address must also
	// be set.
	OcagentExporterEnabled bool `protobuf:"varint,11,opt,name=ocagent_exporter_enabled,json=ocagentExporterEnabled,proto3" json:"ocagent_exporter_enabled,omitempty"`
	// The address of the OpenCensus Agent, if its exporter is enabled, in gRPC
	// format: https://github.com/grpc/grpc/blob/master/doc/naming.md
	OcagentAddress string `protobuf:"bytes,12,opt,name=ocagent_address,json=ocagentAddress,proto3" json:"ocagent_address,omitempty"`
	// List of incoming trace context headers we will accept. First one found
	// wins.
	IncomingTraceContext []OpenCensusConfig_TraceContext `protobuf:"varint,8,rep,packed,name=incoming_trace_context,json=incomingTraceContext,proto3,enum=envoy.config.trace.v2.OpenCensusConfig_TraceContext" json:"incoming_trace_context,omitempty"`
	// List of outgoing trace context headers we will produce.
	OutgoingTraceContext []OpenCensusConfig_TraceContext `protobuf:"varint,9,rep,packed,name=outgoing_trace_context,json=outgoingTraceContext,proto3,enum=envoy.config.trace.v2.OpenCensusConfig_TraceContext" json:"outgoing_trace_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *OpenCensusConfig) Reset()         { *m = OpenCensusConfig{} }
func (m *OpenCensusConfig) String() string { return proto.CompactTextString(m) }
func (*OpenCensusConfig) ProtoMessage()    {}
func (*OpenCensusConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{5}
}

func (m *OpenCensusConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenCensusConfig.Unmarshal(m, b)
}
func (m *OpenCensusConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenCensusConfig.Marshal(b, m, deterministic)
}
func (m *OpenCensusConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenCensusConfig.Merge(m, src)
}
func (m *OpenCensusConfig) XXX_Size() int {
	return xxx_messageInfo_OpenCensusConfig.Size(m)
}
func (m *OpenCensusConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenCensusConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OpenCensusConfig proto.InternalMessageInfo

func (m *OpenCensusConfig) GetTraceConfig() *v1.TraceConfig {
	if m != nil {
		return m.TraceConfig
	}
	return nil
}

func (m *OpenCensusConfig) GetStdoutExporterEnabled() bool {
	if m != nil {
		return m.StdoutExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetStackdriverExporterEnabled() bool {
	if m != nil {
		return m.StackdriverExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetStackdriverProjectId() string {
	if m != nil {
		return m.StackdriverProjectId
	}
	return ""
}

func (m *OpenCensusConfig) GetStackdriverAddress() string {
	if m != nil {
		return m.StackdriverAddress
	}
	return ""
}

func (m *OpenCensusConfig) GetZipkinExporterEnabled() bool {
	if m != nil {
		return m.ZipkinExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetZipkinUrl() string {
	if m != nil {
		return m.ZipkinUrl
	}
	return ""
}

func (m *OpenCensusConfig) GetOcagentExporterEnabled() bool {
	if m != nil {
		return m.OcagentExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetOcagentAddress() string {
	if m != nil {
		return m.OcagentAddress
	}
	return ""
}

func (m *OpenCensusConfig) GetIncomingTraceContext() []OpenCensusConfig_TraceContext {
	if m != nil {
		return m.IncomingTraceContext
	}
	return nil
}

func (m *OpenCensusConfig) GetOutgoingTraceContext() []OpenCensusConfig_TraceContext {
	if m != nil {
		return m.OutgoingTraceContext
	}
	return nil
}

// [#not-implemented-hide:]
// Configuration for AWS X-Ray tracer.
type XRayConfig struct {
	// The endpoint of the X-Ray Daemon where the spans will be sent. Since by default daemon
	// listens to localhost:2000, so the default value is 127.0.0.1:2000.
	DaemonEndpoint string `protobuf:"bytes,1,opt,name=daemon_endpoint,json=daemonEndpoint,proto3" json:"daemon_endpoint,omitempty"`
	// The custom name to name a X-Ray segment. By default will use cluster name.
	SegmentName string `protobuf:"bytes,2,opt,name=segment_name,json=segmentName,proto3" json:"segment_name,omitempty"`
	// The location of custom sampling rule json file.
	SamplingRuleManifest *core.DataSource `protobuf:"bytes,3,opt,name=sampling_rule_manifest,json=samplingRuleManifest,proto3" json:"sampling_rule_manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *XRayConfig) Reset()         { *m = XRayConfig{} }
func (m *XRayConfig) String() string { return proto.CompactTextString(m) }
func (*XRayConfig) ProtoMessage()    {}
func (*XRayConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0785d24fc8ab55c7, []int{6}
}

func (m *XRayConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XRayConfig.Unmarshal(m, b)
}
func (m *XRayConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XRayConfig.Marshal(b, m, deterministic)
}
func (m *XRayConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XRayConfig.Merge(m, src)
}
func (m *XRayConfig) XXX_Size() int {
	return xxx_messageInfo_XRayConfig.Size(m)
}
func (m *XRayConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_XRayConfig.DiscardUnknown(m)
}

var xxx_messageInfo_XRayConfig proto.InternalMessageInfo

func (m *XRayConfig) GetDaemonEndpoint() string {
	if m != nil {
		return m.DaemonEndpoint
	}
	return ""
}

func (m *XRayConfig) GetSegmentName() string {
	if m != nil {
		return m.SegmentName
	}
	return ""
}

func (m *XRayConfig) GetSamplingRuleManifest() *core.DataSource {
	if m != nil {
		return m.SamplingRuleManifest
	}
	return nil
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
	return fileDescriptor_0785d24fc8ab55c7, []int{7}
}

func (m *TraceServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceServiceConfig.Unmarshal(m, b)
}
func (m *TraceServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceServiceConfig.Marshal(b, m, deterministic)
}
func (m *TraceServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceServiceConfig.Merge(m, src)
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
	proto.RegisterEnum("envoy.config.trace.v2.ZipkinConfig_CollectorEndpointVersion", ZipkinConfig_CollectorEndpointVersion_name, ZipkinConfig_CollectorEndpointVersion_value)
	proto.RegisterEnum("envoy.config.trace.v2.OpenCensusConfig_TraceContext", OpenCensusConfig_TraceContext_name, OpenCensusConfig_TraceContext_value)
	proto.RegisterType((*Tracing)(nil), "envoy.config.trace.v2.Tracing")
	proto.RegisterType((*Tracing_Http)(nil), "envoy.config.trace.v2.Tracing.Http")
	proto.RegisterType((*LightstepConfig)(nil), "envoy.config.trace.v2.LightstepConfig")
	proto.RegisterType((*ZipkinConfig)(nil), "envoy.config.trace.v2.ZipkinConfig")
	proto.RegisterType((*DynamicOtConfig)(nil), "envoy.config.trace.v2.DynamicOtConfig")
	proto.RegisterType((*DatadogConfig)(nil), "envoy.config.trace.v2.DatadogConfig")
	proto.RegisterType((*OpenCensusConfig)(nil), "envoy.config.trace.v2.OpenCensusConfig")
	proto.RegisterType((*XRayConfig)(nil), "envoy.config.trace.v2.XRayConfig")
	proto.RegisterType((*TraceServiceConfig)(nil), "envoy.config.trace.v2.TraceServiceConfig")
}

func init() { proto.RegisterFile("envoy/config/trace/v2/trace.proto", fileDescriptor_0785d24fc8ab55c7) }

var fileDescriptor_0785d24fc8ab55c7 = []byte{
	// 1100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x51, 0x6e, 0xdb, 0x46,
	0x13, 0x36, 0x65, 0xc5, 0x96, 0x47, 0xb2, 0x44, 0x6f, 0x1c, 0x5b, 0xbf, 0xfe, 0x24, 0xb0, 0x95,
	0x22, 0x35, 0x8a, 0x82, 0xaa, 0x65, 0x37, 0x4d, 0x81, 0x3e, 0xd4, 0x92, 0xd5, 0xd8, 0x6e, 0x2a,
	0x09, 0x94, 0x62, 0xa4, 0x7d, 0x61, 0x57, 0xe4, 0x9a, 0x66, 0x4c, 0x71, 0xd9, 0xe5, 0x4a, 0xb5,
	0x82, 0x9e, 0xa0, 0xe7, 0xe8, 0x19, 0x8a, 0xde, 0xa0, 0x67, 0xe9, 0x15, 0xf2, 0x54, 0x70, 0x77,
	0x69, 0x31, 0x94, 0x0d, 0x14, 0xf0, 0x1b, 0x77, 0xbe, 0xef, 0x9b, 0x9d, 0x99, 0x9d, 0x19, 0x09,
	0x76, 0x49, 0x30, 0xa5, 0xb3, 0x86, 0x4d, 0x83, 0x0b, 0xcf, 0x6d, 0x70, 0x86, 0x6d, 0xd2, 0x98,
	0x36, 0xe5, 0x87, 0x11, 0x32, 0xca, 0x29, 0x7a, 0x24, 0x28, 0x86, 0xa4, 0x18, 0x12, 0x99, 0x36,
	0x6b, 0x8f, 0xa5, 0x12, 0x87, 0x5e, 0x2c, 0xb0, 0x29, 0x23, 0x8d, 0x11, 0x8e, 0x94, 0xa8, 0xf6,
	0xc9, 0x22, 0xea, 0xb2, 0xd0, 0xb6, 0x22, 0xc2, 0xa6, 0x5e, 0xe2, 0xba, 0xf6, 0x3f, 0x97, 0x52,
	0xd7, 0x27, 0x0d, 0x71, 0x1a, 0x4d, 0x2e, 0x1a, 0x38, 0x98, 0x29, 0xe8, 0x71, 0x16, 0x8a, 0x38,
	0x9b, 0xd8, 0x5c, 0xa1, 0x4f, 0xb3, 0xe8, 0xaf, 0x0c, 0x87, 0x21, 0x61, 0x91, 0xc2, 0x3f, 0xa7,
	0x21, 0x09, 0x6c, 0x12, 0x44, 0x93, 0x48, 0x72, 0x92, 0xd4, 0xf6, 0xe5, 0x87, 0xa5, 0xb2, 0x91,
	0xec, 0xed, 0x29, 0xf6, 0x3d, 0x07, 0x73, 0xd2, 0x48, 0x3e, 0x24, 0x50, 0xff, 0x47, 0x83, 0xd5,
	0x21, 0xc3, 0xb6, 0x17, 0xb8, 0xe8, 0x2b, 0xc8, 0x5f, 0x72, 0x1e, 0x56, 0xb5, 0x1d, 0x6d, 0xaf,
	0xd8, 0x7c, 0x66, 0xdc, 0x5a, 0x15, 0x43, 0xb1, 0x8d, 0x13, 0xce, 0x43, 0x53, 0x08, 0x6a, 0x7f,
	0x68, 0x90, 0x8f, 0x8f, 0xe8, 0xff, 0x90, 0x0f, 0xf0, 0x98, 0x08, 0x0f, 0x6b, 0xad, 0xd5, 0x0f,
	0xad, 0x3c, 0xcb, 0xed, 0x68, 0xa6, 0x30, 0xa2, 0x2f, 0x61, 0x45, 0xfa, 0xaa, 0xe6, 0xc4, 0x05,
	0xdb, 0x86, 0x4c, 0xd1, 0x48, 0x52, 0x34, 0x06, 0xa2, 0x00, 0xad, 0x5c, 0x55, 0x3b, 0x59, 0x32,
	0x15, 0x19, 0x7d, 0x0d, 0x25, 0x3e, 0x0b, 0x89, 0xa3, 0x12, 0xaa, 0x2e, 0x0b, 0xf1, 0xe6, 0x82,
	0xf8, 0x28, 0x98, 0x9d, 0x2c, 0x99, 0x45, 0xc1, 0x6d, 0x0b, 0x6a, 0x6b, 0x1d, 0x8a, 0x52, 0x64,
	0xc5, 0xd6, 0xfa, 0x6f, 0x50, 0x79, 0xed, 0xb9, 0x97, 0x3c, 0xe2, 0x24, 0x94, 0x0c, 0x74, 0x08,
	0x1b, 0x36, 0xf5, 0x7d, 0x62, 0x73, 0xca, 0x2c, 0xdb, 0x9f, 0x44, 0x9c, 0xb0, 0x6c, 0xf4, 0xfa,
	0x0d, 0xa3, 0x2d, 0x09, 0xe8, 0x00, 0x36, 0xb0, 0x6d, 0x93, 0x28, 0xb2, 0x38, 0xbd, 0x22, 0x81,
	0x75, 0xe1, 0xf9, 0x44, 0x24, 0x95, 0x52, 0x55, 0x24, 0x63, 0x18, 0x13, 0xbe, 0xf3, 0x7c, 0x52,
	0xff, 0x7b, 0x19, 0x4a, 0x3f, 0x79, 0xe1, 0x95, 0x17, 0xdc, 0xeb, 0xee, 0x17, 0x80, 0xe6, 0x2a,
	0x12, 0x38, 0x21, 0xf5, 0x02, 0x9e, 0xbd, 0x7c, 0xee, 0xb8, 0xa3, 0x18, 0xe8, 0x39, 0x54, 0x64,
	0x5f, 0x78, 0x8e, 0xb5, 0xdf, 0x7c, 0x39, 0xf2, 0xb8, 0xa8, 0x64, 0xc1, 0x5c, 0x17, 0xe6, 0x53,
	0x47, 0x1a, 0xd1, 0x19, 0x3c, 0x8c, 0x2e, 0x31, 0x23, 0x8e, 0x15, 0x85, 0x38, 0x88, 0x8b, 0xce,
	0xc9, 0x35, 0xaf, 0xe6, 0x45, 0xd5, 0x6b, 0x0b, 0x55, 0x6f, 0x51, 0xea, 0x9f, 0x63, 0x7f, 0x42,
	0xcc, 0x0d, 0x29, 0x1b, 0x84, 0x38, 0x4e, 0x30, 0x16, 0xa1, 0xf7, 0x50, 0x5b, 0x8c, 0xd5, 0x9a,
	0x12, 0x16, 0x79, 0x34, 0xa8, 0x3e, 0xd8, 0xd1, 0xf6, 0xca, 0xcd, 0x6f, 0xee, 0x68, 0xb3, 0x74,
	0xa9, 0x8c, 0x76, 0x36, 0x9d, 0x73, 0xe9, 0xc3, 0xac, 0xda, 0x77, 0x20, 0xf5, 0x1f, 0xa1, 0x7a,
	0x97, 0x0a, 0x6d, 0x42, 0xe9, 0x64, 0x38, 0xec, 0x5b, 0x67, 0x83, 0x5e, 0xd7, 0x3a, 0xdf, 0xd7,
	0x97, 0x6a, 0xb9, 0x82, 0x86, 0xd6, 0x61, 0xed, 0xc6, 0xaa, 0x6b, 0xa8, 0x0c, 0x20, 0x8e, 0x7d,
	0xb3, 0x37, 0xec, 0xe9, 0x39, 0x54, 0x80, 0xfc, 0x2b, 0xb3, 0xdf, 0xd6, 0x97, 0xeb, 0x04, 0x2a,
	0xc7, 0xb3, 0x00, 0x8f, 0x3d, 0xbb, 0xc7, 0xd5, 0x5b, 0xee, 0xc2, 0xaa, 0xef, 0x8d, 0x18, 0x66,
	0xb3, 0xec, 0x0b, 0x26, 0x76, 0xd4, 0xf8, 0x8f, 0xed, 0x9f, 0x34, 0x7e, 0xfd, 0x17, 0x58, 0x3f,
	0xc6, 0x1c, 0x3b, 0xd4, 0xbd, 0x57, 0xc3, 0x7c, 0x06, 0x25, 0xb5, 0x92, 0x2c, 0x31, 0x9b, 0x99,
	0x56, 0x29, 0x2a, 0xb0, 0x8b, 0xc7, 0xa4, 0xfe, 0xd7, 0x0a, 0xe8, 0xbd, 0x90, 0x04, 0x6d, 0xb1,
	0x57, 0xd4, 0xb5, 0xa7, 0x50, 0x4a, 0x6f, 0x14, 0xb5, 0x1e, 0x9e, 0x1b, 0xf3, 0x05, 0x24, 0x53,
	0x48, 0xde, 0x6e, 0x5f, 0xac, 0x08, 0x22, 0xd5, 0x66, 0x91, 0xcf, 0x0f, 0xe8, 0x05, 0x6c, 0x47,
	0xdc, 0xa1, 0x13, 0x6e, 0x91, 0xeb, 0x90, 0x32, 0x4e, 0xe2, 0xb6, 0xc0, 0x23, 0x9f, 0x38, 0x22,
	0xac, 0x82, 0xf9, 0x48, 0xc2, 0x1d, 0x85, 0x76, 0x24, 0x88, 0xbe, 0x85, 0xc7, 0x11, 0xc7, 0xf6,
	0x95, 0xc3, 0xbc, 0x69, 0xac, 0xc9, 0x8a, 0x65, 0x27, 0xd7, 0x52, 0x9c, 0xac, 0x87, 0x43, 0xd8,
	0x4a, 0x7b, 0x08, 0x19, 0x7d, 0x47, 0x6c, 0x6e, 0x79, 0x8e, 0xe8, 0xec, 0x35, 0x73, 0x33, 0x85,
	0xf6, 0x25, 0x78, 0xea, 0xa0, 0x06, 0x3c, 0x4c, 0xab, 0xb0, 0xe3, 0x30, 0x12, 0x45, 0x55, 0x10,
	0x12, 0x94, 0x82, 0x8e, 0x24, 0x12, 0x27, 0xf8, 0x5e, 0x34, 0xee, 0x62, 0x8c, 0x0f, 0x64, 0x82,
	0x12, 0xce, 0x86, 0xf7, 0x04, 0x40, 0xe9, 0x26, 0xcc, 0xaf, 0xae, 0x08, 0xff, 0x6b, 0xd2, 0xf2,
	0x86, 0xf9, 0xe8, 0x25, 0x54, 0xa9, 0x8d, 0x5d, 0x12, 0xdc, 0x52, 0xb8, 0xa2, 0xf0, 0xbb, 0xa5,
	0xf0, 0xac, 0xe3, 0x4f, 0xa1, 0x92, 0x28, 0x93, 0xe8, 0x4b, 0xc2, 0x7b, 0x59, 0x99, 0x93, 0xc8,
	0xdf, 0xc1, 0x96, 0x17, 0xd8, 0x74, 0xec, 0x05, 0xae, 0x75, 0xf3, 0xdc, 0x62, 0xf4, 0x0b, 0x3b,
	0xcb, 0x7b, 0xe5, 0xe6, 0xe1, 0x1d, 0x73, 0x9a, 0x6d, 0x97, 0x9b, 0xc7, 0x8f, 0xb5, 0xe6, 0x66,
	0xe2, 0x33, 0x6d, 0x8d, 0xef, 0xa2, 0x13, 0xee, 0xd2, 0xc5, 0xbb, 0xd6, 0xee, 0x73, 0x57, 0xe2,
	0x33, 0x6d, 0xad, 0xff, 0x0c, 0xa5, 0x8f, 0xee, 0x2e, 0x40, 0xbe, 0xdb, 0xeb, 0x76, 0xf4, 0x25,
	0xb4, 0x01, 0xeb, 0x43, 0xf3, 0xa8, 0xdd, 0xb1, 0xda, 0xbd, 0xee, 0xb0, 0xf3, 0x76, 0xa8, 0x6b,
	0x08, 0x41, 0x39, 0x9e, 0x71, 0x4b, 0xda, 0x5b, 0xa7, 0x5d, 0x3d, 0x87, 0xb6, 0xe1, 0x61, 0xfb,
	0x75, 0xef, 0xcd, 0xb1, 0xf5, 0x31, 0x79, 0x19, 0xad, 0x40, 0xae, 0x75, 0xa0, 0xe7, 0xcf, 0xf2,
	0x85, 0x55, 0xbd, 0x50, 0xff, 0x53, 0x03, 0x78, 0x6b, 0xe2, 0x99, 0xea, 0xf4, 0x2f, 0xa0, 0xe2,
	0x60, 0x32, 0xa6, 0xc1, 0x7c, 0x47, 0x67, 0x26, 0xb5, 0x2c, 0xf1, 0x9b, 0x05, 0xbd, 0x1b, 0xcf,
	0xa9, 0x3b, 0x8e, 0x5f, 0x6a, 0x3e, 0xa7, 0xf1, 0x78, 0x0a, 0x5b, 0x3c, 0x9e, 0x68, 0x00, 0x5b,
	0x11, 0x1e, 0x87, 0x7e, 0x5c, 0x37, 0x36, 0xf1, 0x89, 0x35, 0xc6, 0x81, 0x77, 0x41, 0x22, 0xae,
	0x7e, 0x14, 0x9f, 0xa8, 0xba, 0xe1, 0xd0, 0x8b, 0xcb, 0x15, 0xff, 0x27, 0x31, 0xe2, 0x15, 0x32,
	0xa0, 0x13, 0x66, 0x13, 0x73, 0x33, 0x11, 0x9b, 0x13, 0x9f, 0xfc, 0xa0, 0xa4, 0x75, 0x0c, 0x48,
	0x14, 0x68, 0x20, 0xf7, 0x80, 0x8a, 0xff, 0x7b, 0x28, 0xa5, 0xff, 0xcd, 0xa8, 0xa1, 0x7f, 0x7a,
	0xcb, 0x05, 0xaf, 0x58, 0x68, 0x2b, 0x6d, 0xab, 0xf0, 0xa1, 0xf5, 0xe0, 0x77, 0x2d, 0xa7, 0x6b,
	0x66, 0xd1, 0x4d, 0x99, 0x0f, 0xe0, 0x99, 0x47, 0xa5, 0x34, 0x64, 0xf4, 0x7a, 0x76, 0xfb, 0xf3,
	0xb6, 0x40, 0xc4, 0xd1, 0x8f, 0x77, 0x49, 0x5f, 0x1b, 0xad, 0x88, 0xa5, 0x72, 0xf0, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x09, 0x25, 0xeb, 0xa4, 0xc0, 0x09, 0x00, 0x00,
}
