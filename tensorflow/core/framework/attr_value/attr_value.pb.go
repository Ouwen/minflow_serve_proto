// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/attr_value.proto

/*
Package attr_value is a generated protocol buffer package.

It is generated from these files:
	tensorflow/core/framework/attr_value.proto

It has these top-level messages:
	AttrValue
	NameAttrList
*/
package attr_value

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow3 "github.com/ouwen/minflow_serve_proto/tensorflow/core/framework/tensor"
import tensorflow1 "github.com/ouwen/minflow_serve_proto/tensorflow/core/framework/tensor_shape"
import tensorflow2 "github.com/ouwen/minflow_serve_proto/tensorflow/core/framework/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Protocol buffer representing the value for an attr used to configure an Op.
// Comment indicates the corresponding attr type.  Only the field matching the
// attr type may be filled.
type AttrValue struct {
	// Types that are valid to be assigned to Value:
	//	*AttrValue_S
	//	*AttrValue_I
	//	*AttrValue_F
	//	*AttrValue_B
	//	*AttrValue_Type
	//	*AttrValue_Shape
	//	*AttrValue_Tensor
	//	*AttrValue_List
	//	*AttrValue_Func
	//	*AttrValue_Placeholder
	Value isAttrValue_Value `protobuf_oneof:"value"`
}

func (m *AttrValue) Reset()                    { *m = AttrValue{} }
func (m *AttrValue) String() string            { return proto.CompactTextString(m) }
func (*AttrValue) ProtoMessage()               {}
func (*AttrValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isAttrValue_Value interface {
	isAttrValue_Value()
}

type AttrValue_S struct {
	S []byte `protobuf:"bytes,2,opt,name=s,proto3,oneof"`
}
type AttrValue_I struct {
	I int64 `protobuf:"varint,3,opt,name=i,oneof"`
}
type AttrValue_F struct {
	F float32 `protobuf:"fixed32,4,opt,name=f,oneof"`
}
type AttrValue_B struct {
	B bool `protobuf:"varint,5,opt,name=b,oneof"`
}
type AttrValue_Type struct {
	Type tensorflow2.DataType `protobuf:"varint,6,opt,name=type,enum=tensorflow.DataType,oneof"`
}
type AttrValue_Shape struct {
	Shape *tensorflow1.TensorShapeProto `protobuf:"bytes,7,opt,name=shape,oneof"`
}
type AttrValue_Tensor struct {
	Tensor *tensorflow3.TensorProto `protobuf:"bytes,8,opt,name=tensor,oneof"`
}
type AttrValue_List struct {
	List *AttrValue_ListValue `protobuf:"bytes,1,opt,name=list,oneof"`
}
type AttrValue_Func struct {
	Func *NameAttrList `protobuf:"bytes,10,opt,name=func,oneof"`
}
type AttrValue_Placeholder struct {
	Placeholder string `protobuf:"bytes,9,opt,name=placeholder,oneof"`
}

func (*AttrValue_S) isAttrValue_Value()           {}
func (*AttrValue_I) isAttrValue_Value()           {}
func (*AttrValue_F) isAttrValue_Value()           {}
func (*AttrValue_B) isAttrValue_Value()           {}
func (*AttrValue_Type) isAttrValue_Value()        {}
func (*AttrValue_Shape) isAttrValue_Value()       {}
func (*AttrValue_Tensor) isAttrValue_Value()      {}
func (*AttrValue_List) isAttrValue_Value()        {}
func (*AttrValue_Func) isAttrValue_Value()        {}
func (*AttrValue_Placeholder) isAttrValue_Value() {}

func (m *AttrValue) GetValue() isAttrValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *AttrValue) GetS() []byte {
	if x, ok := m.GetValue().(*AttrValue_S); ok {
		return x.S
	}
	return nil
}

func (m *AttrValue) GetI() int64 {
	if x, ok := m.GetValue().(*AttrValue_I); ok {
		return x.I
	}
	return 0
}

func (m *AttrValue) GetF() float32 {
	if x, ok := m.GetValue().(*AttrValue_F); ok {
		return x.F
	}
	return 0
}

func (m *AttrValue) GetB() bool {
	if x, ok := m.GetValue().(*AttrValue_B); ok {
		return x.B
	}
	return false
}

func (m *AttrValue) GetType() tensorflow2.DataType {
	if x, ok := m.GetValue().(*AttrValue_Type); ok {
		return x.Type
	}
	return tensorflow2.DataType_DT_INVALID
}

func (m *AttrValue) GetShape() *tensorflow1.TensorShapeProto {
	if x, ok := m.GetValue().(*AttrValue_Shape); ok {
		return x.Shape
	}
	return nil
}

