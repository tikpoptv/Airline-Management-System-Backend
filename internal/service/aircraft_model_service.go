package service

import (
	"airline-management-system/internal/models/aircraftmodel"
	"airline-management-system/internal/repository"
	"context"
)

type AircraftModelService interface {
	GetAllAircraftModels(ctx context.Context) ([]aircraftmodel.AircraftModel, error)
}

type aircraftModelService struct {
	repo repository.AircraftModelRepository
}

func NewAircraftModelService(repo repository.AircraftModelRepository) AircraftModelService {
	return &aircraftModelService{repo: repo}
}

func (s *aircraftModelService) GetAllAircraftModels(ctx context.Context) ([]aircraftmodel.AircraftModel, error) {
	return s.repo.GetAllAircraftModels(ctx)
}
