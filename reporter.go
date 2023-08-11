package kirk

import (
	"context"
)

// NewLoggerWithReporter returns a new logger that will report errors with the given reporter.
func NewLoggerWithReporter(logger Logger, reporter Reporter) LoggerWithReporter {
	return LoggerWithReporter{
		logger:   logger,
		reporter: reporter,
	}
}

// Make sure LoggerWithReporter implements Logger.
var _ Logger = LoggerWithReporter{}

// LoggerWithReporter is a logger that will report errors with the given reporter.
type LoggerWithReporter struct {
	logger   Logger
	reporter Reporter
}

// Error logs an error with ERROR severity and reports it.
func (logger LoggerWithReporter) Error(ctx context.Context, err error) {
	logger.logger.Error(ctx, err)
	logger.reporter.ReportError(ctx, err)
}

// Panic logs an error with PANIC severity, reports it, and panics.
func (logger LoggerWithReporter) Panic(ctx context.Context, err error) {
	logger.reporter.ReportError(ctx, err)
	logger.logger.Panic(ctx, err)
}

// Debug logs a message with DEBUG severity.
func (logger LoggerWithReporter) Debug(ctx context.Context, message string) {
	logger.logger.Debug(ctx, message)
}

// Info logs a message with INFO severity.
func (logger LoggerWithReporter) Info(ctx context.Context, message string) {
	logger.logger.Info(ctx, message)
}

// Warn logs a message with WARN severity.
func (logger LoggerWithReporter) Warn(ctx context.Context, message string) {
	logger.logger.Warn(ctx, message)
}
