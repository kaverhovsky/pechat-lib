package logger

import (
	"context"
	pechat_lib "github.com/kaverhovsky/pechat-lib"
	"go.uber.org/zap"
)

func PutIntoContext(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, pechat_lib.LoggerContextKey, l)
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	if l, ok := ctx.Value(pechat_lib.LoggerContextKey).(*zap.Logger); ok {
		l.Debug(msg, fields...)
		return
	}
	Logger().Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	if l, ok := ctx.Value(pechat_lib.LoggerContextKey).(*zap.Logger); ok {
		l.Info(msg, fields...)
		return
	}
	Logger().Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	if l, ok := ctx.Value(pechat_lib.LoggerContextKey).(*zap.Logger); ok {
		l.Warn(msg, fields...)
		return
	}
	Logger().Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	if l, ok := ctx.Value(pechat_lib.LoggerContextKey).(*zap.Logger); ok {
		l.Error(msg, fields...)
		return
	}
	Logger().Error(msg, fields...)
}
