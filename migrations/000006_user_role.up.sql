BEGIN;

CREATE TABLE if not exists user_role (
    user_id integer references users(id),
    role_id integer references roles(id) default 3
);

COMMIT;
