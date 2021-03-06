// Code generated by MockGen. DO NOT EDIT.
// Source: ./eu_vies.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEUVIESAdaptor is a mock of EUVIESAdaptor interface.
type MockEUVIESAdaptor struct {
	ctrl     *gomock.Controller
	recorder *MockEUVIESAdaptorMockRecorder
}

// MockEUVIESAdaptorMockRecorder is the mock recorder for MockEUVIESAdaptor.
type MockEUVIESAdaptorMockRecorder struct {
	mock *MockEUVIESAdaptor
}

// NewMockEUVIESAdaptor creates a new mock instance.
func NewMockEUVIESAdaptor(ctrl *gomock.Controller) *MockEUVIESAdaptor {
	mock := &MockEUVIESAdaptor{ctrl: ctrl}
	mock.recorder = &MockEUVIESAdaptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEUVIESAdaptor) EXPECT() *MockEUVIESAdaptorMockRecorder {
	return m.recorder
}

// ValidateVATID mocks base method.
func (m *MockEUVIESAdaptor) ValidateVATID(ctx context.Context, countryCode, vatID string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateVATID", ctx, countryCode, vatID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateVATID indicates an expected call of ValidateVATID.
func (mr *MockEUVIESAdaptorMockRecorder) ValidateVATID(ctx, countryCode, vatID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateVATID", reflect.TypeOf((*MockEUVIESAdaptor)(nil).ValidateVATID), ctx, countryCode, vatID)
}
