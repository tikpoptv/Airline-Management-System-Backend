package repository

import (
	"airline-management-system/internal/models/flight"

	"gorm.io/gorm"
)

type FlightRepository struct {
	db *gorm.DB
}

func NewFlightRepository(db *gorm.DB) *FlightRepository {
	return &FlightRepository{db}
}

func (r *FlightRepository) GetAllFlights() ([]flight.Flight, error) {
	var flights []flight.Flight
	if err := r.db.Find(&flights).Error; err != nil {
		return nil, err
	}
	return flights, nil
}

func (r *FlightRepository) CreateFlight(flight *flight.Flight) error {
	return r.db.Create(flight).Error
}

func (r *FlightRepository) GetFlightByID(id uint) (*flight.Flight, error) {
	var f flight.Flight
	if err := r.db.First(&f, id).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *FlightRepository) UpdateFlight(id uint, updates map[string]interface{}) (*flight.Flight, error) {
	var f flight.Flight
	if err := r.db.First(&f, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(&f).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *FlightRepository) UpdateFlightDetails(id uint, updates map[string]interface{}) (*flight.Flight, error) {
	var f flight.Flight
	if err := r.db.First(&f, id).Error; err != nil {
		return nil, err
	}

	// อัปเดตข้อมูล
	if err := r.db.Model(&f).Updates(updates).Error; err != nil {
		return nil, err
	}

	// ดึงข้อมูลล่าสุดมาอีกรอบเพื่อให้ได้ค่าใหม่ทั้งหมด
	if err := r.db.First(&f, id).Error; err != nil {
		return nil, err
	}

	return &f, nil
}

func (r *FlightRepository) DeleteFlight(id uint) error {
	result := r.db.Delete(&flight.Flight{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
