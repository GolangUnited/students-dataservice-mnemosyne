package group

const (
	GetGroupByIdQuery        = `SELECT * FROM groups WHERE id = $1`
	AddGroupQuery            = `INSERT INTO groups (name, start_date, end_date) VALUES ($1, $2, $3) RETURNING id`
	UpdateGroupByIdQuery     = `UPDATE groups SET name = $1, start_date = $2, end_date = $3, updated_at = $4 WHERE id = $5`
	DeactivateGroupByIdQuery = `UPDATE groups SET deleted = true, updated_at = $1 WHERE id = $2`
	ActivateGroupByIdQuery   = `UPDATE groups SET deleted = false, updated_at = $1 WHERE id = $2`
	AddUserToGroupQuery      = `INSERT INTO user_groups (user_id, group_id) VALUES ($1, $2)`
	DeleteUserFromGroupQuery = `DELETE FROM user_groups WHERE user_id = $1 AND group_id = $2`
)
