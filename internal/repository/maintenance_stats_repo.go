package repository

import (
	"airline-management-system/internal/models/maintenance"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type MaintenanceStatsRepository struct {
	db *gorm.DB
}

func NewMaintenanceStatsRepository(db *gorm.DB) *MaintenanceStatsRepository {
	return &MaintenanceStatsRepository{db}
}

func (r *MaintenanceStatsRepository) GetMaintenanceStats() (*maintenance.MaintenanceStats, error) {
	var stats maintenance.MaintenanceStats

	// Get counts for different statuses
	if err := r.db.Table("maintenance_log").
		Where("status = ?", "Pending").
		Count(&stats.MaintenanceCount.Scheduled).Error; err != nil {
		return nil, fmt.Errorf("error counting scheduled maintenance: %v", err)
	}

	if err := r.db.Table("maintenance_log").
		Where("status = ?", "In Progress").
		Count(&stats.MaintenanceCount.InProgress).Error; err != nil {
		return nil, fmt.Errorf("error counting in-progress maintenance: %v", err)
	}

	if err := r.db.Table("maintenance_log").
		Where("status = ?", "Completed").
		Count(&stats.MaintenanceCount.Completed).Error; err != nil {
		return nil, fmt.Errorf("error counting completed maintenance: %v", err)
	}

	if err := r.db.Table("maintenance_log").
		Where("status = ?", "Cancelled").
		Count(&stats.MaintenanceCount.Delayed).Error; err != nil {
		return nil, fmt.Errorf("error counting cancelled maintenance: %v", err)
	}

	// Get today's maintenance logs
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := r.db.Table("maintenance_log").
		Select("log_id, aircraft_id, date_of_maintenance, details, maintenance_location, status, assigned_to").
		Where("date_of_maintenance BETWEEN ? AND ?", startOfDay, endOfDay).
		Order("date_of_maintenance ASC").
		Find(&stats.TodayMaintenance).Error

	if err != nil {
		return nil, fmt.Errorf("error fetching today's maintenance: %v", err)
	}

	return &stats, nil
}
