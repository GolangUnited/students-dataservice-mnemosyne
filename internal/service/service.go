package service

import (
	"context"

	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/internal/service/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/models/database"
	dbUser "github.com/NEKETSKY/mnemosyne/models/database/user"
	model "github.com/NEKETSKY/mnemosyne/models/mnemosyne"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

// Mnemosyne has test signatures
type Mnemosyne interface {
	Test(context.Context, model.Request) (model.Response, error)
	GetUserRoles(ctx context.Context, userId int) ([]database.Role, error)
	AddUser(ctx context.Context, user *dbUser.UserFullStuff) (id int, err error)
	GetUsers(ctx context.Context, ur *dbUser.UserRequest) (users []dbUser.UserFullStuff, err error)
	GetUserById(ctx context.Context, id int) (user *dbUser.UserFullStuff, err error)
	GetUserByEmail(ctx context.Context, email string) (user *dbUser.UserFullStuff, err error)
	UpdateUser(ctx context.Context, user *dbUser.UserFullStuff) (err error)
	DeactivateUser(ctx context.Context, id int) (err error)
	ActivateUser(ctx context.Context, id int) (err error)
	GetContactById(ctx context.Context, id int) (c *dbUser.Contact, err error)
	GetResumeById(ctx context.Context, id int) (r *dbUser.Resume, err error)
	UpdateContact(ctx context.Context, contact *dbUser.Contact) (err error)
	UpdateResume(ctx context.Context, resume *dbUser.Resume) (err error)
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
