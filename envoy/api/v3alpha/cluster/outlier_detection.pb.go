// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/cluster/outlier_detection.proto

package envoy_api_v3alpha_cluster

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
	// :ref:`consecutive_local_origin_failure<envoy_api_field_api.v3alpha.cluster.OutlierDetection.consecutive_local_origin_failure>`,
	// :ref:`enforcing_consecutive_local_origin_failure<envoy_api_field_api.v3alpha.cluster.OutlierDetection.enforcing_consecutive_local_origin_failure>`
	// and
	// :ref:`enforcing_local_origin_success_rate<envoy_api_field_api.v3alpha.cluster.OutlierDetection.enforcing_local_origin_success_rate>`.
	// Defaults to false.
	SplitExternalLocalOriginErrors bool `protobuf:"varint,12,opt,name=split_external_local_origin_errors,json=splitExternalLocalOriginErrors,proto3" json:"split_external_local_origin_errors,omitempty"`
	// The number of consecutive locally originated failures before ejection
	// occurs. Defaults to 5. Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_api.v3alpha.cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	ConsecutiveLocalOriginFailure *types.UInt32Value `protobuf:"bytes,13,opt,name=consecutive_local_origin_failure,json=consecutiveLocalOriginFailure,proto3" json:"consecutive_local_origin_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through consecutive locally originated failures. This setting can be
	// used to disable ejection or to ramp it up slowly. Defaults to 100.
	// Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_api.v3alpha.cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to true.
	EnforcingConsecutiveLocalOriginFailure *types.UInt32Value `protobuf:"bytes,14,opt,name=enforcing_consecutive_local_origin_failure,json=enforcingConsecutiveLocalOriginFailure,proto3" json:"enforcing_consecutive_local_origin_failure,omitempty"`
	// The % chance that a host will be actually ejected when an outlier status
	// is detected through success rate statistics for locally originated errors.
	// This setting can be used to disable ejection or to ramp it up slowly. Defaults to 100.
	// Parameter takes effect only when
	// :ref:`split_external_local_origin_errors<envoy_api_field_api.v3alpha.cluster.OutlierDetection.split_external_local_origin_errors>`
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
	return fileDescriptor_559fa0c48c84ccaa, []int{0}
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
	proto.RegisterType((*OutlierDetection)(nil), "envoy.api.v3alpha.cluster.OutlierDetection")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/cluster/outlier_detection.proto", fileDescriptor_559fa0c48c84ccaa)
}

