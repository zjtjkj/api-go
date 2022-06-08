// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: algo/v1/algo.proto

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

// AlgoClient is the client API for Algo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlgoClient interface {
	SetAuth(ctx context.Context, in *SetAuthRequest, opts ...grpc.CallOption) (*SetAuthResponse, error)
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	ListTask(ctx context.Context, in *ListTaskRequest, opts ...grpc.CallOption) (*ListTaskResponse, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error)
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error)
}

type algoClient struct {
	cc grpc.ClientConnInterface
}

func NewAlgoClient(cc grpc.ClientConnInterface) AlgoClient {
	return &algoClient{cc}
}

func (c *algoClient) SetAuth(ctx context.Context, in *SetAuthRequest, opts ...grpc.CallOption) (*SetAuthResponse, error) {
	out := new(SetAuthResponse)
	err := c.cc.Invoke(ctx, "/api.algo.v1.Algo/SetAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/api.algo.v1.Algo/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/api.algo.v1.Algo/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, "/api.algo.v1.Algo/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) ListTask(ctx context.Context, in *ListTaskRequest, opts ...grpc.CallOption) (*ListTaskResponse, error) {
	out := new(ListTaskResponse)
	err := c.cc.Invoke(ctx, "/api.algo.v1.Algo/ListTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error) {
	out := new(DeleteTaskResponse)
	err := c.cc.Invoke(ctx, "/api.algo.v1.Algo/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algoClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error) {
	out := new(GetTaskResponse)
	err := c.cc.Invoke(ctx, "/api.algo.v1.Algo/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlgoServer is the server API for Algo service.
// All implementations must embed UnimplementedAlgoServer
// for forward compatibility
type AlgoServer interface {
	SetAuth(context.Context, *SetAuthRequest) (*SetAuthResponse, error)
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	ListTask(context.Context, *ListTaskRequest) (*ListTaskResponse, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error)
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)
	mustEmbedUnimplementedAlgoServer()
}

// UnimplementedAlgoServer must be embedded to have forward compatible implementations.
type UnimplementedAlgoServer struct {
}

func (UnimplementedAlgoServer) SetAuth(context.Context, *SetAuthRequest) (*SetAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAuth not implemented")
}
func (UnimplementedAlgoServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedAlgoServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedAlgoServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedAlgoServer) ListTask(context.Context, *ListTaskRequest) (*ListTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTask not implemented")
}
func (UnimplementedAlgoServer) DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedAlgoServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedAlgoServer) mustEmbedUnimplementedAlgoServer() {}

// UnsafeAlgoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlgoServer will
// result in compilation errors.
type UnsafeAlgoServer interface {
	mustEmbedUnimplementedAlgoServer()
}

func RegisterAlgoServer(s grpc.ServiceRegistrar, srv AlgoServer) {
	s.RegisterService(&Algo_ServiceDesc, srv)
}

func _Algo_SetAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).SetAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.algo.v1.Algo/SetAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).SetAuth(ctx, req.(*SetAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.algo.v1.Algo/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.algo.v1.Algo/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.algo.v1.Algo/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_ListTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).ListTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.algo.v1.Algo/ListTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).ListTask(ctx, req.(*ListTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.algo.v1.Algo/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algo_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgoServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.algo.v1.Algo/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgoServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Algo_ServiceDesc is the grpc.ServiceDesc for Algo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Algo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.algo.v1.Algo",
	HandlerType: (*AlgoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAuth",
			Handler:    _Algo_SetAuth_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _Algo_GetInfo_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Algo_GetStatus_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _Algo_CreateTask_Handler,
		},
		{
			MethodName: "ListTask",
			Handler:    _Algo_ListTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _Algo_DeleteTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _Algo_GetTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "algo/v1/algo.proto",
}
