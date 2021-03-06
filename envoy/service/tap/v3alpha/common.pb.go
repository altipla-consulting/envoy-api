// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/tap/v3alpha/common.proto

package envoy_service_tap_v3alpha

import (
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v3alpha/core"
	route "github.com/altipla-consulting/envoy-api/envoy/api/v3alpha/route"
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

// Output format. All output is in the form of one or more :ref:`TraceWrapper
// <envoy_api_msg_data.tap.v3alpha.TraceWrapper>` messages. This enumeration indicates
// how those messages are written. Note that not all sinks support all output formats. See
// individual sink documentation for more information.
type OutputSink_Format int32

const (
	// Each message will be written as JSON. Any :ref:`body <envoy_api_msg_data.tap.v3alpha.Body>`
	// data will be present in the :ref:`as_bytes
	// <envoy_api_field_data.tap.v3alpha.Body.as_bytes>` field. This means that body data will be
	// base64 encoded as per the `proto3 JSON mappings
	// <https://developers.google.com/protocol-buffers/docs/proto3#json>`_.
	OutputSink_JSON_BODY_AS_BYTES OutputSink_Format = 0
	// Each message will be written as JSON. Any :ref:`body <envoy_api_msg_data.tap.v3alpha.Body>`
	// data will be present in the :ref:`as_string
	// <envoy_api_field_data.tap.v3alpha.Body.as_string>` field. This means that body data will be
	// string encoded as per the `proto3 JSON mappings
	// <https://developers.google.com/protocol-buffers/docs/proto3#json>`_. This format type is
	// useful when it is known that that body is human readable (e.g., JSON over HTTP) and the
	// user wishes to view it directly without being forced to base64 decode the body.
	OutputSink_JSON_BODY_AS_STRING OutputSink_Format = 1
	// Binary proto format. Note that binary proto is not self-delimiting. If a sink writes
	// multiple binary messages without any length information the data stream will not be
	// useful. However, for certain sinks that are self-delimiting (e.g., one message per file)
	// this output format makes consumption simpler.
	OutputSink_PROTO_BINARY OutputSink_Format = 2
	// Messages are written as a sequence tuples, where each tuple is the message length encoded
	// as a `protobuf 32-bit varint
	// <https://developers.google.com/protocol-buffers/docs/reference/cpp/google.protobuf.io.coded_stream>`_
	// followed by the binary message. The messages can be read back using the language specific
	// protobuf coded stream implementation to obtain the message length and the message.
	OutputSink_PROTO_BINARY_LENGTH_DELIMITED OutputSink_Format = 3
	// Text proto format.
	OutputSink_PROTO_TEXT OutputSink_Format = 4
)

var OutputSink_Format_name = map[int32]string{
	0: "JSON_BODY_AS_BYTES",
	1: "JSON_BODY_AS_STRING",
	2: "PROTO_BINARY",
	3: "PROTO_BINARY_LENGTH_DELIMITED",
	4: "PROTO_TEXT",
}

var OutputSink_Format_value = map[string]int32{
	"JSON_BODY_AS_BYTES":            0,
	"JSON_BODY_AS_STRING":           1,
	"PROTO_BINARY":                  2,
	"PROTO_BINARY_LENGTH_DELIMITED": 3,
	"PROTO_TEXT":                    4,
}

func (x OutputSink_Format) String() string {
	return proto.EnumName(OutputSink_Format_name, int32(x))
}

func (OutputSink_Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ab4dd2506d34575, []int{4, 0}
}

// Tap configuration.
type TapConfig struct {
	// The match configuration. If the configuration matches the data source being tapped, a tap will
	// occur, with the result written to the configured output.
	MatchConfig *MatchPredicate `protobuf:"bytes,1,opt,name=match_config,json=matchConfig,proto3" json:"match_config,omitempty"`
	// The tap output configuration. If a match configuration matches a data source being tapped,
	// a tap will occur and the data will be written to the configured output.
	OutputConfig *OutputConfig `protobuf:"bytes,2,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
	// [#not-implemented-hide:] Specify if Tap matching is enabled. The % of requests\connections for
	// which the tap matching is enabled. When not enabled, the request\connection will not be
	// recorded.
	//
	// .. note::
	//
	//   This field defaults to 100/:ref:`HUNDRED
	//   <envoy_api_enum_type.v3alpha.FractionalPercent.DenominatorType>`.
	TapEnabled           *core.RuntimeFractionalPercent `protobuf:"bytes,3,opt,name=tap_enabled,json=tapEnabled,proto3" json:"tap_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TapConfig) Reset()         { *m = TapConfig{} }
func (m *TapConfig) String() string { return proto.CompactTextString(m) }
func (*TapConfig) ProtoMessage()    {}
func (*TapConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab4dd2506d34575, []int{0}
}

func (m *TapConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TapConfig.Unmarshal(m, b)
}
func (m *TapConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TapConfig.Marshal(b, m, deterministic)
}
func (m *TapConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapConfig.Merge(m, src)
}
func (m *TapConfig) XXX_Size() int {
	return xxx_messageInfo_TapConfig.Size(m)
}
func (m *TapConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TapConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TapConfig proto.InternalMessageInfo

func (m *TapConfig) GetMatchConfig() *MatchPredicate {
	if m != nil {
		return m.MatchConfig
	}
	return nil
}

func (m *TapConfig) GetOutputConfig() *OutputConfig {
	if m != nil {
		return m.OutputConfig
	}
	return nil
}

func (m *TapConfig) GetTapEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.TapEnabled
	}
	return nil
}

// Tap match configuration. This is a recursive structure which allows complex nested match
// configurations to be built using various logical operators.
// [#next-free-field: 9]
type MatchPredicate struct {
	// Types that are valid to be assigned to Rule:
	//	*MatchPredicate_OrMatch
	//	*MatchPredicate_AndMatch
	//	*MatchPredicate_NotMatch
	//	*MatchPredicate_AnyMatch
	//	*MatchPredicate_HttpRequestHeadersMatch
	//	*MatchPredicate_HttpRequestTrailersMatch
	//	*MatchPredicate_HttpResponseHeadersMatch
	//	*MatchPredicate_HttpResponseTrailersMatch
	Rule                 isMatchPredicate_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MatchPredicate) Reset()         { *m = MatchPredicate{} }
func (m *MatchPredicate) String() string { return proto.CompactTextString(m) }
func (*MatchPredicate) ProtoMessage()    {}
func (*MatchPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab4dd2506d34575, []int{1}
}

func (m *MatchPredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchPredicate.Unmarshal(m, b)
}
func (m *MatchPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchPredicate.Marshal(b, m, deterministic)
}
func (m *MatchPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchPredicate.Merge(m, src)
}
func (m *MatchPredicate) XXX_Size() int {
	return xxx_messageInfo_MatchPredicate.Size(m)
}
func (m *MatchPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_MatchPredicate proto.InternalMessageInfo

type isMatchPredicate_Rule interface {
	isMatchPredicate_Rule()
}

type MatchPredicate_OrMatch struct {
	OrMatch *MatchPredicate_MatchSet `protobuf:"bytes,1,opt,name=or_match,json=orMatch,proto3,oneof"`
}

type MatchPredicate_AndMatch struct {
	AndMatch *MatchPredicate_MatchSet `protobuf:"bytes,2,opt,name=and_match,json=andMatch,proto3,oneof"`
}

type MatchPredicate_NotMatch struct {
	NotMatch *MatchPredicate `protobuf:"bytes,3,opt,name=not_match,json=notMatch,proto3,oneof"`
}

type MatchPredicate_AnyMatch struct {
	AnyMatch bool `protobuf:"varint,4,opt,name=any_match,json=anyMatch,proto3,oneof"`
}

type MatchPredicate_HttpRequestHeadersMatch struct {
	HttpRequestHeadersMatch *HttpHeadersMatch `protobuf:"bytes,5,opt,name=http_request_headers_match,json=httpRequestHeadersMatch,proto3,oneof"`
}

type MatchPredicate_HttpRequestTrailersMatch struct {
	HttpRequestTrailersMatch *HttpHeadersMatch `protobuf:"bytes,6,opt,name=http_request_trailers_match,json=httpRequestTrailersMatch,proto3,oneof"`
}

type MatchPredicate_HttpResponseHeadersMatch struct {
	HttpResponseHeadersMatch *HttpHeadersMatch `protobuf:"bytes,7,opt,name=http_response_headers_match,json=httpResponseHeadersMatch,proto3,oneof"`
}

type MatchPredicate_HttpResponseTrailersMatch struct {
	HttpResponseTrailersMatch *HttpHeadersMatch `protobuf:"bytes,8,opt,name=http_response_trailers_match,json=httpResponseTrailersMatch,proto3,oneof"`
}

func (*MatchPredicate_OrMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_AndMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_NotMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_AnyMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpRequestHeadersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpRequestTrailersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpResponseHeadersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpResponseTrailersMatch) isMatchPredicate_Rule() {}

func (m *MatchPredicate) GetRule() isMatchPredicate_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *MatchPredicate) GetOrMatch() *MatchPredicate_MatchSet {
	if x, ok := m.GetRule().(*MatchPredicate_OrMatch); ok {
		return x.OrMatch
	}
	return nil
}

func (m *MatchPredicate) GetAndMatch() *MatchPredicate_MatchSet {
	if x, ok := m.GetRule().(*MatchPredicate_AndMatch); ok {
		return x.AndMatch
	}
	return nil
}

func (m *MatchPredicate) GetNotMatch() *MatchPredicate {
	if x, ok := m.GetRule().(*MatchPredicate_NotMatch); ok {
		return x.NotMatch
	}
	return nil
}

func (m *MatchPredicate) GetAnyMatch() bool {
	if x, ok := m.GetRule().(*MatchPredicate_AnyMatch); ok {
		return x.AnyMatch
	}
	return false
}

func (m *MatchPredicate) GetHttpRequestHeadersMatch() *HttpHeadersMatch {
	if x, ok := m.GetRule().(*MatchPredicate_HttpRequestHeadersMatch); ok {
		return x.HttpRequestHeadersMatch
	}
	return nil
}

func (m *MatchPredicate) GetHttpRequestTrailersMatch() *HttpHeadersMatch {
	if x, ok := m.GetRule().(*MatchPredicate_HttpRequestTrailersMatch); ok {
		return x.HttpRequestTrailersMatch
	}
	return nil
}

func (m *MatchPredicate) GetHttpResponseHeadersMatch() *HttpHeadersMatch {
	if x, ok := m.GetRule().(*MatchPredicate_HttpResponseHeadersMatch); ok {
		return x.HttpResponseHeadersMatch
	}
	return nil
}

func (m *MatchPredicate) GetHttpResponseTrailersMatch() *HttpHeadersMatch {
	if x, ok := m.GetRule().(*MatchPredicate_HttpResponseTrailersMatch); ok {
		return x.HttpResponseTrailersMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MatchPredicate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MatchPredicate_OrMatch)(nil),
		(*MatchPredicate_AndMatch)(nil),
		(*MatchPredicate_NotMatch)(nil),
		(*MatchPredicate_AnyMatch)(nil),
		(*MatchPredicate_HttpRequestHeadersMatch)(nil),
		(*MatchPredicate_HttpRequestTrailersMatch)(nil),
		(*MatchPredicate_HttpResponseHeadersMatch)(nil),
		(*MatchPredicate_HttpResponseTrailersMatch)(nil),
	}
}

// A set of match configurations used for logical operations.
type MatchPredicate_MatchSet struct {
	// The list of rules that make up the set.
	Rules                []*MatchPredicate `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MatchPredicate_MatchSet) Reset()         { *m = MatchPredicate_MatchSet{} }
func (m *MatchPredicate_MatchSet) String() string { return proto.CompactTextString(m) }
func (*MatchPredicate_MatchSet) ProtoMessage()    {}
func (*MatchPredicate_MatchSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab4dd2506d34575, []int{1, 0}
}

func (m *MatchPredicate_MatchSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchPredicate_MatchSet.Unmarshal(m, b)
}
func (m *MatchPredicate_MatchSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchPredicate_MatchSet.Marshal(b, m, deterministic)
}
func (m *MatchPredicate_MatchSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchPredicate_MatchSet.Merge(m, src)
}
func (m *MatchPredicate_MatchSet) XXX_Size() int {
	return xxx_messageInfo_MatchPredicate_MatchSet.Size(m)
}
func (m *MatchPredicate_MatchSet) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchPredicate_MatchSet.DiscardUnknown(m)
}

