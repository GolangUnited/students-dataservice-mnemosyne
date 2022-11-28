package certificate

const AddCertificate = `
	INSERT INTO certificates (certificate_id, user_id, issue_date, expire_date)
	VALUES ($1, $2, $3, $4)
	RETURNING id
`

const GetAllCertificates = `
	SELECT *
	FROM certificates
	`
const UpdateCertificatesById = `
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
