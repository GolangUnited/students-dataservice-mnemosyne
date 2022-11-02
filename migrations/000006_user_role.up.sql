BEGIN;

CREATE TABLE user_role (
    user_id integer references users(id),
    role_id integer references roles(id)
);

COMMIT;
