begin;

DELETE FROM user_role WHERE user_id = 1;
DELETE FROM users WHERE id = 1;

commit;
