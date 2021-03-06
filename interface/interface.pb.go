// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interface/interface.proto

/*
Package gnoi_interface is a generated protocol buffer package.

It is generated from these files:
	interface/interface.proto

It has these top-level messages:
	SetLoopbackModeRequest
	SetLoopbackModeResponse
	GetLoopbackModeRequest
	GetLoopbackModeResponse
	ClearInterfaceCountersRequest
	ClearInterfaceCountersResponse
*/
package gnoi_interface

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import gnoi "github.com/openconfig/gnoi"

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

// SetLoopbackModeRequest requests the provide interface be have its loopback mode
// set to mode.  Mode may be vendor specific.  For example, on a transport
// device, available modes are "none", "mac", "phy", "phy_remote",
// "framer_facility", and "framer_terminal".
type SetLoopbackModeRequest struct {
	Interface *gnoi.Path `protobuf:"bytes,1,opt,name=interface" json:"interface,omitempty"`
	Mode      string     `protobuf:"bytes,2,opt,name=mode" json:"mode,omitempty"`
}

func (m *SetLoopbackModeRequest) Reset()                    { *m = SetLoopbackModeRequest{} }
func (m *SetLoopbackModeRequest) String() string            { return proto.CompactTextString(m) }
func (*SetLoopbackModeRequest) ProtoMessage()               {}
func (*SetLoopbackModeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SetLoopbackModeRequest) GetInterface() *gnoi.Path {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *SetLoopbackModeRequest) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

type SetLoopbackModeResponse struct {
}

