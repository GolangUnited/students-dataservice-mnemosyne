package mnemosyne

import (
	"context"
	modelRole "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/role"
)

func (s *Service) GetRoles(ctx context.Context, roleFilter *modelRole.Filter) ([]*modelRole.DB, error) {
	return s.reposRole.GetRoles(ctx, roleFilter)
}
func (s *Service) CreateRole(ctx context.Context, roleDB *modelRole.DB) (uint32, error) {
	return s.reposRole.AddRole(ctx, roleDB)
}
func (s *Service) DeleteRole(ctx context.Context, roleId uint32) error {
	return s.reposRole.DeleteRole(ctx, roleId)
}
func (s *Service) AddUserToRole(ctx context.Context, userId, roleId uint32) error {
	return s.reposRole.AddUserToRole(ctx, userId, roleId)
}
func (s *Service) DeleteUserFromRole(ctx context.Context, userId, roleId uint32) error {
	return s.reposRole.DeleteUserFromRole(ctx, userId, roleId)
}
