BEGIN;

CREATE TABLE if not exists teams (
    id serial PRIMARY KEY,
    group_id integer references groups(id),
    name varchar(255) not null,
    project_id integer references projects(id),
    start_date timestamp not null,
    end_date timestamp not null,
    mentor_id integer references users(id),    
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now()),
    deleted boolean default false
);

COMMIT;