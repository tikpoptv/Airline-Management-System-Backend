package repository

import (
	"airline-management-system/internal/models/aircraft"
	"airline-management-system/internal/models/airport"
	"airline-management-system/internal/models/crew"
	"airline-management-system/internal/models/dashboard"
	"airline-management-system/internal/models/route"

	"fmt"
	"time"

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

func (r *DashboardRepository) GetTodayCrewSchedule(limit int) ([]dashboard.CrewScheduleResponse, error) {
	var schedules []dashboard.CrewScheduleResponse

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	query := r.db.Table("flight_crew_assignment fca").
		Select(`
			c.crew_id,
			c.name,
			c.role,
			fca.role_in_flight,
			f.flight_number as flight_code,
			fromAirport.name as from_airport,
			toAirport.name as to_airport,
			f.departure_time,
			f.arrival_time,
			f.flight_status as status
		`).
		Joins("JOIN crew c ON c.crew_id = fca.crew_id").
		Joins("JOIN flight f ON f.flight_id = fca.flight_id").
		Joins("JOIN route r ON r.route_id = f.route_id").
		Joins("JOIN airport fromAirport ON fromAirport.airport_id = r.from_airport").
		Joins("JOIN airport toAirport ON toAirport.airport_id = r.to_airport").
		Where("f.departure_time BETWEEN ? AND ?", startOfDay, endOfDay).
		Order("f.departure_time ASC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Scan(&schedules).Error; err != nil {
		return nil, fmt.Errorf("error fetching today's crew schedule: %v", err)
	}

	return schedules, nil
}
