// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/type.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/type.proto

It has these top-level messages:
	Response
	Error
*/
package common

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

type Response struct {
	Code int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Error struct {
	Code int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "common.Response")
	proto.RegisterType((*Error)(nil), "common.Error")
}

func init() { proto.RegisterFile("common/type.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83,
	0x08, 0x29, 0x19, 0x70, 0x71, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x09, 0x71,
	0xb1, 0x24, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0xd9, 0x42, 0x02,
	0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x92, 0x2e,
	0x17, 0xab, 0x6b, 0x51, 0x51, 0x7e, 0x11, 0x71, 0xca, 0x93, 0xd8, 0xc0, 0xf6, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xec, 0xaa, 0xda, 0xfb, 0x84, 0x00, 0x00, 0x00,
}
