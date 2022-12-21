package operations

import (
	"context"

	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/auth"
)

type role string

var (
	admin          role = "admin"
	mentor         role = "mentor"
	student        role = "student"
	roleOperations      = map[role][]interface{}{
		admin: {
			mentor,
			student,
			"add_mentor",
			"delete_mentor",
			"view_deleted",
		},
		mentor: {
			student,
			"view_all_users",
			"delete_student",
			"activate",
		},
		student: {
			"view_all_mentors",
			"add_student",
			"view_own_profile",
			"edit_own_profile",
		},
	}
)

func init() {
	// fill role operations inner roles
	for r := range roleOperations {
		setOps := make(map[string]struct{}, len(roleOperations))
		innerOps := innerOperations(r)
		for _, innerOp := range innerOps {
			setOps[innerOp] = struct{}{}
		}
		roleOperations[r] = make([]interface{}, 0, len(setOps))
		for sop := range setOps {
			roleOperations[r] = append(roleOperations[r], sop)
		}
	}
}

// innerOperations find inner role operations recursive
func innerOperations(r role) []string {
	innerOps := make([]string, 0, len(roleOperations[r]))
	for _, iop := range roleOperations[r] {
		if innerRole, ok := iop.(role); ok {
			innerOps = append(innerOps, innerOperations(innerRole)...)
		} else {
			innerOps = append(innerOps, iop.(string))
		}
	}

	return innerOps
}

// operations return set of all user operations
func operations(ctx context.Context) map[string]struct{} {
	user := auth.GetUser(ctx)
	ops := make(map[string]struct{}, len(roleOperations[admin]))
	for _, ur := range user.Roles {
		if userOperations, ok := roleOperations[role(ur.Code)]; ok {
			for _, uop := range userOperations {
				ops[uop.(string)] = struct{}{}
			}
		}
	}

	return ops
}

// CheckAccess verify access operation for user
func CheckAccess(ctx context.Context, operation string) bool {
	ops := operations(ctx)
	if _, ok := ops[operation]; ok {
		return true
	}

	return false
}
