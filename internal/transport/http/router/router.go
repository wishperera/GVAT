package router

import (
	"context"
	"github.com/wishperera/GVAT/internal/application"
	"github.com/wishperera/GVAT/internal/container"
	"github.com/wishperera/GVAT/internal/pkg/log"
	"github.com/wishperera/GVAT/internal/transport/http/middleware"
	"net/http"
	"time"
)

const (
	routerShutDownTimeoutSeconds = 5
)

type Router struct {
	log    log.Logger
	mux    *http.ServeMux
	server *http.Server
	closed bool
	ready  chan struct{}
}

func (r *Router) Init(c container.Container) error {
	r.log = c.Resolve(application.ModuleLogger).(log.Logger).NewLog("router")
	contextMiddleWare := c.Resolve(application.ModuleContextExtractionMiddleware).(middleware.Middleware)
	loggingMiddleware := c.Resolve(application.ModuleLoggingMiddleWare).(middleware.Middleware)
	vatValidationHandler := c.Resolve(application.ModuleVatIDValidationHandler).(http.Handler)

	r.mux = http.NewServeMux()
	r.mux.Handle("/validate/",
		contextMiddleWare.Handle(loggingMiddleware.Handle(vatValidationHandler)))

	serverConfig := c.GetModuleConfig(application.ModuleRouter).(*Config)
	r.server = &http.Server{
		Addr:         "0.0.0.0:" + serverConfig.Port,
		Handler:      r.mux,
		ReadTimeout:  time.Second * time.Duration(serverConfig.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(serverConfig.WriteTimeout),
	}
	r.ready = make(chan struct{}, 1)

	return nil
}

func (r *Router) Stop() error {
	r.closed = true
	close(r.ready)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*routerShutDownTimeoutSeconds)
	defer cancel()
	err := r.server.Shutdown(ctx)
	if err != nil {
		r.log.ErrorContext(ctx, "graceful shutdown failed", r.log.Param("err", err))
		return err
	}

	return nil
}

func (r *Router) Run() error {
	go func() {
		err := r.server.ListenAndServe()
		if err != nil && !r.closed {
			r.log.Fatal("failed to initialize server", r.log.Param("err", err))
		}
	}()
	r.ready <- struct{}{}
	return nil
}

func (r *Router) Ready() chan struct{} {
	return r.ready
}
