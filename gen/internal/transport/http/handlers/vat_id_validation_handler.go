package handlers

import (
	"fmt"
	"github.com/wishperera/GVAT/gen/internal/application"
	"github.com/wishperera/GVAT/gen/internal/container"
	"github.com/wishperera/GVAT/gen/internal/domain/services"
	"github.com/wishperera/GVAT/gen/internal/pkg/log"
	"github.com/wishperera/GVAT/gen/internal/transport/http/response"
	"net/http"
	"strings"
)

type VATIDValidation struct {
	log      log.Logger
	services struct {
		vatValidation services.ValidateVATID
	}
	errorHandle ExceptionHandler
}

func (v *VATIDValidation) Init(c container.Container) error {
	v.log = c.Resolve(application.ModuleLogger).(log.Logger).NewLog("handlers.vat-id")
	v.services.vatValidation = c.Resolve(application.ModuleVATIDValidationService).(services.ValidateVATID)
	v.errorHandle = c.Resolve(application.ModuleExceptionHandler).(ExceptionHandler)

	return nil
}

func (v *VATIDValidation) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := strings.TrimPrefix(request.URL.Path, HandlePathValidateVATID)
	method := request.Method
	ctx := request.Context()

	if method != http.MethodGet {
		v.errorHandle.HandleException(ctx, writer, MethodNotAllowed{fmt.Errorf("method: %s not allowed", request.Method)})
	}

	err := v.services.vatValidation.Validate(ctx, id)
	if err != nil {
		v.errorHandle.HandleException(ctx, writer, err)
		return
	}

	err = response.WriteSuccessResponse(writer, nil, nil, http.StatusNoContent)
	if err != nil {
		v.errorHandle.HandleException(ctx, writer, err)
		return
	}
}
