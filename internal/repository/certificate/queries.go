package certificate

const AddCertificate = `
	INSERT INTO certificates (user_id, issue_date, expire_date)
	VALUES ($1, $2, $3)
	RETURNING id
`

const GetCertificateById = `
	SELECT *
	FROM certificates
	WHERE id = $1 and deleted = false
	`

const GetAllCertificates = `
	SELECT id, user_id, issue_date, expire_date
	FROM certificates
	WHERE user_id = $1
	`
const UpdateCertificateById = `
	UPDATE certificates
	SET user_id = $1, 
		issue_date = $2, 
		expire_date = $3
	WHERE id = $4
`
const ActivateById = `
	UPDATE certificates 
	SET deleted = false,
		updated_at = $2
	WHERE id = $1
`
const DeactivateById = `
	UPDATE certificates 
	SET deleted = true,
	updated_at = $2
	WHERE id = $1
`
