BEGIN;

CREATE TABLE if not exists projects (
    id serial PRIMARY KEY,
    name varchar(255) not null,
    description text not null,
    git_url varchar(255),
    team_id integer references teams(id),
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now()),
    deleted boolean default false
);

COMMIT;
