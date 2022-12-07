package service

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/internal/service/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/models/database"
	modelGroup "github.com/NEKETSKY/mnemosyne/models/database/group"
	model "github.com/NEKETSKY/mnemosyne/models/mnemosyne"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

type Interview interface {
	CreateInterview(ctx context.Context, interviewModel database.Interview) (interviewDb database.Interview, err error)
	GetInterviews(ctx context.Context, interviewerId uint, studentId uint) (interviews []database.Interview, err error)
	GetInterviewById(ctx context.Context, interviewId uint) (interview database.Interview, err error)
	UpdateInterview(ctx context.Context, interview database.Interview) (interviewDb database.Interview, err error)
	DeactivateInterview(ctx context.Context, interviewId uint) (interview database.Interview, err error)
	ActivateInterview(ctx context.Context, interviewId uint) (interview database.Interview, err error)
}

// Mnemosyne has test signatures
type Mnemosyne interface {
	Interview
	Test(context.Context, model.Request) (model.Response, error)
	GetUserRoles(ctx context.Context, userId int) ([]database.Role, error)
	Group
}

type Group interface {
	GetGroup(context.Context, uint32) (*modelGroup.DB, error)
	GetGroups(context.Context, *modelGroup.Filter) ([]*modelGroup.DB, error)
	CreateGroup(context.Context, *modelGroup.DB) (uint32, error)
	UpdateGroup(context.Context, *modelGroup.DB) error
	DeactivateGroup(context.Context, uint32) error
	ActivateGroup(context.Context, uint32) error
	AddUserToGroup(ctx context.Context, userId, groupId uint32) error
	DeleteUserFromGroup(ctx context.Context, userId, groupId uint32) error
}

// Service represents service level
type Service struct {
	Mnemosyne
}

// NewService created new service with repository
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Mnemosyne: mnemosyne.NewService(repos.Mnemosyne, repos.Role, repos.User, repos.Interview, repos.Group),
	}
}
