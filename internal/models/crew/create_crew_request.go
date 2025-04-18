package crew

type CreateCrewRequest struct {
	Name               string  `json:"name" validate:"required"`
	PassportNumber     string  `json:"passport_number" validate:"required"`
	Role               string  `json:"role" validate:"required,oneof=Pilot Co-Pilot Attendant Technician"`
	LicenseExpiryDate  string  `json:"license_expiry_date" validate:"required"`
	PassportExpiryDate string  `json:"passport_expiry_date" validate:"required"`
	FlightHours        float64 `json:"flight_hours"`
	UserID             *uint   `json:"user_id"` // optional
}
