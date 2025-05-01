package airlineowner

import "time"

type AirlineOwner struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	Alliance    string    `json:"alliance"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
