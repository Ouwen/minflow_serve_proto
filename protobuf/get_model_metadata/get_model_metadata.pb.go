// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/get_model_metadata.proto

/*
Package get_model_metadata is a generated protocol buffer package.

It is generated from these files:
	protobuf/get_model_metadata.proto

It has these top-level messages:
	SignatureDefMap
	GetModelMetadataRequest
	GetModelMetadataResponse
*/
package get_model_metadata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import tensorflow11 "github.com/ouwen/minflow_serve_proto/tensorflow/core/protobuf/meta_graph"
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

// Message returned for "signature_def" field.
type SignatureDefMap struct {
	SignatureDef map[string]*tensorflow11.SignatureDef `protobuf:"bytes,1,rep,name=signature_def,json=signatureDef" json:"signature_def,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SignatureDefMap) Reset()                    { *m = SignatureDefMap{} }
func (m *SignatureDefMap) String() string            { return proto.CompactTextString(m) }
func (*SignatureDefMap) ProtoMessage()               {}
func (*SignatureDefMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignatureDefMap) GetSignatureDef() map[string]*tensorflow11.SignatureDef {
	if m != nil {
		return m.SignatureDef
	}
	return nil
}

type GetModelMetadataRequest struct {
	// Model Specification indicating which model we are querying for metadata.
	ModelSpec *tensorflow_serving.ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec" json:"model_spec,omitempty"`
	// Metadata fields to get. Currently supported: "signature_def".
	MetadataField []string `protobuf:"bytes,2,rep,name=metadata_field,json=metadataField" json:"metadata_field,omitempty"`
}

func (m *GetModelMetadataRequest) Reset()                    { *m = GetModelMetadataRequest{} }
func (m *GetModelMetadataRequest) String() string            { return proto.CompactTextString(m) }
func (*GetModelMetadataRequest) ProtoMessage()               {}
func (*GetModelMetadataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetModelMetadataRequest) GetModelSpec() *tensorflow_serving.ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *GetModelMetadataRequest) GetMetadataField() []string {
	if m != nil {
		return m.MetadataField
	}
	return nil
}

