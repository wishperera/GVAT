package application

const (
	ModuleLogger = "modules.logger"

	// adaptor module definitions
	ModuleEUVIESAdaptor = "modules.adaptors.eu-vies"

	// service module definitions
	ModuleVATIDValidationService = "modules.services.vat-id"

	// middlewares
	ModuleContextExtractionMiddleware = "modules.middlewares.context-extraction"
	ModuleMetricsMiddleware           = "modules.middlewares.metrics"
	ModuleLoggingMiddleWare           = "modules.middlewares.logging"

	// handlers
	ModuleExceptionHandler       = "modules.handlers.exception"
	ModuleVatIDValidationHandler = "modules.handlers.vat"

	ModuleRouter = "modules.router"
)
