package payment

import "time"

type Payment struct {
	Id          string
	PaymentType string
	Date        time.Time
	Amount      int
}
