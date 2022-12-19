package auth

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/models/database/role"
)

type User struct {
	Id    int
	Roles []role.DB
}

type userKey struct{}

// SetUser adds the request User to the context
func SetUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userKey{}, user)
}

// GetUser returns user in the context
func GetUser(ctx context.Context) User {
	return getContextUser(ctx)
}

func getContextUser(ctx context.Context) User {
	contextUser, _ := ctx.Value(userKey{}).(User)
	return contextUser
}
