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
