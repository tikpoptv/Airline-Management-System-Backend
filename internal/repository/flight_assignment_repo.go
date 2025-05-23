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
		Select(`c.crew_id, c.name, c.passport_number, c.role, fca.role_in_flight, c.status`).
		Joins(`JOIN crew c ON c.crew_id = fca.crew_id`).
		Where("fca.flight_id = ?", flightID).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *FlightAssignmentRepository) GetAssignedFlightsByCrewID(crewID uint) ([]assignment.GetFlightCrewAssignment, error) {
	var assignments []assignment.GetFlightCrewAssignment
	err := r.db.
		Preload("Flight").
		Preload("Flight.Aircraft").
		Preload("Flight.Route").
		Preload("Flight.Route.FromAirport").
		Preload("Flight.Route.ToAirport").
		Preload("Crew").
		Where("crew_id = ?", crewID).
		Find(&assignments).Error
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

func (r *FlightAssignmentRepository) GetCrewForPassenger(flightID uint) (*assignment.FlightCrewInfoResponse, error) {
	var result assignment.FlightCrewInfoResponse

	// Get flight code
	if err := r.db.Table("flight").
		Select("flight_id, flight_code").
		Where("flight_id = ?", flightID).
		First(&result).Error; err != nil {
		return nil, err
	}

	// Get crew members
	var crewMembers []assignment.CrewInfoResponse
	err := r.db.Table("flight_crew_assignment AS fca").
		Select("c.name, c.role, fca.role_in_flight").
		Joins("JOIN crew c ON c.crew_id = fca.crew_id").
		Where("fca.flight_id = ? AND c.status = ?", flightID, "active").
		Scan(&crewMembers).Error

	if err != nil {
		return nil, err
	}

	result.CrewMembers = crewMembers
	return &result, nil
}

func (r *FlightAssignmentRepository) GetDB() *gorm.DB {
	return r.db
}
