package aircraft

type CreateAircraftRequest struct {
	Model             string `json:"model" validate:"required"`
	ManufactureYear   int    `json:"manufacture_year" validate:"required,min=1950,max=2100"`
	Capacity          int    `json:"capacity" validate:"required,gt=0"`
	AirlineOwner      string `json:"airline_owner" validate:"required"`
	MaintenanceStatus string `json:"maintenance_status" validate:"required,oneof='Operational' 'In Maintenance' 'Retired'"`
	AircraftHistory   string `json:"aircraft_history"`
}
