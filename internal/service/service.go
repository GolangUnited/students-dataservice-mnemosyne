package service

import (
	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/internal/service/mnemosyne"
	model "github.com/NEKETSKY/mnemosyne/models/mnemosyne"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

type Mnemosyne interface {
	Test(*gin.Context, model.Request) (model.Response, error)
}

type Service struct {
	Mnemosyne
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Mnemosyne: mnemosyne.NewService(repos.Mnemosyne),
	}
}
