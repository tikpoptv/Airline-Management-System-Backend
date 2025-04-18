package route

type CreateRouteRequest struct {
	FromAirportID     uint    `json:"from_airport_id" validate:"required"`
	ToAirportID       uint    `json:"to_airport_id" validate:"required,nefield=FromAirportID"`
	Distance          float64 `json:"distance" validate:"required,gt=0"`
	EstimatedDuration string  `json:"estimated_duration" validate:"required"`
}
