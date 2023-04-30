package kirk

import (
	"go.uber.org/zap"
)

// NewLoggerFromSugaredZap returns an adapted zap SugaredLogger that implements Logger.
func NewLoggerFromSugaredZap(zap zap.SugaredLogger) ZapLoggerAdapter {
	return ZapLoggerAdapter{zap: zap}
}
