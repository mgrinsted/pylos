package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mgrinsted/pylos/internal/handlers"
)

func SetupRoutes(h *handlers.Handlers) http.Handler {
	r := chi.NewRouter()

	// r.Route("/numbers", func(r chi.Router) {
	// 	r.Get("/", h.Number.GetAll)
	// 	r.Get("/{id}", h.Number.GetByID)
	// })

	r.Route("/customers", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			customers, err := h.TenantCustomer.GetAll()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Example: encode as JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customers)
		})
		// r.Get("/{id}", h.TenantCustomer.GetByID)
	})

	return r
}
