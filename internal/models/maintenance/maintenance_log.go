package maintenance

import "airline-management-system/internal/models/aircraft"

type MaintenanceLog struct {
	LogID               uint              `gorm:"column:log_id;primaryKey" json:"log_id"`
	AircraftID          uint              `gorm:"column:aircraft_id" json:"-"`
	Aircraft            aircraft.Aircraft `gorm:"foreignKey:AircraftID" json:"aircraft"`
	DateOfMaintenance   string            `gorm:"column:date_of_maintenance" json:"date_of_maintenance"`
	Details             string            `json:"details"`
	MaintenanceLocation string            `json:"maintenance_location"`
}

func (MaintenanceLog) TableName() string {
	return "maintenance_log"
}
