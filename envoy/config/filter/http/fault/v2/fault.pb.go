// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/fault/v2/fault.proto

package envoy_config_filter_http_fault_v2

import (
	fmt "fmt"
	route "github.com/altipla-consulting/envoy-api/envoy/api/v2/route"
	v2 "github.com/altipla-consulting/envoy-api/envoy/config/filter/fault/v2"
	_type "github.com/altipla-consulting/envoy-api/envoy/type"
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

type FaultAbort struct {
	// Types that are valid to be assigned to ErrorType:
	//	*FaultAbort_HttpStatus
	ErrorType isFaultAbort_ErrorType `protobuf_oneof:"error_type"`
	// The percentage of requests/operations/connections that will be aborted with the error code
	// provided.
	Percentage           *_type.FractionalPercent `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FaultAbort) Reset()         { *m = FaultAbort{} }
func (m *FaultAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort) ProtoMessage()    {}
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{0}
}

func (m *FaultAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultAbort.Unmarshal(m, b)
}
func (m *FaultAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultAbort.Marshal(b, m, deterministic)
}
func (m *FaultAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort.Merge(m, src)
}
func (m *FaultAbort) XXX_Size() int {
	return xxx_messageInfo_FaultAbort.Size(m)
}
func (m *FaultAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort proto.InternalMessageInfo

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
}

type FaultAbort_HttpStatus struct {
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := m.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *FaultAbort) GetPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultAbort) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultAbort_HttpStatus)(nil),
	}
}

// [#next-free-field: 14]
type HTTPFault struct {
	// If specified, the filter will inject delays based on the values in the
	// object.
	Delay *v2.FaultDelay `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	// If specified, the filter will abort requests based on the values in
	// the object. At least *abort* or *delay* must be specified.
	Abort *FaultAbort `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
	// Specifies the name of the (destination) upstream cluster that the
	// filter should match on. Fault injection will be restricted to requests
	// bound to the specific upstream cluster.
	UpstreamCluster string `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	// Specifies a set of headers that the filter should match on. The fault
	// injection filter can be applied selectively to requests that match a set of
	// headers specified in the fault filter config. The chances of actual fault
	// injection further depend on the value of the :ref:`percentage
	// <envoy_api_field_config.filter.http.fault.v2.FaultAbort.percentage>` field.
	// The filter will check the request's headers against all the specified
	// headers in the filter config. A match will happen if all the headers in the
	// config are present in the request with the same values (or based on
	// presence if the *value* field is not in the config).
	Headers []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	// Faults are injected for the specified list of downstream hosts. If this
	// setting is not set, faults are injected for all downstream nodes.
	// Downstream node name is taken from :ref:`the HTTP
	// x-envoy-downstream-service-node
	// <config_http_conn_man_headers_downstream-service-node>` header and compared
	// against downstream_nodes list.
	DownstreamNodes []string `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes,proto3" json:"downstream_nodes,omitempty"`
	// The maximum number of faults that can be active at a single time via the configured fault
	// filter. Note that because this setting can be overridden at the route level, it's possible
	// for the number of active faults to be greater than this value (if injected via a different
	// route). If not specified, defaults to unlimited. This setting can be overridden via
	// `runtime <config_http_filters_fault_injection_runtime>` and any faults that are not injected
	// due to overflow will be indicated via the `faults_overflow
	// <config_http_filters_fault_injection_stats>` stat.
	//
	// .. attention::
	//   Like other :ref:`circuit breakers <arch_overview_circuit_break>` in Envoy, this is a fuzzy
	//   limit. It's possible for the number of active faults to rise slightly above the configured
	//   amount due to the implementation details.
	MaxActiveFaults *types.UInt32Value `protobuf:"bytes,6,opt,name=max_active_faults,json=maxActiveFaults,proto3" json:"max_active_faults,omitempty"`
	// The response rate limit to be applied to the response body of the stream. When configured,
	// the percentage can be overridden by the :ref:`fault.http.rate_limit.response_percent
	// <config_http_filters_fault_injection_runtime>` runtime key.
	//
	// .. attention::
	//  This is a per-stream limit versus a connection level limit. This means that concurrent streams
	//  will each get an independent limit.
	ResponseRateLimit *v2.FaultRateLimit `protobuf:"bytes,7,opt,name=response_rate_limit,json=responseRateLimit,proto3" json:"response_rate_limit,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.delay.fixed_delay_percent
	DelayPercentRuntime string `protobuf:"bytes,8,opt,name=delay_percent_runtime,json=delayPercentRuntime,proto3" json:"delay_percent_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.abort.abort_percent
	AbortPercentRuntime string `protobuf:"bytes,9,opt,name=abort_percent_runtime,json=abortPercentRuntime,proto3" json:"abort_percent_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.delay.fixed_duration_ms
	DelayDurationRuntime string `protobuf:"bytes,10,opt,name=delay_duration_runtime,json=delayDurationRuntime,proto3" json:"delay_duration_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.abort.http_status
	AbortHttpStatusRuntime string `protobuf:"bytes,11,opt,name=abort_http_status_runtime,json=abortHttpStatusRuntime,proto3" json:"abort_http_status_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.max_active_faults
	MaxActiveFaultsRuntime string `protobuf:"bytes,12,opt,name=max_active_faults_runtime,json=maxActiveFaultsRuntime,proto3" json:"max_active_faults_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.rate_limit.response_percent
	ResponseRateLimitPercentRuntime string   `protobuf:"bytes,13,opt,name=response_rate_limit_percent_runtime,json=responseRateLimitPercentRuntime,proto3" json:"response_rate_limit_percent_runtime,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *HTTPFault) Reset()         { *m = HTTPFault{} }
func (m *HTTPFault) String() string { return proto.CompactTextString(m) }
func (*HTTPFault) ProtoMessage()    {}
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{1}
}

func (m *HTTPFault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPFault.Unmarshal(m, b)
}
func (m *HTTPFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPFault.Marshal(b, m, deterministic)
}
func (m *HTTPFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFault.Merge(m, src)
}
func (m *HTTPFault) XXX_Size() int {
	return xxx_messageInfo_HTTPFault.Size(m)
}
func (m *HTTPFault) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFault.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFault proto.InternalMessageInfo

func (m *HTTPFault) GetDelay() *v2.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFault) GetAbort() *FaultAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *HTTPFault) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *HTTPFault) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPFault) GetDownstreamNodes() []string {
	if m != nil {
		return m.DownstreamNodes
	}
	return nil
}

func (m *HTTPFault) GetMaxActiveFaults() *types.UInt32Value {
	if m != nil {
		return m.MaxActiveFaults
	}
	return nil
}

func (m *HTTPFault) GetResponseRateLimit() *v2.FaultRateLimit {
	if m != nil {
		return m.ResponseRateLimit
	}
	return nil
}

func (m *HTTPFault) GetDelayPercentRuntime() string {
	if m != nil {
		return m.DelayPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortPercentRuntime() string {
	if m != nil {
		return m.AbortPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetDelayDurationRuntime() string {
	if m != nil {
		return m.DelayDurationRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortHttpStatusRuntime() string {
	if m != nil {
		return m.AbortHttpStatusRuntime
	}
	return ""
}

func (m *HTTPFault) GetMaxActiveFaultsRuntime() string {
	if m != nil {
		return m.MaxActiveFaultsRuntime
	}
	return ""
}

func (m *HTTPFault) GetResponseRateLimitPercentRuntime() string {
	if m != nil {
		return m.ResponseRateLimitPercentRuntime
	}
	return ""
}

func init() {
	proto.RegisterType((*FaultAbort)(nil), "envoy.config.filter.http.fault.v2.FaultAbort")
	proto.RegisterType((*HTTPFault)(nil), "envoy.config.filter.http.fault.v2.HTTPFault")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/fault/v2/fault.proto", fileDescriptor_26070db6b6576d5c)
}

var fileDescriptor_26070db6b6576d5c = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x5f, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0xd3, 0xa6, 0xed, 0x9a, 0xaa, 0x8d, 0x0b, 0xc5, 0x44, 0x50, 0xd2, 0xf2, 0x62,
	0x10, 0x5d, 0x4b, 0x2e, 0x2f, 0x08, 0x01, 0xaa, 0x5b, 0x55, 0x01, 0x15, 0x14, 0x99, 0xc2, 0x13,
	0x92, 0xb5, 0x89, 0x27, 0x89, 0x25, 0xc7, 0x6b, 0xad, 0xd7, 0x69, 0x72, 0x0f, 0xae, 0xc1, 0x3d,
	0x38, 0x02, 0x47, 0x41, 0x3c, 0xa1, 0x9d, 0xb5, 0xd3, 0xd0, 0x44, 0x2a, 0x2f, 0x95, 0xbb, 0xf3,
	0xfd, 0xe6, 0xcf, 0x37, 0x13, 0x72, 0x04, 0xe9, 0x98, 0x4f, 0xdd, 0x1e, 0x4f, 0xfb, 0xf1, 0xc0,
	0xed, 0xc7, 0x89, 0x04, 0xe1, 0x0e, 0xa5, 0xcc, 0xdc, 0x3e, 0x2b, 0x12, 0xe9, 0x8e, 0x3d, 0xfd,
	0x41, 0x33, 0xc1, 0x25, 0xb7, 0x0e, 0x50, 0x4e, 0xb5, 0x9c, 0x6a, 0x39, 0x55, 0x72, 0xaa, 0x55,
	0x63, 0xaf, 0xb9, 0xaf, 0x33, 0xb2, 0x2c, 0x56, 0xb0, 0xe0, 0x85, 0x04, 0xfd, 0x57, 0xa7, 0x68,
	0x3a, 0xcb, 0x2a, 0x2e, 0x2b, 0xd6, 0xb4, 0xb5, 0x52, 0x4e, 0x33, 0x70, 0x33, 0x10, 0x3d, 0x48,
	0xab, 0xc8, 0xfe, 0x80, 0xf3, 0x41, 0x02, 0x2e, 0xfe, 0xd7, 0x2d, 0xfa, 0xee, 0x95, 0x60, 0x59,
	0x06, 0x22, 0x2f, 0xe3, 0x0f, 0xc6, 0x2c, 0x89, 0x23, 0x26, 0xc1, 0xad, 0x3e, 0x74, 0xe0, 0xf0,
	0xbb, 0x41, 0xc8, 0xb9, 0x2a, 0x71, 0xd2, 0xe5, 0x42, 0x5a, 0x94, 0x98, 0xaa, 0xf9, 0x30, 0x97,
	0x4c, 0x16, 0xb9, 0xbd, 0xd2, 0x32, 0x9c, 0x2d, 0xdf, 0xfc, 0xe3, 0x6f, 0x3c, 0xaf, 0xef, 0xfc,
	0x5a, 0x75, 0x7e, 0x1a, 0xed, 0x3b, 0x01, 0x51, 0x8a, 0xcf, 0x28, 0xb0, 0xde, 0x10, 0x52, 0x36,
	0xc2, 0x06, 0x60, 0xd7, 0x5a, 0x86, 0x63, 0x7a, 0x8f, 0xa9, 0xf6, 0x44, 0xb5, 0x49, 0xcf, 0x05,
	0xeb, 0xc9, 0x98, 0xa7, 0x2c, 0xe9, 0x68, 0x5d, 0x30, 0x07, 0xf8, 0x0d, 0x42, 0x40, 0x08, 0x2e,
	0x42, 0xa5, 0xb5, 0x6a, 0xbf, 0x7d, 0xe3, 0xc3, 0xea, 0x86, 0xb1, 0xb3, 0x72, 0xf8, 0xa3, 0x4e,
	0x36, 0xdb, 0x97, 0x97, 0x1d, 0x6c, 0xcd, 0x7a, 0x4b, 0xd6, 0x22, 0x48, 0xd8, 0xd4, 0x36, 0xb0,
	0x80, 0x43, 0x97, 0x99, 0x5e, 0xf9, 0x4d, 0x91, 0x39, 0x53, 0xfa, 0x40, 0x63, 0xd6, 0x29, 0x59,
	0x63, 0x6a, 0x3c, 0x9c, 0xc7, 0xf4, 0x8e, 0xe8, 0xad, 0x4b, 0xa3, 0xd7, 0x9e, 0x04, 0x9a, 0xb5,
	0x9e, 0x91, 0x9d, 0x22, 0xcb, 0xa5, 0x00, 0x36, 0x0a, 0x7b, 0x49, 0x91, 0x4b, 0x10, 0x38, 0xf0,
	0x66, 0xb0, 0x5d, 0xbd, 0x9f, 0xea, 0x67, 0xeb, 0x35, 0x59, 0x1f, 0x02, 0x8b, 0x40, 0xe4, 0xf6,
	0x6a, 0xab, 0xe6, 0x98, 0xde, 0x41, 0x59, 0x91, 0x65, 0xb1, 0x4a, 0xae, 0xb7, 0xdf, 0x46, 0xc9,
	0x47, 0x26, 0x7b, 0x43, 0x10, 0x41, 0x45, 0xa8, 0x3a, 0x11, 0xbf, 0x4a, 0xcb, 0x4a, 0x29, 0x8f,
	0x20, 0xb7, 0xd7, 0x5a, 0x35, 0x55, 0xe7, 0xfa, 0xfd, 0x93, 0x7a, 0xb6, 0xda, 0xa4, 0x31, 0x62,
	0x93, 0x50, 0x39, 0x3c, 0x86, 0x10, 0x7b, 0xcf, 0xed, 0x3a, 0xce, 0xf8, 0x88, 0xea, 0x8b, 0xa0,
	0xd5, 0x45, 0xd0, 0x2f, 0xef, 0x53, 0x79, 0xec, 0x7d, 0x65, 0x49, 0x01, 0xc1, 0xf6, 0x88, 0x4d,
	0x4e, 0x90, 0xc2, 0x39, 0x73, 0xeb, 0x1b, 0xd9, 0x15, 0x90, 0x67, 0x3c, 0xcd, 0x21, 0x14, 0x4c,
	0x42, 0x98, 0xc4, 0xa3, 0x58, 0xda, 0xeb, 0x98, 0xeb, 0xc5, 0x7f, 0xf8, 0x1d, 0x30, 0x09, 0x17,
	0x8a, 0x09, 0x1a, 0x55, 0xa2, 0xd9, 0x93, 0xe5, 0x91, 0xfb, 0xb8, 0x88, 0xb0, 0x5c, 0x7d, 0x28,
	0x8a, 0x54, 0xc6, 0x23, 0xb0, 0x37, 0xd0, 0xbf, 0x5d, 0x0c, 0x56, 0xf7, 0xa1, 0x43, 0x8a, 0x41,
	0xdf, 0x17, 0x98, 0x4d, 0xcd, 0x60, 0xf0, 0x06, 0xf3, 0x92, 0xec, 0xe9, 0x3a, 0x51, 0x21, 0x98,
	0xba, 0xba, 0x19, 0x44, 0x10, 0xba, 0x87, 0xd1, 0xb3, 0x32, 0x58, 0x51, 0xaf, 0xc8, 0x43, 0x5d,
	0x69, 0xee, 0xf2, 0x67, 0xa0, 0x89, 0xe0, 0x1e, 0x0a, 0xda, 0xb3, 0xbb, 0x9f, 0x43, 0x17, 0x16,
	0x30, 0x43, 0xef, 0x6a, 0xf4, 0x86, 0xd5, 0x15, 0x7a, 0x41, 0x9e, 0x2e, 0x71, 0x7c, 0x61, 0xda,
	0x2d, 0x4c, 0xf2, 0x64, 0xc1, 0xd3, 0x7f, 0x27, 0xf7, 0xdf, 0x11, 0x37, 0xe6, 0x7a, 0x4d, 0x99,
	0xe0, 0x93, 0xe9, 0xed, 0x17, 0xee, 0xeb, 0x9f, 0x7d, 0x47, 0x9d, 0x47, 0xc7, 0xe8, 0xd6, 0xf1,
	0x4e, 0x8e, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x42, 0x1b, 0x0e, 0xff, 0x04, 0x00, 0x00,
}
