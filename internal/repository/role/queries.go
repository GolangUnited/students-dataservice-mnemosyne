package role

const AllRolesQuery = `
	SELECT r.id, r.code FROM roles r
`

const RolesByUserIdQuery = `
	SELECT r.id, r.code FROM roles r 
		JOIN user_role ur ON r.id = ur.role_id 
	WHERE ur.user_id = $1
`

const DeleteRoleForUserQuery = `
	DELETE FROM user_role
	WHERE user_id = $1 AND role_id = (SELECT id FROM roles WHERE code = $2)
`

const AddRoleForUserQuery = `
	INSERT INTO user_role (user_id, role_id)
	VALUES ($1, (SELECT id FROM roles WHERE code = $2))
`

const (
	AddRoleQuery            = `INSERT INTO roles (code) VALUES ($1) RETURNING id`
	DeleteByIdQuery         = `DELETE FROM roles WHERE id = $1`
	AddUserToRoleQuery      = `INSERT INTO user_role (user_id, role_id) VALUES ($1, $2)`
	DeleteUserFromRoleQuery = `DELETE FROM user_role WHERE user_id = $1 AND role_id = $2`
)
