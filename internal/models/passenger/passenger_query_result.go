package passenger

type PassengerQueryResult struct {
	PassengerID     uint   `gorm:"column:passenger_id"`
	Name            string `gorm:"column:name"`
	PassportNumber  string `gorm:"column:passport_number"`
	Nationality     string `gorm:"column:nationality"`
	FlightID        uint   `gorm:"column:flight_id"`
	SpecialRequests string `gorm:"column:special_requests"`
	UserID          *uint  `gorm:"column:user_id"`

	FlightNumber  string `gorm:"column:flight_number"`
	DepartureTime string `gorm:"column:departure_time"`
	ArrivalTime   string `gorm:"column:arrival_time"`

	FromAirportIATA    string `gorm:"column:from_airport_iata"`
	FromAirportName    string `gorm:"column:from_airport_name"`
	FromAirportCity    string `gorm:"column:from_airport_city"`
	FromAirportCountry string `gorm:"column:from_airport_country"`

	ToAirportIATA    string `gorm:"column:to_airport_iata"`
	ToAirportName    string `gorm:"column:to_airport_name"`
	ToAirportCity    string `gorm:"column:to_airport_city"`
	ToAirportCountry string `gorm:"column:to_airport_country"`
}

func (r *PassengerQueryResult) MapToDetailResponse() *PassengerDetailResponse {
	response := &PassengerDetailResponse{
		PassengerID:     r.PassengerID,
		Name:            r.Name,
		PassportNumber:  r.PassportNumber,
		Nationality:     r.Nationality,
		FlightID:        r.FlightID,
		SpecialRequests: r.SpecialRequests,
		UserID:          r.UserID,
	}

	response.FlightDetails.FlightNumber = r.FlightNumber
	response.FlightDetails.DepartureTime = r.DepartureTime
	response.FlightDetails.ArrivalTime = r.ArrivalTime

	response.FlightDetails.Route.FromAirport.IATACode = r.FromAirportIATA
	response.FlightDetails.Route.FromAirport.Name = r.FromAirportName
	response.FlightDetails.Route.FromAirport.City = r.FromAirportCity
	response.FlightDetails.Route.FromAirport.Country = r.FromAirportCountry

	response.FlightDetails.Route.ToAirport.IATACode = r.ToAirportIATA
	response.FlightDetails.Route.ToAirport.Name = r.ToAirportName
	response.FlightDetails.Route.ToAirport.City = r.ToAirportCity
	response.FlightDetails.Route.ToAirport.Country = r.ToAirportCountry

	return response
}
