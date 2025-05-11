package assignment

type CrewInfoResponse struct {
	Name         string `json:"name"`
	Role         string `json:"role"`
	RoleInFlight string `json:"role_in_flight"`
}

type FlightCrewInfoResponse struct {
	FlightID    uint               `json:"flight_id"`
	FlightCode  string             `json:"flight_code"`
	CrewMembers []CrewInfoResponse `json:"crew_members"`
}
