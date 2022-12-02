package user

import (
	"strconv"

	"github.com/NEKETSKY/mnemosyne/pkg/api/common"
	apiUser "github.com/NEKETSKY/mnemosyne/pkg/api/user"
	"github.com/NEKETSKY/mnemosyne/pkg/file"
)

func (d *UserFullStuff) DbToProto() (u *apiUser.User) {
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

// err should be used in Update methods, in another cases it doesn't matter
func (u *UserFullStuff) ProtoToDb(protoUser *apiUser.User) (err error) {
	photoPath := ""
	midName := ""
	innerId := 0
	resumePath := ""
	if protoUser.Photo != nil {
		photoPath, _ = file.Save(protoUser.Photo.Name, protoUser.Photo.Content)
	}
	if protoUser.MiddleName != nil {
		midName = *protoUser.MiddleName
	}
	//err should be used in Update methods, in another cases it doesn't matter
	innerId, err = strconv.Atoi(protoUser.Id)

	u.Id = innerId
	u.LastName = protoUser.LastName
	u.FirstName = protoUser.FirstName
	u.MiddleName = midName
	u.Email = protoUser.Email
	u.Language = protoUser.Language
	u.EnglishLevel = protoUser.EnglishLevel
	u.Photo = photoPath

	if protoUser.Contact != nil {
		u.Telegram = protoUser.Contact.Telegram
		u.Discord = protoUser.Contact.Discord
		u.CommunicationChannel = protoUser.Contact.CommunicationChannel
	}
	if protoUser.Resume != nil {
		if protoUser.Resume.UploadedResume != nil {
			resumePath, _ = file.Save(protoUser.Resume.UploadedResume.Name, protoUser.Resume.UploadedResume.Content)
		}
		u.Experience = protoUser.Resume.Experience
		u.UploadedResume = resumePath
		u.Country = protoUser.Resume.Country
		u.City = protoUser.Resume.City
		u.TimeZone = protoUser.Resume.TimeZone
		u.MentorsNote = protoUser.Resume.MentorsNote
	}
	return
}
