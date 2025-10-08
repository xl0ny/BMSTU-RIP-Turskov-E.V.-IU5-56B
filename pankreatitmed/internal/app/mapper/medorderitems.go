package mapper

import (
	"pankreatitmed/internal/app/ds"
	"pankreatitmed/internal/app/dto/response"
)

func MedOrderItemToSendMedOrderItem(item ds.MedOrderItem) response.SendMedOrderItem {
	return response.SendMedOrderItem{
		ID:             item.ID,
		CriterionID:    item.CriterionID,
		Criterion:      item.Criterion,
		Position:       item.Position,
		ValueNum:       item.ValueNum,
		ValueIndicator: item.ValueIndicator,
	}
}

func MedOrderItemsToSendMedOrderItems(items []ds.MedOrderItem) []response.SendMedOrderItem {
	list := make([]response.SendMedOrderItem, len(items))
	for i, c := range items {
		list[i] = MedOrderItemToSendMedOrderItem(c)
	}
	return list
}
