package router

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/mgrinsted/pylos/internal/handlers"
)

func SetupRoutes(h *handlers.Handlers) http.Handler {
	r := chi.NewRouter()

	templates := loadTemplates()

	// Home route redirects to customers
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/customers/1021", http.StatusFound)
	})

	// Customer routes
	r.Get("/customers/{id}", h.EstateCustomer.CustomerDetailHandler(templates))
	r.Get("/customers", h.EstateCustomer.CustomersPageHandler(templates))

	return r
}

func loadTemplates() *template.Template {
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
		"iter": func(count int) []int {
			s := make([]int, count)
			for i := 0; i < count; i++ {
				s[i] = i
			}
			return s
		},
		"substr": func(s string, start, length int) string {
			if start >= len(s) {
				return ""
			}
			end := start + length
			if end > len(s) {
				end = len(s)
			}
			return s[start:end]
		},
		"upper": func(s string) string {
			return strings.ToUpper(s)
		},
		"default": func(defaultValue, value interface{}) interface{} {
			if value == nil || value == "" {
				return defaultValue
			}
			return value
		},
		"eq": func(a, b interface{}) bool {
			return a == b
		},
		"gt": func(a, b int) bool {
			return a > b
		},
		"lt": func(a, b int) bool {
			return a < b
		},
	}

	tmpl := template.New("").Funcs(funcMap)

	// Parse all templates
	tmpl = template.Must(tmpl.ParseGlob("templates/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("templates/**/*.html"))

	return tmpl
}
