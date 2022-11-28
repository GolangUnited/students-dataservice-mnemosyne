package user

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/NEKETSKY/mnemosyne/pkg/api/common"
	apiUser "github.com/NEKETSKY/mnemosyne/pkg/api/user"
	"github.com/NEKETSKY/mnemosyne/pkg/file"
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

type contact struct {
	Id                   int    `db:"user_id"`
	Telegram             string `db:"telegram"`
	Discord              string `db:"discord"`
	CommunicationChannel string `db:"communication_channel"`
}

type resume struct {
	Id             int    `db:"user_id"`
	Experience     string `db:"experience"`
	UploadedResume string `db:"uploaded_resume"`
	Country        string `db:"country"`
	City           string `db:"city"`
	TimeZone       string `db:"time_zone"`
	MentorsNote    string `db:"mentors_note"`
}
type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func dbToProto(d *userFullStuff) (u *apiUser.User) {
	u = &apiUser.User{
		Id:           strconv.Itoa(d.Id),
		LastName:     d.LastName,
		FirstName:    d.FirstName,
		MiddleName:   &d.MiddleName,
		Email:        d.Email,
		Language:     d.Language,
		EnglishLevel: d.EnglishLevel,
		Photo:        &common.File{Name: d.Photo},
	}
	u.Contact = &apiUser.Contact{
		Telegram:             d.Telegram,
		Discord:              d.Discord,
		CommunicationChannel: d.CommunicationChannel,
	}
	u.Resume = &apiUser.Resume{
		MentorsNote:    d.MentorsNote,
		Experience:     d.Experience,
		Country:        d.Country,
		City:           d.City,
		TimeZone:       d.TimeZone,
		UploadedResume: &common.File{Name: d.UploadedResume},
	}

	return
}

// Add new user to db using proto struct User
func (u *UserRepository) AddUser(ctx context.Context, user *apiUser.User) (userId *apiUser.Id, err error) {

	innerUserId := 0
	photoPath := ""
	if user.Photo != nil {
		photoPath, _ = file.Save(user.Photo.Name, user.Photo.Content)
	}
	row := u.db.QueryRow(ctx, AddUser, user.LastName, user.FirstName, user.MiddleName, user.Email, user.Language, user.EnglishLevel, photoPath)
	if err != nil {
		return nil, errors.Wrap(err, "fail on trying to perform QueryRow")
	}

	err = row.Scan(&innerUserId)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't insert the new user's information and get the id of the new user")
	}
	//add role Student for everybody by default
	_, err = u.db.Exec(ctx, AddRoleStudent, innerUserId)
	if err != nil {
		err = errors.Wrap(err, "failed to give to a user role Student in the system")
	}

	//if exists push contact info to database
	if user.Contact != nil {
		_, err = u.db.Exec(ctx, AddContactById, innerUserId, user.Contact.Telegram, user.Contact.Discord, user.Contact.CommunicationChannel)
		if err != nil {
			err = errors.Wrap(err, "failed to insert user's contacts")
		}
	}

	//if exists push resume info to database
	resumePath := ""
	if user.Resume != nil {
		if user.Resume.UploadedResume != nil {
			resumePath, _ = file.Save(user.Resume.UploadedResume.Name, user.Resume.UploadedResume.Content)
		}
		_, err = u.db.Exec(ctx, AddResumeById, innerUserId, user.Resume.Experience, resumePath, user.Resume.Country, user.Resume.City, user.Resume.TimeZone, user.Resume.MentorsNote)
		if err != nil {
			err = errors.Wrap(err, "failed to insert user's resume")
		}
	}
	userId = &(apiUser.Id{Id: strconv.Itoa(innerUserId)})
	return
}

// Get users using different filters
func (u *UserRepository) GetUsers(ctx context.Context, ur *apiUser.UserRequest) (users *apiUser.Users, err error) {

	var b strings.Builder
	var innerApiUserSlice []*apiUser.User
	var rows pgx.Rows

	b.WriteString(GetUsersFull)
	if ur.Role.Role != "" {
		b.WriteString(SelectByRole)
	}
	if !ur.Option.WithDeleted {
		b.WriteString(AliveUsers)
	}

	b.WriteString(`
	order by u.id asc`)

	if ur.Role.Role != "" {
		rows, err = u.db.Query(ctx, b.String(), ur.Role.Role)
	} else {
		rows, err = u.db.Query(ctx, b.String())
	}
	if err != nil {
		return nil, errors.Wrap(err, "unable to get users from db")
	}
	innerUsers, err := pgx.CollectRows(rows, pgx.RowToStructByName[userFullStuff])
	if err != nil {
		return nil, errors.Wrap(err, "unable to collect rows ")
	}

	for _, innerUser := range innerUsers {
		temp := innerUser
		innerApiUser := dbToProto(&temp)
		if !ur.Option.WithContacts {
			innerApiUser.Contact = nil
		}
		if !ur.Option.WithResume {
			innerApiUser.Resume = nil
		}
		innerApiUserSlice = append(innerApiUserSlice, innerApiUser)
	}
	users = &apiUser.Users{Users: innerApiUserSlice}
	return users, err
}

