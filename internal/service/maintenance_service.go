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

func (s *MaintenanceService) GetAllLogs() ([]maintenance.MaintenanceLog, error) {
	return s.repo.GetAllLogs()
}

func (s *MaintenanceService) CreateLog(req *maintenance.CreateMaintenanceLogRequest) (*maintenance.MaintenanceLog, error) {
	// ตรวจสอบว่า format ถูกต้อง
	if _, err := time.Parse(time.RFC3339, req.DateOfMaintenance); err != nil {
		return nil, errors.New("invalid datetime format: use RFC3339")
	}

	log := &maintenance.MaintenanceLog{
		AircraftID:          req.AircraftID,
		DateOfMaintenance:   req.DateOfMaintenance,
		Details:             req.Details,
		MaintenanceLocation: req.MaintenanceLocation,
	}

	if err := s.repo.CreateLog(log); err != nil {
		return nil, err
	}

	return log, nil
}

func (s *MaintenanceService) GetLogByID(id uint) (*maintenance.MaintenanceLog, error) {
	return s.repo.GetLogByID(id)
}

func (s *MaintenanceService) UpdateLog(id uint, req *maintenance.UpdateMaintenanceLogRequest) error {
	updates := make(map[string]interface{})

	if req.AircraftID != nil {
		updates["aircraft_id"] = *req.AircraftID
	}
	if req.DateOfMaintenance != nil {
		updates["date_of_maintenance"] = *req.DateOfMaintenance
	}
	if req.Details != nil {
		updates["details"] = *req.Details
	}
	if req.MaintenanceLocation != nil {
		updates["maintenance_location"] = *req.MaintenanceLocation
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	tx := s.repo.DB().Model(&maintenance.MaintenanceLog{}).
		Where("log_id = ?", id).
		Updates(updates)

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("maintenance log not found")
	}

	return nil
}
