package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/schemas"
	"gorm.io/gorm"
)

type UpdateOpportunitieUsecase struct {
	db *gorm.DB
	gou *GetOpportunitieUsecase
}

func NewUpdateOpportunitieUsecase(db *gorm.DB, gou * GetOpportunitieUsecase) *UpdateOpportunitieUsecase {
	return &UpdateOpportunitieUsecase{
		db: db,
		gou: gou,
	}
}

func (u *UpdateOpportunitieUsecase) UpdateOpportunitie(id string, o schemas.Opening) (schemas.Opening, error) {
	// db := config.GetDb()
	uo, err := u.gou.GetOpportunitie(id)
	if err != nil {
		return schemas.Opening{}, err
	}
	uo.Role = o.Role
	uo.Company = o.Company
	uo.Location = o.Location
	uo.Remote = o.Remote
	uo.Link = o.Link
	uo.Salary = o.Salary
	if err := u.db.Save(&uo).Error; err != nil {
		return schemas.Opening{}, err
	}
	return uo, nil
}
