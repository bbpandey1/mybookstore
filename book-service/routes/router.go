package routes

import (
	"mybookstore/book-service/handlers"

	"github.com/gorilla/mux"
)

// SetupRouter initializes all API routes and returns the router
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Book endpoints
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books", handlers.AddBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/bestsellers/day", handlers.GetBestSellerByDay).Methods("GET")
	r.HandleFunc("/bestsellers/week", handlers.GetBestSellerByWeek).Methods("GET")
	r.HandleFunc("/bestsellers/year", handlers.GetBestSellerByYear).Methods("GET")
	return r
}
