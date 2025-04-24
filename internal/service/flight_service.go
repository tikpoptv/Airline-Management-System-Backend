package service

import (
	"airline-management-system/internal/models/flight"
	"airline-management-system/internal/repository"
	"errors"
	"time"
)

type FlightService struct {
	repo *repository.FlightRepository
}

func NewFlightService(repo *repository.FlightRepository) *FlightService {
	return &FlightService{repo}
}

func (s *FlightService) GetAllFlights() ([]flight.Flight, error) {
	return s.repo.GetAllFlights()
}

func (s *FlightService) CreateFlight(req *flight.CreateFlightRequest) (*flight.Flight, error) {
	// ตรวจสอบ flight_status
	validStatuses := map[string]bool{
		"Scheduled": true, "Boarding": true, "Delayed": true, "Cancelled": true, "Completed": true,
	}
	if !validStatuses[req.FlightStatus] {
		return nil, errors.New("flight_status must be one of: Scheduled, Boarding, Delayed, Cancelled, Completed")
	}

	// แปลงเวลา
	dep, err := time.Parse(time.RFC3339, req.DepartureTime)
	if err != nil {
		return nil, errors.New("invalid departure_time format")
	}
	arr, err := time.Parse(time.RFC3339, req.ArrivalTime)
	if err != nil {
		return nil, errors.New("invalid arrival_time format")
	}
	if !arr.After(dep) {
		return nil, errors.New("arrival_time must be after departure_time")
	}

	newFlight := &flight.Flight{
		FlightNumber:       req.FlightNumber,
		AircraftID:         req.AircraftID,
		RouteID:            req.RouteID,
		DepartureTime:      req.DepartureTime,
		ArrivalTime:        req.ArrivalTime,
		FlightStatus:       req.FlightStatus,
		CancellationReason: req.CancellationReason,
	}

	if err := s.repo.CreateFlight(newFlight); err != nil {
		return nil, err
	}
	return newFlight, nil
}

func (s *FlightService) GetFlightByID(id uint) (*flight.Flight, error) {
	return s.repo.GetFlightByID(id)
}

func (s *FlightService) UpdateFlight(id uint, req *flight.UpdateFlightRequest) (*flight.Flight, error) {
	updates := make(map[string]interface{})

	if req.FlightStatus != nil {
		status := *req.FlightStatus
		validStatuses := map[string]bool{
			"Scheduled": true, "Boarding": true, "Delayed": true, "Cancelled": true, "Completed": true,
		}
		if !validStatuses[status] {
			return nil, errors.New("flight_status must be one of: Scheduled, Boarding, Delayed, Cancelled, Completed")
		}
		updates["flight_status"] = status
	}

	if req.CancellationReason != nil {
		updates["cancellation_reason"] = *req.CancellationReason
	}

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	return s.repo.UpdateFlight(id, updates)
}

func (s *FlightService) UpdateFlightDetails(id uint, req *flight.UpdateFlightDetailsRequest) (*flight.Flight, error) {
	updates := make(map[string]interface{})

	if req.FlightNumber != nil {
		updates["flight_number"] = *req.FlightNumber
	}
	if req.AircraftID != nil {
		updates["aircraft_id"] = *req.AircraftID
	}
	if req.RouteID != nil {
		updates["route_id"] = *req.RouteID
	}
	if req.DepartureTime != nil {
		if _, err := time.Parse(time.RFC3339, *req.DepartureTime); err != nil {
			return nil, errors.New("invalid departure_time format")
		}
		updates["departure_time"] = *req.DepartureTime
	}
	if req.ArrivalTime != nil {
		if _, err := time.Parse(time.RFC3339, *req.ArrivalTime); err != nil {
			return nil, errors.New("invalid arrival_time format")
		}
		updates["arrival_time"] = *req.ArrivalTime
	}

	// Validate: arrival > departure (ถ้ามีทั้งสอง)
	if req.DepartureTime != nil && req.ArrivalTime != nil {
		dep, _ := time.Parse(time.RFC3339, *req.DepartureTime)
		arr, _ := time.Parse(time.RFC3339, *req.ArrivalTime)
		if !arr.After(dep) {
			return nil, errors.New("arrival_time must be after departure_time")
		}
	}

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	return s.repo.UpdateFlightDetails(id, updates)
}

func (s *FlightService) DeleteFlight(id uint) error {
	return s.repo.DeleteFlight(id)
}

func (s *FlightService) GetFlightsByAircraftID(aircraftID uint) ([]flight.Flight, error) {
	return s.repo.GetFlightsByAircraftID(aircraftID)
}
