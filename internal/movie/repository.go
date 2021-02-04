package movie

import (
	"context"
	"math"

	"github.com/jinzhu/gorm"
)

// Repository Interface
type Repository interface {
	GetByID(ctx context.Context, id int) (Movie, error)
	GetBySpec(ctx context.Context, spec FilterSpec) ([]Movie, error)
	Upsert(ctx context.Context, ar Movie) (Movie, error)
	DeleteByID(ctx context.Context, id int) error
}

// NewRepository will implement MovieRepository Interface
func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) GetByID(ctx context.Context, id int) (res Movie, err error) {
	result := r.db.First(&res, id)
	return res, result.Error
}

func (r *repository) GetBySpec(ctx context.Context, spec FilterSpec) (res []Movie, err error) {
	query := r.db

	if spec.Title != "" {
		query = query.Where("title ILIKE ?", "%"+spec.Title+"%")
	}

	if spec.YearTo != 0 || spec.YearFrom != 0 {
		if spec.YearTo == 0 {
			spec.YearTo = math.MaxInt16
		}
		query = query.Or("year >= ? AND year <= ?", spec.YearFrom, spec.YearTo)
	}

	result := query.Find(&res)

	return res, result.Error
}

func (r *repository) Upsert(ctx context.Context, movie Movie) (res Movie, err error) {
	result := r.db.Save(&movie)
	return movie, result.Error
}

func (r *repository) DeleteByID(ctx context.Context, id int) (err error) {
	result := r.db.Delete(&Movie{ID: id})
	return result.Error
}
