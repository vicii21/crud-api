CREATE TABLE product(
    id bigserial unique primary key not null,
    name varchar(50) not null,
    short_description varchar(120) not null,
    description varchar(255) not null,
    price decimal(12,2) not null,
    quantity bigint not null,
    created timestamp not null,
    updated timestamp, 
);