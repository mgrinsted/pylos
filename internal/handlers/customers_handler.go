package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/mgrinsted/pylos/internal/services"
)

type EstateCustomerHandler struct {
	service *services.EstateCustomerService
}

type Pagination struct {
	Page       int
	PageSize   int
	Total      int
	TotalPages int
}

func NewEstateCustomerHandler(service *services.EstateCustomerService) *EstateCustomerHandler {
	return &EstateCustomerHandler{
		service: service,
	}
}

// CustomersPageHandler returns an HTTP handler that renders the customer page
func (h *EstateCustomerHandler) CustomersPageHandler(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle pagination input (default page 1)
		pageStr := r.URL.Query().Get("page")
		page := 1
		if pageStr != "" {
			if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
				page = p
			}
		}

		pageSize := 10 // default page size
		limitStr := r.URL.Query().Get("limit")
		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
				pageSize = l
			}
		}

		offset := (page - 1) * pageSize

		customers, total, err := h.service.GetPaginated(pageSize, offset)
		if err != nil {
			http.Error(w, "Unable to load customers", http.StatusInternalServerError)
			log.Printf("Customer query failed: %v", err)
			return
		}

		pagination := Pagination{
			Page:       page,
			PageSize:   pageSize,
			Total:      total,
			TotalPages: (total + pageSize - 1) / pageSize,
		}

		data := map[string]interface{}{
			"Title":       "Customers - Admin Portal",
			"CurrentPage": "customers",
			"Customers":   customers,
			"Pagination":  pagination,
		}

		// Determine which template to execute
		var templateName string
		if r.Header.Get("HX-Request") == "true" {
			switch r.Header.Get("HX-Target") {
			case "customer-table-container":
				templateName = "partials/table_container"
			default:
				// For HTMX requests targeting other elements, return the full content
				templateName = "content"
			}
		} else {
			// For full page requests, return the complete page
			templateName = "customers"
		}

		if err := templates.ExecuteTemplate(w, templateName, data); err != nil {
			log.Printf("Template execution error: %v", err)
			return
		}
	}
}
