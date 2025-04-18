package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger" // ✅ เพิ่มตรงนี้
)

func InitDB() *gorm.DB {
	dsn := GetEnv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // ✅ เปิดโหมดแสดง query
	})
	if err != nil {
		log.Fatal("ไม่สามารถเชื่อมต่อฐานข้อมูล:", err)
	}
	return db
}
