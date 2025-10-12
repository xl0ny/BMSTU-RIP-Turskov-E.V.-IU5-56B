package repository

import (
	"fmt"
	"pankreatitmed/internal/app/ds"
	"pankreatitmed/internal/app/dto/request"
	"time"
)

func (r *Repository) CountItems(orderID uint) (int64, error) {
	var cnt int64
	return cnt, r.db.Model(&ds.PankreatitOrderItem{}).Where("med_order_id = ?", orderID).Count(&cnt).Error
}

func (r *Repository) IsPankreatitOrderDeleted(MedOrderID uint) (bool, error) {
	var o ds.PankreatitOrder
	err := r.db.First(&o, "id = ?", MedOrderID).Error
	if err == nil {
		return o.Status == "deleted", nil
	}
	return true, err
}

func (r *Repository) IsPankreatitOrderDraft(MedOrderID uint) (bool, error) {
	var o ds.PankreatitOrder
	err := r.db.First(&o, "id = ?", MedOrderID).Error
	if err == nil {
		return o.Status == "draft", nil
	}
	return true, err
}

func (r *Repository) IsPankreatitOrderFormed(MedOrderID uint) (bool, error) {
	var o ds.PankreatitOrder
	err := r.db.First(&o, "id = ?", MedOrderID).Error
	if err == nil {
		return o.Status == "formed", nil
	}
	return true, err
}

// ----------------------------------------------------------

func (r *Repository) GetOrCreateDraftPankreatitOrder(creatorID uint) (*ds.PankreatitOrder, error) {
	var o ds.PankreatitOrder
	println("GetOrCreateDraftPankreatitOrder")
	if err := r.db.Where("creator_id = ? AND status = 'draft'", creatorID).First(&o).Error; err == nil {
		return &o, nil
	}
	o = ds.PankreatitOrder{Status: "draft", CreatorID: creatorID}
	return &o, r.db.Create(&o).Error
}

func (r *Repository) GetPankreatitOrders(status string, start, end time.Time) ([]ds.PankreatitOrder, error) {
	var orders []ds.PankreatitOrder
	err := r.db.Where("(created_at >= ? AND created_at < ?) AND status = ?", start, end, status).Find(&orders).Error
	return orders, err
}

func (r *Repository) GetPankreatitOrderWithItems(MedOrderID uint) (ds.PankreatitOrder, []ds.PankreatitOrderItem, error) {
	var o ds.PankreatitOrder
	if err := r.db.First(&o, MedOrderID).Error; err != nil {
		return ds.PankreatitOrder{}, nil, err
	}
	var items []ds.PankreatitOrderItem
	if err := r.db.Where("med_order_id = ?", MedOrderID).Order("id").Find(&items).Error; err != nil {
		return ds.PankreatitOrder{}, nil, err
	}
	return o, items, nil
}

func (r *Repository) UpdatePankreatitOrder(id uint, order *request.UpdatePankreatitOrder) error {
	return r.db.Model(&ds.PankreatitOrder{}).Where("id = ?", id).Updates(order).Error
}

func (r *Repository) FormPankreatitOrder(id uint) error {
	return r.db.Model(&ds.PankreatitOrder{}).Where("id = ?", id).UpdateColumns(map[string]any{
		"status":    "formed",
		"formed_at": time.Now(),
	}).Error
}

func (r *Repository) EndOrCancelPankreatitOrder(id, moderator uint, status string) error {
	return r.db.Model(&ds.PankreatitOrder{}).Where("id = ?", id).UpdateColumns(map[string]any{
		"status":       status,
		"finished_at":  time.Now(),
		"moderator_id": moderator,
	}).Error
}

func (r *Repository) SoftDeleteOrderSQL(orderID uint) error {
	sql := `UPDATE pankreatitorders SET status='deleted', formed_at = NOW() WHERE id=$1 AND status = 'draft'`

	tx := r.db.Exec(sql, orderID)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("pankreatitorder %d not updated", orderID)
	}
	return nil
}

func (r *Repository) SetRansonAndRisk(orderID uint, score int, risk string) error {
	return r.db.Model(&ds.PankreatitOrder{}).
		Where("id = ?", orderID).
		Updates(map[string]interface{}{
			"ranson_score":   score,
			"mortality_risk": risk,
		}).Error
}
