package services

type Services struct {
	Criteria             CriteriaService
	PankreatitOrders     PankreatitOrdersService
	PankreatitOrderItems PankreatitOrderItemsService
	MedUsers             MedUsersService
}

type Reps struct {
	CriteriaRepo             CriteriaRepoPort
	PankreatitOrdersRepo     PankreatitOrdersRepoPort
	PankreatitOrderItemsRepo PankreatitOrderItemsRepoPort
	MedUsersRepo             MedUsersRepoPort
}

func NewServices(d Reps) *Services {
	return &Services{
		Criteria:             NewCriteriaService(d.CriteriaRepo),
		PankreatitOrders:     NewPankreatitOrdersService(d.PankreatitOrdersRepo),
		PankreatitOrderItems: NewPankreatitOrderItemsService(d.PankreatitOrderItemsRepo),
		MedUsers:             NewMedUsersService(d.MedUsersRepo),
	}
}
