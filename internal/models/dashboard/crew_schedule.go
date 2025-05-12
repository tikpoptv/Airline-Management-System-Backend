package dashboard

import (
	"time"
)

type CrewScheduleResponse struct {
	CrewID        uint      `json:"crew_id" gorm:"column:crew_id"`
	Name          string    `json:"name" gorm:"column:name"`
	Role          string    `json:"role" gorm:"column:role"`
	RoleInFlight  string    `json:"role_in_flight" gorm:"column:role_in_flight"`
	FlightCode    string    `json:"flight_code" gorm:"column:flight_code"`
	FromAirport   string    `json:"from_airport" gorm:"column:from_airport"`
	ToAirport     string    `json:"to_airport" gorm:"column:to_airport"`
	DepartureTime time.Time `json:"departure_time" gorm:"column:departure_time"`
	ArrivalTime   time.Time `json:"arrival_time" gorm:"column:arrival_time"`
	Status        string    `json:"status" gorm:"column:status"`
}
