
package main

import (
	"ca1/tracker"
	"fmt"
	"os"
)

func main() {

	db := tracker.NewTicketDb()

	db.CreateTicket(tracker.NewTicket(101, "Login Bug", "High", "", "Open"))
	db.CreateTicket(tracker.NewTicket(102, "Payment Bug", "Critical", "", "Open"))

	if len(os.Args) < 2 {
		fmt.Println("create / display / assign / search / close")
		return
	}

	switch os.Args[1] {

	case "create":
		fmt.Println("Tickets Created")
		tracker.SaveFile()

	case "display":
		db.DisplayAll()

	case "assign":
		db.AssignDeveloper(101, "Arun")
		db.DisplayAll()

	case "search":
		ticket, _ := db.SearchTicket(101)
		ticket.Display()

	case "close":
		db.CloseTicket(102)
		db.DisplayAll()
	}
}