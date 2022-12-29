package validate

import (
	"errors"

	"net/url"

	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/interview"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/project"
)

var (
	ErrTooLongRating   = errors.New("too many characters written in field 'SubjectiveRating', must be less then 30")
	ErrTooLongEngLevel = errors.New("too many characters written in field 'DeterminedEnglishLevel', must be less then 50")
	ErrNotUrl          = errors.New("value of field 'GitUrl' is not a link")
)

func IsValid(i interface{}) (bool, error) {
	switch v := i.(type) {
	case *interview.InterviewRequest:
		rating := []rune(v.GetSubjectiveRating())

		if len(rating) > 30 {
			return false, ErrTooLongRating
		}
		engLevel := []rune(v.GetDeterminedEnglishLevel())
		if len(engLevel) > 50 {
			return false, ErrTooLongEngLevel
		}
	case *project.ProjectRequest:
		_, err := url.ParseRequestURI(v.GetGitUrl())
		if err != nil {
			return false, ErrNotUrl
		}
	}
	return true, nil
}
