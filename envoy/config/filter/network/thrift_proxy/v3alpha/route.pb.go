// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v3alpha/route.proto

package envoy_config_filter_network_thrift_proxy_v3alpha

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

type RouteConfiguration struct {
	// The name of the route configuration. Reserved for future use in asynchronous route discovery.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of routes that will be matched, in order, against incoming requests. The first route
	// that matches will be used.
	Routes               []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	// Route matching parameters.
	Match *RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Route request to some upstream cluster.
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	// Types that are valid to be assigned to MatchSpecifier:
	//	*RouteMatch_MethodName
	//	*RouteMatch_ServiceName
	MatchSpecifier isRouteMatch_MatchSpecifier `protobuf_oneof:"match_specifier"`
	// Inverts whatever matching is done in the :ref:`method_name
	// <envoy_api_field_config.filter.network.thrift_proxy.v3alpha.RouteMatch.method_name>` or
	// :ref:`service_name
	// <envoy_api_field_config.filter.network.thrift_proxy.v3alpha.RouteMatch.service_name>` fields.
	// Cannot be combined with wildcard matching as that would result in routes never being matched.
	//
	// .. note::
	//
	//   This does not invert matching done as part of the :ref:`headers field
	//   <envoy_api_field_config.filter.network.thrift_proxy.v3alpha.RouteMatch.headers>` field. To
	//   invert header matching, see :ref:`invert_match
	//   <envoy_api_field_api.v3alpha.route.HeaderMatcher.invert_match>`.
	Invert bool `protobuf:"varint,3,opt,name=invert,proto3" json:"invert,omitempty"`
	// Specifies a set of headers that the route should match on. The router will check the request’s
	// headers against all the specified headers in the route config. A match will happen if all the
	// headers in the route are present in the request with the same values (or based on presence if
	// the value field is not in the config). Note that this only applies for Thrift transports and/or
	// protocols that support headers.
	Headers              []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

type isRouteMatch_MatchSpecifier interface {
	isRouteMatch_MatchSpecifier()
}

type RouteMatch_MethodName struct {
	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3,oneof"`
}

type RouteMatch_ServiceName struct {
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3,oneof"`
}

func (*RouteMatch_MethodName) isRouteMatch_MatchSpecifier() {}

func (*RouteMatch_ServiceName) isRouteMatch_MatchSpecifier() {}

func (m *RouteMatch) GetMatchSpecifier() isRouteMatch_MatchSpecifier {
	if m != nil {
		return m.MatchSpecifier
	}
	return nil
}

func (m *RouteMatch) GetMethodName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_MethodName); ok {
		return x.MethodName
	}
	return ""
}

func (m *RouteMatch) GetServiceName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_ServiceName); ok {
		return x.ServiceName
	}
	return ""
}

func (m *RouteMatch) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

func (m *RouteMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteMatch) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteMatch_MethodName)(nil),
		(*RouteMatch_ServiceName)(nil),
	}
}

// [#next-free-field: 6]
type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	// Optional endpoint metadata match criteria used by the subset load balancer. Only endpoints in
	// the upstream cluster with metadata matching what is set in this field will be considered.
	// Note that this will be merged with what's provided in :ref: `WeightedCluster.MetadataMatch
	// <envoy_api_field_config.filter.network.thrift_proxy.v3alpha.WeightedCluster.ClusterWeight.metadata_match>`,
	// with values there taking precedence. Keys and values should be provided under the "envoy.lb"
	// metadata key.
	MetadataMatch *core.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	// Specifies a set of rate limit configurations that could be applied to the route.
	// N.B. Thrift service or method name matching can be achieved by specifying a RequestHeaders
	// action with the header name ":method-name".
	RateLimits []*route.RateLimit `protobuf:"bytes,4,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	// Strip the service prefix from the method name, if there's a prefix. For
	// example, the method call Service:method would end up being just method.
	StripServiceName     bool     `protobuf:"varint,5,opt,name=strip_service_name,json=stripServiceName,proto3" json:"strip_service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *RouteAction) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *RouteAction) GetRateLimits() []*route.RateLimit {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func (m *RouteAction) GetStripServiceName() bool {
	if m != nil {
		return m.StripServiceName
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
}

// Allows for specification of multiple upstream clusters along with weights that indicate the
// percentage of traffic to be forwarded to each cluster. The router selects an upstream cluster
// based on these weights.
type WeightedCluster struct {
	// Specifies one or more upstream clusters associated with the route.
	Clusters             []*WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *WeightedCluster) Reset()         { *m = WeightedCluster{} }
func (m *WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster) ProtoMessage()    {}
func (*WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{4}
}

