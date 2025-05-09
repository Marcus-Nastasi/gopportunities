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
	uo.Role = o.Role
	uo.Company = o.Company
	uo.Location = o.Location
	uo.Remote = o.Remote
	uo.Link = o.Link
	uo.Salary = o.Salary
	if err := db.Save(&uo).Error; err != nil {
		return schemas.Opening{}, err
	}
	return uo, nil
}
