package usecase

import (
	"context"

	"github.com/qoharu/go-clean-ddd/domain"
	"github.com/qoharu/go-clean-ddd/movie/repository"
)

type movieUseCase struct {
	movieRepo repository.MovieRepository
}

// NewMovieUseCase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewMovieUseCase(movieRepo repository.MovieRepository) domain.MovieUseCase {
	return &movieUseCase{
		movieRepo: movieRepo,
	}
}

//FindById ...
func (us *movieUseCase) FindByID(context context.Context, id int) (res domain.Movie, err error) {
	res, err = us.movieRepo.GetByID(context, id)

	return res, nil
}

//FindByTitle ...
func (us *movieUseCase) FindByTitle(context context.Context, titleKeyword string) (res []domain.Movie, err error) {
	res, err = us.movieRepo.GetBySpec(context, repository.MovieFilterSpec{Title: titleKeyword})

	return res, nil
}

//Add ...
func (us *movieUseCase) Add(context context.Context, newMovie domain.Movie) (res domain.Movie, err error) {
	res, err = us.movieRepo.Upsert(context, newMovie)

	return res, nil
}

//Update ...
func (us *movieUseCase) Update(context context.Context, editedMovie domain.Movie) (res domain.Movie, err error) {
	res, err = us.movieRepo.Upsert(context, editedMovie)

	return res, nil
}

//Remove ...
func (us *movieUseCase) Delete(context context.Context, id int) (err error) {
	err = us.movieRepo.DeleteByID(context, id)

	return err
}
