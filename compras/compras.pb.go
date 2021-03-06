// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: compras.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CompraPyme_EsPrioritario int32

const (
	CompraPyme_NO CompraPyme_EsPrioritario = 0
	CompraPyme_SI CompraPyme_EsPrioritario = 1
)

// Enum value maps for CompraPyme_EsPrioritario.
var (
	CompraPyme_EsPrioritario_name = map[int32]string{
		0: "NO",
		1: "SI",
	}
	CompraPyme_EsPrioritario_value = map[string]int32{
		"NO": 0,
		"SI": 1,
	}
)

func (x CompraPyme_EsPrioritario) Enum() *CompraPyme_EsPrioritario {
	p := new(CompraPyme_EsPrioritario)
	*p = x
	return p
}

func (x CompraPyme_EsPrioritario) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompraPyme_EsPrioritario) Descriptor() protoreflect.EnumDescriptor {
	return file_compras_proto_enumTypes[0].Descriptor()
}

func (CompraPyme_EsPrioritario) Type() protoreflect.EnumType {
	return &file_compras_proto_enumTypes[0]
}

func (x CompraPyme_EsPrioritario) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompraPyme_EsPrioritario.Descriptor instead.
func (CompraPyme_EsPrioritario) EnumDescriptor() ([]byte, []int) {
	return file_compras_proto_rawDescGZIP(), []int{0, 0}
}

type CompraPyme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nombre    string                   `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Costo     float64                  `protobuf:"fixed64,3,opt,name=costo,proto3" json:"costo,omitempty"`
	Origen    string                   `protobuf:"bytes,4,opt,name=origen,proto3" json:"origen,omitempty"`
	Destino   string                   `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
	Prioridad CompraPyme_EsPrioritario `protobuf:"varint,6,opt,name=prioridad,proto3,enum=grpc.CompraPyme_EsPrioritario" json:"prioridad,omitempty"`
}

func (x *CompraPyme) Reset() {
	*x = CompraPyme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compras_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompraPyme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompraPyme) ProtoMessage() {}

func (x *CompraPyme) ProtoReflect() protoreflect.Message {
	mi := &file_compras_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompraPyme.ProtoReflect.Descriptor instead.
func (*CompraPyme) Descriptor() ([]byte, []int) {
	return file_compras_proto_rawDescGZIP(), []int{0}
}

func (x *CompraPyme) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompraPyme) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *CompraPyme) GetCosto() float64 {
	if x != nil {
		return x.Costo
	}
	return 0
}

func (x *CompraPyme) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *CompraPyme) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *CompraPyme) GetPrioridad() CompraPyme_EsPrioritario {
	if x != nil {
		return x.Prioridad
	}
	return CompraPyme_NO
}

type CompraRetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nombre  string  `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Costo   float64 `protobuf:"fixed64,3,opt,name=costo,proto3" json:"costo,omitempty"`
	Origen  string  `protobuf:"bytes,4,opt,name=origen,proto3" json:"origen,omitempty"`
	Destino string  `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *CompraRetail) Reset() {
	*x = CompraRetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compras_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompraRetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompraRetail) ProtoMessage() {}

func (x *CompraRetail) ProtoReflect() protoreflect.Message {
	mi := &file_compras_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompraRetail.ProtoReflect.Descriptor instead.
func (*CompraRetail) Descriptor() ([]byte, []int) {
	return file_compras_proto_rawDescGZIP(), []int{1}
}

func (x *CompraRetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompraRetail) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *CompraRetail) GetCosto() float64 {
	if x != nil {
		return x.Costo
	}
	return 0
}

func (x *CompraRetail) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *CompraRetail) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

type ComprarEnPyme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comprapyme *CompraPyme `protobuf:"bytes,1,opt,name=comprapyme,proto3" json:"comprapyme,omitempty"`
}

func (x *ComprarEnPyme) Reset() {
	*x = ComprarEnPyme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compras_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComprarEnPyme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComprarEnPyme) ProtoMessage() {}

