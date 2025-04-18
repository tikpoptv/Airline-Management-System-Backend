package flight

import (
	"airline-management-system/internal/models/aircraft"
	"airline-management-system/internal/models/route"
)

type Flight struct {
	ID                 uint    `gorm:"column:flight_id;primaryKey" json:"flight_id"`
	FlightNumber       string  `json:"flight_number"`
	AircraftID         uint    `gorm:"column:aircraft_id" json:"-"`
	RouteID            uint    `gorm:"column:route_id" json:"-"`
	DepartureTime      string  `json:"departure_time"`
	ArrivalTime        string  `json:"arrival_time"`
	FlightStatus       string  `json:"flight_status"`
	CancellationReason *string `json:"cancellation_reason"`

	Aircraft aircraft.Aircraft `gorm:"foreignKey:AircraftID" json:"aircraft"`
	Route    route.Route       `gorm:"foreignKey:RouteID" json:"route"`
}

func (Flight) TableName() string {
	return "flight"
}
