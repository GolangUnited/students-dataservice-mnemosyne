package user

import (
	"context"
	"strings"
	"time"

	dbUser "github.com/NEKETSKY/mnemosyne/models/database/user"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type keyTx struct{}

func injectTx(ctx context.Context, tx *pgx.Tx) context.Context {
	return context.WithValue(ctx, keyTx{}, tx)
}
func extractTx(ctx context.Context) (tx *pgx.Tx) {
	if tx, ok := ctx.Value(keyTx{}).(*pgx.Tx); ok {
		return tx
	}
	return nil
}

// Add new user to db using proto struct User
func (u *UserRepository) AddUser(ctx context.Context, user *dbUser.UserFullStuff) (userId int, err error) {

	tr, _ := u.db.Begin(ctx)

	row := tr.QueryRow(ctx, AddUser, user.LastName, user.FirstName, user.MiddleName, user.Email, user.Language, user.EnglishLevel, user.Photo)
	if err != nil {
		return 0, errors.Wrap(err, "fail on trying to perform QueryRow")
	}
	err = row.Scan(&userId)
	if err != nil {
		return 0, errors.Wrap(err, "couldn't insert the new user's information and get the id of the new user")
	}
	//add role Student for everybody by default
	_, err = tr.Exec(ctx, AddRoleStudent, userId)
	if err != nil {
		_ = tr.Rollback(ctx)
		return 0, errors.Wrap(err, "failed to give to a user role Student in the system, try to add user again")

	}
	//inserting contacts info to database
	_, err = tr.Exec(ctx, AddContactById, userId, user.Telegram, user.Discord, user.CommunicationChannel)
	if err != nil {
		_ = tr.Rollback(ctx)
		return 0, errors.Wrap(err, "failed to insert user's contacts")

	}

	//inserting resume info to database
	_, err = tr.Exec(ctx, AddResumeById, userId, user.Experience, user.UploadedResume, user.Country, user.City, user.TimeZone, user.MentorsNote)
	if err != nil {
		_ = tr.Rollback(ctx)
		return 0, errors.Wrap(err, "failed to insert user's resume")
	}
	err = tr.Commit(ctx)
	return
}

// Get users using different filters
func (u *UserRepository) GetUsers(ctx context.Context, ur *dbUser.UserRequest) (users []dbUser.UserFullStuff, err error) {

	var b strings.Builder
	var rows pgx.Rows
	if ur.WithContacts && ur.WithResume {
		b.WriteString(GetUsersFull)
	} else if ur.WithContacts {
		b.WriteString(GetUsersWithContacts)
	} else if ur.WithResume {
		b.WriteString(GetUsersWithResume)
	} else {
		b.WriteString(GetUsers)
	}

	if ur.Role != "" {
		b.WriteString(SelectByRole)
	}
	if !ur.WithDeleted {
		b.WriteString(AliveUsers)
	}

	b.WriteString(OrderAsc)

	if ur.Role != "" {
		rows, err = u.db.Query(ctx, b.String(), ur.Role)
	} else {
		rows, err = u.db.Query(ctx, b.String())
	}

	if err != nil {
		return nil, errors.Wrap(err, "unable to get users from db")
	}
	users, err = pgx.CollectRows(rows, pgx.RowToStructByName[dbUser.UserFullStuff])
	if err != nil {
		return nil, errors.Wrap(err, "unable to collect rows ")
	}

	return users, err
}

func (u *UserRepository) GetUserById(ctx context.Context, userId int) (user *dbUser.UserFullStuff, err error) {

	b := strings.Builder{}
	b.WriteString(GetUsersFull)
	b.WriteString(GetUserById)
	rows, err := u.db.Query(ctx, b.String(), userId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get user by id from the db")
	}

	innerUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbUser.UserFullStuff])
	if err != nil {
		return nil, errors.Wrap(err, "GetUserById CollectRows error")
	}
	user = &innerUser
	return
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (user *dbUser.UserFullStuff, err error) {

	b := strings.Builder{}
	b.WriteString(GetUsersFull)
	b.WriteString(GetUserByEmail)
	rows, err := u.db.Query(ctx, b.String(), email)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get user by email from the db")
	}

	innerUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbUser.UserFullStuff])
	if err != nil {
		return nil, errors.Wrap(err, "GetUserByEmail CollectRows error")
	}
	user = &innerUser
	return
}

