package user

import (
	"strconv"
	"strings"

	apiUser "github.com/NEKETSKY/mnemosyne/pkg/api/user"
	"github.com/pkg/errors"
)

func DbUserToProtoUser(d *BaseUser) (u *apiUser.User) {
	u = &(apiUser.User{Id: strconv.Itoa(d.Id),
		LastName:     d.LastName,
		FirstName:    d.FirstName,
		MiddleName:   &d.MiddleName,
		Email:        d.Email,
		Language:     d.Language,
		EnglishLevel: d.EnglishLevel,
		Photo:        d.Photo,
		Contact: &apiUser.Contact{
			Telegram:             d.Contact.Telegram,
			Discord:              d.Contact.Discord,
			CommunicationChannel: d.Contact.CommunicationChannel,
		},
		Resume: &apiUser.Resume{},
	})
	return
}
func ProtoUserToDbUser(u *apiUser.User) (d *BaseUser, err error) {
	var innerId int
	innerId, err = strconv.Atoi(u.Id)
	if err != nil {
		err = errors.Wrap(err, "invalid user's id value")
	}
	var engLevel strings.Builder
	for index, value := range u.EnglishLevel {
		if index > 2 {
			break
		}
		engLevel.WriteRune(value)
	}
	d = &BaseUser{
		Id:           innerId,
		LastName:     u.LastName,
		FirstName:    u.FirstName,
		MiddleName:   *u.MiddleName,
		Email:        u.Email,
		Language:     u.Language,
		EnglishLevel: engLevel.String(),
		Photo:        u.Photo,
	}
	return
}
