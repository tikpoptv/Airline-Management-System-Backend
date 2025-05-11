package service

import (
	"errors"
	"strings"

	"airline-management-system/internal/models/assignment"
	"airline-management-system/internal/repository"

	"gorm.io/gorm"
)

type FlightAssignmentService struct {
	repo *repository.FlightAssignmentRepository
	db   *gorm.DB
}

func NewFlightAssignmentService(repo *repository.FlightAssignmentRepository) *FlightAssignmentService {
	return &FlightAssignmentService{
		repo: repo,
		db:   repo.GetDB(),
	}
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

func (s *FlightAssignmentService) GetFlightCrewInfo(flightID uint, userID uint) (*assignment.FlightCrewInfoResponse, error) {
	// Check if user is admin or crew
	var roleCount int64
	s.db.Table("users").
		Where("user_id = ? AND role IN ('admin', 'crew')", userID).
		Count(&roleCount)

	if roleCount == 0 {
		return nil, errors.New("unauthorized access")
	}

	return s.repo.GetCrewForPassenger(flightID)
}
