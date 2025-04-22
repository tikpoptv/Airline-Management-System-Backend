package maintenance

type CreateMaintenanceLogRequest struct {
	AircraftID          uint   `json:"aircraft_id" validate:"required"`
	DateOfMaintenance   string `json:"date_of_maintenance" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Details             string `json:"details"`
	MaintenanceLocation string `json:"maintenance_location"`
	Status              string `json:"status" validate:"omitempty,oneof=Pending In Progress Completed Cancelled"`
	AssignedTo          *uint  `json:"assigned_to"`
}
