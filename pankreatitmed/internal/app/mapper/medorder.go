package mapper

import (
	"pankreatitmed/internal/app/ds"
	"pankreatitmed/internal/app/dto/response"
)

func MedOrderToSendMedOrder(o *ds.MedOrder, amount uint) response.SendCartMedOrder {
	return response.SendCartMedOrder{
		MedOrderId:     o.ID,
		CriteriaAmount: amount,
	}
}

func MedOrderToSendMedOrders(mo ds.MedOrder) response.SendMedOrders {
	return response.SendMedOrders{
		ID:            mo.ID,
		Status:        mo.Status,
		FormedAt:      mo.FormedAt,
		FinishedAt:    mo.FinishedAt,
		RansonScore:   mo.RansonScore,
		MortalityRisk: mo.MortalityRisk,
	}
}

func MedOrdersToSendMedOrders(mos []ds.MedOrder) []response.SendMedOrders {
	list := make([]response.SendMedOrders, len(mos))
	for i, c := range mos {
		list[i] = MedOrderToSendMedOrders(c)
	}
	return list
}

func MedOrderToSendMedOrderWithItems(mo ds.MedOrder, itms []ds.MedOrderItem) response.SendMedOrder {
	return response.SendMedOrder{
		ID:            mo.ID,
		Status:        mo.Status,
		CreatorID:     mo.CreatorID,
		FormedAt:      mo.FormedAt,
		FinishedAt:    mo.FinishedAt,
		ModeratorID:   mo.ModeratorID,
		RansonScore:   mo.RansonScore,
		MortalityRisk: mo.MortalityRisk,
		Items:         MedOrderItemsToSendMedOrderItems(itms),
	}
}
