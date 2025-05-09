package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/schemas"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func CreateOpportunitie(o schemas.Opening) (schemas.Opening, error) {
	err := db.Create(&o).Error
	if err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}
