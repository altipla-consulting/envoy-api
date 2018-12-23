// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/accesslog/v2/file.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Custom configuration for an :ref:`AccessLog <envoy_api_msg_config.filter.accesslog.v2.AccessLog>`
// that writes log entries directly to a file. Configures the built-in *envoy.file_access_log*
// AccessLog.
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
	return fileDescriptor_file_c2c17a447909f09b, []int{0}
}
func (m *FileAccessLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileAccessLog.Unmarshal(m, b)
}
func (m *FileAccessLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileAccessLog.Marshal(b, m, deterministic)
}
func (dst *FileAccessLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileAccessLog.Merge(dst, src)
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

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FileAccessLog) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FileAccessLog_OneofMarshaler, _FileAccessLog_OneofUnmarshaler, _FileAccessLog_OneofSizer, []interface{}{
		(*FileAccessLog_Format)(nil),
		(*FileAccessLog_JsonFormat)(nil),
	}
}

func _FileAccessLog_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FileAccessLog)
	// access_log_format
	switch x := m.AccessLogFormat.(type) {
	case *FileAccessLog_Format:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Format)
	case *FileAccessLog_JsonFormat:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.JsonFormat); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("FileAccessLog.AccessLogFormat has unexpected type %T", x)
	}
	return nil
}

func _FileAccessLog_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FileAccessLog)
	switch tag {
	case 2: // access_log_format.format
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.AccessLogFormat = &FileAccessLog_Format{x}
		return true, err
	case 3: // access_log_format.json_format
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types.Struct)
		err := b.DecodeMessage(msg)
		m.AccessLogFormat = &FileAccessLog_JsonFormat{msg}
		return true, err
	default:
		return false, nil
	}
}

func _FileAccessLog_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FileAccessLog)
	// access_log_format
	switch x := m.AccessLogFormat.(type) {
	case *FileAccessLog_Format:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Format)))
		n += len(x.Format)
	case *FileAccessLog_JsonFormat:
		s := proto.Size(x.JsonFormat)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*FileAccessLog)(nil), "envoy.config.accesslog.v2.FileAccessLog")
}

func init() {
	proto.RegisterFile("envoy/config/accesslog/v2/file.proto", fileDescriptor_file_c2c17a447909f09b)
}

var fileDescriptor_file_c2c17a447909f09b = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0x4c, 0x4e, 0x4e, 0x2d, 0x2e, 0xce,
	0xc9, 0x4f, 0xd7, 0x2f, 0x33, 0xd2, 0x4f, 0xcb, 0xcc, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x04, 0xab, 0xd2, 0x83, 0xa8, 0xd2, 0x83, 0xab, 0xd2, 0x2b, 0x33, 0x92, 0x12, 0x2f,
	0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x7a, 0xa4, 0x64, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xbc, 0xa4, 0xd2, 0x34, 0xfd, 0xe2, 0x92, 0xa2, 0xd2, 0xe4,
	0x12, 0x88, 0xac, 0xd2, 0x4c, 0x46, 0x2e, 0x5e, 0xb7, 0xcc, 0x9c, 0x54, 0x47, 0xb0, 0x59, 0x3e,
	0xf9, 0xe9, 0x42, 0xb2, 0x5c, 0x2c, 0x05, 0x89, 0x25, 0x19, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x4e, 0x9c, 0xbb, 0x5e, 0x1e, 0x60, 0x66, 0x29, 0x62, 0x52, 0x60, 0x0c, 0x02, 0x0b, 0x0b, 0x49,
	0x70, 0xb1, 0xa5, 0xe5, 0x17, 0xe5, 0x26, 0x96, 0x48, 0x30, 0x81, 0x14, 0x78, 0x30, 0x04, 0x41,
	0xf9, 0x42, 0x56, 0x5c, 0xdc, 0x59, 0xc5, 0xf9, 0x79, 0xf1, 0x50, 0x69, 0x66, 0x05, 0x46, 0x0d,
	0x6e, 0x23, 0x71, 0x3d, 0x88, 0xf5, 0x7a, 0x30, 0xeb, 0xf5, 0x82, 0xc1, 0xd6, 0x7b, 0x30, 0x04,
	0x71, 0x81, 0x54, 0xbb, 0x81, 0x15, 0x3b, 0x09, 0x73, 0x09, 0x42, 0x7c, 0x13, 0x9f, 0x93, 0x9f,
	0x0e, 0x35, 0xc1, 0x89, 0x25, 0x8a, 0xa9, 0xcc, 0x28, 0x89, 0x0d, 0xac, 0xd3, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xb2, 0x4c, 0xa9, 0xda, 0x22, 0x01, 0x00, 0x00,
}
