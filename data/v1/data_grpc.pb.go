// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: api/data/v1/data.proto

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

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	AddData(ctx context.Context, in *AddDataRequest, opts ...grpc.CallOption) (*AddDataResponse, error)
	DeleteData(ctx context.Context, in *DeleteDataRequest, opts ...grpc.CallOption) (*DeleteDataResponse, error)
	UpdateData(ctx context.Context, in *UpdateDataRequest, opts ...grpc.CallOption) (*UpdateDataResponse, error)
	FindData(ctx context.Context, in *FindDataRequest, opts ...grpc.CallOption) (*FindDataResponse, error)
	BatchDeleteData(ctx context.Context, in *BatchDeleteDataRequest, opts ...grpc.CallOption) (*BatchDeleteDataResponse, error)
	GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) AddData(ctx context.Context, in *AddDataRequest, opts ...grpc.CallOption) (*AddDataResponse, error) {
	out := new(AddDataResponse)
	err := c.cc.Invoke(ctx, "/data.DataService/AddData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) DeleteData(ctx context.Context, in *DeleteDataRequest, opts ...grpc.CallOption) (*DeleteDataResponse, error) {
	out := new(DeleteDataResponse)
	err := c.cc.Invoke(ctx, "/data.DataService/DeleteData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) UpdateData(ctx context.Context, in *UpdateDataRequest, opts ...grpc.CallOption) (*UpdateDataResponse, error) {
	out := new(UpdateDataResponse)
	err := c.cc.Invoke(ctx, "/data.DataService/UpdateData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) FindData(ctx context.Context, in *FindDataRequest, opts ...grpc.CallOption) (*FindDataResponse, error) {
	out := new(FindDataResponse)
	err := c.cc.Invoke(ctx, "/data.DataService/FindData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) BatchDeleteData(ctx context.Context, in *BatchDeleteDataRequest, opts ...grpc.CallOption) (*BatchDeleteDataResponse, error) {
	out := new(BatchDeleteDataResponse)
	err := c.cc.Invoke(ctx, "/data.DataService/BatchDeleteData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error) {
	out := new(GetDataResponse)
	err := c.cc.Invoke(ctx, "/data.DataService/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	AddData(context.Context, *AddDataRequest) (*AddDataResponse, error)
	DeleteData(context.Context, *DeleteDataRequest) (*DeleteDataResponse, error)
	UpdateData(context.Context, *UpdateDataRequest) (*UpdateDataResponse, error)
	FindData(context.Context, *FindDataRequest) (*FindDataResponse, error)
	BatchDeleteData(context.Context, *BatchDeleteDataRequest) (*BatchDeleteDataResponse, error)
	GetData(context.Context, *GetDataRequest) (*GetDataResponse, error)
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) AddData(context.Context, *AddDataRequest) (*AddDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddData not implemented")
}
func (UnimplementedDataServiceServer) DeleteData(context.Context, *DeleteDataRequest) (*DeleteDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteData not implemented")
}
func (UnimplementedDataServiceServer) UpdateData(context.Context, *UpdateDataRequest) (*UpdateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateData not implemented")
}
func (UnimplementedDataServiceServer) FindData(context.Context, *FindDataRequest) (*FindDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindData not implemented")
}
func (UnimplementedDataServiceServer) BatchDeleteData(context.Context, *BatchDeleteDataRequest) (*BatchDeleteDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteData not implemented")
}
func (UnimplementedDataServiceServer) GetData(context.Context, *GetDataRequest) (*GetDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_AddData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).AddData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.DataService/AddData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).AddData(ctx, req.(*AddDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_DeleteData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).DeleteData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.DataService/DeleteData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).DeleteData(ctx, req.(*DeleteDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_UpdateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).UpdateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.DataService/UpdateData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).UpdateData(ctx, req.(*UpdateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_FindData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).FindData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.DataService/FindData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).FindData(ctx, req.(*FindDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_BatchDeleteData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDeleteDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).BatchDeleteData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.DataService/BatchDeleteData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).BatchDeleteData(ctx, req.(*BatchDeleteDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.DataService/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).GetData(ctx, req.(*GetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "data.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddData",
			Handler:    _DataService_AddData_Handler,
		},
		{
			MethodName: "DeleteData",
			Handler:    _DataService_DeleteData_Handler,
		},
		{
			MethodName: "UpdateData",
			Handler:    _DataService_UpdateData_Handler,
		},
		{
			MethodName: "FindData",
			Handler:    _DataService_FindData_Handler,
		},
		{
			MethodName: "BatchDeleteData",
			Handler:    _DataService_BatchDeleteData_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _DataService_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/data/v1/data.proto",
}
