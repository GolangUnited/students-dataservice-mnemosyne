package team

import (
	"context"
	"github.com/GolangUnited/students-dataservice-mnemosyne/models/database/team"
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

func (r *Repository) GetTeamById(ctx context.Context, teamId uint32) (*team.DB, error) {
	rows, err := r.db.Query(ctx, GetTeamByIdQuery, teamId)
	if err != nil {
		return nil, errors.Wrap(err, "GetTeamById query error")
	}

	teamDB, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[team.DB])
	if err != nil {
		return nil, errors.Wrap(err, "GetTeamById CollectOneRow error")
	}

	return &teamDB, nil
}

func (r *Repository) GetTeams(ctx context.Context, teamFilter *team.Filter) ([]*team.DB, error) {
	sb := sqlbuilder.Select("grs.*").From("teams grs")
	if teamFilter.UserId > 0 {
		sb.Join("user_teams ut",
			"grs.id = ut.team_id",
			sb.Equal("ut.user_id", teamFilter.UserId),
		)
	}
	if teamFilter.Name != "" {
		sb.Where(
			sb.Like(
				"name",
				"%"+teamFilter.Name+"%",
			),
		)
	}
	if teamFilter.MentorId > 0 {
		sb.Where(
			sb.Equal(
				"mentor_id",
				teamFilter.MentorId,
			),
		)
	}
	if teamFilter.StartDate.Unix() != 0 {
		sb.Where(sb.GreaterEqualThan(
			"start_date",
			teamFilter.StartDate,
		))
	}
	if teamFilter.EndDate.Unix() != 0 {
		sb.Where(sb.LessEqualThan(
			"end_date",
			teamFilter.EndDate,
		))
	}
	sb.Where(sb.Equal("deleted", teamFilter.Deleted))
	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetTeams query error")
	}
	grsDB, err := pgx.CollectRows(rows, pgx.RowToStructByName[team.DB])
	if err != nil {
		return nil, errors.Wrap(err, "GetTeams - unable to collect rows ")
	}

	teamsDB := make([]*team.DB, 0, len(grsDB))
	for i := 0; i < len(grsDB); i++ {
		teamsDB = append(teamsDB, &grsDB[i])
	}

	return teamsDB, err
}

func (r *Repository) AddTeam(ctx context.Context, teamDB *team.DB) (teamId uint32, err error) {
	row := r.db.QueryRow(
		ctx, AddTeamQuery, teamDB.Name, teamDB.StartDate, teamDB.EndDate, teamDB.GroupId, teamDB.MentorId,
	)
	err = row.Scan(&teamId)
	if err != nil {
		return 0, errors.Wrap(err, "AddTeam error while query executing")
	}

	return
}

func (r *Repository) UpdateTeam(ctx context.Context, teamDB *team.DB) error {
	_, err := r.db.Exec(ctx, UpdateTeamByIdQuery,
		teamDB.Name, teamDB.StartDate, teamDB.EndDate, teamDB.GroupId, teamDB.MentorId, time.Now(), teamDB.Id,
	)
	if err != nil {
		return errors.Wrap(err, "UpdateTeam - unable to execute update statement")
	}

	return err
}

func (r *Repository) DeactivateTeam(ctx context.Context, teamId uint32) error {
	_, err := r.db.Exec(ctx, DeactivateTeamByIdQuery, time.Now(), teamId)
	if err != nil {
		return errors.Wrapf(err, "DeactivateTeam - unable to set team %d as deleted", teamId)
	}

	return err
}

func (r *Repository) ActivateTeam(ctx context.Context, teamId uint32) error {
	_, err := r.db.Exec(ctx, ActivateTeamByIdQuery, time.Now(), teamId)
	if err != nil {
		return errors.Wrapf(err, "ActivateTeam - unable to set team %d as active", teamId)
	}

	return err
}

func (r *Repository) AddUserToTeam(ctx context.Context, userId, teamId uint32) error {
	_, err := r.db.Exec(ctx, AddUserToTeamQuery, userId, teamId)
	if err != nil {
		return errors.Wrap(err, "AddUserToTeam - unable to execute")
	}

	return err
}

func (r *Repository) DeleteUserFromTeam(ctx context.Context, userId, teamId uint32) error {
	_, err := r.db.Exec(ctx, DeleteUserFromTeamQuery, userId, teamId)
	if err != nil {
		return errors.Wrap(err, "DeleteUserFromTeam - unable to execute")
	}

	return err
}
