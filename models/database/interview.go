package database

import (
	"encoding/json"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/interview"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Interview struct {
	Id                     uint32                 `json:"id" db:"id"`
	InterviewerId          uint32                 `json:"interviewer_id" db:"interviewer_id"`
	StudentId              uint32                 `json:"student_id" db:"student_id"`
	InterviewDate          time.Time              `json:"interview_date" db:"interview_date"`
	Grade                  uint32                 `json:"grade" db:"grade"`
	SubjectiveRating       string                 `json:"subjective_rating" db:"subjective_rating"`
	Notes                  string                 `json:"notes,omitempty" db:"notes"`
	DeterminedEnglishLevel string                 `json:"determined_english_level" db:"determined_english_level"`
	MainTask               uint32                 `json:"main_task,omitempty" db:"main_task"`
	Question               map[string]interface{} `json:"question" db:"question"`
	CreatedAt              time.Time              `json:"-" db:"created_at"`
	UpdatedAt              time.Time              `json:"-" db:"updated_at"`
	Deleted                bool                   `json:"-" db:"deleted"`
}

func InterviewFromProto(protoInterview *interview.InterviewRequest) (interviewDb Interview, err error) {
	var question map[string]interface{}
	err = json.Unmarshal([]byte(protoInterview.GetQuestion()), &question)
	if err != nil {
		return Interview{}, err
	}
	interviewDb = Interview{
		Id:                     protoInterview.GetId(),
		InterviewerId:          protoInterview.GetInterviewerId(),
		StudentId:              protoInterview.GetStudentId(),
		InterviewDate:          protoInterview.GetInterviewDate().AsTime(),
		Grade:                  protoInterview.GetGrade(),
		SubjectiveRating:       protoInterview.GetSubjectiveRating(),
		Notes:                  protoInterview.GetNotes(),
		DeterminedEnglishLevel: protoInterview.GetDeterminedEnglishLevel(),
		MainTask:               protoInterview.GetMainTask(),
		Question:               question,
	}
	return interviewDb, err
}

func (i *Interview) ToProto() *interview.InterviewResponse {
	question, _ := json.Marshal(i.Question)
	return &interview.InterviewResponse{
		Id:                     i.Id,
		InterviewerId:          i.InterviewerId,
		StudentId:              i.StudentId,
		InterviewDate:          timestamppb.New(i.InterviewDate),
		Grade:                  i.Grade,
		SubjectiveRating:       i.SubjectiveRating,
		Notes:                  i.Notes,
		DeterminedEnglishLevel: i.DeterminedEnglishLevel,
		MainTask:               i.MainTask,
		Question:               string(question),
		CreatedAt:              timestamppb.New(i.CreatedAt),
		UpdatedAt:              timestamppb.New(i.UpdatedAt),
		Deleted:                i.Deleted,
	}
}
