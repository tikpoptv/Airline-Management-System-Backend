package aircraft

type Aircraft struct {
	ID                uint   `gorm:"column:aircraft_id;primaryKey" json:"aircraft_id"`
	Model             string `json:"model"`
	ManufactureYear   int    `json:"manufacture_year"`
	Capacity          int    `json:"capacity"`
	AirlineOwner      string `json:"airline_owner"`
	MaintenanceStatus string `json:"maintenance_status"`
	AircraftHistory   string `json:"aircraft_history"`
}

func (Aircraft) TableName() string {
	return "aircraft"
}
