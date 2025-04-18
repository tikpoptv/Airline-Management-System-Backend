package aircraft

type UpdateAircraftRequest struct {
	Model             *string `json:"model"`
	ManufactureYear   *int    `json:"manufacture_year"`   // Optional
	Capacity          *int    `json:"capacity"`           // Optional
	AirlineOwner      *string `json:"airline_owner"`      // Optional
	MaintenanceStatus *string `json:"maintenance_status"` // Optional
	AircraftHistory   *string `json:"aircraft_history"`   // Optional
}
