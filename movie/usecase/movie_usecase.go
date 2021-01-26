package usecase

import (
	"context"
	"github.com/qoharu/go-clean-ddd/movie"

	"github.com/qoharu/go-clean-ddd/movie/repository"
)

type movieUseCase struct {
	movieRepo repository.MovieRepository
}

// NewMovieUseCase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewMovieUseCase(movieRepo repository.MovieRepository) movie.UseCase {
	return &movieUseCase{
		movieRepo: movieRepo,
	}
}

//FindById ...
func (us *movieUseCase) FindByID(context context.Context, id int) (res movie.Movie, err error) {
	res, err = us.movieRepo.GetByID(context, id)
	return res, err
}

//FindByTitle ...
func (us *movieUseCase) FindByTitle(context context.Context, titleKeyword string) (res []movie.Movie, err error) {
	res, err = us.movieRepo.GetBySpec(context, repository.MovieFilterSpec{Title: titleKeyword})
	return res, err
}

//Add ...
func (us *movieUseCase) Add(context context.Context, newMovie movie.Movie) (res movie.Movie, err error) {
	res, err = us.movieRepo.Upsert(context, newMovie)
	return res, err
}

//Update ...
func (us *movieUseCase) Update(context context.Context, editedMovie movie.Movie) (res movie.Movie, err error) {
	res, err = us.movieRepo.Upsert(context, editedMovie)
	return res, err
}

//Remove ...
func (us *movieUseCase) Delete(context context.Context, id int) (err error) {
	err = us.movieRepo.DeleteByID(context, id)
	return err
}
