package repository

import (
	"airline-management-system/internal/database"
	"airline-management-system/internal/models/airlineowner"
	"context"
)

type AirlineOwnerRepository interface {
	GetAllAirlineOwners(ctx context.Context) ([]airlineowner.AirlineOwner, error)
}

type airlineOwnerRepository struct {
	db *database.DBService
}

func NewAirlineOwnerRepository(db *database.DBService) AirlineOwnerRepository {
	return &airlineOwnerRepository{db: db}
}

func (r *airlineOwnerRepository) GetAllAirlineOwners(ctx context.Context) ([]airlineowner.AirlineOwner, error) {
	var owners []airlineowner.AirlineOwner

	query := `
        SELECT id, name, country, alliance, description, created_at, updated_at
        FROM airline_owner
        ORDER BY name
    `

	result := r.db.QueryContext(ctx, query)
	if result.Error != nil {
		return nil, result.Error
	}

	if err := result.Scan(&owners).Error; err != nil {
		return nil, err
	}

	return owners, nil
}
