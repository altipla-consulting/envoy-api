// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/rbac/v2/rbac.proto

package v2

import (
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v2/core"
	route "github.com/altipla-consulting/envoy-api/envoy/api/v2/route"
	matcher "github.com/altipla-consulting/envoy-api/envoy/type/matcher"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Should we do safe-list or block-list style access control?
type RBAC_Action int32

const (
	// The policies grant access to principals. The rest is denied. This is safe-list style
	// access control. This is the default type.
	RBAC_ALLOW RBAC_Action = 0
	// The policies deny access to principals. The rest is allowed. This is block-list style
	// access control.
	RBAC_DENY RBAC_Action = 1
)

var RBAC_Action_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}

var RBAC_Action_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x RBAC_Action) String() string {
	return proto.EnumName(RBAC_Action_name, int32(x))
}

func (RBAC_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{0, 0}
}

// Role Based Access Control (RBAC) provides service-level and method-level access control for a
// service. RBAC policies are additive. The policies are examined in order. A request is allowed
// once a matching policy is found (suppose the `action` is ALLOW).
//
// Here is an example of RBAC configuration. It has two policies:
//
// * Service account "cluster.local/ns/default/sa/admin" has full access to the service, and so
//   does "cluster.local/ns/default/sa/superuser".
//
// * Any user can read ("GET") the service at paths with prefix "/products", so long as the
//   destination port is either 80 or 443.
//
//  .. code-block:: yaml
//
//   action: ALLOW
//   policies:
//     "service-admin":
//       permissions:
//         - any: true
//       principals:
//         - authenticated:
//             principal_name:
//               exact: "cluster.local/ns/default/sa/admin"
//         - authenticated:
//             principal_name:
//               exact: "cluster.local/ns/default/sa/superuser"
//     "product-viewer":
//       permissions:
//           - and_rules:
//               rules:
//                 - header: { name: ":method", exact_match: "GET" }
//                 - header: { name: ":path", regex_match: "/products(/.*)?" }
//                 - or_rules:
//                     rules:
//                       - destination_port: 80
//                       - destination_port: 443
//       principals:
//         - any: true
//
type RBAC struct {
	// The action to take if a policy matches. The request is allowed if and only if:
	//
	//   * `action` is "ALLOWED" and at least one policy matches
	//   * `action` is "DENY" and none of the policies match
	Action RBAC_Action `protobuf:"varint,1,opt,name=action,proto3,enum=envoy.config.rbac.v2.RBAC_Action" json:"action,omitempty"`
	// Maps from policy name to policy. A match occurs when at least one policy matches the request.
	Policies             map[string]*Policy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetAction() RBAC_Action {
	if m != nil {
		return m.Action
	}
	return RBAC_ALLOW
}

func (m *RBAC) GetPolicies() map[string]*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

// Policy specifies a role and the principals that are assigned/denied the role. A policy matches if
// and only if at least one of its permissions match the action taking place AND at least one of its
// principals match the downstream.
type Policy struct {
	// Required. The set of permissions that define a role. Each permission is matched with OR
	// semantics. To match all actions for this policy, a single Permission with the `any` field set
	// to true should be used.
	Permissions []*Permission `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// Required. The set of principals that are assigned/denied the role based on “action”. Each
	// principal is matched with OR semantics. To match all downstreams for this policy, a single
	// Principal with the `any` field set to true should be used.
	Principals           []*Principal `protobuf:"bytes,2,rep,name=principals,proto3" json:"principals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{1}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Policy) GetPrincipals() []*Principal {
	if m != nil {
		return m.Principals
	}
	return nil
}

// Permission defines an action (or actions) that a principal can take.
type Permission struct {
	// Types that are valid to be assigned to Rule:
	//	*Permission_AndRules
	//	*Permission_OrRules
	//	*Permission_Any
	//	*Permission_Header
	//	*Permission_DestinationIp
	//	*Permission_DestinationPort
	//	*Permission_Metadata
	//	*Permission_NotRule
	//	*Permission_RequestedServerName
	Rule                 isPermission_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{2}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

type isPermission_Rule interface {
	isPermission_Rule()
}

type Permission_AndRules struct {
	AndRules *Permission_Set `protobuf:"bytes,1,opt,name=and_rules,json=andRules,proto3,oneof"`
}

type Permission_OrRules struct {
	OrRules *Permission_Set `protobuf:"bytes,2,opt,name=or_rules,json=orRules,proto3,oneof"`
}

type Permission_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type Permission_Header struct {
	Header *route.HeaderMatcher `protobuf:"bytes,4,opt,name=header,proto3,oneof"`
}

type Permission_DestinationIp struct {
	DestinationIp *core.CidrRange `protobuf:"bytes,5,opt,name=destination_ip,json=destinationIp,proto3,oneof"`
}

type Permission_DestinationPort struct {
	DestinationPort uint32 `protobuf:"varint,6,opt,name=destination_port,json=destinationPort,proto3,oneof"`
}

type Permission_Metadata struct {
	Metadata *matcher.MetadataMatcher `protobuf:"bytes,7,opt,name=metadata,proto3,oneof"`
}

type Permission_NotRule struct {
	NotRule *Permission `protobuf:"bytes,8,opt,name=not_rule,json=notRule,proto3,oneof"`
}

type Permission_RequestedServerName struct {
	RequestedServerName *matcher.StringMatcher `protobuf:"bytes,9,opt,name=requested_server_name,json=requestedServerName,proto3,oneof"`
}

func (*Permission_AndRules) isPermission_Rule() {}

func (*Permission_OrRules) isPermission_Rule() {}

func (*Permission_Any) isPermission_Rule() {}

func (*Permission_Header) isPermission_Rule() {}

func (*Permission_DestinationIp) isPermission_Rule() {}

func (*Permission_DestinationPort) isPermission_Rule() {}

func (*Permission_Metadata) isPermission_Rule() {}

func (*Permission_NotRule) isPermission_Rule() {}

func (*Permission_RequestedServerName) isPermission_Rule() {}

func (m *Permission) GetRule() isPermission_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *Permission) GetAndRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_AndRules); ok {
		return x.AndRules
	}
	return nil
}

