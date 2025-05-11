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
