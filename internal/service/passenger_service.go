package service

import (
	"airline-management-system/internal/models/passenger"
	"airline-management-system/internal/repository"
	"fmt"
)

type PassengerService struct {
	repo *repository.PassengerRepository
}

func NewPassengerService(repo *repository.PassengerRepository) *PassengerService {
	return &PassengerService{repo}
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
