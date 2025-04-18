package crew

type CrewFlightHoursResponse struct {
	CrewID      uint    `json:"crew_id"`
	Name        string  `json:"name"`
	FlightHours float64 `json:"flight_hours"`
}
