package team

import (
	"github.com/NEKETSKY/mnemosyne/pkg/api/team"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Filter struct {
	UserId    uint32
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Deleted   bool
	MentorId  uint32
}

func (f *Filter) FromProto(teamFilter *team.TeamListFilter) {
	if teamFilter == nil {
		return
	}
	f.UserId = teamFilter.GetUserId()
	f.Name = teamFilter.GetName()
	f.StartDate = teamFilter.GetStartDate().AsTime()
	f.EndDate = teamFilter.GetEndDate().AsTime()
	f.Deleted = teamFilter.GetDeleted()
	f.MentorId = teamFilter.GetMentorId()
}

func (f *Filter) ToProto() *team.TeamListFilter {
	return &team.TeamListFilter{
		UserId:    f.UserId,
		Name:      f.Name,
		StartDate: timestamppb.New(f.StartDate),
		EndDate:   timestamppb.New(f.EndDate),
		Deleted:   f.Deleted,
	}
}
