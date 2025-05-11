package passenger

type Passenger struct {
	ID              uint   `gorm:"column:passenger_id;primaryKey"`
	Name            string `gorm:"column:name;not null"`
	PassportNumber  string `gorm:"column:passport_number;not null;unique"`
	Nationality     string `gorm:"column:nationality;not null"`
	FlightID        uint   `gorm:"column:flight_id;not null"`
	SpecialRequests string `gorm:"column:special_requests"`
	UserID          *uint  `gorm:"column:user_id"`
}

func (Passenger) TableName() string {
	return "passenger"
}
