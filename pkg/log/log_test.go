package log

import (
	"context"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestLoggerFromContext(t *testing.T) {
	ctx := context.Background()
	logger := zap.L()
	ContextWithLogger(ctx, logger)

	tests := []struct {
		name string
		ctx  context.Context
		want *zap.Logger
	}{
		{
			name: "Context with logger",
			ctx:  ctx,
			want: logger,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoggerFromContext(tt.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoggerFromContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContextWithLogger(t *testing.T) {
	ctx := context.Background()
	logger := zap.L()
	ctxWithLogger := ContextWithLogger(ctx, logger)

	tests := []struct {
		name   string
		ctx    context.Context
		logger *zap.Logger
		want   context.Context
	}{
		{
			name:   "Logger from context",
			ctx:    ctx,
			logger: logger,
			want:   ctxWithLogger,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContextWithLogger(tt.ctx, tt.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContextWithLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
