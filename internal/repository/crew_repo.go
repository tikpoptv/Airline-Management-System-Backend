package repository

import (
	"airline-management-system/internal/models/crew"

	"gorm.io/gorm"
)

type CrewRepository struct {
	db *gorm.DB
}

func NewCrewRepository(db *gorm.DB) *CrewRepository {
	return &CrewRepository{db}
}

func (r *CrewRepository) GetAllCrew() ([]crew.Crew, error) {
	var crews []crew.Crew
	if err := r.db.Preload("User").Find(&crews).Error; err != nil {
		return nil, err
	}
	return crews, nil
}
