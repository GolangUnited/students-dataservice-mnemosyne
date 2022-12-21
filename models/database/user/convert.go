package user

import (
	"strconv"

	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/common"
	apiUser "github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/user"
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

	//err should be used in Update methods, in another cases it doesn't matter
	innerId, err := strconv.Atoi(protoUser.Id)

	u.Id = innerId
	u.LastName = protoUser.GetLastName()
	u.FirstName = protoUser.GetFirstName()
	u.MiddleName = protoUser.GetMiddleName()
	u.Email = protoUser.GetEmail()
	u.Language = protoUser.GetLanguage()
	u.EnglishLevel = protoUser.GetEnglishLevel()

	u.Telegram = protoUser.Contact.GetTelegram()
	u.Discord = protoUser.Contact.GetDiscord()
	u.CommunicationChannel = protoUser.Contact.GetCommunicationChannel()

	u.Experience = protoUser.Resume.GetExperience()

	u.Country = protoUser.Resume.GetCountry()
	u.City = protoUser.Resume.GetCity()
	u.TimeZone = protoUser.Resume.GetTimeZone()
	u.MentorsNote = protoUser.Resume.GetMentorsNote()

	return
}
