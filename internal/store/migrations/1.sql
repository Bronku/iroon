create table cake (
    id integer primary key autoincrement,
    name text not null unique,
    price integer not null,
    category text,
    availability text
);

create table customer_order (
    id integer primary key autoincrement,
    name text,
    surname text,
    phone text,
    location text,
    order_date string not null,
    delivery_date string not null,
    status text,
    paid integer
);

create table ordered_cake (
    customer_order integer references customer_order (id) not null,
    cake integer references cake (id) not null,
    amount integer not null,
    primary key (customer_order, cake)
);

create table user (
    login text primary key,
    password text,
    role text,
    salt text
);

create table session (
    token text primary key,
    user text,
    expiration text
);

create virtual table order_fts using fts5 (
    id,
    name,
    surname,
    phone,
    location,
    order_date,
    delivery_date,
    status,
    paid,
    content = 'customer_order'
);

create trigger order_fts_insert after insert on customer_order begin
insert into
    order_fts (
        id,
        name,
        surname,
        phone,
        location,
        order_date,
        delivery_date,
        status,
        paid
    )
values
    (
        new.id,
        new.name,
        new.surname,
        new.phone,
        new.location,
        new.order_date,
        new.delivery_date,
        new.status,
        new.paid
    );

end;

create trigger order_fts_delete after delete on customer_order begin
delete from order_fts
where
    id = old.id;

end;

create trigger order_fts_update after
update on customer_order begin
update customer_order
set
    id = new.id,
    name = new.name,
    surname = new.surname,
    phone = new.phone,
    location = new.location,
    order_date = new.order_date,
    delivery_date = new.delivery_date,
    status = new.status,
    paid = new.paid
where
    id = old.id;

end;

pragma user_version = 1;
