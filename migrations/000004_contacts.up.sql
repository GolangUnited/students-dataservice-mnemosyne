begin;
create table if not exists Contacts (
    User_ID uuid primary key references Users(ID),
    Email varchar(30),
    Telegram varchar(20),
    Discord varchar(20),
    Communication_Channel varchar(10),
    Created_At timestamp,
    Updated_At timestamp,
    Deleted boolean default false
) ;
commit;