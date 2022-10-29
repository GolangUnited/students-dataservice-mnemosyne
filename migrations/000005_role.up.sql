BEGIN;

CREATE TABLE if not exists Roles (
    id serial PRIMARY KEY,
    Name varchar(20) not null
);

COMMIT;
