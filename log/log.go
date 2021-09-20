package log

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	options = []zap.Option{
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}

	logger *zap.SugaredLogger
)

func init() {
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	logger = l.WithOptions(options...).Sugar()
}

// Error outputs an error log.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf outputs a formatted error log.
func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

// Info outputs an informational log.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof outputs a formatted info log.
func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Fatal outputs an error log and exits.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf outputs a formatted error log and exits.
func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}
