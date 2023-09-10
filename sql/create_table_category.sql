CREATE TABLE category(
    id bigserial primary key not null,
    category_name varchar(50) not null,
    created_at timestamp not null default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP
);