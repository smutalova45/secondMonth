package main

import (
	"database/sql"
	"fmt"
	"main/hw4/inventory"
	"main/hw4/user"

	_ "github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	adduser = iota + 1
	getalluser
	getbyid
	deleteuser
)

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	inv := inventory.New(db)

	for true {
		var cmd int
		fmt.Println("1.ADD Users 2.Getallusers 3.Getby ID 4.Delete by idz")
		fmt.Scan(&cmd)
		switch cmd {
		case adduser:
			var u user.User
			fmt.Println("inserting data to database : ")
			fmt.Println("--------------------")
			fmt.Print("enter id : ")
			_, err := fmt.Scan(&u.Id)
			if err != nil {
				fmt.Println(err.Error())

			}
			fmt.Print("Enter First Name : ")
			_, err = fmt.Scan(&u.FirstName)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Print("Enter last name: ")
			_, err = fmt.Scan(&u.LastName)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Print("enter email: ")
			_, err = fmt.Scan(&u.Email)
			if err != nil {
				fmt.Println(err.Error())
			}

			err = inv.AddUser(u)
			if err != nil {
				fmt.Println("Error inserting user:", err)
			} else {
				fmt.Println("User added successfully")
			}

		case getalluser:
			fmt.Println("----------------")
			fmt.Println(" Users list ")
			users1, err := inv.GetAllUsers()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(users1)

		case getbyid:
			fmt.Println("----------------")
			var id int
			fmt.Print("enter id : ")
			fmt.Scan(&id)
			u1, err := inv.GetUserById(id)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("User information:\nID: %d\nFirst Name: %s\nLast Name: %s\nEmail: %s\n", u1.Id, u1.FirstName, u1.LastName, u1.Email)
			}
		case deleteuser:
			fmt.Println("----------------")
			fmt.Print("enter user id to delete : ")
			var id int
			fmt.Scan(&id)
			err := inv.DeleteUserById(id)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("DELETED")
				
				

			}
			users1, err := inv.GetAllUsers()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(users1)
		}
	}

}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=person password=12345 database=users sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
