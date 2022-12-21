package handler

import (
	"context"
	modelProject "github.com/NEKETSKY/mnemosyne/models/database/project"
	"github.com/NEKETSKY/mnemosyne/pkg/api/common"
	"github.com/NEKETSKY/mnemosyne/pkg/api/project"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetProject(ctx context.Context, in *project.Id) (*project.ProjectResponse, error) {
	projectDB, err := h.services.GetProject(ctx, in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return projectDB.ToProtoResponse(), nil
}

func (h *Handler) GetProjects(ctx context.Context, in *project.ProjectsRequest) (*project.Projects, error) {
	projectsProto := &project.Projects{}
	projectFilter := &modelProject.Filter{}
	projectFilter.FromProto(in.GetFilter())
	projectsDB, err := h.services.GetProjects(ctx, projectFilter)
	if err != nil {
		return projectsProto, status.Error(codes.Internal, err.Error())
	}

	projectsResponse := make([]*project.ProjectResponse, 0, len(projectsDB))
	for _, projectDB := range projectsDB {
		projectsResponse = append(projectsResponse, projectDB.ToProtoResponse())
	}
	projectsProto.Projects = projectsResponse

	return projectsProto, nil
}

func (h *Handler) CreateProject(ctx context.Context, in *project.ProjectRequest) (projectIdProto *project.Id, err error) {
	projectIdProto = &project.Id{}
	projectDB := &modelProject.DB{}
	projectDB.FromProtoRequest(in)
	projectId, err := h.services.CreateProject(ctx, projectDB)
	if err != nil {
		return projectIdProto, status.Error(codes.Internal, err.Error())
	}
	projectIdProto.Id = projectId

	return
}

func (h *Handler) UpdateProject(ctx context.Context, in *project.ProjectRequest) (*common.Empty, error) {
	projectDB := &modelProject.DB{}
	projectDB.FromProtoRequest(in)
	err := h.services.UpdateProject(ctx, projectDB)
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) DeactivateProject(ctx context.Context, in *project.Id) (*common.Empty, error) {
	err := h.services.DeactivateProject(ctx, in.GetId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) ActivateProject(ctx context.Context, in *project.Id) (*common.Empty, error) {
	err := h.services.ActivateProject(ctx, in.GetId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}
