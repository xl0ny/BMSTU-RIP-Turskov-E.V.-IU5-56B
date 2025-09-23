package ds

type OrderItem struct {
	ID          uint `gorm:"primaryKey"`
	OrderID     uint `gorm:"not null;index:idx_order_item,unique"`
	CriterionID uint `gorm:"not null;index:idx_order_item,unique"`
	Position    int  `gorm:"not null;default:0"`
	ValueNum    *float64
	ValueText   bool
}
