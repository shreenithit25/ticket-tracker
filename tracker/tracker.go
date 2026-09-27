package tracker

import "fmt"

type Ticket struct {
	TicketID    int
	Title       string
	Severity    string
	AssignedDev string
	Status      string
}

type TicketDb []Ticket

func (this Ticket) Display() {
	fmt.Println("Ticket ID: ", this.TicketID)
	fmt.Println("Title: ", this.Title)
	fmt.Println("Severity: ", this.Severity)
	fmt.Println("Assigned Developer: ", this.AssignedDev)
	fmt.Println("Status: ", this.Status)
}

func NewTicket(TicketID int, Title string, Severity string, AssignedDev string, Status string) Ticket {
	return Ticket{
		TicketID:    TicketID,
		Title:       Title,
		Severity:    Severity,
		AssignedDev: AssignedDev,
		Status:      Status,
	}
}

func NewTicketDb() TicketDb {
	return []Ticket{}
}

func (this *TicketDb) CreateTicket(TicketID int, Title string, Severity string, AssignedDev string, Status string) error {

	if Severity == "Low" || Severity == "Medium" || Severity == "High" || Severity == "Critical" {

		t := NewTicket(TicketID, Title, Severity, AssignedDev, Status)

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

func (this *TicketDb) AssignDeveloper(id int, devName string) error {

	for i := range *this {

		if (*this)[i].TicketID == id {
			(*this)[i].AssignedDev = devName
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
