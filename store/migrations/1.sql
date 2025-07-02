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
