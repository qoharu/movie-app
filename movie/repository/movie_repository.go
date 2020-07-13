package repository

import (
	"context"
	"math"

	"github.com/jinzhu/gorm"

	"github.com/qoharu/go-clean-ddd/domain"
)

type movieRepository struct {
	Conn *gorm.DB
}

// NewMovieRepository will implement MovieRepository Interface
func NewMovieRepository(Conn *gorm.DB) MovieRepository {
	return &movieRepository{Conn}
}

func (m *movieRepository) GetByID(ctx context.Context, id int) (res domain.Movie, err error) {

	result := m.Conn.First(&res, id)

	return res, result.Error
}

func (m *movieRepository) GetBySpec(ctx context.Context, spec MovieFilterSpec) (res []domain.Movie, err error) {
	query := m.Conn

	if spec.Title != "" {
		query = query.Where("title ILIKE ?", "%"+spec.Title+"%")
	}

	if (spec.YearTo != 0 || spec.YearFrom != 0)  {
		if (spec.YearTo == 0) {
			spec.YearTo = math.MaxInt16
		}
		query = query.Or("year >= ? AND year <= ?", spec.YearFrom, spec.YearTo)
	}

	result := query.Find(&res)

	return res, result.Error
}

func (m *movieRepository) Upsert(ctx context.Context, movie domain.Movie) (res domain.Movie, err error) {

	result := m.Conn.Save(&movie)

	return movie, result.Error
}

func (m *movieRepository) DeleteByID(ctx context.Context, id int) (err error) {

	result := m.Conn.Delete(&domain.Movie{ID: id})

	return result.Error
}
