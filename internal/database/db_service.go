package database

import (
	"context"

	"gorm.io/gorm"
)

type DBService struct {
	db *gorm.DB
}

func NewDBService(db *gorm.DB) *DBService {
	return &DBService{db: db}
}

func (d *DBService) QueryContext(ctx context.Context, query string, args ...interface{}) *gorm.DB {
	return d.db.Raw(query, args...)
}
