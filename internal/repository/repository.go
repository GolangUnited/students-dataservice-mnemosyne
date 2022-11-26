package repository

import (
	"context"

	"github.com/NEKETSKY/mnemosyne/internal/repository/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/internal/repository/role"
	"github.com/NEKETSKY/mnemosyne/internal/repository/user"
	"github.com/NEKETSKY/mnemosyne/models/database"
	dbUser "github.com/NEKETSKY/mnemosyne/models/database/user"
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
	GetUserById(ctx context.Context, userId int) (user dbUser.BaseUser, err error)
	GetUserByEmail(ctx context.Context, userEmail string) (user dbUser.BaseUser, err error)
	UpdateUserById(ctx context.Context, user dbUser.BaseUser) (err error)
	ActivateUserById(ctx context.Context, userId int) (err error)
	DeactivateUserById(ctx context.Context, userId int) (err error)
}

type Repository struct {
	Mnemosyne
	Role
	User
}

// NewRepository created Repository struct
func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Mnemosyne: mnemosyne.NewMnemosyne(db),
		Role:      role.NewRoleRepository(db),
		User:      user.NewUserRepository(db),
	}
}
