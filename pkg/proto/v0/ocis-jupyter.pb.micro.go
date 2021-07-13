// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: ocis-jupyter.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for JupyterNotebookSupport service

func NewJupyterNotebookSupportEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "JupyterNotebookSupport.GenerateHTML",
			Path:    []string{"/api/v0/convert"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for JupyterNotebookSupport service

type JupyterNotebookSupportService interface {
	GenerateHTML(ctx context.Context, in *JupyterNotebookJSON, opts ...client.CallOption) (*JupyterNotebookHTML, error)
}

type jupyterNotebookSupportService struct {
	c    client.Client
	name string
}

func NewJupyterNotebookSupportService(name string, c client.Client) JupyterNotebookSupportService {
	return &jupyterNotebookSupportService{
		c:    c,
		name: name,
	}
}

func (c *jupyterNotebookSupportService) GenerateHTML(ctx context.Context, in *JupyterNotebookJSON, opts ...client.CallOption) (*JupyterNotebookHTML, error) {
	req := c.c.NewRequest(c.name, "JupyterNotebookSupport.GenerateHTML", in)
	out := new(JupyterNotebookHTML)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JupyterNotebookSupport service

type JupyterNotebookSupportHandler interface {
	GenerateHTML(context.Context, *JupyterNotebookJSON, *JupyterNotebookHTML) error
}

func RegisterJupyterNotebookSupportHandler(s server.Server, hdlr JupyterNotebookSupportHandler, opts ...server.HandlerOption) error {
	type jupyterNotebookSupport interface {
		GenerateHTML(ctx context.Context, in *JupyterNotebookJSON, out *JupyterNotebookHTML) error
	}
	type JupyterNotebookSupport struct {
		jupyterNotebookSupport
	}
	h := &jupyterNotebookSupportHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "JupyterNotebookSupport.GenerateHTML",
		Path:    []string{"/api/v0/convert"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&JupyterNotebookSupport{h}, opts...))
}

type jupyterNotebookSupportHandler struct {
	JupyterNotebookSupportHandler
}

func (h *jupyterNotebookSupportHandler) GenerateHTML(ctx context.Context, in *JupyterNotebookJSON, out *JupyterNotebookHTML) error {
	return h.JupyterNotebookSupportHandler.GenerateHTML(ctx, in, out)
}