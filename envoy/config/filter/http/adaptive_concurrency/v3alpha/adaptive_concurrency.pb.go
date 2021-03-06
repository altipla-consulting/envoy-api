// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/adaptive_concurrency/v3alpha/adaptive_concurrency.proto

package envoy_config_filter_http_adaptive_concurrency_v3alpha

import (
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v3alpha/core"
	v3alpha "github.com/altipla-consulting/envoy-api/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/googleapis/google/api"
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

// Configuration parameters for the gradient controller.
type GradientControllerConfig struct {
	// The percentile to use when summarizing aggregated samples. Defaults to p50.
	SampleAggregatePercentile *v3alpha.Percent                                            `protobuf:"bytes,1,opt,name=sample_aggregate_percentile,json=sampleAggregatePercentile,proto3" json:"sample_aggregate_percentile,omitempty"`
	ConcurrencyLimitParams    *GradientControllerConfig_ConcurrencyLimitCalculationParams `protobuf:"bytes,2,opt,name=concurrency_limit_params,json=concurrencyLimitParams,proto3" json:"concurrency_limit_params,omitempty"`
	MinRttCalcParams          *GradientControllerConfig_MinimumRTTCalculationParams       `protobuf:"bytes,3,opt,name=min_rtt_calc_params,json=minRttCalcParams,proto3" json:"min_rtt_calc_params,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                                                    `json:"-"`
	XXX_unrecognized          []byte                                                      `json:"-"`
	XXX_sizecache             int32                                                       `json:"-"`
}

func (m *GradientControllerConfig) Reset()         { *m = GradientControllerConfig{} }
func (m *GradientControllerConfig) String() string { return proto.CompactTextString(m) }
func (*GradientControllerConfig) ProtoMessage()    {}
func (*GradientControllerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8176400d3397e66d, []int{0}
}

func (m *GradientControllerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientControllerConfig.Unmarshal(m, b)
}
func (m *GradientControllerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientControllerConfig.Marshal(b, m, deterministic)
}
func (m *GradientControllerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientControllerConfig.Merge(m, src)
}
func (m *GradientControllerConfig) XXX_Size() int {
	return xxx_messageInfo_GradientControllerConfig.Size(m)
}
func (m *GradientControllerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientControllerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GradientControllerConfig proto.InternalMessageInfo

func (m *GradientControllerConfig) GetSampleAggregatePercentile() *v3alpha.Percent {
	if m != nil {
		return m.SampleAggregatePercentile
	}
	return nil
}

func (m *GradientControllerConfig) GetConcurrencyLimitParams() *GradientControllerConfig_ConcurrencyLimitCalculationParams {
	if m != nil {
		return m.ConcurrencyLimitParams
	}
	return nil
}

func (m *GradientControllerConfig) GetMinRttCalcParams() *GradientControllerConfig_MinimumRTTCalculationParams {
	if m != nil {
		return m.MinRttCalcParams
	}
	return nil
}

// Parameters controlling the periodic recalculation of the concurrency limit from sampled request
// latencies.
type GradientControllerConfig_ConcurrencyLimitCalculationParams struct {
	// The maximum value the gradient is allowed to take. This influences how aggressively the
	// concurrency limit can increase. Defaults to 2.0.
	MaxGradient *types.DoubleValue `protobuf:"bytes,1,opt,name=max_gradient,json=maxGradient,proto3" json:"max_gradient,omitempty"`
	// The allowed upper-bound on the calculated concurrency limit. Defaults to 1000.
	MaxConcurrencyLimit *types.UInt32Value `protobuf:"bytes,2,opt,name=max_concurrency_limit,json=maxConcurrencyLimit,proto3" json:"max_concurrency_limit,omitempty"`
	// The period of time samples are taken to recalculate the concurrency limit.
	ConcurrencyUpdateInterval *types.Duration `protobuf:"bytes,3,opt,name=concurrency_update_interval,json=concurrencyUpdateInterval,proto3" json:"concurrency_update_interval,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}        `json:"-"`
	XXX_unrecognized          []byte          `json:"-"`
	XXX_sizecache             int32           `json:"-"`
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) Reset() {
	*m = GradientControllerConfig_ConcurrencyLimitCalculationParams{}
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) String() string {
	return proto.CompactTextString(m)
}
func (*GradientControllerConfig_ConcurrencyLimitCalculationParams) ProtoMessage() {}
func (*GradientControllerConfig_ConcurrencyLimitCalculationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8176400d3397e66d, []int{0, 0}
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Unmarshal(m, b)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Marshal(b, m, deterministic)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Merge(m, src)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_Size() int {
	return xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.Size(m)
}
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams.DiscardUnknown(m)
}

var xxx_messageInfo_GradientControllerConfig_ConcurrencyLimitCalculationParams proto.InternalMessageInfo

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) GetMaxGradient() *types.DoubleValue {
	if m != nil {
		return m.MaxGradient
	}
	return nil
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) GetMaxConcurrencyLimit() *types.UInt32Value {
	if m != nil {
		return m.MaxConcurrencyLimit
	}
	return nil
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) GetConcurrencyUpdateInterval() *types.Duration {
	if m != nil {
		return m.ConcurrencyUpdateInterval
	}
	return nil
}

