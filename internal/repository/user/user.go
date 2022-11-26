package user

import (
	"context"
	"strconv"
	"strings"
	"time"

	dbUser "github.com/NEKETSKY/mnemosyne/models/database/user"
	apiUser "github.com/NEKETSKY/mnemosyne/pkg/api/user"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type userFullStuff struct {
	Id                   int    `db:"id"`
	LastName             string `db:"last_name"`
	FirstName            string `db:"first_name"`
	MiddleName           string `db:"middle_name"`
	Email                string `db:"email"`
	Language             string `db:"language"`
	EnglishLevel         string `db:"english_level"`
	Photo                string `db:"photo"`
	Telegram             string `db:"telegram"`
	Discord              string `db:"discord"`
	CommunicationChannel string `db:"communication_channel"`
	Experience           string `db:"experience"`
	UploadedResume       string `db:"uploaded_resume"`
	Country              string `db:"country"`
	City                 string `db:"city"`
	TimeZone             string `db:"time_zone"`
	MentorsNote          string `db:"mentors_note"`
}

//func parseUserFullStuff(*userFullStuff) {}

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Add new user to db using struct User
func (u *UserRepository) AddUser(ctx context.Context, user *apiUser.User) (userId *apiUser.Id, err error) {

	innerUserId := 0
	row := u.db.QueryRow(ctx, AddUser, user.LastName, user.FirstName, user.MiddleName, user.Email, user.Language, user.EnglishLevel, user.Photo)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't insert the new user's information")
	}

	err = row.Scan(&innerUserId)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get the id of new the user")
	}
	//add role Student for everybody by default
	_, err = u.db.Exec(ctx, AddRoleStudent, innerUserId)
	if err != nil {
		err = errors.Wrap(err, "failed to give to a user role Sudent in the system")
	}

	//if exists push contact info to database
	if user.Contact != nil {
		_, err = u.db.Exec(ctx, AddContactById, innerUserId, user.Contact.Telegram, user.Contact.Discord, user.Contact.CommunicationChannel)
		if err != nil {
			err = errors.Wrap(err, "failed to insert user's contacts")
		}
	}

	//if exists push resume info to database
	if user.Resume != nil {

		_, err = u.db.Exec(ctx, AddResumeById, innerUserId, user.Resume.Experience, user.Resume.UploadedResume, user.Resume.Country, user.Resume.City, user.Resume.TimeZone, user.Resume.MentorsNote)
		if err != nil {
			err = errors.Wrap(err, "failed to insert user's resume")
		}
	}
	userId = &(apiUser.Id{Id: strconv.Itoa(innerUserId)})
	return
}

func (u *UserRepository) GetUsers(ctx context.Context, ur *apiUser.UserRequest) (users *apiUser.Users, err error) {

	var b strings.Builder
	var innerApiUserSlice []*apiUser.User

	b.WriteString(GetUsersFull)
	switch {
	case ur.Role.Role == "student":
		b.WriteString(SelectStudents)
	case ur.Role.Role == "mentor":
		b.WriteString(SelectMentors)
	}

	if !ur.Option.WithDeleted {
		b.WriteString(WithoutDeleted)
	}

	b.WriteString(`
	order by u.id asc`)
	rows, _ := u.db.Query(ctx, b.String())

	if err != nil {
		return nil, errors.Wrap(err, "unable to get users from db")
	}
	innerUsers, err := pgx.CollectRows(rows, pgx.RowToStructByName[userFullStuff])
	if err != nil {
		return nil, errors.Wrap(err, "unable to collect rows ")
	}

	for _, innerUser := range innerUsers {
		temp := innerUser
		innerApiUser := &apiUser.User{
			Id:           strconv.Itoa(temp.Id),
			LastName:     temp.LastName,
			FirstName:    temp.FirstName,
			MiddleName:   &temp.MiddleName,
			Email:        temp.Email,
			Language:     temp.Language,
			EnglishLevel: temp.EnglishLevel,
			Photo:        temp.Photo,
		}
		switch {
		case ur.Option.WithContacts:
			innerApiUser.Contact = &apiUser.Contact{
				Telegram:             temp.Telegram,
				Discord:              temp.Discord,
				CommunicationChannel: temp.CommunicationChannel,
			}
		case ur.Option.WithResume:
			innerApiUser.Resume = &apiUser.Resume{
				UploadedResume: temp.UploadedResume,
				MentorsNote:    temp.MentorsNote,
				Experience:     temp.Experience,
				Country:        temp.Country,
				City:           temp.City,
				TimeZone:       temp.TimeZone,
			}
		}
		innerApiUserSlice = append(innerApiUserSlice, innerApiUser)
	}
	users = &apiUser.Users{Users: innerApiUserSlice}
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
