package mnemosyne

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Mnemosyne struct {
	db *pgx.Conn
}

// NewMnemosyne created Mnemosyne struct
func NewMnemosyne(db *pgx.Conn) *Mnemosyne {
	return &Mnemosyne{
		db: db,
	}
}

// Test is test demo function
func (c *Mnemosyne) Test(ctx context.Context) (err error) {
	_ = ctx

	return
}
