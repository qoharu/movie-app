package movie

import (
	"context"
)

// UseCase ...
type UseCase interface {
	FindByID(context context.Context, id int) (Movie, error)
	FindByTitle(context context.Context, title string) ([]Movie, error)
	Add(context context.Context, movie Movie) (Movie, error)
	Update(context context.Context, editedMovie Movie) (Movie, error)
	Delete(context context.Context, id int) error
}

type useCase struct {
	repo Repository
}

// NewUseCase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewUseCase(repo Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

//FindByID ...
func (us *useCase) FindByID(context context.Context, id int) (res Movie, err error) {
	res, err = us.repo.GetByID(context, id)
	return res, err
}

//FindByTitle ...
func (us *useCase) FindByTitle(context context.Context, titleKeyword string) (res []Movie, err error) {
	res, err = us.repo.GetBySpec(context, FilterSpec{Title: titleKeyword})
	return res, err
}

//Add ...
func (us *useCase) Add(context context.Context, newMovie Movie) (res Movie, err error) {
	res, err = us.repo.Upsert(context, newMovie)
	return res, err
}

//Update ...
func (us *useCase) Update(context context.Context, editedMovie Movie) (res Movie, err error) {
	res, err = us.repo.Upsert(context, editedMovie)
	return res, err
}

//Remove ...
func (us *useCase) Delete(context context.Context, id int) (err error) {
	err = us.repo.DeleteByID(context, id)
	return err
}
