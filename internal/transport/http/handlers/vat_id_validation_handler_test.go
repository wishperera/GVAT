package handlers

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/wishperera/GVAT/internal/application"
	"github.com/wishperera/GVAT/internal/container"
	"github.com/wishperera/GVAT/internal/mocks"
	"github.com/wishperera/GVAT/internal/services"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//nolint //test func
func TestVATIDValidation_ServeHTTP(t *testing.T) {
	type fields struct {
		container container.Container
	}
	type args struct {
		id string
	}

	type expected struct {
		body       string
		statusCode int
	}

	body := `{"vatId":"%s","valid":%v}`

	validGermanID := "DE123456789"
	invalidGermanID := "DE321321321"
	malformedGermanID := "DE123"
	nonGermanValidID := "FR1231231230"

	mockLog := mocks.NewMockLog()
	ctrl := gomock.NewController(t)
	mockAdaptor := mocks.NewMockEUVIESAdaptor(ctrl)
	mockAdaptor.EXPECT().ValidateVATID(gomock.Any(), gomock.Eq(validGermanID[:2]), gomock.Eq(validGermanID[2:])).
		Return(true, nil).AnyTimes()
	mockAdaptor.EXPECT().ValidateVATID(gomock.Any(), gomock.Eq(invalidGermanID[:2]), gomock.Eq(invalidGermanID[2:])).
		Return(false, nil).AnyTimes()

	service := new(services.ValidateVAT)
	exceptionController := new(Exception)

	mockContainer := mocks.NewMockContainer(ctrl)
	mockContainer.EXPECT().Resolve(application.ModuleLogger).Return(mockLog).AnyTimes()
	mockContainer.EXPECT().Resolve(application.ModuleEUVIESAdaptor).Return(mockAdaptor).AnyTimes()
	mockContainer.EXPECT().Resolve(application.ModuleExceptionHandler).Return(exceptionController).AnyTimes()

	err := exceptionController.Init(mockContainer)
	if err != nil {
		t.Error(err)
		return
	}

	err = service.Init(mockContainer)
	if err != nil {
		t.Error(err)
		return
	}

	dummyURL := "http://somet-server/validate/"

	mockContainer.EXPECT().Resolve(application.ModuleVATIDValidationService).Return(service).AnyTimes()

	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "valid german id",
			fields: fields{
				container: mockContainer,
			},
			args: args{
				id: validGermanID,
			},
			expected: expected{
				body:       fmt.Sprintf(body, validGermanID, true),
				statusCode: http.StatusOK,
			},
		},
		{
			name: "invalid german id",
			fields: fields{
				container: mockContainer,
			},
			args: args{
				id: invalidGermanID,
			},
			expected: expected{
				body:       fmt.Sprintf(body, invalidGermanID, false),
				statusCode: http.StatusOK,
			},
		},
		{
			name: "malformed german id",
			fields: fields{
				container: mockContainer,
			},
			args: args{
				id: malformedGermanID,
			},
			expected: expected{
				body:       fmt.Sprintf(body, malformedGermanID, false),
				statusCode: http.StatusOK,
			},
		},
		{
			name: "non german id",
			fields: fields{
				container: mockContainer,
			},
			args: args{
				id: nonGermanValidID,
			},
			expected: expected{
				body:       fmt.Sprintf(body, nonGermanValidID, false),
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		temp := tt
		t.Run(temp.name, func(t *testing.T) {
			v := &VATIDValidation{}
			err = v.Init(temp.fields.container)
			if err != nil {
				t.Error(err)
				return
			}

			recorder := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodGet, dummyURL+temp.args.id, nil)
			v.ServeHTTP(recorder, req)

			resp := recorder.Result()
			if resp.StatusCode != temp.expected.statusCode {
				t.Errorf("expected status: %d got: %d", temp.expected.statusCode, resp.StatusCode)
				return
			}
			defer resp.Body.Close()

			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
				return
			}

			respBodyStr := string(respBody)
			if respBodyStr != temp.expected.body {
				t.Errorf("expected status: %s got: %s", temp.expected.body, respBodyStr)
				return
			}

		})
	}
}
