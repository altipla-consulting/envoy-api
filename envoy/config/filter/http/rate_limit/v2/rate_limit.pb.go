// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/rate_limit/v2/rate_limit.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2 "github.com/altipla-consulting/envoy-api/envoy/config/ratelimit/v2"
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

type RateLimit struct {
	// The rate limit domain to use when calling the rate limit service.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Specifies the rate limit configurations to be applied with the same
	// stage number. If not set, the default stage number is 0.
	//
	// .. note::
	//
	//  The filter supports a range of 0 - 10 inclusively for stage numbers.
	Stage uint32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	// The type of requests the filter should apply to. The supported
	// types are *internal*, *external* or *both*. A request is considered internal if
	// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>` is set to true. If
	// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>` is not set or false, a
	// request is considered external. The filter defaults to *both*, and it will apply to all request
	// types.
	RequestType string `protobuf:"bytes,3,opt,name=request_type,json=requestType,proto3" json:"request_type,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *types.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny bool `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	// Specifies whether a `RESOURCE_EXHAUSTED` gRPC code must be returned instead
	// of the default `UNAVAILABLE` gRPC code for a rate limited gRPC call. The
	// HTTP code will be 200 for a gRPC response.
	RateLimitedAsResourceExhausted bool `protobuf:"varint,6,opt,name=rate_limited_as_resource_exhausted,json=rateLimitedAsResourceExhausted,proto3" json:"rate_limited_as_resource_exhausted,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	// [#comment:TODO(ramaraochavali): Make this required as part of cleanup of deprecated ratelimit
	// service config in bootstrap.]
	RateLimitService     *v2.RateLimitServiceConfig `protobuf:"bytes,7,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_rate_limit_cd9a675ce61b3eb9, []int{0}
}
func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (dst *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(dst, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetStage() uint32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RateLimit) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *RateLimit) GetTimeout() *types.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitedAsResourceExhausted() bool {
	if m != nil {
		return m.RateLimitedAsResourceExhausted
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v2.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.http.rate_limit.v2.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/rate_limit/v2/rate_limit.proto", fileDescriptor_rate_limit_cd9a675ce61b3eb9)
}

var fileDescriptor_rate_limit_cd9a675ce61b3eb9 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0xbf, 0x6e, 0xd4, 0x30,
	0x18, 0x57, 0x72, 0xe9, 0x95, 0x73, 0x41, 0x80, 0x85, 0x84, 0xe9, 0xd0, 0xa6, 0x45, 0x42, 0x51,
	0x07, 0x5b, 0x84, 0x01, 0x31, 0x72, 0x94, 0x05, 0xc1, 0x62, 0x98, 0x58, 0x2c, 0xf7, 0xfc, 0x25,
	0xb5, 0x94, 0xc4, 0xc1, 0x76, 0x22, 0xf2, 0x02, 0x3c, 0x03, 0xcf, 0xc2, 0xc4, 0x9b, 0x30, 0xf3,
	0x16, 0x28, 0x89, 0xd3, 0x6b, 0xd9, 0x3e, 0xff, 0xfe, 0xd9, 0xdf, 0xcf, 0xe8, 0x35, 0x34, 0xbd,
	0x19, 0xd8, 0xce, 0x34, 0x85, 0x2e, 0x59, 0xa1, 0x2b, 0x0f, 0x96, 0x5d, 0x7b, 0xdf, 0x32, 0x2b,
	0x3d, 0x88, 0x4a, 0xd7, 0xda, 0xb3, 0x3e, 0xbf, 0x75, 0xa2, 0xad, 0x35, 0xde, 0xe0, 0x17, 0x93,
	0x91, 0xce, 0x46, 0x3a, 0x1b, 0xe9, 0x68, 0xa4, 0xb7, 0xa4, 0x7d, 0x7e, 0xfc, 0xfc, 0xce, 0x05,
	0x23, 0xb7, 0xcf, 0xac, 0xdc, 0x1c, 0x76, 0x7c, 0x52, 0x1a, 0x53, 0x56, 0xc0, 0xa6, 0xd3, 0x55,
	0x57, 0x30, 0xd5, 0x59, 0xe9, 0xb5, 0x69, 0x02, 0xff, 0xb4, 0x97, 0x95, 0x56, 0xd2, 0x03, 0x5b,
	0x86, 0x40, 0x3c, 0x29, 0x4d, 0x69, 0xa6, 0x91, 0x8d, 0xd3, 0x8c, 0x9e, 0xff, 0x58, 0xa1, 0x0d,
	0x97, 0x1e, 0x3e, 0x8e, 0x37, 0xe1, 0x33, 0xb4, 0x56, 0xa6, 0x96, 0xba, 0x21, 0x51, 0x1a, 0x65,
	0x9b, 0xed, 0xe6, 0xd7, 0xdf, 0xdf, 0xab, 0xc4, 0xc6, 0x69, 0xc4, 0x03, 0x81, 0x4f, 0xd1, 0x81,
	0xf3, 0xb2, 0x04, 0x12, 0xa7, 0x51, 0xf6, 0x20, 0x28, 0x2e, 0x62, 0x82, 0xf8, 0x8c, 0xe3, 0x33,
	0x74, 0xdf, 0xc2, 0xb7, 0x0e, 0x9c, 0x17, 0x7e, 0x68, 0x81, 0xac, 0xc6, 0x24, 0x7e, 0x14, 0xb0,
	0x2f, 0x43, 0x0b, 0xf8, 0x0d, 0x3a, 0xf4, 0xba, 0x06, 0xd3, 0x79, 0x92, 0xa4, 0x51, 0x76, 0x94,
	0x3f, 0xa3, 0xf3, 0x56, 0x74, 0xd9, 0x8a, 0x5e, 0x86, 0xad, 0xb6, 0xc9, 0xcf, 0x3f, 0xa7, 0x11,
	0x5f, 0xf4, 0xf8, 0x02, 0x3d, 0x2e, 0xa4, 0xae, 0x3a, 0x0b, 0xa2, 0x36, 0x0a, 0x84, 0x82, 0x66,
	0x20, 0x07, 0x69, 0x94, 0xdd, 0xe3, 0x0f, 0x03, 0xf1, 0xc9, 0x28, 0xb8, 0x84, 0x66, 0xc0, 0x1f,
	0xd0, 0xf9, 0xbe, 0x60, 0x50, 0x42, 0x3a, 0x61, 0xc1, 0x99, 0xce, 0xee, 0x40, 0xc0, 0xf7, 0x6b,
	0xd9, 0x39, 0x0f, 0x8a, 0xac, 0x27, 0xf3, 0x89, 0x5d, 0x4a, 0x00, 0xf5, 0xd6, 0xf1, 0x20, 0x7b,
	0xbf, 0xa8, 0xb0, 0x40, 0x78, 0x9f, 0x25, 0x1c, 0xd8, 0x5e, 0xef, 0x80, 0x1c, 0x4e, 0xaf, 0x7f,
	0x49, 0xef, 0x7c, 0xf0, 0xcd, 0xc7, 0xd1, 0x3e, 0xa7, 0x37, 0xdd, 0x7e, 0x9e, 0x2d, 0xef, 0x26,
	0x0d, 0x7f, 0x64, 0xff, 0xc3, 0xb7, 0xc9, 0xd7, 0xb8, 0xcf, 0xaf, 0xd6, 0x53, 0x01, 0xaf, 0xfe,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xdd, 0x1b, 0x97, 0x6c, 0x02, 0x00, 0x00,
}
