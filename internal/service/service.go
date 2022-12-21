package service

import (
	"context"

	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/repository"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/service/mnemosyne"
	"github.com/GolangUnited/students-dataservice-mnemosyne/models/database"
	modelGroup "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/group"
	modelRole "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/role"
	modelTeam "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/team"
	dbUser "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/user"
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

type Mnemosyne interface {
	Interview
	Certificate
	User
	Group
	Role
	Team
}

type Certificate interface {
	CreateCertificate(ctx context.Context, certificate database.Certificate) (certificateId uint32, err error)
	GetCertificates(ctx context.Context, userId uint32) (certificates []database.Certificate, err error)
	UpdateCertificate(ctx context.Context, certificate database.Certificate) (err error)
	DeactivateCertificate(ctx context.Context, certificateId uint32) (err error)
	ActivateCertificate(ctx context.Context, certificateId uint32) (err error)
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

type Role interface {
	GetUserRoles(ctx context.Context, userId int) ([]modelRole.DB, error)
	GetRoles(context.Context, *modelRole.Filter) ([]*modelRole.DB, error)
	CreateRole(context.Context, *modelRole.DB) (uint32, error)
	DeleteRole(context.Context, uint32) error
	AddUserToRole(ctx context.Context, userId, roleId uint32) error
	DeleteUserFromRole(ctx context.Context, userId, roleId uint32) error
}

type Team interface {
	GetTeam(context.Context, uint32) (*modelTeam.DB, error)
	GetTeams(context.Context, *modelTeam.Filter) ([]*modelTeam.DB, error)
	CreateTeam(context.Context, *modelTeam.DB) (uint32, error)
	UpdateTeam(context.Context, *modelTeam.DB) error
	DeactivateTeam(context.Context, uint32) error
	ActivateTeam(context.Context, uint32) error
	AddUserToTeam(ctx context.Context, userId, teamId uint32) error
	DeleteUserFromTeam(ctx context.Context, userId, teamId uint32) error
}

// Service represents service level
type Service struct {
	//suggest to move here interfaces User, Group etc. and to remove Mnemosyne interface
	Mnemosyne
}

// NewService created new service with repository
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Mnemosyne: mnemosyne.NewService(repos.Role, repos.User, repos.Interview, repos.Group, repos.Certificate, repos.Team),
	}
}
