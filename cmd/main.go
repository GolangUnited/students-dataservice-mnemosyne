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
	"google.golang.org/grpc"
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

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	ctx = log.ContextWithLogger(ctx, logger)
	defer cancel()

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(ctx, services)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	grpcService := grpc.NewServer()
	grpcServer := server.NewGrpc(ctx, grpcService)
	go func() {
		if err = grpcServer.Run(cfg.GrpcPort, handlers); err != nil {
			sugar.Info(err.Error())
			quit <- nil
		}
	}()

	restServer := server.NewRest(ctx)
	go func() {
		if err = restServer.Run(cfg.GrpcPort, cfg.RestPort); err != nil {
			sugar.Info(err.Error())
			quit <- nil
		}
	}()

	sugar.Info("App Started")

	s := <-quit
	sugar.Infof("Got signal %v, attempting graceful shutdown", s)
	cancel()
	sugar.Info("Context is stopped")
	grpcService.GracefulStop()
	sugar.Info("gRPC graceful stopped")
	err = restServer.RestServer().Shutdown(ctx)
	if err != nil {
		sugar.Infof("error rest server shutdown: %s", err.Error())
	} else {
		sugar.Info("Rest server stopped")
	}

	sugar.Info("App Shutting Down")
}