var fileDescriptor_559fa0c48c84ccaa = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0x5d, 0x4f, 0xdb, 0x3a,
	0x18, 0xc7, 0x4f, 0x39, 0xbc, 0xf4, 0x98, 0x73, 0x80, 0xe3, 0x31, 0x70, 0x81, 0x31, 0xc6, 0xc4,
	0xc6, 0xb8, 0x48, 0x35, 0x10, 0xd7, 0x93, 0x3a, 0xca, 0x5e, 0xb4, 0x89, 0xae, 0x30, 0xd0, 0xc4,
	0x24, 0xcb, 0xa4, 0x4f, 0x53, 0x4f, 0x49, 0x9c, 0xd9, 0x4e, 0x09, 0x57, 0xd3, 0xbe, 0xce, 0x3e,
	0xc4, 0x3e, 0x18, 0x57, 0x53, 0xf3, 0xd2, 0x26, 0x6d, 0xd8, 0x92, 0xbb, 0x4a, 0x79, 0xfe, 0xbf,
	0xdf, 0xe3, 0xc7, 0x8e, 0x1b, 0xf4, 0x1c, 0xdc, 0xbe, 0xb8, 0xa9, 0x33, 0x8f, 0xd7, 0xfb, 0x07,
	0xcc, 0xf6, 0x7a, 0xac, 0x6e, 0xda, 0xbe, 0xd2, 0x20, 0xeb, 0xc2, 0xd7, 0x36, 0x07, 0x49, 0x3b,
	0xa0, 0xc1, 0xd4, 0x5c, 0xb8, 0x86, 0x27, 0x85, 0x16, 0xb8, 0x16, 0x46, 0x0c, 0xe6, 0x71, 0x23,
	0x8e, 0x18, 0x71, 0x64, 0x6d, 0xd3, 0x12, 0xc2, 0xb2, 0xa1, 0x1e, 0x16, 0x5e, 0xf9, 0xdd, 0x7a,
	0xc7, 0x97, 0x6c, 0x14, 0x9d, 0x7c, 0x7e, 0x2d, 0x99, 0xe7, 0x81, 0x54, 0xf1, 0xf3, 0xd5, 0x3e,
	0xb3, 0x79, 0x87, 0x69, 0xa8, 0x27, 0x3f, 0xa2, 0x07, 0xdb, 0x3f, 0x17, 0xd1, 0xd2, 0x49, 0xd4,
	0xcf, 0x51, 0xd2, 0x0e, 0x6e, 0xa2, 0x45, 0x53, 0xb8, 0x0a, 0x4c, 0x5f, 0xf3, 0x3e, 0xd0, 0xc3,
	0x20, 0x20, 0x95, 0xad, 0xca, 0xee, 0xfc, 0xfe, 0x86, 0x11, 0x79, 0x8c, 0xc4, 0x63, 0x7c, 0x7c,
	0xe3, 0xea, 0x83, 0xfd, 0x73, 0x66, 0xfb, 0xd0, 0x5e, 0x48, 0x85, 0x0e, 0x83, 0x00, 0xbf, 0x40,
	0x55, 0xee, 0x6a, 0x90, 0x7d, 0x66, 0x93, 0xa9, 0x30, 0x5f, 0x9b, 0xc8, 0x1f, 0xc5, 0xeb, 0x68,
	0x54, 0x6f, 0x1b, 0x33, 0x3f, 0x2a, 0x53, 0x7b, 0x7f, 0xb5, 0x87, 0x21, 0xfc, 0x01, 0xe1, 0x2b,
	0xa6, 0x80, 0xc2, 0x97, 0xa8, 0x31, 0xaa, 0xb9, 0x03, 0xe4, 0xef, 0xe2, 0xa8, 0xa5, 0x41, 0xbc,
	0x19, 0xa7, 0xcf, 0xb8, 0x03, 0xf8, 0x02, 0x2d, 0x3b, 0x2c, 0x18, 0x11, 0x3d, 0x90, 0x26, 0xb8,
	0x9a, 0x4c, 0xff, 0x79, 0x7d, 0x8d, 0xb9, 0xdb, 0xc6, 0xf4, 0xde, 0x14, 0xe9, 0xb4, 0xb1, 0xc3,
	0x82, 0x84, 0xda, 0x8a, 0x00, 0x98, 0xa1, 0x1a, 0xb8, 0x5d, 0x21, 0x4d, 0xee, 0x5a, 0x74, 0x7c,
	0x7a, 0x33, 0x65, 0xe8, 0xab, 0x43, 0xce, 0xcb, 0xec, 0x3c, 0x2f, 0xd1, 0xca, 0x48, 0xa1, 0x7c,
	0xd3, 0x04, 0xa5, 0xa8, 0x64, 0x1a, 0xc8, 0x6c, 0x19, 0xfe, 0xf2, 0x10, 0x72, 0x1a, 0x31, 0xda,
	0x4c, 0x03, 0xfe, 0x84, 0xd6, 0xd2, 0x48, 0xea, 0x70, 0x97, 0x3b, 0xbe, 0x43, 0x7b, 0x42, 0x69,
	0x45, 0xe6, 0x0a, 0x6c, 0xff, 0xaa, 0x1a, 0xe1, 0xde, 0x47, 0xe9, 0xd7, 0x83, 0x30, 0xbe, 0x44,
	0xeb, 0x19, 0xb4, 0x84, 0xaf, 0x3e, 0x28, 0x4d, 0xfb, 0xc2, 0xf6, 0x1d, 0x20, 0xd5, 0x02, 0x6c,
	0x92, 0x62, 0xb7, 0xa3, 0xf8, 0x79, 0x98, 0xc6, 0x17, 0xa8, 0x96, 0x81, 0x2b, 0xdd, 0x81, 0x3e,
	0xed, 0x32, 0x53, 0x0b, 0x49, 0xfe, 0x29, 0x80, 0x5e, 0x49, 0xa1, 0x4f, 0x07, 0xe1, 0xe3, 0x30,
	0x8b, 0x3f, 0xa3, 0xf5, 0xf4, 0x36, 0x5a, 0x4c, 0xc3, 0x35, 0xbb, 0xa1, 0x5d, 0xc6, 0x6d, 0x5f,
	0x02, 0x41, 0x05, 0xd0, 0xb5, 0x14, 0xe0, 0x55, 0x94, 0x3f, 0x8e, 0xe2, 0x38, 0x40, 0x3b, 0xf9,
	0xc7, 0x65, 0xdc, 0x33, 0x5f, 0x66, 0x6b, 0xb7, 0xf3, 0x8e, 0xce, 0x98, 0xf9, 0x2d, 0xda, 0x56,
	0x9e, 0xcd, 0x35, 0x85, 0x40, 0x83, 0x74, 0x99, 0x4d, 0x6d, 0x61, 0x32, 0x9b, 0x0a, 0xc9, 0x2d,
	0xee, 0x52, 0x90, 0x52, 0x48, 0x45, 0xfe, 0xdd, 0xaa, 0xec, 0x56, 0xdb, 0x9b, 0x61, 0x65, 0x33,
	0x2e, 0x7c, 0x37, 0xa8, 0x3b, 0x09, 0xcb, 0x9a, 0x61, 0x15, 0x06, 0xb4, 0x95, 0xee, 0x3d, 0x03,
	0x4a, 0x16, 0xf0, 0x5f, 0x81, 0x41, 0x3d, 0x48, 0x51, 0x52, 0x96, 0xa4, 0xe5, 0xef, 0x15, 0xb4,
	0x97, 0x3f, 0xad, 0x5c, 0xe3, 0x42, 0x99, 0x91, 0x3d, 0xc9, 0x1b, 0x59, 0x4e, 0x0f, 0x0a, 0x3d,
	0x1e, 0xb5, 0x90, 0xd1, 0x66, 0xde, 0xc4, 0xc5, 0x32, 0xee, 0x87, 0x43, 0x62, 0x4a, 0x98, 0x7e,
	0x29, 0x2d, 0xb4, 0x11, 0x2f, 0x2a, 0xb9, 0xa8, 0x98, 0x05, 0x54, 0xf7, 0x24, 0xa8, 0x9e, 0xb0,
	0x3b, 0x64, 0xa9, 0x8c, 0x6d, 0x2d, 0x46, 0xb5, 0x86, 0xa4, 0xb3, 0x04, 0x34, 0x10, 0x8d, 0x56,
	0x37, 0xa9, 0x24, 0xff, 0x97, 0x12, 0x0d, 0x51, 0xc7, 0xe3, 0x46, 0xfc, 0x0d, 0x3d, 0xfb, 0x9d,
	0x28, 0x33, 0x59, 0x82, 0xcb, 0x58, 0x77, 0xee, 0xb6, 0xa6, 0xa6, 0x3b, 0x38, 0xb2, 0x39, 0xda,
	0xec, 0x6d, 0x77, 0xaf, 0xc8, 0x91, 0x9d, 0x98, 0x66, 0xe6, 0xce, 0xb3, 0xd0, 0xa3, 0x1c, 0xcd,
	0xd8, 0xcd, 0xb7, 0x5c, 0xc0, 0xb3, 0x39, 0xe1, 0xc9, 0xdc, 0x7f, 0x8d, 0x06, 0x7a, 0xca, 0x85,
	0x11, 0x7e, 0x39, 0x78, 0x52, 0x04, 0x37, 0xc6, 0x9d, 0x1f, 0x11, 0x8d, 0xfb, 0xe3, 0x7f, 0xf4,
	0xad, 0x81, 0xb0, 0x55, 0xb9, 0x9a, 0x0d, 0xcd, 0x07, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfe,
	0x7d, 0xb1, 0x64, 0xb4, 0x08, 0x00, 0x00,
}
