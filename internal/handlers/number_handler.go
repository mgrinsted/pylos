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

type EstateNumbersHandler struct {
	service *services.EstateNumbersService
}

func NewEstateNumbersHandler(service *services.EstateNumbersService) *EstateNumbersHandler {
	return &EstateNumbersHandler{
		service: service,
	}
}
