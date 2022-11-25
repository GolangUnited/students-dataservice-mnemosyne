package user

import (
	"context"
	"time"

	dbUser "github.com/NEKETSKY/mnemosyne/models/database/user"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

// type userFullStuff struct {
// 	id                   int    `json:"id" db:"id"`
// 	lastName             string `json:"last_name" db:"last_name"`
// 	firstName            string `json:"first_name" db:"first_name"`
// 	middleName           string `json:"middle_name,omitempty" db:"middle_name"`
// 	email                string `json:"email" db:"email"`
// 	language             string `json:"language" db:"language"`
// 	englishLevel         string `json:"english_level" db:"english_level"`
// 	photo                string `json:"photo" db:"photo"`
// 	telegram             string `json:"telegram" db:"telegram"`
// 	discord              string `json:"discord" db:"discord"`
// 	communicationChannel string `json:"communication_channel" db:"communication_channel"`
// 	experience           string `json:"experience" db:"experience"`
// 	uploadedResume       string `json:"uploaded_resume" db:"uploaded_resume"`
// 	country              string `json:"country" db:"country"`
// 	city                 string `json:"city" db:"city"`
// 	timeZone             string `json:"time_zone" db:"time_zone"`
// 	mentorsNote          string `json:"mentors_note" db:"mentors_note"`
// }

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Add new user to db using struct User
func (u *UserRepository) AddUser(ctx context.Context, user dbUser.BaseUser) (userId int, err error) {
	row := u.db.QueryRow(ctx, AddUser, user.LastName, user.FirstName, user.MiddleName, user.Email, user.Language, user.EnglishLevel, user.Photo)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't insert the new user's information")
	}
	err = row.Scan(&userId)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't get the id of new the user")
	}

	return
}

func (u *UserRepository) GetUsers(ctx context.Context, ur *dbUser.UserRequest) (users []dbUser.BaseUser, err error) {

	rows, _ := u.db.Query(ctx, GetUsersFull, ur.WithDeleted)

	if err != nil {
		return nil, errors.Wrap(err, "unable to get users from db")
	}
	users, err = pgx.CollectRows(rows, pgx.RowToStructByName[dbUser.BaseUser])

	if err != nil {
		return nil, errors.Wrap(err, "unable to collect rows ")
	}

	return users, err
}

func (u *UserRepository) GetUserById(ctx context.Context, userId int) (user dbUser.BaseUser, err error) {

	rows, err := u.db.Query(ctx, GetUserById, userId)
	if err != nil {
		return dbUser.BaseUser{}, errors.Wrap(err, "unable to get user by id from the db")
	}

	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[dbUser.BaseUser])
	if err != nil {
		return dbUser.BaseUser{}, errors.Wrap(err, "GetUserById CollectRows error")
	}
	return
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, userEmail string) (user dbUser.BaseUser, err error) {

	rows, err := u.db.Query(ctx, GetUserByEmail, userEmail)
	if err != nil {
		return dbUser.BaseUser{}, errors.Wrap(err, "unable to get user by email from the db")
	}

	user, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[dbUser.BaseUser])
	if err != nil {
		return dbUser.BaseUser{}, errors.Wrap(err, "GetUserByEmail CollectRows error")
	}

	return
}

func (u *UserRepository) UpdateUserById(ctx context.Context, user dbUser.BaseUser) (err error) {
	_, err = u.db.Exec(ctx, UpdateUserById, user.LastName, user.FirstName, user.MiddleName, user.Language, user.EnglishLevel, user.Photo, time.Now(), user.Id)
	if err != nil {
		return errors.Wrap(err, "unable to execute update statement")
	}
	return err
}

func (u *UserRepository) ActivateUserById(ctx context.Context, userId int) (err error) {
	_, err = u.db.Exec(ctx, ActivateById, userId, time.Now())
	if err != nil {
		return errors.Wrapf(err, "unable to set user %d as active", userId)
	}
	return err
}

func (u *UserRepository) DeactivateUserById(ctx context.Context, userId int) (err error) {
	_, err = u.db.Exec(ctx, DeactivateById, userId, time.Now())
	if err != nil {
		return errors.Wrapf(err, "unable to set user %d as deleted", userId)
	}
	return err
}
