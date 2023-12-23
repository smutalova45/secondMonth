package main

import (
	"database/sql"
	"fmt"
	_ "main/lesson4/country"
	"main/lesson4/inventory"

	_ "github.com/google/uuid"
	_ "github.com/lib/pq"
)

// DATABASE CONNECTION

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	// c := []country.Country{
	// 	{
	// 		ID:       uuid.New(),
	// 		Name:     "China",
	// 		Currency: "Yang",
	// 	},
	// 	{
	// 		ID:       uuid.New(),
	// 		Name:     "USA",
	// 		Currency: "usd",
	// 	},
	// 	{
	// 		ID:       uuid.New(),
	// 		Name:     "Uzbekistan",
	// 		Currency: "uzs",
	// 	},
	// 	{
	// 		ID:       uuid.New(),
	// 		Name:     "Russia",
	// 		Currency: "rub",
	// 	},
	// 	{
	// 		ID:       uuid.New(),
	// 		Name:     "Australia",
	// 		Currency: "aud",
	// 	},
	// 	{
	// 		ID:       uuid.New(),
	// 		Name:     "New Zealand",
	// 		Currency: "nzd",
	// 	},
	// }
	// // for _, country := range c {
	// // 	inv.InsertCountry(country)

	// }
	// if err := inv.InsertCountry(country); err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println("DATA ADDED")
	inv := inventory.New(db)
	// countries,err:=inv.GetAllCountries()
	// if err!=nil{
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(countries)
	c, err := inv.GetCountryById("bc2a15a4-4ca6-414d-be89-18bcc4b3f88c")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("-----------")
		fmt.Println(c)
	}

	inv.DeleteById("bc2a15a4-4ca6-414d-be89-18bcc4b3f88c")

}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=person password=12345 database=country sslmode=disable")
	if err != nil {
		return nil, err

	}
	return db, nil

}
