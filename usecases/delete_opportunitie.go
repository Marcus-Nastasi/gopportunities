package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/repository"
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

type DeleteOpportunitieUsecase struct {
	repo *repository.Repository
}

func NewDeleteOpportunitieUsecase(repo *repository.Repository) *DeleteOpportunitieUsecase {
	return &DeleteOpportunitieUsecase{
		repo: repo,
	}
}

func (u *DeleteOpportunitieUsecase) DeleteOpportunitie(id string) (schemas.Opening, error) {
	// db := config.GetDb()
	o, err := u.repo.DeleteOpportunitie(id)
	if err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}
