BEGIN;

CREATE TABLE user_role (
    User_ID uuid references Users(ID),
    Role_ID integer references Roles(ID)
);

COMMIT;