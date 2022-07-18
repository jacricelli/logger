package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetNewInstance() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.StacktraceKey = zapcore.OmitKey
	config.EncoderConfig = encoderConfig

	logger, err := config.Build(zap.WithCaller(false))
	if err != nil {
		return nil, err
	}

	return logger, nil
}
