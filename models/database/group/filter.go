package group

import (
	"github.com/NEKETSKY/mnemosyne/pkg/api/group"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Filter struct {
	UserId    uint32
	Name      string
	StartDate time.Time
	EndDate   time.Time
}

func ConvertToFilter(gf *group.Filter) *Filter {
	return &Filter{
		UserId:    gf.GetUserId(),
		Name:      gf.GetName(),
		StartDate: gf.GetStartDate().AsTime(),
		EndDate:   gf.GetEndDate().AsTime(),
	}
}

func ConvertFromFilter(f *Filter) *group.Filter {
	return &group.Filter{
		UserId:    f.UserId,
		Name:      &f.Name,
		StartDate: timestamppb.New(f.StartDate),
		EndDate:   timestamppb.New(f.EndDate),
	}
}
