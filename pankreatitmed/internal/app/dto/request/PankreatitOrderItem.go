package request

type GetPankreatitOrderItem struct {
	MedOrderID  uint `form:"med_order_id" binding:"required"`
	CriterionID uint `form:"criterion_id" binding:"required"`
}

type PankreatitOrderItemDelete struct {
	MedOrderID  uint `json:"med_order_id" binding:"required"`
	CriterionID uint `json:"criterion_id" binding:"required"`
}

type PankreatitOrderItemUpdate struct {
	Position *uint    `json:"position" binding:"omitempty"`
	ValueNum *float64 `json:"value_num" binding:"omitempty"`
}
