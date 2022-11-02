BEGIN;

CREATE TABLE if not exists lessons (
    id serial PRIMARY KEY,
    presentation text,
    video text,
    lesson_date timestamp not null,
    homework text,
    lecturer_id integer references users(id),
    group_id integer references groups(id),
    language varchar(255) not null,
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now()),
    deleted boolean default false
);

COMMIT;
