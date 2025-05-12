package repository

import (
	"airline-management-system/internal/models/passenger"
	"fmt"

	"gorm.io/gorm"
)

type PassengerRepository struct {
	db *gorm.DB
}

func NewPassengerRepository(db *gorm.DB) *PassengerRepository {
	return &PassengerRepository{db}
}

// GetAllPassengers retrieves all passengers with pagination
func (r *PassengerRepository) GetAllPassengers(page, pageSize int) ([]passenger.PassengerResponse, int64, error) {
	var passengers []passenger.PassengerResponse
	var total int64

	// Count total records
	if err := r.db.Table("passenger").Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("error counting passengers: %v", err)
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Get paginated records
	err := r.db.Table("passenger").
		Select("passenger_id, name, passport_number, nationality, flight_id, special_requests, user_id").
		Order("passenger_id ASC").
		Limit(pageSize).
		Offset(offset).
		Scan(&passengers).Error

	if err != nil {
		return nil, 0, fmt.Errorf("error fetching passengers: %v", err)
	}

	return passengers, total, nil
}

// GetPassengersByFlightID retrieves passengers for a specific flight
func (r *PassengerRepository) GetPassengersByFlightID(flightID uint) ([]passenger.PassengerResponse, error) {
	var passengers []passenger.PassengerResponse
	err := r.db.Table("passenger").
		Select("passenger_id, name, passport_number, nationality, flight_id, special_requests, user_id").
		Where("flight_id = ?", flightID).
		Order("passenger_id ASC").
		Scan(&passengers).Error

	if err != nil {
		return nil, fmt.Errorf("error fetching passengers: %v", err)
	}

	return passengers, nil
}

// GetPassengerByID retrieves a specific passenger with flight details
func (r *PassengerRepository) GetPassengerByID(id uint) (*passenger.PassengerDetailResponse, error) {
	var result passenger.PassengerQueryResult

	err := r.db.Table("passenger p").
		Select(`
			p.passenger_id, p.name, p.passport_number, p.nationality, 
			p.flight_id, p.special_requests, p.user_id,
			f.flight_number, f.departure_time, f.arrival_time,
			dep.iata_code as from_airport_iata, dep.name as from_airport_name, 
			dep.city as from_airport_city, dep.country as from_airport_country,
			arr.iata_code as to_airport_iata, arr.name as to_airport_name, 
			arr.city as to_airport_city, arr.country as to_airport_country
		`).
		Joins("JOIN flight f ON f.flight_id = p.flight_id").
		Joins("JOIN route r ON r.route_id = f.route_id").
		Joins("JOIN airport dep ON dep.airport_id = r.from_airport").
		Joins("JOIN airport arr ON arr.airport_id = r.to_airport").
		Where("p.passenger_id = ?", id).
		First(&result).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching passenger: %v", err)
	}

	return result.MapToDetailResponse(), nil
}

// SearchPassengers searches for passengers based on name, passport number, or nationality
func (r *PassengerRepository) SearchPassengers(query string, page, pageSize int) ([]passenger.PassengerResponse, int64, error) {
	var passengers []passenger.PassengerResponse
	var total int64

	// Prepare search query
	searchQuery := "%" + query + "%"

	// Count total matching records
	if err := r.db.Table("passenger").
		Where("name ILIKE ? OR passport_number ILIKE ? OR nationality ILIKE ?", searchQuery, searchQuery, searchQuery).
		Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("error counting search results: %v", err)
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Get paginated records matching search criteria
	err := r.db.Table("passenger").
		Select("passenger_id, name, passport_number, nationality, flight_id, special_requests, user_id").
		Where("name ILIKE ? OR passport_number ILIKE ? OR nationality ILIKE ?", searchQuery, searchQuery, searchQuery).
		Order("passenger_id ASC").
		Limit(pageSize).
		Offset(offset).
		Scan(&passengers).Error

	if err != nil {
		return nil, 0, fmt.Errorf("error searching passengers: %v", err)
	}

	return passengers, total, nil
}
