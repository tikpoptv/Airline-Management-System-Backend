package payment

import (
	"airline-management-system/internal/models/ticket"
	"time"
)

type Payment struct {
	PaymentID     uint          `json:"payment_id" gorm:"column:payment_id;primaryKey"`
	TicketID      uint          `json:"ticket_id" gorm:"column:ticket_id"`
	Ticket        ticket.Ticket `json:"ticket" gorm:"foreignKey:TicketID"`
	PaymentMethod string        `json:"payment_method"`
	Amount        float64       `json:"amount"`
	PaymentDate   time.Time     `json:"payment_date"`
	PaymentStatus string        `json:"payment_status"`
	TransactionID string        `json:"transaction_id"`
}

func (Payment) TableName() string {
	return "payment"
}
