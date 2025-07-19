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
	GetByID(ctx context.Context, id int) (*models_estate.CustomerDetailDTO, error)
	// Add any other methods needed for customer operations
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
func (r *estateCustomerRepository) GetByID(ctx context.Context, id int) (*models_estate.CustomerDetailDTO, error) {
	var customer models_estate.CustomerDetailDTO

	/*
		ID				int
		Name           	*string
		AffiliateID    	*string
		Status         	*int
		Created        	*time.Time
		Storage        	*string
		SFTP           	*int
		SupportEmail   	*string
		SupportNumber  	*string
	*/

	query := `
		SELECT
			c.id,
			c.name,
			c.affiliate_id,
			a.name AS affiliate_name,
			c.status,
			c.storage,
			c.created,
			c.sftp,
			c.supportEmail,
			c.supportNumber
		FROM tranquility_estate.customer c
		LEFT JOIN tranquility_estate.affiliates a ON c.affiliate_id = a.id
		WHERE c.id = ?
	`

	err := r.db.GetContext(ctx, &customer, query, id)
	if err != nil {
		return nil, err
	}

	customer.StatusLabel, customer.StatusColor, customer.StatusBadge = models.MapStatus(customer.Status)

	return &customer, nil
}