func (m *Permission) GetOrRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_OrRules); ok {
		return x.OrRules
	}
	return nil
}

func (m *Permission) GetAny() bool {
	if x, ok := m.GetRule().(*Permission_Any); ok {
		return x.Any
	}
	return false
}

func (m *Permission) GetHeader() *route.HeaderMatcher {
	if x, ok := m.GetRule().(*Permission_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Permission) GetDestinationIp() *core.CidrRange {
	if x, ok := m.GetRule().(*Permission_DestinationIp); ok {
		return x.DestinationIp
	}
	return nil
}

func (m *Permission) GetDestinationPort() uint32 {
	if x, ok := m.GetRule().(*Permission_DestinationPort); ok {
		return x.DestinationPort
	}
	return 0
}

func (m *Permission) GetMetadata() *matcher.MetadataMatcher {
	if x, ok := m.GetRule().(*Permission_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Permission) GetNotRule() *Permission {
	if x, ok := m.GetRule().(*Permission_NotRule); ok {
		return x.NotRule
	}
	return nil
}

func (m *Permission) GetRequestedServerName() *matcher.StringMatcher {
	if x, ok := m.GetRule().(*Permission_RequestedServerName); ok {
		return x.RequestedServerName
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Permission) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Permission_AndRules)(nil),
		(*Permission_OrRules)(nil),
		(*Permission_Any)(nil),
		(*Permission_Header)(nil),
		(*Permission_DestinationIp)(nil),
		(*Permission_DestinationPort)(nil),
		(*Permission_Metadata)(nil),
		(*Permission_NotRule)(nil),
		(*Permission_RequestedServerName)(nil),
	}
}

