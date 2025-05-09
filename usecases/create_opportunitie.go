package usecases

import (
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

func CreateOpportunitie(o schemas.Opening) (co schemas.Opening, err error) {
	if err = db.Create(&co).Error; err != nil {
		return schemas.Opening{}, err
	}
	return
}
