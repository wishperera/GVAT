// Code generated by MockGen. DO NOT EDIT.
// Source: validate_vat.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidateVATID is a mock of ValidateVATID interface.
type MockValidateVATID struct {
	ctrl     *gomock.Controller
	recorder *MockValidateVATIDMockRecorder
}

// MockValidateVATIDMockRecorder is the mock recorder for MockValidateVATID.
type MockValidateVATIDMockRecorder struct {
	mock *MockValidateVATID
}

// NewMockValidateVATID creates a new mock instance.
func NewMockValidateVATID(ctrl *gomock.Controller) *MockValidateVATID {
	mock := &MockValidateVATID{ctrl: ctrl}
	mock.recorder = &MockValidateVATIDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidateVATID) EXPECT() *MockValidateVATIDMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockValidateVATID) Validate(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockValidateVATIDMockRecorder) Validate(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidateVATID)(nil).Validate), ctx, id)
}