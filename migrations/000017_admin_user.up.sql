begin;

INSERT INTO users(id, last_name, first_name, email, language, english_level)
VALUES (1, 'admin', 'admin', 'admin@mnemosyne.info', 'ru', '');

INSERT INTO user_role(user_id, role_id)
VALUES (1, 1);

commit;
