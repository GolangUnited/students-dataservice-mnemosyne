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
	UPDATE users
	SET last_name = $1, 
		first_name = $2, 
		middle_name = $3,
		language = $4, 
		english_level = $5, 
		photo = $6, 
		updated_at = $7
	WHERE id = $8
`
const ActivateById = `
	UPDATE users 
	SET deleted = false,
		updated_at = $2
	WHERE id = $1
`
const DeactivateById = `
	UPDATE users 
	SET deleted = true,
	updated_at = $2
	WHERE id = $1
`
