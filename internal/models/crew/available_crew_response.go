package crew

type AvailableCrewResponse struct {
	CrewID       uint    `json:"crew_id"`
	Name         string  `json:"name"`
	Role         string  `json:"role"`
	FlightHours  float64 `json:"flight_hours"`
	Status       string  `json:"status"`
	LicenseValid bool    `json:"license_valid"`
}
