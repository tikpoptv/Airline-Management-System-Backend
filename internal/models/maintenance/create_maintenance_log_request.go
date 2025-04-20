package maintenance

type CreateMaintenanceLogRequest struct {
	AircraftID          uint   `json:"aircraft_id" validate:"required"`
	DateOfMaintenance   string `json:"date_of_maintenance" validate:"required"`
	Details             string `json:"details"`
	MaintenanceLocation string `json:"maintenance_location"`
}
