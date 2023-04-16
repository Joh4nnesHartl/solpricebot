package log

import (
	"os"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	StackdriverKey = "STACKDRIVER_ENABLED"

	commonLoggerFields = []zap.Field{}
)

var Default, _ = GetLogger()

func GetLogger() (*zap.Logger, error) {
	isStackdriver := false
	if v, ok := os.LookupEnv(StackdriverKey); ok {
		if v == "true" {
			isStackdriver = true
		}
	}

	var cfg zap.Config
	if isStackdriver {
		cfg = zapdriver.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		cfg = zapdriver.NewDevelopmentConfig()
		cfg.Encoding = "console"
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, err := cfg.Build(zap.AddCallerSkip(1))
	return logger, err
}

func AddCommonFields(fields ...zap.Field) {
	commonLoggerFields = append(commonLoggerFields, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	fields = append(fields, commonLoggerFields...)
	Default.Fatal(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	fields = append(fields, commonLoggerFields...)
	Default.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	fields = append(fields, commonLoggerFields...)
	Default.Warn(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	fields = append(fields, commonLoggerFields...)
	Default.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	fields = append(fields, commonLoggerFields...)
	Default.Debug(msg, fields...)
}
