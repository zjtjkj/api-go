// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: schedule/v1/schedule.proto

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

// ScheduleClient is the client API for Schedule service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScheduleClient interface {
	CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*CreateScheduleReply, error)
	UpdateSchedule(ctx context.Context, in *UpdateScheduleRequest, opts ...grpc.CallOption) (*UpdateScheduleReply, error)
	DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...grpc.CallOption) (*DeleteScheduleReply, error)
	GetSchedule(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*GetScheduleReply, error)
	ListSchedule(ctx context.Context, in *ListScheduleRequest, opts ...grpc.CallOption) (*ListScheduleReply, error)
	LockSchedule(ctx context.Context, in *LockScheduleRequest, opts ...grpc.CallOption) (*LockScheduleReply, error)
	UnlockSchedule(ctx context.Context, in *UnlockScheduleRequest, opts ...grpc.CallOption) (*UnlockScheduleReply, error)
}

type scheduleClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduleClient(cc grpc.ClientConnInterface) ScheduleClient {
	return &scheduleClient{cc}
}

func (c *scheduleClient) CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*CreateScheduleReply, error) {
	out := new(CreateScheduleReply)
	err := c.cc.Invoke(ctx, "/api.schedule.v1.Schedule/CreateSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) UpdateSchedule(ctx context.Context, in *UpdateScheduleRequest, opts ...grpc.CallOption) (*UpdateScheduleReply, error) {
	out := new(UpdateScheduleReply)
	err := c.cc.Invoke(ctx, "/api.schedule.v1.Schedule/UpdateSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...grpc.CallOption) (*DeleteScheduleReply, error) {
	out := new(DeleteScheduleReply)
	err := c.cc.Invoke(ctx, "/api.schedule.v1.Schedule/DeleteSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) GetSchedule(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*GetScheduleReply, error) {
	out := new(GetScheduleReply)
	err := c.cc.Invoke(ctx, "/api.schedule.v1.Schedule/GetSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) ListSchedule(ctx context.Context, in *ListScheduleRequest, opts ...grpc.CallOption) (*ListScheduleReply, error) {
	out := new(ListScheduleReply)
	err := c.cc.Invoke(ctx, "/api.schedule.v1.Schedule/ListSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) LockSchedule(ctx context.Context, in *LockScheduleRequest, opts ...grpc.CallOption) (*LockScheduleReply, error) {
	out := new(LockScheduleReply)
	err := c.cc.Invoke(ctx, "/api.schedule.v1.Schedule/LockSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleClient) UnlockSchedule(ctx context.Context, in *UnlockScheduleRequest, opts ...grpc.CallOption) (*UnlockScheduleReply, error) {
	out := new(UnlockScheduleReply)
	err := c.cc.Invoke(ctx, "/api.schedule.v1.Schedule/UnlockSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleServer is the server API for Schedule service.
// All implementations must embed UnimplementedScheduleServer
// for forward compatibility
type ScheduleServer interface {
	CreateSchedule(context.Context, *CreateScheduleRequest) (*CreateScheduleReply, error)
	UpdateSchedule(context.Context, *UpdateScheduleRequest) (*UpdateScheduleReply, error)
	DeleteSchedule(context.Context, *DeleteScheduleRequest) (*DeleteScheduleReply, error)
	GetSchedule(context.Context, *GetScheduleRequest) (*GetScheduleReply, error)
	ListSchedule(context.Context, *ListScheduleRequest) (*ListScheduleReply, error)
	LockSchedule(context.Context, *LockScheduleRequest) (*LockScheduleReply, error)
	UnlockSchedule(context.Context, *UnlockScheduleRequest) (*UnlockScheduleReply, error)
	mustEmbedUnimplementedScheduleServer()
}

// UnimplementedScheduleServer must be embedded to have forward compatible implementations.
type UnimplementedScheduleServer struct {
}

func (UnimplementedScheduleServer) CreateSchedule(context.Context, *CreateScheduleRequest) (*CreateScheduleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchedule not implemented")
}
func (UnimplementedScheduleServer) UpdateSchedule(context.Context, *UpdateScheduleRequest) (*UpdateScheduleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSchedule not implemented")
}
func (UnimplementedScheduleServer) DeleteSchedule(context.Context, *DeleteScheduleRequest) (*DeleteScheduleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSchedule not implemented")
}
func (UnimplementedScheduleServer) GetSchedule(context.Context, *GetScheduleRequest) (*GetScheduleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchedule not implemented")
}
func (UnimplementedScheduleServer) ListSchedule(context.Context, *ListScheduleRequest) (*ListScheduleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchedule not implemented")
}
func (UnimplementedScheduleServer) LockSchedule(context.Context, *LockScheduleRequest) (*LockScheduleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockSchedule not implemented")
}
func (UnimplementedScheduleServer) UnlockSchedule(context.Context, *UnlockScheduleRequest) (*UnlockScheduleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockSchedule not implemented")
}
func (UnimplementedScheduleServer) mustEmbedUnimplementedScheduleServer() {}

// UnsafeScheduleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScheduleServer will
// result in compilation errors.
type UnsafeScheduleServer interface {
	mustEmbedUnimplementedScheduleServer()
}

func RegisterScheduleServer(s grpc.ServiceRegistrar, srv ScheduleServer) {
	s.RegisterService(&Schedule_ServiceDesc, srv)
}

func _Schedule_CreateSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).CreateSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.schedule.v1.Schedule/CreateSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).CreateSchedule(ctx, req.(*CreateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_UpdateSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).UpdateSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.schedule.v1.Schedule/UpdateSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).UpdateSchedule(ctx, req.(*UpdateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_DeleteSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).DeleteSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.schedule.v1.Schedule/DeleteSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).DeleteSchedule(ctx, req.(*DeleteScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_GetSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).GetSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.schedule.v1.Schedule/GetSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).GetSchedule(ctx, req.(*GetScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_ListSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).ListSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.schedule.v1.Schedule/ListSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).ListSchedule(ctx, req.(*ListScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_LockSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).LockSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.schedule.v1.Schedule/LockSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).LockSchedule(ctx, req.(*LockScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Schedule_UnlockSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServer).UnlockSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.schedule.v1.Schedule/UnlockSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServer).UnlockSchedule(ctx, req.(*UnlockScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Schedule_ServiceDesc is the grpc.ServiceDesc for Schedule service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Schedule_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.schedule.v1.Schedule",
	HandlerType: (*ScheduleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSchedule",
			Handler:    _Schedule_CreateSchedule_Handler,
		},
		{
			MethodName: "UpdateSchedule",
			Handler:    _Schedule_UpdateSchedule_Handler,
		},
		{
			MethodName: "DeleteSchedule",
			Handler:    _Schedule_DeleteSchedule_Handler,
		},
		{
			MethodName: "GetSchedule",
			Handler:    _Schedule_GetSchedule_Handler,
		},
		{
			MethodName: "ListSchedule",
			Handler:    _Schedule_ListSchedule_Handler,
		},
		{
			MethodName: "LockSchedule",
			Handler:    _Schedule_LockSchedule_Handler,
		},
		{
			MethodName: "UnlockSchedule",
			Handler:    _Schedule_UnlockSchedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schedule/v1/schedule.proto",
}
