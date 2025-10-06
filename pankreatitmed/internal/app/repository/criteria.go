package repository

import (
	"errors"
	"pankreatitmed/internal/app/ds"
	"pankreatitmed/internal/app/dto/request"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *Repository) CreateCriterion(c *ds.Criterion) error {
	return r.db.Create(c).Error
}

func (r *Repository) UpdateCriterion(id uint,c *request.UpdateCriterion) error {
	return r.db.Model(&ds.Criterion{}).Where("id = ?", id).Updates(c).Error
}

func (r *Repository) DeleteCriterion(id uint) error {
	return r.db.Delete(&ds.Criterion{}, id).Error
}

func (r *Repository) AddItem(orderID, criterionID uint) error {
	var lastOI ds.MedOrderItem
	r.db.Last(&lastOI, "med_order_id = ?", orderID)
	item := ds.MedOrderItem{MedOrderID: orderID, CriterionID: criterionID, Position: lastOI.Position + 1}
	return r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&item).Error
}

//func (r *Repository) AddItem(orderID, criterionID uint) error {
//	var lastOI ds.MedOrderItem
//	r.db.Last(&lastOI, "med_order_id = ?", orderID)
//	item := ds.MedOrderItem{MedOrderID: orderID, CriterionID: criterionID, Position: lastOI.Position + 1}
//	return r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&item).Error
//}
