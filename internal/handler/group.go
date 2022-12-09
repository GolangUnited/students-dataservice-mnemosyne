package handler

import (
	"context"
	modelGroup "github.com/NEKETSKY/mnemosyne/models/database/group"
	"github.com/NEKETSKY/mnemosyne/pkg/api/common"
	"github.com/NEKETSKY/mnemosyne/pkg/api/group"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetGroup(ctx context.Context, in *group.Id) (*group.GroupResponse, error) {
	groupDB, err := h.services.GetGroup(ctx, in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return groupDB.ToProtoResponse(), nil
}

func (h *Handler) GetGroups(ctx context.Context, in *group.GroupsRequest) (*group.Groups, error) {
	groupsProto := &group.Groups{}
	groupFilter := &modelGroup.Filter{}
	groupFilter.FromProto(in.GetFilter())
	groupsDB, err := h.services.GetGroups(ctx, groupFilter)
	if err != nil {
		return groupsProto, status.Error(codes.Internal, err.Error())
	}

	groupsResponse := make([]*group.GroupResponse, 0, len(groupsDB))
	for _, groupDB := range groupsDB {
		groupsResponse = append(groupsResponse, groupDB.ToProtoResponse())
	}
	groupsProto.Groups = groupsResponse

	return groupsProto, nil
}

func (h *Handler) CreateGroup(ctx context.Context, in *group.GroupRequest) (groupIdProto *group.Id, err error) {
	groupIdProto = &group.Id{}
	groupDB := &modelGroup.DB{}
	groupDB.FromProtoRequest(in)
	groupId, err := h.services.Mnemosyne.CreateGroup(ctx, groupDB)
	if err != nil {
		return groupIdProto, status.Error(codes.Internal, err.Error())
	}
	groupIdProto.Id = groupId

	return
}

func (h *Handler) UpdateGroup(ctx context.Context, in *group.GroupRequest) (*common.Empty, error) {
	groupDB := &modelGroup.DB{}
	groupDB.FromProtoRequest(in)
	err := h.services.Mnemosyne.UpdateGroup(ctx, groupDB)
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) DeactivateGroup(ctx context.Context, in *group.Id) (*common.Empty, error) {
	err := h.services.Mnemosyne.DeactivateGroup(ctx, in.GetId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) ActivateGroup(ctx context.Context, in *group.Id) (*common.Empty, error) {
	err := h.services.Mnemosyne.ActivateGroup(ctx, in.GetId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) AddUserToGroup(ctx context.Context, in *group.UserGroupRequest) (*common.Empty, error) {
	err := h.services.Mnemosyne.AddUserToGroup(ctx, in.GetUserId(), in.GetGroupId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) DeleteUserFromGroup(ctx context.Context, in *group.UserGroupRequest) (*common.Empty, error) {
	err := h.services.Mnemosyne.DeleteUserFromGroup(ctx, in.GetUserId(), in.GetGroupId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}
