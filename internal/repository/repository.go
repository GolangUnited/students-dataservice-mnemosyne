package repository

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/internal/repository/certificate"
	"github.com/NEKETSKY/mnemosyne/internal/repository/interview"

	"github.com/NEKETSKY/mnemosyne/internal/repository/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/internal/repository/role"
	"github.com/NEKETSKY/mnemosyne/internal/repository/user"
	"github.com/NEKETSKY/mnemosyne/models/database"
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
	AddUser(ctx context.Context, user database.User) (userId int, err error)
	GetAllUsers(ctx context.Context) (users []database.User, err error)
	GetUserById(ctx context.Context, userId int) (user database.User, err error)
	GetUserByEmail(ctx context.Context, userEmail string) (user database.User, err error)
	UpdateUserById(ctx context.Context, user database.User) (err error)
	ActivateUserById(ctx context.Context, userId int) (err error)
	DeactivateUserById(ctx context.Context, userId int) (err error)
}

type Interview interface {
	AddInterview(ctx context.Context, interview database.Interview) (interviewId int, err error)
	GetInterviews(ctx context.Context, interviewerId int, studentId int) (interviews []database.Interview, err error)
	GetInterviewById(ctx context.Context, interviewId int) (interview database.Interview, err error)
	UpdateInterviewById(ctx context.Context, interview database.Interview) (err error)
	DeactivateInterviewById(ctx context.Context, interviewId int) (err error)
	ActivateInterviewById(ctx context.Context, interviewId int) (err error)
}

type Certificate interface {
	AddCertificate(ctx context.Context, certificate database.Certificate) (certificateId int, err error)
	GetCertificates(ctx context.Context) (certificates []database.Certificate, err error)
	UpdateCertificatesById(ctx context.Context, certificate database.Certificate) (err error)
	DeactivateCertificateById(ctx context.Context, certificateId int) (err error)
	ActivateCertificateById(ctx context.Context, certificateId int) (err error)
}

type Repository struct {
	Mnemosyne
	Role
	User
	Interview
	Certificate
}

// NewRepository created Repository struct
func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Mnemosyne:   mnemosyne.NewMnemosyne(db),
		Role:        role.NewRoleRepository(db),
		User:        user.NewUserRepository(db),
		Interview:   interview.NewInterviewRepository(db),
		Certificate: certificate.NewCertificateRepository(db),
	}
}
