package repository

import (
	"context"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/repository/certificate"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/repository/group"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/repository/interview"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/repository/lessons"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/repository/project"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/repository/role"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/repository/team"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/repository/user"
	"github.com/GolangUnited/students-dataservice-mnemosyne/models/database"
	modelGroup "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/group"
	modelLesson "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/lessons"
	modelProject "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/project"
	modelRole "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/role"
	modelTeam "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/team"
	dbUser "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/user"
	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository.go

type Role interface {
	GetAllRoles(ctx context.Context) (roles []modelRole.DB, err error)
	GetUserRoles(ctx context.Context, userId int) (roles []modelRole.DB, err error)
	DeleteUserRoleByCode(ctx context.Context, userId int, roleCode int) (err error)
	AddUserRoleByCode(ctx context.Context, userId int, roleCode int) (err error)

	GetRoles(context.Context, *modelRole.Filter) ([]*modelRole.DB, error)
	AddRole(context.Context, *modelRole.DB) (uint32, error)
	DeleteRole(context.Context, uint32) error
	AddUserToRole(ctx context.Context, userId, roleId uint32) error
	DeleteUserFromRole(ctx context.Context, userId, roleId uint32) error
}

type User interface {
	AddUser(ctx context.Context, user *dbUser.UserFullStuff) (userId int, err error)
	GetUsers(ctx context.Context, ur *dbUser.UserRequest) (users []dbUser.UserFullStuff, err error)
	GetUserById(ctx context.Context, userId int) (user *dbUser.UserFullStuff, err error)
	GetUserByEmail(ctx context.Context, userEmail string) (user *dbUser.UserFullStuff, err error)
	UpdateUserById(ctx context.Context, user *dbUser.UserFullStuff) (err error)
	ActivateUserById(ctx context.Context, userId int) (err error)
	DeactivateUserById(ctx context.Context, userId int) (err error)
	GetContactById(ctx context.Context, id int) (c *dbUser.Contact, err error)
	GetResumeById(ctx context.Context, id int) (r *dbUser.Resume, err error)
	UpdateContact(ctx context.Context, contact *dbUser.Contact) (err error)
	UpdateResume(ctx context.Context, resume *dbUser.Resume) (err error)
	DeleteContact(ctx context.Context, id int) (err error)
	DeleteResume(ctx context.Context, id int) (err error)
}

type Interview interface {
	AddInterview(ctx context.Context, interview database.Interview) (interviewId int, err error)
	GetInterviews(ctx context.Context, interviewerId uint, studentId uint) (interviews []database.Interview, err error)
	GetInterviewById(ctx context.Context, interviewId uint) (interview database.Interview, err error)
	UpdateInterviewById(ctx context.Context, interview database.Interview) (err error)
	DeactivateInterviewById(ctx context.Context, interviewId uint) (err error)
	ActivateInterviewById(ctx context.Context, interviewId uint) (err error)
}

type Group interface {
	GetGroupById(context.Context, uint32) (*modelGroup.DB, error)
	GetGroups(context.Context, *modelGroup.Filter) ([]*modelGroup.DB, error)
	AddGroup(context.Context, *modelGroup.DB) (uint32, error)
	UpdateGroup(context.Context, *modelGroup.DB) error
	DeactivateGroup(context.Context, uint32) error
	ActivateGroup(context.Context, uint32) error
	AddUserToGroup(ctx context.Context, userId, groupId uint32) error
	DeleteUserFromGroup(ctx context.Context, userId, groupId uint32) error
}

type Certificate interface {
	CreateCertificate(ctx context.Context, certificate database.Certificate) (certificateId uint32, err error)
	GetCertificateById(ctx context.Context, certificateId uint32) (certificate database.Certificate, err error)
	GetCertificates(ctx context.Context, userId uint32) (certificates []database.Certificate, err error)
	UpdateCertificates(ctx context.Context, certificate database.Certificate) (err error)
	DeactivateCertificate(ctx context.Context, certificateId uint32) (err error)
	ActivateCertificate(ctx context.Context, certificateId uint32) (err error)
}

type Team interface {
	GetTeamById(context.Context, uint32) (*modelTeam.DB, error)
	GetTeams(context.Context, *modelTeam.Filter) ([]*modelTeam.DB, error)
	AddTeam(context.Context, *modelTeam.DB) (uint32, error)
	UpdateTeam(context.Context, *modelTeam.DB) error
	DeactivateTeam(context.Context, uint32) error
	ActivateTeam(context.Context, uint32) error
	AddUserToTeam(ctx context.Context, userId, teamId uint32) error
	DeleteUserFromTeam(ctx context.Context, userId, teamId uint32) error
}

type Lesson interface {
	GetLessonById(context.Context, uint32) (*modelLesson.Lessons, error)
	GetLessons(context.Context, *modelLesson.Filter) ([]*modelLesson.Lessons, error)
	AddLesson(context.Context, *modelLesson.Lessons) (uint32, error)
	UpdateLesson(context.Context, *modelLesson.Lessons) error
	DeactivateLesson(context.Context, uint32) error
	ActivateLesson(context.Context, uint32) error
}

type Project interface {
	GetProjectById(context.Context, uint32) (*modelProject.DB, error)
	GetProjects(context.Context, *modelProject.Filter) ([]*modelProject.DB, error)
	AddProject(context.Context, *modelProject.DB) (uint32, error)
	UpdateProject(context.Context, *modelProject.DB) error
	DeactivateProject(context.Context, uint32) error
	ActivateProject(context.Context, uint32) error
}

type Repository struct {
	Role
	User
	Interview
	Certificate
	Group
	Team
	Lesson
	Project
}

// NewRepository created Repository struct
func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Role:        role.NewRoleRepository(db),
		User:        user.NewUserRepository(db),
		Interview:   interview.NewInterviewRepository(db),
		Certificate: certificate.NewCertificateRepository(db),
		Group:       group.NewRepository(db),
		Team:        team.NewRepository(db),
		Lesson:      lessons.NewRepository(db),
		Project:     project.NewRepository(db),
	}
}
