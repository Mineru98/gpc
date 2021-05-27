// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// GPCClient is the client API for GPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GPCClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	CallProgram(ctx context.Context, in *ProgramCallRequest, opts ...grpc.CallOption) (*ProgramCallReply, error)
}

type gPCClient struct {
	cc grpc.ClientConnInterface
}

func NewGPCClient(cc grpc.ClientConnInterface) GPCClient {
	return &gPCClient{cc}
}

func (c *gPCClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/gpc.GPC/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gPCClient) CallProgram(ctx context.Context, in *ProgramCallRequest, opts ...grpc.CallOption) (*ProgramCallReply, error) {
	out := new(ProgramCallReply)
	err := c.cc.Invoke(ctx, "/gpc.GPC/CallProgram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GPCServer is the server API for GPC service.
// All implementations must embed UnimplementedGPCServer
// for forward compatibility
type GPCServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	CallProgram(context.Context, *ProgramCallRequest) (*ProgramCallReply, error)
	mustEmbedUnimplementedGPCServer()
}

// UnimplementedGPCServer must be embedded to have forward compatible implementations.
type UnimplementedGPCServer struct {
}

func (UnimplementedGPCServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGPCServer) CallProgram(context.Context, *ProgramCallRequest) (*ProgramCallReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallProgram not implemented")
}
func (UnimplementedGPCServer) mustEmbedUnimplementedGPCServer() {}

// UnsafeGPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GPCServer will
// result in compilation errors.
type UnsafeGPCServer interface {
	mustEmbedUnimplementedGPCServer()
}

func RegisterGPCServer(s grpc.ServiceRegistrar, srv GPCServer) {
	s.RegisterService(&GPC_ServiceDesc, srv)
}

func _GPC_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPCServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gpc.GPC/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPCServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GPC_CallProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProgramCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GPCServer).CallProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gpc.GPC/CallProgram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GPCServer).CallProgram(ctx, req.(*ProgramCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GPC_ServiceDesc is the grpc.ServiceDesc for GPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gpc.GPC",
	HandlerType: (*GPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GPC_SayHello_Handler,
		},
		{
			MethodName: "CallProgram",
			Handler:    _GPC_CallProgram_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gpc.proto",
}
