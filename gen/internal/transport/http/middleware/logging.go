package middleware

import (
	"github.com/wishperera/GVAT/gen/internal/application"
	"github.com/wishperera/GVAT/gen/internal/container"
	"github.com/wishperera/GVAT/gen/internal/pkg/log"
	"net/http"
	"net/http/httputil"
)

type LoggingMiddleware struct {
	log  log.Logger
	next http.Handler
}

func (l *LoggingMiddleware) Init(c container.Container) error {
	l.log = c.Resolve(application.ModuleLogger).(log.Logger)
	l.next = c.Resolve(application.ModuleVatIDValidationHandler).(http.Handler)

	return nil
}

func (l *LoggingMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		l.log.ErrorContext(ctx, "failed to dump request due", l.log.Param("err", err))
	} else {
		l.log.DebugContext(ctx, "incoming request", l.log.Param("record", string(dump)))
	}

	l.next.ServeHTTP(writer, request)
}
