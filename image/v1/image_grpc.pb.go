// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: image/v1/image.proto

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

// ImageClient is the client API for Image service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageClient interface {
	CreateImage(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*CreateImageReply, error)
	UpdateImage(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*UpdateImageReply, error)
	DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*DeleteImageReply, error)
	GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*GetImageReply, error)
	GetImageByNameVersion(ctx context.Context, in *GetImageByNameVersionRequest, opts ...grpc.CallOption) (*GetImageByNameVersionReply, error)
	ListImage(ctx context.Context, in *ListImageRequest, opts ...grpc.CallOption) (*ListImageReply, error)
	ListAvailableImage(ctx context.Context, in *ListAvailableImageRequest, opts ...grpc.CallOption) (*ListAvailableImageReply, error)
	ListAvailableImageByName(ctx context.Context, in *ListAvailableImageByNameRequest, opts ...grpc.CallOption) (*ListAvailableImageByNameReply, error)
	LockImage(ctx context.Context, in *LockImageRequest, opts ...grpc.CallOption) (*LockImageReply, error)
	UnlockImage(ctx context.Context, in *UnlockImageRequest, opts ...grpc.CallOption) (*UnlockImageReply, error)
}

type imageClient struct {
	cc grpc.ClientConnInterface
}

func NewImageClient(cc grpc.ClientConnInterface) ImageClient {
	return &imageClient{cc}
}

