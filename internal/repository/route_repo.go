package repository

import (
	"airline-management-system/internal/models/route"

	"gorm.io/gorm"
)

type RouteRepository struct {
	db *gorm.DB
}

func NewRouteRepository(db *gorm.DB) *RouteRepository {
	return &RouteRepository{db}
}

func (r *RouteRepository) GetAllRoutes() ([]route.Route, error) {
	var routes []route.Route
	err := r.db.
		Preload("FromAirport").
		Preload("ToAirport").
		Find(&routes).Error
	if err != nil {
		return nil, err
	}
	return routes, nil
}

func (r *RouteRepository) CreateRoute(entity *route.RouteEntity) error {
	return r.db.Create(entity).Error
}

func (r *RouteRepository) UpdateRouteStatus(routeID uint, status string) error {
	return r.db.Model(&route.RouteEntity{}).Where("route_id = ?", routeID).Update("status", status).Error
}
