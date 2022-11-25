package database

import "time"

type Interview struct {
	Id                     int                    `json:"id" db:"id"`
	InterviewerId          int                    `json:"interviewer_id" db:"interviewer_id"`
	StudentId              int                    `json:"student_id" db:"student_id"`
	InterviewDate          time.Time              `json:"interview_date" db:"interview_date"`
	Grade                  int                    `json:"grade" db:"grade"`
	SubjectiveRating       string                 `json:"subjective_rating" db:"subjective_rating"`
	Notes                  string                 `json:"notes,omitempty" db:"notes"`
	DeterminedEnglishLevel string                 `json:"determined_english_level" db:"determined_english_level"`
	MainTask               int                    `json:"main_task,omitempty" db:"main_task"`
	Question               map[string]interface{} `json:"question" db:"question"`
	CreatedAt              time.Time              `json:"-" db:"created_at"`
	UpdatedAt              time.Time              `json:"-" db:"updated_at"`
	Deleted                bool                   `json:"-" db:"deleted"`
}
