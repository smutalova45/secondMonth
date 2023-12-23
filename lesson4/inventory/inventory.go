package inventory

import (
	"database/sql"
	

	"main/lesson4/country"

	_"github.com/google/uuid"
	_ "github.com/google/uuid"
)

type Inventory struct {
	db *sql.DB
}

func New(db *sql.DB) Inventory {
	return Inventory{
		db: db,
	}
}

func (i Inventory) InsertCountry(country country.Country) error {

	if _, err := i.db.Exec(`insert into countries values ($1,$2,$3)`, country.ID, country.Name, country.Currency); err != nil {
		return err
	}
	return nil

}

func (i Inventory) GetAllCountries() ([]country.Country, error) {
	rows, err := i.db.Query(`select * from countries`)
	if err != nil {
		return nil, err
	}
	cs := []country.Country{}
	for rows.Next() {
		c := country.Country{}
		if err = rows.Scan(&c.ID, &c.Name, &c.Currency); err != nil {

			return nil, err
		}

		cs = append(cs, c)

	}

	return cs, nil

}
func (i Inventory) GetCountryById(id string) (country.Country, error) {
	country := country.Country{}
	// id = uuid.New()
	row := i.db.QueryRow(`select id, name_of_c,currency from countries where id=$1`, id)
	if err := row.Scan(&country.ID, &country.Name,&country.Currency); err != nil {

		return country, err
	}
	return country, nil

}
func (i Inventory) DeleteById(id string) error {
	id = "23e0c72b-71c7-475e-8938-a6c9adb86957"
	if _, err := i.db.Exec(`delete from countries where id=$1`, id); err != nil {

		return err
	}
	return nil

}
