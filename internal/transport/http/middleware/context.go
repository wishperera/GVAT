package middleware

import (
	"context"
	"fmt"
	"github.com/wishperera/GVAT/internal/application"
	"github.com/wishperera/GVAT/internal/container"
	"github.com/wishperera/GVAT/internal/domain"
	"github.com/wishperera/GVAT/internal/pkg/uuid"
	"github.com/wishperera/GVAT/internal/transport/http/handlers/exception"
	"net/http"
)

type ContextModifier struct {
	errorHandle exception.Handler
}

func (c *ContextModifier) Init(di container.Container) error {
	c.errorHandle = di.Resolve(application.ModuleExceptionHandler).(exception.Handler)
	return nil
}

func (c *ContextModifier) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
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

		next.ServeHTTP(writer, request)
	})
}
