// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/thrift/router/v2alpha1/router.proto

package v2alpha1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Router struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Router) Reset()         { *m = Router{} }
func (m *Router) String() string { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()    {}
func (*Router) Descriptor() ([]byte, []int) {
	return fileDescriptor_router_027bcbf3aaf7672c, []int{0}
}
func (m *Router) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Router.Unmarshal(m, b)
}
func (m *Router) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Router.Marshal(b, m, deterministic)
}
func (dst *Router) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Router.Merge(dst, src)
}
func (m *Router) XXX_Size() int {
	return xxx_messageInfo_Router.Size(m)
}
func (m *Router) XXX_DiscardUnknown() {
	xxx_messageInfo_Router.DiscardUnknown(m)
}

var xxx_messageInfo_Router proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Router)(nil), "envoy.config.filter.thrift.router.v2alpha1.Router")
}

func init() {
	proto.RegisterFile("envoy/config/filter/thrift/router/v2alpha1/router.proto", fileDescriptor_router_027bcbf3aaf7672c)
}

var fileDescriptor_router_027bcbf3aaf7672c = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0x2f, 0xc9, 0x28, 0xca, 0x4c, 0x2b, 0xd1, 0x2f, 0xca, 0x2f, 0x05, 0xf1, 0xca, 0x8c, 0x12, 0x73,
	0x0a, 0x32, 0x12, 0x0d, 0xa1, 0x7c, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x2d, 0xb0, 0x46,
	0x3d, 0x88, 0x46, 0x3d, 0x88, 0x46, 0x3d, 0x88, 0x46, 0x3d, 0xa8, 0x42, 0x98, 0x46, 0x25, 0x0e,
	0x2e, 0xb6, 0x20, 0xb0, 0x90, 0x13, 0x57, 0x14, 0x07, 0x4c, 0x34, 0x89, 0x0d, 0x6c, 0x90, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xfc, 0x97, 0x92, 0x83, 0x00, 0x00, 0x00,
}
