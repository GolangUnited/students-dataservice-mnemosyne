BEGIN;

CREATE TABLE if not exists projects (
    id serial PRIMARY KEY,
    name varchar(30),
    description text,
    External_doc inet,
    Mentor_id uuid references Users(id),
    Created_At timestamp,
    Updated_At timestamp,
    Deleted boolean default false
);

COMMIT;