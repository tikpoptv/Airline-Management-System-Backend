package crew

type UpdateCrewRequest struct {
	Name               *string  `json:"name"`
	PassportNumber     *string  `json:"passport_number"`
	Role               *string  `json:"role" validate:"omitempty,oneof=Pilot Co-Pilot Attendant Technician"`
	LicenseExpiryDate  *string  `json:"license_expiry_date"`
	PassportExpiryDate *string  `json:"passport_expiry_date"`
	FlightHours        *float64 `json:"flight_hours"`
	UserID             *uint    `json:"user_id"`
	Status             *string  `json:"status" validate:"omitempty,oneof=active inactive on_leave suspended retired"`
}
