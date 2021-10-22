package log

import (
	"context"
	"github.com/wishperera/GVAT/gen/internal/container"
)

type Log struct {
	prefix string
}

func (l *Log) Init(c container.Container) error {
	config := c.Resolve()
}

func (l Log) InfoContext(ctx context.Context, message string, params ...param) {
	panic("implement me")
}

func (l Log) ErrorContext(ctx context.Context, message string, params ...param) {
	panic("implement me")
}

func (l Log) DebugContext(ctx context.Context, message string, params ...param) {
	panic("implement me")
}

func (l Log) FatalContext(ctx context.Context, message string, params ...param) {
	panic("implement me")
}

func (l Log) TraceContext(ctx context.Context, message string, params ...param) {
	panic("implement me")
}

func (l Log) Info(message string, params ...param) {
	panic("implement me")
}

func (l Log) Error(message string, params ...param) {
	panic("implement me")
}

func (l Log) Debug(message string, param ...param) {
	panic("implement me")
}

func (l Log) Fatal(message string, param ...param) {
	panic("implement me")
}

func (l Log) Trace(message string, param ...param) {
	panic("implement me")
}

func (l Log) Param(key, value interface{}) param {
	panic("implement me")
}
