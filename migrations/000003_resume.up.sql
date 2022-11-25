begin;
create table if not exists resume(
    user_id integer primary key references users(id),
    experience text not null,
    uploaded_resume text not null,
    country varchar(20) not null,
    city varchar(20) not null,
    time_zone varchar(10) not null,
    mentors_note text not null,
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now()),
    deleted boolean default false
) ;
commit;
