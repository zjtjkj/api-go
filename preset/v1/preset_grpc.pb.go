// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: preset/v1/preset.proto

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

// PresetClient is the client API for Preset service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PresetClient interface {
	CreatePreset(ctx context.Context, in *CreatePresetRequest, opts ...grpc.CallOption) (*CreatePresetReply, error)
	DeletePreset(ctx context.Context, in *DeletePresetRequest, opts ...grpc.CallOption) (*DeletePresetReply, error)
	GetPreset(ctx context.Context, in *GetPresetRequest, opts ...grpc.CallOption) (*GetPresetReply, error)
	GetPresetByAlgoId(ctx context.Context, in *GetPresetByAlgoIdRequest, opts ...grpc.CallOption) (*GetPresetByAlgoIdReply, error)
}

type presetClient struct {
	cc grpc.ClientConnInterface
}

func NewPresetClient(cc grpc.ClientConnInterface) PresetClient {
	return &presetClient{cc}
}

func (c *presetClient) CreatePreset(ctx context.Context, in *CreatePresetRequest, opts ...grpc.CallOption) (*CreatePresetReply, error) {
	out := new(CreatePresetReply)
	err := c.cc.Invoke(ctx, "/api.preset.v1.Preset/CreatePreset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetClient) DeletePreset(ctx context.Context, in *DeletePresetRequest, opts ...grpc.CallOption) (*DeletePresetReply, error) {
	out := new(DeletePresetReply)
	err := c.cc.Invoke(ctx, "/api.preset.v1.Preset/DeletePreset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetClient) GetPreset(ctx context.Context, in *GetPresetRequest, opts ...grpc.CallOption) (*GetPresetReply, error) {
	out := new(GetPresetReply)
	err := c.cc.Invoke(ctx, "/api.preset.v1.Preset/GetPreset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetClient) GetPresetByAlgoId(ctx context.Context, in *GetPresetByAlgoIdRequest, opts ...grpc.CallOption) (*GetPresetByAlgoIdReply, error) {
	out := new(GetPresetByAlgoIdReply)
	err := c.cc.Invoke(ctx, "/api.preset.v1.Preset/GetPresetByAlgoId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PresetServer is the server API for Preset service.
// All implementations must embed UnimplementedPresetServer
// for forward compatibility
type PresetServer interface {
	CreatePreset(context.Context, *CreatePresetRequest) (*CreatePresetReply, error)
	DeletePreset(context.Context, *DeletePresetRequest) (*DeletePresetReply, error)
	GetPreset(context.Context, *GetPresetRequest) (*GetPresetReply, error)
	GetPresetByAlgoId(context.Context, *GetPresetByAlgoIdRequest) (*GetPresetByAlgoIdReply, error)
	mustEmbedUnimplementedPresetServer()
}

// UnimplementedPresetServer must be embedded to have forward compatible implementations.
type UnimplementedPresetServer struct {
}

func (UnimplementedPresetServer) CreatePreset(context.Context, *CreatePresetRequest) (*CreatePresetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePreset not implemented")
}
func (UnimplementedPresetServer) DeletePreset(context.Context, *DeletePresetRequest) (*DeletePresetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePreset not implemented")
}
func (UnimplementedPresetServer) GetPreset(context.Context, *GetPresetRequest) (*GetPresetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreset not implemented")
}
func (UnimplementedPresetServer) GetPresetByAlgoId(context.Context, *GetPresetByAlgoIdRequest) (*GetPresetByAlgoIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPresetByAlgoId not implemented")
}
func (UnimplementedPresetServer) mustEmbedUnimplementedPresetServer() {}

// UnsafePresetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PresetServer will
// result in compilation errors.
type UnsafePresetServer interface {
	mustEmbedUnimplementedPresetServer()
}

func RegisterPresetServer(s grpc.ServiceRegistrar, srv PresetServer) {
	s.RegisterService(&Preset_ServiceDesc, srv)
}

func _Preset_CreatePreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetServer).CreatePreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.preset.v1.Preset/CreatePreset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetServer).CreatePreset(ctx, req.(*CreatePresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Preset_DeletePreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetServer).DeletePreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.preset.v1.Preset/DeletePreset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetServer).DeletePreset(ctx, req.(*DeletePresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Preset_GetPreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetServer).GetPreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.preset.v1.Preset/GetPreset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetServer).GetPreset(ctx, req.(*GetPresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Preset_GetPresetByAlgoId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPresetByAlgoIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetServer).GetPresetByAlgoId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.preset.v1.Preset/GetPresetByAlgoId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetServer).GetPresetByAlgoId(ctx, req.(*GetPresetByAlgoIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Preset_ServiceDesc is the grpc.ServiceDesc for Preset service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Preset_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.preset.v1.Preset",
	HandlerType: (*PresetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePreset",
			Handler:    _Preset_CreatePreset_Handler,
		},
		{
			MethodName: "DeletePreset",
			Handler:    _Preset_DeletePreset_Handler,
		},
		{
			MethodName: "GetPreset",
			Handler:    _Preset_GetPreset_Handler,
		},
		{
			MethodName: "GetPresetByAlgoId",
			Handler:    _Preset_GetPresetByAlgoId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "preset/v1/preset.proto",
}
