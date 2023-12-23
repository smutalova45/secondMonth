package user

import _ "github.com/google/uuid"

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}
