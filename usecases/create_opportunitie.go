package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/repository"
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

type CreateOpportunitieUsecase struct {
	repo *repository.Repository
}

func NewCreateOpportunitieUsecase(repo *repository.Repository) *CreateOpportunitieUsecase {
	return &CreateOpportunitieUsecase{
		repo: repo,
	}
}

func (u *CreateOpportunitieUsecase) CreateOpportunitie(o schemas.Opening) (schemas.Opening, error) {
	created, err := u.repo.CreateOpportunitie(o) 
	if err != nil {
		return schemas.Opening{}, err
	}
	return created, nil
}
