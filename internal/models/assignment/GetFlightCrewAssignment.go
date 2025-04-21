package assignment

import (
	"airline-management-system/internal/models/crew"
	"airline-management-system/internal/models/flight"
)

type GetFlightCrewAssignment struct {
	FlightID     uint   `json:"-" gorm:"column:flight_id;primaryKey"`
	CrewID       uint   `json:"-" gorm:"column:crew_id;primaryKey"`
	RoleInFlight string `json:"role_in_flight"`

	Flight flight.Flight `gorm:"foreignKey:FlightID;references:ID" json:"flight"`
	Crew   crew.Crew     `gorm:"foreignKey:CrewID;references:ID" json:"crew"`
}

func (GetFlightCrewAssignment) TableName() string {
	return "flight_crew_assignment"
}
