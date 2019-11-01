// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/rate_limit/v3alpha/rate_limit.proto

package envoy_config_filter_network_rate_limit_v3alpha

import (
	fmt "fmt"
	ratelimit "github.com/altipla-consulting/envoy-api/envoy/api/v3alpha/ratelimit"
	v3alpha "github.com/altipla-consulting/envoy-api/envoy/config/ratelimit/v3alpha"
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

// [#next-free-field: 7]
type RateLimit struct {
	// The prefix to use when emitting :ref:`statistics <config_network_filters_rate_limit_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The rate limit domain to use in the rate limit service request.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// The rate limit descriptor list to use in the rate limit service request.
	Descriptors []*ratelimit.RateLimitDescriptor `protobuf:"bytes,3,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *types.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny bool `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	RateLimitService     *v3alpha.RateLimitServiceConfig `protobuf:"bytes,6,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_180c61393aaa1c89, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
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

func (m *RateLimit) GetRateLimitService() *v3alpha.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.network.rate_limit.v3alpha.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/rate_limit/v3alpha/rate_limit.proto", fileDescriptor_180c61393aaa1c89)
}

var fileDescriptor_180c61393aaa1c89 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x8f, 0x94, 0x30,
	0x1c, 0xc5, 0x53, 0x66, 0x77, 0x76, 0xb6, 0x24, 0xba, 0xf6, 0x22, 0xee, 0x41, 0x89, 0x27, 0xa2,
	0x49, 0x6b, 0x76, 0x12, 0x4f, 0x26, 0x26, 0x38, 0x47, 0x4d, 0x08, 0x5e, 0x4d, 0x48, 0x77, 0xf8,
	0x33, 0x36, 0x42, 0x4b, 0x4a, 0xc1, 0xe1, 0x03, 0x78, 0xf1, 0xe8, 0xc7, 0x9d, 0x93, 0x19, 0x5a,
	0x60, 0xf0, 0xe6, 0xad, 0xf4, 0xbd, 0xfe, 0xfe, 0x7d, 0xaf, 0xe0, 0x8f, 0x20, 0x3b, 0xd5, 0xb3,
	0xbd, 0x92, 0x85, 0x38, 0xb0, 0x42, 0x94, 0x06, 0x34, 0x93, 0x60, 0x7e, 0x2a, 0xfd, 0x83, 0x69,
	0x6e, 0x20, 0x2b, 0x45, 0x25, 0x0c, 0xeb, 0xb6, 0xbc, 0xac, 0xbf, 0xf3, 0x8b, 0x2d, 0x5a, 0x6b,
	0x65, 0x14, 0xa1, 0x03, 0x80, 0x5a, 0x00, 0xb5, 0x00, 0xea, 0x00, 0xf4, 0xc2, 0xed, 0x00, 0xf7,
	0x6f, 0xed, 0x40, 0x5e, 0x8b, 0x05, 0xd3, 0x4e, 0x99, 0x56, 0x16, 0x7e, 0x1f, 0x2d, 0x6e, 0x37,
	0xfb, 0xa6, 0x93, 0x65, 0xe3, 0x9c, 0x2f, 0x0f, 0x4a, 0x1d, 0x4a, 0x60, 0xc3, 0xd7, 0x63, 0x5b,
	0xb0, 0xbc, 0xd5, 0xdc, 0x08, 0x25, 0x9d, 0xfe, 0xbc, 0xe3, 0xa5, 0xc8, 0xb9, 0x01, 0x36, 0x2e,
	0xac, 0xf0, 0xfa, 0xd7, 0x0a, 0xdf, 0xa6, 0xdc, 0xc0, 0xe7, 0x33, 0x98, 0x44, 0xd8, 0x6f, 0x0c,
	0x37, 0x59, 0xad, 0xa1, 0x10, 0xc7, 0x00, 0x85, 0x28, 0xba, 0x8d, 0x6f, 0x4e, 0xf1, 0x95, 0xf6,
	0x42, 0x94, 0xe2, 0xb3, 0x96, 0x0c, 0x12, 0x79, 0x85, 0xd7, 0xb9, 0xaa, 0xb8, 0x90, 0x81, 0xb7,
	0x34, 0xb9, 0x6d, 0xf2, 0x0d, 0xfb, 0x39, 0x34, 0x7b, 0x2d, 0x6a, 0xa3, 0x74, 0x13, 0xac, 0xc2,
	0x55, 0xe4, 0x3f, 0xbc, 0x73, 0x75, 0xf1, 0x5a, 0x8c, 0x8d, 0xd0, 0x39, 0xf4, 0x74, 0x8f, 0xdd,
	0x74, 0x30, 0xde, 0x9c, 0xe2, 0xeb, 0x3f, 0xc8, 0xdb, 0xa0, 0xf4, 0x12, 0x47, 0xb6, 0xf8, 0xc6,
	0x88, 0x0a, 0x54, 0x6b, 0x82, 0xab, 0x10, 0x45, 0xfe, 0xc3, 0x0b, 0x6a, 0x1b, 0xa0, 0x63, 0x03,
	0x74, 0xe7, 0x1a, 0x48, 0x47, 0x27, 0x79, 0x83, 0x9f, 0x15, 0x5c, 0x94, 0xad, 0x86, 0xac, 0x52,
	0x39, 0x64, 0x39, 0xc8, 0x3e, 0xb8, 0x0e, 0x51, 0xb4, 0x49, 0x9f, 0x3a, 0xe1, 0x8b, 0xca, 0x61,
	0x07, 0xb2, 0x27, 0x12, 0x93, 0xf9, 0xf5, 0xb2, 0x06, 0x74, 0x27, 0xf6, 0x10, 0xac, 0x87, 0x59,
	0xef, 0x97, 0x8f, 0x3e, 0x07, 0x18, 0x23, 0x4d, 0x41, 0xbe, 0xda, 0x73, 0x9f, 0x06, 0xe3, 0x90,
	0xe5, 0x37, 0xf2, 0xee, 0x50, 0x7a, 0xa7, 0xff, 0x71, 0xc4, 0x09, 0xfe, 0x20, 0x94, 0xe5, 0xd6,
	0x5a, 0x1d, 0xfb, 0xff, 0xfc, 0xaf, 0xe2, 0x27, 0xd3, 0xcc, 0xe4, 0x5c, 0x40, 0x82, 0x1e, 0xd7,
	0x43, 0x13, 0xdb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0x6d, 0xd4, 0x48, 0xe3, 0x02, 0x00,
	0x00,
}