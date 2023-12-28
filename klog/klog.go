package klog

import (
	"context"

	"github.com/KyberNetwork/logger"
)

func DefaultLogger() Logger {
	return logger.DefaultLogger()
}

type Logger logger.Logger

var log logger.Logger

type Configuration struct {
	EnableConsole    bool
	EnableJSONFormat bool
	ConsoleLevel     string
	EnableFile       bool
	FileJSONFormat   bool
	FileLevel        string
	FileLocation     string
}

type LoggerBackend logger.LoggerBackend

const (
	// LoggerBackendZap logging using Uber's zap backend
	LoggerBackendZap = LoggerBackend(logger.LoggerBackendZap)
	// LoggerBackendLogrus logging using logrus backend
	LoggerBackendLogrus = LoggerBackend(logger.LoggerBackendLogrus)
)

func InitLogger(config Configuration, backend LoggerBackend) (Logger, error) {
	var err error
	log, err = logger.InitLogger(logger.Configuration{
		EnableConsole:    config.EnableConsole,
		EnableJSONFormat: config.EnableJSONFormat,
		ConsoleLevel:     config.ConsoleLevel,
		EnableFile:       config.EnableFile,
		FileJSONFormat:   config.FileJSONFormat,
		FileLevel:        config.FileLevel,
		FileLocation:     config.FileLocation,
	}, logger.LoggerBackend(backend))
	return log, err
}

func Log() Logger {
	if log == nil {
		log = DefaultLogger()
	}
	return log
}

func NewLogger(config Configuration, backend LoggerBackend) (Logger, error) {
	return logger.NewLogger(logger.Configuration{
		EnableConsole:    config.EnableConsole,
		EnableJSONFormat: config.EnableJSONFormat,
		ConsoleLevel:     config.ConsoleLevel,
		EnableFile:       config.EnableFile,
		FileJSONFormat:   config.FileJSONFormat,
		FileLevel:        config.FileLevel,
		FileLocation:     config.FileLocation,
	}, logger.LoggerBackend(backend))
}

type CtxKeyLogger struct{}

var ctxKeyLogger CtxKeyLogger

func LoggerFromCtx(ctx context.Context) Logger {
	if ctx == nil {
		return Log()
	}
	ctxLog, _ := ctx.Value(ctxKeyLogger).(Logger)
	if ctxLog != nil {
		return ctxLog
	}
	return Log()
}

func CtxWithLogger(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, ctxKeyLogger, log)
}

func Debug(ctx context.Context, msg string) {
	LoggerFromCtx(ctx).Debug(msg)
}

func Debugf(ctx context.Context, format string, args ...any) {
	LoggerFromCtx(ctx).Debugf(format, args...)
}

func Info(ctx context.Context, msg string) {
	LoggerFromCtx(ctx).Info(msg)
}

func Infof(ctx context.Context, format string, args ...any) {
	LoggerFromCtx(ctx).Infof(format, args...)
}

func Infoln(ctx context.Context, msg string) {
	LoggerFromCtx(ctx).Infoln(msg)
}

func Warn(ctx context.Context, msg string) {
	LoggerFromCtx(ctx).Warn(msg)
}

func Warnf(ctx context.Context, format string, args ...any) {
	LoggerFromCtx(ctx).Warnf(format, args...)
}

func Error(ctx context.Context, msg string) {
	LoggerFromCtx(ctx).Error(msg)
}

func Errorf(ctx context.Context, format string, args ...any) {
	LoggerFromCtx(ctx).Errorf(format, args...)
}

func Fatal(ctx context.Context, msg string) {
	LoggerFromCtx(ctx).Fatal(msg)
}

func Fatalf(ctx context.Context, format string, args ...any) {
	LoggerFromCtx(ctx).Fatalf(format, args...)
}

type Fields logger.Fields

func WithFields(ctx context.Context, keyValues Fields) Logger {
	return LoggerFromCtx(ctx).WithFields(logger.Fields(keyValues))
}

func GetDelegate(ctx context.Context) any {
	return LoggerFromCtx(ctx).GetDelegate()
}

func SetLogLevel(ctx context.Context, level string) error {
	return LoggerFromCtx(ctx).SetLogLevel(level)
}
