BEGIN;

CREATE TABLE if not exists teams (
    id serial PRIMARY KEY,
    group_id integer references groups(id),
    name varchar(20) not null,
    project_id integer references projects(id),
    start_date timestamp not null,
    end_date timestamp not null,
    mentor_id uuid references users(id),    
    Created_At timestamp not null default (now()),
    Updated_At timestamp not null default (now()),
    Deleted boolean default false
);

COMMIT;