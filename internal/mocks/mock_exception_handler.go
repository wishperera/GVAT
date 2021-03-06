// Code generated by MockGen. DO NOT EDIT.
// Source: ./exception.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// HandleException mocks base method.
func (m *MockHandler) HandleException(ctx context.Context, w http.ResponseWriter, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleException", ctx, w, err)
}

// HandleException indicates an expected call of HandleException.
func (mr *MockHandlerMockRecorder) HandleException(ctx, w, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleException", reflect.TypeOf((*MockHandler)(nil).HandleException), ctx, w, err)
}
