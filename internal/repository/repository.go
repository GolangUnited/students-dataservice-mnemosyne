package repository

import (
	"context"

	"github.com/NEKETSKY/mnemosyne/internal/repository/group"
	"github.com/NEKETSKY/mnemosyne/internal/repository/interview"
	"github.com/NEKETSKY/mnemosyne/internal/repository/team"
	modelGroup "github.com/NEKETSKY/mnemosyne/models/database/group"
	modelRole "github.com/NEKETSKY/mnemosyne/models/database/role"
	modelTeam "github.com/NEKETSKY/mnemosyne/models/database/team"

	"github.com/NEKETSKY/mnemosyne/internal/repository/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/internal/repository/role"
	"github.com/NEKETSKY/mnemosyne/internal/repository/user"
	"github.com/NEKETSKY/mnemosyne/models/database"
	dbUser "github.com/NEKETSKY/mnemosyne/models/database/user"
	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository.go

// Mnemosyne represents test interface
type Mnemosyne interface {
	Test(ctx context.Context) error
}

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

type Repository struct {
	Mnemosyne
	Role
	User
	Interview
	Group
	Team
}

// NewRepository created Repository struct
func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Mnemosyne: mnemosyne.NewMnemosyne(db),
		Role:      role.NewRoleRepository(db),
		User:      user.NewUserRepository(db),
		Interview: interview.NewInterviewRepository(db),
		Group:     group.NewRepository(db),
		Team:      team.NewRepository(db),
	}
}
