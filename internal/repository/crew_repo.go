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

func (r *CrewRepository) GetAllCrew() ([]crew.GetCrew, error) {
	var crews []crew.GetCrew
	if err := r.db.Preload("User").Find(&crews).Error; err != nil {
		return nil, err
	}
	return crews, nil
}

func (r *CrewRepository) CreateCrew(c *crew.Crew) error {
	return r.db.Create(c).Error
}

func (r *CrewRepository) GetCrewByID(id uint) (*crew.GetCrew, error) {
	var c crew.GetCrew
	if err := r.db.Preload("User").First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}
