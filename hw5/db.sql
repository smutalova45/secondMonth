CREATE TABLE category(
    id int primary key not null,
    namec varchar(10),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

CREATE TABLE products(
    id int ,
    name_ varchar(10),
    category_id int references category(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp

);


