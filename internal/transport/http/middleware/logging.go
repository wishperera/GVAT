package middleware

import (
	"github.com/wishperera/GVAT/internal/application"
	"github.com/wishperera/GVAT/internal/container"
	"github.com/wishperera/GVAT/internal/pkg/log"
	"net/http"
	"net/http/httputil"
)

type LoggingMiddleware struct {
	log log.Logger
}

func (l *LoggingMiddleware) Init(c container.Container) error {
	l.log = c.Resolve(application.ModuleLogger).(log.Logger)
	return nil
}

func (l *LoggingMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			l.log.ErrorContext(ctx, "failed to dump request due", l.log.Param("err", err))
		} else {
			l.log.DebugContext(ctx, "incoming request", l.log.Param("record", string(dump)))
		}

		next.ServeHTTP(writer, request)
	})
}
