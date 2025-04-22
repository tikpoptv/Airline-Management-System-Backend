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

// func (s *MaintenanceService) UpdateLog(id uint, req *maintenance.UpdateMaintenanceLogRequest) error {
// 	updates := make(map[string]interface{})

// 	if req.AircraftID != nil {
// 		updates["aircraft_id"] = *req.AircraftID
// 	}
// 	if req.DateOfMaintenance != nil {
// 		updates["date_of_maintenance"] = *req.DateOfMaintenance
// 	}
// 	if req.Details != nil {
// 		updates["details"] = *req.Details
// 	}
// 	if req.MaintenanceLocation != nil {
// 		updates["maintenance_location"] = *req.MaintenanceLocation
// 	}

// 	if len(updates) == 0 {
// 		return errors.New("no fields to update")
// 	}

// 	tx := s.repo.DB().Model(&maintenance.MaintenanceLog{}).
// 		Where("log_id = ?", id).
// 		Updates(updates)

// 	if tx.Error != nil {
// 		return tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return errors.New("maintenance log not found")
// 	}

// 	return nil
// }

// func (s *MaintenanceService) DeleteLogByID(id uint) error {
// 	found, err := s.repo.DeleteLogByID(id)
// 	if err != nil {
// 		return err
// 	}
// 	if !found {
// 		return errors.New("maintenance log not found")
// 	}
// 	return nil
// }
