package main

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/configs"
	"github.com/NEKETSKY/mnemosyne/internal/handler/grpc"
	"github.com/NEKETSKY/mnemosyne/internal/handler/rest"
	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/internal/service"
	"github.com/NEKETSKY/mnemosyne/models/server"
	"github.com/NEKETSKY/mnemosyne/pkg/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

// @title Mnemosyne App API
// @version 1.0
// @description API Server for Mnemosyne application

// @host localhost:8000
// @BasePath /

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
	restHandlers := rest.NewHandler(ctx, services)
	grpcHandlers := grpc.NewHandler(ctx, services)
	quit := make(chan os.Signal, 1)

	grpcServer := new(server.Grpc)
	go func() {
		if err = grpcServer.Run(cfg.GrpcPort, grpcHandlers); err != nil {
			sugar.Info(err.Error())
			quit <- nil
		}
	}()

	restServer := new(server.Rest)
	go func() {
		if err := restServer.Run(cfg.RestPort, restHandlers.InitRoutes()); err != nil {
			sugar.Info(err.Error())
			quit <- nil
		}
	}()

	sugar.Info("App Started")
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	sugar.Info("App Shutting Down")

	if err := restServer.Shutdown(ctx); err != nil {
		sugar.Info(err.Error())
	}
}
