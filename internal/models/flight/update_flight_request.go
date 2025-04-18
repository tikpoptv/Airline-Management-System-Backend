package flight

type UpdateFlightRequest struct {
	FlightStatus       *string `json:"flight_status"`       
	CancellationReason *string `json:"cancellation_reason"`
}
