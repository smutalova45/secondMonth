package main

import (
	"database/sql"
	"fmt"
	_ "main/hw5/category"
	"main/hw5/inventory"

	_ "github.com/lib/pq"
)

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	inv := inventory.New(db)
	//CATEGORY
	//adding

	// var id int
	// var name string
	// fmt.Print("name :")
	// fmt.Scan(&name)
	// fmt.Print("id for this user :")
	// fmt.Scan(&id)
	// s := category.Category{
	// 	Id:   id,
	// 	Name: name,
	// }
	// err = inv.AddCategory(s)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("Category added successfully:")
	// 	fmt.Printf("ID: %d\nName: %s\nCreated At: %s\nUpdated At: %s\n", s.Id, s.Name, s.CreatedAt, s.UpdatedAt)
	// }

	//getting
	// s, err := inv.GetCategory()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(s)

	//Updating

	// id := 2
	// f := category.Category{
	// 	Id:   id,
	// 	Name: name,
	// }
	// _, err = inv.UpdateCategory(f)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	//deleting
	// err = inv.DeleteCategory(2)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("deleted")

	//PRODUCTS
	//add
	// var name string
	// var id int
	// var categoryid int
	// fmt.Print("enter name: ")
	// fmt.Scan(&name)
	// fmt.Print("id :")
	// fmt.Scan(&id)
	// fmt.Print("category id of this product: ")
	// fmt.Scan(&categoryid)

	// st := product.Product{
	// 	Id:        id,
	// 	Name:      name,
	// 	CtegoryId: categoryid,
	// }
	// err = inv.Addproduct(st)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("added")
	// }

	//updating

	// var id int
	// fmt.Print("which product do you want to update:( enter id) ")
	// fmt.Scan(&id)
	// ps := product.Product{
	// 	Id: id,
	// }
	// _, err = inv.UpdateProduct(ps)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(" updated ")
	// }

	//get
	f, err := inv.GetProduct()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(f)
	//deleting

	id := 1
	err = inv.DeleteProduct(id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("deleted")
	}

}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=person password=12345 database=products sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
