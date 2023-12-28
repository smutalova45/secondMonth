CREATE TABLE company(
    id int primary key not null,
    namecompany varchar(15)
);

CREATE TABLE employers(
  id uuid ,
  name_ varchar(10),
  companyid int references company(id),
  phone varchar(15),
  
);