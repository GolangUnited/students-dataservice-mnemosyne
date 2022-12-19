package handler

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/pkg/auth"
	"github.com/NEKETSKY/mnemosyne/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
)

var (
	userIdArg = "userid"
)

func (h *Handler) Auth(ctx context.Context) (context.Context, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userIdMD, ok := md[userIdArg]; ok {
			userId, err := strconv.ParseUint(userIdMD[0], 10, 0)
			if err != nil || userId <= 0 {
				return nil, status.Error(codes.InvalidArgument, userIdArg+" not correct")
			}

			userRoles, err := h.services.GetUserRoles(ctx, int(userId))
			if err != nil {
				logger.Infof("get user (%d) roles error: %s", userId, err.Error())
				return nil, status.Error(codes.PermissionDenied, err.Error())
			}
			if len(userRoles) == 0 {
				logger.Infof("user (%d) not found", userId)
				return nil, status.Error(codes.PermissionDenied, "user not found")
			}

			ctx = auth.SetUser(ctx, auth.User{
				Id:    int(userId),
				Roles: userRoles,
			})

			return ctx, nil
		}
	}

	return nil, status.Error(codes.Unauthenticated, "meta "+userIdArg+" not found")
}
