package services

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/wishperera/GVAT/internal/application"
	"github.com/wishperera/GVAT/internal/container"
	"github.com/wishperera/GVAT/internal/domain"
	"github.com/wishperera/GVAT/internal/domain/adaptors"
	"github.com/wishperera/GVAT/internal/mocks"
	"github.com/wishperera/GVAT/internal/pkg/log"
	"testing"
)

func TestValidateVAT_validateFormat(t *testing.T) {
	type fields struct {
		log           log.Logger
		euViesAdaptor adaptors.EUVIESAdaptor
	}
	type args struct {
		in0 context.Context
		id  string
	}

	ctrl := gomock.NewController(t)
	mockLog := mocks.NewMockLog()
	mockAdaptor := mocks.NewMockEUVIESAdaptor(ctrl)

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValid bool
		wantErr   bool
	}{
		{
			name: "valid german id",
			fields: fields{
				mockLog,
				mockAdaptor,
			},
			args: args{
				context.Background(),
				"DE123456789",
			},
			wantErr:   false,
			wantValid: true,
		},
		{
			name: "valid non-german id",
			fields: fields{
				mockLog,
				mockAdaptor,
			},
			args: args{
				context.Background(),
				"FR123456789",
			},
			wantErr:   false,
			wantValid: false,
		},
		{
			name: "invalid german vat id",
			fields: fields{
				mockLog,
				mockAdaptor,
			},
			args: args{
				context.Background(),
				"DE1234567",
			},
			wantErr:   false,
			wantValid: false,
		},
	}
	for _, tt := range tests {
		temp := tt
		t.Run(temp.name, func(t *testing.T) {
			v := &ValidateVAT{
				log: temp.fields.log,
			}
			v.adaptors.euVies = temp.fields.euViesAdaptor

			valid, err := v.validateFormat(temp.args.in0, temp.args.id)
			if (err != nil) != temp.wantErr {
				t.Errorf("validateFormat() error = %v, wantErr %v", err != nil, temp.wantErr)
			}

			if valid != temp.wantValid {
				t.Errorf("validity expected: %v got: %v", temp.wantValid, valid)
			}
		})
	}
}

func TestValidateVAT_checkAgainstVIES(t *testing.T) {
	type fields struct {
		container container.Container
	}
	type args struct {
		ctx         context.Context
		countryCode string
		id          string
	}

	ctrl := gomock.NewController(t)
	mockLog := mocks.NewMockLog()
	mockAdaptor := mocks.NewMockEUVIESAdaptor(ctrl)
	mockContainer := mocks.NewMockContainer(ctrl)

	mockContainer.EXPECT().Resolve(application.ModuleLogger).Return(mockLog).AnyTimes()
	mockContainer.EXPECT().Resolve(application.ModuleEUVIESAdaptor).Return(mockAdaptor).AnyTimes()

	validVAT := "DE123456789"
	invalidVAT := "DE111111111"
	serviceFailureVAT := "AB123"

	mockAdaptor.EXPECT().ValidateVATID(gomock.Any(), gomock.Any(), validVAT).Return(true, nil)
	mockAdaptor.EXPECT().ValidateVATID(gomock.Any(), gomock.Any(), invalidVAT).Return(false, nil)
	mockAdaptor.EXPECT().ValidateVATID(gomock.Any(), gomock.Any(), serviceFailureVAT).Return(false, errors.New(
		"service failed"))

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValid bool
		wantErr   bool
	}{
		{
			name: "happy path",
			fields: fields{
				container: mockContainer,
			},
			args: args{
				ctx:         context.Background(),
				countryCode: domain.CountryCodeGermany,
				id:          validVAT,
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "invalid vat id",
			fields: fields{
				container: mockContainer,
			},
			args: args{
				ctx:         context.Background(),
				countryCode: domain.CountryCodeGermany,
				id:          invalidVAT,
			},
			wantValid: false,
			wantErr:   false,
		},
		{
			name: "service returns error",
			fields: fields{
				container: mockContainer,
			},
			args: args{
				ctx:         context.Background(),
				countryCode: domain.CountryCodeGermany,
				id:          serviceFailureVAT,
			},
			wantValid: false,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		temp := tt
		t.Run(temp.name, func(t *testing.T) {
			v := &ValidateVAT{}
			err := v.Init(temp.fields.container)
			if err != nil {
				t.Error(err)
				t.Fail()
			}

			gotValid, err := v.checkAgainstVIES(temp.args.ctx, temp.args.countryCode, temp.args.id)
			if (err != nil) != temp.wantErr {
				t.Errorf("checkAgainstVIES() error = %v, wantErr %v", err, temp.wantErr)
				return
			}
			if gotValid != temp.wantValid {
				t.Errorf("checkAgainstVIES() gotValid = %v, want %v", gotValid, temp.wantValid)
			}
		})
	}
}
