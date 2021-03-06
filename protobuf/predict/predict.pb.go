// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/predict.proto

/*
Package predict is a generated protocol buffer package.

It is generated from these files:
	protobuf/predict.proto

It has these top-level messages:
	PredictRequest
	PredictResponse
*/
package predict

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow3 "github.com/ouwen/minflow_serve_proto/tensorflow/core/framework/tensor"
import tensorflow_serving "github.com/ouwen/minflow_serve_proto/protobuf/model"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PredictRequest specifies which TensorFlow model to run, as well as
// how inputs are mapped to tensors and how outputs are filtered before
// returning to user.
type PredictRequest struct {
	// Model Specification.
	ModelSpec *tensorflow_serving.ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec" json:"model_spec,omitempty"`
	// Input tensors.
	// Names of input tensor are alias names. The mapping from aliases to real
	// input tensor names is expected to be stored as named generic signature
	// under the key "inputs" in the model export.
	// Each alias listed in a generic signature named "inputs" should be provided
	// exactly once in order to run the prediction.
	Inputs map[string]*tensorflow3.TensorProto `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Output filter.
	// Names specified are alias names. The mapping from aliases to real output
	// tensor names is expected to be stored as named generic signature under
	// the key "outputs" in the model export.
	// Only tensors specified here will be run/fetched and returned, with the
	// exception that when none is specified, all tensors specified in the
	// named signature will be run/fetched and returned.
	OutputFilter []string `protobuf:"bytes,3,rep,name=output_filter,json=outputFilter" json:"output_filter,omitempty"`
}

func (m *PredictRequest) Reset()                    { *m = PredictRequest{} }
func (m *PredictRequest) String() string            { return proto.CompactTextString(m) }
func (*PredictRequest) ProtoMessage()               {}
func (*PredictRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PredictRequest) GetModelSpec() *tensorflow_serving.ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *PredictRequest) GetInputs() map[string]*tensorflow3.TensorProto {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PredictRequest) GetOutputFilter() []string {
	if m != nil {
		return m.OutputFilter
	}
	return nil
}

// Response for PredictRequest on successful run.
type PredictResponse struct {
	// Output tensors.
	Outputs map[string]*tensorflow3.TensorProto `protobuf:"bytes,1,rep,name=outputs" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PredictResponse) Reset()                    { *m = PredictResponse{} }
func (m *PredictResponse) String() string            { return proto.CompactTextString(m) }
func (*PredictResponse) ProtoMessage()               {}
func (*PredictResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PredictResponse) GetOutputs() map[string]*tensorflow3.TensorProto {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictRequest)(nil), "tensorflow.serving.PredictRequest")
	proto.RegisterType((*PredictResponse)(nil), "tensorflow.serving.PredictResponse")
}

func init() { proto.RegisterFile("protobuf/predict.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x18, 0xc5, 0x49, 0xc2, 0xed, 0xa5, 0xd3, 0xde, 0xab, 0x0c, 0xa2, 0x21, 0x20, 0x94, 0x0a, 0xd2,
	0x8d, 0x13, 0xa9, 0x08, 0x52, 0x5c, 0x09, 0x16, 0x14, 0xc4, 0x32, 0x75, 0xe5, 0x26, 0xb4, 0xe9,
	0x97, 0x1a, 0x9a, 0xcc, 0x8c, 0xf3, 0xa7, 0xa5, 0x4f, 0xe1, 0xbb, 0xf8, 0x74, 0x2e, 0x25, 0x33,
	0xb6, 0xd6, 0x3f, 0xb8, 0x72, 0xf7, 0x71, 0xe6, 0xfb, 0xcd, 0x39, 0x27, 0x19, 0xb4, 0x2b, 0x24,
	0xd7, 0x7c, 0x6c, 0xb2, 0x58, 0x48, 0x98, 0xe4, 0xa9, 0x26, 0x56, 0xc0, 0x58, 0x03, 0x53, 0x5c,
	0x66, 0x05, 0x5f, 0x10, 0x05, 0x72, 0x9e, 0xb3, 0x69, 0x74, 0xf8, 0xae, 0xc5, 0x29, 0x97, 0x10,
	0x67, 0x72, 0x54, 0xc2, 0x82, 0xcb, 0x59, 0xec, 0x4e, 0x1c, 0x1b, 0xed, 0xac, 0xef, 0x2c, 0xf9,
	0x04, 0x0a, 0xa7, 0xb6, 0x9f, 0x7c, 0xf4, 0x7f, 0xe0, 0x3c, 0x28, 0x3c, 0x1a, 0x50, 0x1a, 0x9f,
	0x23, 0x64, 0x37, 0x12, 0x25, 0x20, 0x0d, 0xbd, 0x96, 0xd7, 0x69, 0x74, 0xf7, 0xc9, 0x57, 0x67,
	0x72, 0x53, 0x6d, 0x0d, 0x05, 0xa4, 0xb4, 0x5e, 0xae, 0x46, 0xdc, 0x47, 0xb5, 0x9c, 0x09, 0xa3,
	0x55, 0xe8, 0xb7, 0x82, 0x4e, 0xa3, 0x4b, 0xbe, 0x23, 0x3f, 0x3a, 0x92, 0x2b, 0x0b, 0x5c, 0x32,
	0x2d, 0x97, 0xf4, 0x8d, 0xc6, 0x07, 0xe8, 0x1f, 0x37, 0x5a, 0x18, 0x9d, 0x64, 0x79, 0xa1, 0x41,
	0x86, 0x41, 0x2b, 0xe8, 0xd4, 0x69, 0xd3, 0x89, 0x7d, 0xab, 0x45, 0x14, 0x35, 0x36, 0x58, 0xbc,
	0x8d, 0x82, 0x19, 0x2c, 0x6d, 0xe4, 0x3a, 0xad, 0x46, 0x7c, 0x84, 0xfe, 0xcc, 0x47, 0x85, 0x81,
	0xd0, 0xb7, 0x35, 0xf6, 0x36, 0xc3, 0xdc, 0xd9, 0x71, 0x50, 0x7d, 0x06, 0xea, 0xb6, 0x7a, 0xfe,
	0x99, 0xd7, 0x7e, 0xf6, 0xd0, 0xd6, 0x3a, 0x9f, 0x12, 0x9c, 0x29, 0xc0, 0xd7, 0xe8, 0xaf, 0xf3,
	0x55, 0xa1, 0x67, 0x5b, 0x1d, 0xff, 0xd8, 0xca, 0x51, 0xe4, 0xd6, 0x21, 0xae, 0xd7, 0xea, 0x82,
	0x68, 0x88, 0x9a, 0x9b, 0x07, 0xbf, 0x12, 0xfa, 0xa2, 0x77, 0x7f, 0x3a, 0xcd, 0xf5, 0x83, 0x19,
	0x93, 0x94, 0x97, 0x31, 0x37, 0x0b, 0x60, 0x71, 0x99, 0xb3, 0x0a, 0x49, 0xaa, 0x78, 0x90, 0xd8,
	0xdf, 0x1d, 0x7f, 0x7e, 0x57, 0x2f, 0x9e, 0x37, 0xae, 0x59, 0xf1, 0xe4, 0x35, 0x00, 0x00, 0xff,
	0xff, 0x3d, 0xf3, 0x52, 0xb3, 0x75, 0x02, 0x00, 0x00,
}
