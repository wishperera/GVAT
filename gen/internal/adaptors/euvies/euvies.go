package euvies

import (
	"context"
	"github.com/wishperera/GVAT/gen/internal/application"
	"github.com/wishperera/GVAT/gen/internal/container"
	"github.com/wishperera/GVAT/gen/internal/pkg/log"
	"net/http"
)

type Adaptor struct {
	log        log.Logger
	client     *http.Client
	maxRetries int
	baseURL    string
}

func (e *Adaptor) Init(c container.Container) error {
	e.log = c.Resolve(application.ModuleLogger).(log.Logger).NewLog("adaptors.euvies")
	config := c.GetModuleConfig(application.ModuleEUVIESAdaptor).(*Config)
	e.client = &http.Client{
		Timeout: config.Timeout,
	}
	e.baseURL = config.URL
	e.maxRetries = config.MaxRetries

	return nil
}

func (e *Adaptor) ValidateVATID(ctx context.Context, vatID string) (valid bool, err error) {
	// todo - implement me
	return true, nil
}
