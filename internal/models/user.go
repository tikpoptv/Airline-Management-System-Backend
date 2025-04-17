package models

import "time"

type User struct {
	ID             uint       `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	Username       string     `gorm:"uniqueIndex" json:"username"`
	Email          string     `gorm:"uniqueIndex" json:"email"`
	HashedPassword string     `json:"-"`
	Role           string     `json:"role"`
	IsActive       bool       `gorm:"default:true" json:"is_active"`
	LastLogin      *time.Time `json:"last_login,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `gorm:"column:updated_at;<-:false" json:"updated_at"`
}
