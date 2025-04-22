package repository

import (
	"airline-management-system/internal/models/maintenance"

	"gorm.io/gorm"
)

type MaintenanceTaskRepository struct {
	db *gorm.DB
}

func NewMaintenanceTaskRepository(db *gorm.DB) *MaintenanceTaskRepository {
	return &MaintenanceTaskRepository{db}
}

func (r *MaintenanceTaskRepository) GetTasksByUser(userID uint) ([]maintenance.MaintenanceLog, error) {
	var tasks []maintenance.MaintenanceLog
	err := r.db.
		Preload("Aircraft").
		Preload("AssignedUser").
		Where("assigned_to = ?", userID).
		Order("log_id DESC").
		Find(&tasks).Error

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *MaintenanceTaskRepository) UpdateTaskStatus(logID uint, update map[string]interface{}) error {
	tx := r.db.Model(&maintenance.MaintenanceLog{}).
		Where("log_id = ?", logID).
		Updates(update)

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *MaintenanceTaskRepository) IsTaskOwnedByUser(logID uint, userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&maintenance.MaintenanceLog{}).
		Where("log_id = ? AND assigned_to = ?", logID, userID).
		Count(&count).Error
	return count > 0, err
}
