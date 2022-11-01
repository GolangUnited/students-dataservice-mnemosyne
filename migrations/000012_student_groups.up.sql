BEGIN;

CREATE TABLE student_groups (
    student_id integer references users(id),
    group_id integer references groups(id)
);

COMMIT;