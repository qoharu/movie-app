package movie

import (
	"context"
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

// UseCase ...
type UseCase interface {
	FindByID(context context.Context, id int) (Movie, error)
	FindByTitle(context context.Context, title string) ([]Movie, error)
	Add(context context.Context, movie Movie) (Movie, error)
	Update(context context.Context, editedMovie Movie) (Movie, error)
	Delete(context context.Context, id int) error
}
