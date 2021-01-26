package repository

import (
	"context"
	"github.com/qoharu/go-clean-ddd/movie"
)

// MovieFilterSpec ...
type MovieFilterSpec struct {
	Title    string
	YearFrom int
	YearTo   int
}

// MovieRepository Interface
type MovieRepository interface {
	GetByID(ctx context.Context, id int) (movie.Movie, error)
	GetBySpec(ctx context.Context, spec MovieFilterSpec) ([]movie.Movie, error)
	Upsert(ctx context.Context, ar movie.Movie) (movie.Movie, error)
	DeleteByID(ctx context.Context, id int) error
}
