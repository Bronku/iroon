create table session (
    token text primary key,
    user_name text,
    created text,
    lastAccess text
);

PRAGMA user_version = 2;
