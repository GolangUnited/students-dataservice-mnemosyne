package user

const AddUser = `
	INSERT INTO users (last_name, first_name, middle_name, email, language,english_level,photo)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id
`

const GetUserById = `
	SELECT *
	FROM users
	WHERE users.id = $1
`

const GetUserByEmail = `
	SELECT *
	FROM users
	WHERE users.email = $1
`
const GetAllUsers = `
	SELECT *
	FROM users
	`
const UpdateUserById = `
	UPDATE users u 
	SET u.last_name = $1, 
		u.first_name = $2, 
		u.middle_name = $3,
		u.language = $4, 
		u.english_level = $5, 
		u.photo = $6, 
		u.updated_at = $7
	WHERE u.id = $8
`
const ActivateById = `
	UPDATE users u 
	SET u.deleted = false,
		u.updated_at = $2
	WHERE u.id = $1
`
const DeactivateById = `
	UPDATE users u 
	SET u.deleted = true,
	u.updated_at = $2
	WHERE u.id = $1
`
