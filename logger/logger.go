package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *zap.Logger
)

func SetupLogger(mode, level string) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var zapLevel zapcore.Level
	if err := zapLevel.Set(level); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	loggerConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapLevel),
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	switch mode {
	case "development":
		loggerConfig.EncoderConfig.CallerKey = "caller"
		loggerConfig.EncoderConfig.StacktraceKey = "stacktrace"
		loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		loggerConfig.Development = true
		loggerConfig.Encoding = "console"
	default:
		loggerConfig.Development = false
		loggerConfig.Encoding = "json"
	}

	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	globalLogger = logger
}

func Logger() *zap.Logger {
	return globalLogger
}
