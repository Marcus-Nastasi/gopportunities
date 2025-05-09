package usecases

import "github.com/Marcus-Nastasi/gopportunities/schemas"

func UpdateOpportunitie(id string, o schemas.Opening) (schemas.Opening, error) {
	if err := db.Update(id, &o).Error; err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}
