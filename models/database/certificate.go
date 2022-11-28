package database

import "time"

type Certificate struct {
	Id         int       `json:"id" db:"id"`
	UserId     string    `json:"user_id" db:"user_id"`
	IssueDate  time.Time `json:"issue_date" db:"issue_date"`
	ExpireDate time.Time `json:"expire_date" db:"expire_date"`
}
