package lessons

import (
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/lessons"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Lessons struct {
	Id           uint32    `json:"id" db:"id"`
	Presentation string    `json:"presentation" db:"presentation"`
	VideoLink    string    `json:"video_link" db:"video_link"`
	LessonDate   time.Time `json:"lesson_date" db:"lesson_date"`
	Homework     string    `json:"homework" db:"homework"`
	LecturerId   uint32    `json:"lecturer_id" db:"lecturer_id"`
	GroupId      uint32    `json:"group_id" db:"group_id"`
	Language     string    `json:"language" db:"language"`
	CreatedAt    time.Time `json:"-" db:"created_at"`
	UpdatedAt    time.Time `json:"-" db:"updated_at"`
	Deleted      bool      `json:"-" db:"deleted"`
}

func (l *Lessons) LessonFromProto(protoLesson *lessons.LessonRequest) {
	l.Id = protoLesson.GetId()
	l.Presentation = protoLesson.GetPresentation()
	l.VideoLink = protoLesson.GetVideoLink()
	l.LessonDate = protoLesson.LessonDate.AsTime()
	l.Homework = protoLesson.GetHomework()
	l.LecturerId = protoLesson.GetLecturerId()
	l.GroupId = protoLesson.GetGroupId()
	l.Language = protoLesson.GetLanguage()

}
func (l *Lessons) LessonToProto() *lessons.LessonResponse {
	return &lessons.LessonResponse{
		Id:           l.Id,
		Presentation: l.Presentation,
		VideoLink:    l.VideoLink,
		LessonDate:   timestamppb.New(l.LessonDate),
		Homework:     l.Homework,
		LecturerId:   l.LecturerId,
		GroupId:      l.GroupId,
		Language:     l.Language,
		CreatedAt:    timestamppb.New(l.CreatedAt),
		UpdatedAt:    timestamppb.New(l.UpdatedAt),
		Deleted:      l.Deleted,
	}
}
