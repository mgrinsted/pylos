package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mgrinsted/pylos/internal/models"
	models_estate "github.com/mgrinsted/pylos/internal/models/estate"
)

type EstateCustomerRepository interface {
	GetAllSummary(ctx context.Context, limit, offset int) ([]models_estate.CustomerSummaryDTO, error)
	CountAll(ctx context.Context) (int, error)
}

type estateCustomerRepository struct {
	db *sqlx.DB
}

func NewEstateCustomerRepository(db *sqlx.DB) EstateCustomerRepository {
	return &estateCustomerRepository{db: db}
}

func (r *estateCustomerRepository) GetAllSummary(ctx context.Context, limit, offset int) ([]models_estate.CustomerSummaryDTO, error) {
	var customers []models_estate.CustomerSummaryDTO

	query := `
		SELECT
			c.id,
			c.name,
			c.affiliate_id,
			a.name AS affiliate_name,
			c.namespace,
			c.status,
			c.storage
		FROM tranquility_estate.customer c
		LEFT JOIN tranquility_estate.affiliates a ON c.affiliate_id = a.id
		LIMIT ? OFFSET ?
	`

	err := r.db.SelectContext(ctx, &customers, query, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := range customers {
		customers[i].StatusLabel, customers[i].StatusColor, customers[i].StatusBadge = models.MapStatus(customers[i].Status)
	}

	return customers, nil
}

func (r *estateCustomerRepository) CountAll(ctx context.Context) (int, error) {
	var count int
	err := r.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM tranquility_estate.customer")
	return count, err
}
