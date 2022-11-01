begin;
create table if not exists сontacts (
    user_id integer primary key references users(id),
    email varchar(255) not null,
    telegram varchar(255) not null,
    discord varchar(255) not null,
    communication_channel varchar(10),
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now()),
    deleted boolean default false
) ;
commit;