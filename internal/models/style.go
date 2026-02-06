package models

// Style represents a tattoo style.
type Style struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Origin      string `json:"origin"`
	Popularity  string `json:"popularity"` // "high", "medium", "low"
}
