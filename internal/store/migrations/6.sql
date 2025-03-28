drop table session;

create table session (
    token text pirmary key,
    user text,
    expiration text
);

pragma user_version = 6;
