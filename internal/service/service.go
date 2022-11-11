package service

import (
	"context"

	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/internal/service/mnemosyne"
	model "github.com/NEKETSKY/mnemosyne/models/mnemosyne"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

// Mnemosyne has test signatures
type Mnemosyne interface {
	Test(context.Context, model.Request) (model.Response, error)
}

// Service represents service level
type Service struct {
	Mnemosyne
}

// NewService created new service with repository
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Mnemosyne: mnemosyne.NewService(repos.Mnemosyne, repos.Role, repos.User),
	}
}
