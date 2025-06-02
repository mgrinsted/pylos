package handlers

import (
	"github.com/mgrinsted/pylos/internal/services"
)

type TenantNumbersHandler struct {
	service *services.TenantNumbersService
}

func NewTenantNumbersHandler(service *services.TenantNumbersService) *TenantNumbersHandler {
	return &TenantNumbersHandler{
		service: service,
	}
}
