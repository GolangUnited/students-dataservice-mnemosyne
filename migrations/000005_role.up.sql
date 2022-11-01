BEGIN;

CREATE TABLE if not exists roles (
    id serial PRIMARY KEY,
    name varchar(255) not null
);

COMMIT;
