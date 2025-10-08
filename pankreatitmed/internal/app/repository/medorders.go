package repository

import (
	"fmt"
	"pankreatitmed/internal/app/ds"
	"pankreatitmed/internal/app/dto/request"
	"time"
)

func (r *Repository) CountItems(orderID uint) (int64, error) {
	var cnt int64
	return cnt, r.db.Model(&ds.MedOrderItem{}).Where("med_order_id = ?", orderID).Count(&cnt).Error
}

func (r *Repository) IsMedOrderDeleted(MedOrderID uint) (bool, error) {
	var o ds.MedOrder
	err := r.db.First(&o, "id = ?", MedOrderID).Error
	if err == nil {
		return o.Status == "deleted", nil
	}
	return true, err
}

func (r *Repository) IsMedOrderDraft(MedOrderID uint) (bool, error) {
	var o ds.MedOrder
	err := r.db.First(&o, "id = ?", MedOrderID).Error
	if err == nil {
		return o.Status == "draft", nil
	}
	return true, err
}

func (r *Repository) IsMedOrderFormed(MedOrderID uint) (bool, error) {
	var o ds.MedOrder
	err := r.db.First(&o, "id = ?", MedOrderID).Error
	if err == nil {
		return o.Status == "formed", nil
	}
	return true, err
}

// ----------------------------------------------------------

func (r *Repository) GetOrCreateDraftMedOrder(creatorID uint) (*ds.MedOrder, error) {
	var o ds.MedOrder
	println("GetOrCreateDraftMedOrder")
	if err := r.db.Where("creator_id = ? AND status = 'draft'", creatorID).First(&o).Error; err == nil {
		return &o, nil
	}
	o = ds.MedOrder{Status: "draft", CreatorID: creatorID}
	return &o, r.db.Create(&o).Error
}

func (r *Repository) GetMedOrders(status string, start, end time.Time) ([]ds.MedOrder, error) {
	var orders []ds.MedOrder
	err := r.db.Where("(created_at >= ? AND created_at < ?) AND status = ?", start, end, status).Find(&orders).Error
	return orders, err
}

func (r *Repository) GetMedOrderWithItems(MedOrderID uint) (ds.MedOrder, []ds.MedOrderItem, error) {
	var o ds.MedOrder
	if err := r.db.First(&o, MedOrderID).Error; err != nil {
		return ds.MedOrder{}, nil, err
	}
	var items []ds.MedOrderItem
	if err := r.db.Where("med_order_id = ?", MedOrderID).Order("id").Find(&items).Error; err != nil {
		return ds.MedOrder{}, nil, err
	}
	return o, items, nil
}

func (r *Repository) UpdateMedOrder(id uint, order *request.UpdateMedOrder) error {
	return r.db.Model(&ds.MedOrder{}).Where("id = ?", id).Updates(order).Error
}

func (r *Repository) FormMedOrder(id uint) error {
	return r.db.Model(&ds.MedOrder{}).Where("id = ?", id).UpdateColumns(map[string]any{
		"status":    "formed",
		"formed_at": time.Now(),
	}).Error
}

func (r *Repository) EndOrCancelMedOrder(id, moderator uint, status string) error {
	return r.db.Model(&ds.MedOrder{}).Where("id = ?", id).UpdateColumns(map[string]any{
		"status":       status,
		"finished_at":  time.Now(),
		"moderator_id": moderator,
	}).Error
}

func (r *Repository) SoftDeleteOrderSQL(orderID uint) error {
	sql := `UPDATE medorders SET status='deleted', formed_at = NOW() WHERE id=$1 AND status = 'draft'`

	tx := r.db.Exec(sql, orderID)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("medorder %d not updated", orderID)
	}
	return nil
}

func (r *Repository) SetRansonAndRisk(orderID uint, score int, risk string) error {
	return r.db.Model(&ds.MedOrder{}).
		Where("id = ?", orderID).
		Updates(map[string]interface{}{
			"ranson_score":   score,
			"mortality_risk": risk,
		}).Error
}
