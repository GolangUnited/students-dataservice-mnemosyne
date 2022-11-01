BEGIN;

CREATE TABLE student_teams (
    student_id integer references users(id),
    team_id integer references teams(id)
);

COMMIT;