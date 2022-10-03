package repository

import (
	"github.com/NEKETSKY/mnemosyne/internal/repository/mnemosyne"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository.go

type Mnemosyne interface {
	Test(ctx *gin.Context) error
}

type Repository struct {
	Mnemosyne
}

func NewRepository() *Repository {
	return &Repository{
		Mnemosyne: mnemosyne.NewMnemosyne(),
	}
}
