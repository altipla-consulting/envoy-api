// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/eds.proto

package envoy_api_v2

import (
	context "context"
	fmt "fmt"
	endpoint "github.com/altipla-consulting/envoy-api/envoy/api/v2/endpoint"
	_type "github.com/altipla-consulting/envoy-api/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/googleapis/google/api"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// Each route from RDS will map to a single cluster or traffic split across
// clusters using weights expressed in the RDS WeightedCluster.
//
// With EDS, each cluster is treated independently from a LB perspective, with
// LB taking place between the Localities within a cluster and at a finer
// granularity between the hosts within a locality. The percentage of traffic
// for each endpoint is determined by both its load_balancing_weight, and the
// load_balancing_weight of its locality. First, a locality will be selected,
// then an endpoint within that locality will be chose based on its weight.
// [#next-free-field: 6]
type ClusterLoadAssignment struct {
	// Name of the cluster. This will be the :ref:`service_name
	// <envoy_api_field_Cluster.EdsClusterConfig.service_name>` value if specified
	// in the cluster :ref:`EdsClusterConfig
	// <envoy_api_msg_Cluster.EdsClusterConfig>`.
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// List of endpoints to load balance to.
	Endpoints []*endpoint.LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// Map of named endpoints that can be referenced in LocalityLbEndpoints.
	NamedEndpoints map[string]*endpoint.Endpoint `protobuf:"bytes,5,rep,name=named_endpoints,json=namedEndpoints,proto3" json:"named_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Load balancing policy settings.
	Policy               *ClusterLoadAssignment_Policy `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ClusterLoadAssignment) Reset()         { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()    {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c8760f38742c17f, []int{0}
}

func (m *ClusterLoadAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment.Merge(m, src)
}
func (m *ClusterLoadAssignment) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment.Size(m)
}
func (m *ClusterLoadAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment proto.InternalMessageInfo

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*endpoint.LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetNamedEndpoints() map[string]*endpoint.Endpoint {
	if m != nil {
		return m.NamedEndpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

// Load balancing policy settings.
// [#next-free-field: 6]
type ClusterLoadAssignment_Policy struct {
	// Action to trim the overall incoming traffic to protect the upstream
	// hosts. This action allows protection in case the hosts are unable to
	// recover from an outage, or unable to autoscale or unable to handle
	// incoming traffic volume for any reason.
	//
	// At the client each category is applied one after the other to generate
	// the 'actual' drop percentage on all outgoing traffic. For example:
	//
	// .. code-block:: json
	//
	//  { "drop_overloads": [
	//      { "category": "throttle", "drop_percentage": 60 }
	//      { "category": "lb", "drop_percentage": 50 }
	//  ]}
	//
	// The actual drop percentages applied to the traffic at the clients will be
	//    "throttle"_drop = 60%
	//    "lb"_drop = 20%  // 50% of the remaining 'actual' load, which is 40%.
	//    actual_outgoing_load = 20% // remaining after applying all categories.
	DropOverloads []*ClusterLoadAssignment_Policy_DropOverload `protobuf:"bytes,2,rep,name=drop_overloads,json=dropOverloads,proto3" json:"drop_overloads,omitempty"`
	// Priority levels and localities are considered overprovisioned with this
	// factor (in percentage). This means that we don't consider a priority
	// level or locality unhealthy until the percentage of healthy hosts
	// multiplied by the overprovisioning factor drops below 100.
	// With the default value 140(1.4), Envoy doesn't consider a priority level
	// or a locality unhealthy until their percentage of healthy hosts drops
	// below 72%. For example:
	//
	// .. code-block:: json
	//
	//  { "overprovisioning_factor": 100 }
	//
	// Read more at :ref:`priority levels <arch_overview_load_balancing_priority_levels>` and
	// :ref:`localities <arch_overview_load_balancing_locality_weighted_lb>`.
	OverprovisioningFactor *types.UInt32Value `protobuf:"bytes,3,opt,name=overprovisioning_factor,json=overprovisioningFactor,proto3" json:"overprovisioning_factor,omitempty"`
	// The max time until which the endpoints from this assignment can be used.
	// If no new assignments are received before this time expires the endpoints
	// are considered stale and should be marked unhealthy.
	// Defaults to 0 which means endpoints never go stale.
	EndpointStaleAfter *types.Duration `protobuf:"bytes,4,opt,name=endpoint_stale_after,json=endpointStaleAfter,proto3" json:"endpoint_stale_after,omitempty"`
	// The flag to disable overprovisioning. If it is set to true,
	// :ref:`overprovisioning factor
	// <arch_overview_load_balancing_overprovisioning_factor>` will be ignored
	// and Envoy will not perform graceful failover between priority levels or
	// localities as endpoints become unhealthy. Otherwise Envoy will perform
	// graceful failover as :ref:`overprovisioning factor
	// <arch_overview_load_balancing_overprovisioning_factor>` suggests.
	// [#next-major-version: Unify with overprovisioning config as a single message.]
	// [#not-implemented-hide:]
	DisableOverprovisioning bool     `protobuf:"varint,5,opt,name=disable_overprovisioning,json=disableOverprovisioning,proto3" json:"disable_overprovisioning,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ClusterLoadAssignment_Policy) Reset()         { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c8760f38742c17f, []int{0, 0}
}

func (m *ClusterLoadAssignment_Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Size(m)
}
func (m *ClusterLoadAssignment_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy) GetDropOverloads() []*ClusterLoadAssignment_Policy_DropOverload {
	if m != nil {
		return m.DropOverloads
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetOverprovisioningFactor() *types.UInt32Value {
	if m != nil {
		return m.OverprovisioningFactor
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetEndpointStaleAfter() *types.Duration {
	if m != nil {
		return m.EndpointStaleAfter
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetDisableOverprovisioning() bool {
	if m != nil {
		return m.DisableOverprovisioning
	}
	return false
}

type ClusterLoadAssignment_Policy_DropOverload struct {
	// Identifier for the policy specifying the drop.
	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	// Percentage of traffic that should be dropped for the category.
	DropPercentage       *_type.FractionalPercent `protobuf:"bytes,2,opt,name=drop_percentage,json=dropPercentage,proto3" json:"drop_percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ClusterLoadAssignment_Policy_DropOverload) Reset() {
	*m = ClusterLoadAssignment_Policy_DropOverload{}
}
func (m *ClusterLoadAssignment_Policy_DropOverload) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy_DropOverload) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy_DropOverload) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c8760f38742c17f, []int{0, 0, 0}
}

func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Size(m)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy_DropOverload) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ClusterLoadAssignment_Policy_DropOverload) GetDropPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.DropPercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.api.v2.ClusterLoadAssignment")
	proto.RegisterMapType((map[string]*endpoint.Endpoint)(nil), "envoy.api.v2.ClusterLoadAssignment.NamedEndpointsEntry")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy")
	proto.RegisterType((*ClusterLoadAssignment_Policy_DropOverload)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy.DropOverload")
}

