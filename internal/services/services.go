package services

import "github.com/jmoiron/sqlx"

type Services struct {
	TenantCustomerService *TenantCustomerService
	TenantNumbersService  *TenantNumbersService
}

func SetupServices(db *sqlx.DB) *Services {
	return &Services{
		TenantCustomerService: NewTenantCustomerService(db),
		TenantNumbersService:  NewTenantNumbersService(db),
	}
}
