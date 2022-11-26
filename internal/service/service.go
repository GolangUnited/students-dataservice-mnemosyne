package service

import (
	"context"

	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/internal/service/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/models/database"
	dbUser "github.com/NEKETSKY/mnemosyne/models/database/user"
	model "github.com/NEKETSKY/mnemosyne/models/mnemosyne"
	apiUser "github.com/NEKETSKY/mnemosyne/pkg/api/user"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

// Mnemosyne has test signatures
type Mnemosyne interface {
	Test(context.Context, model.Request) (model.Response, error)
	GetUserRoles(ctx context.Context, userId int) ([]database.Role, error)
	AddUser(ctx context.Context, user *apiUser.User) (id *apiUser.Id, err error)
	GetUsers(ctx context.Context, ur *apiUser.UserRequest) (users *apiUser.Users, err error)
	GetUserById(ctx context.Context, id int) (user dbUser.BaseUser, err error)
	GetUserByEmail(ctx context.Context, email string) (user dbUser.BaseUser, err error)
	UpdateUser(ctx context.Context, user dbUser.BaseUser) (ok bool, err error)
	DeleteUser(ctx context.Context, id int) (ok bool, err error)
	ActivateUser(ctx context.Context, id int) (ok bool, err error)
}

// Service represents service level
type Service struct {
	Mnemosyne
}

// NewService created new service with repository
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Mnemosyne: mnemosyne.NewService(repos.Mnemosyne, repos.Role, repos.User, repos.Interview),
	}
}