var xxx_messageInfo_MatchPredicate_MatchSet proto.InternalMessageInfo

func (m *MatchPredicate_MatchSet) GetRules() []*MatchPredicate {
	if m != nil {
		return m.Rules
	}
	return nil
}

// HTTP headers match configuration.
type HttpHeadersMatch struct {
	// HTTP headers to match.
	Headers              []*route.HeaderMatcher `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HttpHeadersMatch) Reset()         { *m = HttpHeadersMatch{} }
func (m *HttpHeadersMatch) String() string { return proto.CompactTextString(m) }
func (*HttpHeadersMatch) ProtoMessage()    {}
func (*HttpHeadersMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab4dd2506d34575, []int{2}
}

func (m *HttpHeadersMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpHeadersMatch.Unmarshal(m, b)
}
func (m *HttpHeadersMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpHeadersMatch.Marshal(b, m, deterministic)
}
func (m *HttpHeadersMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpHeadersMatch.Merge(m, src)
}
func (m *HttpHeadersMatch) XXX_Size() int {
	return xxx_messageInfo_HttpHeadersMatch.Size(m)
}
func (m *HttpHeadersMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpHeadersMatch.DiscardUnknown(m)
}

var xxx_messageInfo_HttpHeadersMatch proto.InternalMessageInfo

func (m *HttpHeadersMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

// Tap output configuration.
type OutputConfig struct {
	// Output sinks for tap data. Currently a single sink is allowed in the list. Once multiple
	// sink types are supported this constraint will be relaxed.
	Sinks []*OutputSink `protobuf:"bytes,1,rep,name=sinks,proto3" json:"sinks,omitempty"`
	// For buffered tapping, the maximum amount of received body that will be buffered prior to
	// truncation. If truncation occurs, the :ref:`truncated
	// <envoy_api_field_data.tap.v3alpha.Body.truncated>` field will be set. If not specified, the
	// default is 1KiB.
	MaxBufferedRxBytes *types.UInt32Value `protobuf:"bytes,2,opt,name=max_buffered_rx_bytes,json=maxBufferedRxBytes,proto3" json:"max_buffered_rx_bytes,omitempty"`
	// For buffered tapping, the maximum amount of transmitted body that will be buffered prior to
	// truncation. If truncation occurs, the :ref:`truncated
	// <envoy_api_field_data.tap.v3alpha.Body.truncated>` field will be set. If not specified, the
	// default is 1KiB.
	MaxBufferedTxBytes *types.UInt32Value `protobuf:"bytes,3,opt,name=max_buffered_tx_bytes,json=maxBufferedTxBytes,proto3" json:"max_buffered_tx_bytes,omitempty"`
	// Indicates whether taps produce a single buffered message per tap, or multiple streamed
	// messages per tap in the emitted :ref:`TraceWrapper
	// <envoy_api_msg_data.tap.v3alpha.TraceWrapper>` messages. Note that streamed tapping does not
	// mean that no buffering takes place. Buffering may be required if data is processed before a
	// match can be determined. See the HTTP tap filter :ref:`streaming
	// <config_http_filters_tap_streaming>` documentation for more information.
	Streaming            bool     `protobuf:"varint,4,opt,name=streaming,proto3" json:"streaming,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputConfig) Reset()         { *m = OutputConfig{} }
func (m *OutputConfig) String() string { return proto.CompactTextString(m) }
func (*OutputConfig) ProtoMessage()    {}
func (*OutputConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab4dd2506d34575, []int{3}
}

func (m *OutputConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputConfig.Unmarshal(m, b)
}
func (m *OutputConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputConfig.Marshal(b, m, deterministic)
}
func (m *OutputConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputConfig.Merge(m, src)
}
func (m *OutputConfig) XXX_Size() int {
	return xxx_messageInfo_OutputConfig.Size(m)
}
func (m *OutputConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OutputConfig proto.InternalMessageInfo

func (m *OutputConfig) GetSinks() []*OutputSink {
	if m != nil {
		return m.Sinks
	}
	return nil
}

func (m *OutputConfig) GetMaxBufferedRxBytes() *types.UInt32Value {
	if m != nil {
		return m.MaxBufferedRxBytes
	}
	return nil
}

func (m *OutputConfig) GetMaxBufferedTxBytes() *types.UInt32Value {
	if m != nil {
		return m.MaxBufferedTxBytes
	}
	return nil
}

func (m *OutputConfig) GetStreaming() bool {
	if m != nil {
		return m.Streaming
	}
	return false
}

// Tap output sink configuration.
type OutputSink struct {
	// Sink output format.
	Format OutputSink_Format `protobuf:"varint,1,opt,name=format,proto3,enum=envoy.service.tap.v3alpha.OutputSink_Format" json:"format,omitempty"`
	// Types that are valid to be assigned to OutputSinkType:
	//	*OutputSink_StreamingAdmin
	//	*OutputSink_FilePerTap
	//	*OutputSink_StreamingGrpc
	OutputSinkType       isOutputSink_OutputSinkType `protobuf_oneof:"output_sink_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *OutputSink) Reset()         { *m = OutputSink{} }
func (m *OutputSink) String() string { return proto.CompactTextString(m) }
func (*OutputSink) ProtoMessage()    {}
func (*OutputSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab4dd2506d34575, []int{4}
}

func (m *OutputSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputSink.Unmarshal(m, b)
}
func (m *OutputSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputSink.Marshal(b, m, deterministic)
}
func (m *OutputSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputSink.Merge(m, src)
}
func (m *OutputSink) XXX_Size() int {
	return xxx_messageInfo_OutputSink.Size(m)
}
func (m *OutputSink) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputSink.DiscardUnknown(m)
}

