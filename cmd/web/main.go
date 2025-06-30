package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mgrinsted/pylos/internal/db"
	"github.com/mgrinsted/pylos/internal/handlers"
	"github.com/mgrinsted/pylos/internal/router"
	"github.com/mgrinsted/pylos/internal/services"
)

func main() {

	// Load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the database
	database := db.Connect()
	defer database.Close()

	// setup the services and handlers
	services := services.SetupServices(database)
	handlers := handlers.SetupHandlers(services)

	// setup the routes
	r := router.SetupRoutes(handlers)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
