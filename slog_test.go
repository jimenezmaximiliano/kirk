package kirk_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jimenezmaximiliano/kirk"
)

func TestLogInfoWithFields(test *testing.T) {
	var outputBuffer bytes.Buffer
	logger := getLogger(&outputBuffer)
	ctx := context.Background()

	ctx = kirk.ContextWithFields(ctx, kirk.LoggerFields{
		"obladi": "oblada",
	})

	logger.Info(ctx, "test message")

	output := getLoggerOutput(test, &outputBuffer)

	assert.Equal(test, "test message", output["msg"])
	assert.Equal(test, "oblada", output["obladi"])
	assert.Equal(test, "INFO", output["level"])
}

func getLogger(outputBuffer *bytes.Buffer) kirk.SlogAdapter {
	jsonHandler := slog.NewJSONHandler(outputBuffer, nil)

	return kirk.NewSlogLogger(jsonHandler)
}

func getLoggerOutput(test *testing.T, outputBuffer *bytes.Buffer) map[string]string {
	jsonOutput := map[string]string{}
	err := json.Unmarshal(outputBuffer.Bytes(), &jsonOutput)
	require.NoError(test, err)

	return jsonOutput
}
