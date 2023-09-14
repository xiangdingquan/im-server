package log

import (
	"context"
	"fmt"
)

// Logger is used for logging formatted messages.
type Logger interface {
	// Debugf logs messages at DEBUG level.
	Debugf(format string, args ...interface{})
	// Infof logs messages at INFO level.
	Infof(format string, args ...interface{})
	// Warnf logs messages at WARN level.
	Warnf(format string, args ...interface{})
	// Errorf logs messages at ERROR level.
	Errorf(format string, args ...interface{})
	// Fatalf logs messages at FATAL level.
	Fatalf(format string, args ...interface{})
}

type DefaultLogger struct {
}

func NewDefaultLogger() Logger {
	return new(DefaultLogger)
}

func (m *DefaultLogger) Debugf(format string, args ...interface{}) {
	h.Log(context.Background(), _debugLevel, KVString(_log, fmt.Sprintf(format, args...)))
}

func (m *DefaultLogger) Infof(format string, args ...interface{}) {
	h.Log(context.Background(), _infoLevel, KVString(_log, fmt.Sprintf(format, args...)))
}

func (m *DefaultLogger) Warnf(format string, args ...interface{}) {
	h.Log(context.Background(), _warnLevel, KVString(_log, fmt.Sprintf(format, args...)))
}

func (m *DefaultLogger) Errorf(format string, args ...interface{}) {
	h.Log(context.Background(), _errorLevel, KVString(_log, fmt.Sprintf(format, args...)))
}

func (m *DefaultLogger) Fatalf(format string, args ...interface{}) {
	h.Log(context.Background(), _fatalLevel, KVString(_log, fmt.Sprintf(format, args...)))
}
