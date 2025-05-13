package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/repository"
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

type UpdateOpportunitieUsecase struct {
	repo *repository.Repository
}

func NewUpdateOpportunitieUsecase(repo *repository.Repository) *UpdateOpportunitieUsecase {
	return &UpdateOpportunitieUsecase{
		repo: repo,
	}
}

func (u *UpdateOpportunitieUsecase) UpdateOpportunitie(id string, o schemas.Opening) (schemas.Opening, error) {
	uo, err := u.repo.UpdateOpportunitie(id, o)
	if err != nil {
		return schemas.Opening{}, err
	}
	return uo, nil
}
