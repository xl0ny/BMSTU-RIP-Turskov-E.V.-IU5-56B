package ds

func (MedOrderItem) TableName() string { return "medorderitems" }

type MedOrderItem struct {
	ID             uint      `gorm:"primaryKey"`
	MedOrderID     uint      `gorm:"not null;index:idx_order_item,unique"`
	MedOrder       MedOrder  `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:RESTRICT;"`
	CriterionID    uint      `gorm:"not null;index:idx_order_item,unique"`
	Criterion      Criterion `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:RESTRICT;"`
	Position       int       `gorm:"not null;default:0"`
	ValueNum       *float64  `gorm:"type:numeric(10,3)"`
	ValueIndicator bool      `gorm:"not null;default:false"`
}
