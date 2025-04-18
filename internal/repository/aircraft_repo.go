package repository

import (
	"airline-management-system/internal/models/aircraft"

	"gorm.io/gorm"
)

type AircraftRepository struct {
	db *gorm.DB
}

func NewAircraftRepository(db *gorm.DB) *AircraftRepository {
	return &AircraftRepository{db}
}

func (r *AircraftRepository) GetAllAircraft() ([]aircraft.Aircraft, error) {
	var aircrafts []aircraft.Aircraft
	if err := r.db.Find(&aircrafts).Error; err != nil {
		return nil, err
	}
	return aircrafts, nil
}

func (r *AircraftRepository) CreateAircraft(ac *aircraft.Aircraft) error {
	return r.db.Create(ac).Error
}

func (r *AircraftRepository) GetAircraftByID(id uint) (*aircraft.Aircraft, error) {
	var result aircraft.Aircraft
	if err := r.db.First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *AircraftRepository) UpdateAircraft(id uint, updated map[string]interface{}) (*aircraft.Aircraft, error) {
	var ac aircraft.Aircraft
	if err := r.db.First(&ac, id).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&ac).Updates(updated).Error; err != nil {
		return nil, err
	}

	return &ac, nil
}

func (r *AircraftRepository) DeleteAircraft(id uint) error {
	var ac aircraft.Aircraft
	if err := r.db.First(&ac, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&ac).Error
}
