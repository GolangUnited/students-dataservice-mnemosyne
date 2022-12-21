package mnemosyne

import (
	"context"
	modelProject "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/project"
)

func (s *Service) GetProject(ctx context.Context, projectId uint32) (*modelProject.DB, error) {
	return s.reposProject.GetProjectById(ctx, projectId)
}

func (s *Service) GetProjects(ctx context.Context, groupFilter *modelProject.Filter) ([]*modelProject.DB, error) {
	return s.reposProject.GetProjects(ctx, groupFilter)
}

func (s *Service) CreateProject(ctx context.Context, projectDB *modelProject.DB) (uint32, error) {
	return s.reposProject.AddProject(ctx, projectDB)
}

func (s *Service) UpdateProject(ctx context.Context, projectDB *modelProject.DB) error {
	return s.reposProject.UpdateProject(ctx, projectDB)
}

func (s *Service) DeactivateProject(ctx context.Context, projectId uint32) error {
	return s.reposProject.DeactivateProject(ctx, projectId)
}

func (s *Service) ActivateProject(ctx context.Context, projectId uint32) error {
	return s.reposProject.ActivateProject(ctx, projectId)
}
