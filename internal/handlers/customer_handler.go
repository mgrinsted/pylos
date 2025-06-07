package handlers

import (
	"log"

	models "github.com/mgrinsted/pylos/internal/models/estate"
	"github.com/mgrinsted/pylos/internal/services"
)

type TenantCustomerHandler struct {
	service *services.TenantCustomerService
}
type EstateCustomerHandler struct {
	service *services.EstateCustomerService
}

func NewTenantCustomerHandler(service *services.TenantCustomerService) *TenantCustomerHandler {
	return &TenantCustomerHandler{
		service: service,
	}
}
// NewEstateCustomerHandler creates a new EstateCustomerHandler with the given service.
func NewEstateCustomerHandler(service *services.EstateCustomerService) *EstateCustomerHandler {
	return &EstateCustomerHandler{
		service: service,
	}
}

// GetAll retrieves all customer names from the database.

func (h *EstateCustomerHandler) GetAll() ([]models.EstateCustomer, error) {
	rows, err := h.service.GetAll()
	if err != nil {
		log.Printf("Error querying customers: %v", err)
		return nil, err
	}
	return rows, nil
}
