package movie

import (
	"time"
)

// Movie ...
type Movie struct {
	ID        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title     string    `json:"title"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FilterSpec ...
type FilterSpec struct {
	Title    string
	YearFrom int
	YearTo   int
}