type GetModelMetadataResponse struct {
	// Model Specification indicating which model this metadata belongs to.
	ModelSpec *tensorflow_serving.ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec" json:"model_spec,omitempty"`
	// Map of metadata field name to metadata field. The options for metadata
	// field name are listed in GetModelMetadataRequest. Currently supported:
	// "signature_def".
	Metadata map[string]*google_protobuf.Any `protobuf:"bytes,2,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GetModelMetadataResponse) Reset()                    { *m = GetModelMetadataResponse{} }
func (m *GetModelMetadataResponse) String() string            { return proto.CompactTextString(m) }
func (*GetModelMetadataResponse) ProtoMessage()               {}
func (*GetModelMetadataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetModelMetadataResponse) GetModelSpec() *tensorflow_serving.ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *GetModelMetadataResponse) GetMetadata() map[string]*google_protobuf.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*SignatureDefMap)(nil), "tensorflow.serving.SignatureDefMap")
	proto.RegisterType((*GetModelMetadataRequest)(nil), "tensorflow.serving.GetModelMetadataRequest")
	proto.RegisterType((*GetModelMetadataResponse)(nil), "tensorflow.serving.GetModelMetadataResponse")
}

func init() { proto.RegisterFile("protobuf/get_model_metadata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0x6b, 0xdb, 0x30,
	0x1c, 0x45, 0x0e, 0x1b, 0x8b, 0xbc, 0xec, 0x8f, 0x09, 0xcc, 0x33, 0x0c, 0xbc, 0xc0, 0xc0, 0xdb,
	0x41, 0x06, 0x8f, 0xc1, 0x08, 0x3b, 0x6c, 0xa3, 0x69, 0x4f, 0x39, 0xd4, 0x81, 0x42, 0x73, 0x31,
	0x8a, 0xfd, 0xb3, 0x63, 0x6a, 0x4b, 0xae, 0x25, 0x27, 0xf8, 0xd2, 0x4b, 0xbf, 0x5c, 0x3f, 0x52,
	0x8f, 0xc5, 0x76, 0x9c, 0x38, 0x4d, 0xda, 0x4b, 0x6f, 0xf2, 0xfb, 0xbd, 0xa7, 0xf7, 0xa4, 0x27,
	0xe3, 0xaf, 0x59, 0xce, 0x25, 0x5f, 0x14, 0xa1, 0x1d, 0x81, 0xf4, 0x52, 0x1e, 0x40, 0xe2, 0xa5,
	0x20, 0x69, 0x40, 0x25, 0x25, 0xf5, 0x4c, 0xd3, 0x24, 0x30, 0xc1, 0xf3, 0x30, 0xe1, 0x6b, 0x22,
	0x20, 0x5f, 0xc5, 0x2c, 0x32, 0x3e, 0x47, 0x9c, 0x47, 0x09, 0xd8, 0x5b, 0x35, 0x65, 0x65, 0x43,
	0x37, 0xbe, 0xef, 0xe8, 0xb6, 0xcf, 0xf3, 0x0e, 0xa7, 0xda, 0xd7, 0x8b, 0x72, 0x9a, 0x2d, 0x37,
	0xd4, 0xe1, 0x6e, 0x54, 0x19, 0x37, 0xe8, 0xe8, 0x0e, 0xe1, 0xf7, 0xb3, 0x38, 0x62, 0x54, 0x16,
	0x39, 0x9c, 0x40, 0x38, 0xa5, 0x99, 0x36, 0xc7, 0x03, 0xd1, 0x42, 0x5e, 0x00, 0xa1, 0x8e, 0xcc,
	0x9e, 0xa5, 0x3a, 0xbf, 0xc8, 0x61, 0x36, 0xf2, 0x48, 0xbb, 0xf7, 0x3d, 0x61, 0x32, 0x2f, 0xdd,
	0xb7, 0xa2, 0x03, 0x19, 0x97, 0xf8, 0xe3, 0x01, 0x45, 0xfb, 0x80, 0x7b, 0x57, 0x50, 0xea, 0xc8,
	0x44, 0x56, 0xdf, 0xad, 0x96, 0x1a, 0xc1, 0xaf, 0x56, 0x34, 0x29, 0x40, 0x57, 0x4c, 0x64, 0xa9,
	0x8e, 0xde, 0xb5, 0xee, 0xea, 0xdd, 0x86, 0x36, 0x56, 0x7e, 0xa3, 0xd1, 0x0d, 0xfe, 0x74, 0x06,
	0x72, 0x5a, 0x1d, 0x6e, 0xba, 0xb9, 0x54, 0x17, 0xae, 0x0b, 0x10, 0x52, 0xfb, 0x83, 0x71, 0x73,
	0xdb, 0x22, 0x03, 0xbf, 0xf6, 0x51, 0x9d, 0x2f, 0xc7, 0x8e, 0x53, 0xab, 0x67, 0x19, 0xf8, 0x6e,
	0x3f, 0x6d, 0x97, 0xda, 0x37, 0xfc, 0xae, 0x6d, 0xc9, 0x0b, 0x63, 0x48, 0x02, 0x5d, 0x31, 0x7b,
	0x56, 0xdf, 0x1d, 0xb4, 0xe8, 0x69, 0x05, 0x8e, 0x6e, 0x15, 0xac, 0x1f, 0x06, 0x10, 0x19, 0x67,
	0x02, 0x5e, 0x98, 0xe0, 0x02, 0xbf, 0x69, 0xbd, 0x6a, 0x6f, 0xd5, 0x19, 0x1f, 0xd3, 0x3e, 0xe5,
	0x4e, 0x5a, 0xa0, 0x69, 0x64, 0xbb, 0x97, 0x71, 0x8e, 0x07, 0x7b, 0xa3, 0x23, 0x4d, 0xfc, 0xd8,
	0x6f, 0x62, 0x48, 0x9a, 0xc7, 0x48, 0xda, 0xd7, 0x44, 0xfe, 0xb1, 0xb2, 0xd3, 0xc2, 0xff, 0xc9,
	0xfc, 0x6f, 0x14, 0xcb, 0x65, 0xb1, 0x20, 0x3e, 0x4f, 0x6d, 0x5e, 0xac, 0x81, 0xd9, 0x69, 0xcc,
	0xaa, 0x9c, 0x5e, 0x95, 0x13, 0xbc, 0x5a, 0x69, 0x3f, 0xf3, 0x2b, 0xdc, 0x23, 0xb4, 0x78, 0x5d,
	0xcf, 0x7f, 0x3e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x40, 0x4e, 0x4e, 0x33, 0x03, 0x00, 0x00,
}