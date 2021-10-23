package log

import (
	"context"
	"fmt"
	"github.com/wishperera/GVAT/gen/internal/application"
	"github.com/wishperera/GVAT/gen/internal/container"
	pkgCtx "github.com/wishperera/GVAT/gen/internal/pkg/context"
	"log"
)

type logLevel int

func (l logLevel) String() string {
	switch l {
	case 1:
		return "TRACE"
	case 2:
		return "DEBUG"
	case 3:
		return "INFO"
	case 4:
		return "ERROR"
	case 5:
		return "FATAL"
	default:
		return ""
	}
}

func parseLogLevel(str string) (l logLevel, err error) {
	switch str {
	case "TRACE":
		return logLevelTrace, nil
	case "DEBUG":
		return logLevelDebug, nil
	case "INFO":
		return logLevelInfo, nil
	case "ERROR":
		return logLevelError, nil
	case "FATAL":
		return logLevelFatal, nil
	default:
		return l, fmt.Errorf("unknown log level: %s", str)
	}
}

const (
	logLevelTrace logLevel = iota + 1
	logLevelDebug
	logLevelInfo
	logLevelError
	logLevelFatal
)

const (
	// logMessageFormat : [trace-id][log level][prefix][message][params]
	logMessageFormat = "[%s] [%s] %s [%s][%s]"
)

type Log struct {
	prefix string
	level  logLevel
}

func (l *Log) Init(c container.Container) error {
	config := c.Resolve(application.ModuleConfigLogger).(*Config)
	l.level = config.Level
	l.prefix = "[application]"

	return nil
}

func (l *Log) NewLog(prefix string) Logger {
	newLog := new(Log)
	newLog.level = l.level
	newLog.prefix = l.prefix + fmt.Sprintf(" [%s]", prefix)

	return newLog
}

func (l Log) InfoContext(ctx context.Context, message string, params ...param) {
	if l.level <= logLevelInfo {
		log.Println(composeMessage(ctx, logLevelInfo, l.prefix, message, params...))
	}
}

func (l Log) ErrorContext(ctx context.Context, message string, params ...param) {
	if l.level <= logLevelError {
		log.Println(composeMessage(ctx, logLevelError, l.prefix, message, params...))
	}
}

func (l Log) DebugContext(ctx context.Context, message string, params ...param) {
	if l.level <= logLevelDebug {
		log.Println(composeMessage(ctx, logLevelDebug, l.prefix, message, params...))
	}
}

func (l Log) FatalContext(ctx context.Context, message string, params ...param) {
	if l.level <= logLevelFatal {
		log.Fatalln(composeMessage(ctx, logLevelFatal, l.prefix, message, params...))
	}
}

func (l Log) TraceContext(ctx context.Context, message string, params ...param) {
	if l.level <= logLevelTrace {
		log.Fatalln(composeMessage(ctx, logLevelTrace, l.prefix, message, params...))
	}
}

func (l Log) Info(message string, params ...param) {
	l.InfoContext(context.Background(), message, params...)
}

func (l Log) Error(message string, params ...param) {
	l.ErrorContext(context.Background(), message, params...)
}

func (l Log) Debug(message string, params ...param) {
	l.DebugContext(context.Background(), message, params...)
}

func (l Log) Fatal(message string, params ...param) {
	l.FatalContext(context.Background(), message, params...)
}

func (l Log) Trace(message string, params ...param) {
	l.TraceContext(context.Background(), message, params...)
}

func (l Log) Param(key, value interface{}) param {
	return param{
		key:   key,
		value: value,
	}
}

func composeMessage(ctx context.Context, level logLevel, prefix, message string, params ...param) string {
	paramString := ""
	for _, v := range params {
		paramString += v.String()
	}
	str := fmt.Sprintf(logMessageFormat, level, pkgCtx.ExtractTrace(ctx), prefix, message, paramString)

	return str
}
