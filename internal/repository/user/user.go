package user

import (
	"context"
	"time"

	"github.com/NEKETSKY/mnemosyne/models/database"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type UserRepository struct {
	db *pgx.Conn
}
type id struct {
	userid int
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Add new user to db using struct User
func (u *UserRepository) AddUser(ctx context.Context, user database.User) (userId int, err error) {
	rows, err := u.db.Query(ctx, AddUser, user.LastName, user.FirstName, user.MiddleName, user.Email, user.Language, user.EnglishLevel, user.Photo)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't insert the new user's information")
	}
	ids, err := pgx.CollectRows(rows, pgx.RowToStructByPos[id])
	if err != nil {
		return 0, errors.Wrap(err, "couldn't get the id of new the user")
	}
	userId = ids[0].userid
	return
}

func (u *UserRepository) GetAllUsers(ctx context.Context) (users []database.User, err error) {
	rows, _ := u.db.Query(ctx, GetAllUsers)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get users from db")
	}
	users, err = pgx.CollectRows(rows, pgx.RowToStructByPos[database.User])
	if err != nil {
		return nil, errors.Wrap(err, "unable to collect rows ")
	}
	return users, err
}

func (u *UserRepository) GetUserById(ctx context.Context, userId int) (user database.User, err error) {

	rows, err := u.db.Query(ctx, GetUserById, userId)
	if err != nil {
		return database.User{}, errors.Wrap(err, "unable to get user by id from the db")
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[database.User])
	if err != nil {
		return database.User{}, errors.Wrap(err, "GetUserById CollectRows error")
	}
	user = users[0]
	return
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, userEmail string) (user database.User, err error) {

	rows, err := u.db.Query(ctx, GetUserById, userEmail)
	if err != nil {
		return database.User{}, errors.Wrap(err, "unable to get user by email from the db")
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[database.User])
	if err != nil {
		return database.User{}, errors.Wrap(err, "GetUserByEmail CollectRows error")
	}
	user = users[0]
	return
}

func (u *UserRepository) UpdateUserById(ctx context.Context, user database.User) (err error) {
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
