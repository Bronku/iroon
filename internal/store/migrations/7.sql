create table user (login text primary key, password text, role text);

drop table session;

create table session (
    token text primary key,
    user text,
    expiration text
);

pragma user_version = 7;
