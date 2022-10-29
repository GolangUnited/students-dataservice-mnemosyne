begin;
create table if not exists Interview (
ID serial primary key,
Interviewer_ID uuid references Users(id),
Student_ID uuid references Users(id),
Date timestamp not null default (now()),
Grade integer not null,
Subjective_Rating varchar(30),
Notes text,
Determined_English_level varchar(3) not null,
Main_Task integer,
Question1 integer,
Question2 integer,
Question3 integer,
Question4 integer,
Question5 integer,
Question6 integer,
Question7 integer,
Created_At timestamp not null default (now()),
Updated_At timestamp not null default (now()),
Deleted boolean default false
);
commit;