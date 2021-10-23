package bootstrap

import (
	"github.com/wishperera/GVAT/gen/internal/adaptors/euvies"
	"github.com/wishperera/GVAT/gen/internal/application"
	"github.com/wishperera/GVAT/gen/internal/container"
	"github.com/wishperera/GVAT/gen/internal/pkg/log"
	"github.com/wishperera/GVAT/gen/internal/services"
	"github.com/wishperera/GVAT/gen/internal/transport/http/handlers"
	"github.com/wishperera/GVAT/gen/internal/transport/http/middleware"
	"github.com/wishperera/GVAT/gen/internal/transport/http/router"
)

func bindAndInit(c container.AppContainer) {
	c.SetModuleConfig(application.ModuleLogger, new(log.Config))
	c.SetModuleConfig(application.ModuleRouter, new(router.Config))
	c.SetModuleConfig(application.ModuleEUVIESAdaptor, new(euvies.Config))

	c.Bind(application.ModuleLogger, new(log.Log))
	c.Bind(application.ModuleEUVIESAdaptor, new(euvies.Adaptor))
	c.Bind(application.ModuleVATIDValidationService, new(services.ValidateVAT))
	c.Bind(application.ModuleContextExtractionMiddleware, new(middleware.ContextModifier))
	c.Bind(application.ModuleLoggingMiddleWare, new(middleware.LoggingMiddleware))
	c.Bind(application.ModuleExceptionHandler, new(handlers.Exception))
	c.Bind(application.ModuleVatIDValidationHandler, new(handlers.VATIDValidation))
	c.Bind(application.ModuleRouter, new(router.Router))

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
