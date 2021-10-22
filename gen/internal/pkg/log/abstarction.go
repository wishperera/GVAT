package log

import (
	"context"
	"fmt"
)

type param struct {
	key   interface{}
	value interface{}
}

func (p param) String() string {
	return fmt.Sprintf("%v:%v", p.key, p.value)
}

type Logger interface {
	NewLog(prefix string) Logger
	InfoContext(ctx context.Context, message string, params ...param)
	ErrorContext(ctx context.Context, message string, params ...param)
	DebugContext(ctx context.Context, message string, params ...param)
	FatalContext(ctx context.Context, message string, params ...param)
	TraceContext(ctx context.Context, message string, params ...param)
	Info(message string, params ...param)
	Error(message string, params ...param)
	Debug(message string, param ...param)
	Fatal(message string, param ...param)
	Trace(message string, param ...param)
	Param(key, value interface{}) param
}
