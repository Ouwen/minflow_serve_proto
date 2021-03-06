// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/node_def.proto

/*
Package node_def is a generated protocol buffer package.

It is generated from these files:
	tensorflow/core/framework/node_def.proto

It has these top-level messages:
	NodeDef
*/
package node_def

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow4 "github.com/ouwen/minflow_serve_proto/tensorflow/core/framework/attr_value"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NodeDef struct {
	// The name given to this operator. Used for naming inputs,
	// logging, visualization, etc.  Unique within a single GraphDef.
	// Must match the regexp "[A-Za-z0-9.][A-Za-z0-9_./]*".
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The operation name.  There may be custom parameters in attrs.
	// Op names starting with an underscore are reserved for internal use.
	Op string `protobuf:"bytes,2,opt,name=op" json:"op,omitempty"`
	// Each input is "node:src_output" with "node" being a string name and
	// "src_output" indicating which output tensor to use from "node". If
	// "src_output" is 0 the ":0" suffix can be omitted.  Regular inputs
	// may optionally be followed by control inputs that have the format
	// "^node".
	Input []string `protobuf:"bytes,3,rep,name=input" json:"input,omitempty"`
	// A (possibly partial) specification for the device on which this
	// node should be placed.
	// The expected syntax for this string is as follows:
	//
	// DEVICE_SPEC ::= PARTIAL_SPEC
	//
	// PARTIAL_SPEC ::= ("/" CONSTRAINT) *
	// CONSTRAINT ::= ("job:" JOB_NAME)
	//              | ("replica:" [1-9][0-9]*)
	//              | ("task:" [1-9][0-9]*)
	//              | ( ("gpu" | "cpu") ":" ([1-9][0-9]* | "*") )
	//
	// Valid values for this string include:
	// * "/job:worker/replica:0/task:1/gpu:3"  (full specification)
	// * "/job:worker/gpu:3"                   (partial specification)
	// * ""                                    (no specification)
	//
	// If the constraints do not resolve to a single device (or if this
	// field is empty or not present), the runtime will attempt to
	// choose a device automatically.
	Device string `protobuf:"bytes,4,opt,name=device" json:"device,omitempty"`
	// Operation-specific graph-construction-time configuration.
	// Note that this should include all attrs defined in the
	// corresponding OpDef, including those with a value matching
	// the default -- this allows the default to change and makes
	// NodeDefs easier to interpret on their own.  However, if
	// an attr with a default is not specified in this list, the
	// default will be used.
	// The "names" (keys) must match the regexp "[a-z][a-z0-9_]+" (and
	// one of the names from the corresponding OpDef's attr field).
	// The values must have a type matching the corresponding OpDef
	// attr's type field.
	// TODO(josh11b): Add some examples here showing best practices.
	Attr map[string]*tensorflow4.AttrValue `protobuf:"bytes,5,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NodeDef) Reset()                    { *m = NodeDef{} }
func (m *NodeDef) String() string            { return proto.CompactTextString(m) }
func (*NodeDef) ProtoMessage()               {}
func (*NodeDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NodeDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeDef) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *NodeDef) GetInput() []string {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *NodeDef) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *NodeDef) GetAttr() map[string]*tensorflow4.AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeDef)(nil), "tensorflow.NodeDef")
}

func init() { proto.RegisterFile("tensorflow/core/framework/node_def.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xd9, 0xa4, 0xad, 0x64, 0x0a, 0x22, 0x8b, 0xca, 0x52, 0x10, 0x8a, 0xa7, 0xa0, 0x90,
	0xc5, 0x7a, 0x11, 0x6f, 0x16, 0xc5, 0x5b, 0x29, 0x39, 0x78, 0xf0, 0x12, 0xd2, 0x66, 0x52, 0x43,
	0x9b, 0x9d, 0xb0, 0xdd, 0x34, 0xf4, 0xcf, 0xfa, 0x3b, 0x3c, 0xca, 0x6e, 0x42, 0xdb, 0x8b, 0x78,
	0x9b, 0x99, 0xfd, 0xde, 0xce, 0xe3, 0x0d, 0x84, 0x06, 0xd5, 0x96, 0x74, 0xbe, 0xa1, 0x46, 0x2e,
	0x49, 0xa3, 0xcc, 0x75, 0x5a, 0x62, 0x43, 0x7a, 0x2d, 0x15, 0x65, 0x98, 0x64, 0x98, 0x47, 0x95,
	0x26, 0x43, 0x1c, 0x8e, 0xe4, 0xe8, 0xee, 0x6f, 0x55, 0x6a, 0x8c, 0x4e, 0x76, 0xe9, 0xa6, 0xc6,
	0x56, 0x77, 0xfb, 0xcd, 0xe0, 0x6c, 0x46, 0x19, 0xbe, 0x62, 0xce, 0x39, 0xf4, 0x54, 0x5a, 0xa2,
	0x60, 0x63, 0x16, 0x06, 0xb1, 0xab, 0xf9, 0x39, 0x78, 0x54, 0x09, 0xcf, 0x4d, 0x3c, 0xaa, 0xf8,
	0x25, 0xf4, 0x0b, 0x55, 0xd5, 0x46, 0xf8, 0x63, 0x3f, 0x0c, 0xe2, 0xb6, 0xe1, 0xd7, 0x30, 0xc8,
	0x70, 0x57, 0x2c, 0x51, 0xf4, 0x1c, 0xd9, 0x75, 0xfc, 0x01, 0x7a, 0x76, 0xa3, 0xe8, 0x8f, 0xfd,
	0x70, 0x38, 0xb9, 0x89, 0x8e, 0xc6, 0xa2, 0x6e, 0x69, 0xf4, 0x62, 0x8c, 0x7e, 0x53, 0x46, 0xef,
	0x63, 0x87, 0x8e, 0x66, 0x10, 0x1c, 0x46, 0xfc, 0x02, 0xfc, 0x35, 0xee, 0x3b, 0x43, 0xb6, 0xe4,
	0xf7, 0xd0, 0x77, 0xf6, 0x9d, 0xa5, 0xe1, 0xe4, 0xea, 0xf4, 0x4b, 0xab, 0xfb, 0xb0, 0x8f, 0x71,
	0xcb, 0x3c, 0x7b, 0x4f, 0x6c, 0xba, 0x05, 0x41, 0x7a, 0x75, 0x8a, 0x1d, 0xd2, 0x98, 0x06, 0xd6,
	0xc4, 0xdc, 0xe6, 0x30, 0x67, 0x9f, 0xef, 0xab, 0xc2, 0x7c, 0xd5, 0x8b, 0x68, 0x49, 0xa5, 0xa4,
	0xba, 0x41, 0x25, 0xcb, 0x42, 0x59, 0x41, 0xb2, 0x45, 0xbd, 0xc3, 0xc4, 0xe5, 0x25, 0xff, 0x3f,
	0xc8, 0x0f, 0x63, 0x8b, 0x81, 0x83, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x58, 0xa1, 0x59,
	0x7e, 0xc0, 0x01, 0x00, 0x00,
}
