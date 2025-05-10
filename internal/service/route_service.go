package service

import (
	"airline-management-system/internal/models/route"
	"airline-management-system/internal/repository"
	"errors"
)

type RouteService struct {
	repo *repository.RouteRepository
}

func NewRouteService(repo *repository.RouteRepository) *RouteService {
	return &RouteService{repo}
}

func (s *RouteService) GetAllRoutes() ([]route.Route, error) {
	return s.repo.GetAllRoutes()
}

func (s *RouteService) CreateRoute(req *route.CreateRouteRequest) (*route.RouteBasicResponse, error) {
	if req.FromAirportID == req.ToAirportID {
		return nil, errors.New("from_airport_id and to_airport_id cannot be the same")
	}

	newRoute := &route.RouteEntity{
		FromAirportID:     req.FromAirportID,
		ToAirportID:       req.ToAirportID,
		Distance:          req.Distance,
		EstimatedDuration: req.EstimatedDuration,
		Status:            req.Status,
	}

	if err := s.repo.CreateRoute(newRoute); err != nil {
		return nil, err
	}

	return &route.RouteBasicResponse{
		RouteID:           newRoute.ID,
		FromAirportID:     newRoute.FromAirportID,
		ToAirportID:       newRoute.ToAirportID,
		Distance:          newRoute.Distance,
		EstimatedDuration: newRoute.EstimatedDuration,
		Status:            newRoute.Status,
	}, nil
}

func (s *RouteService) UpdateRouteStatus(routeID uint, req *route.UpdateRouteStatusRequest) error {
	if req.Status != "active" && req.Status != "inactive" {
		return errors.New("status must be either 'active' or 'inactive'")
	}

	return s.repo.UpdateRouteStatus(routeID, req.Status)
}
