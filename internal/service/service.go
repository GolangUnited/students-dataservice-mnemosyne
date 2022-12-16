package service

import (
	"context"

	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/internal/service/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/models/database"
	model "github.com/NEKETSKY/mnemosyne/models/mnemosyne"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

// Mnemosyne has test signatures
type Mnemosyne interface {
	Test(context.Context, model.Request) (model.Response, error)
	GetUserRoles(ctx context.Context, userId int) ([]database.Role, error)
	Certificate
}

type Certificate interface {
	CreateCertificate(ctx context.Context, certificate database.Certificate) (certificateId uint32, err error)
	GetCertificates(ctx context.Context, userId uint32) (certificates []database.Certificate, err error)
	UpdateCertificate(ctx context.Context, certificate database.Certificate) (err error)
	DeactivateCertificate(ctx context.Context, certificateId uint32) (err error)
	ActivateCertificate(ctx context.Context, certificateId uint32) (err error)
}

// Service represents service level
type Service struct {
	Mnemosyne
}

// NewService created new service with repository
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Mnemosyne: mnemosyne.NewService(repos.Mnemosyne, repos.Role, repos.User, repos.Interview, repos.Certificate),
	}
}
