BEGIN;

CREATE TABLE Student_teams (
    Student_ID uuid references Users(ID),
    Team_ID integer references Teams(ID)
);

COMMIT;