package services

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	models "github.com/mgrinsted/pylos/internal/models/estate"
)

type TenantCustomerService struct {
	db *sqlx.DB
}

func NewTenantCustomerService(db *sqlx.DB) *TenantCustomerService {
	return &TenantCustomerService{
		db: db,
	}
}

// GetAll retrieves all customer names from the database.
func (s *TenantCustomerService) GetAll() ([]models.EstateCustomer, error) {
	var customers []models.EstateCustomer
	err := s.db.Select(&customers, "SELECT * FROM tranquility_estate.customer")
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// GetByID retrieves a customer by their ID.
func (s *TenantCustomerService) GetByID(id int) (string, error) {
	var name string
	err := s.db.QueryRow("SELECT name FROM tranquility_estate.customer WHERE id = ?", id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // No customer found with the given ID
		}
		return "", err
	}
	return name, nil
}
