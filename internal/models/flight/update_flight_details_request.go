package flight

type UpdateFlightDetailsRequest struct {
	FlightNumber  *string `json:"flight_number"`
	AircraftID    *uint   `json:"aircraft_id"`
	RouteID       *uint   `json:"route_id"`
	DepartureTime *string `json:"departure_time"` // RFC3339
	ArrivalTime   *string `json:"arrival_time"`   // RFC3339
}
