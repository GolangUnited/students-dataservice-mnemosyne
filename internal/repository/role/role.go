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
	rows, err := r.db.Query(ctx, AllRolesQuery)
	if err != nil {
		return nil, err
	}
	result := make([]database.Role, 0)
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		role := database.Role{Id: int(values[0].(int32)), Name: values[1].(string)}
		result = append(result, role)
	}
	return result, err
}

func (r *RoleRepository) GetUserRoles(ctx context.Context, userId int) (roles []database.Role, err error) {
	rows, err := r.db.Query(ctx, RolesByUserIdQuery, userId)
	if err != nil {
		return nil, err
	}
	result := make([]database.Role, 0)
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		role := database.Role{Id: int(values[0].(int32)), Name: values[1].(string)}
		result = append(result, role)
	}
	return result, err
}

func (r *RoleRepository) DeleteUserRole(ctx context.Context, userId int, roleId int) (err error) {
	_, err = r.db.Query(ctx, DeleteRoleForUserQuery, userId, roleId)
	if err != nil {
		return err
	}
	return err
}
func (r *RoleRepository) AddUserRole(ctx context.Context, userId int, roleId int) (err error) {
	_, err = r.db.Query(ctx, AddRoleForUserQuery, userId, roleId)
	if err != nil {
		return err
	}
	return err
}
