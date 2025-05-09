package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/config"
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

func GetOpportunitie(id string) (schemas.Opening, error) {
	db := config.GetDb()
	var o schemas.Opening
	if err := db.Find(&o, id).Error; err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}
