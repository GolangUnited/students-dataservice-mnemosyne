BEGIN;

CREATE TABLE if not exists teams (
    id serial PRIMARY KEY,
    group_id integer references groups(id),
    name varchar(20),
    project_id integer references projects(id),
    start_date timestamp,
    end_date timestamp,
    mentor_id uuid references users(id),    
    Created_At timestamp,
    Updated_At timestamp,
    Deleted boolean default false
);

COMMIT;