begin;
create table if not exists resume(
    user_id integer primary key references users(id),
    experience text,
    uploaded_resume text,
    country varchar(20),
    city varchar(20),
    time_zone varchar(10),
    mentors_note text,
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now()),
    deleted boolean default false
) ;
commit;