// Parameters controlling the periodic minRTT recalculation.
type GradientControllerConfig_MinimumRTTCalculationParams struct {
	// The time interval between recalculating the minimum request round-trip time.
	Interval *types.Duration `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	// The number of requests to aggregate/sample during the minRTT recalculation window before
	// updating. Defaults to 50.
	RequestCount *types.UInt32Value `protobuf:"bytes,2,opt,name=request_count,json=requestCount,proto3" json:"request_count,omitempty"`
	// Randomized time delta that will be introduced to the start of the minRTT calculation window.
	// This is represented as a percentage of the interval duration. Defaults to 15%.
	//
	// Example: If the interval is 10s and the jitter is 15%, the next window will begin
	// somewhere in the range (10s - 11.5s).
	Jitter               *v3alpha.Percent `protobuf:"bytes,3,opt,name=jitter,proto3" json:"jitter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) Reset() {
	*m = GradientControllerConfig_MinimumRTTCalculationParams{}
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) String() string {
	return proto.CompactTextString(m)
}
func (*GradientControllerConfig_MinimumRTTCalculationParams) ProtoMessage() {}
func (*GradientControllerConfig_MinimumRTTCalculationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8176400d3397e66d, []int{0, 1}
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Unmarshal(m, b)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Marshal(b, m, deterministic)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Merge(m, src)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_Size() int {
	return xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.Size(m)
}
func (m *GradientControllerConfig_MinimumRTTCalculationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams.DiscardUnknown(m)
}

var xxx_messageInfo_GradientControllerConfig_MinimumRTTCalculationParams proto.InternalMessageInfo

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetInterval() *types.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetRequestCount() *types.UInt32Value {
	if m != nil {
		return m.RequestCount
	}
	return nil
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) GetJitter() *v3alpha.Percent {
	if m != nil {
		return m.Jitter
	}
	return nil
}

type AdaptiveConcurrency struct {
	// Types that are valid to be assigned to ConcurrencyControllerConfig:
	//	*AdaptiveConcurrency_GradientControllerConfig
	ConcurrencyControllerConfig isAdaptiveConcurrency_ConcurrencyControllerConfig `protobuf_oneof:"concurrency_controller_config"`
	// If set to false, the adaptive concurrency filter will operate as a pass-through filter. If the
	// message is unspecified, the filter will be enabled.
	Enabled              *core.RuntimeFeatureFlag `protobuf:"bytes,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AdaptiveConcurrency) Reset()         { *m = AdaptiveConcurrency{} }
func (m *AdaptiveConcurrency) String() string { return proto.CompactTextString(m) }
func (*AdaptiveConcurrency) ProtoMessage()    {}
func (*AdaptiveConcurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor_8176400d3397e66d, []int{1}
}

func (m *AdaptiveConcurrency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdaptiveConcurrency.Unmarshal(m, b)
}
func (m *AdaptiveConcurrency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdaptiveConcurrency.Marshal(b, m, deterministic)
}
func (m *AdaptiveConcurrency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdaptiveConcurrency.Merge(m, src)
}
func (m *AdaptiveConcurrency) XXX_Size() int {
	return xxx_messageInfo_AdaptiveConcurrency.Size(m)
}
func (m *AdaptiveConcurrency) XXX_DiscardUnknown() {
	xxx_messageInfo_AdaptiveConcurrency.DiscardUnknown(m)
}

var xxx_messageInfo_AdaptiveConcurrency proto.InternalMessageInfo

type isAdaptiveConcurrency_ConcurrencyControllerConfig interface {
	isAdaptiveConcurrency_ConcurrencyControllerConfig()
}

type AdaptiveConcurrency_GradientControllerConfig struct {
	GradientControllerConfig *GradientControllerConfig `protobuf:"bytes,1,opt,name=gradient_controller_config,json=gradientControllerConfig,proto3,oneof"`
}

func (*AdaptiveConcurrency_GradientControllerConfig) isAdaptiveConcurrency_ConcurrencyControllerConfig() {
}

func (m *AdaptiveConcurrency) GetConcurrencyControllerConfig() isAdaptiveConcurrency_ConcurrencyControllerConfig {
	if m != nil {
		return m.ConcurrencyControllerConfig
	}
	return nil
}

func (m *AdaptiveConcurrency) GetGradientControllerConfig() *GradientControllerConfig {
	if x, ok := m.GetConcurrencyControllerConfig().(*AdaptiveConcurrency_GradientControllerConfig); ok {
		return x.GradientControllerConfig
	}
	return nil
}

func (m *AdaptiveConcurrency) GetEnabled() *core.RuntimeFeatureFlag {
	if m != nil {
		return m.Enabled
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdaptiveConcurrency) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdaptiveConcurrency_GradientControllerConfig)(nil),
	}
}

func init() {
	proto.RegisterType((*GradientControllerConfig)(nil), "envoy.config.filter.http.adaptive_concurrency.v3alpha.GradientControllerConfig")
	proto.RegisterType((*GradientControllerConfig_ConcurrencyLimitCalculationParams)(nil), "envoy.config.filter.http.adaptive_concurrency.v3alpha.GradientControllerConfig.ConcurrencyLimitCalculationParams")
	proto.RegisterType((*GradientControllerConfig_MinimumRTTCalculationParams)(nil), "envoy.config.filter.http.adaptive_concurrency.v3alpha.GradientControllerConfig.MinimumRTTCalculationParams")
	proto.RegisterType((*AdaptiveConcurrency)(nil), "envoy.config.filter.http.adaptive_concurrency.v3alpha.AdaptiveConcurrency")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/adaptive_concurrency/v3alpha/adaptive_concurrency.proto", fileDescriptor_8176400d3397e66d)
}

var fileDescriptor_8176400d3397e66d = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x4e, 0xdb, 0x4a,
	0x14, 0x8d, 0x03, 0x8f, 0xf0, 0x06, 0xde, 0x13, 0x32, 0x7a, 0xaf, 0x21, 0xa1, 0x15, 0xa0, 0x2e,
	0x2a, 0x16, 0xb6, 0x44, 0xd4, 0x75, 0x85, 0x83, 0x68, 0xa9, 0x5a, 0x11, 0x59, 0x50, 0xa9, 0xea,
	0xc2, 0xba, 0x71, 0x2e, 0x66, 0xda, 0xf1, 0xcc, 0x30, 0x19, 0xa7, 0xc9, 0x2f, 0xf4, 0x0f, 0xba,
	0xe8, 0xae, 0xab, 0xee, 0xdb, 0xff, 0xe9, 0x1f, 0x54, 0xea, 0xa6, 0xca, 0xaa, 0xb2, 0x67, 0x1c,
	0x22, 0x02, 0x54, 0x20, 0xb2, 0x8a, 0x75, 0xef, 0x39, 0xf7, 0xdc, 0x73, 0x8f, 0x4d, 0x3a, 0xc8,
	0x07, 0x62, 0xe4, 0xc7, 0x82, 0x9f, 0xd0, 0xc4, 0x3f, 0xa1, 0x4c, 0xa3, 0xf2, 0x4f, 0xb5, 0x96,
	0x3e, 0xf4, 0x40, 0x6a, 0x3a, 0xc0, 0x28, 0x16, 0x3c, 0xce, 0x94, 0x42, 0x1e, 0x8f, 0xfc, 0x41,
	0x0b, 0x98, 0x3c, 0x85, 0x4b, 0x8b, 0x9e, 0x54, 0x42, 0x0b, 0xf7, 0x71, 0xc1, 0xe8, 0x19, 0x46,
	0xcf, 0x30, 0x7a, 0x39, 0xa3, 0x77, 0x29, 0xc8, 0x32, 0x36, 0x36, 0x8d, 0x10, 0x90, 0x74, 0x32,
	0x24, 0x16, 0x0a, 0xfd, 0x2e, 0xf4, 0xd1, 0x30, 0x37, 0x36, 0x4c, 0x8b, 0x1e, 0x49, 0x9c, 0xf4,
	0x48, 0x54, 0x31, 0x72, 0x6d, 0x3b, 0xd6, 0x13, 0x21, 0x12, 0x86, 0x05, 0x0b, 0x70, 0x2e, 0x34,
	0x68, 0x2a, 0x78, 0xdf, 0x56, 0x1f, 0xd8, 0x6a, 0xf1, 0xd4, 0xcd, 0x4e, 0xfc, 0x5e, 0xa6, 0x8a,
	0x86, 0xab, 0xea, 0xef, 0x15, 0x48, 0x89, 0xaa, 0xc4, 0xdf, 0x1b, 0x00, 0xa3, 0x3d, 0xd0, 0xe8,
	0x97, 0x7f, 0x4c, 0x61, 0xeb, 0x67, 0x8d, 0xd4, 0x9f, 0x2a, 0xe8, 0x51, 0xe4, 0xba, 0x2d, 0xb8,
	0x56, 0x82, 0x31, 0x54, 0xed, 0xc2, 0x01, 0xf7, 0x0d, 0x69, 0xf6, 0x21, 0x95, 0x0c, 0x23, 0x48,
	0x12, 0x85, 0x09, 0x68, 0x8c, 0xac, 0x6a, 0xca, 0xb0, 0xee, 0x6c, 0x38, 0x8f, 0x96, 0x76, 0x9a,
	0x9e, 0x71, 0x2d, 0xdf, 0xad, 0xb4, 0xc4, 0xeb, 0x98, 0xae, 0x70, 0xcd, 0xe0, 0x77, 0x4b, 0x78,
	0x67, 0x82, 0x76, 0xbf, 0x3a, 0xa4, 0x3e, 0xe5, 0x66, 0xc4, 0x68, 0x4a, 0x75, 0x24, 0x41, 0x41,
	0xda, 0xaf, 0x57, 0x0b, 0xea, 0x33, 0xef, 0x56, 0x07, 0xf1, 0xae, 0x5a, 0xc8, 0x6b, 0x9f, 0x37,
	0xbf, 0xc8, 0xc7, 0xb5, 0x81, 0xc5, 0x19, 0x2b, 0x8c, 0xec, 0x14, 0x83, 0x83, 0xc5, 0x71, 0xf0,
	0xd7, 0x07, 0xa7, 0xba, 0xe2, 0x84, 0xff, 0xc7, 0x17, 0x9a, 0x4d, 0x87, 0xfb, 0xd9, 0x21, 0xab,
	0x29, 0xe5, 0x91, 0xd2, 0x3a, 0x8a, 0x81, 0xc5, 0xa5, 0xe4, 0xb9, 0x42, 0xf2, 0xbb, 0xbb, 0x96,
	0xfc, 0x92, 0x72, 0x9a, 0x66, 0x69, 0x78, 0x74, 0x74, 0x9d, 0xd8, 0x95, 0x94, 0xf2, 0x50, 0x17,
	0xfb, 0x98, 0x5a, 0xe3, 0x5b, 0x95, 0x6c, 0xfe, 0x71, 0x5d, 0xf7, 0x90, 0x2c, 0xa7, 0x30, 0x8c,
	0x12, 0x3b, 0xdd, 0x9e, 0x74, 0xdd, 0x33, 0x71, 0xf2, 0xca, 0x38, 0x79, 0x7b, 0x22, 0xeb, 0x32,
	0x7c, 0x05, 0x2c, 0xc3, 0xe0, 0xdf, 0x71, 0xb0, 0xe4, 0xfe, 0xbd, 0x59, 0x29, 0x7e, 0x3f, 0x9e,
	0x84, 0x4b, 0x29, 0x0c, 0x4b, 0xf9, 0xee, 0x6b, 0xf2, 0x5f, 0x4e, 0x38, 0x73, 0x58, 0x7b, 0xd1,
	0x59, 0xe6, 0xe3, 0x03, 0xae, 0x5b, 0x3b, 0x86, 0xb9, 0x36, 0x0e, 0xe6, 0xb7, 0xab, 0x1b, 0x95,
	0x70, 0x35, 0x85, 0xe1, 0x45, 0xf1, 0x2e, 0x92, 0xe6, 0x34, 0x6d, 0x26, 0xf3, 0x18, 0x47, 0x94,
	0x6b, 0x54, 0x03, 0x60, 0xd6, 0xff, 0xb5, 0x59, 0xe9, 0xf6, 0x4d, 0x09, 0xc8, 0x38, 0xa8, 0x7d,
	0x71, 0xe6, 0x17, 0x9d, 0xed, 0x4a, 0xb8, 0x36, 0xc5, 0x74, 0x5c, 0x10, 0x1d, 0x58, 0x9e, 0xc6,
	0x77, 0x87, 0x34, 0xaf, 0x31, 0xdd, 0xdd, 0x25, 0x8b, 0x93, 0x99, 0xce, 0x4d, 0x66, 0x4e, 0x60,
	0xee, 0x73, 0xf2, 0x8f, 0xc2, 0xb3, 0x0c, 0xfb, 0x3a, 0x8a, 0x45, 0xc6, 0x6f, 0x68, 0xce, 0xb2,
	0xc5, 0xb6, 0x73, 0xa8, 0xdb, 0x22, 0x0b, 0x6f, 0xa9, 0xd6, 0xa8, 0xac, 0x01, 0xd7, 0xbe, 0x8e,
	0xb6, 0x75, 0xeb, 0x53, 0x95, 0xac, 0xee, 0xda, 0x38, 0x4e, 0xf9, 0xec, 0x7e, 0x74, 0x48, 0xa3,
	0xcc, 0x42, 0x7e, 0x43, 0x1b, 0xc5, 0xc8, 0xa4, 0xd9, 0xae, 0x7b, 0x78, 0xc7, 0x11, 0x3f, 0x8f,
	0xf1, 0xb3, 0x4a, 0x58, 0x4f, 0xae, 0xfa, 0x18, 0xed, 0x91, 0x1a, 0x72, 0xe8, 0x32, 0xec, 0x59,
	0xbb, 0xb6, 0xad, 0x0e, 0x90, 0x74, 0x32, 0x23, 0xff, 0xee, 0x7a, 0x61, 0xc6, 0x35, 0x4d, 0x71,
	0x1f, 0x41, 0x67, 0x0a, 0xf7, 0x19, 0x24, 0x61, 0x09, 0x0d, 0x1e, 0x92, 0xfb, 0xd3, 0x21, 0x9a,
	0xd9, 0xd1, 0x9d, 0xfb, 0x15, 0x38, 0x01, 0x90, 0x36, 0x15, 0x86, 0x5e, 0x2a, 0x31, 0x1c, 0xdd,
	0x6e, 0xe3, 0xa0, 0x7e, 0x89, 0xc7, 0x9d, 0xfc, 0xb6, 0x1d, 0xa7, 0xbb, 0x50, 0x1c, 0xb9, 0xf5,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x78, 0xab, 0x99, 0xee, 0xc6, 0x06, 0x00, 0x00,
}
