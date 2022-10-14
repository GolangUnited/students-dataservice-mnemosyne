package main

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/configs"
	"github.com/NEKETSKY/mnemosyne/internal/handler"
	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/internal/service"
	"github.com/NEKETSKY/mnemosyne/models/server"
	"github.com/NEKETSKY/mnemosyne/pkg/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	sugar := logger.Sugar()

	cfg, err := configs.Init()
	if err != nil {
		sugar.Fatalf("error init config: %s", err.Error())
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	ctx = log.ContextWithLogger(ctx, logger)
	defer func() {
		stop()
		sugar.Info("Context is stopped")
	}()

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(ctx, services)
	quit := make(chan os.Signal, 1)

	grpcServer := new(server.Grpc)
	go func() {
		if err = grpcServer.Run(ctx, cfg.GrpcPort, handlers); err != nil {
			sugar.Info(err.Error())
			quit <- nil
		}
	}()

	restServer := new(server.Rest)
	go func() {
		if err = restServer.Run(ctx, cfg.GrpcPort, cfg.RestPort); err != nil {
			sugar.Info(err.Error())
			quit <- nil
		}
	}()

	sugar.Info("App Started")
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	sugar.Info("App Shutting Down")
}
