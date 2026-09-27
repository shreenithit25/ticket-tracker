package main

import (
	"ca1/tracker"
	"fmt"
)

func main() {

	db := tracker.NewTicketDb()

	err := db.CreateTicket(101, "Login Bug", "High", "", "Open")
	if err != nil {
		fmt.Println(err)
	}

	err = db.CreateTicket(102, "Payment Bug", "Critical", "", "Open")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All Tickets")
	db.DisplayAll()

	// Assign Arun to Ticket 101
	err = db.AssignDeveloper(101, "Arun")
	if err != nil {
		fmt.Println(err)
	}

	// Assign Arun to Ticket 102
	err = db.AssignDeveloper(102, "Arun")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("After Assigning Developer")
	db.DisplayAll()

	ticket, err := db.SearchTicket(101)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Search Result")
		ticket.Display()
	}

	err = db.CloseTicket(102)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("After Closing Ticket 102")
	db.DisplayAll()
}