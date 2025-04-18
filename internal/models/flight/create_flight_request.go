package flight

type CreateFlightRequest struct {
	FlightNumber       string  `json:"flight_number" validate:"required"`
	AircraftID         uint    `json:"aircraft_id" validate:"required"`
	RouteID            uint    `json:"route_id" validate:"required"`
	DepartureTime      string  `json:"departure_time" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	ArrivalTime        string  `json:"arrival_time" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	FlightStatus       string  `json:"flight_status" validate:"required"`
	CancellationReason *string `json:"cancellation_reason"`
}
