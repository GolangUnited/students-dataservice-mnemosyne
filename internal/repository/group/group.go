package group

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/models/database/group"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"time"
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetGroupById(ctx context.Context, groupId uint32) (*group.DB, error) {
	rows, err := r.db.Query(ctx, GetGroupByIdQuery, groupId)
	if err != nil {
		return nil, errors.Wrap(err, "GetGroupById query error")
	}

	grDB, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[group.DB])
	if err != nil {
		return nil, errors.Wrap(err, "GetGroupById CollectOneRow error")
	}

	return &grDB, nil
}

func (r *Repository) GetGroups(ctx context.Context, groupFilter *group.Filter) ([]*group.DB, error) {
	sb := sqlbuilder.Select("grs.*").From("groups grs")
	if groupFilter.UserId > 0 {
		sb.Join("user_groups sgs",
			"grs.id = sgs.group_id",
			sb.Equal("sgs.user_id", groupFilter.UserId),
		)
	}
	if groupFilter.Name != "" {
		sb.Where(
			sb.Like(
				"name",
				"%"+groupFilter.Name+"%",
			),
		)
	}
	if groupFilter.StartDate.Unix() != 0 {
		sb.Where(sb.GreaterEqualThan(
			"start_date",
			groupFilter.StartDate,
		))
	}
	if groupFilter.EndDate.Unix() != 0 {
		sb.Where(sb.LessEqualThan(
			"end_date",
			groupFilter.EndDate,
		))
	}
	sb.Where(sb.Equal("deleted", groupFilter.Deleted))
	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetGroups query error")
	}
	grsDB, err := pgx.CollectRows(rows, pgx.RowToStructByName[group.DB])
	if err != nil {
		return nil, errors.Wrap(err, "GetGroups - unable to collect rows ")
	}

	groupsDB := make([]*group.DB, 0, len(grsDB))
	for i := 0; i < len(grsDB); i++ {
		groupsDB = append(groupsDB, &grsDB[i])
	}

	return groupsDB, err
}

func (r *Repository) AddGroup(ctx context.Context, groupDB *group.DB) (groupId uint32, err error) {
	row := r.db.QueryRow(ctx, AddGroupQuery, groupDB.Name, groupDB.StartDate, groupDB.EndDate)
	err = row.Scan(&groupId)
	if err != nil {
		return 0, errors.Wrap(err, "AddGroup error while query executing")
	}

	return
}

func (r *Repository) UpdateGroup(ctx context.Context, groupDB *group.DB) error {
	_, err := r.db.Exec(ctx, UpdateGroupByIdQuery,
		groupDB.Name, groupDB.StartDate, groupDB.EndDate, time.Now(), groupDB.Id,
	)
	if err != nil {
		return errors.Wrap(err, "UpdateGroup - unable to execute update statement")
	}

	return err
}

func (r *Repository) DeactivateGroup(ctx context.Context, groupId uint32) error {
	_, err := r.db.Exec(ctx, DeactivateGroupByIdQuery, time.Now(), groupId)
	if err != nil {
		return errors.Wrapf(err, "DeactivateGroup - unable to set group %d as deleted", groupId)
	}

	return err
}

func (r *Repository) ActivateGroup(ctx context.Context, groupId uint32) error {
	_, err := r.db.Exec(ctx, ActivateGroupByIdQuery, time.Now(), groupId)
	if err != nil {
		return errors.Wrapf(err, "ActivateGroup - unable to set group %d as active", groupId)
	}

	return err
}

func (r *Repository) AddUserToGroup(ctx context.Context, userId, groupId uint32) error {
	_, err := r.db.Exec(ctx, AddUserToGroupQuery, userId, groupId)
	if err != nil {
		return errors.Wrap(err, "AddUserToGroup - unable to execute")
	}

	return err
}

func (r *Repository) DeleteUserFromGroup(ctx context.Context, userId, groupId uint32) error {
	_, err := r.db.Exec(ctx, DeleteUserFromGroupQuery, userId, groupId)
	if err != nil {
		return errors.Wrap(err, "DeleteUserFromGroup - unable to execute")
	}

	return err
}
