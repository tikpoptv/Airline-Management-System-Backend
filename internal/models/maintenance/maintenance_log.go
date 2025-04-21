package maintenance

import (
	"airline-management-system/internal/models/aircraft"
	"airline-management-system/internal/models/user"
	"time"
)

type MaintenanceLog struct {
	ID                  uint      `gorm:"column:log_id;primaryKey" json:"log_id"`
	AircraftID          uint      `json:"aircraft_id"`
	DateOfMaintenance   time.Time `json:"date_of_maintenance"`
	Details             string    `json:"details"`
	MaintenanceLocation string    `json:"maintenance_location"`
	Status              string    `json:"status"` // NEW
	AssignedTo          *uint     `json:"-"`

	AssignedUser *user.User        `gorm:"foreignKey:AssignedTo" json:"assigned_user"`
	Aircraft     aircraft.Aircraft `gorm:"foreignKey:AircraftID" json:"aircraft"`
}

func (MaintenanceLog) TableName() string {
	return "maintenance_log"
}
