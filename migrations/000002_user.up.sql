begin;
create table if not exists users (
id serial primary key,
last_name varchar(255) not NULL,
first_name varchar(255) not null,
middle_name varchar(255) not null,
email varchar(255) not null UNIQUE,
language varchar(80) not null,
english_level varchar(50) not null,
photo text not null,
created_at timestamp not null default (now()),
updated_at timestamp not null default (now()),
deleted boolean default false
);
commit;
