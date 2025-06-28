package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/mgrinsted/pylos/internal/handlers"
)

func CustomersPageHandler(customerHandler *handlers.EstateCustomerHandler, templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customers, err := customerHandler.GetAll()
		if err != nil {
			http.Error(w, "Could not fetch customers", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"Title":     "Customers",
			"Customers": customers,
		}
log.Printf("Customers type: %T, count: %d", data["Customers"], len(customers))

		if err := templates.ExecuteTemplate(w, "layout", data); err != nil {
			log.Printf("Template execution error: %v", err)
			http.Error(w, "Template execution error", http.StatusInternalServerError)
		}
	}
}
