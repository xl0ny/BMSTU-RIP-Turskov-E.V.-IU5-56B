package repository

import (
	"errors"
	"pankreatitmed/internal/app/ds"

	"gorm.io/gorm"
)

func (r *Repository) GetCriteria(q string) ([]ds.Criterion, error) {
	var list []ds.Criterion
	db := r.db.Model(&ds.Criterion{}).Where("status = 'active'").Order("id ASC")
	if q != "" {
		db = db.Where("name ILIKE ?", "%"+q+"%")
	}
	return list, db.Find(&list).Error
}

func (r *Repository) GetCriterionByID(id uint) (*ds.Criterion, error) {
	var c ds.Criterion
	err := r.db.First(&c, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &c, err
}
