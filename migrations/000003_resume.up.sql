begin;
create table if not exists resume(
    user_id integer primary key references users(id),
    experience text,
    uploaded_resume text,
    country varchar(255),
    city varchar(255),
    time_zone varchar(255),
    mentors_note text,
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now()),
    deleted boolean default false
) ;
commit;
