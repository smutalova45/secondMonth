CREATE TABLE ticket(
    id uuid primary key ,
    from_city varchar(30),
    to_city varchar(30),
    date_of_fly timestamp
    

);
CREATE TABLE users(
    id uuid  primary key not null,
    firstname varchar(15),
    lastname varchar(15),
    email varchar(25),
    ticket_id uuid references ticket(id),
    phone varchar(14)
);