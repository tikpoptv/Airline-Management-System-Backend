package repository

import (
	"airline-management-system/internal/models/airport"

	"gorm.io/gorm"
)

type AirportRepository struct {
	db *gorm.DB
}

func NewAirportRepository(db *gorm.DB) *AirportRepository {
	return &AirportRepository{db}
}

func (r *AirportRepository) GetAllAirports() ([]airport.Airport, error) {
	var airports []airport.Airport
	if err := r.db.Find(&airports).Error; err != nil {
		return nil, err
	}
	return airports, nil
}

func (r *AirportRepository) CreateAirport(a *airport.Airport) error {
	return r.db.Create(a).Error
}
