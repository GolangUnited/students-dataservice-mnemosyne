begin;

DELETE FROM roles WHERE id IN (1, 2, 3);

commit;
