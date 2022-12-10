package interview

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/models/database"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"time"
)

type InterviewRepository struct {
	db *pgx.Conn
}

func NewInterviewRepository(db *pgx.Conn) *InterviewRepository {
	return &InterviewRepository{
		db: db,
	}
}

func (i *InterviewRepository) AddInterview(ctx context.Context, interview database.Interview) (interviewId int, err error) {
	row := i.db.QueryRow(
		ctx,
		AddInterviewQuery,
		interview.InterviewerId,
		interview.StudentId,
		interview.InterviewDate,
		interview.Grade,
		interview.SubjectiveRating,
		interview.Notes,
		interview.DeterminedEnglishLevel,
		interview.MainTask,
		interview.Question,
	)
	if err != nil {
		return 0, errors.Wrap(err, "AddInterview query error")
	}
	err = row.Scan(&interviewId)
	if err != nil {
		return 0, errors.Wrap(err, "AddInterview error while query executing")
	}
	return
}

func (i *InterviewRepository) GetInterviews(ctx context.Context, interviewerId uint, studentId uint) (interviews []database.Interview, err error) {
	sb := sqlbuilder.Select("*").From("interview")
	if interviewerId > 0 {
		sb.Where("interviewer_id = 1")
	}
	if studentId > 0 {
		sb.Where("student_id = 1")
	}
	query := sb.String()

	rows, _ := i.db.Query(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "GetAllInterviews query error")
	}
	interviews, err = pgx.CollectRows(rows, pgx.RowToStructByName[database.Interview])
	if err != nil {
		return nil, errors.Wrap(err, "GetAllInterviews - unable to collect rows ")
	}
	return interviews, err
}

func (i *InterviewRepository) GetInterviewById(ctx context.Context, interviewId uint) (interview database.Interview, err error) {

	rows, err := i.db.Query(ctx, GetInterviewByIdQuery, interviewId)
	if err != nil {
		return database.Interview{}, errors.Wrap(err, "GetInterviewById query error")
	}
	interview, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[database.Interview])
	if err != nil {
		return database.Interview{}, errors.Wrap(err, "GetInterviewById CollectOneRow error")
	}
	return interview, err
}

func (i *InterviewRepository) UpdateInterviewById(ctx context.Context, interview database.Interview) (err error) {
	_, err = i.db.Exec(
		ctx,
		UpdateInterviewByIdQuery,
		interview.InterviewerId,
		interview.StudentId,
		interview.InterviewDate,
		interview.Grade,
		interview.SubjectiveRating,
		interview.Notes,
		interview.DeterminedEnglishLevel,
		interview.MainTask,
		interview.Question,
		time.Now(),
		interview.Id,
	)
	if err != nil {
		return errors.Wrap(err, "UpdateInterviewById - unable to execute update statement")
	}
	return err
}

func (i *InterviewRepository) DeactivateInterviewById(ctx context.Context, interviewId uint) (err error) {
	_, err = i.db.Exec(ctx, DeactivateByIdQuery, interviewId, time.Now())
	if err != nil {
		return errors.Wrapf(err, "DeactivateInterviewById - unable to set interview %d as deleted", interviewId)
	}
	return err
}

func (i *InterviewRepository) ActivateInterviewById(ctx context.Context, interviewId uint) (err error) {
	_, err = i.db.Exec(ctx, ActivateByIdQuery, interviewId, time.Now())
	if err != nil {
		return errors.Wrapf(err, "ActivateInterviewById - unable to set interview %d as active", interviewId)
	}
	return err
}
