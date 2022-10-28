package mnemosyne

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Mnemosyne struct {
	db *pgx.Conn
}

func NewMnemosyne(db *pgx.Conn) *Mnemosyne {
	return &Mnemosyne{
		db: db,
	}
}

func (c *Mnemosyne) Test(ctx context.Context) (err error) {
	_ = ctx

	return
}
