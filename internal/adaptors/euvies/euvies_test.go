package euvies

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/wishperera/GVAT/internal/application"
	"github.com/wishperera/GVAT/internal/mocks"
	"io/ioutil"
	"testing"
)

func TestAdaptor_generateRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLog := mocks.NewMockLog()
	config := new(Config)
	config.URL = "random url"
	mockContainer := mocks.NewMockContainer(ctrl)
	mockContainer.EXPECT().GetModuleConfig(application.ModuleEUVIESAdaptor).Return(config).AnyTimes()
	mockContainer.EXPECT().Resolve(application.ModuleLogger).Return(mockLog).AnyTimes()

	adaptor := Adaptor{}
	err := adaptor.Init(mockContainer)
	if err != nil {
		t.Errorf("failed to initialize adaptor due: %s", err)
		return
	}

	countryCode := "DE"
	vatID := "123456"

	req, err := adaptor.generateRequest(context.Background(), countryCode, vatID)
	if err != nil {
		t.Error(err)
		return
	}

	reqURL := req.URL.Path
	expectedURL := config.URL + vatCheckEndpoint
	if reqURL != expectedURL {
		t.Logf("exptected: %s received: %s", expectedURL, reqURL)
		t.Fail()
		return
	}

	reqBodyByt, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Error(err)
		return
	}
	defer req.Body.Close()

	expectedBody := `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
   <soapenv:Header/>
   <soapenv:Body>
      <urn:checkVat>
         <urn:countryCode>DE</urn:countryCode>
         <urn:vatNumber>123456</urn:vatNumber>
      </urn:checkVat>
   </soapenv:Body>
</soapenv:Envelope>
`

	if expectedBody != string(reqBodyByt) {
		t.Logf("exptected: %+v received: %+v", expectedBody, reqBodyByt)
		t.Fail()
		return
	}
}
