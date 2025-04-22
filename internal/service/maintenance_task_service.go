package service

import (
	"airline-management-system/internal/models/maintenance"
	"airline-management-system/internal/repository"
)

type MaintenanceTaskService struct {
	repo *repository.MaintenanceTaskRepository
}

func NewMaintenanceTaskService(repo *repository.MaintenanceTaskRepository) *MaintenanceTaskService {
	return &MaintenanceTaskService{repo}
}

func (s *MaintenanceTaskService) GetTasksByUser(userID uint) ([]maintenance.MaintenanceLog, error) {
	return s.repo.GetTasksByUser(userID)
}
