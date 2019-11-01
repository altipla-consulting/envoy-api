// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/cluster/outlier_detection.proto

package envoy_api_v2_cluster

import (
	fmt "fmt"
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

// See the :ref:`architecture overview <arch_overview_outlier_detection>` for
// more information on outlier detection.
// [#next-free-field: 21]
type OutlierDetection struct {
	// The number of consecutive 5xx responses or local origin errors that are mapped
	// to 5xx error codes before a consecutive 5xx ejection
	// occurs. Defaults to 5.
	Consecutive_5Xx *types.UInt32Value `protobuf:"bytes,1,opt,name=consecutive_5xx,json=consecutive5xx,proto3" json:"consecutive_5xx,omitempty"`
	// The time interval between ejection analysis sweeps. This can result in
	// both new ejections as well as hosts being returned to service. Defaults
	// to 10000ms or 10s.
	Interval *types.Duration `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	// The base time that a host is ejected for. The real time is equal to the
	// base time multiplied by the number of times the host has been ejected.
	// Defaults to 30000ms or 30s.
	BaseEjectionTime *types.Duration `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime,proto3" json:"base_ejection_time,omitempty"`
	// The maximum % of an upstream cluster that can be ejected due to outlier
	// detection. Defaults to 10% but will eject at least one host regardless of the value.
	MaxEjectionPercent *types.UInt32Value `protobuf:"bytes,4,opt,name=max_ejection_percent,json=maxEjectionPercent,proto3" json:"max_ejection_percent,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive 5xx. This setting can be used to disable
	// ejection or to ramp it up slowly. Defaults to 100.
	EnforcingConsecutive_5Xx *types.UInt32Value `protobuf:"bytes,5,opt,name=enforcing_consecutive_5xx,json=enforcingConsecutive5xx,proto3" json:"enforcing_consecutive_5xx,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics. This setting can be used to
	// disable ejection or to ramp it up slowly. Defaults to 100.
	EnforcingSuccessRate *types.UInt32Value `protobuf:"bytes,6,opt,name=enforcing_success_rate,json=enforcingSuccessRate,proto3" json:"enforcing_success_rate,omitempty"`
	// The number of hosts in a cluster that must have enough request volume to
	// detect success rate outliers. If the number of hosts is less than this
	// setting, outlier detection via success rate statistics is not performed
	// for any host in the cluster. Defaults to 5.
	SuccessRateMinimumHosts *types.UInt32Value `protobuf:"bytes,7,opt,name=success_rate_minimum_hosts,json=successRateMinimumHosts,proto3" json:"success_rate_minimum_hosts,omitempty"`
	// The minimum number of total requests that must be collected in one
	// interval (as defined by the interval duration above) to include this host
	// in success rate based outlier detection. If the volume is lower than this
	// setting, outlier detection via success rate statistics is not performed
	// for that host. Defaults to 100.
	SuccessRateRequestVolume *types.UInt32Value `protobuf:"bytes,8,opt,name=success_rate_request_volume,json=successRateRequestVolume,proto3" json:"success_rate_request_volume,omitempty"`
	// This factor is used to determine the ejection threshold for success rate
	// outlier ejection. The ejection threshold is the difference between the
	// mean success rate, and the product of this factor and the standard
	// deviation of the mean success rate: mean - (stdev *
	// success_rate_stdev_factor). This factor is divided by a thousand to get a
	// double. That is, if the desired factor is 1.9, the runtime value should
	// be 1900. Defaults to 1900.
	SuccessRateStdevFactor *types.UInt32Value `protobuf:"bytes,9,opt,name=success_rate_stdev_factor,json=successRateStdevFactor,proto3" json:"success_rate_stdev_factor,omitempty"`
	// The number of consecutive gateway failures (502, 503, 504 status codes)
	// before a consecutive gateway failure ejection occurs. Defaults to 5.
	ConsecutiveGatewayFailure *types.UInt32Value `protobuf:"bytes,10,opt,name=consecutive_gateway_failure,json=consecutiveGatewayFailure,proto3" json:"consecutive_gateway_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive gateway failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 0.
	EnforcingConsecutiveGatewayFailure *types.UInt32Value `protobuf:"bytes,11,opt,name=enforcing_consecutive_gateway_failure,json=enforcingConsecutiveGatewayFailure,proto3" json:"enforcing_consecutive_gateway_failure,omitempty"`
	// Determines whether to distinguish local origin failures from external errors. If set to true
	// the following configuration parameters are taken into account:
	// :ref:`consecutive_local_origin_failure<envoy_api_field_cluster.OutlierDetection.consecutive_local_origin_failure>`,
	// :ref:`enforcing_consecutive_local_origin_failure<envoy_api_field_cluster.OutlierDetection.enforcing_consecutive_local_origin_failure>`
	// and
	// :ref:`enforcing_local_origin_success_rate<envoy_api_field_cluster.OutlierDetection.enforcing_local_origin_success_rate>`.
	// Defaults to false.
	SplitExternalLocalOriginErrors bool `protobuf:"varint,12,opt,name=split_external_local_origin_errors,json=splitExternalLocalOriginErrors,proto3" json:"split_external_local_origin_errors,omitempty"`
	// The number of consecutive locally originated failures before ejection
	// occurs. Defaults to 5. Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	ConsecutiveLocalOriginFailure *types.UInt32Value `protobuf:"bytes,13,opt,name=consecutive_local_origin_failure,json=consecutiveLocalOriginFailure,proto3" json:"consecutive_local_origin_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive locally originated failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 100.
	// Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	EnforcingConsecutiveLocalOriginFailure *types.UInt32Value `protobuf:"bytes,14,opt,name=enforcing_consecutive_local_origin_failure,json=enforcingConsecutiveLocalOriginFailure,proto3" json:"enforcing_consecutive_local_origin_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics for locally originated errors.
	// This setting can be used to disable ejection or to ramp it up slowly. Defaults to 100.
	// Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	EnforcingLocalOriginSuccessRate *types.UInt32Value `protobuf:"bytes,15,opt,name=enforcing_local_origin_success_rate,json=enforcingLocalOriginSuccessRate,proto3" json:"enforcing_local_origin_success_rate,omitempty"`
	// The failure percentage to use when determining failure percentage-based outlier detection. If
	// the failure percentage of a given host is greater than or equal to this value, it will be
	// ejected. Defaults to 85.
	FailurePercentageThreshold *types.UInt32Value `protobuf:"bytes,16,opt,name=failure_percentage_threshold,json=failurePercentageThreshold,proto3" json:"failure_percentage_threshold,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status is detected through
	// failure percentage statistics. This setting can be used to disable ejection or to ramp it up
	// slowly. Defaults to 0.
	//
	// [#next-major-version: setting this without setting failure_percentage_threshold should be
	// invalid in v4.]
	EnforcingFailurePercentage *types.UInt32Value `protobuf:"bytes,17,opt,name=enforcing_failure_percentage,json=enforcingFailurePercentage,proto3" json:"enforcing_failure_percentage,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status is detected through
	// local-origin failure percentage statistics. This setting can be used to disable ejection or to
	// ramp it up slowly. Defaults to 0.
	EnforcingFailurePercentageLocalOrigin *types.UInt32Value `protobuf:"bytes,18,opt,name=enforcing_failure_percentage_local_origin,json=enforcingFailurePercentageLocalOrigin,proto3" json:"enforcing_failure_percentage_local_origin,omitempty"`
	// The minimum number of hosts in a cluster in order to perform failure percentage-based ejection.
	// If the total number of hosts in the cluster is less than this value, failure percentage-based
	// ejection will not be performed. Defaults to 5.
	FailurePercentageMinimumHosts *types.UInt32Value `protobuf:"bytes,19,opt,name=failure_percentage_minimum_hosts,json=failurePercentageMinimumHosts,proto3" json:"failure_percentage_minimum_hosts,omitempty"`
	// The minimum number of total requests that must be collected in one interval (as defined by the
	// interval duration above) to perform failure percentage-based ejection for this host. If the
	// volume is lower than this setting, failure percentage-based ejection will not be performed for
	// this host. Defaults to 50.
	FailurePercentageRequestVolume *types.UInt32Value `protobuf:"bytes,20,opt,name=failure_percentage_request_volume,json=failurePercentageRequestVolume,proto3" json:"failure_percentage_request_volume,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}           `json:"-"`
	XXX_unrecognized               []byte             `json:"-"`
	XXX_sizecache                  int32              `json:"-"`
}

func (m *OutlierDetection) Reset()         { *m = OutlierDetection{} }
func (m *OutlierDetection) String() string { return proto.CompactTextString(m) }
func (*OutlierDetection) ProtoMessage()    {}
func (*OutlierDetection) Descriptor() ([]byte, []int) {
	return fileDescriptor_56cd87362a3f00c9, []int{0}
}

func (m *OutlierDetection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierDetection.Unmarshal(m, b)
}
func (m *OutlierDetection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierDetection.Marshal(b, m, deterministic)
}
func (m *OutlierDetection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierDetection.Merge(m, src)
}
func (m *OutlierDetection) XXX_Size() int {
	return xxx_messageInfo_OutlierDetection.Size(m)
}
func (m *OutlierDetection) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierDetection.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierDetection proto.InternalMessageInfo

func (m *OutlierDetection) GetConsecutive_5Xx() *types.UInt32Value {
	if m != nil {
		return m.Consecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetInterval() *types.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *OutlierDetection) GetBaseEjectionTime() *types.Duration {
	if m != nil {
		return m.BaseEjectionTime
	}
	return nil
}

func (m *OutlierDetection) GetMaxEjectionPercent() *types.UInt32Value {
	if m != nil {
		return m.MaxEjectionPercent
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutive_5Xx() *types.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutive_5Xx
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingSuccessRate() *types.UInt32Value {
	if m != nil {
		return m.EnforcingSuccessRate
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateMinimumHosts() *types.UInt32Value {
	if m != nil {
		return m.SuccessRateMinimumHosts
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateRequestVolume() *types.UInt32Value {
	if m != nil {
		return m.SuccessRateRequestVolume
	}
	return nil
}

func (m *OutlierDetection) GetSuccessRateStdevFactor() *types.UInt32Value {
	if m != nil {
		return m.SuccessRateStdevFactor
	}
	return nil
}

func (m *OutlierDetection) GetConsecutiveGatewayFailure() *types.UInt32Value {
	if m != nil {
		return m.ConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveGatewayFailure() *types.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveGatewayFailure
	}
	return nil
}

func (m *OutlierDetection) GetSplitExternalLocalOriginErrors() bool {
	if m != nil {
		return m.SplitExternalLocalOriginErrors
	}
	return false
}

func (m *OutlierDetection) GetConsecutiveLocalOriginFailure() *types.UInt32Value {
	if m != nil {
		return m.ConsecutiveLocalOriginFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingConsecutiveLocalOriginFailure() *types.UInt32Value {
	if m != nil {
		return m.EnforcingConsecutiveLocalOriginFailure
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingLocalOriginSuccessRate() *types.UInt32Value {
	if m != nil {
		return m.EnforcingLocalOriginSuccessRate
	}
	return nil
}

func (m *OutlierDetection) GetFailurePercentageThreshold() *types.UInt32Value {
	if m != nil {
		return m.FailurePercentageThreshold
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingFailurePercentage() *types.UInt32Value {
	if m != nil {
		return m.EnforcingFailurePercentage
	}
	return nil
}

func (m *OutlierDetection) GetEnforcingFailurePercentageLocalOrigin() *types.UInt32Value {
	if m != nil {
		return m.EnforcingFailurePercentageLocalOrigin
	}
	return nil
}

func (m *OutlierDetection) GetFailurePercentageMinimumHosts() *types.UInt32Value {
	if m != nil {
		return m.FailurePercentageMinimumHosts
	}
	return nil
}

func (m *OutlierDetection) GetFailurePercentageRequestVolume() *types.UInt32Value {
	if m != nil {
		return m.FailurePercentageRequestVolume
	}
	return nil
}

func init() {
	proto.RegisterType((*OutlierDetection)(nil), "envoy.api.v2.cluster.OutlierDetection")
}

func init() {
	proto.RegisterFile("envoy/api/v2/cluster/outlier_detection.proto", fileDescriptor_56cd87362a3f00c9)
}

var fileDescriptor_56cd87362a3f00c9 = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0xc7, 0x37, 0x59, 0x3e, 0xb2, 0xc3, 0x2e, 0xb0, 0xb3, 0x59, 0x70, 0x80, 0x65, 0x69, 0x2a,
	0x2a, 0x8a, 0x2a, 0x47, 0x0a, 0xe2, 0xba, 0x6a, 0x20, 0xf4, 0x43, 0x6d, 0x49, 0x03, 0x05, 0x55,
	0x54, 0x1a, 0x0d, 0xce, 0x89, 0x99, 0xca, 0xf6, 0xb8, 0x33, 0x63, 0x63, 0xae, 0xaa, 0xbe, 0x0e,
	0x0f, 0xd1, 0x07, 0xea, 0x23, 0x70, 0x55, 0xc5, 0x8e, 0x13, 0x3b, 0x31, 0xad, 0x7d, 0x17, 0x69,
	0xce, 0xff, 0xf7, 0x3b, 0x73, 0x3c, 0x9e, 0x18, 0x3d, 0x01, 0xc7, 0xe7, 0x37, 0x0d, 0xea, 0xb2,
	0x86, 0xdf, 0x6c, 0x18, 0x96, 0x27, 0x15, 0x88, 0x06, 0xf7, 0x94, 0xc5, 0x40, 0x90, 0x1e, 0x28,
	0x30, 0x14, 0xe3, 0x8e, 0xee, 0x0a, 0xae, 0x38, 0xae, 0x86, 0xd5, 0x3a, 0x75, 0x99, 0xee, 0x37,
	0xf5, 0x61, 0xf5, 0xda, 0xa6, 0xc9, 0xb9, 0x69, 0x41, 0x23, 0xac, 0xb9, 0xf4, 0xfa, 0x8d, 0x9e,
	0x27, 0xe8, 0x38, 0x35, 0xbd, 0x7e, 0x2d, 0xa8, 0xeb, 0x82, 0x90, 0xc3, 0xf5, 0x55, 0x9f, 0x5a,
	0xac, 0x47, 0x15, 0x34, 0xe2, 0x1f, 0xd1, 0x42, 0xfd, 0xdb, 0x12, 0x5a, 0x3e, 0x8e, 0x5a, 0x39,
	0x8c, 0x3b, 0xc1, 0x6d, 0xb4, 0x64, 0x70, 0x47, 0x82, 0xe1, 0x29, 0xe6, 0x03, 0xd9, 0x0f, 0x02,
	0xad, 0xb4, 0x55, 0xda, 0x59, 0x68, 0x6e, 0xe8, 0x91, 0x47, 0x8f, 0x3d, 0xfa, 0xfb, 0x97, 0x8e,
	0xda, 0x6b, 0x9e, 0x51, 0xcb, 0x83, 0xee, 0x62, 0x22, 0xb4, 0x1f, 0x04, 0xf8, 0x29, 0xaa, 0x30,
	0x47, 0x81, 0xf0, 0xa9, 0xa5, 0x95, 0xc3, 0x7c, 0x6d, 0x2a, 0x7f, 0x38, 0xdc, 0x47, 0xab, 0x72,
	0xd7, 0x9a, 0xbd, 0x2d, 0x95, 0x77, 0x7f, 0xeb, 0x8e, 0x42, 0xf8, 0x1d, 0xc2, 0x97, 0x54, 0x02,
	0x81, 0x4f, 0x51, 0x63, 0x44, 0x31, 0x1b, 0xb4, 0xdf, 0xf3, 0xa3, 0x96, 0x07, 0xf1, 0xf6, 0x30,
	0x7d, 0xca, 0x6c, 0xc0, 0xe7, 0xa8, 0x6a, 0xd3, 0x60, 0x4c, 0x74, 0x41, 0x18, 0xe0, 0x28, 0x6d,
	0xe6, 0xd7, 0xfb, 0x6b, 0xcd, 0xdf, 0xb5, 0x66, 0x76, 0xcb, 0x5a, 0xaf, 0x8b, 0x6d, 0x1a, 0xc4,
	0xd4, 0x4e, 0x04, 0xc0, 0x14, 0xd5, 0xc0, 0xe9, 0x73, 0x61, 0x30, 0xc7, 0x24, 0x93, 0xd3, 0x9b,
	0x2d, 0x42, 0x5f, 0x1d, 0x71, 0x0e, 0xd2, 0xf3, 0xbc, 0x40, 0x2b, 0x63, 0x85, 0xf4, 0x0c, 0x03,
	0xa4, 0x24, 0x82, 0x2a, 0xd0, 0xe6, 0x8a, 0xf0, 0xab, 0x23, 0xc8, 0x49, 0xc4, 0xe8, 0x52, 0x05,
	0xf8, 0x03, 0x5a, 0x4b, 0x22, 0x89, 0xcd, 0x1c, 0x66, 0x7b, 0x36, 0xb9, 0xe2, 0x52, 0x49, 0x6d,
	0x3e, 0xc7, 0xe3, 0x5f, 0x95, 0x63, 0xdc, 0x9b, 0x28, 0xfd, 0x62, 0x10, 0xc6, 0x17, 0x68, 0x3d,
	0x85, 0x16, 0xf0, 0xd9, 0x03, 0xa9, 0x88, 0xcf, 0x2d, 0xcf, 0x06, 0xad, 0x92, 0x83, 0xad, 0x25,
	0xd8, 0xdd, 0x28, 0x7e, 0x16, 0xa6, 0xf1, 0x39, 0xaa, 0xa5, 0xe0, 0x52, 0xf5, 0xc0, 0x27, 0x7d,
	0x6a, 0x28, 0x2e, 0xb4, 0x3f, 0x72, 0xa0, 0x57, 0x12, 0xe8, 0x93, 0x41, 0xf8, 0x28, 0xcc, 0xe2,
	0x8f, 0x68, 0x3d, 0xf9, 0x18, 0x4d, 0xaa, 0xe0, 0x9a, 0xde, 0x90, 0x3e, 0x65, 0x96, 0x27, 0x40,
	0x43, 0x39, 0xd0, 0xb5, 0x04, 0xe0, 0x79, 0x94, 0x3f, 0x8a, 0xe2, 0x38, 0x40, 0xdb, 0xd9, 0xc7,
	0x65, 0xd2, 0xb3, 0x50, 0xe4, 0xd1, 0xd6, 0xb3, 0x8e, 0xce, 0x84, 0xf9, 0x15, 0xaa, 0x4b, 0xd7,
	0x62, 0x8a, 0x40, 0xa0, 0x40, 0x38, 0xd4, 0x22, 0x16, 0x37, 0xa8, 0x45, 0xb8, 0x60, 0x26, 0x73,
	0x08, 0x08, 0xc1, 0x85, 0xd4, 0xfe, 0xdc, 0x2a, 0xed, 0x54, 0xba, 0x9b, 0x61, 0x65, 0x7b, 0x58,
	0xf8, 0x7a, 0x50, 0x77, 0x1c, 0x96, 0xb5, 0xc3, 0x2a, 0x0c, 0x68, 0x2b, 0xd9, 0x7b, 0x0a, 0x14,
	0x6f, 0xe0, 0xaf, 0x1c, 0x83, 0xfa, 0x2f, 0x41, 0x49, 0x58, 0xe2, 0x96, 0xbf, 0x96, 0xd0, 0x6e,
	0xf6, 0xb4, 0x32, 0x8d, 0x8b, 0x45, 0x46, 0xf6, 0x28, 0x6b, 0x64, 0x19, 0x3d, 0x48, 0xf4, 0x70,
	0xdc, 0x42, 0x4a, 0x9b, 0x7a, 0x13, 0x97, 0x8a, 0xb8, 0xff, 0x1f, 0x11, 0x13, 0xc2, 0xe4, 0x4b,
	0x69, 0xa2, 0x8d, 0xe1, 0xa6, 0xe2, 0x8b, 0x8a, 0x9a, 0x40, 0xd4, 0x95, 0x00, 0x79, 0xc5, 0xad,
	0x9e, 0xb6, 0x5c, 0xc4, 0xb6, 0x36, 0x44, 0x75, 0x46, 0xa4, 0xd3, 0x18, 0x34, 0x10, 0x8d, 0x77,
	0x37, 0xad, 0xd4, 0xfe, 0x2e, 0x24, 0x1a, 0xa1, 0x8e, 0x26, 0x8d, 0xf8, 0x0b, 0x7a, 0xfc, 0x33,
	0x51, 0x6a, 0xb2, 0x1a, 0x2e, 0x62, 0xdd, 0xbe, 0xdf, 0x9a, 0x98, 0xee, 0xe0, 0xc8, 0x66, 0x68,
	0xd3, 0xb7, 0xdd, 0x3f, 0x79, 0x8e, 0xec, 0xd4, 0x34, 0x53, 0x77, 0x9e, 0x89, 0x1e, 0x64, 0x68,
	0x26, 0x6e, 0xbe, 0x6a, 0x0e, 0xcf, 0xe6, 0x94, 0x27, 0x75, 0xff, 0xb5, 0x38, 0xaa, 0x33, 0xae,
	0x87, 0x1f, 0x0d, 0xae, 0xe0, 0xc1, 0x8d, 0x9e, 0xf5, 0xfd, 0xd0, 0xfa, 0x77, 0xf2, 0x3f, 0xbe,
	0x33, 0x70, 0x75, 0x4a, 0xb7, 0xe5, 0x95, 0x76, 0x58, 0xff, 0xcc, 0x65, 0xfa, 0x59, 0x53, 0x3f,
	0x88, 0xea, 0xdf, 0x9e, 0x7c, 0xbf, 0x6f, 0xe1, 0x72, 0x2e, 0x6c, 0x73, 0xef, 0x47, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x15, 0x3e, 0x57, 0x67, 0xd7, 0x08, 0x00, 0x00,
}
