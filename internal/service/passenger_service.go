package service

import (
	"airline-management-system/internal/models/passenger"
	"airline-management-system/internal/repository"
	"fmt"
	"math"
)

type PassengerService struct {
	repo *repository.PassengerRepository
}

func NewPassengerService(repo *repository.PassengerRepository) *PassengerService {
	return &PassengerService{repo}
}

func (s *PassengerService) GetAllPassengers(page, pageSize int) (*passenger.AllPassengersResponse, error) {
	// Validate page and pageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// Get passengers from repository
	passengers, total, err := s.repo.GetAllPassengers(page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get passengers: %v", err)
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	// Create response
	response := &passenger.AllPassengersResponse{
		Passengers: passengers,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}

	return response, nil
}

func (s *PassengerService) GetPassengersByFlightID(flightID uint) ([]passenger.PassengerResponse, error) {
	passengers, err := s.repo.GetPassengersByFlightID(flightID)
	if err != nil {
		return nil, fmt.Errorf("failed to get passengers: %v", err)
	}
	return passengers, nil
}

func (s *PassengerService) GetPassengerByID(id uint) (*passenger.PassengerDetailResponse, error) {
	passenger, err := s.repo.GetPassengerByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get passenger: %v", err)
	}
	if passenger == nil {
		return nil, fmt.Errorf("passenger not found")
	}
	return passenger, nil
}

func (s *PassengerService) SearchPassengers(query string, page, pageSize int) (*passenger.AllPassengersResponse, error) {
	// Validate page and pageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// Get passengers from repository
	passengers, total, err := s.repo.SearchPassengers(query, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to search passengers: %v", err)
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	// Create response
	response := &passenger.AllPassengersResponse{
		Passengers: passengers,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}

	return response, nil
}
