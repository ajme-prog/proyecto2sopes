// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notificador.proto

/*
Package notificador is a generated protocol buffer package.

It is generated from these files:
	notificador.proto

It has these top-level messages:
	DatosRequest
	Responsevalor
	Datosjson
	CorreoResponse
*/
package notificador

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Definimos el request del correo
type DatosRequest struct {
	Nombre        string `protobuf:"bytes,1,opt,name=nombre" json:"nombre,omitempty"`
	Departamento  string `protobuf:"bytes,2,opt,name=departamento" json:"departamento,omitempty"`
	Edad          int32  `protobuf:"varint,3,opt,name=edad" json:"edad,omitempty"`
	Formacontagio string `protobuf:"bytes,4,opt,name=formacontagio" json:"formacontagio,omitempty"`
	Estado        string `protobuf:"bytes,5,opt,name=estado" json:"estado,omitempty"`
}

func (m *DatosRequest) Reset()                    { *m = DatosRequest{} }
func (m *DatosRequest) String() string            { return proto.CompactTextString(m) }
func (*DatosRequest) ProtoMessage()               {}
func (*DatosRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DatosRequest) GetNombre() string {
	if m != nil {
		return m.Nombre
	}
	return ""
}

func (m *DatosRequest) GetDepartamento() string {
	if m != nil {
		return m.Departamento
	}
	return ""
}

func (m *DatosRequest) GetEdad() int32 {
	if m != nil {
		return m.Edad
	}
	return 0
}

func (m *DatosRequest) GetFormacontagio() string {
	if m != nil {
		return m.Formacontagio
	}
	return ""
}

func (m *DatosRequest) GetEstado() string {
	if m != nil {
		return m.Estado
	}
	return ""
}

type Responsevalor struct {
	Enviado bool `protobuf:"varint,1,opt,name=enviado" json:"enviado,omitempty"`
}

func (m *Responsevalor) Reset()                    { *m = Responsevalor{} }
func (m *Responsevalor) String() string            { return proto.CompactTextString(m) }
func (*Responsevalor) ProtoMessage()               {}
func (*Responsevalor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Responsevalor) GetEnviado() bool {
	if m != nil {
		return m.Enviado
	}
	return false
}

type Datosjson struct {
	Jsondatos string `protobuf:"bytes,1,opt,name=jsondatos" json:"jsondatos,omitempty"`
}

func (m *Datosjson) Reset()                    { *m = Datosjson{} }
func (m *Datosjson) String() string            { return proto.CompactTextString(m) }
func (*Datosjson) ProtoMessage()               {}
func (*Datosjson) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Datosjson) GetJsondatos() string {
	if m != nil {
		return m.Jsondatos
	}
	return ""
}

// Definimos el response del correo
type CorreoResponse struct {
	Enviado bool `protobuf:"varint,1,opt,name=enviado" json:"enviado,omitempty"`
}

func (m *CorreoResponse) Reset()                    { *m = CorreoResponse{} }
func (m *CorreoResponse) String() string            { return proto.CompactTextString(m) }
func (*CorreoResponse) ProtoMessage()               {}
func (*CorreoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CorreoResponse) GetEnviado() bool {
	if m != nil {
		return m.Enviado
	}
	return false
}

