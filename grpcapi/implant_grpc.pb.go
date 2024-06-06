// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: grpcapi/implant.proto

package grpcapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Implant_FetchCommand_FullMethodName    = "/grpcapi.Implant/FetchCommand"
	Implant_SendOutput_FullMethodName      = "/grpcapi.Implant/SendOutput"
	Implant_RegisterImplant_FullMethodName = "/grpcapi.Implant/RegisterImplant"
)

// ImplantClient is the client API for Implant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImplantClient interface {
	FetchCommand(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Command, error)
	SendOutput(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Empty, error)
	RegisterImplant(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Identity, error)
}

type implantClient struct {
	cc grpc.ClientConnInterface
}

func NewImplantClient(cc grpc.ClientConnInterface) ImplantClient {
	return &implantClient{cc}
}

func (c *implantClient) FetchCommand(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Command, error) {
	out := new(Command)
	err := c.cc.Invoke(ctx, Implant_FetchCommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *implantClient) SendOutput(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Implant_SendOutput_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *implantClient) RegisterImplant(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := c.cc.Invoke(ctx, Implant_RegisterImplant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImplantServer is the server API for Implant service.
// All implementations must embed UnimplementedImplantServer
// for forward compatibility
type ImplantServer interface {
	FetchCommand(context.Context, *Identity) (*Command, error)
	SendOutput(context.Context, *Command) (*Empty, error)
	RegisterImplant(context.Context, *Identity) (*Identity, error)
	mustEmbedUnimplementedImplantServer()
}

// UnimplementedImplantServer must be embedded to have forward compatible implementations.
type UnimplementedImplantServer struct {
}

func (UnimplementedImplantServer) FetchCommand(context.Context, *Identity) (*Command, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCommand not implemented")
}
func (UnimplementedImplantServer) SendOutput(context.Context, *Command) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOutput not implemented")
}
func (UnimplementedImplantServer) RegisterImplant(context.Context, *Identity) (*Identity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterImplant not implemented")
}
func (UnimplementedImplantServer) mustEmbedUnimplementedImplantServer() {}

// UnsafeImplantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImplantServer will
// result in compilation errors.
type UnsafeImplantServer interface {
	mustEmbedUnimplementedImplantServer()
}

func RegisterImplantServer(s grpc.ServiceRegistrar, srv ImplantServer) {
	s.RegisterService(&Implant_ServiceDesc, srv)
}

func _Implant_FetchCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImplantServer).FetchCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Implant_FetchCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImplantServer).FetchCommand(ctx, req.(*Identity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Implant_SendOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImplantServer).SendOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Implant_SendOutput_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImplantServer).SendOutput(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Implant_RegisterImplant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImplantServer).RegisterImplant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Implant_RegisterImplant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImplantServer).RegisterImplant(ctx, req.(*Identity))
	}
	return interceptor(ctx, in, info, handler)
}

// Implant_ServiceDesc is the grpc.ServiceDesc for Implant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Implant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.Implant",
	HandlerType: (*ImplantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchCommand",
			Handler:    _Implant_FetchCommand_Handler,
		},
		{
			MethodName: "SendOutput",
			Handler:    _Implant_SendOutput_Handler,
		},
		{
			MethodName: "RegisterImplant",
			Handler:    _Implant_RegisterImplant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcapi/implant.proto",
}

const (
	Admin_RunCommand_FullMethodName = "/grpcapi.Admin/RunCommand"
)

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Command, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Command, error) {
	out := new(Command)
	err := c.cc.Invoke(ctx, Admin_RunCommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	RunCommand(context.Context, *Command) (*Command, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) RunCommand(context.Context, *Command) (*Command, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCommand not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_RunCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).RunCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunCommand",
			Handler:    _Admin_RunCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcapi/implant.proto",
}
