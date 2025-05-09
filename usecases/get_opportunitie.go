package usecases

import "github.com/Marcus-Nastasi/gopportunities/schemas"

func GetOpportunitie(id string) (o schemas.Opening, err error) {
	if err = db.Find(&o, id).Error; err != nil {
		return schemas.Opening{}, err
	}
	return
}
