CREATE TABLE countries(
    id uuid primary key not null,
    name_of_c varchar(20),
    currency varchar(10)
);
CREATE TABLE cities(
    id uuid primary key,
    population_of int default 0,
    country_id uuid references countries(id),
    name_of_city varchar(20)
);