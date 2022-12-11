begin;
create table if not exists сontacts (
    user_id integer primary key references users(id),   
    telegram varchar(255),
    discord varchar(255),
    communication_channel varchar(255),
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now()),
    deleted boolean default false
) ;
commit;
