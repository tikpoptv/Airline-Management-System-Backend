package service

import (
	"airline-management-system/internal/models/aircraft"
	"airline-management-system/internal/repository"
	"errors"
	"time"
)

type AircraftService struct {
	repo *repository.AircraftRepository
}

func NewAircraftService(repo *repository.AircraftRepository) *AircraftService {
	return &AircraftService{repo}
}

func (s *AircraftService) GetAllAircraft() ([]aircraft.Aircraft, error) {
	return s.repo.GetAllAircraft()
}

func (s *AircraftService) CreateAircraft(req *aircraft.CreateAircraftRequest) (*aircraft.Aircraft, error) {
	currentYear := time.Now().Year()
	if req.ManufactureYear < 1950 || req.ManufactureYear > currentYear {
		return nil, errors.New("invalid manufacture year")
	}

	newAircraft := &aircraft.Aircraft{
		Model:             req.Model,
		ManufactureYear:   req.ManufactureYear,
		Capacity:          req.Capacity,
		AirlineOwner:      req.AirlineOwner,
		MaintenanceStatus: req.MaintenanceStatus,
		AircraftHistory:   req.AircraftHistory,
	}

	if err := s.repo.CreateAircraft(newAircraft); err != nil {
		return nil, err
	}
	return newAircraft, nil
}

func (s *AircraftService) GetAircraftByID(id uint) (*aircraft.Aircraft, error) {
	return s.repo.GetAircraftByID(id)
}

func (s *AircraftService) UpdateAircraft(id uint, req *aircraft.UpdateAircraftRequest) (*aircraft.Aircraft, error) {
	updates := make(map[string]interface{})

	if req.Model != nil {
		updates["model"] = *req.Model
	}
	if req.ManufactureYear != nil {
		updates["manufacture_year"] = *req.ManufactureYear
	}
	if req.Capacity != nil {
		updates["capacity"] = *req.Capacity
	}
	if req.AirlineOwner != nil {
		updates["airline_owner"] = *req.AirlineOwner
	}
	if req.MaintenanceStatus != nil {
		status := *req.MaintenanceStatus
		valid := status == "Operational" || status == "In Maintenance" || status == "Retired"
		if !valid {
			return nil, errors.New("maintenance_status must be one of: Operational, In Maintenance, Retired")
		}
		updates["maintenance_status"] = status
	}
	if req.AircraftHistory != nil {
		updates["aircraft_history"] = *req.AircraftHistory
	}

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	return s.repo.UpdateAircraft(id, updates)
}

func (s *AircraftService) DeleteAircraft(id uint) error {
	return s.repo.DeleteAircraft(id)
}
