package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/mgrinsted/pylos/internal/repository"
)

type Services struct {
	EstateCustomerService *EstateCustomerService
}

func SetupServices(db *sqlx.DB) *Services {
	estateCustomerRepo := repository.NewEstateCustomerRepository(db)
	
	return &Services{
		EstateCustomerService: NewEstateCustomerService(estateCustomerRepo),
	}
}
