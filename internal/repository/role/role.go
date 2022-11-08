package role

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/models/database"
	"github.com/jackc/pgx/v5"
)

type RoleRepository struct {
	db *pgx.Conn
}

func NewRoleRepository(db *pgx.Conn) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) GetAllRoles(ctx context.Context) (roles []database.Role, err error) {
	rows, _ := r.db.Query(ctx, AllRolesQuery)
	if err != nil {
		return nil, err
	}
	roles, err = pgx.CollectRows(rows, pgx.RowToStructByPos[database.Role])
	if err != nil {
		return nil, err
	}
	return roles, err
}

func (r *RoleRepository) GetUserRoles(ctx context.Context, userId int) (roles []database.Role, err error) {
	rows, err := r.db.Query(ctx, RolesByUserIdQuery, userId)
	if err != nil {
		return nil, err
	}
	roles, err = pgx.CollectRows(rows, pgx.RowToStructByPos[database.Role])
	if err != nil {
		return nil, err
	}
	return roles, err
}

func (r *RoleRepository) DeleteUserRoleByCode(ctx context.Context, userId int, roleCode int) (err error) {
	_, err = r.db.Query(ctx, DeleteRoleForUserQuery, userId, roleCode)
	if err != nil {
		return err
	}
	return err
}
func (r *RoleRepository) AddUserRoleByCode(ctx context.Context, userId int, roleCode int) (err error) {
	_, err = r.db.Query(ctx, AddRoleForUserQuery, userId, roleCode)
	if err != nil {
		return err
	}
	return err
}
