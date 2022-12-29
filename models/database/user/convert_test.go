package user

import (
	"testing"

	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/common"
	proto "github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/user"
	"github.com/stretchr/testify/assert"
)

type TestData struct {
	TestValue *UserFullStuff
	Expected  *proto.User
}

var tt = []TestData{
	TestData{TestValue: &UserFullStuff{
		Id:                   1,
		LastName:             "Adminton",
		FirstName:            "Admin",
		MiddleName:           "",
		Email:                "admin@gmail.com",
		Language:             "En,Ru,By, Ua",
		EnglishLevel:         "B1+",
		Photo:                "some/local/way/to/photo",
		Telegram:             "@superAdmin",
		Discord:              "discord_acc",
		CommunicationChannel: "tg",
		Experience:           "huge bla bla bla",
		UploadedResume:       "some/local/way/to/uploaded/resume",
		Country:              "Cracozhia",
		City:                 "Zootopia",
		TimeZone:             "+01 GMT",
		MentorsNote:          "another bla bla",
	},
		Expected: &proto.User{
			Id:           "1",
			LastName:     "Adminton",
			FirstName:    "Admin",
			MiddleName:   nil,
			Email:        "admin@gmail.com",
			Language:     "En,Ru,By, Ua",
			EnglishLevel: "B1+",
			Photo: &common.File{
				Name: "some/local/way/to/photo",
			},
			Resume: &proto.Resume{
				UploadedResume: &common.File{
					Name: "some/local/way/to/uploaded/resume",
				},
				MentorsNote: "another bla bla",
				Experience:  "huge bla bla bla",
				Country:     "Cracozhia",
				City:        "Zootopia",
				TimeZone:    "+01 GMT",
			},
			Contact: &proto.Contact{
				Telegram:             "@superAdmin",
				Discord:              "discord_acc",
				CommunicationChannel: "tg",
			},
		},
	},
	TestData{TestValue: &UserFullStuff{
		Id:                   1,
		LastName:             "Adminton",
		FirstName:            "Admin",
		MiddleName:           "",
		Email:                "admin@gmail.com",
		Language:             "En,Ru,By, Ua",
		EnglishLevel:         "B1+",
		Photo:                "some/local/way/to/photo",
		Telegram:             "",
		Discord:              "",
		CommunicationChannel: "",
		Experience:           "",
		UploadedResume:       "",
		Country:              "",
		City:                 "",
		TimeZone:             "",
		MentorsNote:          "",
	},
		Expected: &proto.User{
			Id:           "1",
			LastName:     "Adminton",
			FirstName:    "Admin",
			MiddleName:   nil,
			Email:        "admin@gmail.com",
			Language:     "En,Ru,By, Ua",
			EnglishLevel: "B1+",
			Photo: &common.File{
				Name: "some/local/way/to/photo",
			},
			Resume: &proto.Resume{
				UploadedResume: &common.File{
					Name: "",
				},
				MentorsNote: "",
				Experience:  "",
				Country:     "",
				City:        "",
				TimeZone:    "",
			},
			Contact: &proto.Contact{
				Telegram:             "",
				Discord:              "",
				CommunicationChannel: "",
			},
		},
	},
}

func TestDbToProto(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tt {
		test.Expected.MiddleName = &test.TestValue.MiddleName
		assert.Equal(test.Expected, test.TestValue.DbToProto())
	}
}
