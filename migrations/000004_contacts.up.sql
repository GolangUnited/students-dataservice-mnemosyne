begin;
create table if not exists Contacts (
    User_ID uuid primary key references Users(ID),
    Email varchar(30) not null,
    Telegram varchar(20) not null,
    Discord varchar(20) not null,
    Communication_Channel varchar(10),
    Created_At timestamp not null default (now()),
    Updated_At timestamp not null default (now()),
    Deleted boolean default false
) ;
commit;