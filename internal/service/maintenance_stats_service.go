package service

import (
	"airline-management-system/internal/models/maintenance"
	"airline-management-system/internal/repository"
)

type MaintenanceStatsService struct {
	repo *repository.MaintenanceStatsRepository
}

func NewMaintenanceStatsService(repo *repository.MaintenanceStatsRepository) *MaintenanceStatsService {
	return &MaintenanceStatsService{repo}
}

func (s *MaintenanceStatsService) GetMaintenanceStats() (*maintenance.MaintenanceStats, error) {
	return s.repo.GetMaintenanceStats()
}
