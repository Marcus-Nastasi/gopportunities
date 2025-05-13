package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/repository"
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

type GetOpportunitieUsecase struct {
	repo *repository.Repository
}

func NewGetOpportunitieUsecase(repo *repository.Repository) *GetOpportunitieUsecase {
	return &GetOpportunitieUsecase{
		repo: repo,
	}
}

func (u *GetOpportunitieUsecase) GetOpportunitie(id string) (schemas.Opening, error) {
	o, err := u.repo.GetOpportunitie(id)
	if err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}
