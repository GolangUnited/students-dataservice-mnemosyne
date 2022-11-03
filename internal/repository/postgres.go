package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

// Config represents postgres connect config
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SslMode  string
}

// NewPostgresDB created postgres pgx connection
func NewPostgresDB(ctx context.Context, cfg Config) (*pgx.Conn, error) {
	db, err := pgx.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SslMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
