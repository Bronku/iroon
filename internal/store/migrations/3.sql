alter table cake
add column category text;

alter table cake
add column availability text;

PRAGMA user_version = 3;
