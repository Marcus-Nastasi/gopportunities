package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/schemas"
	"gorm.io/gorm"
)

type GetOpportunitieUsecase struct {
	db *gorm.DB
}

func NewGetOpportunitieUsecase(db *gorm.DB) *GetOpportunitieUsecase {
	return &GetOpportunitieUsecase{
		db: db,
	}
}

func (u *GetOpportunitieUsecase) GetOpportunitie(id string) (schemas.Opening, error) {
	// db := config.GetDb()
	var o schemas.Opening
	if err := u.db.Find(&o, id).Error; err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}
