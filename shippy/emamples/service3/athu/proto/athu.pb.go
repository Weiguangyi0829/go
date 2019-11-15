// Code generated by protoc-gen-go. DO NOT EDIT.
// source: athu.proto

package go_micro_srv_athu

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
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce99ddae7552a6d, []int{0}
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

func (m *Request) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce99ddae7552a6d, []int{1}
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

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.athu.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.athu.Response")
}

func init() { proto.RegisterFile("athu.proto", fileDescriptor_3ce99ddae7552a6d) }

var fileDescriptor_3ce99ddae7552a6d = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xbd, 0x4e, 0xc4, 0x30,
	0x10, 0x84, 0x15, 0xfe, 0xce, 0xb7, 0xdd, 0x59, 0x14, 0x56, 0x68, 0xd0, 0x55, 0x54, 0x2e, 0xa0,
	0xa4, 0x42, 0x5c, 0x09, 0x12, 0xb2, 0x10, 0xbd, 0xb1, 0x16, 0xce, 0x3a, 0xe2, 0x0d, 0xbb, 0x36,
	0xaf, 0xcb, 0xab, 0xa0, 0x38, 0x81, 0x86, 0x74, 0x29, 0xbf, 0x99, 0x9d, 0xd1, 0x48, 0x0b, 0xe0,
	0xf3, 0xbe, 0xd8, 0x9e, 0x29, 0x93, 0xde, 0xbc, 0x93, 0xed, 0x62, 0x60, 0xb2, 0xc2, 0x5f, 0x76,
	0x30, 0xb6, 0xb7, 0xb0, 0x72, 0xf8, 0x59, 0x50, 0xb2, 0x6e, 0x41, 0x15, 0x41, 0x4e, 0xbe, 0x43,
	0xd3, 0x5c, 0x36, 0x57, 0x6b, 0xf7, 0xc7, 0xfa, 0x1c, 0x4e, 0x33, 0x1d, 0x30, 0x99, 0xa3, 0x6a,
	0x8c, 0xb0, 0x7d, 0x02, 0xe5, 0x50, 0x7a, 0x4a, 0x82, 0xda, 0xc0, 0x4a, 0x4a, 0x08, 0x28, 0x52,
	0xc3, 0xca, 0xfd, 0xe2, 0x7c, 0x76, 0x50, 0x91, 0x99, 0xd8, 0x1c, 0x8f, 0x6a, 0x85, 0xeb, 0xef,
	0x06, 0x4e, 0xee, 0xf2, 0xbe, 0xe8, 0x1d, 0xac, 0x1f, 0xfd, 0x01, 0x9f, 0xeb, 0x6d, 0x6b, 0xff,
	0x0d, 0xb7, 0xd3, 0xea, 0xf6, 0x62, 0xd6, 0x9b, 0x46, 0xdd, 0x83, 0xda, 0xe1, 0xc7, 0xc2, 0x92,
	0x07, 0xd8, 0xbc, 0x20, 0xc7, 0xb7, 0x18, 0x7c, 0x8e, 0x94, 0x96, 0xb5, 0xbd, 0x9e, 0xd5, 0x57,
	0xdc, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xee, 0x3f, 0x05, 0xe5, 0x98, 0x01, 0x00, 0x00,
}