func (c *imageClient) CreateImage(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*CreateImageReply, error) {
	out := new(CreateImageReply)
	err := c.cc.Invoke(ctx, "/api.image.v1.Image/CreateImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) UpdateImage(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*UpdateImageReply, error) {
	out := new(UpdateImageReply)
	err := c.cc.Invoke(ctx, "/api.image.v1.Image/UpdateImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*DeleteImageReply, error) {
	out := new(DeleteImageReply)
	err := c.cc.Invoke(ctx, "/api.image.v1.Image/DeleteImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*GetImageReply, error) {
	out := new(GetImageReply)
	err := c.cc.Invoke(ctx, "/api.image.v1.Image/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) GetImageByNameVersion(ctx context.Context, in *GetImageByNameVersionRequest, opts ...grpc.CallOption) (*GetImageByNameVersionReply, error) {
	out := new(GetImageByNameVersionReply)
	err := c.cc.Invoke(ctx, "/api.image.v1.Image/GetImageByNameVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) ListImage(ctx context.Context, in *ListImageRequest, opts ...grpc.CallOption) (*ListImageReply, error) {
	out := new(ListImageReply)
	err := c.cc.Invoke(ctx, "/api.image.v1.Image/ListImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) ListAvailableImage(ctx context.Context, in *ListAvailableImageRequest, opts ...grpc.CallOption) (*ListAvailableImageReply, error) {
	out := new(ListAvailableImageReply)
	err := c.cc.Invoke(ctx, "/api.image.v1.Image/ListAvailableImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) ListAvailableImageByName(ctx context.Context, in *ListAvailableImageByNameRequest, opts ...grpc.CallOption) (*ListAvailableImageByNameReply, error) {
	out := new(ListAvailableImageByNameReply)
	err := c.cc.Invoke(ctx, "/api.image.v1.Image/ListAvailableImageByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) LockImage(ctx context.Context, in *LockImageRequest, opts ...grpc.CallOption) (*LockImageReply, error) {
	out := new(LockImageReply)
	err := c.cc.Invoke(ctx, "/api.image.v1.Image/LockImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) UnlockImage(ctx context.Context, in *UnlockImageRequest, opts ...grpc.CallOption) (*UnlockImageReply, error) {
	out := new(UnlockImageReply)
	err := c.cc.Invoke(ctx, "/api.image.v1.Image/UnlockImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServer is the server API for Image service.
// All implementations must embed UnimplementedImageServer
// for forward compatibility
type ImageServer interface {
	CreateImage(context.Context, *CreateImageRequest) (*CreateImageReply, error)
	UpdateImage(context.Context, *UpdateImageRequest) (*UpdateImageReply, error)
	DeleteImage(context.Context, *DeleteImageRequest) (*DeleteImageReply, error)
	GetImage(context.Context, *GetImageRequest) (*GetImageReply, error)
	GetImageByNameVersion(context.Context, *GetImageByNameVersionRequest) (*GetImageByNameVersionReply, error)
	ListImage(context.Context, *ListImageRequest) (*ListImageReply, error)
	ListAvailableImage(context.Context, *ListAvailableImageRequest) (*ListAvailableImageReply, error)
	ListAvailableImageByName(context.Context, *ListAvailableImageByNameRequest) (*ListAvailableImageByNameReply, error)
	LockImage(context.Context, *LockImageRequest) (*LockImageReply, error)
	UnlockImage(context.Context, *UnlockImageRequest) (*UnlockImageReply, error)
	mustEmbedUnimplementedImageServer()
}

// UnimplementedImageServer must be embedded to have forward compatible implementations.
type UnimplementedImageServer struct {
}

func (UnimplementedImageServer) CreateImage(context.Context, *CreateImageRequest) (*CreateImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImage not implemented")
}
func (UnimplementedImageServer) UpdateImage(context.Context, *UpdateImageRequest) (*UpdateImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateImage not implemented")
}
func (UnimplementedImageServer) DeleteImage(context.Context, *DeleteImageRequest) (*DeleteImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}
func (UnimplementedImageServer) GetImage(context.Context, *GetImageRequest) (*GetImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedImageServer) GetImageByNameVersion(context.Context, *GetImageByNameVersionRequest) (*GetImageByNameVersionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageByNameVersion not implemented")
}
func (UnimplementedImageServer) ListImage(context.Context, *ListImageRequest) (*ListImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImage not implemented")
}
func (UnimplementedImageServer) ListAvailableImage(context.Context, *ListAvailableImageRequest) (*ListAvailableImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAvailableImage not implemented")
}
func (UnimplementedImageServer) ListAvailableImageByName(context.Context, *ListAvailableImageByNameRequest) (*ListAvailableImageByNameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAvailableImageByName not implemented")
}
func (UnimplementedImageServer) LockImage(context.Context, *LockImageRequest) (*LockImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockImage not implemented")
}
func (UnimplementedImageServer) UnlockImage(context.Context, *UnlockImageRequest) (*UnlockImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockImage not implemented")
}
func (UnimplementedImageServer) mustEmbedUnimplementedImageServer() {}

// UnsafeImageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServer will
// result in compilation errors.
type UnsafeImageServer interface {
	mustEmbedUnimplementedImageServer()
}

func RegisterImageServer(s grpc.ServiceRegistrar, srv ImageServer) {
	s.RegisterService(&Image_ServiceDesc, srv)
}

func _Image_CreateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).CreateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.image.v1.Image/CreateImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).CreateImage(ctx, req.(*CreateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_UpdateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).UpdateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.image.v1.Image/UpdateImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).UpdateImage(ctx, req.(*UpdateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.image.v1.Image/DeleteImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).DeleteImage(ctx, req.(*DeleteImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.image.v1.Image/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).GetImage(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_GetImageByNameVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageByNameVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).GetImageByNameVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.image.v1.Image/GetImageByNameVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).GetImageByNameVersion(ctx, req.(*GetImageByNameVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_ListImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).ListImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.image.v1.Image/ListImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).ListImage(ctx, req.(*ListImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_ListAvailableImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAvailableImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).ListAvailableImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.image.v1.Image/ListAvailableImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).ListAvailableImage(ctx, req.(*ListAvailableImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_ListAvailableImageByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAvailableImageByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).ListAvailableImageByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.image.v1.Image/ListAvailableImageByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).ListAvailableImageByName(ctx, req.(*ListAvailableImageByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_LockImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).LockImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.image.v1.Image/LockImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).LockImage(ctx, req.(*LockImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_UnlockImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).UnlockImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.image.v1.Image/UnlockImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).UnlockImage(ctx, req.(*UnlockImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Image_ServiceDesc is the grpc.ServiceDesc for Image service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Image_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.image.v1.Image",
	HandlerType: (*ImageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateImage",
			Handler:    _Image_CreateImage_Handler,
		},
		{
			MethodName: "UpdateImage",
			Handler:    _Image_UpdateImage_Handler,
		},
		{
			MethodName: "DeleteImage",
			Handler:    _Image_DeleteImage_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _Image_GetImage_Handler,
		},
		{
			MethodName: "GetImageByNameVersion",
			Handler:    _Image_GetImageByNameVersion_Handler,
		},
		{
			MethodName: "ListImage",
			Handler:    _Image_ListImage_Handler,
		},
		{
			MethodName: "ListAvailableImage",
			Handler:    _Image_ListAvailableImage_Handler,
		},
		{
			MethodName: "ListAvailableImageByName",
			Handler:    _Image_ListAvailableImageByName_Handler,
		},
		{
			MethodName: "LockImage",
			Handler:    _Image_LockImage_Handler,
		},
		{
			MethodName: "UnlockImage",
			Handler:    _Image_UnlockImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image/v1/image.proto",
}
