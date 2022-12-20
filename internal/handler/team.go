package handler

import (
	"context"
	modelTeam "github.com/NEKETSKY/mnemosyne/models/database/team"
	"github.com/NEKETSKY/mnemosyne/pkg/api/common"
	"github.com/NEKETSKY/mnemosyne/pkg/api/team"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetTeam(ctx context.Context, in *team.Id) (resp *team.TeamResponse, err error) {
	teamDb, err := h.services.GetTeam(ctx, in.GetId())
	if err != nil {
		return &team.TeamResponse{}, status.Error(codes.Internal, err.Error())
	}

	return teamDb.ToProtoResponse(), nil
}

func (h *Handler) GetTeams(ctx context.Context, in *team.TeamListFilter) (*team.Teams, error) {
	teamsProto := &team.Teams{}
	teamFilter := &modelTeam.Filter{}
	teamFilter.FromProto(in)
	teamsDB, err := h.services.GetTeams(ctx, teamFilter)
	if err != nil {
		return teamsProto, status.Error(codes.Internal, err.Error())
	}

	teamsResponse := make([]*team.TeamResponse, 0, len(teamsDB))
	for _, teamDB := range teamsDB {
		teamsResponse = append(teamsResponse, teamDB.ToProtoResponse())
	}
	teamsProto.Teams = teamsResponse

	return teamsProto, nil
}

func (h *Handler) CreateTeam(ctx context.Context, in *team.TeamRequest) (teamIdProto *team.Id, err error) {
	teamIdProto = &team.Id{}
	teamDB := &modelTeam.DB{}
	teamDB.FromProtoRequest(in)
	teamId, err := h.services.Mnemosyne.CreateTeam(ctx, teamDB)
	if err != nil {
		return teamIdProto, status.Error(codes.Internal, err.Error())
	}
	teamIdProto.Id = teamId

	return
}

func (h *Handler) UpdateTeam(ctx context.Context, in *team.TeamRequest) (*common.Empty, error) {
	teamDB := &modelTeam.DB{}
	teamDB.FromProtoRequest(in)
	err := h.services.Mnemosyne.UpdateTeam(ctx, teamDB)
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) DeactivateTeam(ctx context.Context, in *team.Id) (*common.Empty, error) {
	err := h.services.Mnemosyne.DeactivateTeam(ctx, in.GetId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) ActivateTeam(ctx context.Context, in *team.Id) (*common.Empty, error) {
	err := h.services.Mnemosyne.ActivateTeam(ctx, in.GetId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) AddUserToTeam(ctx context.Context, in *team.UserTeamRequest) (*common.Empty, error) {
	err := h.services.Mnemosyne.AddUserToTeam(ctx, in.GetUserId(), in.GetTeamId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) DeleteUserFromTeam(ctx context.Context, in *team.UserTeamRequest) (*common.Empty, error) {
	err := h.services.Mnemosyne.DeleteUserFromTeam(ctx, in.GetUserId(), in.GetTeamId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}
