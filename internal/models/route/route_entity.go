package route

type RouteEntity struct {
	ID                uint    `gorm:"column:route_id;primaryKey"`
	FromAirportID     uint    `gorm:"column:from_airport"`
	ToAirportID       uint    `gorm:"column:to_airport"`
	Distance          float64 `gorm:"column:distance"`
	EstimatedDuration string  `gorm:"column:estimated_duration"`
}

func (RouteEntity) TableName() string {
	return "route"
}
