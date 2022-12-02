package group

import (
	"github.com/NEKETSKY/mnemosyne/pkg/api/group"
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

func ConvertToDB(gr *group.Group) *DB {
	return &DB{
		Id:        gr.GetId(),
		Name:      gr.GetName(),
		StartDate: gr.GetStartDate().AsTime(),
		EndDate:   gr.GetEndDate().AsTime(),
	}
}

func ConvertFromDB(db *DB) *group.Group {
	return &group.Group{
		Id:        db.Id,
		Name:      db.Name,
		StartDate: timestamppb.New(db.StartDate),
		EndDate:   timestamppb.New(db.EndDate),
	}
}
