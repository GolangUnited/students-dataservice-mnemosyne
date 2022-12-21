package project

import (
	"github.com/NEKETSKY/mnemosyne/pkg/api/project"
)

type Filter struct {
	Name    string
	TeamId  uint32
	Deleted bool
}

func (f *Filter) FromProto(pf *project.Filter) {
	if pf == nil {
		return
	}
	f.TeamId = pf.GetTeamId()
	f.Name = pf.GetName()
	f.Deleted = pf.GetDeleted()
}

func (f *Filter) ToProto() *project.Filter {
	return &project.Filter{
		TeamId:  f.TeamId,
		Name:    &f.Name,
		Deleted: f.Deleted,
	}
}
