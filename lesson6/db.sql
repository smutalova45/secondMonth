CREATE TABLE countries(
    id int primary key,
    namecountry varchar(15)
);
CREATE TABLE cities(
    id int ,
    namecity varchar(15),
    population varchar(20),
    id_country int references countries(id)
);