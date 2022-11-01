BEGIN;

CREATE TABLE if not exists projects (
    id serial PRIMARY KEY,
    name varchar(255) not null,
    description text not null,
    external_doc inet not null,
    mentor_id integer references users(id),
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now()),
    deleted boolean default false
);

COMMIT;