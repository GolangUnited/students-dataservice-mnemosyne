package user

const WithoutDeleted = `
where u.deleted = false`

const SelectStudents = `
left join user_role ur on ur.user_id = u.id
inner join roles ro on ro.id = ur.role_id and ro.code='student'
`
const SelectMentors = `
left join user_role ur on ur.user_id = u.id
inner join roles ro on ro.id = ur.role_id and ro.code='mentor'
`

const AddUser = `
	INSERT INTO users (last_name, first_name, middle_name, email, language,english_level,photo)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id
`
const AddContactById = `
insert into contacts (user_id, telegram, discord, communication_channel) values ($1, $2, $3, $4)
`
const AddResumeById = `
insert into resume (user_id, experience, uploaded_resume, country, city, time_zone, mentors_note) values ($1, $2, $3, $4, $5, $6, $7)
`
const AddRoleStudent = `
insert into user_role (user_id, role_id) values ($1, 3)
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
const GetUsersFull = `
SELECT
u.id as id,
u.last_name as last_name,
u.first_name as first_name,
u.middle_name as middle_name,
u.email as email,
u."language" as "language",    
u.english_level as english_level,  
u.photo as photo,
case when c.telegram is null then '' end as telegram,
case when c.discord is null then '' end as discord,
case when c.communication_channel is null then '' end as communication_channel,
case when r.experience is null then '' end as experience,
case when r.uploaded_resume is null then '' end as uploaded_resume,
case when r.country is null then '' end as country,
case when r.city is null then '' end as city,
case when r.time_zone is null then '' end as time_zone,
case when r.mentors_note is null then '' end as mentors_note
FROM
users u
left join сontacts c on c.user_id = u.id
left join resume r on r.user_id = u.id
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