func (x *ComprarEnPyme) ProtoReflect() protoreflect.Message {
	mi := &file_compras_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComprarEnPyme.ProtoReflect.Descriptor instead.
func (*ComprarEnPyme) Descriptor() ([]byte, []int) {
	return file_compras_proto_rawDescGZIP(), []int{2}
}

func (x *ComprarEnPyme) GetComprapyme() *CompraPyme {
	if x != nil {
		return x.Comprapyme
	}
	return nil
}

type ComprarEnPymeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Seguimiento string `protobuf:"bytes,2,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
}

func (x *ComprarEnPymeResp) Reset() {
	*x = ComprarEnPymeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compras_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComprarEnPymeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComprarEnPymeResp) ProtoMessage() {}

func (x *ComprarEnPymeResp) ProtoReflect() protoreflect.Message {
	mi := &file_compras_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComprarEnPymeResp.ProtoReflect.Descriptor instead.
func (*ComprarEnPymeResp) Descriptor() ([]byte, []int) {
	return file_compras_proto_rawDescGZIP(), []int{3}
}

func (x *ComprarEnPymeResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ComprarEnPymeResp) GetSeguimiento() string {
	if x != nil {
		return x.Seguimiento
	}
	return ""
}

type ComprarEnRetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Compraretail *CompraRetail `protobuf:"bytes,1,opt,name=compraretail,proto3" json:"compraretail,omitempty"`
}

func (x *ComprarEnRetail) Reset() {
	*x = ComprarEnRetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compras_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComprarEnRetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComprarEnRetail) ProtoMessage() {}

func (x *ComprarEnRetail) ProtoReflect() protoreflect.Message {
	mi := &file_compras_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComprarEnRetail.ProtoReflect.Descriptor instead.
func (*ComprarEnRetail) Descriptor() ([]byte, []int) {
	return file_compras_proto_rawDescGZIP(), []int{4}
}

func (x *ComprarEnRetail) GetCompraretail() *CompraRetail {
	if x != nil {
		return x.Compraretail
	}
	return nil
}

type ComprarEnRetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ComprarEnRetailResp) Reset() {
	*x = ComprarEnRetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compras_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComprarEnRetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComprarEnRetailResp) ProtoMessage() {}

func (x *ComprarEnRetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_compras_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComprarEnRetailResp.ProtoReflect.Descriptor instead.
func (*ComprarEnRetailResp) Descriptor() ([]byte, []int) {
	return file_compras_proto_rawDescGZIP(), []int{5}
}

func (x *ComprarEnRetailResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_compras_proto protoreflect.FileDescriptor

var file_compras_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0xdb, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x61,
	0x50, 0x79, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x73, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6f, 0x73,
	0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x6f, 0x12, 0x3c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x64, 0x61,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x72, 0x61, 0x50, 0x79, 0x6d, 0x65, 0x2e, 0x45, 0x73, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x64,
	0x61, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x45, 0x73, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61,
	0x72, 0x69, 0x6f, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x4f, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x53,
	0x49, 0x10, 0x01, 0x22, 0x7e, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x61, 0x52, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x73, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6f, 0x73, 0x74,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x6f, 0x22, 0x41, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x61, 0x72, 0x45, 0x6e,
	0x50, 0x79, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x61, 0x70, 0x79,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x61, 0x50, 0x79, 0x6d, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x61, 0x70, 0x79, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x61,
	0x72, 0x45, 0x6e, 0x50, 0x79, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x49, 0x0a,
	0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x61, 0x72, 0x45, 0x6e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x36, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x61, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x61, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x61, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70,
	0x72, 0x61, 0x72, 0x45, 0x6e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0x86, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x69, 0x6f, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x05, 0x43, 0x50, 0x79, 0x6d, 0x65, 0x12, 0x13, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x61, 0x72, 0x45, 0x6e, 0x50, 0x79,
	0x6d, 0x65, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x61,
	0x72, 0x45, 0x6e, 0x50, 0x79, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x07, 0x43,
	0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x61, 0x72, 0x45, 0x6e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x19, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x61, 0x72, 0x45, 0x6e, 0x52, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_compras_proto_rawDescOnce sync.Once
	file_compras_proto_rawDescData = file_compras_proto_rawDesc
)

func file_compras_proto_rawDescGZIP() []byte {
	file_compras_proto_rawDescOnce.Do(func() {
		file_compras_proto_rawDescData = protoimpl.X.CompressGZIP(file_compras_proto_rawDescData)
	})
	return file_compras_proto_rawDescData
}

var file_compras_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_compras_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_compras_proto_goTypes = []interface{}{
	(CompraPyme_EsPrioritario)(0), // 0: grpc.CompraPyme.EsPrioritario
	(*CompraPyme)(nil),            // 1: grpc.CompraPyme
	(*CompraRetail)(nil),          // 2: grpc.CompraRetail
	(*ComprarEnPyme)(nil),         // 3: grpc.ComprarEnPyme
	(*ComprarEnPymeResp)(nil),     // 4: grpc.ComprarEnPymeResp
	(*ComprarEnRetail)(nil),       // 5: grpc.ComprarEnRetail
	(*ComprarEnRetailResp)(nil),   // 6: grpc.ComprarEnRetailResp
}
var file_compras_proto_depIdxs = []int32{
	0, // 0: grpc.CompraPyme.prioridad:type_name -> grpc.CompraPyme.EsPrioritario
	1, // 1: grpc.ComprarEnPyme.comprapyme:type_name -> grpc.CompraPyme
	2, // 2: grpc.ComprarEnRetail.compraretail:type_name -> grpc.CompraRetail
	3, // 3: grpc.ServicioClientes.CPyme:input_type -> grpc.ComprarEnPyme
	5, // 4: grpc.ServicioClientes.CRetail:input_type -> grpc.ComprarEnRetail
	4, // 5: grpc.ServicioClientes.CPyme:output_type -> grpc.ComprarEnPymeResp
	6, // 6: grpc.ServicioClientes.CRetail:output_type -> grpc.ComprarEnRetailResp
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_compras_proto_init() }
func file_compras_proto_init() {
	if File_compras_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_compras_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompraPyme); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_compras_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompraRetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_compras_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComprarEnPyme); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_compras_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComprarEnPymeResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_compras_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComprarEnRetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_compras_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComprarEnRetailResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_compras_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_compras_proto_goTypes,
		DependencyIndexes: file_compras_proto_depIdxs,
		EnumInfos:         file_compras_proto_enumTypes,
		MessageInfos:      file_compras_proto_msgTypes,
	}.Build()
	File_compras_proto = out.File
	file_compras_proto_rawDesc = nil
	file_compras_proto_goTypes = nil
	file_compras_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServicioClientesClient is the client API for ServicioClientes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServicioClientesClient interface {
	CPyme(ctx context.Context, in *ComprarEnPyme, opts ...grpc.CallOption) (*ComprarEnPymeResp, error)
	CRetail(ctx context.Context, in *ComprarEnRetail, opts ...grpc.CallOption) (*ComprarEnRetailResp, error)
}

type servicioClientesClient struct {
	cc grpc.ClientConnInterface
}

func NewServicioClientesClient(cc grpc.ClientConnInterface) ServicioClientesClient {
	return &servicioClientesClient{cc}
}

func (c *servicioClientesClient) CPyme(ctx context.Context, in *ComprarEnPyme, opts ...grpc.CallOption) (*ComprarEnPymeResp, error) {
	out := new(ComprarEnPymeResp)
	err := c.cc.Invoke(ctx, "/grpc.ServicioClientes/CPyme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicioClientesClient) CRetail(ctx context.Context, in *ComprarEnRetail, opts ...grpc.CallOption) (*ComprarEnRetailResp, error) {
	out := new(ComprarEnRetailResp)
	err := c.cc.Invoke(ctx, "/grpc.ServicioClientes/CRetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicioClientesServer is the server API for ServicioClientes service.
type ServicioClientesServer interface {
	CPyme(context.Context, *ComprarEnPyme) (*ComprarEnPymeResp, error)
	CRetail(context.Context, *ComprarEnRetail) (*ComprarEnRetailResp, error)
}

// UnimplementedServicioClientesServer can be embedded to have forward compatible implementations.
type UnimplementedServicioClientesServer struct {
}

func (*UnimplementedServicioClientesServer) CPyme(context.Context, *ComprarEnPyme) (*ComprarEnPymeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CPyme not implemented")
}
func (*UnimplementedServicioClientesServer) CRetail(context.Context, *ComprarEnRetail) (*ComprarEnRetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CRetail not implemented")
}

func RegisterServicioClientesServer(s *grpc.Server, srv ServicioClientesServer) {
	s.RegisterService(&_ServicioClientes_serviceDesc, srv)
}

func _ServicioClientes_CPyme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComprarEnPyme)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioClientesServer).CPyme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ServicioClientes/CPyme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioClientesServer).CPyme(ctx, req.(*ComprarEnPyme))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicioClientes_CRetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComprarEnRetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioClientesServer).CRetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ServicioClientes/CRetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioClientesServer).CRetail(ctx, req.(*ComprarEnRetail))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServicioClientes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ServicioClientes",
	HandlerType: (*ServicioClientesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CPyme",
			Handler:    _ServicioClientes_CPyme_Handler,
		},
		{
			MethodName: "CRetail",
			Handler:    _ServicioClientes_CRetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "compras.proto",
}
