package role

import (
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/role"
)

type DB struct {
	Id   uint32 `db:"id"`
	Code string `db:"code"`
}

func (d *DB) FromProto(rr *role.Role) {
	if rr == nil {
		return
	}
	d.Id = rr.GetId()
	d.Code = rr.GetCode()
}

func (d *DB) ToProto() *role.Role {
	return &role.Role{
		Id:   d.Id,
		Code: d.Code,
	}
}
