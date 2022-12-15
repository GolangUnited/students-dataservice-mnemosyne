BEGIN;

CREATE TABLE if not exists roles (
    id serial PRIMARY KEY,
    code varchar(50) not null UNIQUE
);

COMMIT;
