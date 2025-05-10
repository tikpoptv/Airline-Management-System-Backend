package airport

type CreateAirportRequest struct {
	IATACode  string  `json:"iata_code" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	City      string  `json:"city" validate:"required"`
	Country   string  `json:"country" validate:"required"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone" validate:"required"`
	Status    string  `json:"status" validate:"required,oneof=active inactive"`
}

type UpdateAirportRequest struct {
	Name      string  `json:"name" validate:"required"`
	City      string  `json:"city" validate:"required"`
	Country   string  `json:"country" validate:"required"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone" validate:"required"`
	Status    string  `json:"status" validate:"required,oneof=active inactive"`
}
