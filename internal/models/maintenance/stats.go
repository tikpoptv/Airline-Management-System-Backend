package maintenance

import "time"

type MaintenanceStats struct {
	MaintenanceCount struct {
		Scheduled  int64 `json:"scheduled"`   // Pending
		InProgress int64 `json:"in_progress"` // In Progress
		Completed  int64 `json:"completed"`   // Completed
		Delayed    int64 `json:"delayed"`     // Cancelled
	} `json:"maintenance_stats"`

	TodayMaintenance []MaintenanceLogToday `json:"today_maintenance"`
}

type MaintenanceLogToday struct {
	LogID               uint      `json:"log_id" gorm:"column:log_id"`
	AircraftID          uint      `json:"aircraft_id" gorm:"column:aircraft_id"`
	DateOfMaintenance   time.Time `json:"date_of_maintenance" gorm:"column:date_of_maintenance"`
	Details             string    `json:"details" gorm:"column:details"`
	MaintenanceLocation string    `json:"maintenance_location" gorm:"column:maintenance_location"`
	Status              string    `json:"status" gorm:"column:status"`
	AssignedTo          *uint     `json:"assigned_to" gorm:"column:assigned_to"` // nullable
}

func (MaintenanceLogToday) TableName() string {
	return "maintenance_log"
}
