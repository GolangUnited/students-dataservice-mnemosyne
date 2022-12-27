BEGIN;

CREATE TABLE user_lessons (
    user_id integer references users(id),
    lesson_id integer references lessons(id)
);

COMMIT;
