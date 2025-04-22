package service

import (
	"airline-management-system/internal/models/assignment"
	"airline-management-system/internal/models/maintenance"
	"airline-management-system/internal/repository"
	"errors"
)

type CrewAssignmentService struct {
	crewRepo       *repository.CrewRepository
	assignmentRepo *repository.FlightAssignmentRepository
}

func NewCrewAssignmentService(
	crewRepo *repository.CrewRepository,
	assignmentRepo *repository.FlightAssignmentRepository,
) *CrewAssignmentService {
	return &CrewAssignmentService{
		crewRepo:       crewRepo,
		assignmentRepo: assignmentRepo,
	}
}

func (s *CrewAssignmentService) GetAssignedFlightsByUserID(userID uint) ([]assignment.GetFlightCrewAssignment, error) {
	crewID, err := s.crewRepo.GetCrewIDByUserID(userID)
	if err != nil {
		return nil, errors.New("crew not found for this user")
	}
	return s.assignmentRepo.GetAssignedFlightsByCrewID(crewID)
}

func (s *MaintenanceTaskService) UpdateTaskStatus(logID uint, userID uint, req *maintenance.UpdateTaskStatusRequest) error {
	ok, err := s.repo.IsTaskOwnedByUser(logID, userID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("you are not assigned to this task")
	}

	update := map[string]interface{}{
		"status": req.Status,
	}
	if req.Details != nil {
		update["details"] = *req.Details
	}

	return s.repo.UpdateTaskStatus(logID, update)
}
