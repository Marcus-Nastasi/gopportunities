package usecases

import "github.com/Marcus-Nastasi/gopportunities/schemas"

func GetOpportunities() (o []schemas.Opening, err error) {
	if err = db.Find(&o).Error; err != nil {
		return nil, err
	}
	return
}
