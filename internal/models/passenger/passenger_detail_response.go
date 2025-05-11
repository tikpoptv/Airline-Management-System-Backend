package passenger

type PassengerDetailResponse struct {
	PassengerID    uint   `json:"passenger_id"`
	Name           string `json:"name"`
	PassportNumber string `json:"passport_number"`
	Nationality    string `json:"nationality"`
	FlightID       uint   `json:"flight_id"`
	FlightDetails  struct {
		FlightNumber  string `json:"flight_number"`
		DepartureTime string `json:"departure_time"`
		ArrivalTime   string `json:"arrival_time"`
		Route         struct {
			FromAirport struct {
				IATACode string `json:"iata_code"`
				Name     string `json:"name"`
				City     string `json:"city"`
				Country  string `json:"country"`
			} `json:"from_airport"`
			ToAirport struct {
				IATACode string `json:"iata_code"`
				Name     string `json:"name"`
				City     string `json:"city"`
				Country  string `json:"country"`
			} `json:"to_airport"`
		} `json:"route"`
	} `json:"flight_details"`
	SpecialRequests string `json:"special_requests,omitempty"`
	UserID          *uint  `json:"user_id,omitempty"`
}
