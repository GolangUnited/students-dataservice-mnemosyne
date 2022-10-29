BEGIN;

CREATE TABLE if not exists projects (
    id serial PRIMARY KEY,
    name varchar(30) not null,
    description text not null,
    External_doc inet not null,
    Mentor_id uuid references Users(id),
    Created_At timestamp not null default (now()),
    Updated_At timestamp not null default (now()),
    Deleted boolean default false
);

COMMIT;