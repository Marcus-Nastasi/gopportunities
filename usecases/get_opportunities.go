package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/repository"
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

type GetOpportunitiesUsecase struct {
	repo *repository.Repository
}

func NewGetOpportunitiesUsecase(repo *repository.Repository) *GetOpportunitiesUsecase {
	return &GetOpportunitiesUsecase{
		repo: repo,
	}
}

func (u *GetOpportunitiesUsecase) GetOpportunities() ([]schemas.Opening, error) {
	o, err := u.repo.GetOpportunities()
	if err != nil {
		return nil, err
	}
	return o, nil
}