func (u *UserRepository) GetUserById(ctx context.Context, userId int) (user *apiUser.User, err error) {

	b := strings.Builder{}
	b.WriteString(GetUsersFull)
	b.WriteString(GetUserById)
	rows, err := u.db.Query(ctx, b.String(), userId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get user by id from the db")
	}

	innerUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[userFullStuff])
	if err != nil {
		return nil, errors.Wrap(err, "GetUserById CollectRows error")
	}
	user = dbToProto(&innerUser)
	return
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, userEmail string) (user *apiUser.User, err error) {

	b := strings.Builder{}
	b.WriteString(GetUsersFull)
	b.WriteString(GetUserByEmail)
	rows, err := u.db.Query(ctx, b.String(), userEmail)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get user by email from the db")
	}

	innerUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[userFullStuff])
	if err != nil {
		return nil, errors.Wrap(err, "GetUserByEmail CollectRows error")
	}
	user = dbToProto(&innerUser)
	return
}

func (u *UserRepository) UpdateUserById(ctx context.Context, user *apiUser.User) (err error) {
	photoPath := ""
	if user.Photo != nil {
		photoPath, _ = file.Save(user.Photo.Name, user.Photo.Content)
	}
	innerId, err := strconv.Atoi(user.Id)
	if err != nil {
		return errors.Wrap(err, "invalid user's id")
	}
	_, err = u.db.Exec(ctx, UpdateUserById, user.LastName, user.FirstName, user.MiddleName, user.Language, user.EnglishLevel, photoPath, time.Now(), innerId)
	if err != nil {
		return errors.Wrap(err, "unable to update basic user's info")
	}

	if user.Contact != nil {
		user.Contact.Id = user.Id
		err = u.UpdateContact(ctx, user.Contact)
		if err != nil {
			return errors.Wrap(err, "unable to update user's contact info")
		}
	}
	if user.Resume != nil {
		user.Resume.Id = user.Id
		err = u.UpdateResume(ctx, user.Resume)
		if err != nil {
			return errors.Wrap(err, "unable to update user's resume info")
		}
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

func (u *UserRepository) GetEmailById(ctx context.Context, id int) (email *apiUser.Email, err error) {

	row := u.db.QueryRow(ctx, GetEmailById, id)
	var innerEmail string
	err = row.Scan(&innerEmail)
	email = &apiUser.Email{Email: innerEmail}
	return
}

func (u *UserRepository) UpdateContact(ctx context.Context, contact *apiUser.Contact) (err error) {
	innerId, err := strconv.Atoi(contact.Id)
	if err != nil {
		return errors.Wrap(err, "invalid user's id")
	}
	_, err = u.db.Exec(ctx, UpdateContactById, contact.Telegram, contact.Discord, contact.CommunicationChannel, time.Now(), innerId)
	if err != nil {
		return errors.Wrap(err, "unable to update user's contact info")
	}
	return err
}

func (u *UserRepository) UpdateResume(ctx context.Context, resume *apiUser.Resume) (err error) {
	resumePath, err := file.Save(resume.UploadedResume.Name, resume.UploadedResume.Content)
	if err != nil {
		return errors.Wrap(err, "failed to save file with resume")
	}
	innerId, err := strconv.Atoi(resume.Id)
	if err != nil {
		return errors.Wrap(err, "invalid user's id")
	}
	_, err = u.db.Exec(ctx, UpdateResumeById, resume.Experience, resumePath, resume.Country, resume.City, resume.TimeZone, resume.MentorsNote, time.Now(), innerId)
	if err != nil {
		return errors.Wrap(err, "failed to update user's resume")
	}
	return err
}

func (u *UserRepository) GetContactById(ctx context.Context, id int) (c *apiUser.Contact, err error) {
	row, err := u.db.Query(ctx, GetContactById, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a row from db")
	}

	innerContact, err := pgx.CollectOneRow(row, pgx.RowToStructByName[contact])
	if err != nil {
		return nil, errors.Wrap(err, "failed to collect fields into struct")
	}
	c = &apiUser.Contact{
		Id:                   strconv.Itoa(innerContact.Id),
		Telegram:             innerContact.Telegram,
		Discord:              innerContact.Discord,
		CommunicationChannel: innerContact.CommunicationChannel,
	}
	return
}

func (u *UserRepository) GetResumeById(ctx context.Context, id int) (r *apiUser.Resume, err error) {
	row, err := u.db.Query(ctx, GetResumeById, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a row from db")
	}

	innerResume, err := pgx.CollectOneRow(row, pgx.RowToStructByName[resume])
	if err != nil {
		return nil, errors.Wrap(err, "failed to collect fields into struct")
	}
	r = &apiUser.Resume{
		Id:             strconv.Itoa(innerResume.Id),
		Experience:     innerResume.Experience,
		UploadedResume: &common.File{Name: innerResume.UploadedResume},
		Country:        innerResume.Country,
		City:           innerResume.City,
		TimeZone:       innerResume.TimeZone,
		MentorsNote:    innerResume.MentorsNote,
	}
	return
}
