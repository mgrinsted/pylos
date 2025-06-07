package services

import "github.com/jmoiron/sqlx"

type Services struct {
	TenantCustomerService *TenantCustomerService
	TenantNumbersService  *TenantNumbersService
	EstateCustomerService *EstateCustomerService
	EstateNumbersService  *EstateNumbersService
}

func SetupServices(db *sqlx.DB) *Services {
	return &Services{
		TenantCustomerService: NewTenantCustomerService(db),
		TenantNumbersService:  NewTenantNumbersService(db),
		EstateCustomerService: NewEstateCustomerService(db),
		EstateNumbersService:  NewEstateNumbersService(db),
	}
}
