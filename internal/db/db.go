package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect() *sqlx.DB {
	dsn := os.Getenv("DB_HOST")
	if dsn == "" {
		log.Fatal("Database URL not set")
	}
	if os.Getenv("DB_USERNAME") != "" && os.Getenv("DB_PASSWORD") != "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), dsn, os.Getenv("DB_NAME"))
	} else {
		dsn = fmt.Sprintf("%s/%s", dsn, os.Getenv("DB_NAME"))
	}

	if os.Getenv("DB_SSL") == "true" {
		dsn += "?parseTime=true&tls=true"
	} else {
		dsn += "?parseTime=true&tls=false"
	}

	if os.Getenv("DEBUG") == "true" {
		log.Printf("Connecting to database with DSN: %s", dsn)
	} else {
		log.Println("Connecting to database...")
	}

	// ✅ Use sqlx here
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	fmt.Println("Connected to the database ✅")

	return db
}