var xxx_messageInfo_OutputSink proto.InternalMessageInfo

func (m *OutputSink) GetFormat() OutputSink_Format {
	if m != nil {
		return m.Format
	}
	return OutputSink_JSON_BODY_AS_BYTES
}

type isOutputSink_OutputSinkType interface {
	isOutputSink_OutputSinkType()
}

type OutputSink_StreamingAdmin struct {
	StreamingAdmin *StreamingAdminSink `protobuf:"bytes,2,opt,name=streaming_admin,json=streamingAdmin,proto3,oneof"`
}

type OutputSink_FilePerTap struct {
	FilePerTap *FilePerTapSink `protobuf:"bytes,3,opt,name=file_per_tap,json=filePerTap,proto3,oneof"`
}

type OutputSink_StreamingGrpc struct {
	StreamingGrpc *StreamingGrpcSink `protobuf:"bytes,4,opt,name=streaming_grpc,json=streamingGrpc,proto3,oneof"`
}

func (*OutputSink_StreamingAdmin) isOutputSink_OutputSinkType() {}

func (*OutputSink_FilePerTap) isOutputSink_OutputSinkType() {}

func (*OutputSink_StreamingGrpc) isOutputSink_OutputSinkType() {}

func (m *OutputSink) GetOutputSinkType() isOutputSink_OutputSinkType {
	if m != nil {
		return m.OutputSinkType
	}
	return nil
}

