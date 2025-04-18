package airport

type AirportPreload struct {
	ID       uint   `gorm:"column:airport_id;primaryKey" json:"airport_id"`
	IATACode string `gorm:"column:iata_code" json:"iata_code"`
	Name     string `json:"name"`
	City     string `json:"city"`
	Country  string `json:"country"`
}

func (AirportPreload) TableName() string {
	return "airport"
}
