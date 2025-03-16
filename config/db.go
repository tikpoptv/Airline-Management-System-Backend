package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

func InitDB() *gorm.DB {
    dsn := GetEnv("DB_DSN")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("ไม่สามารถเชื่อมต่อฐานข้อมูล:", err)
    }
    return db
}
