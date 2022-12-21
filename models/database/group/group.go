package group

import (
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/group"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type DB struct {
	Id        uint32    `db:"id"`
	Name      string    `db:"name"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted   bool      `db:"deleted"`
}

func (d *DB) FromProtoRequest(gr *group.GroupRequest) {
	if gr == nil {
		return
	}
	d.Id = gr.GetId()
	d.Name = gr.GetName()
	d.StartDate = gr.GetStartDate().AsTime()
	d.EndDate = gr.GetEndDate().AsTime()
}

func (d *DB) ToProtoRequest() *group.GroupRequest {
	return &group.GroupRequest{
		Id:        d.Id,
		Name:      d.Name,
		StartDate: timestamppb.New(d.StartDate),
		EndDate:   timestamppb.New(d.EndDate),
	}
}

func (d *DB) FromProtoResponse(gr *group.GroupResponse) {
	if gr == nil {
		return
	}
	d.Id = gr.GetId()
	d.Name = gr.GetName()
	d.StartDate = gr.GetStartDate().AsTime()
	d.EndDate = gr.GetEndDate().AsTime()
	d.CreatedAt = gr.GetCreatedAt().AsTime()
	d.UpdatedAt = gr.GetUpdatedAt().AsTime()
	d.Deleted = gr.GetDeleted()
}

func (d *DB) ToProtoResponse() *group.GroupResponse {
	return &group.GroupResponse{
		Id:        d.Id,
		Name:      d.Name,
		StartDate: timestamppb.New(d.StartDate),
		EndDate:   timestamppb.New(d.EndDate),
		CreatedAt: timestamppb.New(d.CreatedAt),
		UpdatedAt: timestamppb.New(d.UpdatedAt),
		Deleted:   d.Deleted,
	}
}
