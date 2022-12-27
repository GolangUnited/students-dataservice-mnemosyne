package mnemosyne

import (
	"context"

	modelRole "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/role"

	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/repository"
)

//go:generate mockgen -source=mnemosyne.go -destination=mocks/mnemosyne.go

// Service implemented Mnemosyne interface
type Service struct {
	reposRole        repository.Role
	reposUser        repository.User
	reposInterview   repository.Interview
	reposGroup       repository.Group
	reposCertificate repository.Certificate
	reposTeam        repository.Team
	reposLesson      repository.Lesson
}

// NewService created Service struct
func NewService(
	reposRole repository.Role,
	reposUser repository.User,
	reposInterview repository.Interview,
	reposGroup repository.Group,
	reposCertificate repository.Certificate,
	reposTeam repository.Team,
	reposLesson repository.Lesson,
) *Service {
	return &Service{
		reposRole:        reposRole,
		reposUser:        reposUser,
		reposInterview:   reposInterview,
		reposGroup:       reposGroup,
		reposCertificate: reposCertificate,
		reposTeam:        reposTeam,
		reposLesson:      reposLesson,
	}
}

// GetUserRoles get all user roles
func (s *Service) GetUserRoles(ctx context.Context, userId int) (userRoles []modelRole.DB, err error) {
	userRoles, err = s.reposRole.GetUserRoles(ctx, userId)
	if err != nil {
		return nil, err
	}
	return
}
