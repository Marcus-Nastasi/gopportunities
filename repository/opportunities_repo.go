package repository

import (
	"github.com/Marcus-Nastasi/gopportunities/schemas"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r *Repository) GetOpportunities() ([]schemas.Opening, error) {
	var o []schemas.Opening
	if err := r.db.Find(&o).Error; err != nil {
		return nil, err
	}
	return o, nil
}


func (r *Repository) GetOpportunitie(id string) (schemas.Opening, error) {
	var o schemas.Opening
	if err := r.db.Find(&o, id).Error; err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}


func (r *Repository) CreateOpportunitie(o schemas.Opening) (schemas.Opening, error) {
	if err := r.db.Create(&o).Error; err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}

func (r *Repository) UpdateOpportunitie(id string, o schemas.Opening) (schemas.Opening, error) {
	uo, err := r.GetOpportunitie(id)
	if err != nil {
		return schemas.Opening{}, err
	}
	uo.Update(&o)
	if err := r.db.Save(&uo).Error; err != nil {
		return schemas.Opening{}, err
	}
	return uo, nil
}

func (r *Repository) DeleteOpportunitie(id string) (schemas.Opening, error) {
	o, err := r.GetOpportunitie(id)
	if err != nil {
		return schemas.Opening{}, err
	}
	if err = r.db.Delete(&o, id).Error; err != nil {
		return schemas.Opening{}, err
	}
	return o, nil
}
