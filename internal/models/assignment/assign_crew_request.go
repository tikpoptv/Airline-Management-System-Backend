package assignment

type AssignCrewRequest struct {
	CrewID       uint   `json:"crew_id" validate:"required"`
	RoleInFlight string `json:"role_in_flight" validate:"required"`
}