func (u *UserRepository) UpdateUserById(ctx context.Context, user *dbUser.UserFullStuff) (err error) {

	tr, _ := u.db.Begin(ctx)
	injectTx(ctx, &tr)

	_, err = tr.Exec(ctx, UpdateUserById, user.LastName, user.FirstName, user.MiddleName, user.Language, user.EnglishLevel, user.Photo, time.Now(), user.Id)
	if err != nil {
		return errors.Wrap(err, "unable to update basic user's info")
	}

	err = u.UpdateContact(ctx, &dbUser.Contact{
		Id:                   user.Id,
		Telegram:             user.Telegram,
		Discord:              user.Discord,
		CommunicationChannel: user.CommunicationChannel,
	})
	if err != nil {
		return errors.Wrap(err, "unable to update user's contact info")
	}

	err = u.UpdateResume(ctx, &dbUser.Resume{
		Id:             user.Id,
		Experience:     user.Experience,
		UploadedResume: user.UploadedResume,
		Country:        user.Country,
		City:           user.City,
		TimeZone:       user.TimeZone,
		MentorsNote:    user.MentorsNote,
	})
	if err != nil {
		return errors.Wrap(err, "unable to update user's resume info")
	}

	err = tr.Commit(ctx)
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

func (u *UserRepository) UpdateContact(ctx context.Context, contact *dbUser.Contact) (err error) {
	var tr pgx.Tx
	if extractTx(ctx) != nil {
		tr = *(extractTx(ctx))
	} else {
		tr, _ = u.db.Begin(ctx)
		defer tr.Commit(ctx)
	}

	_, err = tr.Exec(ctx, UpdateContactById, contact.Telegram, contact.Discord, contact.CommunicationChannel, time.Now(), contact.Id)
	if err != nil {
		_ = tr.Rollback(ctx)
		return errors.Wrap(err, "unable to update user's contact info")
	}

	return err
}

func (u *UserRepository) UpdateResume(ctx context.Context, resume *dbUser.Resume) (err error) {
	var tr pgx.Tx
	if extractTx(ctx) != nil {
		tr = *(extractTx(ctx))
	} else {
		tr, _ = u.db.Begin(ctx)
		defer tr.Commit(ctx)
	}

	_, err = tr.Exec(ctx, UpdateResumeById, resume.Experience, resume.UploadedResume, resume.Country, resume.City, resume.TimeZone, resume.MentorsNote, time.Now(), resume.Id)
	if err != nil {
		_ = tr.Rollback(ctx)
		return errors.Wrap(err, "failed to update user's resume")
	}
	return err
}

func (u *UserRepository) GetContactById(ctx context.Context, id int) (c *dbUser.Contact, err error) {
	row, err := u.db.Query(ctx, GetContactById, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a row from db")
	}

	innerContact, err := pgx.CollectOneRow(row, pgx.RowToStructByName[dbUser.Contact])
	if err != nil {
		return nil, errors.Wrap(err, "failed to collect fields into struct")
	}
	c = &innerContact
	return
}

func (u *UserRepository) GetResumeById(ctx context.Context, id int) (r *dbUser.Resume, err error) {
	row, err := u.db.Query(ctx, GetResumeById, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a row from db")
	}

	innerResume, err := pgx.CollectOneRow(row, pgx.RowToStructByName[dbUser.Resume])
	if err != nil {
		return nil, errors.Wrap(err, "failed to collect fields into struct")
	}
	r = &innerResume
	return
}
