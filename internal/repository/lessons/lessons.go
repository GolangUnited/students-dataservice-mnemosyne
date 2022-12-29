package lessons

import (
	"context"
	"github.com/GolangUnited/students-dataservice-mnemosyne/models/database/lessons"
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

func (r *Repository) GetLessonById(ctx context.Context, lessonId uint32) (*lessons.Lessons, error) {
	rows, err := r.db.Query(ctx, GetLessonByIdQuery, lessonId)
	if err != nil {
		return nil, errors.Wrap(err, "GetLessonById query error")
	}

	lessonDB, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[lessons.Lessons])
	if err != nil {
		return nil, errors.Wrap(err, "GetLessonById CollectOneRow error")
	}

	return &lessonDB, nil
}

func (r *Repository) GetLessons(ctx context.Context, lessonFilter *lessons.Filter) ([]*lessons.Lessons, error) {
	sb := sqlbuilder.Select("less.*").From("lessons less")

	if lessonFilter.Presentation != "" {
		sb.Where(
			sb.Like(
				"presentation",
				"%"+lessonFilter.Presentation+"%",
			),
		)
	}
	if lessonFilter.VideoLink != "" {
		sb.Where(
			sb.Like(
				"video_link",
				"%"+lessonFilter.VideoLink+"%",
			),
		)
	}
	if lessonFilter.LessonDate.Unix() != 0 {
		sb.Where(sb.GreaterEqualThan(
			"lesson_date",
			lessonFilter.LessonDate,
		))
	}
	if lessonFilter.Homework != "" {
		sb.Where(
			sb.Like(
				"homework",
				"%"+lessonFilter.Homework+"%",
			),
		)
	}
	if lessonFilter.LecturerId > 0 {
		sb.Where(
			sb.Equal(
				"lecturer_id",
				lessonFilter.LecturerId,
			),
		)
	}
	if lessonFilter.GroupId > 0 {
		sb.Where(
			sb.Equal(
				"group_id",
				lessonFilter.GroupId,
			),
		)
	}
	if lessonFilter.Language != "" {
		sb.Where(
			sb.Like(
				"language",
				"%"+lessonFilter.Language+"%",
			),
		)
	}

	sb.Where(sb.Equal("deleted", lessonFilter.Deleted))
	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "GetLessons query error")
	}
	lessDB, err := pgx.CollectRows(rows, pgx.RowToStructByName[lessons.Lessons])
	if err != nil {
		return nil, errors.Wrap(err, "GetLessons - unable to collect rows ")
	}

	lessonsDB := make([]*lessons.Lessons, 0, len(lessDB))
	for i := 0; i < len(lessDB); i++ {
		lessonsDB = append(lessonsDB, &lessDB[i])
	}

	return lessonsDB, err
}

func (r *Repository) AddLesson(ctx context.Context, lessonDB *lessons.Lessons) (lessonId uint32, err error) {
	row := r.db.QueryRow(
		ctx, AddLessonQuery,
		lessonDB.Presentation,
		lessonDB.VideoLink,
		lessonDB.LessonDate,
		lessonDB.Homework,
		lessonDB.LecturerId,
		lessonDB.GroupId,
		lessonDB.Language,
	)
	err = row.Scan(&lessonId)
	if err != nil {
		return 0, errors.Wrap(err, "AddLesson error while query executing")
	}

	return
}

func (r *Repository) UpdateLesson(ctx context.Context, lessonDB *lessons.Lessons) error {
	_, err := r.db.Exec(
		ctx, UpdateLessonByIdQuery,
		lessonDB.Presentation,
		lessonDB.VideoLink,
		lessonDB.LessonDate,
		lessonDB.Homework,
		lessonDB.LecturerId,
		lessonDB.GroupId,
		lessonDB.Language,
		time.Now(),
		lessonDB.Id,
	)
	if err != nil {
		return errors.Wrap(err, "UpdateLesson - unable to execute update statement")
	}

	return err
}

func (r *Repository) DeactivateLesson(ctx context.Context, lessonId uint32) error {
	_, err := r.db.Exec(ctx, DeactivateLessonByIdQuery, time.Now(), lessonId)
	if err != nil {
		return errors.Wrapf(err, "DeactivateLesson - unable to set lesson %d as deleted", lessonId)
	}

	return err
}

func (r *Repository) ActivateLesson(ctx context.Context, lessonId uint32) error {
	_, err := r.db.Exec(ctx, ActivateLessonByIdQuery, time.Now(), lessonId)
	if err != nil {
		return errors.Wrapf(err, "ActivateLesson - unable to set lesson %d as active", lessonId)
	}

	return err
}
