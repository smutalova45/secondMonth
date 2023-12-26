package main

import (
	"database/sql"
	"fmt"
	"main/hw6/inventory"
	"main/hw6/ticket"
	"main/hw6/user"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	adduser = iota + 1
	getalluser
	getuserbyid
	updateuser
	deleteuser
	back1
)
const (
	addticket = iota + 1
	getAllTickets
	getticketbyid
	updateTicket
	deleteTicketById
	back
)
const (
	tickets = iota + 1
	users
	report
	finish
)

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	inv := inventory.New(db)

	for true {
		fmt.Println("1.Tickets 2.Users 3.Report 4.Finish")
		var cmd int
		fmt.Print("ENTER : ")
		fmt.Scan(&cmd)
		switch cmd {
		case tickets:
			fmt.Println("Tickets")
			for true {
				var cmd1 int
				fmt.Println("1. Add ticket, 2.Get all tickets 3. Get By Id 4. Update ticket 5. Delete ticket 6.back to main menu")
				fmt.Print("enter cmd :")
				fmt.Scan(&cmd1)
				switch cmd1 {
				case addticket:
					var t ticket.Ticket
					fmt.Println("------------------")
					var idStr string

					_, err := uuid.Parse(idStr)
					if err != nil {
						fmt.Println("Error parsing UUID:", err)
					}
					fmt.Print("ENTER | Id ticket :")
					if _, err := fmt.Scan(&idStr); err != nil {
						fmt.Println(err.Error())
					}

					fmt.Print("ENTER | (which country) From : ")
					if _, err = fmt.Scan(&t.From); err != nil {
						fmt.Println(err.Error())
					}
					fmt.Print("ENTER | (which country) To : ")
					if _, err = fmt.Scan(&t.To); err != nil {
						fmt.Println(err.Error())
					}

					fmt.Print("ENTER | Date of fly : ")
					var data string
					fmt.Scan(&data)

					DateOfFly, err := time.Parse("2006-01-02", data)
					if err != nil {
						fmt.Print(err.Error())
					}
					t.DateOfFly = DateOfFly
					fmt.Println(t.DateOfFly)

					err = inv.AddTicket(t)
					if err != nil {
						fmt.Println("error in inserting value ", err)
					} else {
						fmt.Println("added succesfully")
					}

				case getAllTickets:
					fmt.Println("-------------------------")
					fmt.Println("list of tickets")
					tickets1, err := inv.GetAllTickets()
					if err != nil {
						fmt.Println(err.Error())
					}
					fmt.Println(tickets1)

				case getticketbyid:
					fmt.Println("-------------------------")
					var idstr string
					fmt.Println("enter id :")
					fmt.Scan(&idstr)
					idt, err := uuid.Parse(idstr)

					if err != nil {
						fmt.Println(err.Error())
					}

					t1, err := inv.GetTicketById(idt)
					if err != nil {
						fmt.Println(err.Error())
					} else {

						fmt.Printf("Ticket information:\nId: %s\nFrom: %s\nTo: %s\nDate: %s\n", idstr, t1.From, t1.To, t1.DateOfFly)
					}
					//UPDATE
				case updateTicket:

					var id string
					fmt.Println("enter id (uuid)")

					fmt.Scan(&id)

					_, err := uuid.Parse(id)
					if err != nil {
						fmt.Print(err.Error())
					}
					var date string
					fmt.Println("enter date (yyyy-mm-dd):")
					fmt.Scan(&date)
					dateof, err := time.Parse("2006-01-02", date)
					if err != nil {
						fmt.Println(err.Error())
					}
					t := ticket.Ticket{
						Id:        id,
						DateOfFly: dateof,
					}

					_, err = inv.UpdateTicket(t)
					if err != nil {
						fmt.Println(err.Error())
					}
					fmt.Println("updated")

				case deleteTicketById:
					fmt.Println("-------------------------")

					var idString string
					fmt.Print(" enter id in uuid type: ")
					fmt.Scan(&idString)
					_, err := uuid.Parse(idString)
					if err != nil {
						fmt.Println("error parsing UUID:", err)
						return
					}

					err = inv.DeleteByIdTicket(idString)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("DELETED")
					}

				case back:
					break
				default:
					fmt.Println(" wrong cmd entered ")
				}
				if cmd1 == back {
					break
				}

			}
		case users:
			fmt.Println("users")
			for true {
				var cmd2 int
				fmt.Println("1. Add user, 2.Get all users 3. Get By Id 4. Update user 5. Delete user 6. Back to main menu")
				fmt.Print("enter cmd :")
				fmt.Scan(&cmd2)
				switch cmd2 {

				case adduser:
					var u user.User
					fmt.Print(u)
				case getalluser:
					users1, err := inv.GetAllUsers()
					if err != nil {
						fmt.Println(err.Error())
					}
					fmt.Println(users1)
				case getuserbyid:

					var idstr string
					fmt.Print("enter id: *in uuid: ")
					fmt.Scan(&idstr)
					idt, err := uuid.Parse(idstr)
					if err != nil {
						fmt.Println(err.Error())
					}
					u1, err := inv.GetUserById(idt)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("user information : id :", idstr, "firstname : ", u1.FirstName, "lastname: ", u1.LastName, "email:", u1.Email, "ticket_id: ", u1.TicketId)
					}
				case updateuser:

					// fmt.Println("enter id (uuid)")

					// fmt.Scan(&id)

					// _, err := uuid.Parse(id)
					// if err != nil {
					// 	fmt.Print(err.Error())
					// }
					// var date string
					// fmt.Println("enter date (yyyy-mm-dd):")
					// fmt.Scan(&date)
					// dateof, err := time.Parse("2006-01-02", date)
					// if err != nil {
					// 	fmt.Println(err.Error())
					// }
					// t := ticket.Ticket{
					// 	Id:        id,
					// 	DateOfFly: dateof,
					// }

					// _, err = inv.UpdateTicket(t)
					// if err != nil {
					// 	fmt.Println(err.Error())
					// }
					// fmt.Println("updated")
					fmt.Println("enter id (uuid)")
					var id string
					fmt.Scan(&id)
					_, err := uuid.Parse(id)
					if err != nil {
						fmt.Println(err.Error())
					}
					

				case deleteuser:
					var idstr string
					fmt.Println("id enter: ")
					fmt.Scan(&idstr)
					_, err := uuid.Parse(idstr)
					if err != nil {
						fmt.Println(err.Error())
					}
					err = inv.DeleteUserById(idstr)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println("deleted")
					}

				case back1:
					break
				default:
					fmt.Println(" wrong cmd entered")
				}

				if cmd2 == back1 {
					break
				}

			}
		case report:
			fmt.Println("----------------")
			fmt.Println("  Report : ")
			fmt.Print("enter from ")
			var from string
			fmt.Scan(&from)
			fmt.Println("enter to ")
			var to string
			fmt.Scan(&to)

			user, ticket, err := inv.Report(from, to)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(user)
			fmt.Println(ticket)

		case finish:
			return
		default:
			fmt.Println(" wrong cmd entered ")
		}

	}

}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=person password=12345 database=homework6 sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
