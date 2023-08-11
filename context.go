package kirk

import (
	"context"
)

// ContextWithFields returns a new context with the given fields.
func ContextWithFields(ctx context.Context, fields LoggerFields) context.Context {
	if fields == nil {

		return ctx
	}

	existingFields := FieldsFromCtx(ctx)
	if len(existingFields) == 0 {

		return context.WithValue(ctx, FieldsKey, fields)
	}

	return context.WithValue(ctx, FieldsKey, mergeFields(existingFields, fields))
}

// FieldsFromCtx returns the fields from the given context.
func FieldsFromCtx(ctx context.Context) LoggerFields {
	if ctx == nil {
		return nil
	}

	fields, ok := ctx.Value(FieldsKey).(LoggerFields)
	if !ok {
		return nil
	}

	return fields
}

func mergeFields(a LoggerFields, b LoggerFields) LoggerFields {
	if a == nil && b == nil {
		return nil
	}

	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	merged := make(LoggerFields, len(a)+len(b))

	for k, v := range a {
		merged[k] = v
	}

	for k, v := range b {
		merged[k] = v
	}

	return merged
}
