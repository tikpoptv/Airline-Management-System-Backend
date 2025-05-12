package service

import (
	"airline-management-system/internal/models/dashboard"
	"airline-management-system/internal/repository"
	"fmt"
)

type DashboardService struct {
	repo *repository.DashboardRepository
}

func NewDashboardService(repo *repository.DashboardRepository) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetDashboardStats() (*dashboard.DashboardStats, error) {
	return s.repo.GetDashboardStats()
}

func (s *DashboardService) GetTodayCrewSchedule(limit int) ([]dashboard.CrewScheduleResponse, error) {
	if limit < 0 {
		return nil, fmt.Errorf("invalid limit: %d", limit)
	}

	schedules, err := s.repo.GetTodayCrewSchedule(limit)
	if err != nil {
		return nil, fmt.Errorf("error from repository: %v", err)
	}

	return schedules, nil
}
