// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/accesslog/v3alpha/file.proto

package envoy_config_accesslog_v3alpha

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

// Custom configuration for an :ref:`AccessLog
// <envoy_api_msg_config.filter.accesslog.v3alpha.AccessLog>` that writes log entries directly to a
// file. Configures the built-in *envoy.file_access_log* AccessLog.
type FileAccessLog struct {
	// A path to a local file to which to write the access log entries.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Access log format. Envoy supports :ref:`custom access log formats
	// <config_access_log_format>` as well as a :ref:`default format
	// <config_access_log_default_format>`.
	//
	// Types that are valid to be assigned to AccessLogFormat:
	//	*FileAccessLog_Format
	//	*FileAccessLog_JsonFormat
	AccessLogFormat      isFileAccessLog_AccessLogFormat `protobuf_oneof:"access_log_format"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FileAccessLog) Reset()         { *m = FileAccessLog{} }
func (m *FileAccessLog) String() string { return proto.CompactTextString(m) }
func (*FileAccessLog) ProtoMessage()    {}
func (*FileAccessLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d015b66e55289b5, []int{0}
}

func (m *FileAccessLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileAccessLog.Unmarshal(m, b)
}
func (m *FileAccessLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileAccessLog.Marshal(b, m, deterministic)
}
func (m *FileAccessLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileAccessLog.Merge(m, src)
}
func (m *FileAccessLog) XXX_Size() int {
	return xxx_messageInfo_FileAccessLog.Size(m)
}
func (m *FileAccessLog) XXX_DiscardUnknown() {
	xxx_messageInfo_FileAccessLog.DiscardUnknown(m)
}

var xxx_messageInfo_FileAccessLog proto.InternalMessageInfo

func (m *FileAccessLog) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type isFileAccessLog_AccessLogFormat interface {
	isFileAccessLog_AccessLogFormat()
}

type FileAccessLog_Format struct {
	Format string `protobuf:"bytes,2,opt,name=format,proto3,oneof"`
}

type FileAccessLog_JsonFormat struct {
	JsonFormat *types.Struct `protobuf:"bytes,3,opt,name=json_format,json=jsonFormat,proto3,oneof"`
}

func (*FileAccessLog_Format) isFileAccessLog_AccessLogFormat() {}

func (*FileAccessLog_JsonFormat) isFileAccessLog_AccessLogFormat() {}

func (m *FileAccessLog) GetAccessLogFormat() isFileAccessLog_AccessLogFormat {
	if m != nil {
		return m.AccessLogFormat
	}
	return nil
}

func (m *FileAccessLog) GetFormat() string {
	if x, ok := m.GetAccessLogFormat().(*FileAccessLog_Format); ok {
		return x.Format
	}
	return ""
}

func (m *FileAccessLog) GetJsonFormat() *types.Struct {
	if x, ok := m.GetAccessLogFormat().(*FileAccessLog_JsonFormat); ok {
		return x.JsonFormat
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FileAccessLog) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FileAccessLog_Format)(nil),
		(*FileAccessLog_JsonFormat)(nil),
	}
}

func init() {
	proto.RegisterType((*FileAccessLog)(nil), "envoy.config.accesslog.v3alpha.FileAccessLog")
}

func init() {
	proto.RegisterFile("envoy/config/accesslog/v3alpha/file.proto", fileDescriptor_4d015b66e55289b5)
}

var fileDescriptor_4d015b66e55289b5 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xeb, 0x52, 0x15, 0xd5, 0x15, 0x03, 0x61, 0x68, 0x54, 0x10, 0x8a, 0x98, 0x8a, 0x84,
	0xce, 0x12, 0xdd, 0x60, 0xc2, 0x43, 0xd5, 0x81, 0xa1, 0x0a, 0x3f, 0xa0, 0x72, 0x83, 0xe3, 0x1a,
	0x99, 0x5c, 0xe4, 0xb8, 0x11, 0xfd, 0x25, 0xfc, 0x57, 0x26, 0x94, 0x73, 0xca, 0xc8, 0x66, 0xfb,
	0x7d, 0xbe, 0xf7, 0xde, 0xf1, 0x7b, 0x5d, 0xb5, 0x78, 0x14, 0x05, 0x56, 0xa5, 0x35, 0x42, 0x15,
	0x85, 0x6e, 0x1a, 0x87, 0x46, 0xb4, 0x4b, 0xe5, 0xea, 0xbd, 0x12, 0xa5, 0x75, 0x1a, 0x6a, 0x8f,
	0x01, 0x93, 0x5b, 0x42, 0x21, 0xa2, 0xf0, 0x87, 0x42, 0x8f, 0xce, 0x6f, 0x0c, 0xa2, 0x71, 0x5a,
	0x10, 0xbd, 0x3b, 0x94, 0xa2, 0x09, 0xfe, 0x50, 0x84, 0xf8, 0x7b, 0x3e, 0x6b, 0x95, 0xb3, 0xef,
	0x2a, 0x68, 0x71, 0x3a, 0x44, 0xe1, 0xee, 0x9b, 0xf1, 0x8b, 0x95, 0x75, 0xfa, 0x85, 0x06, 0xbe,
	0xa2, 0x49, 0xae, 0xf9, 0xa8, 0x56, 0x61, 0x9f, 0xb2, 0x8c, 0x2d, 0x26, 0xf2, 0xfc, 0x47, 0x8e,
	0xfc, 0x30, 0x63, 0x39, 0x3d, 0x26, 0x29, 0x1f, 0x97, 0xe8, 0x3f, 0x55, 0x48, 0x87, 0x9d, 0xbc,
	0x1e, 0xe4, 0xfd, 0x3d, 0x79, 0xe2, 0xd3, 0x8f, 0x06, 0xab, 0x6d, 0x2f, 0x9f, 0x65, 0x6c, 0x31,
	0x7d, 0x9c, 0x41, 0x4c, 0x05, 0xa7, 0x54, 0xf0, 0x46, 0xa9, 0xd6, 0x83, 0x9c, 0x77, 0xf4, 0x8a,
	0x60, 0x79, 0xc5, 0x2f, 0x63, 0xa1, 0xad, 0x43, 0xd3, 0x4f, 0x90, 0xcf, 0xfc, 0xc1, 0x22, 0x50,
	0xeb, 0xda, 0xe3, 0xd7, 0x11, 0xfe, 0x5f, 0x80, 0x9c, 0x74, 0x35, 0x36, 0x9d, 0xcf, 0x86, 0xed,
	0xc6, 0x64, 0xb8, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x82, 0x8e, 0x1f, 0x7c, 0x61, 0x01, 0x00,
	0x00,
}
