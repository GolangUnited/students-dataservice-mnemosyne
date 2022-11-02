begin;
create table if not exists interview (
id serial primary key,
interviewer_id integer references users(id),
student_id integer references users(id),
interview_date timestamp not null default (now()),
grade integer not null,
subjective_rating varchar(30),
notes text,
determined_english_level varchar(3) not null,
main_task integer,
question json,
created_at timestamp not null default (now()),
updated_at timestamp not null default (now()),
deleted boolean default false
);
commit;
