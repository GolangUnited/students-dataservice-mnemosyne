package database

import "time"

type User struct {
	Id           int       `json:"id" db:"id"`
	LastName     string    `json:"last_name" db:"last_name"`
	FirstName    string    `json:"first_name" db:"first_name"`
	MiddleName   string    `json:"middle_name,omitempty" db:"middle_name"`
	Email        string    `json:"email" db:"email"`
	Language     string    `json:"language" db:"language"`
	EnglishLevel string    `json:"english_level" db:"english_level"`
	Photo        string    `json:"photo" db:"photo"`
	CreatedAt    time.Time `json:"-" db:"created_at"`
	UpdatedAt    time.Time `json:"-" db:"updated_at"`
	Deleted      bool      `json:"-" db:"deleted"`
}
