// Code generated by protoc-gen-go. DO NOT EDIT.
// source: x.proto

package pb

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

type ExampleFuncRequest struct {
	ExampleRequestString string   `protobuf:"bytes,1,opt,name=example_request_string,json=exampleRequestString,proto3" json:"example_request_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExampleFuncRequest) Reset()         { *m = ExampleFuncRequest{} }
func (m *ExampleFuncRequest) String() string { return proto.CompactTextString(m) }
func (*ExampleFuncRequest) ProtoMessage()    {}
func (*ExampleFuncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43f7fb48deeb84e9, []int{0}
}

func (m *ExampleFuncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExampleFuncRequest.Unmarshal(m, b)
}
func (m *ExampleFuncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExampleFuncRequest.Marshal(b, m, deterministic)
}
func (m *ExampleFuncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleFuncRequest.Merge(m, src)
}
func (m *ExampleFuncRequest) XXX_Size() int {
	return xxx_messageInfo_ExampleFuncRequest.Size(m)
}
func (m *ExampleFuncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleFuncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleFuncRequest proto.InternalMessageInfo

func (m *ExampleFuncRequest) GetExampleRequestString() string {
	if m != nil {
		return m.ExampleRequestString
	}
	return ""
}

type ExampleFuncResponse struct {
	ExampleResponseString string   `protobuf:"bytes,1,opt,name=example_response_string,json=exampleResponseString,proto3" json:"example_response_string,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ExampleFuncResponse) Reset()         { *m = ExampleFuncResponse{} }
func (m *ExampleFuncResponse) String() string { return proto.CompactTextString(m) }
func (*ExampleFuncResponse) ProtoMessage()    {}
func (*ExampleFuncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43f7fb48deeb84e9, []int{1}
}

func (m *ExampleFuncResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExampleFuncResponse.Unmarshal(m, b)
}
func (m *ExampleFuncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExampleFuncResponse.Marshal(b, m, deterministic)
}
func (m *ExampleFuncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleFuncResponse.Merge(m, src)
}
func (m *ExampleFuncResponse) XXX_Size() int {
	return xxx_messageInfo_ExampleFuncResponse.Size(m)
}
func (m *ExampleFuncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleFuncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleFuncResponse proto.InternalMessageInfo

func (m *ExampleFuncResponse) GetExampleResponseString() string {
	if m != nil {
		return m.ExampleResponseString
	}
	return ""
}

func init() {
	proto.RegisterType((*ExampleFuncRequest)(nil), "pb.ExampleFuncRequest")
	proto.RegisterType((*ExampleFuncResponse)(nil), "pb.ExampleFuncResponse")
}

func init() { proto.RegisterFile("x.proto", fileDescriptor_43f7fb48deeb84e9) }

var fileDescriptor_43f7fb48deeb84e9 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x62, 0xaf, 0xd0, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xf2, 0xe2, 0x12, 0x72, 0xad, 0x48, 0xcc, 0x2d,
	0xc8, 0x49, 0x75, 0x2b, 0xcd, 0x4b, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x32, 0xe1,
	0x12, 0x4b, 0x85, 0x88, 0xc6, 0x17, 0x41, 0x84, 0xe2, 0x8b, 0x4b, 0x8a, 0x32, 0xf3, 0xd2, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x44, 0xa0, 0xb2, 0x50, 0xf5, 0xc1, 0x60, 0x39, 0x25, 0x5f,
	0x2e, 0x61, 0x14, 0xb3, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xcc, 0xb8, 0xc4, 0x11, 0x86,
	0x41, 0xc4, 0x50, 0x4d, 0x13, 0x85, 0x9b, 0x06, 0x91, 0x85, 0x18, 0x67, 0xe4, 0xcd, 0xc5, 0x0e,
	0x35, 0x4e, 0xc8, 0x81, 0x8b, 0x1b, 0xc9, 0x64, 0x21, 0x31, 0xbd, 0x82, 0x24, 0x3d, 0x4c, 0x67,
	0x4b, 0x89, 0x63, 0x88, 0x43, 0x0c, 0x54, 0x62, 0x70, 0xe2, 0x8e, 0xe2, 0x2c, 0x06, 0x8b, 0xeb,
	0x17, 0x24, 0x25, 0xb1, 0x81, 0xfd, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xb8, 0x67,
	0x07, 0x0a, 0x01, 0x00, 0x00,
}
