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
}

// NewService created Service struct
func NewService(mnemosyne repository.Mnemosyne, reposRole repository.Role) *Service {
	return &Service{
		mnemosyne: mnemosyne,
		reposRole: reposRole,
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