func (m *WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster.Unmarshal(m, b)
}
func (m *WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster.Merge(m, src)
}
func (m *WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster.Size(m)
}
func (m *WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster proto.InternalMessageInfo

func (m *WeightedCluster) GetClusters() []*WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type WeightedCluster_ClusterWeight struct {
	// Name of the upstream cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When a request matches the route, the choice of an upstream cluster is determined by its
	// weight. The sum of weights across all entries in the clusters array determines the total
	// weight.
	Weight *types.UInt32Value `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	// Optional endpoint metadata match criteria used by the subset load balancer. Only endpoints in
	// the upstream cluster with metadata matching what is set in this field, combined with what's
	// provided in :ref: `RouteAction's metadata_match
	// <envoy_api_field_config.filter.network.thrift_proxy.v3alpha.RouteAction.metadata_match>`,
	// will be considered. Values here will take precedence. Keys and values should be provided
	// under the "envoy.lb" metadata key.
	MetadataMatch        *core.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WeightedCluster_ClusterWeight) Reset()         { *m = WeightedCluster_ClusterWeight{} }
func (m *WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{4, 0}
}

func (m *WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Size(m)
}
func (m *WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WeightedCluster_ClusterWeight) GetWeight() *types.UInt32Value {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *WeightedCluster_ClusterWeight) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.RouteAction")
	proto.RegisterType((*WeightedCluster)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.WeightedCluster")
	proto.RegisterType((*WeightedCluster_ClusterWeight)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v3alpha/route.proto", fileDescriptor_82b3b4ebd82e248c)
}

var fileDescriptor_82b3b4ebd82e248c = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x9e, 0xd3, 0xb5, 0x2b, 0xaf, 0x8c, 0x6d, 0x3e, 0x8c, 0xaa, 0x20, 0xd4, 0x75, 0x12, 0xaa,
	0x10, 0x72, 0xd0, 0x76, 0xe0, 0xb2, 0x21, 0x96, 0x1d, 0x28, 0x12, 0xfb, 0x21, 0x23, 0xe0, 0x02,
	0x8a, 0xbc, 0xd4, 0x6d, 0x2d, 0xd2, 0x38, 0x72, 0x9c, 0x96, 0xfd, 0x0b, 0x1c, 0xf9, 0x5b, 0xb8,
	0x73, 0x41, 0xfc, 0x4d, 0x68, 0x27, 0x14, 0xdb, 0xe9, 0xba, 0x01, 0x87, 0x0e, 0x4e, 0x71, 0x9e,
	0xdf, 0xf7, 0x7d, 0xef, 0x7d, 0xcf, 0x36, 0xec, 0xf1, 0x64, 0x22, 0xcf, 0xfd, 0x48, 0x26, 0x03,
	0x31, 0xf4, 0x07, 0x22, 0xd6, 0x5c, 0xf9, 0x09, 0xd7, 0x53, 0xa9, 0x3e, 0xfa, 0x7a, 0xa4, 0xc4,
	0x40, 0x87, 0xa9, 0x92, 0x9f, 0xce, 0xfd, 0xc9, 0x2e, 0x8b, 0xd3, 0x11, 0xf3, 0x95, 0xcc, 0x35,
	0x27, 0xa9, 0x92, 0x5a, 0xe2, 0x27, 0x06, 0x4d, 0x2c, 0x9a, 0x58, 0x34, 0x71, 0x68, 0x32, 0x8f,
	0x26, 0x0e, 0xdd, 0xda, 0xb2, 0x7a, 0x2c, 0x15, 0x33, 0xc2, 0x48, 0x2a, 0xee, 0x9f, 0xb1, 0xcc,
	0x91, 0xb6, 0xb6, 0x7f, 0x4f, 0x31, 0x9a, 0xf3, 0xca, 0xad, 0x07, 0x43, 0x29, 0x87, 0x31, 0xf7,
	0xcd, 0xdf, 0x59, 0x3e, 0xf0, 0xa7, 0x8a, 0xa5, 0x29, 0x57, 0x99, 0xdb, 0xbf, 0x3b, 0x61, 0xb1,
	0xe8, 0x33, 0xcd, 0xfd, 0x72, 0x61, 0x37, 0x3a, 0xe7, 0x80, 0x69, 0xc1, 0x73, 0x68, 0x6a, 0xce,
	0x15, 0xd3, 0x42, 0x26, 0x18, 0xc3, 0x72, 0xc2, 0xc6, 0xbc, 0x89, 0xda, 0xa8, 0x7b, 0x8b, 0x9a,
	0x35, 0x3e, 0x81, 0x9a, 0x51, 0xcc, 0x9a, 0x5e, 0xbb, 0xd2, 0x6d, 0xec, 0x3c, 0x25, 0x8b, 0x76,
	0x4b, 0x8c, 0x12, 0x75, 0x34, 0x9d, 0xef, 0x08, 0xaa, 0x26, 0x82, 0xdf, 0x43, 0x75, 0xcc, 0x74,
	0x34, 0x32, 0x7a, 0x8d, 0x9d, 0xbd, 0x1b, 0x32, 0x1f, 0x15, 0x1c, 0x41, 0xfd, 0x22, 0xa8, 0x7e,
	0x46, 0xde, 0x3a, 0xa2, 0x96, 0x14, 0x7f, 0x80, 0xaa, 0x51, 0x6c, 0x7a, 0x86, 0x7d, 0xff, 0x86,
	0xec, 0x07, 0x51, 0x61, 0xcd, 0x3c, 0xbd, 0x61, 0xed, 0xfc, 0x40, 0x00, 0x97, 0xf2, 0x78, 0x0b,
	0x1a, 0x63, 0xae, 0x47, 0xb2, 0x1f, 0x5e, 0x3a, 0xd8, 0x5b, 0xa2, 0x60, 0x83, 0xc7, 0x85, 0x93,
	0xdb, 0x70, 0x3b, 0xe3, 0x6a, 0x22, 0x22, 0x6e, 0x73, 0x3c, 0x97, 0xd3, 0x70, 0x51, 0x93, 0xb4,
	0x09, 0x35, 0x91, 0x4c, 0xb8, 0xd2, 0xcd, 0x4a, 0x1b, 0x75, 0xeb, 0xd4, 0xfd, 0xe1, 0xe7, 0xb0,
	0x32, 0xe2, 0xac, 0xcf, 0x55, 0xd6, 0x5c, 0x36, 0x73, 0x78, 0xe8, 0xfa, 0x61, 0xa9, 0x98, 0x15,
	0x6c, 0x8f, 0x46, 0xcf, 0xe4, 0x99, 0xb2, 0xb8, 0xa2, 0x25, 0x2c, 0xd8, 0x84, 0x35, 0x63, 0x4c,
	0x98, 0xa5, 0x3c, 0x12, 0x03, 0xc1, 0x15, 0xae, 0xfc, 0x0c, 0x50, 0xe7, 0xc2, 0x83, 0xc6, 0x5c,
	0xa7, 0x78, 0x1b, 0x56, 0xa2, 0x38, 0xcf, 0x34, 0x57, 0xb6, 0x8b, 0x60, 0xe5, 0x22, 0x58, 0x56,
	0x5e, 0x1b, 0xf5, 0x96, 0x68, 0xb9, 0x83, 0x53, 0xd8, 0x98, 0x72, 0x31, 0x1c, 0x69, 0xde, 0x0f,
	0x5d, 0x2c, 0x73, 0x46, 0x1f, 0x2c, 0x6e, 0xf4, 0x3b, 0x47, 0x75, 0x68, 0x99, 0x7a, 0x4b, 0x74,
	0x7d, 0x7a, 0x35, 0x94, 0xe1, 0x17, 0x70, 0x67, 0xcc, 0x35, 0xeb, 0x33, 0xcd, 0x42, 0x7b, 0x6a,
	0x2a, 0x46, 0xae, 0xfd, 0x07, 0x1f, 0x8a, 0xbb, 0x44, 0x8e, 0x5c, 0x36, 0x5d, 0x2d, 0x71, 0x76,
	0x52, 0x87, 0xd0, 0x50, 0x4c, 0xf3, 0x30, 0x16, 0x63, 0xa1, 0x4b, 0x37, 0x3b, 0x7f, 0x75, 0x93,
	0x32, 0xcd, 0x5f, 0x15, 0xa9, 0x14, 0x54, 0xb9, 0xcc, 0xf0, 0x63, 0xc0, 0x99, 0x56, 0x22, 0x0d,
	0xaf, 0x4c, 0xb4, 0x6a, 0x46, 0xb6, 0x6e, 0x76, 0x5e, 0x5f, 0x0e, 0x35, 0x68, 0xc2, 0x86, 0x33,
	0xe9, 0xba, 0xf9, 0xdf, 0x3c, 0x58, 0xbb, 0xd6, 0x3d, 0xce, 0xa1, 0x3e, 0xb3, 0x14, 0x99, 0xea,
	0x4e, 0xfe, 0xd9, 0x52, 0xe2, 0xbe, 0x36, 0x6c, 0x4e, 0xf3, 0x17, 0xe4, 0xd5, 0x11, 0x9d, 0x49,
	0xb5, 0xbe, 0x22, 0x58, 0xbd, 0x92, 0x85, 0xef, 0xcd, 0x3f, 0x07, 0xb3, 0x63, 0xe0, 0xde, 0x85,
	0x7d, 0xa8, 0xd9, 0x19, 0xb9, 0xb1, 0xdf, 0x27, 0xf6, 0x2d, 0x22, 0xe5, 0x5b, 0x44, 0xde, 0xbc,
	0x4c, 0xf4, 0xee, 0xce, 0x5b, 0x16, 0xe7, 0xdc, 0x80, 0x1f, 0x79, 0x5d, 0x44, 0x1d, 0xe8, 0xbf,
	0x8d, 0x33, 0x38, 0x86, 0x67, 0x42, 0x5a, 0x90, 0x75, 0x60, 0x51, 0xab, 0x02, 0x7b, 0x8d, 0x4f,
	0x8b, 0xb2, 0x4f, 0xd1, 0x59, 0xcd, 0xd4, 0xbf, 0xfb, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x8e,
	0xad, 0x2b, 0x13, 0x06, 0x00, 0x00,
}
