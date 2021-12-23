// Code generated by protoc-gen-go. DO NOT EDIT.
// source: db.proto

package protos

import (
	fmt "fmt"
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

// Represents value stored in key/value database
// User provided value is automatically (de)serialized
type DbValue struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Version              int32    `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	IsNative             bool     `protobuf:"varint,4,opt,name=is_native,json=isNative,proto3" json:"is_native,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DbValue) Reset()         { *m = DbValue{} }
func (m *DbValue) String() string { return proto.CompactTextString(m) }
func (*DbValue) ProtoMessage()    {}
func (*DbValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8817812184a13374, []int{0}
}

func (m *DbValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DbValue.Unmarshal(m, b)
}
func (m *DbValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DbValue.Marshal(b, m, deterministic)
}
func (m *DbValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DbValue.Merge(m, src)
}
func (m *DbValue) XXX_Size() int {
	return xxx_messageInfo_DbValue.Size(m)
}
func (m *DbValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DbValue.DiscardUnknown(m)
}

var xxx_messageInfo_DbValue proto.InternalMessageInfo

func (m *DbValue) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DbValue) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *DbValue) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *DbValue) GetIsNative() bool {
	if m != nil {
		return m.IsNative
	}
	return false
}

func init() {
	proto.RegisterType((*DbValue)(nil), "DbValue")
}

func init() { proto.RegisterFile("db.proto", fileDescriptor_8817812184a13374) }

var fileDescriptor_8817812184a13374 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xb1, 0xaa, 0xc2, 0x30,
	0x18, 0x46, 0xc9, 0xbd, 0xd5, 0xb6, 0xc1, 0x29, 0x53, 0xc0, 0x25, 0x38, 0x05, 0xd1, 0x66, 0xf0,
	0x0d, 0xc4, 0xd9, 0x21, 0x83, 0x83, 0x8b, 0x24, 0x6d, 0x68, 0x7f, 0x68, 0xf3, 0x97, 0x24, 0x2d,
	0xf8, 0xf6, 0x62, 0xc0, 0xe9, 0x3b, 0xe7, 0x9b, 0x0e, 0xad, 0x3a, 0xdb, 0xcc, 0x01, 0x13, 0x1e,
	0x06, 0x5a, 0xde, 0xec, 0xc3, 0x8c, 0x8b, 0x63, 0x8c, 0x16, 0xe9, 0x3d, 0x3b, 0x4e, 0x04, 0x91,
	0xb5, 0xce, 0xfc, 0xfd, 0x3a, 0x93, 0x0c, 0xff, 0x13, 0x44, 0xee, 0x74, 0x66, 0xc6, 0x69, 0xb9,
	0xba, 0x10, 0x01, 0x3d, 0xff, 0x17, 0x44, 0x6e, 0xf4, 0x4f, 0xd9, 0x9e, 0xd6, 0x10, 0x5f, 0xde,
	0x24, 0x58, 0x1d, 0x2f, 0x04, 0x91, 0x95, 0xae, 0x20, 0xde, 0xb3, 0x5f, 0x4f, 0xcf, 0x63, 0x0f,
	0x69, 0x58, 0x6c, 0xd3, 0xe2, 0xa4, 0x70, 0x75, 0x61, 0x04, 0xef, 0xce, 0x13, 0x78, 0xf0, 0xbd,
	0xea, 0x11, 0x47, 0x15, 0x43, 0xab, 0x72, 0x56, 0xb4, 0xdb, 0xbc, 0x97, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x92, 0x3d, 0xce, 0x31, 0xaa, 0x00, 0x00, 0x00,
}
