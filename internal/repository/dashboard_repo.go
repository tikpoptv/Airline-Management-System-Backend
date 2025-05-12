package repository

import (
	"airline-management-system/internal/models/aircraft"
	"airline-management-system/internal/models/airport"
	"airline-management-system/internal/models/crew"
	"airline-management-system/internal/models/dashboard"
	"airline-management-system/internal/models/route"

	"gorm.io/gorm"
)

type DashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepository {
	return &DashboardRepository{db}
}

func (r *DashboardRepository) GetDashboardStats() (*dashboard.DashboardStats, error) {
	var stats dashboard.DashboardStats

	// Get Aircraft Stats
	r.db.Model(&aircraft.Aircraft{}).Count(&stats.TotalAircrafts)
	r.db.Model(&aircraft.Aircraft{}).Where("maintenance_status = ?", "Operational").Count(&stats.ActiveAircrafts)
	r.db.Model(&aircraft.Aircraft{}).Where("maintenance_status = ?", "In Maintenance").Count(&stats.MaintenanceAircrafts)

	// Get Crew Stats
	r.db.Model(&crew.Crew{}).Count(&stats.TotalCrews)
	r.db.Model(&crew.Crew{}).Where("status = ?", "active").Count(&stats.ActiveCrews)

	// Get Route & Airport Stats
	r.db.Model(&route.Route{}).Count(&stats.TotalRoutes)
	r.db.Model(&route.Route{}).Where("status = ?", "active").Count(&stats.ActiveRoutes)
	r.db.Model(&airport.Airport{}).Count(&stats.TotalAirports)

	return &stats, nil
}
