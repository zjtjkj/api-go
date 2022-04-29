// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: api/application/v1/application.proto

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

// ApplicationServiceClient is the client API for ApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplicationServiceClient interface {
	CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error)
	DeleteApp(ctx context.Context, in *DeleteAppRequest, opts ...grpc.CallOption) (*DeleteAppResponse, error)
	UpdateApp(ctx context.Context, in *UpdateAppRequest, opts ...grpc.CallOption) (*UpdateAppResponse, error)
	FindApp(ctx context.Context, in *FindAppRequest, opts ...grpc.CallOption) (*FindAppResponse, error)
	GetApp(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*GetAppResponse, error)
}

type applicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationServiceClient(cc grpc.ClientConnInterface) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error) {
	out := new(CreateAppResponse)
	err := c.cc.Invoke(ctx, "/application.ApplicationService/CreateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) DeleteApp(ctx context.Context, in *DeleteAppRequest, opts ...grpc.CallOption) (*DeleteAppResponse, error) {
	out := new(DeleteAppResponse)
	err := c.cc.Invoke(ctx, "/application.ApplicationService/DeleteApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) UpdateApp(ctx context.Context, in *UpdateAppRequest, opts ...grpc.CallOption) (*UpdateAppResponse, error) {
	out := new(UpdateAppResponse)
	err := c.cc.Invoke(ctx, "/application.ApplicationService/UpdateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) FindApp(ctx context.Context, in *FindAppRequest, opts ...grpc.CallOption) (*FindAppResponse, error) {
	out := new(FindAppResponse)
	err := c.cc.Invoke(ctx, "/application.ApplicationService/FindApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) GetApp(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*GetAppResponse, error) {
	out := new(GetAppResponse)
	err := c.cc.Invoke(ctx, "/application.ApplicationService/GetApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationServiceServer is the server API for ApplicationService service.
// All implementations must embed UnimplementedApplicationServiceServer
// for forward compatibility
type ApplicationServiceServer interface {
	CreateApp(context.Context, *CreateAppRequest) (*CreateAppResponse, error)
	DeleteApp(context.Context, *DeleteAppRequest) (*DeleteAppResponse, error)
	UpdateApp(context.Context, *UpdateAppRequest) (*UpdateAppResponse, error)
	FindApp(context.Context, *FindAppRequest) (*FindAppResponse, error)
	GetApp(context.Context, *GetAppRequest) (*GetAppResponse, error)
	mustEmbedUnimplementedApplicationServiceServer()
}

// UnimplementedApplicationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApplicationServiceServer struct {
}

func (UnimplementedApplicationServiceServer) CreateApp(context.Context, *CreateAppRequest) (*CreateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApp not implemented")
}
func (UnimplementedApplicationServiceServer) DeleteApp(context.Context, *DeleteAppRequest) (*DeleteAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApp not implemented")
}
func (UnimplementedApplicationServiceServer) UpdateApp(context.Context, *UpdateAppRequest) (*UpdateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApp not implemented")
}
func (UnimplementedApplicationServiceServer) FindApp(context.Context, *FindAppRequest) (*FindAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindApp not implemented")
}
func (UnimplementedApplicationServiceServer) GetApp(context.Context, *GetAppRequest) (*GetAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApp not implemented")
}
func (UnimplementedApplicationServiceServer) mustEmbedUnimplementedApplicationServiceServer() {}

// UnsafeApplicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplicationServiceServer will
// result in compilation errors.
type UnsafeApplicationServiceServer interface {
	mustEmbedUnimplementedApplicationServiceServer()
}

func RegisterApplicationServiceServer(s grpc.ServiceRegistrar, srv ApplicationServiceServer) {
	s.RegisterService(&ApplicationService_ServiceDesc, srv)
}

func _ApplicationService_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).CreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/CreateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).CreateApp(ctx, req.(*CreateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_DeleteApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).DeleteApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/DeleteApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).DeleteApp(ctx, req.(*DeleteAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_UpdateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).UpdateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/UpdateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).UpdateApp(ctx, req.(*UpdateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_FindApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).FindApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/FindApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).FindApp(ctx, req.(*FindAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_GetApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).GetApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/application.ApplicationService/GetApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).GetApp(ctx, req.(*GetAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApplicationService_ServiceDesc is the grpc.ServiceDesc for ApplicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApplicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "application.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApp",
			Handler:    _ApplicationService_CreateApp_Handler,
		},
		{
			MethodName: "DeleteApp",
			Handler:    _ApplicationService_DeleteApp_Handler,
		},
		{
			MethodName: "UpdateApp",
			Handler:    _ApplicationService_UpdateApp_Handler,
		},
		{
			MethodName: "FindApp",
			Handler:    _ApplicationService_FindApp_Handler,
		},
		{
			MethodName: "GetApp",
			Handler:    _ApplicationService_GetApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/application/v1/application.proto",
}
