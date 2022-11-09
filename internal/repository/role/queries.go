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
