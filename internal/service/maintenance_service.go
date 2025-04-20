package service

import (
	"airline-management-system/internal/models/maintenance"
	"airline-management-system/internal/repository"
)

type MaintenanceService struct {
	repo *repository.MaintenanceRepository
}

func NewMaintenanceService(repo *repository.MaintenanceRepository) *MaintenanceService {
	return &MaintenanceService{repo}
}

func (s *MaintenanceService) GetAllLogs() ([]maintenance.MaintenanceLog, error) {
	return s.repo.GetAllLogs()
}
