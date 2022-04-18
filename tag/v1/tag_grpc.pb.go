// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: api/tag/v1/tag.proto

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

// TagClient is the client API for Tag service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagClient interface {
	CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*CreateTagReply, error)
	UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*UpdateTagReply, error)
	DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*DeleteTagReply, error)
	GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*GetTagReply, error)
	ListTag(ctx context.Context, in *ListTagRequest, opts ...grpc.CallOption) (*ListTagReply, error)
}

type tagClient struct {
	cc grpc.ClientConnInterface
}

func NewTagClient(cc grpc.ClientConnInterface) TagClient {
	return &tagClient{cc}
}

func (c *tagClient) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*CreateTagReply, error) {
	out := new(CreateTagReply)
	err := c.cc.Invoke(ctx, "/api.tag.v1.Tag/CreateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagClient) UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*UpdateTagReply, error) {
	out := new(UpdateTagReply)
	err := c.cc.Invoke(ctx, "/api.tag.v1.Tag/UpdateTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagClient) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...grpc.CallOption) (*DeleteTagReply, error) {
	out := new(DeleteTagReply)
	err := c.cc.Invoke(ctx, "/api.tag.v1.Tag/DeleteTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagClient) GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*GetTagReply, error) {
	out := new(GetTagReply)
	err := c.cc.Invoke(ctx, "/api.tag.v1.Tag/GetTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagClient) ListTag(ctx context.Context, in *ListTagRequest, opts ...grpc.CallOption) (*ListTagReply, error) {
	out := new(ListTagReply)
	err := c.cc.Invoke(ctx, "/api.tag.v1.Tag/ListTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServer is the server API for Tag service.
// All implementations must embed UnimplementedTagServer
// for forward compatibility
type TagServer interface {
	CreateTag(context.Context, *CreateTagRequest) (*CreateTagReply, error)
	UpdateTag(context.Context, *UpdateTagRequest) (*UpdateTagReply, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*DeleteTagReply, error)
	GetTag(context.Context, *GetTagRequest) (*GetTagReply, error)
	ListTag(context.Context, *ListTagRequest) (*ListTagReply, error)
	mustEmbedUnimplementedTagServer()
}

// UnimplementedTagServer must be embedded to have forward compatible implementations.
type UnimplementedTagServer struct {
}

func (UnimplementedTagServer) CreateTag(context.Context, *CreateTagRequest) (*CreateTagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTag not implemented")
}
func (UnimplementedTagServer) UpdateTag(context.Context, *UpdateTagRequest) (*UpdateTagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedTagServer) DeleteTag(context.Context, *DeleteTagRequest) (*DeleteTagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}
func (UnimplementedTagServer) GetTag(context.Context, *GetTagRequest) (*GetTagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}
func (UnimplementedTagServer) ListTag(context.Context, *ListTagRequest) (*ListTagReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTag not implemented")
}
func (UnimplementedTagServer) mustEmbedUnimplementedTagServer() {}

// UnsafeTagServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServer will
// result in compilation errors.
type UnsafeTagServer interface {
	mustEmbedUnimplementedTagServer()
}

func RegisterTagServer(s grpc.ServiceRegistrar, srv TagServer) {
	s.RegisterService(&Tag_ServiceDesc, srv)
}

func _Tag_CreateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServer).CreateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.tag.v1.Tag/CreateTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServer).CreateTag(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tag_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.tag.v1.Tag/UpdateTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServer).UpdateTag(ctx, req.(*UpdateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tag_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.tag.v1.Tag/DeleteTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServer).DeleteTag(ctx, req.(*DeleteTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tag_GetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServer).GetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.tag.v1.Tag/GetTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServer).GetTag(ctx, req.(*GetTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tag_ListTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServer).ListTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.tag.v1.Tag/ListTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServer).ListTag(ctx, req.(*ListTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tag_ServiceDesc is the grpc.ServiceDesc for Tag service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tag_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.tag.v1.Tag",
	HandlerType: (*TagServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTag",
			Handler:    _Tag_CreateTag_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _Tag_UpdateTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _Tag_DeleteTag_Handler,
		},
		{
			MethodName: "GetTag",
			Handler:    _Tag_GetTag_Handler,
		},
		{
			MethodName: "ListTag",
			Handler:    _Tag_ListTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/tag/v1/tag.proto",
}
