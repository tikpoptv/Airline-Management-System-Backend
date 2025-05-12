package repository

import (
	"airline-management-system/internal/models/flight"
	"errors"
	"strings"
	"time"

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
	if err := r.db.
		Preload("Aircraft").
		Preload("Route").
		Preload("Route.FromAirport").
		Preload("Route.ToAirport").
		Find(&flights).Error; err != nil {
		return nil, err
	}
	return flights, nil
}

func (r *FlightRepository) CreateFlight(flight *flight.Flight) error {
	return r.db.Create(flight).Error
}

func (r *FlightRepository) GetFlightByID(id uint) (*flight.Flight, error) {
	var f flight.Flight
	if err := r.db.
		Preload("Aircraft").
		Preload("Route").
		Preload("Route.FromAirport").
		Preload("Route.ToAirport").
		First(&f, id).Error; err != nil {
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

func (r *FlightRepository) GetFlightsByAircraftID(aircraftID uint) ([]flight.Flight, error) {
	var flights []flight.Flight
	err := r.db.
		Preload("Aircraft").
		Preload("Route").
		Preload("Route.FromAirport").
		Preload("Route.ToAirport").
		Where("aircraft_id = ?", aircraftID).
		Order("departure_time DESC").
		Find(&flights).Error
	if err != nil {
		return nil, err
	}
	return flights, nil
}

func (r *FlightRepository) GetTodayFlights(status string) ([]flight.Flight, error) {
	var flights []flight.Flight

	// Get today's date in UTC
	now := time.Now().UTC()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.Add(24 * time.Hour)

	// Build query
	query := r.db.
		Preload("Aircraft").
		Preload("Route").
		Preload("Route.FromAirport").
		Preload("Route.ToAirport").
		Where("departure_time >= ? AND departure_time < ?", startOfDay.Format(time.RFC3339), endOfDay.Format(time.RFC3339))

	// Add status filter if not "all"
	if status != "all" {
		validStatuses := map[string]bool{
			"active":    true,
			"delayed":   true,
			"cancelled": true,
		}
		if !validStatuses[status] {
			return nil, errors.New("invalid status. must be one of: all, active, delayed, cancelled")
		}

		if status == "active" {
			query = query.Where("flight_status IN ?", []string{"Scheduled", "Boarding"})
		} else {
			query = query.Where("flight_status = ?", strings.Title(status))
		}
	}

	// Execute query
	err := query.Order("departure_time ASC").Find(&flights).Error
	if err != nil {
		return nil, err
	}

	return flights, nil
}
