package router

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mgrinsted/pylos/internal/handlers"
	"github.com/mgrinsted/pylos/internal/web"
)

func SetupRoutes(h *handlers.Handlers) http.Handler {
	r := chi.NewRouter()

	templates := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/sidebar.html",
		"templates/topbar.html",
		"templates/customers.html",
	))

	// add an index page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Welcome to the Pylos API!"))
	})

	// r.Route("/numbers", func(r chi.Router) {
	// 	r.Get("/", h.Number.GetAll)
	// 	r.Get("/{id}", h.Number.GetByID)
	// })

	// r.Route("/customers", func(r chi.Router) {
	// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 		customers, err := h.EstateCustomer.GetAll()
	// 		if err != nil {
	// 			http.Error(w, err.Error(), http.StatusInternalServerError)
	// 			return
	// 		}
	// 		// Example: encode as JSON
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(customers)
	// 	})
	// 	// r.Get("/{id}", h.TenantCustomer.GetByID)
	// })

	r.Get("/customers", web.CustomersPageHandler(h.EstateCustomer, templates))

	return r
}
