// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPresetCreatePreset = "/api.preset.v1.Preset/CreatePreset"
const OperationPresetDeletePreset = "/api.preset.v1.Preset/DeletePreset"
const OperationPresetGetPreset = "/api.preset.v1.Preset/GetPreset"
const OperationPresetGetPresetByAlgoId = "/api.preset.v1.Preset/GetPresetByAlgoId"

type PresetHTTPServer interface {
	CreatePreset(context.Context, *CreatePresetRequest) (*CreatePresetReply, error)
	DeletePreset(context.Context, *DeletePresetRequest) (*DeletePresetReply, error)
	GetPreset(context.Context, *GetPresetRequest) (*GetPresetReply, error)
	GetPresetByAlgoId(context.Context, *GetPresetByAlgoIdRequest) (*GetPresetByAlgoIdReply, error)
}

func RegisterPresetHTTPServer(s *http.Server, srv PresetHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/preset", _Preset_CreatePreset0_HTTP_Handler(srv))
	r.DELETE("/api/v1/preset", _Preset_DeletePreset0_HTTP_Handler(srv))
	r.GET("/api/v1/preset/{id}", _Preset_GetPreset0_HTTP_Handler(srv))
	r.GET("/api/v1/preset/algo/{id}", _Preset_GetPresetByAlgoId0_HTTP_Handler(srv))
}

func _Preset_CreatePreset0_HTTP_Handler(srv PresetHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePresetRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPresetCreatePreset)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePreset(ctx, req.(*CreatePresetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePresetReply)
		return ctx.Result(200, reply)
	}
}

func _Preset_DeletePreset0_HTTP_Handler(srv PresetHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePresetRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPresetDeletePreset)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePreset(ctx, req.(*DeletePresetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeletePresetReply)
		return ctx.Result(200, reply)
	}
}

func _Preset_GetPreset0_HTTP_Handler(srv PresetHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPresetRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPresetGetPreset)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPreset(ctx, req.(*GetPresetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPresetReply)
		return ctx.Result(200, reply)
	}
}

func _Preset_GetPresetByAlgoId0_HTTP_Handler(srv PresetHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPresetByAlgoIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPresetGetPresetByAlgoId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPresetByAlgoId(ctx, req.(*GetPresetByAlgoIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPresetByAlgoIdReply)
		return ctx.Result(200, reply)
	}
}

type PresetHTTPClient interface {
	CreatePreset(ctx context.Context, req *CreatePresetRequest, opts ...http.CallOption) (rsp *CreatePresetReply, err error)
	DeletePreset(ctx context.Context, req *DeletePresetRequest, opts ...http.CallOption) (rsp *DeletePresetReply, err error)
	GetPreset(ctx context.Context, req *GetPresetRequest, opts ...http.CallOption) (rsp *GetPresetReply, err error)
	GetPresetByAlgoId(ctx context.Context, req *GetPresetByAlgoIdRequest, opts ...http.CallOption) (rsp *GetPresetByAlgoIdReply, err error)
}

type PresetHTTPClientImpl struct {
	cc *http.Client
}

func NewPresetHTTPClient(client *http.Client) PresetHTTPClient {
	return &PresetHTTPClientImpl{client}
}

func (c *PresetHTTPClientImpl) CreatePreset(ctx context.Context, in *CreatePresetRequest, opts ...http.CallOption) (*CreatePresetReply, error) {
	var out CreatePresetReply
	pattern := "/api/v1/preset"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPresetCreatePreset))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PresetHTTPClientImpl) DeletePreset(ctx context.Context, in *DeletePresetRequest, opts ...http.CallOption) (*DeletePresetReply, error) {
	var out DeletePresetReply
	pattern := "/api/v1/preset"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPresetDeletePreset))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PresetHTTPClientImpl) GetPreset(ctx context.Context, in *GetPresetRequest, opts ...http.CallOption) (*GetPresetReply, error) {
	var out GetPresetReply
	pattern := "/api/v1/preset/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPresetGetPreset))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PresetHTTPClientImpl) GetPresetByAlgoId(ctx context.Context, in *GetPresetByAlgoIdRequest, opts ...http.CallOption) (*GetPresetByAlgoIdReply, error) {
	var out GetPresetByAlgoIdReply
	pattern := "/api/v1/preset/algo/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPresetGetPresetByAlgoId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
