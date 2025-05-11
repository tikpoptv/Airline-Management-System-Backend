package service

import (
	"airline-management-system/internal/models/crew"
	"airline-management-system/internal/repository"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type CrewService struct {
	repo       *repository.CrewRepository
	flightRepo *repository.FlightRepository
}

func NewCrewService(repo *repository.CrewRepository, flightRepo *repository.FlightRepository) *CrewService {
	return &CrewService{
		repo:       repo,
		flightRepo: flightRepo,
	}
}

func (s *CrewService) GetAllCrew() ([]crew.GetCrew, error) {
	return s.repo.GetAllCrew()
}

func (s *CrewService) CreateCrew(req *crew.CreateCrewRequest) (*crew.Crew, error) {
	newCrew := &crew.Crew{
		Name:               req.Name,
		PassportNumber:     req.PassportNumber,
		Role:               req.Role,
		LicenseExpiryDate:  req.LicenseExpiryDate,
		PassportExpiryDate: req.PassportExpiryDate,
		FlightHours:        req.FlightHours,
		UserID:             req.UserID,
		Status:             req.Status,
	}

	if err := s.repo.CreateCrew(newCrew); err != nil {
		return nil, err
	}

	return newCrew, nil
}

func (s *CrewService) GetCrewByID(id uint) (*crew.GetCrew, error) {
	return s.repo.GetCrewByID(id)
}

func (s *CrewService) UpdateCrew(id uint, req *crew.UpdateCrewRequest) error {
	updates := make(map[string]interface{})

	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.PassportNumber != nil {
		updates["passport_number"] = *req.PassportNumber
	}
	if req.Role != nil {
		role := strings.TrimSpace(*req.Role)
		if role != "Pilot" && role != "Co-Pilot" && role != "Attendant" && role != "Technician" {
			return errors.New("invalid role")
		}
		updates["role"] = role
	}
	if req.LicenseExpiryDate != nil {
		updates["license_expiry_date"] = *req.LicenseExpiryDate
	}
	if req.PassportExpiryDate != nil {
		updates["passport_expiry_date"] = *req.PassportExpiryDate
	}
	if req.FlightHours != nil {
		updates["flight_hours"] = *req.FlightHours
	}
	if req.UserID != nil {
		updates["user_id"] = *req.UserID
	}
	if req.Status != nil {
		status := strings.TrimSpace(*req.Status)
		if status != "active" && status != "inactive" && status != "on_leave" && status != "suspended" && status != "retired" {
			return errors.New("invalid status")
		}
		updates["status"] = status
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	return s.repo.UpdateCrew(id, updates)
}

func (s *CrewService) DeleteCrew(id uint) (bool, error) {
	return s.repo.DeleteCrew(id)
}

func (s *CrewService) GetCrewFlightHours(id uint) (*crew.CrewFlightHoursResponse, error) {
	return s.repo.GetCrewFlightHours(id)
}

func (s *CrewService) UpdateCrewProfileFromUser(userID uint, req *crew.UpdateCrewProfileRequest) error {
	crewID, err := s.repo.GetCrewIDByUserID(userID)
	if err != nil {
		return errors.New("crew not found for this user")
	}

	// Validate date
	licenseDate, err := time.Parse("2006-01-02", req.LicenseExpiryDate)
	if err != nil || licenseDate.Before(time.Now()) {
		return errors.New("invalid license expiry date")
	}
	passportDate, err := time.Parse("2006-01-02", req.PassportExpiryDate)
	if err != nil || passportDate.Before(time.Now()) {
		return errors.New("invalid passport expiry date")
	}

	// Prepare update map
	update := map[string]interface{}{
		"name":                 req.Name,
		"passport_number":      req.PassportNumber,
		"role":                 req.Role,
		"license_expiry_date":  licenseDate,
		"passport_expiry_date": passportDate,
	}

	return s.repo.UpdateCrewProfileByID(crewID, update)
}

func (s *CrewService) GetAvailableCrewsForFlight(flightID uint) ([]crew.AvailableCrewResponse, error) {
	if flightID == 0 {
		return nil, errors.New("invalid flight ID")
	}

	// Get flight departure time
	flight, err := s.flightRepo.GetFlightByID(flightID)
	if err != nil {
		return nil, fmt.Errorf("error getting flight: %v", err)
	}
	if flight == nil {
		return nil, errors.New("flight not found")
	}

	// Parse departure time
	departureTime, err := time.Parse(time.RFC3339, flight.DepartureTime)
	if err != nil {
		return nil, fmt.Errorf("invalid departure time format: %v", err)
	}

	// Convert both times to UTC for comparison
	now := time.Now().UTC()
	departureTimeUTC := departureTime.UTC()

	// For debugging
	fmt.Printf("Now: %v, Departure: %v\n", now, departureTimeUTC)

	crews, err := s.repo.GetAvailableCrewsForFlight(flightID, departureTime)
	if err != nil {
		return nil, fmt.Errorf("error getting available crews: %v", err)
	}

	return crews, nil
}

func (s *CrewService) GetCrewProfileByUserID(userID uint) (*crew.GetCrew, error) {
	// หา crew_id จาก user_id
	crewID, err := s.repo.GetCrewIDByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("crew not found")
		}
		return nil, fmt.Errorf("error getting crew ID: %v", err)
	}

	// ดึงข้อมูล crew พร้อม user information
	crewProfile, err := s.repo.GetCrewByID(crewID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("crew not found")
		}
		return nil, fmt.Errorf("error getting crew profile: %v", err)
	}

	return crewProfile, nil
}
