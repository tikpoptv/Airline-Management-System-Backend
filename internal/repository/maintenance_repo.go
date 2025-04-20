package repository

import (
	"airline-management-system/internal/models/maintenance"

	"gorm.io/gorm"
)

type MaintenanceRepository struct {
	db *gorm.DB
}

func NewMaintenanceRepository(db *gorm.DB) *MaintenanceRepository {
	return &MaintenanceRepository{db}
}

func (r *MaintenanceRepository) GetAllLogs() ([]maintenance.MaintenanceLog, error) {
	var logs []maintenance.MaintenanceLog
	err := r.db.Preload("Aircraft").
		Order("date_of_maintenance DESC").
		Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}
