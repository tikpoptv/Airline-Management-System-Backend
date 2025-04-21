package ticket

type Ticket struct {
	TicketID      uint   `json:"ticket_id" gorm:"column:ticket_id;primaryKey"`
	SeatNumber    string `json:"seat_number"`
	TicketStatus  string `json:"ticket_status"`
	CheckInStatus string `json:"check_in_status"`
	FlightID      uint   `json:"flight_id"`
	PassengerID   uint   `json:"passenger_id"`
}

func (Ticket) TableName() string {
	return "ticket"
}
