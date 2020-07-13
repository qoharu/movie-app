package repository

import (
	"context"

	"github.com/qoharu/go-clean-ddd/domain"
)

// MovieRepository Interface
type MovieRepository interface {
	GetByID(ctx context.Context, id int) (domain.Movie, error)
	GetBySpec(ctx context.Context, spec MovieFilterSpec) ([]domain.Movie, error)
	Upsert(ctx context.Context, ar domain.Movie) (domain.Movie, error)
	DeleteByID(ctx context.Context, id int) error
}

// MovieFilterSpec ...
type MovieFilterSpec struct {
	Title    string
	YearFrom int
	YearTo   int
}
