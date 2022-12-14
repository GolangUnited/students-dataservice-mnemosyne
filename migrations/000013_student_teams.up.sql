BEGIN;

CREATE TABLE if not exists user_teams (
    user_id integer references users(id),
    team_id integer references teams(id),
    UNIQUE (user_id, team_id)
);

COMMIT;
