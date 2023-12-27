package inventory

import (
	"database/sql"
	"fmt"
	"main/lesson7/payment"

	"github.com/google/uuid"
)

type Inventory struct {
	db *sql.DB
}

func New(db *sql.DB) Inventory {
	return Inventory{
		db: db,
	}

}
func (i Inventory) Add(payment payment.Payment) error {
	payment.Id = uuid.New().String()
	_, err := i.db.Exec(`insert into payments values( $1, $2, $3, $4)`, payment.Id, payment.PaymentType, payment.Date, payment.Amount)
	if err != nil {
		return err
	}
	return nil

}
func (i Inventory) GetAll() ([]payment.Payment, error) {
	rows, err := i.db.Query(`select * from payments `)
	if err != nil {
		return nil, err
	}
	pay := []payment.Payment{}
	for rows.Next() {
		p := payment.Payment{}
		if err = rows.Scan(&p.Id, &p.PaymentType, &p.Date, &p.Amount); err != nil {
			return nil, err
		}
		pay = append(pay, p)
	}
	return pay, nil

}
func (i Inventory) Update(p payment.Payment) (payment.Payment, error) {
	if _, err := i.db.Exec(`update payments set payment_type = $1, date = $2 where id =$3`,p.Id); err != nil {
		return payment.Payment{}, err
	}
	return p, nil

}
func (i Inventory) Delete(id string) error {
	if _, err := i.db.Exec(`delete from payments where id = $1`, id); err != nil {
		return err
	}
	return nil

}
func (i Inventory) List(paymenttype string) ([]payment.Payment, error) {
	rows, err := i.db.Query(`select id, payment_type, date_, amount_ where payment_type = $1`, paymenttype)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var payments []payment.Payment
	for rows.Next() {
		var p payment.Payment
		if err := rows.Scan(&p.Id, &p.PaymentType, &p.Date, &p.Amount); err != nil {
			fmt.Println(err.Error())
		}
		payments = append(payments, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return payments, nil

}
