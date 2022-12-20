package team

const (
	GetTeamByIdQuery    = `SELECT * FROM teams WHERE id = $1`
	AddTeamQuery        = `INSERT INTO teams (name, start_date, end_date, group_id, mentor_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	UpdateTeamByIdQuery = `UPDATE teams 
								SET name = $1, 
								    start_date = $2, 
								    end_date = $3, 
								    updated_at = $4, 
								    group_id = $5, 
								    mentor_id = $6 
								WHERE id = $7`
	DeactivateTeamByIdQuery = `UPDATE teams SET deleted = true, updated_at = $1 WHERE id = $2`
	ActivateTeamByIdQuery   = `UPDATE teams SET deleted = false, updated_at = $1 WHERE id = $2`
	AddUserToTeamQuery      = `INSERT INTO user_teams (user_id, team_id) VALUES ($1, $2)`
	DeleteUserFromTeamQuery = `DELETE FROM user_teams WHERE user_id = $1 AND team_id = $2`
)
