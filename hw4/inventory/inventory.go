package inventory

import (
	"database/sql"
	_"fmt"
	"main/hw4/user"
)

type Inventory struct {
	db *sql.DB
}

func New(db *sql.DB) Inventory {
	return Inventory{
		db: db,
	}

}
func (i Inventory) AddUser(user user.User) error {
	if _, err := i.db.Exec(`insert into users values ($1,$2,$3,$4)`, user.Id, user.FirstName, user.LastName, user.Email); err != nil {
		return err
	}
	return nil

}
func (i Inventory) GetAllUsers() ([]user.User, error) {
	rows, err := i.db.Query(`select *from users`)
	if err != nil {
		return nil, err
	}
	us := []user.User{}
	for rows.Next() {
		u := user.User{}
		if err = rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email); err != nil {
			return nil, err
		}
		us = append(us, u)

	}
	return us, nil
}

func (i Inventory) GetUserById(id int) (user.User,error) {
	user:=user.User{}
	row:=i.db.QueryRow(`select id,firstname,lastname,email from users where id=$1`,id)
	if err:=row.Scan(&user.Id,&user.FirstName,&user.LastName,&user.Email);err!=nil{
		return user,err
	}
	return user,nil
}

func (i Inventory) DeleteUserById(id int) error {
	
	if _, err := i.db.Exec(`delete from users where id=$1`, id); err != nil {
		return err
	}
	return nil

}
