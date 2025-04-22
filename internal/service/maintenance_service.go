package service

import (
	"airline-management-system/internal/models/maintenance"
	"airline-management-system/internal/repository"
	"errors"
	"time"
)

type MaintenanceService struct {
	repo *repository.MaintenanceRepository
}

func NewMaintenanceService(repo *repository.MaintenanceRepository) *MaintenanceService {
	return &MaintenanceService{repo}
}

func (s *MaintenanceService) GetAllLogs(filters map[string]interface{}) ([]maintenance.MaintenanceLog, error) {
	return s.repo.GetAllLogs(filters)
}

func (s *MaintenanceService) CreateLog(req *maintenance.CreateMaintenanceLogRequest) (*maintenance.MaintenanceLog, error) {
	date, err := time.Parse(time.RFC3339, req.DateOfMaintenance)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	log := &maintenance.MaintenanceLog{
		AircraftID:          req.AircraftID,
		DateOfMaintenance:   date,
		Details:             req.Details,
		MaintenanceLocation: req.MaintenanceLocation,
		Status:              req.Status,
		AssignedTo:          req.AssignedTo,
	}

	if log.Status == "" {
		log.Status = "Pending"
	}

	err = s.repo.CreateLog(log)
	if err != nil {
		return nil, err
	}
	return log, nil
}

func (s *MaintenanceService) GetLogByID(id uint) (*maintenance.MaintenanceLog, error) {
	return s.repo.GetLogByID(id)
}
