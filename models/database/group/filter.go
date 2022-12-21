package group

import (
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/group"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Filter struct {
	UserId    uint32
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Deleted   bool
}

func (f *Filter) FromProto(gf *group.Filter) {
	if gf == nil {
		return
	}
	f.UserId = gf.GetUserId()
	f.Name = gf.GetName()
	f.StartDate = gf.GetStartDate().AsTime()
	f.EndDate = gf.GetEndDate().AsTime()
	f.Deleted = gf.GetDeleted()
}

func (f *Filter) ToProto() *group.Filter {
	return &group.Filter{
		UserId:    f.UserId,
		Name:      &f.Name,
		StartDate: timestamppb.New(f.StartDate),
		EndDate:   timestamppb.New(f.EndDate),
		Deleted:   f.Deleted,
	}
}
