package service

import (
	"context"

	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/internal/service/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/models/database"
	modelGroup "github.com/NEKETSKY/mnemosyne/models/database/group"
	dbUser "github.com/NEKETSKY/mnemosyne/models/database/user"
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
	User
	Group
}

type User interface {
	AddUser(ctx context.Context, user *dbUser.TransitUser) (id int, err error)
	GetUsers(ctx context.Context, ur *dbUser.UserRequest) (users []dbUser.UserFullStuff, err error)
	GetUserById(ctx context.Context, id int) (user *dbUser.UserFullStuff, err error)
	GetUserByEmail(ctx context.Context, email string) (user *dbUser.UserFullStuff, err error)
	UpdateUser(ctx context.Context, user *dbUser.TransitUser) (err error)
	DeactivateUser(ctx context.Context, id int) (err error)
	ActivateUser(ctx context.Context, id int) (err error)
	GetContactById(ctx context.Context, id int) (c *dbUser.Contact, err error)
	GetResumeById(ctx context.Context, id int) (r *dbUser.Resume, err error)
	UpdateContact(ctx context.Context, contact *dbUser.Contact) (err error)
	UpdateResume(ctx context.Context, resume *dbUser.TransitResume) (err error)
	DeleteContact(ctx context.Context, id int) (err error)
	DeleteResume(ctx context.Context, id int) (err error)
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
