package euvies

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"github.com/wishperera/GVAT/internal/application"
	"github.com/wishperera/GVAT/internal/container"
	"github.com/wishperera/GVAT/internal/pkg/log"
	"github.com/wishperera/GVAT/internal/workerpool"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"
)

const (
	vatCheckEndpoint = "/checkVatService"

	failedToComposeRequestDue = "failed to compose request due: %s"
)

type Adaptor struct {
	log        log.Logger
	client     *http.Client
	maxRetries int
	baseURL    string
	template   *template.Template
	pool       *workerpool.Pool
	wait       chan struct{}
}

func (e *Adaptor) Init(c container.Container) error {
	e.log = c.Resolve(application.ModuleLogger).(log.Logger).NewLog("adaptors.euvies")
	config := c.GetModuleConfig(application.ModuleEUVIESAdaptor).(*Config)
	e.client = &http.Client{
		Timeout: time.Millisecond * time.Duration(config.Timeout),
	}
	e.baseURL = config.URL
	e.maxRetries = config.MaxRetries

	temp, err := template.New("VatRequest").Parse(requestTemplate)
	if err != nil {
		return fmt.Errorf("failed to generate request template due: %s", err)
	}
	e.template = temp
	e.pool, err = workerpool.NewPool(workerpool.Config{
		MaxWorkers:       config.MaxWorkers,
		QueueSize:        config.QueueSize,
		WorkerBufferSize: config.WorkerBuffer,
		Process:          e.performWithRetry,
	}, e.log)
	if err != nil {
		return fmt.Errorf("failed to initialize worker pool due: %s", err)
	}
	e.wait = make(chan struct{}, 1)

	return nil
}

func (e *Adaptor) Run() error {
	e.pool.Init()
	e.wait <- struct{}{}
	return nil
}

func (e *Adaptor) Ready() chan struct{} {
	return e.wait
}

func (e *Adaptor) Stop() error {
	e.pool.ShutDown()
	close(e.wait)
	return nil
}

func (e *Adaptor) ValidateVATID(ctx context.Context, countryCode, vatID string) (valid bool, err error) {
	request, err := e.generateRequest(ctx, countryCode, vatID)
	if err != nil {
		return false, err
	}

	outChan, errChan, err := e.pool.ExecuteJob(ctx, request)
	if err != nil {
		return false, DependencyError{
			fmt.Errorf("failed to process request due: %s", err),
		}
	}

	var out interface{}

	select {
	case out = <-outChan:
		response := out.(ResponseBody)
		return response.Valid, nil
	case err = <-errChan:
		return false, err
	}
}

// performWithRetry: performs the request with retry
// todo - use exponential backoff
func (e *Adaptor) performWithRetry(ctx context.Context, input interface{}) (response interface{}, err error) {
	in := input.(*http.Request)
	var res *http.Response
	for i := 0; i < e.maxRetries+1; i++ {
		res, err = e.client.Do(in)
		if err != nil {
			e.log.ErrorContext(ctx, "failed to perform request",
				e.log.Param("err", err),
				e.log.Param("attempt", i))
			continue
		}

		if res.StatusCode == http.StatusOK {
			break
		}
		res.Body.Close()
	}
	if err != nil {
		return response, DependencyError{
			fmt.Errorf("failed to connect to euvies databse due: %s", err),
		}
	}
	defer res.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return response, DependencyError{
			fmt.Errorf("failed to read response body due: %s", err),
		}
	}

	xmlResponse := Response{}
	err = xml.Unmarshal(body, &xmlResponse)
	if err != nil {
		return response, DependencyError{
			fmt.Errorf("failed to unmarshal response due: %s", err),
		}
	}

	if xmlResponse.SoapBody.FaultDetails != nil {
		return response, ValidationError{
			fmt.Errorf("eu vies returned error: %s", xmlResponse.SoapBody.FaultDetails.Faultstring),
		}
	}

	return *xmlResponse.SoapBody.Resp, nil
}

func (e *Adaptor) generateRequest(ctx context.Context, countryCode, vatID string) (req *http.Request, err error) {
	// construct request with template
	requestBody := new(Request)
	requestBody.CountryCode = countryCode
	requestBody.VATNumber = vatID

	doc := new(bytes.Buffer)
	err = e.template.Execute(doc, requestBody)
	if err != nil {
		return nil, fmt.Errorf(failedToComposeRequestDue, err)
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, e.baseURL+vatCheckEndpoint, doc)
	if err != nil {
		return nil, fmt.Errorf(failedToComposeRequestDue, err)
	}

	return r, nil
}
