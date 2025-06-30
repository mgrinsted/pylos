package handlers

import "github.com/mgrinsted/pylos/internal/services"

type Handlers struct {
	EstateCustomer *EstateCustomerHandler
}

func SetupHandlers(services *services.Services) *Handlers {
	return &Handlers{
		EstateCustomer: NewEstateCustomerHandler(services.EstateCustomerService),
	}
}
