package mnemosyne

import (
	"context"

	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/models/database"
	"github.com/NEKETSKY/mnemosyne/models/mnemosyne"
	apiUser "github.com/NEKETSKY/mnemosyne/pkg/api/user"
)

//go:generate mockgen -source=mnemosyne.go -destination=mocks/mnemosyne.go

// Service implemented Mnemosyne interface
type Service struct {
	mnemosyne      repository.Mnemosyne
	reposRole      repository.Role
	reposUser      repository.User
	reposInterview repository.Interview
}

// NewService created Service struct
func NewService(mnemosyne repository.Mnemosyne, reposRole repository.Role, reposUser repository.User, reposInterview repository.Interview) *Service {
	return &Service{
		mnemosyne:      mnemosyne,
		reposRole:      reposRole,
		reposUser:      reposUser,
		reposInterview: reposInterview,
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

func (s *Service) AddUser(ctx context.Context, user *apiUser.User) (id *apiUser.Id, err error) {
	id, err = s.reposUser.AddUser(ctx, user)
	return
}

func (s *Service) GetUsers(ctx context.Context, ur *apiUser.UserRequest) (users *apiUser.Users, err error) {
	users, err = s.reposUser.GetUsers(ctx, ur)
	return
}

func (s *Service) GetUserById(ctx context.Context, id int) (user *apiUser.User, err error) {
	user, err = s.reposUser.GetUserById(ctx, id)
	return
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (user *apiUser.User, err error) {
	user, err = s.reposUser.GetUserByEmail(ctx, email)
	return
}
func (s *Service) UpdateUser(ctx context.Context, user *apiUser.User) (ok bool, err error) {
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

func (s *Service) GetContactById(ctx context.Context, id int) (c *apiUser.Contact, err error) {
	c, err = s.reposUser.GetContactById(ctx, id)
	return
}
func (s *Service) GetResumeById(ctx context.Context, id int) (r *apiUser.Resume, err error) {
	r, err = s.reposUser.GetResumeById(ctx, id)
	return
}
func (s *Service) UpdateContact(ctx context.Context, contact *apiUser.Contact) (ok bool, err error) {
	err = s.reposUser.UpdateContact(ctx, contact)
	if err == nil {
		ok = true
	}
	return
}
func (s *Service) UpdateResume(ctx context.Context, resume *apiUser.Resume) (ok bool, err error) {
	err = s.reposUser.UpdateResume(ctx, resume)
	if err == nil {
		ok = true
	}
	return
}
func (s *Service) GetEmailById(ctx context.Context, id int) (email *apiUser.Email, err error) {
	email, err = s.reposUser.GetEmailById(ctx, id)
	return
}
