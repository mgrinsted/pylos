package services

import (
	"database/sql"

	"github.com/mgrinsted/pylos/internal/utils/reflectutil"

	"github.com/jmoiron/sqlx"
	models "github.com/mgrinsted/pylos/internal/models/estate"
)

type TenantCustomerService struct {
	db *sqlx.DB
}

type EstateCustomerService struct {
	db      *sqlx.DB
	db_name string
}

func NewTenantCustomerService(db *sqlx.DB) *TenantCustomerService {
	return &TenantCustomerService{
		db: db,
	}
}

func NewEstateCustomerService(db *sqlx.DB, db_name string) *EstateCustomerService {
	return &EstateCustomerService{
		db:      db,
		db_name: db_name,
	}
}

// GetAll retrieves all customer names from the database.
func (s *EstateCustomerService) GetAll() ([]models.CustomerSummaryDTO, error) {
	var customers []models.CustomerSummaryDTO

	columns := reflectutil.GetDBFields(models.CustomerSummaryDTO{})

	query := "SELECT " + columns + " FROM tranquility_estate.customer"

	err := s.db.Select(&customers, query)
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
