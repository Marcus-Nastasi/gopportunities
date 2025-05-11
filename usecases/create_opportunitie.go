package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/schemas"
	"gorm.io/gorm"
)

type CreateOpportunitieUsecase struct {
	db *gorm.DB
}

func NewCreateOpportunitieUsecase(db *gorm.DB) *CreateOpportunitieUsecase {
	return &CreateOpportunitieUsecase{
		db: db,
	}
}

func (u *CreateOpportunitieUsecase) CreateOpportunitie(o schemas.Opening) (schemas.Opening, error) {
	// db := config.GetDb()
	if err := u.db.Create(&o).Error; err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}
