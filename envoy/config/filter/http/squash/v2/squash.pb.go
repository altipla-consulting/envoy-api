// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/squash/v2/squash.proto

package v2

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
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

// [#proto-status: experimental]
type Squash struct {
	// The name of the cluster that hosts the Squash server.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// When the filter requests the Squash server to create a DebugAttachment, it will use this
	// structure as template for the body of the request. It can contain reference to environment
	// variables in the form of '{{ ENV_VAR_NAME }}'. These can be used to provide the Squash server
	// with more information to find the process to attach the debugger to. For example, in a
	// Istio/k8s environment, this will contain information on the pod:
	//
	// .. code-block:: json
	//
	//  {
	//    "spec": {
	//      "attachment": {
	//        "pod": "{{ POD_NAME }}",
	//        "namespace": "{{ POD_NAMESPACE }}"
	//      },
	//      "match_request": true
	//    }
	//  }
	//
	// (where POD_NAME, POD_NAMESPACE are configured in the pod via the Downward API)
	AttachmentTemplate *types.Struct `protobuf:"bytes,2,opt,name=attachment_template,json=attachmentTemplate,proto3" json:"attachment_template,omitempty"`
	// The timeout for individual requests sent to the Squash cluster. Defaults to 1 second.
	RequestTimeout *types.Duration `protobuf:"bytes,3,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	// The total timeout Squash will delay a request and wait for it to be attached. Defaults to 60
	// seconds.
	AttachmentTimeout *types.Duration `protobuf:"bytes,4,opt,name=attachment_timeout,json=attachmentTimeout,proto3" json:"attachment_timeout,omitempty"`
	// Amount of time to poll for the status of the attachment object in the Squash server
	// (to check if has been attached). Defaults to 1 second.
	AttachmentPollPeriod *types.Duration `protobuf:"bytes,5,opt,name=attachment_poll_period,json=attachmentPollPeriod,proto3" json:"attachment_poll_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Squash) Reset()         { *m = Squash{} }
func (m *Squash) String() string { return proto.CompactTextString(m) }
func (*Squash) ProtoMessage()    {}
func (*Squash) Descriptor() ([]byte, []int) {
	return fileDescriptor_63fc8434388b1e13, []int{0}
}

func (m *Squash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Squash.Unmarshal(m, b)
}
func (m *Squash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Squash.Marshal(b, m, deterministic)
}
func (m *Squash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Squash.Merge(m, src)
}
func (m *Squash) XXX_Size() int {
	return xxx_messageInfo_Squash.Size(m)
}
func (m *Squash) XXX_DiscardUnknown() {
	xxx_messageInfo_Squash.DiscardUnknown(m)
}

var xxx_messageInfo_Squash proto.InternalMessageInfo

func (m *Squash) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Squash) GetAttachmentTemplate() *types.Struct {
	if m != nil {
		return m.AttachmentTemplate
	}
	return nil
}

