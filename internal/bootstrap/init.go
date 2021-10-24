package bootstrap

import (
	euvies2 "github.com/wishperera/GVAT/internal/adaptors/euvies"
	"github.com/wishperera/GVAT/internal/application"
	"github.com/wishperera/GVAT/internal/container"
	log2 "github.com/wishperera/GVAT/internal/pkg/log"
	"github.com/wishperera/GVAT/internal/services"
	handlers2 "github.com/wishperera/GVAT/internal/transport/http/handlers"
	middleware2 "github.com/wishperera/GVAT/internal/transport/http/middleware"
	router2 "github.com/wishperera/GVAT/internal/transport/http/router"
)

func bindAndInit(c container.AppContainer) {
	c.SetModuleConfig(application.ModuleLogger, new(log2.Config))
	c.SetModuleConfig(application.ModuleRouter, new(router2.Config))
	c.SetModuleConfig(application.ModuleEUVIESAdaptor, new(euvies2.Config))

	c.Bind(application.ModuleLogger, new(log2.Log))
	c.Bind(application.ModuleEUVIESAdaptor, new(euvies2.Adaptor))
	c.Bind(application.ModuleVATIDValidationService, new(services.ValidateVAT))
	c.Bind(application.ModuleContextExtractionMiddleware, new(middleware2.ContextModifier))
	c.Bind(application.ModuleLoggingMiddleWare, new(middleware2.LoggingMiddleware))
	c.Bind(application.ModuleExceptionHandler, new(handlers2.Exception))
	c.Bind(application.ModuleVatIDValidationHandler, new(handlers2.VATIDValidation))
	c.Bind(application.ModuleRouter, new(router2.Router))

	c.Init(
		application.ModuleLogger,
		application.ModuleEUVIESAdaptor,
		application.ModuleVATIDValidationService,
		application.ModuleExceptionHandler,
		application.ModuleVatIDValidationHandler,
		application.ModuleContextExtractionMiddleware,
		application.ModuleLoggingMiddleWare,
		application.ModuleRouter,
	)
}
