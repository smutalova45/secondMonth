package inventory

import (
	"database/sql"

	"main/hw5/category"
	"main/hw5/product"
	"time"
)

type Inventory struct {
	db *sql.DB
}

func New(db *sql.DB) Inventory {
	return Inventory{
		db: db,
	}
}

func (i Inventory) AddCategory(category category.Category) error {
	category.CreatedAt = time.Now()

	if _, err := i.db.Exec(`insert into category values ($1, $2, $3, $4 )`, category.Id, category.Name, category.CreatedAt, category.UpdatedAt); err != nil {
		return err
	}
	return nil

}

func (i Inventory) GetCategory() ([]category.Category, error) {
	rows, err := i.db.Query(`select * from category`)
	if err != nil {
		return nil, err
	}
	c := []category.Category{}
	for rows.Next() {
		ca := category.Category{}
		if err = rows.Scan(&ca.Id, &ca.Name, &ca.CreatedAt, &ca.UpdatedAt); err != nil {
			return nil, err
		}
		c = append(c, ca)

	}
	return c, nil

}

func (i Inventory) UpdateCategory(c category.Category) (category.Category, error) {
	c.UpdatedAt = time.Now()
	if _, err := i.db.Exec(`update category set namec= $1 , updated_at = $2 where id = $3`, c.Name, c.UpdatedAt, c.Id); err != nil {
		return category.Category{}, err
	}
	return c, nil
}

func (i Inventory) DeleteCategory(id int) error {
	if _, err := i.db.Exec(`delete from category where id = $1 `, id); err != nil {
		return err
	}
	return nil

}

func (i Inventory) Addproduct(product product.Product) error {
	product.CreatedAt = time.Now()

	if _, err := i.db.Exec(`insert into products values ($1, $2, $3, $4 )`, product.Id, product.Name, product.CtegoryId, product.CreatedAt); err != nil {

		return err
	}
	return nil

}

func (i Inventory) GetProduct() ([]product.Product, error) {
	rows, err := i.db.Query(`select * from products `)
	if err != nil {
		return nil, err
	}
	pro := []product.Product{}
	for rows.Next() {
		p := product.Product{}
		if err = rows.Scan(&p.Id, &p.Name, &p.CtegoryId, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		pro = append(pro, p)

	}
	return pro, nil
}

func (i Inventory) UpdateProduct(p product.Product) (product.Product, error) {
	p.UpdatedAt = time.Now()
	if _, err := i.db.Exec(`update products set updated_at = $1 where id = $2 `, p.UpdatedAt, p.Id); err != nil {
		return product.Product{}, err
	}
	return p, nil
}

func (i Inventory) DeleteProduct(id int) error {
	if _, err := i.db.Exec(`delete from products where id = $1`, id); err != nil {
		return err
	}
	return nil

}
