package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true

	// using json format if app_env is production or development
	if os.Getenv("APP_ENV") == "production" || os.Getenv("APP_ENV") == "development" {
		config = zap.NewProductionConfig()
		config.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		config.DisableStacktrace = false
	}

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000000Z")
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)

	return logger.Sugar()
}
