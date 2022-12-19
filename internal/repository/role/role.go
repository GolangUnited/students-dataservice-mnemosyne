package role

import (
	"context"
	modelRole "github.com/NEKETSKY/mnemosyne/models/database/role"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type RoleRepository struct {
	db *pgx.Conn
}

func NewRoleRepository(db *pgx.Conn) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) GetAllRoles(ctx context.Context) (roles []modelRole.DB, err error) {
	rows, _ := r.db.Query(ctx, AllRolesQuery)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllRoles query error")
	}
	roles, err = pgx.CollectRows(rows, pgx.RowToStructByPos[modelRole.DB])
	if err != nil {
		return nil, errors.Wrap(err, "GetAllRoles CollectRows error")
	}
	return roles, err
}

func (r *RoleRepository) GetUserRoles(ctx context.Context, userId int) (roles []modelRole.DB, err error) {
	rows, err := r.db.Query(ctx, RolesByUserIdQuery, userId)
	if err != nil {
		return nil, errors.Wrap(err, "GetUserRoles Query error")
	}
	roles, err = pgx.CollectRows(rows, pgx.RowToStructByPos[modelRole.DB])
	if err != nil {
		return nil, errors.Wrap(err, "GetUserRoles CollectRows error")
	}
	return roles, err
}

func (r *RoleRepository) DeleteUserRoleByCode(ctx context.Context, userId int, roleCode int) (err error) {
	_, err = r.db.Query(ctx, DeleteRoleForUserQuery, userId, roleCode)
	if err != nil {
		return errors.Wrap(err, "DeleteUserRoleByCode Query error")
	}
	return err
}
func (r *RoleRepository) AddUserRoleByCode(ctx context.Context, userId int, roleCode int) (err error) {
	_, err = r.db.Query(ctx, AddRoleForUserQuery, userId, roleCode)
	if err != nil {
		return errors.Wrap(err, "AddUserRoleByCode Query error")
	}
	return err
}

func (r *RoleRepository) GetRoles(ctx context.Context, roleFilter *modelRole.Filter) ([]*modelRole.DB, error) {
	sb := sqlbuilder.Select("rs.*").From("roles rs")
	if roleFilter.Code != "" {
		sb.Where(
			sb.Equal(
				"code",
				roleFilter.Code,
			),
		)
	}
	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetRoles query error")
	}
	rsDB, err := pgx.CollectRows(rows, pgx.RowToStructByName[modelRole.DB])
	if err != nil {
		return nil, errors.Wrap(err, "GetRoles - unable to collect rows ")
	}

	rolesDB := make([]*modelRole.DB, 0, len(rsDB))
	for i := 0; i < len(rsDB); i++ {
		rolesDB = append(rolesDB, &rsDB[i])
	}

	return rolesDB, err
}
func (r *RoleRepository) AddRole(ctx context.Context, roleDB *modelRole.DB) (roleId uint32, err error) {
	row := r.db.QueryRow(ctx, AddRoleQuery, roleDB.Code)
	err = row.Scan(&roleId)
	if err != nil {
		return 0, errors.Wrap(err, "AddRole error while query executing")
	}

	return
}
func (r *RoleRepository) DeleteRole(ctx context.Context, roleId uint32) error {
	_, err := r.db.Exec(ctx, DeleteByIdQuery, roleId)
	if err != nil {
		return errors.Wrapf(err, "DeleteRole - unable deleted")
	}

	return err
}
func (r *RoleRepository) AddUserToRole(ctx context.Context, userId, roleId uint32) error {
	_, err := r.db.Exec(ctx, AddUserToRoleQuery, userId, roleId)
	if err != nil {
		return errors.Wrap(err, "AddUserToRole - unable to execute")
	}

	return err
}
func (r *RoleRepository) DeleteUserFromRole(ctx context.Context, userId, roleId uint32) error {
	_, err := r.db.Exec(ctx, DeleteUserFromRoleQuery, userId, roleId)
	if err != nil {
		return errors.Wrap(err, "DeleteUserFromRole - unable to execute")
	}

	return err
}
