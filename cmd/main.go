// main.go
package main

import (
	"log"
	"net/http"
	"os"
	"payment-gateway/db"
	_ "payment-gateway/docs"
	"payment-gateway/internal/api"
	"payment-gateway/internal/services"
)

// @title           Payment Gateway API
// @version         1.0
// @description     A simple payment gateway API with deposit and withdrawal endpoints
// @host      localhost:8080
func main() {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURL := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"

	err := services.RetryOperation(func() error {
		return db.InitializeDB(dbURL)
	}, 5)

	if err != nil {
		log.Fatalf("Could not initialize DB after multiple attempts: %v", err)
	}

	router := api.SetupRouter()

	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
