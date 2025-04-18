package user

type User struct {
	ID        uint   `gorm:"column:user_id;primaryKey" json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
}

func (User) TableName() string {
	return "users"
}
