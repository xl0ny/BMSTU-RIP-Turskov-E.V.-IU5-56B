package repository

import (
	"pankreatitmed/internal/app/ds"
)

func (r *Repository) DeleteFromOrder(medorder, criterion uint) error {
	return r.db.Where("med_order_id = ? AND criterion_id = ?", medorder, criterion).Delete(&ds.MedOrderItem{}).Error
}

func (r *Repository) UpdateMedOrderItem(medorder, criterion, position uint, val float64) error {
	return r.db.Model(&ds.MedOrderItem{}).UpdateColumns(map[string]any{
		"position": position,
		"value_num": val,
	}).Error
}