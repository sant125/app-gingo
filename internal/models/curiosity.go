package models

// Curiosity represents an interesting tattoo fact or piece of knowledge.
type Curiosity struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"` // "history", "culture", "science", "art"
}
