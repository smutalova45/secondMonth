package main

import (
	"database/sql"
	"fmt"
	"main/lesson7/inventory"

	_ "github.com/lib/pq"
)

func main() {
	db, err := ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	//add
	//get
	//delete
	//update

	//LIST
	fmt.Print("enter payment_type : ")
	var paymentype string
	fmt.Scan(&paymentype)
	inv := inventory.New(db)
	payments, err := inv.List(paymentype)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for i, pay := range payments {
			fmt.Printf("ticket %d: id  : %s , payment_type : %s , date:%s, amount : %d\n", i+1, pay.Id, pay.PaymentType, pay.Date, pay.Amount)

		}

	}

}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=person password=12345 database=payment sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
