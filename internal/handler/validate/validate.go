package validate

import (
	"errors"

	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/interview"
	// "net/url"
	//"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/projects"
)

var (
	ErrTooLongRating   = errors.New("too many characters written in field 'Subjective rating', must be less then 30")
	ErrTooLongEngLevel = errors.New("too many characters written in field 'Determined English Level', must be less then 50")
	ErrNotUrl          = errors.New("value of field 'Git_url' is not a link")
)

func Validate(i interface{}) error {

	switch v := i.(type) {
	case interview.InterviewRequest:
		rating := []rune(v.GetSubjectiveRating())
		if len(rating) > 30 {
			return ErrTooLongRating
		}
		engLevel := []rune(v.GetDeterminedEnglishLevel())
		if len(engLevel) > 50 {
			return ErrTooLongEngLevel
		}
	// case projects.ProjectRequest :
	// 	l:=v.GetGitUrl
	// 	_,err:=url.ParseRequestURI(l)
	// 	if err!=nil{
	// 		return ErrNotUrl
	// 	}
	default:
		return errors.New("type is unknown, can't validate")

	}
	return nil
}
