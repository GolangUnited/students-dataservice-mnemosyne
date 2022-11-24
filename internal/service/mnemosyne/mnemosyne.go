package mnemosyne

import (
	"context"

	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/models/database"
	"github.com/NEKETSKY/mnemosyne/models/mnemosyne"
)

//go:generate mockgen -source=mnemosyne.go -destination=mocks/mnemosyne.go

// Service implemented Mnemosyne interface
type Service struct {
	mnemosyne repository.Mnemosyne
	reposRole repository.Role
	reposUser repository.User
}

// NewService created Service struct
func NewService(mnemosyne repository.Mnemosyne, reposRole repository.Role, reposUser repository.User) *Service {
	return &Service{
		mnemosyne: mnemosyne,
		reposRole: reposRole,
		reposUser: reposUser,
	}
}

// Test is test demo function
func (s *Service) Test(ctx context.Context, req mnemosyne.Request) (resp mnemosyne.Response, err error) {
	_ = ctx
	_ = req
	resp = *mnemosyne.NewResponse()
	return
}

// GetUserRoles get all user roles
func (s *Service) GetUserRoles(ctx context.Context, userId int) (userRoles []database.Role, err error) {
	userRoles, err = s.reposRole.GetUserRoles(ctx, userId)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) AddUser(ctx context.Context, user database.User) (id int, err error) {
	id, err = s.reposUser.AddUser(ctx, user)
	return
}

func (s *Service) GetUsers(ctx context.Context) (users []database.User, err error) {
	users, err = s.reposUser.GetUsers(ctx)
	return
}

func (s *Service) GetUserById(ctx context.Context, id int) (user database.User, err error) {
	user, err = s.reposUser.GetUserById(ctx, id)
	return
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (user database.User, err error) {
	user, err = s.reposUser.GetUserByEmail(ctx, email)
	return
}
func (s *Service) UpdateUser(ctx context.Context, user database.User) (ok bool, err error) {
	err = s.reposUser.UpdateUserById(ctx, user)
	if err == nil {
		ok = true
	}
	return
}
func (s *Service) DeleteUser(ctx context.Context, id int) (ok bool, err error) {
	err = s.reposUser.DeactivateUserById(ctx, id)
	if err == nil {
		ok = true
	}
	return
}

func (s *Service) ActivateUser(ctx context.Context, id int) (ok bool, err error) {
	err = s.reposUser.ActivateUserById(ctx, id)
	if err == nil {
		ok = true
	}
	return
}
