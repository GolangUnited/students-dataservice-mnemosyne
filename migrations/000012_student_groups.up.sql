BEGIN;

CREATE TABLE if not exists student_groups (
    student_id integer references users(id),
    group_id integer references groups(id)
);

COMMIT;
