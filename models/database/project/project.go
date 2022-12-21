package project

import (
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/project"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type DB struct {
	Id          uint32    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	GitUrl      string    `db:"git_url"`
	TeamId      uint32    `db:"team_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Deleted     bool      `db:"deleted"`
}

func (d *DB) FromProtoRequest(pr *project.ProjectRequest) {
	if pr == nil {
		return
	}
	d.Id = pr.GetId()
	d.Name = pr.GetName()
	d.Description = pr.GetDescription()
	d.GitUrl = pr.GetGitUrl()
	d.TeamId = pr.GetTeamId()
}

func (d *DB) ToProtoRequest() *project.ProjectRequest {
	return &project.ProjectRequest{
		Id:          d.Id,
		Name:        d.Name,
		Description: d.Description,
		GitUrl:      d.GitUrl,
		TeamId:      d.TeamId,
	}
}

func (d *DB) FromProtoResponse(pr *project.ProjectResponse) {
	if pr == nil {
		return
	}
	d.Id = pr.GetId()
	d.Name = pr.GetName()
	d.Description = pr.GetDescription()
	d.GitUrl = pr.GetGitUrl()
	d.TeamId = pr.TeamId
	d.CreatedAt = pr.GetCreatedAt().AsTime()
	d.UpdatedAt = pr.GetUpdatedAt().AsTime()
	d.Deleted = pr.GetDeleted()
}

func (d *DB) ToProtoResponse() *project.ProjectResponse {
	return &project.ProjectResponse{
		Id:          d.Id,
		Name:        d.Name,
		Description: d.Description,
		GitUrl:      d.GitUrl,
		TeamId:      d.TeamId,
		CreatedAt:   timestamppb.New(d.CreatedAt),
		UpdatedAt:   timestamppb.New(d.UpdatedAt),
		Deleted:     d.Deleted,
	}
}
