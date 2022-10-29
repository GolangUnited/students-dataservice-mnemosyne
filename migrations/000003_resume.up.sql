begin;
create table if not exists RESUME (
    User_ID uuid primary key references Users(ID),
    Experience text,
    Uploaded_Resume inet,
    Country varchar(20),
    City varchar(20),
    Time_Zone varchar(10),
    Mentors_Note text,
    Created_At timestamp,
    Updated_At timestamp,
    Deleted boolean default false
) ;
commit;