package kirk

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type SlogAdapter struct {
	logger *slog.Logger
}

var _ Logger = SlogAdapter{}

func NewDefaultSlogLogger() SlogAdapter {
	return SlogAdapter{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func NewSlogLogger(handler slog.Handler) SlogAdapter {
	return SlogAdapter{
		logger: slog.New(handler),
	}
}

func (adapter SlogAdapter) Debug(ctx context.Context, message string) {
	fields := getSlogFieldsFromContext(ctx)

	if len(fields) > 0 {
		adapter.logger.DebugContext(ctx, message, fields...)

		return
	}

	adapter.logger.DebugContext(ctx, message)
}

func (adapter SlogAdapter) Info(ctx context.Context, message string) {
	fields := getSlogFieldsFromContext(ctx)

	if len(fields) > 0 {
		adapter.logger.InfoContext(ctx, message, fields...)

		return
	}

	adapter.logger.InfoContext(ctx, message)
}

func (adapter SlogAdapter) Warn(ctx context.Context, message string) {
	fields := getSlogFieldsFromContext(ctx)

	if len(fields) > 0 {
		adapter.logger.WarnContext(ctx, message, fields...)

		return
	}

	adapter.logger.WarnContext(ctx, message)
}

func (adapter SlogAdapter) Error(ctx context.Context, err error) {
	fields := getSlogFieldsFromContext(ctx)

	if len(fields) > 0 {
		adapter.logger.ErrorContext(ctx, fmt.Sprintf("%+v", err), fields...)

		return
	}

	adapter.logger.ErrorContext(ctx, fmt.Sprintf("%+v", err))
}

func (adapter SlogAdapter) Panic(ctx context.Context, err error) {
	fields := getSlogFieldsFromContext(ctx)

	if len(fields) > 0 {
		adapter.logger.ErrorContext(ctx, fmt.Sprintf("%+v", err), fields...)
	} else {
		adapter.logger.ErrorContext(ctx, fmt.Sprintf("%+v", err))
	}

	panic(err)
}

func getSlogFieldsFromContext(ctx context.Context) []any {
	fieldsMap := FieldsFromCtx(ctx)
	var slogFields []any

	for key, value := range fieldsMap {
		slogFields = append(slogFields, key, value)
	}

	if len(slogFields) == 0 {
		return nil
	}

	return slogFields
}
