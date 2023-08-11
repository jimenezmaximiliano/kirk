package kirk

import (
	"context"
)

// LoggerFieldsKey is the type of the key used to store the logger fields in the context.
type LoggerFieldsKey int

// FieldsKey is the key used to store the logger fields in the context.
const FieldsKey LoggerFieldsKey = iota

// LoggerFields represents the fields to be logged.
type LoggerFields map[string]string

// Logger represents the common logging methods supported.
type Logger interface {
	Error(ctx context.Context, err error)
	Panic(ctx context.Context, err error)
	Debug(ctx context.Context, message string)
	Info(ctx context.Context, message string)
	Warn(ctx context.Context, message string)
}

// Reporter represents the common error reporting methods supported.
type Reporter interface {
	ReportError(ctx context.Context, errorToReport error)
}

// LoggerForReporter represents the logger that the reporter will use to log errors when reporting them.
type LoggerForReporter interface {
	Error(ctx context.Context, err error)
}
