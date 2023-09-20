// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package my

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

// MyRPCServiceClient is the client API for MyRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyRPCServiceClient interface {
	GetMy(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error)
}

type myRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyRPCServiceClient(cc grpc.ClientConnInterface) MyRPCServiceClient {
	return &myRPCServiceClient{cc}
}

func (c *myRPCServiceClient) GetMy(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, "/my.MyRPCService/GetMy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyRPCServiceServer is the server API for MyRPCService service.
// All implementations must embed UnimplementedMyRPCServiceServer
// for forward compatibility
type MyRPCServiceServer interface {
	GetMy(context.Context, *MyRequest) (*MyResponse, error)
	mustEmbedUnimplementedMyRPCServiceServer()
}

// UnimplementedMyRPCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMyRPCServiceServer struct {
}

func (UnimplementedMyRPCServiceServer) GetMy(context.Context, *MyRequest) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMy not implemented")
}
func (UnimplementedMyRPCServiceServer) mustEmbedUnimplementedMyRPCServiceServer() {}

// UnsafeMyRPCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyRPCServiceServer will
// result in compilation errors.
type UnsafeMyRPCServiceServer interface {
	mustEmbedUnimplementedMyRPCServiceServer()
}

func RegisterMyRPCServiceServer(s grpc.ServiceRegistrar, srv MyRPCServiceServer) {
	s.RegisterService(&MyRPCService_ServiceDesc, srv)
}

func _MyRPCService_GetMy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyRPCServiceServer).GetMy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/my.MyRPCService/GetMy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyRPCServiceServer).GetMy(ctx, req.(*MyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MyRPCService_ServiceDesc is the grpc.ServiceDesc for MyRPCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyRPCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "my.MyRPCService",
	HandlerType: (*MyRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMy",
			Handler:    _MyRPCService_GetMy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "my.proto",
}