// Used in the `and_rules` and `or_rules` fields in the `rule` oneof. Depending on the context,
// each are applied with the associated behavior.
type Permission_Set struct {
	Rules                []*Permission `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Permission_Set) Reset()         { *m = Permission_Set{} }
func (m *Permission_Set) String() string { return proto.CompactTextString(m) }
func (*Permission_Set) ProtoMessage()    {}
func (*Permission_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{2, 0}
}

func (m *Permission_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission_Set.Unmarshal(m, b)
}
func (m *Permission_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission_Set.Marshal(b, m, deterministic)
}
func (m *Permission_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission_Set.Merge(m, src)
}
func (m *Permission_Set) XXX_Size() int {
	return xxx_messageInfo_Permission_Set.Size(m)
}
func (m *Permission_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Permission_Set proto.InternalMessageInfo

func (m *Permission_Set) GetRules() []*Permission {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Principal defines an identity or a group of identities for a downstream subject.
type Principal struct {
	// Types that are valid to be assigned to Identifier:
	//	*Principal_AndIds
	//	*Principal_OrIds
	//	*Principal_Any
	//	*Principal_Authenticated_
	//	*Principal_SourceIp
	//	*Principal_Header
	//	*Principal_Metadata
	//	*Principal_NotId
	Identifier           isPrincipal_Identifier `protobuf_oneof:"identifier"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{3}
}

func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (m *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(m, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

type isPrincipal_Identifier interface {
	isPrincipal_Identifier()
}

type Principal_AndIds struct {
	AndIds *Principal_Set `protobuf:"bytes,1,opt,name=and_ids,json=andIds,proto3,oneof"`
}

type Principal_OrIds struct {
	OrIds *Principal_Set `protobuf:"bytes,2,opt,name=or_ids,json=orIds,proto3,oneof"`
}

type Principal_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type Principal_Authenticated_ struct {
	Authenticated *Principal_Authenticated `protobuf:"bytes,4,opt,name=authenticated,proto3,oneof"`
}

type Principal_SourceIp struct {
	SourceIp *core.CidrRange `protobuf:"bytes,5,opt,name=source_ip,json=sourceIp,proto3,oneof"`
}

type Principal_Header struct {
	Header *route.HeaderMatcher `protobuf:"bytes,6,opt,name=header,proto3,oneof"`
}

type Principal_Metadata struct {
	Metadata *matcher.MetadataMatcher `protobuf:"bytes,7,opt,name=metadata,proto3,oneof"`
}

type Principal_NotId struct {
	NotId *Principal `protobuf:"bytes,8,opt,name=not_id,json=notId,proto3,oneof"`
}

func (*Principal_AndIds) isPrincipal_Identifier() {}

func (*Principal_OrIds) isPrincipal_Identifier() {}

func (*Principal_Any) isPrincipal_Identifier() {}

func (*Principal_Authenticated_) isPrincipal_Identifier() {}

func (*Principal_SourceIp) isPrincipal_Identifier() {}

func (*Principal_Header) isPrincipal_Identifier() {}

func (*Principal_Metadata) isPrincipal_Identifier() {}

func (*Principal_NotId) isPrincipal_Identifier() {}

func (m *Principal) GetIdentifier() isPrincipal_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Principal) GetAndIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_AndIds); ok {
		return x.AndIds
	}
	return nil
}

func (m *Principal) GetOrIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_OrIds); ok {
		return x.OrIds
	}
	return nil
}

func (m *Principal) GetAny() bool {
	if x, ok := m.GetIdentifier().(*Principal_Any); ok {
		return x.Any
	}
	return false
}

func (m *Principal) GetAuthenticated() *Principal_Authenticated {
	if x, ok := m.GetIdentifier().(*Principal_Authenticated_); ok {
		return x.Authenticated
	}
	return nil
}

func (m *Principal) GetSourceIp() *core.CidrRange {
	if x, ok := m.GetIdentifier().(*Principal_SourceIp); ok {
		return x.SourceIp
	}
	return nil
}

