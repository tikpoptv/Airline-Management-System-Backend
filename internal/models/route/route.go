package route

import "airline-management-system/internal/models/airport"

type Route struct {
	ID                uint                   `gorm:"column:route_id;primaryKey" json:"route_id"`
	FromAirportID     uint                   `gorm:"column:from_airport" json:"-"`
	ToAirportID       uint                   `gorm:"column:to_airport" json:"-"`
	FromAirport       airport.AirportPreload `gorm:"foreignKey:FromAirportID;references:ID" json:"from_airport"`
	ToAirport         airport.AirportPreload `gorm:"foreignKey:ToAirportID;references:ID" json:"to_airport"`
	Distance          float64                `json:"distance"`
	EstimatedDuration string                 `gorm:"type:interval" json:"estimated_duration"`
	Status            string                 `json:"status"`
}

func (Route) TableName() string {
	return "route"
}

type RouteBasicResponse struct {
	RouteID           uint    `json:"route_id"`
	FromAirportID     uint    `json:"from_airport_id"`
	ToAirportID       uint    `json:"to_airport_id"`
	Distance          float64 `json:"distance"`
	EstimatedDuration string  `json:"estimated_duration"`
	Status            string  `json:"status"`
}
