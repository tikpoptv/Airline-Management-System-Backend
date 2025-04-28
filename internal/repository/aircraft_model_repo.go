package repository

import (
	"airline-management-system/config"
	"airline-management-system/internal/database"
	"airline-management-system/internal/models/aircraftmodel"
	"context"
)

type AircraftModelRepository interface {
	GetAllAircraftModels(ctx context.Context) ([]aircraftmodel.AircraftModel, error)
}

type aircraftModelRepository struct {
	db *database.DBService
}

func NewAircraftModelRepository(db *database.DBService) AircraftModelRepository {
	return &aircraftModelRepository{db: db}
}

func (r *aircraftModelRepository) GetAllAircraftModels(ctx context.Context) ([]aircraftmodel.AircraftModel, error) {
	var models []aircraftmodel.AircraftModel

	query := `
        SELECT model_id, model_name, manufacturer, capacity_range, description, created_at, updated_at
        FROM ` + config.TableAircraftModel

	result := r.db.QueryContext(ctx, query)
	if result.Error != nil {
		return nil, result.Error
	}

	if err := result.Scan(&models).Error; err != nil {
		return nil, err
	}

	return models, nil
}
