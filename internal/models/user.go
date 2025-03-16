package models

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"uniqueIndex"`
    Password  string
    Role      string
    CreatedAt time.Time
    UpdatedAt time.Time
}
