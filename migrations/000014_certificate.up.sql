BEGIN;

CREATE TABLE if not exists certificates (
    id serial PRIMARY KEY,
    user_id integer references users(id),
    issue_date timestamp not null,
    expire_date timestamp not null,
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now()),
    deleted boolean default false
);

COMMIT;
