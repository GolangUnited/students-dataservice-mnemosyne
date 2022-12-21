package handler

import (
	"context"
	modelRole "github.com/GolangUnited/students-dataservice-mnemosyne/models/database/role"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/common"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/role"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetRoles(ctx context.Context, in *role.RolesRequest) (*role.Roles, error) {
	rolesProto := &role.Roles{}
	roleFilter := &modelRole.Filter{}
	roleFilter.FromProto(in.GetFilter())
	rolesDB, err := h.services.GetRoles(ctx, roleFilter)
	if err != nil {
		return rolesProto, status.Error(codes.Internal, err.Error())
	}

	roles := make([]*role.Role, 0, len(rolesDB))
	for _, roleDB := range rolesDB {
		roles = append(roles, roleDB.ToProto())
	}
	rolesProto.Roles = roles

	return rolesProto, nil
}

func (h *Handler) CreateRole(ctx context.Context, in *role.Role) (roleIdProto *role.Id, err error) {
	roleIdProto = &role.Id{}
	roleDB := &modelRole.DB{}
	roleDB.FromProto(in)
	roleId, err := h.services.CreateRole(ctx, roleDB)
	if err != nil {
		return roleIdProto, status.Error(codes.Internal, err.Error())
	}
	roleIdProto.Id = roleId

	return
}

func (h *Handler) DeleteRole(ctx context.Context, in *role.Id) (*common.Empty, error) {
	err := h.services.DeleteRole(ctx, in.GetId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) AddUserToRole(ctx context.Context, in *role.UserRoleRequest) (*common.Empty, error) {
	err := h.services.AddUserToRole(ctx, in.GetUserId(), in.GetRoleId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}

func (h *Handler) DeleteUserFromRole(ctx context.Context, in *role.UserRoleRequest) (*common.Empty, error) {
	err := h.services.DeleteUserFromRole(ctx, in.GetUserId(), in.GetRoleId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}

	return emptyProto, nil
}
