package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/config"
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

func CreateOpportunitie(o schemas.Opening) (schemas.Opening, error) {
	db := config.GetDb()
	if err := db.Create(&o).Error; err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}
