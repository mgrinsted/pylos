package handlers

import "github.com/mgrinsted/pylos/internal/services"

// Handlers struct holds all the handlers for the application
type Handlers struct {
	TenantCustomer *TenantCustomerHandler
	TenantNumbers  *TenantNumbersHandler
}

func SetupHandlers(services *services.Services) *Handlers {
	return &Handlers{
		TenantCustomer: NewTenantCustomerHandler(services.TenantCustomerService),
		TenantNumbers:  NewTenantNumbersHandler(services.TenantNumbersService),
	}
}


