package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/schemas"
	"gorm.io/gorm"
)

type DeleteOpportunitieUsecase struct {
	db *gorm.DB
	gou *GetOpportunitieUsecase
}

func NewDeleteOpportunitieUsecase(db *gorm.DB, gou *GetOpportunitieUsecase) *DeleteOpportunitieUsecase {
	return &DeleteOpportunitieUsecase{
		db: db,
		gou: gou,
	}
}

func (u *DeleteOpportunitieUsecase) DeleteOpportunitie(id string) (schemas.Opening, error) {
	// db := config.GetDb()
	o, err := u.gou.GetOpportunitie(id)
	if err != nil {
		return schemas.Opening{}, err
	}
	if err = u.db.Delete(&o, id).Error; err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}
