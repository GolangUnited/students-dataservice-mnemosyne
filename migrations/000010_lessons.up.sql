BEGIN;

CREATE TABLE if not exists lessons (
    id serial PRIMARY KEY,
    presentation inet,
    video inet,
    date timestamp not null,
    homework inet,
    lecturer_id uuid references users(id),
    group_id integer references groups(id),
    Language varchar(5) not null,
    Created_At timestamp not null default (now()),
    Updated_At timestamp not null default (now()),
    Deleted boolean default false
);

COMMIT;