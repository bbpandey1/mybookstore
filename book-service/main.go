package main

import (
	"log"
	"net/http"

	"mybookstore/book-service/db"
	"mybookstore/book-service/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file loaded (may be running in Kubernetes)")
	}

	db.Init()

	r := routes.SetupRouter()

	log.Println("📚 Bookstore API running on port 8080")
	http.ListenAndServe(":8080", r)
}
