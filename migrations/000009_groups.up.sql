BEGIN;

CREATE TABLE if not exists groups (
    id serial PRIMARY KEY,
    name varchar(255) not null,
    start_date timestamp not null,
    end_date timestamp not null,
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now()),
    deleted boolean default false
);

COMMIT;