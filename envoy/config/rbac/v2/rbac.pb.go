// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/rbac/v2/rbac.proto

package envoy_config_rbac_v2

import (
	fmt "fmt"
	core "github.com/altipla-consulting/envoy-api/envoy/api/v2/core"
	route "github.com/altipla-consulting/envoy-api/envoy/api/v2/route"
	matcher "github.com/altipla-consulting/envoy-api/envoy/type/matcher"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
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
// principals match the downstream AND the condition is true if specified.
type Policy struct {
	// Required. The set of permissions that define a role. Each permission is matched with OR
	// semantics. To match all actions for this policy, a single Permission with the `any` field set
	// to true should be used.
	Permissions []*Permission `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// Required. The set of principals that are assigned/denied the role based on “action”. Each
	// principal is matched with OR semantics. To match all downstreams for this policy, a single
	// Principal with the `any` field set to true should be used.
	Principals []*Principal `protobuf:"bytes,2,rep,name=principals,proto3" json:"principals,omitempty"`
	// An optional symbolic expression specifying an access control
	// :ref:`condition <arch_overview_condition>`. The condition is combined
	// with the permissions and the principals as a clause with AND semantics.
	Condition            *v1alpha1.Expr `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
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

func (m *Policy) GetCondition() *v1alpha1.Expr {
	if m != nil {
		return m.Condition
	}
	return nil
}

// Permission defines an action (or actions) that a principal can take.
// [#next-free-field: 10]
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
// [#next-free-field: 9]
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
	// The name of the principal. If set, The URI SAN or DNS SAN in that order is used from the
	// certificate, otherwise the subject field is used. If unset, it applies to any user that is
	// authenticated.
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
	// 885 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xd1, 0x6e, 0x23, 0x35,
	0x14, 0xcd, 0x4c, 0x92, 0xe9, 0xcc, 0x8d, 0x52, 0x82, 0x01, 0x31, 0x0a, 0xb0, 0xcd, 0x66, 0x41,
	0x8a, 0x90, 0x98, 0x51, 0x83, 0x84, 0x16, 0x0a, 0x88, 0x4c, 0xb7, 0x22, 0x41, 0xdd, 0xa5, 0x9a,
	0x0a, 0xad, 0xf6, 0x29, 0x72, 0xc7, 0xde, 0xd4, 0x90, 0xd8, 0x83, 0xc7, 0x89, 0x9a, 0x47, 0x7e,
	0x81, 0x0f, 0xe0, 0x0b, 0xf8, 0x27, 0x7e, 0x03, 0xf5, 0x65, 0x91, 0xed, 0x49, 0x9a, 0x48, 0xdd,
	0xcd, 0x46, 0xe2, 0xa5, 0xb5, 0xe2, 0x73, 0x8e, 0xef, 0xbd, 0x3e, 0xc7, 0x03, 0x47, 0x94, 0x2f,
	0xc4, 0x32, 0xce, 0x04, 0x7f, 0xc9, 0x26, 0xb1, 0xbc, 0xc2, 0x59, 0xbc, 0xe8, 0x9b, 0xff, 0x51,
	0x2e, 0x85, 0x12, 0xe8, 0x7d, 0x03, 0x88, 0x2c, 0x20, 0x32, 0x1b, 0x8b, 0x7e, 0xbb, 0xa4, 0xe1,
	0x9c, 0x69, 0x78, 0x26, 0x24, 0x8d, 0x31, 0x21, 0x92, 0x16, 0x85, 0xa5, 0xb5, 0x1f, 0x6c, 0x01,
	0xa4, 0x98, 0x2b, 0x6a, 0xff, 0x96, 0xfb, 0x0f, 0xed, 0xbe, 0x5a, 0xe6, 0x34, 0x9e, 0x61, 0x95,
	0x5d, 0x53, 0x19, 0xcf, 0xa8, 0xc2, 0x04, 0x2b, 0x5c, 0x42, 0x8e, 0xee, 0x81, 0x14, 0x4a, 0x32,
	0x3e, 0x29, 0x01, 0x9f, 0x4d, 0x84, 0x98, 0x4c, 0xa9, 0x39, 0x84, 0xde, 0xe4, 0x32, 0x5e, 0x1c,
	0xe3, 0x69, 0x7e, 0x8d, 0x8f, 0xe3, 0x62, 0xc9, 0x15, 0xbe, 0x29, 0x61, 0x1f, 0x2e, 0xf0, 0x94,
	0x11, 0xac, 0x68, 0xbc, 0x5a, 0xd8, 0x8d, 0xee, 0x1f, 0x2e, 0xd4, 0xd2, 0x64, 0x70, 0x8a, 0xbe,
	0x06, 0x0f, 0x67, 0x8a, 0x09, 0x1e, 0x3a, 0x1d, 0xa7, 0x77, 0xd8, 0x7f, 0x18, 0xdd, 0xd7, 0x74,
	0xa4, 0xb1, 0xd1, 0xc0, 0x00, 0xd3, 0x92, 0x80, 0x9e, 0x80, 0x9f, 0x8b, 0x29, 0xcb, 0x18, 0x2d,
	0x42, 0xb7, 0x53, 0xed, 0x35, 0xfa, 0xbd, 0x37, 0x90, 0x2f, 0x4a, 0xe8, 0x19, 0x57, 0x72, 0x99,
	0xae, 0x99, 0xed, 0x17, 0xd0, 0xdc, 0xda, 0x42, 0x2d, 0xa8, 0xfe, 0x46, 0x97, 0xa6, 0x9c, 0x20,
	0xd5, 0x4b, 0xd4, 0x87, 0xfa, 0x02, 0x4f, 0xe7, 0x34, 0x74, 0x3b, 0x4e, 0xaf, 0xd1, 0xff, 0xf8,
	0xfe, 0x53, 0x8c, 0xca, 0x32, 0xb5, 0xd0, 0x6f, 0xdc, 0xc7, 0x4e, 0xf7, 0x13, 0xf0, 0x6c, 0xc9,
	0x28, 0x80, 0xfa, 0xe0, 0xfc, 0xfc, 0xe7, 0xe7, 0xad, 0x0a, 0xf2, 0xa1, 0xf6, 0xe4, 0xec, 0xd9,
	0x8b, 0x96, 0xd3, 0xfd, 0xc7, 0x01, 0xcf, 0x92, 0xd0, 0x39, 0x34, 0x72, 0x2a, 0x67, 0xac, 0x28,
	0x98, 0xe0, 0x45, 0xe8, 0x98, 0x6e, 0x3a, 0xaf, 0x39, 0x67, 0x0d, 0x4c, 0xfc, 0xdb, 0xa4, 0xfe,
	0xa7, 0xe3, 0xfa, 0x4e, 0xba, 0x49, 0x47, 0x23, 0x80, 0x5c, 0x32, 0x9e, 0xb1, 0x1c, 0x4f, 0x57,
	0xa3, 0x39, 0x7a, 0x8d, 0xd8, 0x0a, 0xb7, 0xa1, 0xb5, 0x41, 0x46, 0xdf, 0x42, 0x90, 0x09, 0x4e,
	0x98, 0xb9, 0xa1, 0xaa, 0x69, 0xff, 0x41, 0x64, 0xef, 0x3e, 0xc2, 0x39, 0x8b, 0xf4, 0xdd, 0x47,
	0xab, 0xbb, 0x8f, 0xce, 0x6e, 0x72, 0x99, 0xde, 0x11, 0xba, 0x7f, 0xd5, 0x01, 0xee, 0xca, 0x45,
	0xa7, 0x10, 0x60, 0x4e, 0xc6, 0x72, 0x3e, 0xa5, 0x85, 0x99, 0x6f, 0xa3, 0xff, 0xe9, 0xae, 0x1e,
	0xa3, 0x4b, 0xaa, 0x86, 0x95, 0xd4, 0xc7, 0x9c, 0xa4, 0x9a, 0x87, 0x06, 0xe0, 0x0b, 0x59, 0x6a,
	0xb8, 0x7b, 0x69, 0x1c, 0x08, 0x69, 0x25, 0x3e, 0x82, 0x2a, 0xe6, 0x4b, 0xd3, 0x8e, 0x9f, 0x1c,
	0xdc, 0x26, 0xb5, 0x5f, 0x5d, 0xdf, 0x19, 0x56, 0x52, 0xfd, 0x2b, 0x3a, 0x01, 0xef, 0x9a, 0x62,
	0x42, 0x65, 0x58, 0x33, 0xea, 0x2b, 0x43, 0xea, 0x6e, 0x17, 0xfd, 0xc8, 0x06, 0x69, 0x68, 0x10,
	0x4f, 0x6d, 0x32, 0x86, 0x95, 0xb4, 0xa4, 0xa0, 0x33, 0x38, 0x24, 0xb4, 0x50, 0x8c, 0x63, 0xdd,
	0xff, 0x98, 0xe5, 0x61, 0x7d, 0xcb, 0x32, 0xa5, 0x88, 0x0e, 0x6d, 0x74, 0xca, 0x88, 0x4c, 0x31,
	0x9f, 0xd0, 0x61, 0x25, 0x6d, 0x6e, 0xb0, 0x46, 0x39, 0xfa, 0x0a, 0x5a, 0x9b, 0x32, 0xb9, 0x90,
	0x2a, 0xf4, 0x3a, 0x4e, 0xaf, 0x99, 0x04, 0xb7, 0x89, 0xf7, 0x79, 0x2d, 0x7c, 0xf5, 0xaa, 0x3a,
	0xac, 0xa4, 0xef, 0x6c, 0x80, 0x2e, 0x84, 0x54, 0x7a, 0x36, 0xab, 0x20, 0x87, 0x07, 0xe6, 0xe0,
	0x47, 0xe5, 0xc1, 0x3a, 0xc9, 0x51, 0x99, 0xe4, 0xe8, 0x69, 0x89, 0xb9, 0xab, 0x7f, 0x4d, 0x43,
	0xdf, 0x81, 0xcf, 0x85, 0x32, 0xf3, 0x0d, 0x7d, 0x23, 0xb1, 0xd3, 0x86, 0x7a, 0xb4, 0x5c, 0x28,
	0x3d, 0x5b, 0xf4, 0x1c, 0x3e, 0x90, 0xf4, 0xf7, 0x39, 0x2d, 0x14, 0x25, 0xe3, 0x82, 0xca, 0x05,
	0x95, 0x63, 0x8e, 0x67, 0x34, 0x0c, 0xb6, 0x86, 0xb9, 0x55, 0xce, 0xa5, 0x79, 0x58, 0xee, 0x8a,
	0x79, 0x6f, 0xad, 0x70, 0x69, 0x04, 0x9e, 0xe1, 0x19, 0x6d, 0xff, 0x08, 0xd5, 0x4b, 0xaa, 0xd0,
	0x0f, 0x50, 0x5f, 0xd9, 0x67, 0xdf, 0x88, 0x58, 0x62, 0xd2, 0x80, 0x9a, 0x5e, 0xa0, 0xea, 0xbf,
	0x89, 0xd3, 0xfd, 0xbb, 0x0e, 0xc1, 0x3a, 0x02, 0xe8, 0x7b, 0x38, 0xd0, 0xfe, 0x64, 0x64, 0xe5,
	0xce, 0x47, 0x3b, 0x42, 0x53, 0x1a, 0xcb, 0xc3, 0x9c, 0x8c, 0x88, 0x0e, 0x8b, 0x27, 0xa4, 0xa1,
	0xbb, 0xfb, 0xd0, 0xeb, 0x42, 0x6a, 0xf6, 0x1b, 0x5d, 0xf9, 0x0b, 0x34, 0xf1, 0x5c, 0x5d, 0x53,
	0xae, 0x58, 0x86, 0x15, 0x25, 0xa5, 0x39, 0xbf, 0xd8, 0x75, 0xc2, 0x60, 0x93, 0xa4, 0x8d, 0xb6,
	0xa5, 0x82, 0x4e, 0x20, 0x28, 0xc4, 0x5c, 0x66, 0xf4, 0xed, 0xad, 0xea, 0x5b, 0xc2, 0x28, 0xdf,
	0x48, 0x8a, 0xb7, 0x7f, 0x52, 0xfe, 0x07, 0xab, 0x3e, 0x06, 0x4f, 0x5b, 0x95, 0x91, 0xd2, 0xa8,
	0xbb, 0x9e, 0x38, 0x3d, 0x6a, 0x2e, 0xd4, 0x88, 0xb4, 0x13, 0x6b, 0xa6, 0x13, 0xa8, 0xda, 0xbb,
	0xde, 0xf3, 0x81, 0xd4, 0xac, 0xf6, 0x18, 0x9a, 0x5b, 0xc3, 0x45, 0x43, 0x38, 0x5c, 0x3f, 0x9c,
	0xd6, 0xf3, 0xee, 0x5b, 0x7a, 0x3e, 0x6d, 0xae, 0x89, 0xda, 0xeb, 0x3f, 0xd5, 0x7c, 0xa7, 0xe5,
	0x26, 0xef, 0x02, 0x30, 0xa2, 0xe5, 0x5f, 0x32, 0x2a, 0x8d, 0x5d, 0x93, 0x63, 0xe8, 0x32, 0x61,
	0xe5, 0x72, 0x29, 0x6e, 0x96, 0xf7, 0x96, 0x9c, 0x04, 0xe9, 0x15, 0xce, 0x2e, 0xf4, 0x67, 0xf6,
	0xc2, 0xb9, 0xf2, 0xcc, 0xf7, 0xf6, 0xcb, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xce, 0xc7,
	0x9a, 0x6d, 0x08, 0x00, 0x00,
}
