package repository

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/internal/repository/mnemosyne"
	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository.go

type Mnemosyne interface {
	Test(ctx context.Context) error
}

type Repository struct {
	Mnemosyne
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Mnemosyne: mnemosyne.NewMnemosyne(db),
	}
}
