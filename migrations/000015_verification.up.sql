begin;
create table if not exists verification (
    id serial PRIMARY KEY,
    email varchar(255) not null,
    secret_token varchar(255) not null,
    access_token varchar(255) not null,
    expire_date timestamp not null
    );

commit;
