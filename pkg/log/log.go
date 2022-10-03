package log

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ctxLogger struct{}

const LoggerKey = "logger"

// ContextWithLogger adds logger to context
func ContextWithLogger(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext returns logger from context
func LoggerFromContext(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*zap.Logger); ok {
		return l
	}
	return zap.L()
}

// LoggerFromGinContext returns logger from gin context
func LoggerFromGinContext(ctx *gin.Context) *zap.Logger {
	if l, ok := ctx.Get(LoggerKey); ok {
		return l.(*zap.Logger)
	}
	return zap.L()
}
