package http

import (
	"github.com/go-chi/chi"
	"github.com/anaswaratrajan/ocis-jupyter/pkg/assets"
	"github.com/anaswaratrajan/ocis-jupyter/pkg/proto/v0"
	svc "github.com/anaswaratrajan/ocis-jupyter/pkg/service/v0"
	"github.com/anaswaratrajan/ocis-jupyter/pkg/version"
	"github.com/owncloud/ocis/ocis-pkg/account"
	"github.com/owncloud/ocis/ocis-pkg/middleware"
	"github.com/owncloud/ocis/ocis-pkg/service/http"
)

// Server initializes the http service and server.
func Server(opts ...Option) http.Service {
	options := newOptions(opts...)

	service := http.NewService(
		http.Logger(options.Logger),
		http.Name(options.Name),
		http.Version(version.String),
		http.Address(options.Config.HTTP.Addr),
		http.Namespace(options.Config.HTTP.Namespace),
		http.Context(options.Context),
		http.Flags(options.Flags...),
	)

	handle := svc.NewService()

	{
		handle = svc.NewInstrument(handle, options.Metrics)
		handle = svc.NewLogging(handle, options.Logger)
		handle = svc.NewTracing(handle)
	}

	mux := chi.NewMux()

	mux.Use(middleware.RealIP)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.NoCache)
	mux.Use(middleware.Cors)
	mux.Use(middleware.Secure)
	mux.Use(middleware.ExtractAccountUUID(
		account.Logger(options.Logger),
		account.JWTSecret(options.Config.TokenManager.JWTSecret)),
	)

	mux.Use(middleware.Version(
		options.Name,
		version.String,
	))

	mux.Use(middleware.Logger(
		options.Logger,
	))

	mux.Use(middleware.Static(
		options.Config.HTTP.Root,
		assets.New(
			assets.Logger(options.Logger),
			assets.Config(options.Config),
		),
		// Currently this option does not affect anything but might again in the future
		// when the static middleware implements caching again.
		// TTL = 7 days in seconds = 60 * 60 * 24 * 7
		604800))

	mux.Route(options.Config.HTTP.Root, func(r chi.Router) {
		proto.RegisterHelloWeb(r, handle)
	})

	service.Handle(
		"/",
		mux,
	)

	if err := service.Init(); err != nil {
		panic(err)
	}
	return service
}
