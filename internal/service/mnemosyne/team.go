package mnemosyne

import (
	"context"
	modelTeam "github.com/NEKETSKY/mnemosyne/models/database/team"
	//"time"
)

func (s *Service) GetTeam(ctx context.Context, teamId uint32) (*modelTeam.DB, error) {
	return s.reposTeam.GetTeamById(ctx, teamId)
}

func (s *Service) GetTeams(ctx context.Context, teamFilter *modelTeam.Filter) ([]*modelTeam.DB, error) {
	return s.reposTeam.GetTeams(ctx, teamFilter)
}

func (s *Service) CreateTeam(ctx context.Context, teamDB *modelTeam.DB) (uint32, error) {
	return s.reposTeam.AddTeam(ctx, teamDB)
}

func (s *Service) UpdateTeam(ctx context.Context, teamDB *modelTeam.DB) error {
	return s.reposTeam.UpdateTeam(ctx, teamDB)
}

func (s *Service) DeactivateTeam(ctx context.Context, teamId uint32) error {
	return s.reposTeam.DeactivateTeam(ctx, teamId)
}

func (s *Service) ActivateTeam(ctx context.Context, teamId uint32) error {
	return s.reposTeam.ActivateTeam(ctx, teamId)
}

func (s *Service) AddUserToTeam(ctx context.Context, userId, teamId uint32) error {
	return s.reposTeam.AddUserToTeam(ctx, userId, teamId)
}

func (s *Service) DeleteUserFromTeam(ctx context.Context, userId, teamId uint32) error {
	return s.reposTeam.DeleteUserFromTeam(ctx, userId, teamId)
}
