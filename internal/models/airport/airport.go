package airport

type Airport struct {
	ID        uint    `gorm:"column:airport_id;primaryKey" json:"airport_id"`
	IATACode  string  `gorm:"column:iata_code" json:"iata_code"`
	Name      string  `json:"name"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Status    string  `json:"status"`
}

func (Airport) TableName() string {
	return "airport"
}
