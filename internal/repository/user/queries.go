package user

const AddUser = `
	INSERT INTO users (last_name, first_name, middle_name, email, language,english_level,photo)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id
`
const AddContactById = `
insert into сontacts (user_id, telegram, discord, communication_channel) values ($1, $2, $3, $4)
`
const AddResumeById = `
insert into resume (user_id, experience, uploaded_resume, country, city, time_zone, mentors_note) values ($1, $2, $3, $4, $5, $6, $7)
`
const AddRoleStudent = `
insert into user_role (user_id) values ($1)`

const AliveUsers = `
where u.deleted = false`

const SelectByRole = `
left join user_role ur on ur.user_id = u.id
inner join roles ro on ro.id = ur.role_id and ro.code=`
const SelectByGroup = `
left join student_groups sg on sg.student_id = u.id
inner join groups g on g.id = sg.group_id and g.id =`
const SelectByTeam = `
left join student_teams st on st.student_id = u.id
inner join teams t on t.id = st.team_id and t.id =`

const GetUserById = `
	WHERE u.id = $1 and u.deleted = false
`
const GetUserByEmail = `
	WHERE u.email = $1 and u.deleted = false
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
case when u.photo is null then '' 
    else u.photo end as photo,
case when c.telegram is null then '' 
    else c.telegram end as telegram,
case when c.discord is null then '' 
    else c.discord end as discord,
case when c.communication_channel is null then '' 
    else c.communication_channel end as communication_channel,
case when r.experience is null then ''
    else r.experience end as experience,
case when r.uploaded_resume is null then ''
    else r.uploaded_resume end as uploaded_resume,
case when r.country is null then '' 
    else r.country end as country,
case when r.city is null then '' 
    else r.city end as city,
case when r.time_zone is null then '' 
    else r.time_zone end as time_zone,
case when r.mentors_note is null then '' 
    else r.mentors_note end as mentors_note
FROM
users u
left join сontacts c on c.user_id = u.id
left join resume r on r.user_id = u.id
`
const GetUsersWithContacts = `
SELECT
u.id as id,
u.last_name as last_name,
u.first_name as first_name,
u.middle_name as middle_name,
u.email as email,
u."language" as "language",    
u.english_level as english_level,  
case when u.photo is null then '' 
    else u.photo end as photo,
case when c.telegram is null then '' 
    else c.telegram end as telegram,
case when c.discord is null then '' 
    else c.discord end as discord,
case when c.communication_channel is null then '' 
    else c.communication_channel end as communication_channel,
'' as experience,
'' as uploaded_resume,
'' as country,
'' as city,
'' as time_zone,
'' as mentors_note
FROM
users u
left join сontacts c on c.user_id = u.id
`
const GetUsersWithResume = `
SELECT
u.id as id,
u.last_name as last_name,
u.first_name as first_name,
u.middle_name as middle_name,
u.email as email,
u."language" as "language",    
u.english_level as english_level,  
case when u.photo is null then '' 
    else u.photo end as photo,
'' as telegram,
'' as discord,
'' as communication_channel,
case when r.experience is null then ''
    else r.experience end as experience,
case when r.uploaded_resume is null then ''
    else r.uploaded_resume end as uploaded_resume,
case when r.country is null then '' 
    else r.country end as country,
case when r.city is null then '' 
    else r.city end as city,
case when r.time_zone is null then '' 
    else r.time_zone end as time_zone,
case when r.mentors_note is null then '' 
    else r.mentors_note end as mentors_note
FROM
users u
left join resume r on r.user_id = u.id
`
const GetUsers = `
SELECT
u.id as id,
u.last_name as last_name,
u.first_name as first_name,
u.middle_name as middle_name,
u.email as email,
u."language" as "language",    
u.english_level as english_level,  
case when u.photo is null then '' 
    else u.photo end as photo,
'' as telegram,
'' as discord,
'' as communication_channel,
'' as experience,
'' as uploaded_resume,
'' as country,
'' as city,
'' as time_zone,
'' as mentors_note
FROM
users u
left join сontacts c on c.user_id = u.id
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
const GetEmailById = `
select users.email from users where users.id = $1 and users.deleted = false
`
const UpdateContactById = `
	UPDATE сontacts 
	SET telegram = $1, 
		discord = $2, 
		communication_channel = $3,
		updated_at = $4
	WHERE user_id = $5
`

const UpdateResumeById = `
	UPDATE resume
	SET experience = $1, 
		uploaded_resume = $2, 
		country = $3,
		city = $4,
		time_zone = $5,
		mentors_note = $6,
		updated_at = $7
	WHERE user_id = $8
`
const GetContactById = `
select c.user_id, c.telegram, c.discord, c.communication_channel
from сontacts c
where c.user_id = $1 and c.deleted = false
`
const GetResumeById = `
select r.user_id, r.experience, r.uploaded_resume, r.country, r.city, r.time_zone, r.mentors_note
from resume r
where r.user_id =$1 and r.deleted = false
`
const OrderAsc = `
order by u.id asc
`
