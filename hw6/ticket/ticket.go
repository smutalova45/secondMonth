package ticket

import (
	"time"

	_"github.com/google/uuid"
)

type Ticket struct {
	Id        string
	From      string
	To        string
	DateOfFly time.Time
}
