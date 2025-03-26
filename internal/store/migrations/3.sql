alter table cake
add column category text;

alter table cake
add column availability text;

pragma user_version = 3;
