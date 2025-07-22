package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"mybookstore/book-service/db"
	"mybookstore/book-service/middleware" // 👈 import middleware
	"mybookstore/book-service/routes"
)

func main() {
	// Load environment variables from .env (locally only)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found; assuming environment variables are set externally")
	}

	// Initialize PostgreSQL connection
	db.Init()

	// Set up router with all routes
	router := routes.SetupRouter()

	// Apply CORS middleware
	handlerWithCORS := middleware.EnableCORS(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("📚 Book service is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handlerWithCORS))
}
