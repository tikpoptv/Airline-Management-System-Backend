package service

import (
	"airline-management-system/internal/models/assignment"
	"airline-management-system/internal/repository"
)

type CrewAssignmentService struct {
	repo *repository.FlightAssignmentRepository
}

func NewCrewAssignmentService(repo *repository.FlightAssignmentRepository) *CrewAssignmentService {
	return &CrewAssignmentService{repo}
}

func (s *CrewAssignmentService) GetAssignedFlightsByCrewID(crewID uint) ([]assignment.GetFlightCrewAssignment, error) {
	return s.repo.GetAssignedFlightsByCrewID(crewID)
}