func init() {
	proto.RegisterType((*DatosRequest)(nil), "notificador.DatosRequest")
	proto.RegisterType((*Responsevalor)(nil), "notificador.Responsevalor")
	proto.RegisterType((*Datosjson)(nil), "notificador.Datosjson")
	proto.RegisterType((*CorreoResponse)(nil), "notificador.CorreoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Notificador service

type NotificadorClient interface {
	Enviardatos(ctx context.Context, in *DatosRequest, opts ...grpc.CallOption) (*Responsevalor, error)
	Enviarjson(ctx context.Context, in *Datosjson, opts ...grpc.CallOption) (*Responsevalor, error)
}

type notificadorClient struct {
	cc *grpc.ClientConn
}

func NewNotificadorClient(cc *grpc.ClientConn) NotificadorClient {
	return &notificadorClient{cc}
}

func (c *notificadorClient) Enviardatos(ctx context.Context, in *DatosRequest, opts ...grpc.CallOption) (*Responsevalor, error) {
	out := new(Responsevalor)
	err := grpc.Invoke(ctx, "/notificador.Notificador/Enviardatos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificadorClient) Enviarjson(ctx context.Context, in *Datosjson, opts ...grpc.CallOption) (*Responsevalor, error) {
	out := new(Responsevalor)
	err := grpc.Invoke(ctx, "/notificador.Notificador/Enviarjson", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notificador service

type NotificadorServer interface {
	Enviardatos(context.Context, *DatosRequest) (*Responsevalor, error)
	Enviarjson(context.Context, *Datosjson) (*Responsevalor, error)
}

func RegisterNotificadorServer(s *grpc.Server, srv NotificadorServer) {
	s.RegisterService(&_Notificador_serviceDesc, srv)
}

func _Notificador_Enviardatos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificadorServer).Enviardatos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notificador.Notificador/Enviardatos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificadorServer).Enviardatos(ctx, req.(*DatosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notificador_Enviarjson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Datosjson)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificadorServer).Enviarjson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notificador.Notificador/Enviarjson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificadorServer).Enviarjson(ctx, req.(*Datosjson))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notificador_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notificador.Notificador",
	HandlerType: (*NotificadorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enviardatos",
			Handler:    _Notificador_Enviardatos_Handler,
		},
		{
			MethodName: "Enviarjson",
			Handler:    _Notificador_Enviarjson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notificador.proto",
}

func init() { proto.RegisterFile("notificador.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x4e, 0x02, 0x31,
	0x10, 0x80, 0xad, 0x02, 0xca, 0x2c, 0x98, 0x38, 0x07, 0x52, 0x89, 0x07, 0xd2, 0x78, 0x00, 0x0f,
	0x1c, 0xf4, 0x0d, 0xfc, 0x3b, 0x7a, 0xe8, 0x1b, 0x14, 0x3a, 0x98, 0x35, 0x6e, 0x67, 0x6d, 0x2b,
	0xef, 0x62, 0xe2, 0xc3, 0x9a, 0x76, 0x21, 0xec, 0xc6, 0xe8, 0x69, 0x77, 0xbe, 0x69, 0xbf, 0xf9,
	0x29, 0x5c, 0x38, 0x8e, 0xe5, 0xa6, 0x5c, 0x1b, 0xcb, 0x7e, 0x59, 0x7b, 0x8e, 0x8c, 0x45, 0x0b,
	0xa9, 0x6f, 0x01, 0xa3, 0x47, 0x13, 0x39, 0x68, 0xfa, 0xf8, 0xa4, 0x10, 0x71, 0x02, 0x03, 0xc7,
	0xd5, 0xca, 0x93, 0x14, 0x33, 0x31, 0x1f, 0xea, 0x5d, 0x84, 0x0a, 0x46, 0x96, 0x6a, 0xe3, 0xa3,
	0xa9, 0xc8, 0x45, 0x96, 0xc7, 0x39, 0xdb, 0x61, 0x88, 0xd0, 0x23, 0x6b, 0xac, 0x3c, 0x99, 0x89,
	0x79, 0x5f, 0xe7, 0x7f, 0xbc, 0x86, 0xf1, 0x86, 0x7d, 0x65, 0xd6, 0xec, 0xa2, 0x79, 0x2d, 0x59,
	0xf6, 0xf2, 0xc5, 0x2e, 0x4c, 0x55, 0x29, 0x44, 0x63, 0x59, 0xf6, 0x9b, 0xaa, 0x4d, 0xa4, 0x16,
	0x30, 0xd6, 0x14, 0x6a, 0x76, 0x81, 0xb6, 0xe6, 0x9d, 0x3d, 0x4a, 0x38, 0x25, 0xb7, 0x2d, 0xd3,
	0xc9, 0xd4, 0xdf, 0x99, 0xde, 0x87, 0x6a, 0x01, 0xc3, 0x3c, 0xc8, 0x5b, 0x60, 0x87, 0x57, 0x30,
	0x4c, 0x5f, 0x9b, 0xc0, 0x6e, 0x90, 0x03, 0x50, 0x37, 0x70, 0xfe, 0xc0, 0xde, 0x13, 0xef, 0xdd,
	0x7f, 0x6b, 0x6f, 0xbf, 0x04, 0x14, 0x2f, 0x87, 0x85, 0xe1, 0x33, 0x14, 0x4f, 0x29, 0xe5, 0xb3,
	0x0a, 0x2f, 0x97, 0xed, 0x05, 0xb7, 0x37, 0x39, 0x9d, 0x76, 0x52, 0x9d, 0x31, 0xd4, 0x11, 0xde,
	0x03, 0x34, 0x9e, 0xdc, 0xef, 0xe4, 0xb7, 0x26, 0xf1, 0xff, 0x1d, 0xab, 0x41, 0x7e, 0xd0, 0xbb,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x44, 0xa8, 0xbd, 0xe5, 0x01, 0x00, 0x00,
}