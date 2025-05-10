package crew

type Crew struct {
	ID                 uint    `gorm:"column:crew_id;primaryKey"`
	Name               string  `json:"name"`
	PassportNumber     string  `json:"passport_number"`
	Role               string  `json:"role"`
	LicenseExpiryDate  string  `json:"license_expiry_date"`
	PassportExpiryDate string  `json:"passport_expiry_date"`
	FlightHours        float64 `json:"flight_hours"`
	UserID             *uint   `json:"user_id"` // optional
	Status             string  `json:"status" gorm:"default:active"`
}

func (Crew) TableName() string {
	return "crew"
}
