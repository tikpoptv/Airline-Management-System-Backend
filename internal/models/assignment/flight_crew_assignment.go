package assignment

type FlightCrewAssignment struct {
	FlightID     uint   `gorm:"column:flight_id;primaryKey" json:"flight_id"`
	CrewID       uint   `gorm:"column:crew_id;primaryKey" json:"crew_id"`
	RoleInFlight string `gorm:"column:role_in_flight" json:"role_in_flight"`
}

func (FlightCrewAssignment) TableName() string {
	return "flight_crew_assignment"
}
