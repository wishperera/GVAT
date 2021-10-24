package log

import (
	"context"
	"fmt"
)

type Param struct {
	key   interface{}
	value interface{}
}

func (p Param) String() string {
	return fmt.Sprintf("%v:%v", p.key, p.value)
}

type Logger interface {
	NewLog(prefix string) Logger
	InfoContext(ctx context.Context, message string, params ...Param)
	ErrorContext(ctx context.Context, message string, params ...Param)
	DebugContext(ctx context.Context, message string, params ...Param)
	FatalContext(ctx context.Context, message string, params ...Param)
	TraceContext(ctx context.Context, message string, params ...Param)
	Info(message string, params ...Param)
	Error(message string, params ...Param)
	Debug(message string, param ...Param)
	Fatal(message string, param ...Param)
	Trace(message string, param ...Param)
	Param(key, value interface{}) Param
}
