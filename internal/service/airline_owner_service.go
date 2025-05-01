package service

import (
	"airline-management-system/internal/models/airlineowner"
	"airline-management-system/internal/repository"
	"context"
)

type AirlineOwnerService interface {
	GetAllAirlineOwners(ctx context.Context) ([]airlineowner.AirlineOwner, error)
}

type airlineOwnerService struct {
	repo repository.AirlineOwnerRepository
}

func NewAirlineOwnerService(repo repository.AirlineOwnerRepository) AirlineOwnerService {
	return &airlineOwnerService{repo: repo}
}

func (s *airlineOwnerService) GetAllAirlineOwners(ctx context.Context) ([]airlineowner.AirlineOwner, error) {
	return s.repo.GetAllAirlineOwners(ctx)
}
