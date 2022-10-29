BEGIN;

CREATE TABLE if not exists lessons (
    id serial PRIMARY KEY,
    presentation inet,
    video inet,
    date timestamp,
    homework inet,
    lecturer_id uuid references users(id),
    group_id integer references groups(id),
    Language varchar(5),
    Created_At timestamp,
    Updated_At timestamp,
    Deleted boolean default false
);

COMMIT;