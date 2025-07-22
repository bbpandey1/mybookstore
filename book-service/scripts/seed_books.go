package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var sampleBooks = []struct {
	Title  string
	Author string
}{
	{"The Go Programming Language", "Alan A. A. Donovan"},
	{"Clean Code", "Robert C. Martin"},
	{"The Pragmatic Programmer", "Andrew Hunt"},
	{"Design Patterns", "Erich Gamma"},
	{"Effective Go", "Google"},
	{"Go in Action", "William Kennedy"},
	{"Go Web Programming", "Sau Sheong Chang"},
	{"Domain-Driven Design", "Eric Evans"},
	{"Refactoring", "Martin Fowler"},
	{"Introduction to Algorithms", "Thomas H. Cormen"},
}

func main() {
	// Load .env only if not in container
	if !isRunningInDocker() {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️  No .env file found")
		}
	}

	// Get DB config from env
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		requireEnv("DB_USER"),
		requireEnv("DB_PASSWORD"),
		requireEnv("DB_HOST"),
		requireEnv("DB_PORT"),
		requireEnv("DB_NAME"),
		requireEnv("SSL_MODE"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Ensure books table exists
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		quantity INT NOT NULL,
		sold INT DEFAULT 0,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatalf("❌ Failed to create table: %v", err)
	}

	// Check if any books exist
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)
	if err != nil {
		log.Fatalf("❌ Failed to check book count: %v", err)
	}
	if count > 0 {
		log.Println("✅ Books already seeded. Skipping...")
		return
	}

	// Insert books
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		book := sampleBooks[i%len(sampleBooks)]
		quantity := rand.Intn(100) + 1
		sold := rand.Intn(quantity)

		_, err := db.Exec(
			"INSERT INTO books (title, author, quantity, sold, updated_at) VALUES ($1, $2, $3, $4, NOW())",
			book.Title, book.Author, quantity, sold,
		)
		if err != nil {
			log.Printf("⚠️ Failed to insert book %d: %v", i+1, err)
		}
	}

	log.Println("✅ 50 books seeded.")
}

func requireEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		log.Fatalf("❌ Required environment variable %s not set", key)
	}
	return val
}

func isRunningInDocker() bool {
	// Check Docker environment hint files
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	if cgroup, err := os.ReadFile("/proc/1/cgroup"); err == nil && string(cgroup) != "" {
		return true
	}
	return false
}
