package svc

import (
	"context"

	v0proto "github.com/anaswaratrajan/ocis-jupyter/pkg/proto/v0"
	"go.opencensus.io/trace"
)

// NewTracing returns a service that instruments traces.
func NewTracing(next v0proto.JupyterNotebookSupportHandler) v0proto.JupyterNotebookSupportHandler {
	return tracing{
		next: next,
	}
}

type tracing struct {
	next v0proto.JupyterNotebookSupportHandler
}

// Greet implements the HelloHandler interface.
func (t tracing) GenerateHTML(ctx context.Context, req *v0proto.JupyterNotebookJSON, rsp *v0proto.JupyterNotebookHTML) error {
	ctx, span := trace.StartSpan(ctx, "JupyterNotebookSupport.GenerateHTML")
	defer span.End()

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("name", req.JSONString),
	}, "Execute JupyterNotebookSupport.GenerateHTML handler")

	return t.next.GenerateHTML(ctx, req, rsp)
}
