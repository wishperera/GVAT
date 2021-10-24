package handlers

import (
	"context"
	"errors"
	"github.com/wishperera/GVAT/internal/adaptors/euvies"
	"github.com/wishperera/GVAT/internal/application"
	"github.com/wishperera/GVAT/internal/container"
	cont "github.com/wishperera/GVAT/internal/pkg/context"
	"github.com/wishperera/GVAT/internal/pkg/log"
	"github.com/wishperera/GVAT/internal/services"
	"github.com/wishperera/GVAT/internal/transport/http/middleware"
	"github.com/wishperera/GVAT/internal/transport/http/response"
	"net/http"
)

type Exception struct {
	log log.Logger
}

func (e *Exception) Init(c container.Container) error {
	e.log = c.Resolve(application.ModuleLogger).(log.Logger).NewLog("handlers.exception")
	return nil
}

func (e *Exception) HandleException(ctx context.Context, w http.ResponseWriter, err error) {
	e.log.ErrorContext(ctx, "request failed", e.log.Param("err", err))

	exception, code := composeExceptionResponse(ctx, err)
	err = response.WriteExceptionResponse(w, exception, nil, code)
	if err != nil {
		e.log.ErrorContext(ctx, "failed to write exception response", e.log.Param("err", err))
		return
	}
}

type responseFields struct {
	code           int
	httpStatusCode int
	trace          string
}

func mapError(err error) responseFields {
	switch err.(type) {
	case MethodNotAllowed:
		return responseFields{
			httpStatusCode: http.StatusMethodNotAllowed,
		}
	case services.ValidationError, euvies.ValidationError:
		return responseFields{
			code:           errorCodeInvalidRequest,
			httpStatusCode: http.StatusBadRequest,
			trace:          err.Error(),
		}
	case services.DependencyError, euvies.DependencyError:
		return responseFields{
			code:           errorCodeDependencyFailure,
			httpStatusCode: http.StatusFailedDependency,
			trace:          err.Error(),
		}
	case middleware.InvalidHeader:
		return responseFields{
			code:           errorCodeInvalidHeader,
			httpStatusCode: http.StatusBadRequest,
			trace:          err.Error(),
		}
	default:
		if errors.Unwrap(err) == nil {
			return responseFields{
				code:           errorCodeUnknown,
				httpStatusCode: http.StatusInternalServerError,
				trace:          "oops..something went wrong",
			}
		}
		return mapError(errors.Unwrap(err))
	}
}

func composeExceptionResponse(ctx context.Context, err error) (
	exception response.Exception, httpCode int) {
	fields := mapError(err)
	resp := response.Exception{
		Code:    fields.code,
		TraceID: cont.ExtractTrace(ctx),
	}

	if fields.code == errorCodeUnknown {
		resp.Description = "something went wrong"
	} else {
		resp.Description = err.Error()
	}

	return resp, fields.httpStatusCode
}
