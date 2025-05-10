package assignment

type CrewAssignmentResponse struct {
	CrewID         uint   `json:"crew_id"`
	Name           string `json:"name"`
	PassportNumber string `json:"passport_number"`
	Role           string `json:"role"`
	RoleInFlight   string `json:"role_in_flight"`
	Status         string `json:"status"`
}
