package inventory

import (
	"database/sql"
	"main/hw6/ticket"
	"main/hw6/user"

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

//USER

func (i Inventory) AddUser(user user.User) error {

	if _, err := i.db.Exec(`insert into users values($1,$2,$3,$4,$5,$6)`, user.Id, user.FirstName, user.LastName, user.Email, user.TicketId, user.Phone); err != nil {
		return err
	}
	return nil
}
func (i Inventory) GetAllUsers() ([]user.User, error) {
	rows, err := i.db.Query(`select * from users`)
	if err != nil {
		return nil, err
	}
	us := []user.User{}
	for rows.Next() {
		u := user.User{}
		if err = rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.TicketId, &u.Phone); err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, nil

}
func (i Inventory) GetUserById(id uuid.UUID) (user.User, error) {
	user := user.User{}
	row := i.db.QueryRow(`select id,firstname,lastname,email , ticket_id , phone from users where id=$1`, id)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.TicketId, &user.Phone); err != nil {
		return user, err
	}
	return user, nil
}

func (i Inventory) UpdateUsers(u user.User) (user.User, error) {
	if _, err := i.db.Exec(`update users set email = $1 ,phone = $2 where id = $3`, u.Email, u.Phone, u.Id); err != nil {
		return user.User{}, err
	}
	return u, nil

}
func (i Inventory) DeleteUserById(id string) error {

	if _, err := i.db.Exec(`delete from users where id=$1`, id); err != nil {
		return err
	}
	return nil

}

//TICKET

func (i Inventory) AddTicket(ticket ticket.Ticket) error {
	if _, err := i.db.Exec(`insert into ticket values ($1,$2,$3,$4),`, ticket.Id, ticket.From, ticket.To, ticket.DateOfFly); err != nil {
		return err
	}
	return nil
}
func (i Inventory) GetAllTickets() ([]ticket.Ticket, error) {
	rows, err := i.db.Query(`select *from ticket `)
	if err != nil {
		return nil, err
	}
	tickets := []ticket.Ticket{}
	for rows.Next() {
		t := ticket.Ticket{}
		if err = rows.Scan(&t.Id, &t.From, &t.To, &t.DateOfFly); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}

func (i Inventory) GetTicketById(id uuid.UUID) (ticket.Ticket, error) {
	t := ticket.Ticket{}
	row := i.db.QueryRow(`select id, from_city, to_city , date_of_fly from ticket where id=$1`, id)
	if err := row.Scan(&t.Id, &t.From, &t.To, &t.DateOfFly); err != nil {

		return ticket.Ticket{}, err
	}
	return t, nil
}

// func (i Inventory) UpdateTicket(t ticket.Ticket) (ticket.Ticket, error) {
// 	if _, err := i.db.Exec(`update ticket set date_of_fly = &1 where id = &2`, t.DateOfFly, t.Id); err != nil {
// 		return ticket.Ticket{}, err
// 	}
// 	return t, nil
// }
func (i Inventory) UpdateTicket( t ticket.Ticket) (ticket.Ticket, error) {
	_, err := i.db.Exec("UPDATE ticket SET date_of_fly = $1 WHERE id = $2", t.DateOfFly,t.Id)
	if err != nil {
		
	  return ticket.Ticket{}, err
	}
	return t, nil
  }

func (i Inventory) DeleteByIdTicket(id string) error {

	if _, err := i.db.Exec(`delete from ticket where id = &1`, id); err != nil {
		return err
	}
	return nil
}

//REPORT

func (i Inventory) Report(from string, to string) ([]ticket.Ticket, []user.User, error) {
	rows, err := i.db.Query(`select t.id as No,t.from_city,t.to_city,u.firstname,u.lastname,u.email, u.phone from users as u inner join ticket as t on t.id=u.ticket_id where t.from_city = $1 and t.to_city = $2`, from, to)

	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var tickets []ticket.Ticket
	var users []user.User

	for rows.Next() {
		var t ticket.Ticket
		var u user.User
		if err := rows.Scan(&t.Id, &t.From, &t.To, &u.FirstName, &u.LastName, &u.Email, &u.Phone); err != nil {
			return nil, nil, err
		}
		tickets = append(tickets, t)
		users = append(users, u)

	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return tickets, users, nil

}
