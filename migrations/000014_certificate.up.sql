BEGIN;

CREATE TABLE Certificates (
    id serial PRIMARY KEY,
    User_id uuid references Users(id),
    Issue_date timestamp not null,
    Expire_date timestamp not null
);

COMMIT;
