CREATE TABLE companies(
    id int ,
    name_ varchar(20),
    age int,
    address_ varchar(15),
    salary int

);

create or replace function ages(age int)
returns integer
language plpgsql AS
$$ 
begin 
if age>10 and age <23 then
return 1;
else
return 0;
end if;
end;
$$;

CREATE FUNCTION selecting()
RETURNS integer
LANGUAGE plpgsql AS
$$
declare 
result integer = 0;
begin

  SELECT  COUNT(*) into result FROM companies WHERE age between 5 and 11;
  return result;

end;
$$;
