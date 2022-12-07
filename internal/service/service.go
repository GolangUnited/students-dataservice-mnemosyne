package service

import (
	"context"

	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/internal/service/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/models/database"
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
