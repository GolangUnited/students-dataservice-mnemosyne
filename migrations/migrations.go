package migrations

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/repository"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed *.sql
var schemaFs embed.FS

// MigrateUp exec migrate up on run instance
func MigrateUp(ctx context.Context, cfg repository.Config) (err error) {
	_ = ctx

	d, err := iofs.New(schemaFs, ".") // Get migrations from sql folder
	if err != nil {
		return
	}

	db, err := sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SslMode))
	if err != nil {
		return
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			logger.Infof("migrate close db: %s", err.Error())
		}
	}(db)
	driver, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		return
	}

	m, err := migrate.NewWithInstance("iofs", d, cfg.DBName, driver)
	if err != nil {
		return
	}

	errUp := m.Up()
	if errUp != nil {
		logger.Infof("Migrate: %s", errUp.Error())
	} else {
		logger.Info("Migrate Up")
	}

	return
}
