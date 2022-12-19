package team

import (
	"github.com/NEKETSKY/mnemosyne/pkg/api/team"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type DB struct {
	Id        uint32    `json:"id" db:"id"`
	GroupId   uint32    `json:"group_id" db:"group_id"`
	MentorId  uint32    `json:"mentor_id" db:"mentor_id"`
	Name      string    `json:"name" db:"name"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
	Deleted   bool      `json:"-" db:"deleted"`
}

func (d *DB) FromProtoRequest(teamRequest *team.TeamRequest) {
	d.Id = teamRequest.GetId()
	d.Name = teamRequest.GetName()
	d.GroupId = teamRequest.GetGroupId()
	d.MentorId = teamRequest.GetMentorId()
	d.StartDate = teamRequest.GetStartDate().AsTime()
	d.EndDate = teamRequest.GetEndDate().AsTime()
}

func (d *DB) ToProtoResponse() *team.TeamResponse {
	return &team.TeamResponse{
		Id:        d.Id,
		Name:      d.Name,
		GroupId:   d.GroupId,
		MentorId:  d.MentorId,
		StartDate: timestamppb.New(d.StartDate),
		EndDate:   timestamppb.New(d.EndDate),
		CreatedAt: timestamppb.New(d.CreatedAt),
		UpdatedAt: timestamppb.New(d.UpdatedAt),
		Deleted:   d.Deleted,
	}
}
