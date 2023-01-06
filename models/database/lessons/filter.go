package lessons

import (
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/lessons"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Filter struct {
	GroupId      uint32
	Presentation string
	VideoLink    string
	LessonDate   time.Time
	Homework     string
	LecturerId   uint32
	Language     string
	Deleted      bool
}

func (f *Filter) FromProto(lf *lessons.Filter) {
	if lf == nil {
		return
	}
	f.GroupId = lf.GetGroupId()
	f.Presentation = lf.GetPresentation()
	f.VideoLink = lf.GetVideoLink()
	f.LessonDate = lf.LessonDate.AsTime()
	f.Homework = lf.GetHomework()
	f.LecturerId = lf.GetLecturerId()
	f.Language = lf.GetLanguage()
}

func (f *Filter) ToProto() *lessons.Filter {
	return &lessons.Filter{
		GroupId:      f.GroupId,
		Presentation: f.Presentation,
		VideoLink:    f.VideoLink,
		LessonDate:   timestamppb.New(f.LessonDate),
		Homework:     f.Homework,
		LecturerId:   f.LecturerId,
		Language:     f.Language,
		Deleted:      f.Deleted,
	}
}
