package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/config"
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

func DeleteOpportunitie(id string) (schemas.Opening, error) {
	db := config.GetDb()
	o, err := GetOpportunitie(id)
	if err != nil {
		return schemas.Opening{}, err
	}
	if err = db.Delete(&o, id).Error; err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}
