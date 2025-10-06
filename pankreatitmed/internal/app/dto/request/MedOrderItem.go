package request

type MedOrderItemDelete struct {
	MedOrderID  uint `json:"med_order_id" binding:"required"`
	CriterionID uint `json:"criterion_id" binding:"required"`
}

type MedOrderItemUpdate struct {
	MedOrderID  uint    `json:"med_order_id" binding:"required"`
	CriterionID uint    `json:"criterion_id" binding:"required"`
	Position    uint    `json:"position"`
	ValueNum    float64 `json:"value_num"`
}
