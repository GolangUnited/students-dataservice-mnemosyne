package mnemosyne

import (
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
	reposProject     repository.Project
}

// NewService created Service struct
func NewService(
	reposRole repository.Role,
	reposUser repository.User,
	reposInterview repository.Interview,
	reposGroup repository.Group,
	reposCertificate repository.Certificate,
	reposTeam repository.Team,
	reposProject repository.Project,
) *Service {
	return &Service{
		reposRole:        reposRole,
		reposUser:        reposUser,
		reposInterview:   reposInterview,
		reposGroup:       reposGroup,
		reposCertificate: reposCertificate,
		reposTeam:        reposTeam,
		reposProject:     reposProject,
	}
}
