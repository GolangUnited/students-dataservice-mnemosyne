package project

import (
	"context"
	"github.com/GolangUnited/students-dataservice-mnemosyne/models/database/project"
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

func (r *Repository) GetProjectById(ctx context.Context, projectId uint32) (*project.DB, error) {
	rows, err := r.db.Query(ctx, GetProjectByIdQuery, projectId)
	if err != nil {
		return nil, errors.Wrap(err, "GetProjectById query error")
	}

	prDB, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[project.DB])
	if err != nil {
		return nil, errors.Wrap(err, "GetProjectById CollectOneRow error")
	}

	return &prDB, nil
}

func (r *Repository) GetProjects(ctx context.Context, projectFilter *project.Filter) ([]*project.DB, error) {
	sb := sqlbuilder.Select("grs.*").From("projects grs")
	if projectFilter.TeamId > 0 {
		sb.Where(
			sb.Equal(
				"team_id",
				projectFilter.TeamId,
			),
		)
	}
	if projectFilter.Name != "" {
		sb.Where(
			sb.Like(
				"name",
				"%"+projectFilter.Name+"%",
			),
		)
	}
	sb.Where(sb.Equal("deleted", projectFilter.Deleted))
	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetProjects query error")
	}
	prsDB, err := pgx.CollectRows(rows, pgx.RowToStructByName[project.DB])
	if err != nil {
		return nil, errors.Wrap(err, "GetProjects - unable to collect rows ")
	}

	projectsDB := make([]*project.DB, 0, len(prsDB))
	for i := 0; i < len(prsDB); i++ {
		projectsDB = append(projectsDB, &prsDB[i])
	}

	return projectsDB, err
}

func (r *Repository) AddProject(ctx context.Context, projectDB *project.DB) (projectId uint32, err error) {
	row := r.db.QueryRow(ctx, AddProjectQuery, projectDB.Name, projectDB.Description, projectDB.GitUrl, projectDB.TeamId)
	err = row.Scan(&projectId)
	if err != nil {
		return 0, errors.Wrap(err, "AddProject error while query executing")
	}

	return
}

func (r *Repository) UpdateProject(ctx context.Context, projectDB *project.DB) error {
	_, err := r.db.Exec(ctx, UpdateProjectByIdQuery,
		projectDB.Name, projectDB.Description, projectDB.GitUrl, projectDB.TeamId, time.Now(), projectDB.Id,
	)
	if err != nil {
		return errors.Wrap(err, "UpdateProject - unable to execute update statement")
	}

	return err
}

func (r *Repository) DeactivateProject(ctx context.Context, projectId uint32) error {
	_, err := r.db.Exec(ctx, DeactivateProjectByIdQuery, time.Now(), projectId)
	if err != nil {
		return errors.Wrapf(err, "DeactivateProject - unable to set project %d as deleted", projectId)
	}

	return err
}

func (r *Repository) ActivateProject(ctx context.Context, projectId uint32) error {
	_, err := r.db.Exec(ctx, ActivateProjectByIdQuery, time.Now(), projectId)
	if err != nil {
		return errors.Wrapf(err, "ActivateProject - unable to set project %d as active", projectId)
	}

	return err
}
