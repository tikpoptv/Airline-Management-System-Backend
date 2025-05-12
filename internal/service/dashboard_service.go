package service

import (
	"airline-management-system/internal/models/dashboard"
	"airline-management-system/internal/repository"
)

type DashboardService struct {
	repo *repository.DashboardRepository
}

func NewDashboardService(repo *repository.DashboardRepository) *DashboardService {
	return &DashboardService{repo}
}

func (s *DashboardService) GetDashboardStats() (*dashboard.DashboardStats, error) {
	return s.repo.GetDashboardStats()
}
