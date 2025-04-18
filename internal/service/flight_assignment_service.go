package service

import (
	"errors"
	"strings"

	"airline-management-system/internal/models/assignment"
	"airline-management-system/internal/repository"
)

type FlightAssignmentService struct {
	repo *repository.FlightAssignmentRepository
}

func NewFlightAssignmentService(repo *repository.FlightAssignmentRepository) *FlightAssignmentService {
	return &FlightAssignmentService{repo}
}

func (s *FlightAssignmentService) AssignCrewToFlight(flightID uint, req *assignment.AssignCrewRequest) (*assignment.FlightCrewAssignment, error) {
	newAssignment := &assignment.FlightCrewAssignment{
		FlightID:     flightID,
		CrewID:       req.CrewID,
		RoleInFlight: req.RoleInFlight,
	}

	err := s.repo.AssignCrewToFlight(newAssignment)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errors.New("crew already assigned to this flight")
		}
		return nil, err
	}

	return newAssignment, nil
}

func (s *FlightAssignmentService) GetCrewByFlightID(flightID uint) ([]assignment.CrewAssignmentResponse, error) {
	return s.repo.GetCrewByFlightID(flightID)
}