func (m *AttrValue) GetTensor() *tensorflow3.TensorProto {
	if x, ok := m.GetValue().(*AttrValue_Tensor); ok {
		return x.Tensor
	}
	return nil
}

func (m *AttrValue) GetList() *AttrValue_ListValue {
	if x, ok := m.GetValue().(*AttrValue_List); ok {
		return x.List
	}
	return nil
}

func (m *AttrValue) GetFunc() *NameAttrList {
	if x, ok := m.GetValue().(*AttrValue_Func); ok {
		return x.Func
	}
	return nil
}

func (m *AttrValue) GetPlaceholder() string {
	if x, ok := m.GetValue().(*AttrValue_Placeholder); ok {
		return x.Placeholder
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AttrValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AttrValue_OneofMarshaler, _AttrValue_OneofUnmarshaler, _AttrValue_OneofSizer, []interface{}{
		(*AttrValue_S)(nil),
		(*AttrValue_I)(nil),
		(*AttrValue_F)(nil),
		(*AttrValue_B)(nil),
		(*AttrValue_Type)(nil),
		(*AttrValue_Shape)(nil),
		(*AttrValue_Tensor)(nil),
		(*AttrValue_List)(nil),
		(*AttrValue_Func)(nil),
		(*AttrValue_Placeholder)(nil),
	}
}

func _AttrValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AttrValue)
	// value
	switch x := m.Value.(type) {
	case *AttrValue_S:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.S)
	case *AttrValue_I:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.I))
	case *AttrValue_F:
		b.EncodeVarint(4<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.F)))
	case *AttrValue_B:
		t := uint64(0)
		if x.B {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *AttrValue_Type:
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Type))
	case *AttrValue_Shape:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Shape); err != nil {
			return err
		}
	case *AttrValue_Tensor:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tensor); err != nil {
			return err
		}
	case *AttrValue_List:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.List); err != nil {
			return err
		}
	case *AttrValue_Func:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Func); err != nil {
			return err
		}
	case *AttrValue_Placeholder:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Placeholder)
	case nil:
	default:
		return fmt.Errorf("AttrValue.Value has unexpected type %T", x)
	}
	return nil
}

func _AttrValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AttrValue)
	switch tag {
	case 2: // value.s
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &AttrValue_S{x}
		return true, err
	case 3: // value.i
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &AttrValue_I{int64(x)}
		return true, err
	case 4: // value.f
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &AttrValue_F{math.Float32frombits(uint32(x))}
		return true, err
	case 5: // value.b
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &AttrValue_B{x != 0}
		return true, err
	case 6: // value.type
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &AttrValue_Type{tensorflow2.DataType(x)}
		return true, err
	case 7: // value.shape
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(tensorflow1.TensorShapeProto)
		err := b.DecodeMessage(msg)
		m.Value = &AttrValue_Shape{msg}
		return true, err
	case 8: // value.tensor
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(tensorflow3.TensorProto)
		err := b.DecodeMessage(msg)
		m.Value = &AttrValue_Tensor{msg}
		return true, err
	case 1: // value.list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AttrValue_ListValue)
		err := b.DecodeMessage(msg)
		m.Value = &AttrValue_List{msg}
		return true, err
	case 10: // value.func
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NameAttrList)
		err := b.DecodeMessage(msg)
		m.Value = &AttrValue_Func{msg}
		return true, err
	case 9: // value.placeholder
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &AttrValue_Placeholder{x}
		return true, err
	default:
		return false, nil
	}
}

func _AttrValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AttrValue)
	// value
	switch x := m.Value.(type) {
	case *AttrValue_S:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.S)))
		n += len(x.S)
	case *AttrValue_I:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.I))
	case *AttrValue_F:
		n += proto.SizeVarint(4<<3 | proto.WireFixed32)
		n += 4
	case *AttrValue_B:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += 1
	case *AttrValue_Type:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Type))
	case *AttrValue_Shape:
		s := proto.Size(x.Shape)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AttrValue_Tensor:
		s := proto.Size(x.Tensor)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AttrValue_List:
		s := proto.Size(x.List)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AttrValue_Func:
		s := proto.Size(x.Func)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AttrValue_Placeholder:
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Placeholder)))
		n += len(x.Placeholder)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// LINT.IfChange
type AttrValue_ListValue struct {
	S      [][]byte                        `protobuf:"bytes,2,rep,name=s,proto3" json:"s,omitempty"`
	I      []int64                         `protobuf:"varint,3,rep,packed,name=i" json:"i,omitempty"`
	F      []float32                       `protobuf:"fixed32,4,rep,packed,name=f" json:"f,omitempty"`
	B      []bool                          `protobuf:"varint,5,rep,packed,name=b" json:"b,omitempty"`
	Type   []tensorflow2.DataType          `protobuf:"varint,6,rep,packed,name=type,enum=tensorflow.DataType" json:"type,omitempty"`
	Shape  []*tensorflow1.TensorShapeProto `protobuf:"bytes,7,rep,name=shape" json:"shape,omitempty"`
	Tensor []*tensorflow3.TensorProto      `protobuf:"bytes,8,rep,name=tensor" json:"tensor,omitempty"`
	Func   []*NameAttrList                 `protobuf:"bytes,9,rep,name=func" json:"func,omitempty"`
}

