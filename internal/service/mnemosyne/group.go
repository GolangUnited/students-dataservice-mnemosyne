package mnemosyne

import (
	"context"
	modelGroup "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/group"
)

func (s *Service) GetGroup(ctx context.Context, groupId uint32) (*modelGroup.DB, error) {
	return s.reposGroup.GetGroupById(ctx, groupId)
}

func (s *Service) GetGroups(ctx context.Context, groupFilter *modelGroup.Filter) ([]*modelGroup.DB, error) {
	return s.reposGroup.GetGroups(ctx, groupFilter)
}

func (s *Service) CreateGroup(ctx context.Context, groupDB *modelGroup.DB) (uint32, error) {
	return s.reposGroup.AddGroup(ctx, groupDB)
}

func (s *Service) UpdateGroup(ctx context.Context, groupDB *modelGroup.DB) error {
	return s.reposGroup.UpdateGroup(ctx, groupDB)
}

func (s *Service) DeactivateGroup(ctx context.Context, groupId uint32) error {
	return s.reposGroup.DeactivateGroup(ctx, groupId)
}

func (s *Service) ActivateGroup(ctx context.Context, groupId uint32) error {
	return s.reposGroup.ActivateGroup(ctx, groupId)
}

func (s *Service) AddUserToGroup(ctx context.Context, userId, groupId uint32) error {
	return s.reposGroup.AddUserToGroup(ctx, userId, groupId)
}

func (s *Service) DeleteUserFromGroup(ctx context.Context, userId, groupId uint32) error {
	return s.reposGroup.DeleteUserFromGroup(ctx, userId, groupId)
}
