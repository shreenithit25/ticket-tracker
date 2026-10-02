package tracker

import (
	"fmt"
	"os"
)

type Ticket struct {
	TicketID    int
	Title       string
	Severity    string
	AssignedDev string
	Status      string
}

type TicketDb []Ticket

func (this Ticket) Display() {
	fmt.Println(this.TicketID, this.Title, this.Severity, this.AssignedDev, this.Status)
}

func NewTicket(id int, title string, severity string, dev string, status string) Ticket {
	return Ticket{id, title, severity, dev, status}
}

func NewTicketDb() TicketDb {
	return []Ticket{}
}

func (this *TicketDb) CreateTicket(t Ticket) error {
	if t.Severity == "Low" || t.Severity == "Medium" ||
		t.Severity == "High" || t.Severity == "Critical" {
		*this = append(*this, t)
	}
	return nil
}

func (this TicketDb) SearchTicket(id int) (*Ticket, error) {
	for i := range this {
		if this[i].TicketID == id {
			return &this[i], nil
		}
	}
	return nil, nil
}

func (this *TicketDb) AssignDeveloper(id int, dev string) error {
	for i := range *this {
		if (*this)[i].TicketID == id {
			(*this)[i].AssignedDev = dev
			(*this)[i].Status = "In Progress"
		}
	}
	return nil
}

func (this *TicketDb) CloseTicket(id int) error {
	for i := range *this {
		if (*this)[i].TicketID == id {
			(*this)[i].Status = "Closed"
		}
	}
	return nil
}

func (this TicketDb) DisplayAll() {
	for _, v := range this {
		v.Display()
	}
}

func SaveFile() {
	os.WriteFile("tickets.txt", []byte("Tickets saved"), 0644)
}