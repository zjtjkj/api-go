// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: api/region/v1/region.proto

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

// RegionServiceClient is the client API for RegionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegionServiceClient interface {
	AddRegion(ctx context.Context, in *AddRegionRequest, opts ...grpc.CallOption) (*AddRegionResponse, error)
	DeleteRegion(ctx context.Context, in *DeleteRegionRequest, opts ...grpc.CallOption) (*DeleteRegionResponse, error)
	UpdateRegion(ctx context.Context, in *UpdateRegionRequest, opts ...grpc.CallOption) (*UpdateRegionResponse, error)
	ListRegion(ctx context.Context, in *ListRegionRequest, opts ...grpc.CallOption) (*ListRegionResponse, error)
	GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*GetRegionResponse, error)
	GetListRegion(ctx context.Context, in *GetListRegionRequest, opts ...grpc.CallOption) (*GetListRegionResponse, error)
	GetAncestor(ctx context.Context, in *GetAncestorRequest, opts ...grpc.CallOption) (*GetAncestorResponse, error)
}

type regionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegionServiceClient(cc grpc.ClientConnInterface) RegionServiceClient {
	return &regionServiceClient{cc}
}

func (c *regionServiceClient) AddRegion(ctx context.Context, in *AddRegionRequest, opts ...grpc.CallOption) (*AddRegionResponse, error) {
	out := new(AddRegionResponse)
	err := c.cc.Invoke(ctx, "/region.RegionService/AddRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) DeleteRegion(ctx context.Context, in *DeleteRegionRequest, opts ...grpc.CallOption) (*DeleteRegionResponse, error) {
	out := new(DeleteRegionResponse)
	err := c.cc.Invoke(ctx, "/region.RegionService/DeleteRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) UpdateRegion(ctx context.Context, in *UpdateRegionRequest, opts ...grpc.CallOption) (*UpdateRegionResponse, error) {
	out := new(UpdateRegionResponse)
	err := c.cc.Invoke(ctx, "/region.RegionService/UpdateRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) ListRegion(ctx context.Context, in *ListRegionRequest, opts ...grpc.CallOption) (*ListRegionResponse, error) {
	out := new(ListRegionResponse)
	err := c.cc.Invoke(ctx, "/region.RegionService/ListRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*GetRegionResponse, error) {
	out := new(GetRegionResponse)
	err := c.cc.Invoke(ctx, "/region.RegionService/GetRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) GetListRegion(ctx context.Context, in *GetListRegionRequest, opts ...grpc.CallOption) (*GetListRegionResponse, error) {
	out := new(GetListRegionResponse)
	err := c.cc.Invoke(ctx, "/region.RegionService/GetListRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) GetAncestor(ctx context.Context, in *GetAncestorRequest, opts ...grpc.CallOption) (*GetAncestorResponse, error) {
	out := new(GetAncestorResponse)
	err := c.cc.Invoke(ctx, "/region.RegionService/GetAncestor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegionServiceServer is the server API for RegionService service.
// All implementations must embed UnimplementedRegionServiceServer
// for forward compatibility
type RegionServiceServer interface {
	AddRegion(context.Context, *AddRegionRequest) (*AddRegionResponse, error)
	DeleteRegion(context.Context, *DeleteRegionRequest) (*DeleteRegionResponse, error)
	UpdateRegion(context.Context, *UpdateRegionRequest) (*UpdateRegionResponse, error)
	ListRegion(context.Context, *ListRegionRequest) (*ListRegionResponse, error)
	GetRegion(context.Context, *GetRegionRequest) (*GetRegionResponse, error)
	GetListRegion(context.Context, *GetListRegionRequest) (*GetListRegionResponse, error)
	GetAncestor(context.Context, *GetAncestorRequest) (*GetAncestorResponse, error)
	mustEmbedUnimplementedRegionServiceServer()
}

// UnimplementedRegionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegionServiceServer struct {
}

func (UnimplementedRegionServiceServer) AddRegion(context.Context, *AddRegionRequest) (*AddRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRegion not implemented")
}
func (UnimplementedRegionServiceServer) DeleteRegion(context.Context, *DeleteRegionRequest) (*DeleteRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegion not implemented")
}
func (UnimplementedRegionServiceServer) UpdateRegion(context.Context, *UpdateRegionRequest) (*UpdateRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRegion not implemented")
}
func (UnimplementedRegionServiceServer) ListRegion(context.Context, *ListRegionRequest) (*ListRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRegion not implemented")
}
func (UnimplementedRegionServiceServer) GetRegion(context.Context, *GetRegionRequest) (*GetRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegion not implemented")
}
func (UnimplementedRegionServiceServer) GetListRegion(context.Context, *GetListRegionRequest) (*GetListRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListRegion not implemented")
}
func (UnimplementedRegionServiceServer) GetAncestor(context.Context, *GetAncestorRequest) (*GetAncestorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAncestor not implemented")
}
func (UnimplementedRegionServiceServer) mustEmbedUnimplementedRegionServiceServer() {}

// UnsafeRegionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegionServiceServer will
// result in compilation errors.
type UnsafeRegionServiceServer interface {
	mustEmbedUnimplementedRegionServiceServer()
}

func RegisterRegionServiceServer(s grpc.ServiceRegistrar, srv RegionServiceServer) {
	s.RegisterService(&RegionService_ServiceDesc, srv)
}

func _RegionService_AddRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).AddRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/region.RegionService/AddRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).AddRegion(ctx, req.(*AddRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_DeleteRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).DeleteRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/region.RegionService/DeleteRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).DeleteRegion(ctx, req.(*DeleteRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_UpdateRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).UpdateRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/region.RegionService/UpdateRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).UpdateRegion(ctx, req.(*UpdateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_ListRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).ListRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/region.RegionService/ListRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).ListRegion(ctx, req.(*ListRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_GetRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).GetRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/region.RegionService/GetRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).GetRegion(ctx, req.(*GetRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_GetListRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).GetListRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/region.RegionService/GetListRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).GetListRegion(ctx, req.(*GetListRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_GetAncestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAncestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).GetAncestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/region.RegionService/GetAncestor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).GetAncestor(ctx, req.(*GetAncestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegionService_ServiceDesc is the grpc.ServiceDesc for RegionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "region.RegionService",
	HandlerType: (*RegionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRegion",
			Handler:    _RegionService_AddRegion_Handler,
		},
		{
			MethodName: "DeleteRegion",
			Handler:    _RegionService_DeleteRegion_Handler,
		},
		{
			MethodName: "UpdateRegion",
			Handler:    _RegionService_UpdateRegion_Handler,
		},
		{
			MethodName: "ListRegion",
			Handler:    _RegionService_ListRegion_Handler,
		},
		{
			MethodName: "GetRegion",
			Handler:    _RegionService_GetRegion_Handler,
		},
		{
			MethodName: "GetListRegion",
			Handler:    _RegionService_GetListRegion_Handler,
		},
		{
			MethodName: "GetAncestor",
			Handler:    _RegionService_GetAncestor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/region/v1/region.proto",
}
