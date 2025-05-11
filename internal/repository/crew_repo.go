package repository

import (
	"airline-management-system/internal/models/crew"
	"time"

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

func (r *CrewRepository) UpdateCrew(id uint, updates map[string]interface{}) error {
	return r.db.Model(&crew.Crew{}).Where("crew_id = ?", id).Updates(updates).Error
}

func (r *CrewRepository) DeleteCrew(id uint) (bool, error) {
	result := r.db.Delete(&crew.Crew{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (r *CrewRepository) GetCrewFlightHours(id uint) (*crew.CrewFlightHoursResponse, error) {
	var result crew.CrewFlightHoursResponse
	err := r.db.Model(&crew.Crew{}).
		Select("crew_id, name, flight_hours").
		Where("crew_id = ?", id).
		First(&result).Error

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *CrewRepository) GetCrewIDByUserID(userID uint) (uint, error) {
	var c crew.Crew
	err := r.db.Select("crew_id").Where("user_id = ?", userID).First(&c).Error
	if err != nil {
		return 0, err
	}
	return c.ID, nil
}

func (r *CrewRepository) UpdateCrewProfileByID(id uint, data map[string]interface{}) error {
	tx := r.db.Model(&crew.Crew{}).Where("crew_id = ?", id).Updates(data)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *CrewRepository) IsPassportNumberTakenExceptID(passport string, excludeID uint) (bool, error) {
	var count int64
	err := r.db.Model(&crew.Crew{}).
		Where("passport_number = ? AND crew_id != ?", passport, excludeID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *CrewRepository) GetAvailableCrewsForFlight(flightID uint, departureTime time.Time) ([]crew.AvailableCrewResponse, error) {
	var crews []crew.AvailableCrewResponse

	// Get crews that:
	// 1. Are active
	// 2. Have valid license and passport
	// 3. Not assigned to other flights at the same time
	err := r.db.Table("crew c").
		Select(`
			c.crew_id,
			c.name,
			c.role,
			c.flight_hours,
			c.status,
			CASE 
				WHEN c.license_expiry_date > ? THEN true 
				ELSE false 
			END as license_valid
		`, departureTime).
		Where("c.status = ?", "active").
		Where("c.license_expiry_date > ?", departureTime).
		Where("c.passport_expiry_date > ?", departureTime).
		// Exclude crews that are already assigned to other flights at the same time
		Where(`NOT EXISTS (
			SELECT 1 FROM flight_crew_assignment fca
			JOIN flight f ON f.flight_id = fca.flight_id
			WHERE fca.crew_id = c.crew_id
			AND f.flight_id != ?
			AND f.departure_time <= ?
			AND f.arrival_time >= ?
		)`, flightID, departureTime, departureTime).
		Find(&crews).Error

	if err != nil {
		return nil, err
	}

	return crews, nil
}
