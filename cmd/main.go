package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/GolangUnited/students-dataservice-mnemosyne/configs"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/handler"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/repository"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/service"
	"github.com/GolangUnited/students-dataservice-mnemosyne/migrations"
	"github.com/GolangUnited/students-dataservice-mnemosyne/models/server"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/logger"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// init config
	cfg, err := configs.Init()
	if err != nil {
		logger.Fatalf("error init config: %s", err.Error())
	}

	// init context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err = godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	// init postgres db
	dbCfg := repository.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB_NAME"),
		SslMode:  os.Getenv("POSTGRES_SSL"),
	}
	db, err := repository.NewPostgresDB(ctx, dbCfg)
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer func(db *pgx.Conn, ctx context.Context) {
		err = db.Close(ctx)
		if err != nil {
			logger.Infos("error close db conn: %s", err.Error())
		}
	}(db, ctx)

	// init migrations
	err = migrations.MigrateUp(ctx, dbCfg)
	if err != nil {
		logger.Fatalf("error init db migrate: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(ctx, services)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// run grpc server
	grpcServer := server.NewGrpc(ctx, handlers)
	go func() {
		if err = grpcServer.Run(cfg.GrpcPort); err != nil {
			logger.Info(err.Error())
			quit <- nil
		}
	}()

	// run rest server
	restServer := server.NewRest(ctx)
	go func() {
		if err = restServer.Run(cfg.GrpcPort, cfg.RestPort); err != nil {
			logger.Info(err.Error())
			quit <- nil
		}
	}()

	logger.Info("App Started")

	// graceful shutdown
	logger.Infof("Got signal %v, attempting graceful shutdown", <-quit)
	cancel()
	logger.Info("Context is stopped")
	grpcServer.GracefulStop()
	logger.Info("gRPC graceful stopped")
	err = restServer.Shutdown()
	if err != nil {
		logger.Infof("error rest server shutdown: %s", err.Error())
	} else {
		logger.Info("Rest server stopped")
	}

	logger.Info("App Shutting Down")
}