func init() { proto.RegisterFile("envoy/api/v2/eds.proto", fileDescriptor_3c8760f38742c17f) }

var fileDescriptor_3c8760f38742c17f = []byte{
	// 702 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6e, 0xd3, 0x4a,
	0x14, 0xc6, 0x6b, 0xa7, 0xc9, 0x4d, 0xa7, 0xbd, 0x69, 0x35, 0xf7, 0xde, 0xc6, 0xd7, 0x0a, 0x6d,
	0x94, 0x76, 0x11, 0x05, 0xc9, 0x41, 0xa9, 0x50, 0xa1, 0xbb, 0x86, 0x34, 0x02, 0x54, 0xd1, 0xc8,
	0x55, 0x11, 0xab, 0x9a, 0x89, 0x7d, 0x1a, 0x46, 0x38, 0x33, 0x66, 0x3c, 0x31, 0x58, 0xec, 0x58,
	0xb1, 0xe7, 0x2d, 0x78, 0x18, 0x36, 0xac, 0xd9, 0xf1, 0x0e, 0x48, 0x5d, 0x21, 0xff, 0x4d, 0x9d,
	0xb6, 0x88, 0x05, 0xbb, 0xf1, 0x7c, 0xe7, 0xfb, 0xcd, 0x99, 0x73, 0x8e, 0x07, 0x6d, 0x02, 0x0b,
	0x78, 0xd8, 0x25, 0x1e, 0xed, 0x06, 0xbd, 0x2e, 0x38, 0xbe, 0xe1, 0x09, 0x2e, 0x39, 0x5e, 0x8b,
	0xf7, 0x0d, 0xe2, 0x51, 0x23, 0xe8, 0xe9, 0x8d, 0x42, 0x94, 0x43, 0x7d, 0x9b, 0x07, 0x20, 0xc2,
	0x24, 0x56, 0xdf, 0x2d, 0x32, 0x98, 0xe3, 0x71, 0xca, 0x64, 0xbe, 0x48, 0xa3, 0xb4, 0x24, 0x4a,
	0x86, 0x1e, 0x74, 0x3d, 0x10, 0x36, 0xe4, 0x4a, 0x63, 0xc2, 0xf9, 0xc4, 0x85, 0x18, 0x40, 0x18,
	0xe3, 0x92, 0x48, 0xca, 0x59, 0x9a, 0x89, 0xbe, 0x95, 0xaa, 0xf1, 0xd7, 0x78, 0x76, 0xd1, 0x75,
	0x66, 0x22, 0x0e, 0xb8, 0x4d, 0x7f, 0x2b, 0x88, 0xe7, 0x81, 0xc8, 0xfc, 0xf5, 0x80, 0xb8, 0xd4,
	0x21, 0x12, 0xba, 0xd9, 0x22, 0x11, 0x5a, 0x3f, 0x2a, 0xe8, 0xbf, 0x47, 0xee, 0xcc, 0x97, 0x20,
	0x8e, 0x39, 0x71, 0x0e, 0x7d, 0x9f, 0x4e, 0xd8, 0x14, 0x98, 0xc4, 0x1d, 0xb4, 0x66, 0x27, 0x82,
	0xc5, 0xc8, 0x14, 0x34, 0xa5, 0xa9, 0xb4, 0x57, 0xfa, 0x7f, 0x5d, 0xf6, 0x97, 0x85, 0xda, 0x54,
	0xcc, 0xd5, 0x54, 0x7c, 0x46, 0xa6, 0x80, 0x1f, 0xa3, 0x95, 0xec, 0xa2, 0xbe, 0xa6, 0x36, 0x4b,
	0xed, 0xd5, 0x5e, 0xc7, 0xb8, 0x5a, 0x3c, 0x23, 0xaf, 0xc3, 0x31, 0xb7, 0x89, 0x4b, 0x65, 0x78,
	0x3c, 0x3e, 0xca, 0x1c, 0xe6, 0xdc, 0x8c, 0x5f, 0xa2, 0xf5, 0xe8, 0x34, 0xc7, 0x9a, 0xf3, 0xca,
	0x31, 0x6f, 0xbf, 0xc8, 0xbb, 0x31, 0x67, 0x23, 0x4a, 0xc6, 0xc9, 0xb9, 0x47, 0x4c, 0x8a, 0xd0,
	0xac, 0xb1, 0xc2, 0x26, 0xee, 0xa3, 0x8a, 0xc7, 0x5d, 0x6a, 0x87, 0xda, 0x72, 0x53, 0xb9, 0x9e,
	0xe8, 0xcd, 0xe0, 0x51, 0xec, 0x30, 0x53, 0xa7, 0xfe, 0xad, 0x84, 0x2a, 0xc9, 0x16, 0x3e, 0x47,
	0x35, 0x47, 0x70, 0xcf, 0x8a, 0x66, 0xc1, 0xe5, 0xc4, 0xc9, 0xee, 0xbf, 0xff, 0xfb, 0x58, 0x63,
	0x20, 0xb8, 0x77, 0x92, 0xfa, 0xcd, 0xbf, 0x9d, 0x2b, 0x5f, 0x3e, 0x3e, 0x47, 0xf5, 0x08, 0xed,
	0x09, 0x1e, 0x50, 0x9f, 0x72, 0x46, 0xd9, 0xc4, 0xba, 0x20, 0xb6, 0xe4, 0x42, 0x2b, 0xc5, 0xf9,
	0x37, 0x8c, 0xa4, 0xf7, 0x46, 0xd6, 0x7b, 0xe3, 0xec, 0x09, 0x93, 0x7b, 0xbd, 0xe7, 0xc4, 0x9d,
	0x41, 0xdc, 0xaf, 0x8e, 0xda, 0x5c, 0x32, 0x37, 0x17, 0x29, 0xc3, 0x18, 0x82, 0xcf, 0xd0, 0xbf,
	0x59, 0xa9, 0x2d, 0x5f, 0x12, 0x17, 0x2c, 0x72, 0x21, 0x41, 0xa4, 0xc5, 0xf9, 0xff, 0x1a, 0x7c,
	0x90, 0x0e, 0x5e, 0xbf, 0x7a, 0xd9, 0x2f, 0x7f, 0x56, 0xd4, 0xce, 0x92, 0x89, 0x33, 0xc0, 0x69,
	0xe4, 0x3f, 0x8c, 0xec, 0xf8, 0x21, 0xd2, 0x1c, 0xea, 0x93, 0xb1, 0x0b, 0xd6, 0xe2, 0xc1, 0x5a,
	0xb9, 0xa9, 0xb4, 0xab, 0x66, 0x3d, 0xd5, 0x4f, 0x16, 0x64, 0xfd, 0x3d, 0x5a, 0xbb, 0x5a, 0x10,
	0xbc, 0x83, 0xaa, 0x36, 0x91, 0x30, 0xe1, 0x22, 0x5c, 0x1c, 0xc2, 0x5c, 0xc0, 0x43, 0xb4, 0x1e,
	0xb7, 0x21, 0xfd, 0xa9, 0xc8, 0x04, 0x34, 0x35, 0xbe, 0xc1, 0x9d, 0xb4, 0x0f, 0xd1, 0x2f, 0x67,
	0x0c, 0x05, 0xb1, 0xa3, 0xe4, 0x89, 0x3b, 0x4a, 0xe2, 0xcc, 0xb8, 0x79, 0xa3, 0xdc, 0xf4, 0x74,
	0xb9, 0xaa, 0x6c, 0xa8, 0xfa, 0x18, 0xfd, 0x73, 0xc3, 0x28, 0xe1, 0x0d, 0x54, 0x7a, 0x0d, 0x69,
	0x12, 0x66, 0xb4, 0xc4, 0xf7, 0x51, 0x39, 0x88, 0xea, 0x9c, 0x1e, 0xb6, 0x7d, 0xcb, 0xd0, 0x67,
	0x1c, 0x33, 0x89, 0x3e, 0x50, 0x1f, 0x28, 0xbd, 0x2f, 0x2a, 0xd2, 0xb2, 0xfd, 0x41, 0xf6, 0x98,
	0x9c, 0x82, 0x08, 0xa8, 0x0d, 0xf8, 0x05, 0x5a, 0x3f, 0x95, 0x02, 0xc8, 0x74, 0x3e, 0xb7, 0x5b,
	0x45, 0x76, 0x6e, 0x31, 0xe1, 0xcd, 0x0c, 0x7c, 0xa9, 0x6f, 0xdf, 0xaa, 0xfb, 0x1e, 0x67, 0x3e,
	0xb4, 0x96, 0xda, 0xca, 0x3d, 0x05, 0x13, 0x54, 0x1b, 0x80, 0x2b, 0xc9, 0x1c, 0xbc, 0xb3, 0x60,
	0x8c, 0xd4, 0x6b, 0xf4, 0xdd, 0x5f, 0x07, 0x15, 0x8e, 0x98, 0xa1, 0xda, 0x10, 0xa4, 0xfd, 0xea,
	0x0f, 0xe6, 0xde, 0xfa, 0xf0, 0xf5, 0xfb, 0x27, 0xb5, 0xd1, 0xaa, 0x17, 0x9e, 0xde, 0x83, 0xfc,
	0x91, 0x38, 0x50, 0x3a, 0xfd, 0xbb, 0x48, 0xa7, 0x3c, 0x01, 0x79, 0x82, 0xbf, 0x0b, 0x0b, 0xcc,
	0x7e, 0xf5, 0xc8, 0xf1, 0x47, 0xd1, 0x10, 0x8f, 0x94, 0x8f, 0x8a, 0x32, 0xae, 0xc4, 0x03, 0xbd,
	0xf7, 0x33, 0x00, 0x00, 0xff, 0xff, 0x67, 0x4f, 0x40, 0xb4, 0xfb, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndpointDiscoveryServiceClient is the client API for EndpointDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointDiscoveryServiceClient interface {
	// The resource_names field in DiscoveryRequest specifies a list of clusters
	// to subscribe to updates for.
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.EndpointDiscoveryService/DeltaEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceDeltaEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_DeltaEndpointsClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceDeltaEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointDiscoveryServiceServer is the server API for EndpointDiscoveryService service.
type EndpointDiscoveryServiceServer interface {
	// The resource_names field in DiscoveryRequest specifies a list of clusters
	// to subscribe to updates for.
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	DeltaEndpoints(EndpointDiscoveryService_DeltaEndpointsServer) error
	FetchEndpoints(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_DeltaEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).DeltaEndpoints(&endpointDiscoveryServiceDeltaEndpointsServer{stream})
}

type EndpointDiscoveryService_DeltaEndpointsServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceDeltaEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaEndpoints",
			Handler:       _EndpointDiscoveryService_DeltaEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/eds.proto",
}
