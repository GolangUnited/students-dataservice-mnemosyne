begin;
create table if not exists Users (
ID uuid primary key,
Last_Name varchar(80) not NULL,
First_Name varchar(80) not null,
Middle_Name varchar(80) not null,
Language varchar(80) not null,
English_Level varchar(3) not null,
Photo inet,
Created_At timestamp not null default (now()),
Updated_At timestamp not null default (now()),
Deleted boolean default false
);
commit;