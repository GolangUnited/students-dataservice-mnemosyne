package logger

import (
	"log"

	"go.uber.org/zap"
)

type ZapLogger struct {
	localLogger *zap.Logger
}

func NewZapLogger() *ZapLogger {
	localLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("logger init", err)
	}

	return &ZapLogger{localLogger: localLogger}
}

func (l ZapLogger) Infos(args ...interface{}) {
	sugar := l.localLogger.Sugar()
	sugar.Info(args...)
}

func (l ZapLogger) Info(msg string) {
	l.localLogger.Info(msg)
}

func (l ZapLogger) Infof(template string, args ...interface{}) {
	sugar := l.localLogger.Sugar()
	sugar.Infof(template, args...)
}

func (l ZapLogger) Fatalf(template string, args ...interface{}) {
	sugar := l.localLogger.Sugar()
	sugar.Fatalf(template, args...)
}
