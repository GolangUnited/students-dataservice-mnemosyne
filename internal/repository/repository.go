package repository

import (
	"context"

	"github.com/NEKETSKY/mnemosyne/internal/repository/interview"

	"github.com/NEKETSKY/mnemosyne/internal/repository/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/internal/repository/role"
	"github.com/NEKETSKY/mnemosyne/internal/repository/user"
	"github.com/NEKETSKY/mnemosyne/models/database"
	apiUser "github.com/NEKETSKY/mnemosyne/pkg/api/user"
	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository.go

// Mnemosyne represents test interface
type Mnemosyne interface {
	Test(ctx context.Context) error
}

type Role interface {
	GetAllRoles(ctx context.Context) (roles []database.Role, err error)
	GetUserRoles(ctx context.Context, userId int) (roles []database.Role, err error)
	DeleteUserRoleByCode(ctx context.Context, userId int, roleCode int) (err error)
	AddUserRoleByCode(ctx context.Context, userId int, roleCode int) (err error)
}

type User interface {
	AddUser(ctx context.Context, user *apiUser.User) (userId *apiUser.Id, err error)
	GetUsers(ctx context.Context, ur *apiUser.UserRequest) (users *apiUser.Users, err error)
	GetUserById(ctx context.Context, userId int) (user *apiUser.User, err error)
	GetUserByEmail(ctx context.Context, userEmail string) (user *apiUser.User, err error)
	UpdateUserById(ctx context.Context, user *apiUser.User) (err error)
	ActivateUserById(ctx context.Context, userId int) (err error)
	DeactivateUserById(ctx context.Context, userId int) (err error)
	GetContactById(ctx context.Context, id int) (c *apiUser.Contact, err error)
	GetResumeById(ctx context.Context, id int) (r *apiUser.Resume, err error)
	UpdateContact(ctx context.Context, contact *apiUser.Contact) (err error)
	UpdateResume(ctx context.Context, resume *apiUser.Resume) (err error)
	GetEmailById(ctx context.Context, id int) (email *apiUser.Email, err error)
}

type Interview interface {
	AddInterview(ctx context.Context, interview database.Interview) (interviewId int, err error)
	GetInterviews(ctx context.Context, interviewerId int, studentId int) (interviews []database.Interview, err error)
	GetInterviewById(ctx context.Context, interviewId int) (interview database.Interview, err error)
	UpdateInterviewById(ctx context.Context, interview database.Interview) (err error)
	DeactivateInterviewById(ctx context.Context, interviewId int) (err error)
	ActivateInterviewById(ctx context.Context, interviewId int) (err error)
}

type Repository struct {
	Mnemosyne
	Role
	User
	Interview
}

// NewRepository created Repository struct
func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Mnemosyne: mnemosyne.NewMnemosyne(db),
		Role:      role.NewRoleRepository(db),
		User:      user.NewUserRepository(db),
		Interview: interview.NewInterviewRepository(db),
	}
}
