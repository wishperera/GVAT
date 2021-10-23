package middleware

import (
	"context"
	"fmt"
	"github.com/wishperera/GVAT/gen/internal/application"
	"github.com/wishperera/GVAT/gen/internal/container"
	"github.com/wishperera/GVAT/gen/internal/domain"
	"github.com/wishperera/GVAT/gen/internal/pkg/uuid"
	"github.com/wishperera/GVAT/gen/internal/transport/http/handlers"
	"net/http"
)

type ContextModifier struct {
	next        http.Handler
	errorHandle handlers.ExceptionHandler
}

func (c *ContextModifier) Init(di container.Container) error {
	c.next = di.Resolve(application.ModuleLoggingMiddleWare).(http.Handler)
	c.errorHandle = di.Resolve(application.ModuleExceptionHandler).(handlers.ExceptionHandler)

	return nil
}

func (c *ContextModifier) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var traceID uuid.UUID
	var err error

	tid := request.Header.Get(domain.ContextKeyTraceID.String())
	if tid != "" {
		traceID, err = uuid.Parse(tid)
		if err != nil {
			c.errorHandle.HandleException(request.Context(), writer, InvalidHeader{
				fmt.Errorf("invalid header: %s, must be a valid uuid", domain.ContextKeyTraceID),
			})
			return
		}
	} else {
		traceID = uuid.New()
	}

	ctx := request.Context()
	ctx = context.WithValue(ctx, domain.ContextKeyTraceID, traceID)
	request = request.Clone(ctx)

	c.next.ServeHTTP(writer, request)
}
