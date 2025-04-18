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

func (s *CrewService) GetAllCrew() ([]crew.Crew, error) {
	return s.repo.GetAllCrew()
}