func (m *Principal) GetHeader() *route.HeaderMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Principal) GetMetadata() *matcher.MetadataMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Principal) GetNotId() *Principal {
	if x, ok := m.GetIdentifier().(*Principal_NotId); ok {
		return x.NotId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Principal) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Principal_AndIds)(nil),
		(*Principal_OrIds)(nil),
		(*Principal_Any)(nil),
		(*Principal_Authenticated_)(nil),
		(*Principal_SourceIp)(nil),
		(*Principal_Header)(nil),
		(*Principal_Metadata)(nil),
		(*Principal_NotId)(nil),
	}
}

// Used in the `and_ids` and `or_ids` fields in the `identifier` oneof. Depending on the context,
// each are applied with the associated behavior.
type Principal_Set struct {
	Ids                  []*Principal `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Principal_Set) Reset()         { *m = Principal_Set{} }
func (m *Principal_Set) String() string { return proto.CompactTextString(m) }
func (*Principal_Set) ProtoMessage()    {}
func (*Principal_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{3, 0}
}

func (m *Principal_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Set.Unmarshal(m, b)
}
func (m *Principal_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Set.Marshal(b, m, deterministic)
}
func (m *Principal_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Set.Merge(m, src)
}
func (m *Principal_Set) XXX_Size() int {
	return xxx_messageInfo_Principal_Set.Size(m)
}
func (m *Principal_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Set proto.InternalMessageInfo

func (m *Principal_Set) GetIds() []*Principal {
	if m != nil {
		return m.Ids
	}
	return nil
}

// Authentication attributes for a downstream.
type Principal_Authenticated struct {
	// The name of the principal. If set, The URI SAN is used from the certificate, otherwise the
	// subject field is used. If unset, it applies to any user that is authenticated.
	PrincipalName        *matcher.StringMatcher `protobuf:"bytes,2,opt,name=principal_name,json=principalName,proto3" json:"principal_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Principal_Authenticated) Reset()         { *m = Principal_Authenticated{} }
func (m *Principal_Authenticated) String() string { return proto.CompactTextString(m) }
func (*Principal_Authenticated) ProtoMessage()    {}
func (*Principal_Authenticated) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{3, 1}
}

func (m *Principal_Authenticated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Authenticated.Unmarshal(m, b)
}
func (m *Principal_Authenticated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Authenticated.Marshal(b, m, deterministic)
}
func (m *Principal_Authenticated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Authenticated.Merge(m, src)
}
func (m *Principal_Authenticated) XXX_Size() int {
	return xxx_messageInfo_Principal_Authenticated.Size(m)
}
func (m *Principal_Authenticated) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Authenticated.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Authenticated proto.InternalMessageInfo

