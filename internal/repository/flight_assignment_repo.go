package repository

import (
	"airline-management-system/internal/models/assignment"

	"gorm.io/gorm"
)

type FlightAssignmentRepository struct {
	db *gorm.DB
}

func NewFlightAssignmentRepository(db *gorm.DB) *FlightAssignmentRepository {
	return &FlightAssignmentRepository{db}
}

func (r *FlightAssignmentRepository) AssignCrewToFlight(a *assignment.FlightCrewAssignment) error {
	return r.db.Create(a).Error
}

func (r *FlightAssignmentRepository) GetCrewByFlightID(flightID uint) ([]assignment.CrewAssignmentResponse, error) {
	var result []assignment.CrewAssignmentResponse
	err := r.db.Table("flight_crew_assignment AS fca").
		Select(`c.crew_id, c.name, c.passport_number, c.role, fca.role_in_flight`).
		Joins(`JOIN crew c ON c.crew_id = fca.crew_id`).
		Where("fca.flight_id = ?", flightID).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}
	return result, nil
}
