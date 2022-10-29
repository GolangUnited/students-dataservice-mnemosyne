BEGIN;

CREATE TABLE Student_groups (
    Student_ID uuid references Users(ID),
    Group_ID integer references groups(ID)
);

COMMIT;