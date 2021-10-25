package mocks

import (
	"context"
	"github.com/wishperera/GVAT/internal/pkg/log"
)

func NewMockLog() MockLog {
	return MockLog{}
}

type MockLog struct{}

func (m MockLog) NewLog(prefix string) log.Logger {
	return MockLog{}
}

func (m MockLog) InfoContext(ctx context.Context, message string, params ...log.Param) {}

func (m MockLog) ErrorContext(ctx context.Context, message string, params ...log.Param) {}

func (m MockLog) DebugContext(ctx context.Context, message string, params ...log.Param) {}

func (m MockLog) FatalContext(ctx context.Context, message string, params ...log.Param) {}

func (m MockLog) TraceContext(ctx context.Context, message string, params ...log.Param) {}

func (m MockLog) Info(message string, params ...log.Param) {}

func (m MockLog) Error(message string, params ...log.Param) {}

func (m MockLog) Debug(message string, param ...log.Param) {}

func (m MockLog) Fatal(message string, param ...log.Param) {}

func (m MockLog) Trace(message string, param ...log.Param) {}

func (m MockLog) Param(key, value interface{}) log.Param { return log.Param{} }
