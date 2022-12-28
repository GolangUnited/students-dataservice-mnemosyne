package project

const (
	GetProjectByIdQuery        = `SELECT * FROM projects WHERE id = $1`
	AddProjectQuery            = `INSERT INTO projects (name, description, git_url, team_id) VALUES ($1, $2, $3, $4) RETURNING id`
	AddProjectQueryWOTeam      = `INSERT INTO projects (name, description, git_url, team_id) VALUES ($1, $2, $3, null) RETURNING id`
	UpdateProjectByIdQuery     = `UPDATE projects SET name = $1, description = $2, git_url = $3, team_id = $4, updated_at = $5 WHERE id = $6`
	DeactivateProjectByIdQuery = `UPDATE projects SET deleted = true, updated_at = $1 WHERE id = $2`
	ActivateProjectByIdQuery   = `UPDATE projects SET deleted = false, updated_at = $1 WHERE id = $2`
)
