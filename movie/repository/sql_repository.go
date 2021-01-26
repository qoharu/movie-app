package repository

import (
	"context"
	"github.com/qoharu/go-clean-ddd/movie"
	"math"

	"github.com/jinzhu/gorm"
)

type movieSQLRepository struct {
	Conn *gorm.DB
}

// NewMovieSQLRepository will implement MovieRepository Interface
func NewMovieSQLRepository(Conn *gorm.DB) MovieRepository {
	return &movieSQLRepository{Conn}
}

func (m *movieSQLRepository) GetByID(ctx context.Context, id int) (res movie.Movie, err error) {
	result := m.Conn.First(&res, id)
	return res, result.Error
}

func (m *movieSQLRepository) GetBySpec(ctx context.Context, spec MovieFilterSpec) (res []movie.Movie, err error) {
	query := m.Conn

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

func (m *movieSQLRepository) Upsert(ctx context.Context, movie movie.Movie) (res movie.Movie, err error) {
	result := m.Conn.Save(&movie)
	return movie, result.Error
}

func (m *movieSQLRepository) DeleteByID(ctx context.Context, id int) (err error) {
	result := m.Conn.Delete(&movie.Movie{ID: id})
	return result.Error
}
