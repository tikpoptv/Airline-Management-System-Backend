package service

import (
	"airline-management-system/internal/models/airport"
	"airline-management-system/internal/repository"
)

type AirportService struct {
	repo *repository.AirportRepository
}

func NewAirportService(repo *repository.AirportRepository) *AirportService {
	return &AirportService{repo}
}

func (s *AirportService) GetAllAirports() ([]airport.Airport, error) {
	return s.repo.GetAllAirports()
}

func (s *AirportService) CreateAirport(req *airport.CreateAirportRequest) (*airport.Airport, error) {
	newAirport := &airport.Airport{
		IATACode:  req.IATACode,
		Name:      req.Name,
		City:      req.City,
		Country:   req.Country,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Timezone:  req.Timezone,
		Status:    req.Status,
	}

	if err := s.repo.CreateAirport(newAirport); err != nil {
		return nil, err
	}

	return newAirport, nil
}

func (s *AirportService) UpdateAirport(id uint, req *airport.UpdateAirportRequest) error {
	return s.repo.UpdateAirport(id, req)
}
