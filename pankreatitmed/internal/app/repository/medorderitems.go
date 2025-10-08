package repository

import (
	"fmt"
	"pankreatitmed/internal/app/ds"
)

func (r *Repository) DeleteFromOrder(medorder, criterion uint) error {
	fmt.Println("sqfew")
	return r.db.Where("med_order_id = ? AND criterion_id = ?", medorder, criterion).Delete(&ds.MedOrderItem{}).Error
}

func (r *Repository) UpdateMedOrderItem(medorder, criterion uint, position *uint, val *float64) error {
	updates := make(map[string]any)
	if position != nil {
		updates["position"] = *position
	}
	if val != nil {
		updates["value_num"] = *val
	}

	if len(updates) == 0 {
		return nil
	}

	return r.db.Model(&ds.MedOrderItem{}).
		Where("med_order_id = ? AND criterion_id = ?", medorder, criterion).
		Updates(updates).Error
}
