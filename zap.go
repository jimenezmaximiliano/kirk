package kirk

import (
	"go.uber.org/zap"
)

type ZapLoggerAdapter struct {
	zap zap.SugaredLogger
}

var _ Logger = ZapLoggerAdapter{}

// Error logs an error with ERROR severity.
func (zap ZapLoggerAdapter) Error(err error) {
	zap.zap.Error(err)
}

// Panic logs an error with PANIC severity, and panics.
func (zap ZapLoggerAdapter) Panic(err error) {
	zap.zap.Panic(err)
}

// Debug logs a message with DEBUG severity.
func (zap ZapLoggerAdapter) Debug(message string) {
	zap.zap.Debug(message)
}

// Info logs a message with INFO severity.
func (zap ZapLoggerAdapter) Info(message string) {
	zap.zap.Info(message)
}

// Warn logs a message with WARN severity.
func (zap ZapLoggerAdapter) Warn(message string) {
	zap.zap.Warn(message)
}
