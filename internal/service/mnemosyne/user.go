package mnemosyne

import (
	"context"

	dbUser "github.com/NEKETSKY/mnemosyne/models/database/user"
)

func (s *Service) AddUser(ctx context.Context, user *dbUser.UserFullStuff) (id int, err error) {
	id, err = s.reposUser.AddUser(ctx, user)
	return
}

func (s *Service) GetUsers(ctx context.Context, ur *dbUser.UserRequest) (users []dbUser.UserFullStuff, err error) {
	users, err = s.reposUser.GetUsers(ctx, ur)
	return
}

func (s *Service) GetUserById(ctx context.Context, id int) (user *dbUser.UserFullStuff, err error) {
	user, err = s.reposUser.GetUserById(ctx, id)
	return
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (user *dbUser.UserFullStuff, err error) {
	user, err = s.reposUser.GetUserByEmail(ctx, email)
	return
}
func (s *Service) UpdateUser(ctx context.Context, user *dbUser.UserFullStuff) (err error) {
	err = s.reposUser.UpdateUserById(ctx, user)
	return
}
func (s *Service) DeactivateUser(ctx context.Context, id int) (err error) {
	err = s.reposUser.DeactivateUserById(ctx, id)
	return
}

func (s *Service) ActivateUser(ctx context.Context, id int) (err error) {
	err = s.reposUser.ActivateUserById(ctx, id)

	return
}

func (s *Service) GetContactById(ctx context.Context, id int) (c *dbUser.Contact, err error) {
	c, err = s.reposUser.GetContactById(ctx, id)
	return
}
func (s *Service) GetResumeById(ctx context.Context, id int) (r *dbUser.Resume, err error) {
	r, err = s.reposUser.GetResumeById(ctx, id)
	return
}
func (s *Service) UpdateContact(ctx context.Context, contact *dbUser.Contact) (err error) {
	err = s.reposUser.UpdateContact(ctx, contact)

	return
}
func (s *Service) UpdateResume(ctx context.Context, resume *dbUser.Resume) (err error) {
	err = s.reposUser.UpdateResume(ctx, resume)
	return
}
