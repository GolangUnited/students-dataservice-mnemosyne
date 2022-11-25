package interview

const GetInterviewByIdQuery = `
	SELECT *
	FROM interview
	WHERE id = $1
`
const AddInterviewQuery = `
	INSERT INTO interview (
	   interviewer_id, 
	   student_id, 
	   interview_date, 
	   grade,
	   subjective_rating,
	   notes, 
	   determined_english_level, 
	   main_task, 
	   question
	   )
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id
`

const UpdateInterviewByIdQuery = `
	UPDATE interview
	SET interviewer_id = $1,
		student_id = $2,
		interview_date = $3,
		grade = $4,
		subjective_rating = $5,
		notes = $6,
		determined_english_level = $7,
		main_task = $8,
		question = $9,
		updated_at = $10
	WHERE id = $11
`
const DeactivateByIdQuery = `
	UPDATE interview 
	SET deleted = true,
	updated_at = $2
	WHERE id = $1
`
const ActivateByIdQuery = `
	UPDATE interview 
	SET deleted = false,
	updated_at = $2
	WHERE id = $1
`
