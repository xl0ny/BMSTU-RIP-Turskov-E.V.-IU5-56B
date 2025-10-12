package repository

import (
	"context"
	"fmt"

	"pankreatitmed/internal/app/ds"

	"gorm.io/gorm/clause"
)

func (r *Repository) GetOrCreateDraftMedOrder(creatorID uint) (*ds.PankreatitOrder, error) {
	var o ds.PankreatitOrder
	if err := r.db.Where("creator_id = ? AND status = 'draft'", creatorID).First(&o).Error; err == nil {

		return &o, nil
	}
	o = ds.PankreatitOrder{Status: "draft", CreatorID: creatorID}
	return &o, r.db.Create(&o).Error
}

func (r *Repository) AddItem(orderID, criterionID uint) error {
	var lastOI ds.PankreatitOrderItem
	r.db.Last(&lastOI, "pankreatit_order_id = ?", orderID)
	item := ds.PankreatitOrderItem{PankreatitOrderID: orderID, CriterionID: criterionID, Position: lastOI.Position + 1}
	return r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&item).Error
}

func (r *Repository) GetMedOrderWithItems(MedOrderID uint) (ds.PankreatitOrder, []ds.PankreatitOrderItem, error) {
	var o ds.PankreatitOrder
	if err := r.db.First(&o, MedOrderID).Error; err != nil {
		return ds.PankreatitOrder{}, nil, err
	}
	var items []ds.PankreatitOrderItem
	if err := r.db.Where("pankreatit_order_id = ?", MedOrderID).Order("id").Find(&items).Error; err != nil {
		return ds.PankreatitOrder{}, nil, err
	}
	return o, items, nil
}

func (r *Repository) CountItems(orderID uint) (int64, error) {
	var cnt int64
	return cnt, r.db.Model(&ds.PankreatitOrderItem{}).Where("pankreatit_order_id = ?", orderID).Count(&cnt).Error
}

func (r *Repository) SoftDeleteOrderSQL(ctx context.Context, orderID uint) error {
	sql := `UPDATE pankreatitorders SET status='deleted' WHERE id=$1 AND status<>'deleted'`
	tx := r.db.Exec(sql, orderID)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("pankreatitorder %d not updated", orderID)
	}
	return nil
}

func (r *Repository) IsMedOrderDeleted(MedOrderID uint) (bool, error) {
	var o ds.PankreatitOrder
	err := r.db.First(&o, "id = ?", MedOrderID).Error
	if err == nil {
		return o.Status == "deleted", nil
	}
	return true, err
}
