package service

import (
	"airline-management-system/internal/models/crew"
	"airline-management-system/internal/repository"
	"errors"
	"strings"
)

type CrewService struct {
	repo *repository.CrewRepository
}

func NewCrewService(repo *repository.CrewRepository) *CrewService {
	return &CrewService{repo}
}

func (s *CrewService) GetAllCrew() ([]crew.GetCrew, error) {
	return s.repo.GetAllCrew()
}

func (s *CrewService) CreateCrew(req *crew.CreateCrewRequest) (*crew.Crew, error) {
	newCrew := &crew.Crew{
		Name:               req.Name,
		PassportNumber:     req.PassportNumber,
		Role:               req.Role,
		LicenseExpiryDate:  req.LicenseExpiryDate,
		PassportExpiryDate: req.PassportExpiryDate,
		FlightHours:        req.FlightHours,
		UserID:             req.UserID,
	}

	if err := s.repo.CreateCrew(newCrew); err != nil {
		return nil, err
	}

	return newCrew, nil
}

func (s *CrewService) GetCrewByID(id uint) (*crew.GetCrew, error) {
	return s.repo.GetCrewByID(id)
}

func (s *CrewService) UpdateCrew(id uint, req *crew.UpdateCrewRequest) error {
	updates := make(map[string]interface{})

	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.PassportNumber != nil {
		updates["passport_number"] = *req.PassportNumber
	}
	if req.Role != nil {
		role := strings.TrimSpace(*req.Role)
		if role != "Pilot" && role != "Co-Pilot" && role != "Attendant" && role != "Technician" {
			return errors.New("invalid role")
		}
		updates["role"] = role
	}
	if req.LicenseExpiryDate != nil {
		updates["license_expiry_date"] = *req.LicenseExpiryDate
	}
	if req.PassportExpiryDate != nil {
		updates["passport_expiry_date"] = *req.PassportExpiryDate
	}
	if req.FlightHours != nil {
		updates["flight_hours"] = *req.FlightHours
	}
	if req.UserID != nil {
		updates["user_id"] = *req.UserID
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	return s.repo.UpdateCrew(id, updates)
}

func (s *CrewService) DeleteCrew(id uint) (bool, error) {
	return s.repo.DeleteCrew(id)
}

func (s *CrewService) GetCrewFlightHours(id uint) (*crew.CrewFlightHoursResponse, error) {
	return s.repo.GetCrewFlightHours(id)
}
