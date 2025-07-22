package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"mybookstore/book-service/db"
	"mybookstore/book-service/models"

	"github.com/gorilla/mux"
)

// GetBooks returns all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	search := r.URL.Query().Get("search")

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "id"
	}
	order := r.URL.Query().Get("order")
	if order != "desc" && order != "asc" {
		order = "asc" // ✅ default to ascending
	}

	condition := "1=1"
	countArgs := []interface{}{}
	selectArgs := []interface{}{limit, offset}

	if search != "" {
		condition = "(title ILIKE $1 OR author ILIKE $1)"
		countArgs = append(countArgs, "%"+search+"%")
		selectArgs = append(selectArgs, "%"+search+"%") // for $3
	}

	// Get total count
	var total int
	countQuery := "SELECT COUNT(*) FROM books WHERE " + condition
	err := db.Conn.QueryRow(countQuery, countArgs...).Scan(&total)
	if err != nil {
		http.Error(w, "Error getting count", http.StatusInternalServerError)
		return
	}

	// Get paginated books
	var rows *sql.Rows
	if search != "" {
		rows, err = db.Conn.Query(
			fmt.Sprintf("SELECT id, title, author, quantity, sold FROM books WHERE %s ORDER BY %s %s LIMIT $1 OFFSET $2", condition, sort, order),
			selectArgs...,
		)
	} else {
		rows, err = db.Conn.Query(
			fmt.Sprintf("SELECT id, title, author, quantity, sold FROM books WHERE %s ORDER BY %s %s LIMIT $1 OFFSET $2", condition, sort, order),
			selectArgs...,
		)
	}
	if err != nil {
		http.Error(w, "Error fetching books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		rows.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity, &b.Sold)
		books = append(books, b)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"books": books,
		"total": total,
	})
}

// AddBook inserts a new book
func AddBook(w http.ResponseWriter, r *http.Request) {
	var b models.Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Conn.Exec("INSERT INTO books (title, author, quantity) VALUES ($1, $2, $3)", b.Title, b.Author, b.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateBook updates book info (title, author, quantity, sold)
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var b models.Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Conn.Exec(
		"UPDATE books SET title=$1, author=$2, quantity=$3, sold=$4 WHERE id=$5",
		b.Title, b.Author, b.Quantity, b.Sold, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteBook removes a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := db.Conn.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetBestSellerToday returns the book(s) with highest sales today
func GetBestSellerByDay(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT id, title, author, quantity, sold
		FROM books
		WHERE DATE(updated_at) = CURRENT_DATE
		ORDER BY sold DESC
		LIMIT 1;
	`

	row := db.Conn.QueryRow(query)

	var b models.Book
	err := row.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity, &b.Sold)
	if err != nil {
		// No data found or some error
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]models.Book{})
		return
	}

	json.NewEncoder(w).Encode([]models.Book{b})
}

// GetBestSellerByWeek returns the book(s) with highest sales in the last 7 days
func GetBestSellerByWeek(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT id, title, author, quantity, sold
		FROM books
		WHERE updated_at >= CURRENT_DATE - INTERVAL '7 days'
		ORDER BY sold DESC
		LIMIT 1;
	`

	row := db.Conn.QueryRow(query)

	var b models.Book
	err := row.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity, &b.Sold)
	if err != nil {
		// No results
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]models.Book{})
		return
	}

	json.NewEncoder(w).Encode([]models.Book{b})
}

func GetBestSellerByYear(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT id, title, author, quantity, sold
		FROM books
		WHERE updated_at >= CURRENT_DATE - INTERVAL '1 year'
		ORDER BY sold DESC
		LIMIT 1;
	`

	row := db.Conn.QueryRow(query)

	var b models.Book
	err := row.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity, &b.Sold)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]models.Book{})
		return
	}

	json.NewEncoder(w).Encode([]models.Book{b})
}
