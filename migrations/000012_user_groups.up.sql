BEGIN;

CREATE TABLE user_groups (
    user_id integer references users(id),
    group_id integer references groups(id)
);

COMMIT;
