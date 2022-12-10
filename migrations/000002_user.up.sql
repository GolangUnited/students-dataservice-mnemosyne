begin;
create table if not exists users (
id serial primary key,
last_name varchar(255) not NULL,
first_name varchar(255) not null,
middle_name varchar(255),
email varchar(255) not null UNIQUE,
language varchar(80),
english_level varchar(50),
photo text,
created_at timestamp not null default (now()),
updated_at timestamp not null default (now()),
deleted boolean default false
);
commit;
