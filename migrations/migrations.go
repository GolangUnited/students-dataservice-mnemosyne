package migrations

import (
	"context"
	"database/sql"
	"embed"
	"github.com/NEKETSKY/mnemosyne/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	mpgx "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

//go:embed *.sql
var schemaFs embed.FS

// MigrateUp exec migrate up on run instance
func MigrateUp(ctx context.Context, cfg *pgx.ConnConfig) (err error) {
	_ = ctx

	d, err := iofs.New(schemaFs, ".") // Get migrations from sql folder
	if err != nil {
		return
	}

	db := stdlib.OpenDB(*cfg)
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			logger.Infof("migrate close db: %s", err.Error())
		}
	}(db)
	driver, err := mpgx.WithInstance(db, &mpgx.Config{})
	if err != nil {
		return
	}

	m, err := migrate.NewWithInstance("iofs", d, cfg.Database, driver)
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
