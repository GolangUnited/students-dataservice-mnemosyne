package role

import (
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/role"
)

type Filter struct {
	Code string
}

func (f *Filter) FromProto(rf *role.Filter) {
	if rf == nil {
		return
	}
	f.Code = rf.GetCode()
}

func (f *Filter) ToProto() *role.Filter {
	return &role.Filter{
		Code: &f.Code,
	}
}
