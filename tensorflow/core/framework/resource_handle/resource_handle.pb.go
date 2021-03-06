// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/resource_handle.proto

/*
Package resource_handle is a generated protocol buffer package.

It is generated from these files:
	tensorflow/core/framework/resource_handle.proto

It has these top-level messages:
	ResourceHandleProto
*/
package resource_handle

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

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type ResourceHandleProto struct {
	// Unique name for the device containing the resource.
	Device string `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	// Container in which this resource is placed.
	Container string `protobuf:"bytes,2,opt,name=container" json:"container,omitempty"`
	// Unique name of this resource.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// Hash code for the type of the resource. Is only valid in the same device
	// and in the same execution.
	HashCode uint64 `protobuf:"varint,4,opt,name=hash_code,json=hashCode" json:"hash_code,omitempty"`
	// For debug-only, the name of the type pointed to by this handle, if
	// available.
	MaybeTypeName string `protobuf:"bytes,5,opt,name=maybe_type_name,json=maybeTypeName" json:"maybe_type_name,omitempty"`
}

func (m *ResourceHandleProto) Reset()                    { *m = ResourceHandleProto{} }
func (m *ResourceHandleProto) String() string            { return proto.CompactTextString(m) }
func (*ResourceHandleProto) ProtoMessage()               {}
func (*ResourceHandleProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ResourceHandleProto) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ResourceHandleProto) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *ResourceHandleProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHandleProto) GetHashCode() uint64 {
	if m != nil {
		return m.HashCode
	}
	return 0
}

func (m *ResourceHandleProto) GetMaybeTypeName() string {
	if m != nil {
		return m.MaybeTypeName
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceHandleProto)(nil), "tensorflow.ResourceHandleProto")
}

func init() { proto.RegisterFile("tensorflow/core/framework/resource_handle.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0xc9, 0xff, 0xd7, 0xc1, 0x06, 0x54, 0x88, 0x20, 0x01, 0x5d, 0x0c, 0x2e, 0x64, 0x56,
	0xcd, 0xc2, 0x37, 0x18, 0x37, 0xae, 0x06, 0x29, 0xae, 0xdc, 0x84, 0x34, 0xbd, 0x33, 0x2d, 0x4e,
	0x72, 0xcb, 0x6d, 0x3a, 0xa5, 0x2b, 0x5f, 0xc6, 0x87, 0x74, 0x29, 0x73, 0x19, 0x2c, 0xb8, 0x72,
	0x97, 0x9c, 0x73, 0x3e, 0x38, 0xf7, 0x48, 0x93, 0x20, 0xf6, 0x48, 0xdb, 0x3d, 0x8e, 0xc6, 0x23,
	0x81, 0xd9, 0x92, 0x0b, 0x30, 0x22, 0xbd, 0x1b, 0x82, 0x1e, 0x07, 0xf2, 0x60, 0x1b, 0x17, 0xeb,
	0x3d, 0x14, 0x1d, 0x61, 0x42, 0x25, 0x67, 0xe0, 0xfe, 0x53, 0xc8, 0xeb, 0xf2, 0x94, 0x7a, 0xe6,
	0xd0, 0x0b, 0x67, 0x6e, 0xe4, 0xa2, 0x86, 0x43, 0xeb, 0x41, 0x8b, 0xa5, 0x58, 0xe5, 0xe5, 0xe9,
	0xa7, 0xee, 0x64, 0xee, 0x31, 0x26, 0xd7, 0x46, 0x20, 0xfd, 0x8f, 0xad, 0x59, 0x50, 0x4a, 0x66,
	0xd1, 0x05, 0xd0, 0xff, 0xd9, 0xe0, 0xb7, 0xba, 0x95, 0x79, 0xe3, 0xfa, 0xc6, 0x7a, 0xac, 0x41,
	0x67, 0x4b, 0xb1, 0xca, 0xca, 0xf3, 0xa3, 0xf0, 0x84, 0x35, 0xa8, 0x07, 0x79, 0x15, 0xdc, 0x54,
	0x81, 0x4d, 0x53, 0x07, 0x96, 0xd9, 0x33, 0x66, 0x2f, 0x58, 0x7e, 0x9d, 0x3a, 0xd8, 0xb8, 0x00,
	0xeb, 0x0f, 0xa9, 0x91, 0x76, 0xc5, 0x5c, 0xbc, 0xf8, 0x39, 0x72, 0x7d, 0xf9, 0xab, 0xbf, 0x78,
	0xdb, 0xec, 0xda, 0xd4, 0x0c, 0x55, 0xe1, 0x31, 0x18, 0x1c, 0x46, 0x88, 0x26, 0xb4, 0xf1, 0x48,
	0xd9, 0x1e, 0xe8, 0x00, 0x96, 0x67, 0xf8, 0xfb, 0x6c, 0x5f, 0x42, 0x54, 0x0b, 0x66, 0x1e, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x78, 0xb4, 0xbe, 0x17, 0x6d, 0x01, 0x00, 0x00,
}
