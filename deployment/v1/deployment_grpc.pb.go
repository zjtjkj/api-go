// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: api/deployment/v1/deployment.proto

package v1

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

// DeploymentServiceClient is the client API for DeploymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeploymentServiceClient interface {
	StartApp(ctx context.Context, in *StartAppRequest, opts ...grpc.CallOption) (*StartAppResponse, error)
	StopApp(ctx context.Context, in *StopAppRequest, opts ...grpc.CallOption) (*StopAppResponse, error)
}

type deploymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentServiceClient(cc grpc.ClientConnInterface) DeploymentServiceClient {
	return &deploymentServiceClient{cc}
}

func (c *deploymentServiceClient) StartApp(ctx context.Context, in *StartAppRequest, opts ...grpc.CallOption) (*StartAppResponse, error) {
	out := new(StartAppResponse)
	err := c.cc.Invoke(ctx, "/app.DeploymentService/StartApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentServiceClient) StopApp(ctx context.Context, in *StopAppRequest, opts ...grpc.CallOption) (*StopAppResponse, error) {
	out := new(StopAppResponse)
	err := c.cc.Invoke(ctx, "/app.DeploymentService/StopApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentServiceServer is the server API for DeploymentService service.
// All implementations must embed UnimplementedDeploymentServiceServer
// for forward compatibility
type DeploymentServiceServer interface {
	StartApp(context.Context, *StartAppRequest) (*StartAppResponse, error)
	StopApp(context.Context, *StopAppRequest) (*StopAppResponse, error)
	mustEmbedUnimplementedDeploymentServiceServer()
}

// UnimplementedDeploymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeploymentServiceServer struct {
}

func (UnimplementedDeploymentServiceServer) StartApp(context.Context, *StartAppRequest) (*StartAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartApp not implemented")
}
func (UnimplementedDeploymentServiceServer) StopApp(context.Context, *StopAppRequest) (*StopAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopApp not implemented")
}
func (UnimplementedDeploymentServiceServer) mustEmbedUnimplementedDeploymentServiceServer() {}

// UnsafeDeploymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeploymentServiceServer will
// result in compilation errors.
type UnsafeDeploymentServiceServer interface {
	mustEmbedUnimplementedDeploymentServiceServer()
}

func RegisterDeploymentServiceServer(s grpc.ServiceRegistrar, srv DeploymentServiceServer) {
	s.RegisterService(&DeploymentService_ServiceDesc, srv)
}

func _DeploymentService_StartApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServiceServer).StartApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.DeploymentService/StartApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServiceServer).StartApp(ctx, req.(*StartAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentService_StopApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServiceServer).StopApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.DeploymentService/StopApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServiceServer).StopApp(ctx, req.(*StopAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeploymentService_ServiceDesc is the grpc.ServiceDesc for DeploymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeploymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.DeploymentService",
	HandlerType: (*DeploymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartApp",
			Handler:    _DeploymentService_StartApp_Handler,
		},
		{
			MethodName: "StopApp",
			Handler:    _DeploymentService_StopApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/deployment/v1/deployment.proto",
}