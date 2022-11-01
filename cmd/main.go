package main

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/configs"
	"github.com/NEKETSKY/mnemosyne/internal/handler"
	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/internal/service"
	"github.com/NEKETSKY/mnemosyne/models/server"
	"github.com/NEKETSKY/mnemosyne/pkg/logger"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	cfg, err := configs.Init()
	if err != nil {
		logger.Fatalf("error init config: %s", err.Error())
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err = godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(ctx, repository.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB_NAME"),
		SslMode:  os.Getenv("POSTGRES_SSL"),
	})
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer func(db *pgx.Conn, ctx context.Context) {
		err = db.Close(ctx)
		if err != nil {
			logger.Infos("error close db conn: %s", err.Error())
		}
	}(db, ctx)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(ctx, services)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	grpcService := grpc.NewServer()
	grpcServer := server.NewGrpc(ctx, grpcService)
	go func() {
		if err = grpcServer.Run(cfg.GrpcPort, handlers); err != nil {
			logger.Info(err.Error())
			quit <- nil
		}
	}()

	restServer := server.NewRest(ctx)
	go func() {
		if err = restServer.Run(cfg.GrpcPort, cfg.RestPort); err != nil {
			logger.Info(err.Error())
			quit <- nil
		}
	}()

	logger.Info("App Started")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-quit
		logger.Infof("Got signal %v, attempting graceful shutdown", s)
		cancel()
		logger.Info("Context is stopped")
		grpcService.GracefulStop()
		logger.Info("gRPC graceful stopped")
		err = restServer.RestServer().Shutdown(ctx)
		if err != nil {
			logger.Infof("error rest server shutdown: %s", err.Error())
		} else {
			logger.Info("Rest server stopped")
		}
		wg.Done()
	}()
	wg.Wait()

	logger.Info("App Shutting Down")
}
