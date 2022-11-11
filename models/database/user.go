package database

import "time"

type User struct {
	Id           int       `json:"id"`
	LastName     string    `json:"last_name"`
	FirstName    string    `json:"first_name"`
	MiddleName   string    `json:"middle_name,omitempty"`
	Email        string    `json:"email"`
	Language     string    `json:"language"`
	EnglishLevel string    `json:"english_level"`
	Photo        string    `json:"photo"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	Deleted      bool      `json:"-"`
}
