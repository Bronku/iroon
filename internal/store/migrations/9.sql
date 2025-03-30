create virtual table order_fts using fts5 (
    id,
    name,
    surname,
    phone,
    location,
    status,
    order_date,
    delivery_date,
    paid,
    content = 'customer_order'
);

insert into
    order_fts (
        id,
        name,
        surname,
        phone,
        location,
        status,
        order_date,
        delivery_date,
        paid
    )
select
    id,
    name,
    surname,
    phone,
    location,
    status,
    order_date,
    delivery_date,
    paid
from
    customer_order;

create trigger order_fts_insert after insert on customer_order begin
insert into
    order_fts (
        id,
        name,
        surname,
        phone,
        location,
        status,
        order_date,
        delivery_date,
        paid
    )
values
    (
        new.id,
        new.name,
        new.surname,
        new.phone,
        new.location,
        new.status,
        new.order_date,
        new.delivery_date,
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
    status = new.status,
    order_date = new.order_date,
    delivery_date = new.delivery_date,
    paid = new.paid
where
    id = old.id;

end;

pragma user_version = 9;
