package aircraftmodel

import "time"

type AircraftModel struct {
	ModelID       int       `json:"model_id"`
	ModelName     string    `json:"model_name"`
	Manufacturer  string    `json:"manufacturer"`
	CapacityRange string    `json:"capacity_range"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
