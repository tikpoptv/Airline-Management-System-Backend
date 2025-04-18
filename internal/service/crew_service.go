package service

import (
	"airline-management-system/internal/models/crew"
	"airline-management-system/internal/repository"
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
