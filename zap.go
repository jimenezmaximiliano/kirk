package kirk

import (
	"go.uber.org/zap"
)

type zapAdapter struct {
	zap zap.SugaredLogger
}

var _ Logger = zapAdapter{}

func (zap zapAdapter) Error(err error) {
	zap.zap.Error(err)
}

func (zap zapAdapter) Panic(err error) {
	zap.zap.Panic(err)
}

func (zap zapAdapter) Debug(message string) {
	zap.zap.Debug(message)
}

func (zap zapAdapter) Info(message string) {
	zap.zap.Info(message)
}

func (zap zapAdapter) Warn(message string) {
	zap.zap.Warn(message)
}
