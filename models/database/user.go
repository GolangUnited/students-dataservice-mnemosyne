package database

type User struct {
	Id           int     `json:"id" db:"id"`
	LastName     string  `json:"last_name" db:"last_name"`
	FirstName    string  `json:"first_name" db:"first_name"`
	MiddleName   string  `json:"middle_name,omitempty" db:"middle_name"`
	Email        string  `json:"email" db:"email"`
	Language     string  `json:"language" db:"language"`
	EnglishLevel string  `json:"english_level" db:"english_level"`
	Photo        string  `json:"photo" db:"photo"`
	Contact      Contact `json:"contact" db:"-"`
	Resume       Resume  `json:"resume" db:"-"`
}

type Contact struct {
	Id                   int    `json:"id" db:"user_id"`
	Telegram             string `json:"telegram" db:"telegram"`
	Discord              string `json:"discord" db:"discord"`
	CommunicationChannel string `json:"communication_channel" db:"communication_channel"`
}

type Resume struct {
	Id             int    `json:"id" db:"user_id"`
	Experience     string `json:"experience" db:"experience"`
	UploadedResume string `json:"uploaded_resume" db:"uploaded_resume"`
	Country        string `json:"country" db:"country"`
	City           string `json:"city" db:"city"`
	TimeZone       string `json:"time_zone" db:"time_zone"`
	MentorsNote    string `json:"mentors_note" db:"mentors_note"`
}

type UserRequest struct {
	WithContacts bool
	WithResume   bool
	WithDeleted  bool
	Role         string
	FieldName    string
	FieldValue   string
}
