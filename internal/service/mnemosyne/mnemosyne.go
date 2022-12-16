package mnemosyne

import (
	"context"

	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/models/database"
)

//go:generate mockgen -source=mnemosyne.go -destination=mocks/mnemosyne.go

// Service implemented Mnemosyne interface
type Service struct {
	reposRole      repository.Role
	reposUser      repository.User
	reposInterview repository.Interview
	reposGroup     repository.Group
}

// NewService created Service struct
func NewService(reposRole repository.Role, reposUser repository.User,
	reposInterview repository.Interview, reposGroup repository.Group) *Service {
	return &Service{
		reposRole:      reposRole,
		reposUser:      reposUser,
		reposInterview: reposInterview,
		reposGroup:     reposGroup,
	}
}

// GetUserRoles get all user roles
func (s *Service) GetUserRoles(ctx context.Context, userId int) (userRoles []database.Role, err error) {
	userRoles, err = s.reposRole.GetUserRoles(ctx, userId)
	if err != nil {
		return nil, err
	}
	return
}
