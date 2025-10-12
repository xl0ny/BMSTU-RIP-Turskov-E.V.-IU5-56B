package repository

import (
	"fmt"
	"pankreatitmed/internal/app/ds"
)

func (r *Repository) DeleteFromPankreatitOrder(pankreatitorder, criterion uint) error {
	fmt.Println("sqfew")
	return r.db.Where("pankreatit_order_id = ? AND criterion_id = ?", pankreatitorder, criterion).Delete(&ds.PankreatitOrderItem{}).Error
}

func (r *Repository) UpdatePankreatitOrderItem(pankreatitorder, criterion uint, position *uint, val *float64) error {
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

	return r.db.Model(&ds.PankreatitOrderItem{}).
		Where("pankreatit_order_id = ? AND criterion_id = ?", pankreatitorder, criterion).
		Updates(updates).Error
}
