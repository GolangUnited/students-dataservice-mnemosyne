BEGIN;

CREATE TABLE if not exists groups (
    id serial PRIMARY KEY,
    name varchar(30),
    start_date timestamp,
    end_date timestamp,
    Created_At timestamp,
    Updated_At timestamp,
    Deleted boolean default false
);

COMMIT;