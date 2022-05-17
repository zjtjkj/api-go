// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: api/storage/v1/storage.proto

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

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	CreateStorage(ctx context.Context, in *CreateStorageRequest, opts ...grpc.CallOption) (*CreateStorageResponse, error)
	DeleteStorage(ctx context.Context, in *DeleteStorageRequest, opts ...grpc.CallOption) (*DeleteStorageResponse, error)
	UpdateStorage(ctx context.Context, in *UpdateStorageSizeRequest, opts ...grpc.CallOption) (*UpdateStorageSizeResponse, error)
	FindStorage(ctx context.Context, in *FindStorageRequest, opts ...grpc.CallOption) (*FindStorageResponse, error)
	GetStorage(ctx context.Context, in *GetStorageRequest, opts ...grpc.CallOption) (*GetStorageResponse, error)
	GetStorageState(ctx context.Context, in *GetStorageStateRequest, opts ...grpc.CallOption) (*GetStorageStateResponse, error)
	GetStorageMedium(ctx context.Context, in *GetStorageMediumRequest, opts ...grpc.CallOption) (*GetStorageMediumResponse, error)
	UpdateStorageState(ctx context.Context, in *UpdateStorageStateRequest, opts ...grpc.CallOption) (*UpdateStorageStateResponse, error)
	CreateAllocation(ctx context.Context, in *CreateAllocationRequest, opts ...grpc.CallOption) (*CreateAllocationResponse, error)
	DeleteAllocation(ctx context.Context, in *DeleteAllocationRequest, opts ...grpc.CallOption) (*DeleteAllocationResponse, error)
	UpdateAllocation(ctx context.Context, in *UpdateAllocationRequest, opts ...grpc.CallOption) (*UpdateAllocationResponse, error)
	FindAllocation(ctx context.Context, in *FindAllocationRequest, opts ...grpc.CallOption) (*FindAllocationResponse, error)
	GetAllocation(ctx context.Context, in *GetAllocationRequest, opts ...grpc.CallOption) (*GetAllocationResponse, error)
	ListAllocation(ctx context.Context, in *ListAllocationRequest, opts ...grpc.CallOption) (*ListAllocationResponse, error)
	RequestSpace(ctx context.Context, in *RequestSpaceRequest, opts ...grpc.CallOption) (*RequestSpaceResponse, error)
	ReleaseSpace(ctx context.Context, in *ReleaseSpaceRequest, opts ...grpc.CallOption) (*ReleaseSpaceResponse, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) CreateStorage(ctx context.Context, in *CreateStorageRequest, opts ...grpc.CallOption) (*CreateStorageResponse, error) {
	out := new(CreateStorageResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/CreateStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) DeleteStorage(ctx context.Context, in *DeleteStorageRequest, opts ...grpc.CallOption) (*DeleteStorageResponse, error) {
	out := new(DeleteStorageResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/DeleteStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) UpdateStorage(ctx context.Context, in *UpdateStorageSizeRequest, opts ...grpc.CallOption) (*UpdateStorageSizeResponse, error) {
	out := new(UpdateStorageSizeResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/UpdateStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) FindStorage(ctx context.Context, in *FindStorageRequest, opts ...grpc.CallOption) (*FindStorageResponse, error) {
	out := new(FindStorageResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/FindStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetStorage(ctx context.Context, in *GetStorageRequest, opts ...grpc.CallOption) (*GetStorageResponse, error) {
	out := new(GetStorageResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/GetStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetStorageState(ctx context.Context, in *GetStorageStateRequest, opts ...grpc.CallOption) (*GetStorageStateResponse, error) {
	out := new(GetStorageStateResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/GetStorageState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetStorageMedium(ctx context.Context, in *GetStorageMediumRequest, opts ...grpc.CallOption) (*GetStorageMediumResponse, error) {
	out := new(GetStorageMediumResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/GetStorageMedium", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) UpdateStorageState(ctx context.Context, in *UpdateStorageStateRequest, opts ...grpc.CallOption) (*UpdateStorageStateResponse, error) {
	out := new(UpdateStorageStateResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/UpdateStorageState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) CreateAllocation(ctx context.Context, in *CreateAllocationRequest, opts ...grpc.CallOption) (*CreateAllocationResponse, error) {
	out := new(CreateAllocationResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/CreateAllocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) DeleteAllocation(ctx context.Context, in *DeleteAllocationRequest, opts ...grpc.CallOption) (*DeleteAllocationResponse, error) {
	out := new(DeleteAllocationResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/DeleteAllocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) UpdateAllocation(ctx context.Context, in *UpdateAllocationRequest, opts ...grpc.CallOption) (*UpdateAllocationResponse, error) {
	out := new(UpdateAllocationResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/UpdateAllocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) FindAllocation(ctx context.Context, in *FindAllocationRequest, opts ...grpc.CallOption) (*FindAllocationResponse, error) {
	out := new(FindAllocationResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/FindAllocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetAllocation(ctx context.Context, in *GetAllocationRequest, opts ...grpc.CallOption) (*GetAllocationResponse, error) {
	out := new(GetAllocationResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/GetAllocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) ListAllocation(ctx context.Context, in *ListAllocationRequest, opts ...grpc.CallOption) (*ListAllocationResponse, error) {
	out := new(ListAllocationResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/ListAllocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) RequestSpace(ctx context.Context, in *RequestSpaceRequest, opts ...grpc.CallOption) (*RequestSpaceResponse, error) {
	out := new(RequestSpaceResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/RequestSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) ReleaseSpace(ctx context.Context, in *ReleaseSpaceRequest, opts ...grpc.CallOption) (*ReleaseSpaceResponse, error) {
	out := new(ReleaseSpaceResponse)
	err := c.cc.Invoke(ctx, "/api.storage.v1.StorageService/ReleaseSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
// All implementations must embed UnimplementedStorageServiceServer
// for forward compatibility
type StorageServiceServer interface {
	CreateStorage(context.Context, *CreateStorageRequest) (*CreateStorageResponse, error)
	DeleteStorage(context.Context, *DeleteStorageRequest) (*DeleteStorageResponse, error)
	UpdateStorage(context.Context, *UpdateStorageSizeRequest) (*UpdateStorageSizeResponse, error)
	FindStorage(context.Context, *FindStorageRequest) (*FindStorageResponse, error)
	GetStorage(context.Context, *GetStorageRequest) (*GetStorageResponse, error)
	GetStorageState(context.Context, *GetStorageStateRequest) (*GetStorageStateResponse, error)
	GetStorageMedium(context.Context, *GetStorageMediumRequest) (*GetStorageMediumResponse, error)
	UpdateStorageState(context.Context, *UpdateStorageStateRequest) (*UpdateStorageStateResponse, error)
	CreateAllocation(context.Context, *CreateAllocationRequest) (*CreateAllocationResponse, error)
	DeleteAllocation(context.Context, *DeleteAllocationRequest) (*DeleteAllocationResponse, error)
	UpdateAllocation(context.Context, *UpdateAllocationRequest) (*UpdateAllocationResponse, error)
	FindAllocation(context.Context, *FindAllocationRequest) (*FindAllocationResponse, error)
	GetAllocation(context.Context, *GetAllocationRequest) (*GetAllocationResponse, error)
	ListAllocation(context.Context, *ListAllocationRequest) (*ListAllocationResponse, error)
	RequestSpace(context.Context, *RequestSpaceRequest) (*RequestSpaceResponse, error)
	ReleaseSpace(context.Context, *ReleaseSpaceRequest) (*ReleaseSpaceResponse, error)
	mustEmbedUnimplementedStorageServiceServer()
}

// UnimplementedStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServiceServer struct {
}

func (UnimplementedStorageServiceServer) CreateStorage(context.Context, *CreateStorageRequest) (*CreateStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStorage not implemented")
}
func (UnimplementedStorageServiceServer) DeleteStorage(context.Context, *DeleteStorageRequest) (*DeleteStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStorage not implemented")
}
func (UnimplementedStorageServiceServer) UpdateStorage(context.Context, *UpdateStorageSizeRequest) (*UpdateStorageSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStorage not implemented")
}
func (UnimplementedStorageServiceServer) FindStorage(context.Context, *FindStorageRequest) (*FindStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStorage not implemented")
}
func (UnimplementedStorageServiceServer) GetStorage(context.Context, *GetStorageRequest) (*GetStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorage not implemented")
}
func (UnimplementedStorageServiceServer) GetStorageState(context.Context, *GetStorageStateRequest) (*GetStorageStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorageState not implemented")
}
func (UnimplementedStorageServiceServer) GetStorageMedium(context.Context, *GetStorageMediumRequest) (*GetStorageMediumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorageMedium not implemented")
}
func (UnimplementedStorageServiceServer) UpdateStorageState(context.Context, *UpdateStorageStateRequest) (*UpdateStorageStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStorageState not implemented")
}
func (UnimplementedStorageServiceServer) CreateAllocation(context.Context, *CreateAllocationRequest) (*CreateAllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAllocation not implemented")
}
func (UnimplementedStorageServiceServer) DeleteAllocation(context.Context, *DeleteAllocationRequest) (*DeleteAllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllocation not implemented")
}
func (UnimplementedStorageServiceServer) UpdateAllocation(context.Context, *UpdateAllocationRequest) (*UpdateAllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAllocation not implemented")
}
func (UnimplementedStorageServiceServer) FindAllocation(context.Context, *FindAllocationRequest) (*FindAllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllocation not implemented")
}
func (UnimplementedStorageServiceServer) GetAllocation(context.Context, *GetAllocationRequest) (*GetAllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllocation not implemented")
}
func (UnimplementedStorageServiceServer) ListAllocation(context.Context, *ListAllocationRequest) (*ListAllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllocation not implemented")
}
func (UnimplementedStorageServiceServer) RequestSpace(context.Context, *RequestSpaceRequest) (*RequestSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestSpace not implemented")
}
func (UnimplementedStorageServiceServer) ReleaseSpace(context.Context, *ReleaseSpaceRequest) (*ReleaseSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseSpace not implemented")
}
func (UnimplementedStorageServiceServer) mustEmbedUnimplementedStorageServiceServer() {}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_CreateStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).CreateStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/CreateStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).CreateStorage(ctx, req.(*CreateStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_DeleteStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).DeleteStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/DeleteStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).DeleteStorage(ctx, req.(*DeleteStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_UpdateStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStorageSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).UpdateStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/UpdateStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).UpdateStorage(ctx, req.(*UpdateStorageSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_FindStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).FindStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/FindStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).FindStorage(ctx, req.(*FindStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/GetStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetStorage(ctx, req.(*GetStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetStorageState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetStorageState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/GetStorageState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetStorageState(ctx, req.(*GetStorageStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetStorageMedium_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageMediumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetStorageMedium(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/GetStorageMedium",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetStorageMedium(ctx, req.(*GetStorageMediumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_UpdateStorageState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStorageStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).UpdateStorageState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/UpdateStorageState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).UpdateStorageState(ctx, req.(*UpdateStorageStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_CreateAllocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).CreateAllocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/CreateAllocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).CreateAllocation(ctx, req.(*CreateAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_DeleteAllocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).DeleteAllocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/DeleteAllocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).DeleteAllocation(ctx, req.(*DeleteAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_UpdateAllocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).UpdateAllocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/UpdateAllocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).UpdateAllocation(ctx, req.(*UpdateAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_FindAllocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).FindAllocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/FindAllocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).FindAllocation(ctx, req.(*FindAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetAllocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetAllocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/GetAllocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetAllocation(ctx, req.(*GetAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_ListAllocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).ListAllocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/ListAllocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).ListAllocation(ctx, req.(*ListAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_RequestSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).RequestSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/RequestSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).RequestSpace(ctx, req.(*RequestSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_ReleaseSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).ReleaseSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.storage.v1.StorageService/ReleaseSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).ReleaseSpace(ctx, req.(*ReleaseSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.storage.v1.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStorage",
			Handler:    _StorageService_CreateStorage_Handler,
		},
		{
			MethodName: "DeleteStorage",
			Handler:    _StorageService_DeleteStorage_Handler,
		},
		{
			MethodName: "UpdateStorage",
			Handler:    _StorageService_UpdateStorage_Handler,
		},
		{
			MethodName: "FindStorage",
			Handler:    _StorageService_FindStorage_Handler,
		},
		{
			MethodName: "GetStorage",
			Handler:    _StorageService_GetStorage_Handler,
		},
		{
			MethodName: "GetStorageState",
			Handler:    _StorageService_GetStorageState_Handler,
		},
		{
			MethodName: "GetStorageMedium",
			Handler:    _StorageService_GetStorageMedium_Handler,
		},
		{
			MethodName: "UpdateStorageState",
			Handler:    _StorageService_UpdateStorageState_Handler,
		},
		{
			MethodName: "CreateAllocation",
			Handler:    _StorageService_CreateAllocation_Handler,
		},
		{
			MethodName: "DeleteAllocation",
			Handler:    _StorageService_DeleteAllocation_Handler,
		},
		{
			MethodName: "UpdateAllocation",
			Handler:    _StorageService_UpdateAllocation_Handler,
		},
		{
			MethodName: "FindAllocation",
			Handler:    _StorageService_FindAllocation_Handler,
		},
		{
			MethodName: "GetAllocation",
			Handler:    _StorageService_GetAllocation_Handler,
		},
		{
			MethodName: "ListAllocation",
			Handler:    _StorageService_ListAllocation_Handler,
		},
		{
			MethodName: "RequestSpace",
			Handler:    _StorageService_RequestSpace_Handler,
		},
		{
			MethodName: "ReleaseSpace",
			Handler:    _StorageService_ReleaseSpace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/storage/v1/storage.proto",
}
