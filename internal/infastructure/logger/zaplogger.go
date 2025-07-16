package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() (*ZapLogger, error) {
	z, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("error NewZapLogger():%v", err)
	}
	return &ZapLogger{
		logger: z,
	}, nil
}

func (z *ZapLogger) Fatal(massage string, filed ...zap.Field) {
	z.logger.Fatal(massage, filed...)
}

func (z *ZapLogger) Error(massage string, filed ...zap.Field) {
	z.logger.Error(massage, filed...)
}

func (z *ZapLogger) Info(massage string, filed ...zap.Field) {
	z.logger.Info(massage, filed...)
}
