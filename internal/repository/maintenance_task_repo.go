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