func (m *AttrValue_ListValue) Reset()                    { *m = AttrValue_ListValue{} }
func (m *AttrValue_ListValue) String() string            { return proto.CompactTextString(m) }
func (*AttrValue_ListValue) ProtoMessage()               {}
func (*AttrValue_ListValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *AttrValue_ListValue) GetS() [][]byte {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *AttrValue_ListValue) GetI() []int64 {
	if m != nil {
		return m.I
	}
	return nil
}

func (m *AttrValue_ListValue) GetF() []float32 {
	if m != nil {
		return m.F
	}
	return nil
}

func (m *AttrValue_ListValue) GetB() []bool {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *AttrValue_ListValue) GetType() []tensorflow2.DataType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *AttrValue_ListValue) GetShape() []*tensorflow1.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *AttrValue_ListValue) GetTensor() []*tensorflow3.TensorProto {
	if m != nil {
		return m.Tensor
	}
	return nil
}

func (m *AttrValue_ListValue) GetFunc() []*NameAttrList {
	if m != nil {
		return m.Func
	}
	return nil
}

// A list of attr names and their values. The whole list is attached
// with a string name.  E.g., MatMul[T=float].
type NameAttrList struct {
	Name string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Attr map[string]*AttrValue `protobuf:"bytes,2,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NameAttrList) Reset()                    { *m = NameAttrList{} }
func (m *NameAttrList) String() string            { return proto.CompactTextString(m) }
func (*NameAttrList) ProtoMessage()               {}
func (*NameAttrList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NameAttrList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NameAttrList) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func init() {
	proto.RegisterType((*AttrValue)(nil), "tensorflow.AttrValue")
	proto.RegisterType((*AttrValue_ListValue)(nil), "tensorflow.AttrValue.ListValue")
	proto.RegisterType((*NameAttrList)(nil), "tensorflow.NameAttrList")
}

func init() { proto.RegisterFile("tensorflow/core/framework/attr_value.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0x87, 0x3b, 0x99, 0x74, 0xb7, 0x99, 0x96, 0xb5, 0x0c, 0x8a, 0x43, 0x11, 0x0c, 0x05, 0x65,
	0x58, 0x4b, 0x82, 0xf5, 0x0f, 0xe2, 0x9d, 0x45, 0xa1, 0x82, 0x2c, 0x4b, 0x5c, 0xbc, 0xf0, 0xa6,
	0xa4, 0x75, 0xb2, 0x0d, 0xdb, 0x66, 0xca, 0x64, 0xba, 0xa5, 0xf8, 0x00, 0xde, 0xfa, 0x1c, 0x3e,
	0xa1, 0x97, 0x72, 0xce, 0x74, 0xd3, 0x80, 0x76, 0xd7, 0xbb, 0x39, 0x27, 0xdf, 0x6f, 0x72, 0xfa,
	0x65, 0xa6, 0xec, 0xd4, 0xaa, 0xa2, 0xd4, 0x26, 0x5b, 0xe8, 0x4d, 0x3c, 0xd3, 0x46, 0xc5, 0x99,
	0x49, 0x97, 0x6a, 0xa3, 0xcd, 0x55, 0x9c, 0x5a, 0x6b, 0x26, 0xd7, 0xe9, 0x62, 0xad, 0xa2, 0x95,
	0xd1, 0x56, 0x73, 0xb6, 0x67, 0x7b, 0x4f, 0x0f, 0xe7, 0xdc, 0x13, 0x97, 0xe9, 0x0d, 0xee, 0xe2,
	0x26, 0xe5, 0x3c, 0x5d, 0xed, 0xde, 0xd0, 0x7b, 0x72, 0x0b, 0xbd, 0x5d, 0xa9, 0xd2, 0x61, 0xfd,
	0x1f, 0x4d, 0x16, 0xbc, 0xb3, 0xd6, 0x7c, 0x81, 0xe1, 0xf8, 0x09, 0x23, 0xa5, 0xf0, 0x42, 0x22,
	0x3b, 0xe3, 0x46, 0x42, 0x4a, 0xa8, 0x73, 0x41, 0x43, 0x22, 0x29, 0xd4, 0x39, 0xd4, 0x99, 0xf0,
	0x43, 0x22, 0x3d, 0xa8, 0x33, 0xa8, 0xa7, 0xa2, 0x19, 0x12, 0xd9, 0x82, 0x7a, 0xca, 0x4f, 0x99,
	0x0f, 0x9b, 0x8b, 0xa3, 0x90, 0xc8, 0x93, 0xe1, 0xfd, 0x68, 0x3f, 0x43, 0xf4, 0x3e, 0xb5, 0xe9,
	0xc5, 0x76, 0xa5, 0xc6, 0x8d, 0x04, 0x19, 0xfe, 0x92, 0x35, 0x71, 0x5e, 0x71, 0x1c, 0x12, 0xd9,
	0x1e, 0x3e, 0xaa, 0xc3, 0x17, 0xb8, 0xfc, 0x0c, 0x8f, 0xcf, 0x61, 0xcc, 0x71, 0x23, 0x71, 0x30,
	0x7f, 0xce, 0x8e, 0x1c, 0x27, 0x5a, 0x18, 0x7b, 0xf8, 0x77, 0xec, 0x26, 0xb1, 0x03, 0xf9, 0x2b,
	0xe6, 0x2f, 0xf2, 0xd2, 0x0a, 0x82, 0x81, 0xc7, 0xf5, 0x40, 0xf5, 0xcb, 0xa3, 0x4f, 0x79, 0x69,
	0x71, 0x05, 0xf3, 0x01, 0xce, 0x23, 0xe6, 0x67, 0xeb, 0x62, 0x26, 0x18, 0xc6, 0x44, 0x3d, 0x76,
	0x96, 0x2e, 0x15, 0x44, 0x21, 0x04, 0x3c, 0x70, 0xbc, 0xcf, 0xda, 0xab, 0x45, 0x3a, 0x53, 0x73,
	0xbd, 0xf8, 0xa6, 0x8c, 0x08, 0x42, 0x22, 0x83, 0x71, 0x23, 0xa9, 0x37, 0x7b, 0x3f, 0x3d, 0x16,
	0x54, 0x6f, 0xe2, 0x1d, 0x67, 0x9b, 0xca, 0x0e, 0xb8, 0xee, 0x3a, 0xd7, 0x54, 0xd2, 0x91, 0xd7,
	0x25, 0x60, 0xbb, 0xeb, 0x6c, 0x53, 0xe9, 0xb9, 0x4e, 0x06, 0x1d, 0xf0, 0x4d, 0x65, 0xcb, 0x75,
	0xa6, 0x7c, 0x50, 0x19, 0xa7, 0x87, 0x8c, 0x23, 0xea, 0x9c, 0x0f, 0xf7, 0xce, 0xe9, 0x5d, 0xce,
	0x6f, 0x8c, 0xc7, 0x35, 0xe3, 0xf4, 0x16, 0xe3, 0x95, 0xef, 0xc1, 0x4e, 0x5c, 0x80, 0xf8, 0x41,
	0x71, 0x4e, 0xdb, 0xe8, 0x98, 0x35, 0xf1, 0x62, 0xf4, 0x7f, 0x11, 0xd6, 0xa9, 0x3f, 0xe7, 0x9c,
	0xf9, 0x45, 0xba, 0x54, 0xf8, 0xdd, 0x82, 0x04, 0xd7, 0xfc, 0x35, 0xf3, 0xe1, 0x2e, 0xa1, 0xb5,
	0xf6, 0xb0, 0x7f, 0x68, 0x6f, 0xfc, 0xb0, 0x1f, 0x0a, 0x6b, 0xb6, 0x09, 0xf2, 0xbd, 0x33, 0x77,
	0xca, 0xb1, 0xc5, 0xbb, 0x8c, 0x5e, 0xa9, 0xed, 0x6e, 0x5f, 0x58, 0xf2, 0x67, 0xbb, 0x21, 0xf0,
	0xec, 0xb7, 0x87, 0x0f, 0xfe, 0x79, 0x46, 0x12, 0xc7, 0xbc, 0xf5, 0xde, 0x90, 0xd1, 0x77, 0x26,
	0xb4, 0xb9, 0xac, 0x63, 0xd5, 0xf5, 0x1a, 0xdd, 0xab, 0x12, 0xe8, 0xa5, 0x3c, 0x27, 0x5f, 0x3f,
	0x5e, 0xe6, 0x76, 0xbe, 0x9e, 0x46, 0x33, 0xbd, 0x8c, 0xf5, 0x7a, 0xa3, 0x8a, 0x78, 0x99, 0x17,
	0x10, 0x9b, 0x94, 0xca, 0x5c, 0xab, 0x09, 0xde, 0xc5, 0xf8, 0x7f, 0xfe, 0x3f, 0x7e, 0x13, 0x32,
	0x3d, 0x42, 0xfc, 0xc5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x6b, 0xcd, 0x2f, 0x71, 0x04,
	0x00, 0x00,
}