func (m *SetLoopbackModeResponse) Reset()                    { *m = SetLoopbackModeResponse{} }
func (m *SetLoopbackModeResponse) String() string            { return proto.CompactTextString(m) }
func (*SetLoopbackModeResponse) ProtoMessage()               {}
func (*SetLoopbackModeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetLoopbackModeRequest struct {
	Interface *gnoi.Path `protobuf:"bytes,1,opt,name=interface" json:"interface,omitempty"`
}

func (m *GetLoopbackModeRequest) Reset()                    { *m = GetLoopbackModeRequest{} }
func (m *GetLoopbackModeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLoopbackModeRequest) ProtoMessage()               {}
func (*GetLoopbackModeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetLoopbackModeRequest) GetInterface() *gnoi.Path {
	if m != nil {
		return m.Interface
	}
	return nil
}

type GetLoopbackModeResponse struct {
	Mode string `protobuf:"bytes,1,opt,name=mode" json:"mode,omitempty"`
}

func (m *GetLoopbackModeResponse) Reset()                    { *m = GetLoopbackModeResponse{} }
func (m *GetLoopbackModeResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLoopbackModeResponse) ProtoMessage()               {}
func (*GetLoopbackModeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetLoopbackModeResponse) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

type ClearInterfaceCountersRequest struct {
	Interface []*gnoi.Path `protobuf:"bytes,1,rep,name=interface" json:"interface,omitempty"`
}

func (m *ClearInterfaceCountersRequest) Reset()                    { *m = ClearInterfaceCountersRequest{} }
func (m *ClearInterfaceCountersRequest) String() string            { return proto.CompactTextString(m) }
func (*ClearInterfaceCountersRequest) ProtoMessage()               {}
func (*ClearInterfaceCountersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ClearInterfaceCountersRequest) GetInterface() []*gnoi.Path {
	if m != nil {
		return m.Interface
	}
	return nil
}

type ClearInterfaceCountersResponse struct {
}

func (m *ClearInterfaceCountersResponse) Reset()                    { *m = ClearInterfaceCountersResponse{} }
func (m *ClearInterfaceCountersResponse) String() string            { return proto.CompactTextString(m) }
func (*ClearInterfaceCountersResponse) ProtoMessage()               {}
func (*ClearInterfaceCountersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*SetLoopbackModeRequest)(nil), "gnoi.interface.SetLoopbackModeRequest")
	proto.RegisterType((*SetLoopbackModeResponse)(nil), "gnoi.interface.SetLoopbackModeResponse")
	proto.RegisterType((*GetLoopbackModeRequest)(nil), "gnoi.interface.GetLoopbackModeRequest")
	proto.RegisterType((*GetLoopbackModeResponse)(nil), "gnoi.interface.GetLoopbackModeResponse")
	proto.RegisterType((*ClearInterfaceCountersRequest)(nil), "gnoi.interface.ClearInterfaceCountersRequest")
	proto.RegisterType((*ClearInterfaceCountersResponse)(nil), "gnoi.interface.ClearInterfaceCountersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Interface service

type InterfaceClient interface {
	// SetLoopbackMode is used to set the mode of loopback on a interface.
	SetLoopbackMode(ctx context.Context, in *SetLoopbackModeRequest, opts ...grpc.CallOption) (*SetLoopbackModeResponse, error)
	// GetLoopbackMode is used to get the mode of loopback on a interface.
	GetLoopbackMode(ctx context.Context, in *GetLoopbackModeRequest, opts ...grpc.CallOption) (*GetLoopbackModeResponse, error)
	// ClearInterfaceCounters will reset the counters for the provided interface.
	ClearInterfaceCounters(ctx context.Context, in *ClearInterfaceCountersRequest, opts ...grpc.CallOption) (*ClearInterfaceCountersResponse, error)
}

type interfaceClient struct {
	cc *grpc.ClientConn
}

func NewInterfaceClient(cc *grpc.ClientConn) InterfaceClient {
	return &interfaceClient{cc}
}

func (c *interfaceClient) SetLoopbackMode(ctx context.Context, in *SetLoopbackModeRequest, opts ...grpc.CallOption) (*SetLoopbackModeResponse, error) {
	out := new(SetLoopbackModeResponse)
	err := grpc.Invoke(ctx, "/gnoi.interface.Interface/SetLoopbackMode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interfaceClient) GetLoopbackMode(ctx context.Context, in *GetLoopbackModeRequest, opts ...grpc.CallOption) (*GetLoopbackModeResponse, error) {
	out := new(GetLoopbackModeResponse)
	err := grpc.Invoke(ctx, "/gnoi.interface.Interface/GetLoopbackMode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interfaceClient) ClearInterfaceCounters(ctx context.Context, in *ClearInterfaceCountersRequest, opts ...grpc.CallOption) (*ClearInterfaceCountersResponse, error) {
	out := new(ClearInterfaceCountersResponse)
	err := grpc.Invoke(ctx, "/gnoi.interface.Interface/ClearInterfaceCounters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Interface service

type InterfaceServer interface {
	// SetLoopbackMode is used to set the mode of loopback on a interface.
	SetLoopbackMode(context.Context, *SetLoopbackModeRequest) (*SetLoopbackModeResponse, error)
	// GetLoopbackMode is used to get the mode of loopback on a interface.
	GetLoopbackMode(context.Context, *GetLoopbackModeRequest) (*GetLoopbackModeResponse, error)
	// ClearInterfaceCounters will reset the counters for the provided interface.
	ClearInterfaceCounters(context.Context, *ClearInterfaceCountersRequest) (*ClearInterfaceCountersResponse, error)
}

func RegisterInterfaceServer(s *grpc.Server, srv InterfaceServer) {
	s.RegisterService(&_Interface_serviceDesc, srv)
}

func _Interface_SetLoopbackMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLoopbackModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).SetLoopbackMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.interface.Interface/SetLoopbackMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).SetLoopbackMode(ctx, req.(*SetLoopbackModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interface_GetLoopbackMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoopbackModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).GetLoopbackMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.interface.Interface/GetLoopbackMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).GetLoopbackMode(ctx, req.(*GetLoopbackModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interface_ClearInterfaceCounters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearInterfaceCountersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).ClearInterfaceCounters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.interface.Interface/ClearInterfaceCounters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).ClearInterfaceCounters(ctx, req.(*ClearInterfaceCountersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Interface_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.interface.Interface",
	HandlerType: (*InterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLoopbackMode",
			Handler:    _Interface_SetLoopbackMode_Handler,
		},
		{
			MethodName: "GetLoopbackMode",
			Handler:    _Interface_GetLoopbackMode_Handler,
		},
		{
			MethodName: "ClearInterfaceCounters",
			Handler:    _Interface_ClearInterfaceCounters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interface/interface.proto",
}

func init() { proto.RegisterFile("interface/interface.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2b, 0x49,
	0x2d, 0x4a, 0x4b, 0x4c, 0x4e, 0xd5, 0x87, 0xb3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xf8,
	0xd2, 0xf3, 0xf2, 0x33, 0xf5, 0xe0, 0xa2, 0x52, 0x6a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a,
	0xc9, 0xf9, 0xb9, 0xfa, 0xf9, 0x05, 0xa9, 0x79, 0xc9, 0xf9, 0x79, 0x69, 0x99, 0xe9, 0xfa, 0x20,
	0x55, 0xfa, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x10, 0x7d, 0x4a, 0x61, 0x5c, 0x62, 0xc1, 0xa9, 0x25,
	0x3e, 0xf9, 0xf9, 0x05, 0x49, 0x89, 0xc9, 0xd9, 0xbe, 0xf9, 0x29, 0xa9, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x42, 0x1a, 0x5c, 0x9c, 0x70, 0xe3, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d,
	0xb8, 0xf4, 0xc0, 0xb6, 0x04, 0x24, 0x96, 0x64, 0x04, 0x21, 0x24, 0x85, 0x84, 0xb8, 0x58, 0x72,
	0xf3, 0x53, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x49, 0x2e, 0x71,
	0x0c, 0x73, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95, 0x9c, 0xb8, 0xc4, 0xdc, 0x29, 0xb4, 0x52,
	0x49, 0x97, 0x4b, 0xdc, 0x1d, 0xbb, 0xf1, 0x70, 0xd7, 0x30, 0x22, 0xb9, 0xc6, 0x93, 0x4b, 0xd6,
	0x39, 0x27, 0x35, 0xb1, 0xc8, 0x13, 0x66, 0x80, 0x73, 0x7e, 0x29, 0x88, 0x59, 0x8c, 0xc3, 0x66,
	0x66, 0xdc, 0x36, 0x2b, 0x70, 0xc9, 0xe1, 0x32, 0x0a, 0xe2, 0x00, 0xa3, 0x4b, 0x4c, 0x5c, 0x9c,
	0x70, 0x59, 0xa1, 0x14, 0x2e, 0x7e, 0xb4, 0x80, 0x10, 0x52, 0xd3, 0x43, 0x8d, 0x2c, 0x3d, 0xec,
	0x31, 0x20, 0xa5, 0x4e, 0x50, 0x1d, 0x34, 0x44, 0x19, 0x40, 0xb6, 0xb8, 0x13, 0xb2, 0xc5, 0x9d,
	0x48, 0x5b, 0xdc, 0x71, 0xda, 0x52, 0xc9, 0x25, 0x86, 0xdd, 0xef, 0x42, 0xba, 0xe8, 0x86, 0xe0,
	0x0d, 0x6e, 0x29, 0x3d, 0x62, 0x95, 0xc3, 0xac, 0x4e, 0x62, 0x03, 0x27, 0x57, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xfa, 0x35, 0x00, 0x60, 0x03, 0x03, 0x00, 0x00,
}
