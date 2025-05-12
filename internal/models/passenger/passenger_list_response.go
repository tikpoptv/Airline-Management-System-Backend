package passenger

type PassengerResponse struct {
	PassengerID     uint   `json:"passenger_id"`
	Name            string `json:"name"`
	PassportNumber  string `json:"passport_number"`
	Nationality     string `json:"nationality"`
	FlightID        uint   `json:"flight_id"`
	SpecialRequests string `json:"special_requests,omitempty"`
	UserID          *uint  `json:"user_id,omitempty"`
}

type AllPassengersResponse struct {
	Passengers []PassengerResponse `json:"passengers"`
	Total      int64               `json:"total"`
	Page       int                 `json:"page"`
	PageSize   int                 `json:"page_size"`
	TotalPages int                 `json:"total_pages"`
}

type FlightPassengerListResponse struct {
	FlightID    uint                `json:"flight_id"`
	FlightCode  string              `json:"flight_code"`
	FlightRoute string              `json:"flight_route"`
	FlightDate  string              `json:"flight_date"`
	Passengers  []PassengerResponse `json:"passengers"`
	TotalCount  int                 `json:"total_count"`
}
