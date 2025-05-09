package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/config"
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

func UpdateOpportunitie(id string, o schemas.Opening) (schemas.Opening, error) {
	db := config.GetDb()
	uo, err := GetOpportunitie(id)
	if err != nil {
		return schemas.Opening{}, err
	}
	if err := db.Model(&uo).Updates(o).Error; err != nil {
		return schemas.Opening{}, err
	}
	return uo, nil
}
