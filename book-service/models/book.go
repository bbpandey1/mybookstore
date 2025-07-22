package models

// Book represents the structure of a book in the system
type Book struct {
	ID       int    `json:"id"`       // Auto-incremented ID
	Title    string `json:"title"`    // Book title
	Author   string `json:"author"`   // Author name
	Quantity int    `json:"quantity"` // Available stock
	Sold     int    `json:"sold"`     // Total number sold
}
