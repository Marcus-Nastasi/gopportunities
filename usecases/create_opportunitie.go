package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

func CreateOpportunitie(o schemas.Opening) (schemas.Opening, error) {
	if err := db.Create(&o).Error; err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}
