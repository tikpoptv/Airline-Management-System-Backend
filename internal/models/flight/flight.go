package flight

type Flight struct {
	ID                 uint    `gorm:"column:flight_id;primaryKey" json:"flight_id"`
	FlightNumber       string  `json:"flight_number"`
	AircraftID         uint    `json:"aircraft_id"`
	RouteID            uint    `json:"route_id"`
	DepartureTime      string  `json:"departure_time"`
	ArrivalTime        string  `json:"arrival_time"`
	FlightStatus       string  `json:"flight_status"`
	CancellationReason *string `json:"cancellation_reason"`
}

func (Flight) TableName() string {
	return "flight"
}
