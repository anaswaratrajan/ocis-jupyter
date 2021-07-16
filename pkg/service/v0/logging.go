package svc

import (
	"context"
	"time"

	v0proto "github.com/anaswaratrajan/ocis-jupyter/pkg/proto/v0"
	"github.com/owncloud/ocis/ocis-pkg/log"
)

// NewLogging returns a service that logs messages.
func NewLogging(next v0proto.JupyterNotebookSupportHandler, logger log.Logger) v0proto.JupyterNotebookSupportHandler {
	return logging{
		next:   next,
		logger: logger,
	}
}

type logging struct {
	next   v0proto.JupyterNotebookSupportHandler
	logger log.Logger
}

// Greet implements the HelloHandler interface.
func (l logging) GenerateHTML(ctx context.Context, req *v0proto.JupyterNotebookJSON, rsp *v0proto.JupyterNotebookHTML) error {
	start := time.Now()
	err := l.next.GenerateHTML(ctx, req, rsp)

	logger := l.logger.With().
		Str("method", "JupyterNotebookSupport.GenerateHTML").
		Dur("duration", time.Since(start)).
		Logger()

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("Failed to execute")
	} else {
		logger.Debug().
			Msg("")
	}

	return err
}
