// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consumer.proto

package go_micro_srv_user

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

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_376dcb7dada8cc1b, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_376dcb7dada8cc1b, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.Usrv.user.Request")
	proto.RegisterType((*Response)(nil), "go.micro.Usrv.user.Response")
}

func init() { proto.RegisterFile("consumer.proto", fileDescriptor_376dcb7dada8cc1b) }

var fileDescriptor_376dcb7dada8cc1b = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xce, 0xcf, 0x2b,
	0x2e, 0xcd, 0x4d, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0xcf, 0xd7, 0xcb,
	0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x2b, 0x2e, 0x2a, 0xd3, 0x2b, 0x2d, 0x4e, 0x2d, 0x52, 0xb2, 0xe4,
	0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xa4, 0xb8, 0x38, 0x0a, 0x12, 0x8b,
	0x8b, 0xcb, 0xf3, 0x8b, 0x52, 0x24, 0x98, 0xc0, 0xe2, 0x70, 0xbe, 0x92, 0x12, 0x17, 0x47, 0x50,
	0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x18, 0x17, 0x5b, 0x71, 0x49, 0x62, 0x49, 0x69,
	0x31, 0x54, 0x37, 0x94, 0x67, 0x34, 0x91, 0x91, 0x8b, 0x25, 0xb4, 0x38, 0xb5, 0x48, 0xc8, 0x89,
	0x8b, 0xd5, 0x27, 0x3f, 0x3d, 0x33, 0x4f, 0x48, 0x4a, 0x0f, 0xc3, 0x11, 0x7a, 0x50, 0x17, 0x48,
	0x49, 0x63, 0x95, 0x83, 0x58, 0xa1, 0xc4, 0x20, 0xe4, 0x0a, 0xb2, 0x30, 0x3d, 0xb3, 0xb8, 0x24,
	0xb5, 0x88, 0x02, 0x63, 0x92, 0xd8, 0xc0, 0x81, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x30,
	0xb8, 0x44, 0x58, 0x1e, 0x01, 0x00, 0x00,
}
