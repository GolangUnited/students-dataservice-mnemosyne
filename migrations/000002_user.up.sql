begin;
create table if not exists Users (
ID uuid primary key,
Last_Name varchar(80),
First_Name varchar(80),
Middle_Name varchar(80),
Language varchar(80),
English_Level varchar(3),
Photo inet,
Created_At timestamp,
Updated_At timestamp,
Deleted boolean default false
);
commit;