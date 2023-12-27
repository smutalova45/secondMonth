package product

import "time"

type Product struct {
	Id        int
	Name      string
	CtegoryId int
	CreatedAt time.Time
	UpdatedAt time.Time
}
