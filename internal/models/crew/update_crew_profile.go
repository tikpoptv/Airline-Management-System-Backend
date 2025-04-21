package crew

type UpdateCrewProfileRequest struct {
	Name               string `json:"name" validate:"required"`
	PassportNumber     string `json:"passport_number" validate:"required"`
	Role               string `json:"role" validate:"required,oneof=Pilot Co-Pilot Attendant Technician"`
	LicenseExpiryDate  string `json:"license_expiry_date" validate:"required,datetime=2006-01-02"`
	PassportExpiryDate string `json:"passport_expiry_date" validate:"required,datetime=2006-01-02"`
}
