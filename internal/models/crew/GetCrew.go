package crew

import "airline-management-system/internal/models/user"

type GetCrew struct {
	ID                 uint    `gorm:"column:crew_id;primaryKey" json:"crew_id"`
	Name               string  `json:"name"`
	PassportNumber     string  `json:"passport_number"`
	Role               string  `json:"role"`
	LicenseExpiryDate  string  `json:"license_expiry_date"`
	PassportExpiryDate string  `json:"passport_expiry_date"`
	FlightHours        float64 `json:"flight_hours"`

	UserID *uint     `gorm:"column:user_id" json:"-"`
	User   user.User `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

func (GetCrew) TableName() string {
	return "crew"
}
