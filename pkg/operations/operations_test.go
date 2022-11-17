package operations

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/models/database"
	"github.com/NEKETSKY/mnemosyne/pkg/auth"
	"testing"
)

func TestCheckAccessAdmin(t *testing.T) {
	ctx := context.Background()
	roles := []database.Role{
		{
			Id:   1,
			Code: string(admin),
		},
	}
	ctx = auth.SetUser(ctx, auth.User{
		Roles: roles,
	})

	tests := []struct {
		operation string
		want      bool
	}{
		{
			operation: "add_mentor",
			want:      true,
		},
		{
			operation: "view_all_students",
			want:      true,
		},
		{
			operation: "edit_own_profile",
			want:      true,
		},
		{
			operation: "absent_operation",
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.operation, func(t *testing.T) {
			if got := CheckAccess(ctx, tt.operation); got != tt.want {
				t.Errorf("CheckAccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckAccessMentor(t *testing.T) {
	ctx := context.Background()
	roles := []database.Role{
		{
			Code: string(mentor),
		},
	}
	ctx = auth.SetUser(ctx, auth.User{
		Id:    1,
		Roles: roles,
	})

	tests := []struct {
		operation string
		want      bool
	}{
		{
			operation: "add_mentor",
			want:      false,
		},
		{
			operation: "view_all_students",
			want:      true,
		},
		{
			operation: "edit_own_profile",
			want:      true,
		},
		{
			operation: "absent_operation",
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.operation, func(t *testing.T) {
			if got := CheckAccess(ctx, tt.operation); got != tt.want {
				t.Errorf("CheckAccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckAccessStudent(t *testing.T) {
	ctx := context.Background()
	roles := []database.Role{
		{
			Code: string(student),
		},
	}
	ctx = auth.SetUser(ctx, auth.User{
		Id:    1,
		Roles: roles,
	})

	tests := []struct {
		operation string
		want      bool
	}{
		{
			operation: "add_mentor",
			want:      false,
		},
		{
			operation: "view_all_students",
			want:      false,
		},
		{
			operation: "edit_own_profile",
			want:      true,
		},
		{
			operation: "absent_operation",
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.operation, func(t *testing.T) {
			if got := CheckAccess(ctx, tt.operation); got != tt.want {
				t.Errorf("CheckAccess() = %v, want %v", got, tt.want)
			}
		})
	}
}
