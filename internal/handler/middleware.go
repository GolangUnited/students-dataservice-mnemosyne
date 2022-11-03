package handler

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/models/auth"
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

			// todo: fill roles by userId from repository
			var roles []string
			ctx = auth.SetUser(ctx, auth.User{
				Id:    userId,
				Roles: roles,
			})
			return ctx, nil
		}
	}

	return nil, status.Error(codes.Unauthenticated, "meta "+userIdArg+" not found")
}
