package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/config"
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

func GetOpportunities() ([]schemas.Opening, error) {
	db := config.GetDb()
	var o []schemas.Opening
	if err := db.Find(&o).Error; err != nil {
		return nil, err
	}
	return o, nil
}
