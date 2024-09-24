package main

import (
	"hacktiv8-go-final-project/internal/infrastructure/http"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load .env file")
	}

	r := http.SetupRoutes()

	port := os.Getenv("PORT")

	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}
}
