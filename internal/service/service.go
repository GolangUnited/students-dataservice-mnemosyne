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
	AddUser(ctx context.Context, user database.User) (id int, err error)
	GetUsers(ctx context.Context) (users []database.User, err error)
	GetUserById(ctx context.Context, id int) (user database.User, err error)
	GetUserByEmail(ctx context.Context, email string) (user database.User, err error)
	UpdateUser(ctx context.Context, user database.User) (ok bool, err error)
	DeleteUser(ctx context.Context, id int) (ok bool, err error)
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