func (m *Squash) GetRequestTimeout() *types.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentTimeout() *types.Duration {
	if m != nil {
		return m.AttachmentTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentPollPeriod() *types.Duration {
	if m != nil {
		return m.AttachmentPollPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*Squash)(nil), "envoy.config.filter.http.squash.v2.Squash")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/squash/v2/squash.proto", fileDescriptor_63fc8434388b1e13)
}

var fileDescriptor_63fc8434388b1e13 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6e, 0xe2, 0x30,
	0x14, 0x86, 0x95, 0x0c, 0x03, 0x1a, 0x23, 0xcd, 0x68, 0x52, 0x54, 0x52, 0x54, 0xb5, 0x94, 0x15,
	0x2b, 0xbb, 0x4a, 0x6f, 0x10, 0xb5, 0x12, 0xab, 0x2a, 0x02, 0xba, 0xe9, 0x06, 0x85, 0x60, 0x82,
	0x25, 0x93, 0x17, 0x9c, 0xe7, 0xa8, 0xdc, 0xa4, 0xb7, 0xeb, 0xae, 0x87, 0xe8, 0xaa, 0xc2, 0x36,
	0x02, 0xb5, 0x8b, 0xb2, 0xfb, 0xe5, 0x97, 0xef, 0x7b, 0x7f, 0x64, 0x13, 0xc6, 0x8b, 0x1a, 0xb6,
	0x2c, 0x83, 0x62, 0x29, 0x72, 0xb6, 0x14, 0x12, 0xb9, 0x62, 0x2b, 0xc4, 0x92, 0x55, 0x1b, 0x9d,
	0x56, 0x2b, 0x56, 0x47, 0x2e, 0xd1, 0x52, 0x01, 0x42, 0x30, 0x30, 0x00, 0xb5, 0x00, 0xb5, 0x00,
	0xdd, 0x01, 0xd4, 0x7d, 0x56, 0x47, 0xbd, 0xab, 0x1c, 0x20, 0x97, 0x9c, 0x19, 0x62, 0xae, 0x97,
	0x6c, 0xa1, 0x55, 0x8a, 0x02, 0x0a, 0xeb, 0xe8, 0x5d, 0x7e, 0x9d, 0x57, 0xa8, 0x74, 0x86, 0x6e,
	0xda, 0xad, 0x53, 0x29, 0x16, 0x29, 0x72, 0xb6, 0x0f, 0x6e, 0xd0, 0xc9, 0x21, 0x07, 0x13, 0xd9,
	0x2e, 0xd9, 0xd3, 0xc1, 0xbb, 0x4f, 0x9a, 0x13, 0xb3, 0x3a, 0xb8, 0x21, 0xad, 0x4c, 0xea, 0x0a,
	0xb9, 0x0a, 0xbd, 0xbe, 0x37, 0xfc, 0x13, 0xb7, 0x3e, 0xe2, 0x86, 0xf2, 0xfb, 0xde, 0x78, 0x7f,
	0x1e, 0x8c, 0xc8, 0x59, 0x8a, 0x98, 0x66, 0xab, 0x35, 0x2f, 0x70, 0x86, 0x7c, 0x5d, 0xca, 0x14,
	0x79, 0xe8, 0xf7, 0xbd, 0x61, 0x3b, 0xea, 0x52, 0x5b, 0x8c, 0xee, 0x8b, 0xd1, 0x89, 0x29, 0x36,
	0x0e, 0x0e, 0xcc, 0xd4, 0x21, 0xc1, 0x88, 0xfc, 0x53, 0x7c, 0xa3, 0x79, 0x85, 0x33, 0x14, 0x6b,
	0x0e, 0x1a, 0xc3, 0x5f, 0xc6, 0x72, 0xf1, 0xcd, 0x72, 0xef, 0x7e, 0x3f, 0x6e, 0xbc, 0xbe, 0x5d,
	0x7b, 0xe3, 0xbf, 0x8e, 0x9b, 0x5a, 0x2c, 0x78, 0x24, 0xc1, 0x71, 0x27, 0x27, 0x6b, 0x9c, 0x26,
	0xfb, 0x7f, 0x54, 0xcd, 0xf9, 0x9e, 0xc8, 0xf9, 0x91, 0xaf, 0x04, 0x29, 0x67, 0x25, 0x57, 0x02,
	0x16, 0xe1, 0xef, 0xd3, 0x9c, 0x9d, 0x03, 0x9e, 0x80, 0x94, 0x89, 0x81, 0xe3, 0x07, 0x72, 0x2b,
	0x80, 0x9a, 0xeb, 0x2f, 0x15, 0xbc, 0x6c, 0xe9, 0xcf, 0x2f, 0x21, 0x6e, 0xdb, 0x9b, 0x49, 0x76,
	0x8b, 0x12, 0xef, 0xd9, 0xaf, 0xa3, 0x79, 0xd3, 0x6c, 0xbd, 0xfb, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0xf7, 0x44, 0x20, 0xbb, 0x7a, 0x02, 0x00, 0x00,
}
