BEGIN;

CREATE TABLE if not exists groups (
    id serial PRIMARY KEY,
    name varchar(30) not null,
    start_date timestamp not null,
    end_date timestamp not null,
    Created_At timestamp not null default (now()),
    Updated_At timestamp not null default (now()),
    Deleted boolean default false
);

COMMIT;