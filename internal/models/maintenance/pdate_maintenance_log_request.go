package maintenance

type UpdateMaintenanceLogRequest struct {
	AircraftID          *uint   `json:"aircraft_id,omitempty"`
	DateOfMaintenance   *string `json:"date_of_maintenance,omitempty"`
	Details             *string `json:"details,omitempty"`
	MaintenanceLocation *string `json:"maintenance_location,omitempty"`
}
