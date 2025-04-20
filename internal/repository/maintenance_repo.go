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
		Order("log_id ASC").
		Find(&logs).Error

	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *MaintenanceRepository) CreateLog(log *maintenance.MaintenanceLog) error {
	return r.db.Create(log).Error
}

func (r *MaintenanceRepository) GetLogByID(id uint) (*maintenance.MaintenanceLog, error) {
	var log maintenance.MaintenanceLog
	err := r.db.Preload("Aircraft").First(&log, "log_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *MaintenanceRepository) UpdateLog(id uint, updates map[string]interface{}) error {
	return r.db.Model(&maintenance.MaintenanceLog{}).
		Where("log_id = ?", id).
		Updates(updates).Error
}

func (r *MaintenanceRepository) DB() *gorm.DB {
	return r.db
}