func (m *Principal_Authenticated) GetPrincipalName() *matcher.StringMatcher {
	if m != nil {
		return m.PrincipalName
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.rbac.v2.RBAC_Action", RBAC_Action_name, RBAC_Action_value)
	proto.RegisterType((*RBAC)(nil), "envoy.config.rbac.v2.RBAC")
	proto.RegisterMapType((map[string]*Policy)(nil), "envoy.config.rbac.v2.RBAC.PoliciesEntry")
	proto.RegisterType((*Policy)(nil), "envoy.config.rbac.v2.Policy")
	proto.RegisterType((*Permission)(nil), "envoy.config.rbac.v2.Permission")
	proto.RegisterType((*Permission_Set)(nil), "envoy.config.rbac.v2.Permission.Set")
	proto.RegisterType((*Principal)(nil), "envoy.config.rbac.v2.Principal")
	proto.RegisterType((*Principal_Set)(nil), "envoy.config.rbac.v2.Principal.Set")
	proto.RegisterType((*Principal_Authenticated)(nil), "envoy.config.rbac.v2.Principal.Authenticated")
}

func init() { proto.RegisterFile("envoy/config/rbac/v2/rbac.proto", fileDescriptor_e8a2b527e1e731e1) }

var fileDescriptor_e8a2b527e1e731e1 = []byte{
	// 855 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xd1, 0x8e, 0xdb, 0x44,
	0x14, 0x86, 0x63, 0x27, 0xf1, 0x3a, 0x27, 0xca, 0x12, 0x86, 0x22, 0xac, 0x00, 0x6d, 0x9a, 0x72,
	0x11, 0x21, 0xe1, 0x48, 0x41, 0x42, 0x2d, 0x0b, 0x88, 0x78, 0xbb, 0x22, 0x41, 0xdb, 0xb2, 0x9a,
	0x15, 0xaa, 0xca, 0xcd, 0x6a, 0xd6, 0x33, 0xcd, 0x0e, 0x24, 0x33, 0x66, 0x3c, 0x89, 0xc8, 0x25,
	0xaf, 0xc0, 0x03, 0x20, 0xf1, 0x0a, 0x3c, 0x0d, 0xd7, 0x3c, 0x05, 0xda, 0x9b, 0xa2, 0x99, 0x71,
	0xb2, 0xb1, 0xb4, 0x6d, 0x1a, 0xa9, 0x37, 0xc9, 0xc8, 0xf3, 0xff, 0xff, 0x19, 0x1f, 0x7f, 0xc7,
	0x86, 0x7b, 0x4c, 0x2c, 0xe5, 0x6a, 0x90, 0x4a, 0xf1, 0x82, 0x4f, 0x07, 0xea, 0x92, 0xa4, 0x83,
	0xe5, 0xd0, 0xfe, 0xc7, 0x99, 0x92, 0x5a, 0xa2, 0x3b, 0x56, 0x10, 0x3b, 0x41, 0x6c, 0x37, 0x96,
	0xc3, 0xce, 0x07, 0x4b, 0x32, 0xe3, 0x94, 0x68, 0x36, 0x58, 0x2f, 0x9c, 0xbc, 0x73, 0x67, 0x2a,
	0xa7, 0xd2, 0x2e, 0x07, 0x66, 0x55, 0x5c, 0x2d, 0xaa, 0x90, 0x8c, 0x9b, 0xf4, 0x54, 0x2a, 0x36,
	0x20, 0x94, 0x2a, 0x96, 0xe7, 0x85, 0xe0, 0x6e, 0x49, 0xa0, 0xe4, 0x42, 0x33, 0xf7, 0x5b, 0xec,
	0xdf, 0x77, 0xfb, 0x7a, 0x95, 0xb1, 0xc1, 0x9c, 0xe8, 0xf4, 0x8a, 0xa9, 0xc1, 0x9c, 0x69, 0x42,
	0x89, 0x26, 0xe5, 0x1a, 0x25, 0x49, 0xae, 0x15, 0x17, 0x53, 0x27, 0xe8, 0xfd, 0xee, 0x43, 0x0d,
	0x27, 0xa3, 0x63, 0xf4, 0x08, 0x02, 0x92, 0x6a, 0x2e, 0x45, 0xe4, 0x75, 0xbd, 0xfe, 0xe1, 0xf0,
	0x7e, 0x7c, 0xdb, 0x3d, 0xc6, 0x46, 0x1b, 0x8f, 0xac, 0x10, 0x17, 0x06, 0xf4, 0x18, 0xc2, 0x4c,
	0xce, 0x78, 0xca, 0x59, 0x1e, 0xf9, 0xdd, 0x6a, 0xbf, 0x39, 0xec, 0xbf, 0xc6, 0x7c, 0x56, 0x48,
	0x4f, 0x84, 0x56, 0x2b, 0xbc, 0x71, 0x76, 0x9e, 0x43, 0xab, 0xb4, 0x85, 0xda, 0x50, 0xfd, 0x85,
	0xad, 0xec, 0x71, 0x1a, 0xd8, 0x2c, 0xd1, 0x10, 0xea, 0x4b, 0x32, 0x5b, 0xb0, 0xc8, 0xef, 0x7a,
	0xfd, 0xe6, 0xf0, 0xa3, 0xdb, 0xab, 0xd8, 0x94, 0x15, 0x76, 0xd2, 0x2f, 0xfd, 0x87, 0x5e, 0xef,
	0x63, 0x08, 0xdc, 0x91, 0x51, 0x03, 0xea, 0xa3, 0xd3, 0xd3, 0x1f, 0x9e, 0xb5, 0x2b, 0x28, 0x84,
	0xda, 0xe3, 0x93, 0xa7, 0xcf, 0xdb, 0x5e, 0xef, 0x2f, 0x0f, 0x02, 0x67, 0x42, 0xa7, 0xd0, 0xcc,
	0x98, 0x9a, 0xf3, 0x3c, 0xe7, 0x52, 0xe4, 0x91, 0x67, 0xef, 0xa6, 0xfb, 0x8a, 0x3a, 0x1b, 0x61,
	0x12, 0x5e, 0x27, 0xf5, 0x3f, 0x3c, 0x3f, 0xf4, 0xf0, 0xb6, 0x1d, 0x4d, 0x00, 0x32, 0xc5, 0x45,
	0xca, 0x33, 0x32, 0x5b, 0xb7, 0xe6, 0xde, 0x2b, 0xc2, 0xd6, 0xba, 0xad, 0xac, 0x2d, 0x73, 0xef,
	0xcf, 0x3a, 0xc0, 0x4d, 0x41, 0x74, 0x0c, 0x0d, 0x22, 0xe8, 0x85, 0x5a, 0xcc, 0x58, 0x6e, 0x3b,
	0xd4, 0x1c, 0x7e, 0xb2, 0xeb, 0x94, 0xf1, 0x39, 0xd3, 0xe3, 0x0a, 0x0e, 0x89, 0xa0, 0xd8, 0xf8,
	0xd0, 0x08, 0x42, 0xa9, 0x8a, 0x0c, 0x7f, 0xaf, 0x8c, 0x03, 0xa9, 0x5c, 0xc4, 0x87, 0x50, 0x25,
	0x62, 0x15, 0x55, 0xbb, 0x5e, 0x3f, 0x4c, 0x0e, 0xae, 0x93, 0xda, 0xcf, 0x7e, 0xe8, 0x8d, 0x2b,
	0xd8, 0x5c, 0x45, 0x47, 0x10, 0x5c, 0x31, 0x42, 0x99, 0x8a, 0x6a, 0x36, 0x7d, 0x8d, 0x14, 0xc9,
	0xb8, 0x49, 0x75, 0x28, 0x8f, 0xad, 0xe2, 0x89, 0x63, 0x73, 0x5c, 0xc1, 0x85, 0x05, 0x9d, 0xc0,
	0x21, 0x65, 0xb9, 0xe6, 0x82, 0x98, 0x07, 0x77, 0xc1, 0xb3, 0xa8, 0x5e, 0x7a, 0xe8, 0x45, 0x88,
	0x19, 0x9b, 0xf8, 0x98, 0x53, 0x85, 0x89, 0x98, 0xb2, 0x71, 0x05, 0xb7, 0xb6, 0x5c, 0x93, 0x0c,
	0x7d, 0x01, 0xed, 0xed, 0x98, 0x4c, 0x2a, 0x1d, 0x05, 0x5d, 0xaf, 0xdf, 0x4a, 0x1a, 0xd7, 0x49,
	0xf0, 0x69, 0x2d, 0x7a, 0xf9, 0xb2, 0x3a, 0xae, 0xe0, 0x77, 0xb6, 0x44, 0x67, 0x52, 0x69, 0xd3,
	0x9b, 0xf5, 0x28, 0x45, 0x07, 0xb6, 0xf0, 0x83, 0xa2, 0xb0, 0x99, 0xa5, 0xb8, 0x98, 0xa5, 0xf8,
	0x49, 0xa1, 0xb9, 0x39, 0xff, 0xc6, 0x86, 0xbe, 0x86, 0x50, 0x48, 0x6d, 0xfb, 0x1b, 0x85, 0x36,
	0x62, 0x27, 0x48, 0xa6, 0xb5, 0x42, 0x6a, 0xd3, 0x5b, 0xf4, 0x0c, 0xde, 0x57, 0xec, 0xd7, 0x05,
	0xcb, 0x35, 0xa3, 0x17, 0x39, 0x53, 0x4b, 0xa6, 0x2e, 0x04, 0x99, 0xb3, 0xa8, 0x51, 0x6a, 0x66,
	0xe9, 0x38, 0xe7, 0x76, 0xb4, 0x6f, 0x0e, 0xf3, 0xde, 0x26, 0xe1, 0xdc, 0x06, 0x3c, 0x25, 0x73,
	0xd6, 0xf9, 0x0e, 0xaa, 0xe7, 0x4c, 0xa3, 0x6f, 0xa1, 0xbe, 0xc6, 0x67, 0x5f, 0xc8, 0x9d, 0x31,
	0x69, 0x42, 0xcd, 0x2c, 0x50, 0xf5, 0xbf, 0xc4, 0xeb, 0xfd, 0x5d, 0x87, 0xc6, 0x06, 0x62, 0xf4,
	0x0d, 0x1c, 0x18, 0x3e, 0x39, 0x5d, 0xd3, 0xf9, 0x60, 0x07, 0xf6, 0x05, 0x58, 0x01, 0x11, 0x74,
	0x42, 0x73, 0xf4, 0x15, 0x04, 0x52, 0x59, 0xbb, 0xbf, 0x8f, 0xbd, 0x2e, 0x95, 0x71, 0xbf, 0x96,
	0xca, 0x1f, 0xa1, 0x45, 0x16, 0xfa, 0x8a, 0x09, 0xcd, 0x53, 0xa2, 0x19, 0x2d, 0xe0, 0xfc, 0x6c,
	0x57, 0x85, 0xd1, 0xb6, 0xc9, 0x80, 0x56, 0x4a, 0x41, 0x47, 0xd0, 0xc8, 0xe5, 0x42, 0xa5, 0xec,
	0xcd, 0x51, 0x0d, 0x9d, 0x61, 0x92, 0x6d, 0x4d, 0x4a, 0xb0, 0xff, 0xa4, 0xbc, 0x05, 0x54, 0x1f,
	0x42, 0x60, 0x50, 0xe5, 0xb4, 0x00, 0x75, 0xd7, 0x4b, 0xca, 0xb4, 0x5a, 0x48, 0x3d, 0xa1, 0x9d,
	0xc4, 0xc1, 0x74, 0x04, 0x55, 0xf7, 0xac, 0xf7, 0x7c, 0xc5, 0x19, 0x57, 0x87, 0x41, 0xab, 0xd4,
	0x5c, 0x34, 0x86, 0xc3, 0xcd, 0xab, 0xcf, 0x31, 0xef, 0xbf, 0x21, 0xf3, 0xb8, 0xb5, 0x31, 0x1a,
	0xd6, 0xbf, 0xaf, 0x85, 0x5e, 0xdb, 0xc7, 0x35, 0x93, 0x91, 0xbc, 0x0b, 0xc0, 0xa9, 0x29, 0xf2,
	0x82, 0x33, 0x65, 0xa1, 0x4d, 0x1e, 0xfd, 0xf3, 0xef, 0x5d, 0x0f, 0x7a, 0x5c, 0xba, 0xe0, 0x4c,
	0xc9, 0xdf, 0x56, 0xb7, 0x1e, 0x3e, 0x69, 0xe0, 0x4b, 0x92, 0x9e, 0x99, 0x4f, 0xe6, 0x99, 0xf7,
	0x93, 0xbf, 0x1c, 0x5e, 0x06, 0xf6, 0xfb, 0xf9, 0xf9, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e,
	0xad, 0xcd, 0x8d, 0x2c, 0x08, 0x00, 0x00,
}