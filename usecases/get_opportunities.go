package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/schemas"
	"gorm.io/gorm"
)

type GetOpportunitiesUsecase struct {
	db *gorm.DB
}

func NewGetOpportunitiesUsecase(db *gorm.DB) *GetOpportunitiesUsecase {
	return &GetOpportunitiesUsecase{
		db: db,
	}
}

func (u *GetOpportunitiesUsecase) GetOpportunities() ([]schemas.Opening, error) {
	// db := config.GetDb()
	var o []schemas.Opening
	if err := u.db.Find(&o).Error; err != nil {
		return nil, err
	}
	return o, nil
}
