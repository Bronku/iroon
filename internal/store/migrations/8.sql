alter table user
add column salt text;

pragma user_version = 8;