func (m *OutputSink) GetStreamingAdmin() *StreamingAdminSink {
	if x, ok := m.GetOutputSinkType().(*OutputSink_StreamingAdmin); ok {
		return x.StreamingAdmin
	}
	return nil
}

func (m *OutputSink) GetFilePerTap() *FilePerTapSink {
	if x, ok := m.GetOutputSinkType().(*OutputSink_FilePerTap); ok {
		return x.FilePerTap
	}
	return nil
}

func (m *OutputSink) GetStreamingGrpc() *StreamingGrpcSink {
	if x, ok := m.GetOutputSinkType().(*OutputSink_StreamingGrpc); ok {
		return x.StreamingGrpc
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OutputSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OutputSink_StreamingAdmin)(nil),
		(*OutputSink_FilePerTap)(nil),
		(*OutputSink_StreamingGrpc)(nil),
	}
}

// Streaming admin sink configuration.
type StreamingAdminSink struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingAdminSink) Reset()         { *m = StreamingAdminSink{} }
func (m *StreamingAdminSink) String() string { return proto.CompactTextString(m) }
func (*StreamingAdminSink) ProtoMessage()    {}
func (*StreamingAdminSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab4dd2506d34575, []int{5}
}

func (m *StreamingAdminSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingAdminSink.Unmarshal(m, b)
}
func (m *StreamingAdminSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingAdminSink.Marshal(b, m, deterministic)
}
func (m *StreamingAdminSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingAdminSink.Merge(m, src)
}
func (m *StreamingAdminSink) XXX_Size() int {
	return xxx_messageInfo_StreamingAdminSink.Size(m)
}
func (m *StreamingAdminSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingAdminSink.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingAdminSink proto.InternalMessageInfo

// The file per tap sink outputs a discrete file for every tapped stream.
type FilePerTapSink struct {
	// Path prefix. The output file will be of the form <path_prefix>_<id>.pb, where <id> is an
	// identifier distinguishing the recorded trace for stream instances (the Envoy
	// connection ID, HTTP stream ID, etc.).
	PathPrefix           string   `protobuf:"bytes,1,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilePerTapSink) Reset()         { *m = FilePerTapSink{} }
func (m *FilePerTapSink) String() string { return proto.CompactTextString(m) }
func (*FilePerTapSink) ProtoMessage()    {}
func (*FilePerTapSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab4dd2506d34575, []int{6}
}

func (m *FilePerTapSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilePerTapSink.Unmarshal(m, b)
}
func (m *FilePerTapSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilePerTapSink.Marshal(b, m, deterministic)
}
func (m *FilePerTapSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilePerTapSink.Merge(m, src)
}
func (m *FilePerTapSink) XXX_Size() int {
	return xxx_messageInfo_FilePerTapSink.Size(m)
}
func (m *FilePerTapSink) XXX_DiscardUnknown() {
	xxx_messageInfo_FilePerTapSink.DiscardUnknown(m)
}

var xxx_messageInfo_FilePerTapSink proto.InternalMessageInfo

func (m *FilePerTapSink) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

// [#not-implemented-hide:] Streaming gRPC sink configuration sends the taps to an external gRPC
// server.
type StreamingGrpcSink struct {
	// Opaque identifier, that will be sent back to the streaming grpc server.
	TapId string `protobuf:"bytes,1,opt,name=tap_id,json=tapId,proto3" json:"tap_id,omitempty"`
	// The gRPC server that hosts the Tap Sink Service.
	GrpcService          *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StreamingGrpcSink) Reset()         { *m = StreamingGrpcSink{} }
func (m *StreamingGrpcSink) String() string { return proto.CompactTextString(m) }
func (*StreamingGrpcSink) ProtoMessage()    {}
func (*StreamingGrpcSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab4dd2506d34575, []int{7}
}

func (m *StreamingGrpcSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingGrpcSink.Unmarshal(m, b)
}
func (m *StreamingGrpcSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingGrpcSink.Marshal(b, m, deterministic)
}
func (m *StreamingGrpcSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingGrpcSink.Merge(m, src)
}
func (m *StreamingGrpcSink) XXX_Size() int {
	return xxx_messageInfo_StreamingGrpcSink.Size(m)
}
func (m *StreamingGrpcSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingGrpcSink.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingGrpcSink proto.InternalMessageInfo

func (m *StreamingGrpcSink) GetTapId() string {
	if m != nil {
		return m.TapId
	}
	return ""
}

func (m *StreamingGrpcSink) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.service.tap.v3alpha.OutputSink_Format", OutputSink_Format_name, OutputSink_Format_value)
	proto.RegisterType((*TapConfig)(nil), "envoy.service.tap.v3alpha.TapConfig")
	proto.RegisterType((*MatchPredicate)(nil), "envoy.service.tap.v3alpha.MatchPredicate")
	proto.RegisterType((*MatchPredicate_MatchSet)(nil), "envoy.service.tap.v3alpha.MatchPredicate.MatchSet")
	proto.RegisterType((*HttpHeadersMatch)(nil), "envoy.service.tap.v3alpha.HttpHeadersMatch")
	proto.RegisterType((*OutputConfig)(nil), "envoy.service.tap.v3alpha.OutputConfig")
	proto.RegisterType((*OutputSink)(nil), "envoy.service.tap.v3alpha.OutputSink")
	proto.RegisterType((*StreamingAdminSink)(nil), "envoy.service.tap.v3alpha.StreamingAdminSink")
	proto.RegisterType((*FilePerTapSink)(nil), "envoy.service.tap.v3alpha.FilePerTapSink")
	proto.RegisterType((*StreamingGrpcSink)(nil), "envoy.service.tap.v3alpha.StreamingGrpcSink")
}

func init() {
	proto.RegisterFile("envoy/service/tap/v3alpha/common.proto", fileDescriptor_8ab4dd2506d34575)
}

var fileDescriptor_8ab4dd2506d34575 = []byte{
	// 992 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xc1, 0x53, 0xdb, 0xc6,
	0x17, 0xc7, 0x91, 0x0c, 0xc6, 0x3c, 0x3b, 0xfc, 0xf4, 0xdb, 0x36, 0x85, 0x50, 0xda, 0x21, 0xce,
	0x94, 0x90, 0x69, 0x2b, 0x77, 0xe0, 0xd2, 0xe9, 0xa9, 0x28, 0x18, 0xec, 0x4e, 0xc0, 0x8e, 0x2c,
	0x32, 0xe1, 0xb4, 0xb3, 0xb6, 0xd6, 0x46, 0x89, 0xac, 0xdd, 0xae, 0xd6, 0xd4, 0x9e, 0xe9, 0xf4,
	0xd0, 0x63, 0x0f, 0x3d, 0xe4, 0xd2, 0xff, 0xa7, 0xff, 0x53, 0x67, 0x3a, 0x9c, 0x3a, 0xda, 0x5d,
	0x83, 0x0c, 0x85, 0x42, 0x72, 0xf1, 0x58, 0x6f, 0xdf, 0xf7, 0xf3, 0xf6, 0x3d, 0xed, 0x7b, 0x2b,
	0xd8, 0xa4, 0xc9, 0x19, 0x9b, 0xd4, 0x52, 0x2a, 0xce, 0xa2, 0x1e, 0xad, 0x49, 0xc2, 0x6b, 0x67,
	0x3b, 0x24, 0xe6, 0xa7, 0xa4, 0xd6, 0x63, 0xc3, 0x21, 0x4b, 0x5c, 0x2e, 0x98, 0x64, 0xe8, 0x91,
	0xf2, 0x73, 0x8d, 0x9f, 0x2b, 0x09, 0x77, 0x8d, 0xdf, 0xda, 0x63, 0x8d, 0x20, 0x3c, 0xca, 0x49,
	0x05, 0xad, 0x75, 0x49, 0x4a, 0xb5, 0x7a, 0xed, 0xd9, 0x0d, 0x2e, 0x03, 0xc1, 0x7b, 0x78, 0xca,
	0xd4, 0xae, 0x4f, 0xae, 0xbb, 0x0a, 0x36, 0x92, 0x54, 0xff, 0x1a, 0xa7, 0xcf, 0x07, 0x8c, 0x0d,
	0x62, 0x5a, 0x53, 0x4f, 0xdd, 0x51, 0xbf, 0xf6, 0x93, 0x20, 0x9c, 0x53, 0x91, 0x9a, 0xf5, 0x95,
	0x33, 0x12, 0x47, 0x21, 0x91, 0xb4, 0x36, 0xfd, 0xa3, 0x17, 0xaa, 0xbf, 0xdb, 0xb0, 0x14, 0x10,
	0xfe, 0x9c, 0x25, 0xfd, 0x68, 0x80, 0x5e, 0x41, 0x65, 0x48, 0x64, 0xef, 0x14, 0xf7, 0xd4, 0xf3,
	0xaa, 0xb5, 0x61, 0x6d, 0x95, 0xb7, 0x9f, 0xb9, 0x37, 0xe6, 0xea, 0x1e, 0x66, 0xee, 0x6d, 0x41,
	0xc3, 0xa8, 0x47, 0x24, 0xf5, 0x4a, 0xe7, 0xde, 0xc2, 0x6f, 0x96, 0xed, 0x58, 0x7e, 0x59, 0x81,
	0x2e, 0xb8, 0x0f, 0xd8, 0x48, 0xf2, 0x91, 0x9c, 0x82, 0x6d, 0x05, 0x7e, 0x7a, 0x0b, 0xb8, 0xa5,
	0xfc, 0xb5, 0x3e, 0x87, 0xad, 0xb0, 0x9c, 0x1d, 0xbd, 0x84, 0xb2, 0x24, 0x1c, 0xd3, 0x84, 0x74,
	0x63, 0x1a, 0xae, 0x16, 0x14, 0xf5, 0x1b, 0x43, 0x25, 0x3c, 0xba, 0xa0, 0x65, 0xc5, 0x75, 0xfd,
	0x51, 0x22, 0xa3, 0x21, 0xdd, 0x17, 0xa4, 0x27, 0x23, 0x96, 0x90, 0xb8, 0x4d, 0x45, 0x8f, 0x26,
	0xd2, 0x07, 0x49, 0x78, 0x5d, 0x33, 0xaa, 0x7f, 0x16, 0x61, 0x79, 0x36, 0x29, 0xd4, 0x82, 0x12,
	0x13, 0x58, 0xe5, 0x63, 0x2a, 0xb2, 0x7d, 0xe7, 0x8a, 0xe8, 0xc7, 0x0e, 0x95, 0x8d, 0x39, 0x7f,
	0x91, 0x09, 0xf5, 0x84, 0x5e, 0xc2, 0x12, 0x49, 0x42, 0x43, 0xb4, 0x3f, 0x80, 0x58, 0x22, 0x49,
	0xa8, 0x91, 0x0d, 0x58, 0x4a, 0x98, 0x34, 0xc8, 0xc2, 0x3d, 0x5f, 0x5b, 0x46, 0x4a, 0x98, 0xd4,
	0xa4, 0xcd, 0x6c, 0x73, 0x13, 0x43, 0x9a, 0xdf, 0xb0, 0xb6, 0x4a, 0xde, 0xe2, 0xb9, 0x37, 0xff,
	0xc6, 0x2e, 0x59, 0x3a, 0xe2, 0x44, 0xfb, 0xbd, 0x81, 0xb5, 0x53, 0x29, 0x39, 0x16, 0xf4, 0xc7,
	0x11, 0x4d, 0x25, 0x3e, 0xa5, 0x24, 0xa4, 0x22, 0x35, 0xc2, 0x05, 0xb5, 0x85, 0x2f, 0x6f, 0xd9,
	0x42, 0x43, 0x4a, 0xde, 0xd0, 0x1a, 0x05, 0x6c, 0xcc, 0xf9, 0x2b, 0x19, 0xd0, 0xd7, 0xbc, 0xfc,
	0x12, 0x8a, 0xe1, 0xd3, 0x99, 0x58, 0x52, 0x90, 0x28, 0xbe, 0x0c, 0x56, 0x7c, 0x9f, 0x60, 0xab,
	0xb9, 0x60, 0x81, 0xe1, 0x5d, 0x8d, 0x96, 0x72, 0x96, 0xa4, 0xf4, 0x4a, 0x6a, 0x8b, 0x1f, 0x10,
	0x4d, 0x03, 0x67, 0x72, 0x4b, 0x60, 0x7d, 0x36, 0xda, 0x95, 0xe4, 0x4a, 0xef, 0x13, 0xee, 0x51,
	0x3e, 0xdc, 0x4c, 0x76, 0x6b, 0xc7, 0x50, 0x9a, 0x9e, 0x20, 0xd4, 0x84, 0x05, 0x31, 0x8a, 0x69,
	0xba, 0x6a, 0x6d, 0x14, 0xee, 0xdf, 0xe8, 0xef, 0x2c, 0xbb, 0x64, 0xfb, 0x9a, 0xe0, 0x95, 0x61,
	0x3e, 0xfb, 0x83, 0x0a, 0x7f, 0x7b, 0x56, 0x35, 0x00, 0xe7, 0xea, 0xa6, 0xd0, 0xf7, 0xb0, 0x68,
	0xea, 0x68, 0xa2, 0x6d, 0xfe, 0x4b, 0x9f, 0xea, 0x99, 0xa6, 0x75, 0x4a, 0x46, 0x85, 0x3f, 0x95,
	0x55, 0xff, 0xb0, 0xa1, 0x92, 0x1f, 0x0b, 0xe8, 0x00, 0x16, 0xd2, 0x28, 0x79, 0x3b, 0x05, 0x7e,
	0xf1, 0x9f, 0xe3, 0xa4, 0x13, 0x25, 0x6f, 0x3d, 0x38, 0xf7, 0x16, 0xdf, 0x59, 0xf3, 0x25, 0xcb,
	0xb1, 0x7c, 0xad, 0x47, 0x2d, 0x78, 0x38, 0x24, 0x63, 0xdc, 0x1d, 0xf5, 0xfb, 0x54, 0xd0, 0x10,
	0x8b, 0x31, 0xee, 0x4e, 0x24, 0x4d, 0x4d, 0x73, 0xae, 0xbb, 0x7a, 0xbc, 0xba, 0xd3, 0xf1, 0xea,
	0x1e, 0x37, 0x13, 0xb9, 0xb3, 0xfd, 0x8a, 0xc4, 0x23, 0xea, 0xa3, 0x21, 0x19, 0x7b, 0x46, 0xe9,
	0x8f, 0xbd, 0x4c, 0x77, 0x0d, 0x28, 0xa7, 0xc0, 0xc2, 0x3d, 0x81, 0x81, 0x01, 0xae, 0xc3, 0x52,
	0x2a, 0x05, 0x25, 0xc3, 0x28, 0x19, 0xe8, 0xae, 0xf4, 0x2f, 0x0d, 0xd5, 0xbf, 0x0a, 0x00, 0x97,
	0x19, 0xa2, 0x23, 0x28, 0xf6, 0x99, 0x18, 0x12, 0xa9, 0xc6, 0xd5, 0xf2, 0xf6, 0x57, 0x77, 0x2a,
	0x8c, 0xbb, 0xaf, 0x34, 0xea, 0xd5, 0xfe, 0xaa, 0x86, 0xad, 0xa1, 0xa0, 0xd7, 0xf0, 0xbf, 0x8b,
	0x58, 0x98, 0x84, 0xc3, 0x28, 0x31, 0x85, 0xf9, 0xfa, 0x16, 0x70, 0x67, 0xaa, 0xd8, 0xcd, 0x04,
	0x59, 0x80, 0xc6, 0x9c, 0xbf, 0x9c, 0xce, 0x58, 0xd1, 0x21, 0x54, 0xfa, 0x51, 0x4c, 0x31, 0xa7,
	0x02, 0x4b, 0xc2, 0xef, 0x30, 0xb9, 0xf6, 0xa3, 0x98, 0xb6, 0xa9, 0x08, 0x08, 0x37, 0x48, 0xe8,
	0x5f, 0x58, 0xd0, 0x31, 0x5c, 0x06, 0xc0, 0xd9, 0x5d, 0xaa, 0x4a, 0x55, 0xbe, 0xb5, 0x00, 0x17,
	0xfb, 0x3c, 0x10, 0xbc, 0x67, 0x98, 0x0f, 0xd2, 0xbc, 0xb1, 0xfa, 0x0b, 0x14, 0x75, 0x6d, 0xd0,
	0x27, 0x80, 0x7e, 0xe8, 0xb4, 0x8e, 0xb0, 0xd7, 0xda, 0x3b, 0xc1, 0xbb, 0x1d, 0xec, 0x9d, 0x04,
	0xf5, 0x8e, 0x33, 0x87, 0x56, 0xe0, 0xa3, 0x19, 0x7b, 0x27, 0xf0, 0x9b, 0x47, 0x07, 0x8e, 0x85,
	0x1c, 0xa8, 0xb4, 0xfd, 0x56, 0xd0, 0xc2, 0x5e, 0xf3, 0x68, 0xd7, 0x3f, 0x71, 0x6c, 0xf4, 0x18,
	0x3e, 0xcb, 0x5b, 0xf0, 0x8b, 0xfa, 0xd1, 0x41, 0xd0, 0xc0, 0x7b, 0xf5, 0x17, 0xcd, 0xc3, 0x66,
	0x50, 0xdf, 0x73, 0x0a, 0x68, 0x19, 0x40, 0xbb, 0x04, 0xf5, 0xd7, 0x81, 0x33, 0xef, 0xad, 0x80,
	0x63, 0xae, 0xcf, 0xec, 0xb8, 0x62, 0x39, 0xe1, 0xa6, 0xcf, 0x3e, 0x06, 0x74, 0xbd, 0xcc, 0xd5,
	0xef, 0x60, 0x79, 0xb6, 0x4a, 0x68, 0x0b, 0xca, 0x9c, 0xc8, 0x53, 0xcc, 0x05, 0xed, 0x47, 0x63,
	0x75, 0x2a, 0x96, 0xd4, 0x54, 0x17, 0xf6, 0x86, 0xe5, 0x43, 0xb6, 0xd6, 0x56, 0x4b, 0xd5, 0x9f,
	0xe1, 0xff, 0xd7, 0x0a, 0x82, 0x1e, 0x42, 0x31, 0xbb, 0x66, 0xa3, 0x50, 0x2b, 0xfd, 0x05, 0x49,
	0x78, 0x33, 0x44, 0x6d, 0xa8, 0xe4, 0xbf, 0x57, 0xcc, 0x99, 0x78, 0x72, 0xd3, 0xf5, 0xab, 0x70,
	0xda, 0x35, 0xff, 0x9d, 0x30, 0xc8, 0x99, 0xbf, 0x85, 0xa7, 0x11, 0xd3, 0x7a, 0x2e, 0xd8, 0x78,
	0x72, 0xf3, 0x6b, 0xf3, 0xca, 0xcf, 0xd5, 0xd7, 0x58, 0x3b, 0x6b, 0xa0, 0xb6, 0xd5, 0x2d, 0xaa,
	0x4e, 0xda, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0xda, 0x94, 0xb0, 0xd7, 0xbf, 0x09, 0x00, 0x00,
}
