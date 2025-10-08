package services

type Services struct {
	Criteria      CriteriaService
	MedOrders     MedOrdersService
	MedOrderItems MedOrderItemsService
	MedUsers      MedUsersService
}

type Reps struct {
	CriteriaRepo      CriteriaRepoPort
	MedOrdersRepo     MedOrdersRepoPort
	MedOrderItemsRepo MedOrderItemsRepoPort
	MedUsersRepo      MedUsersRepoPort
}

func NewServices(d Reps) *Services {
	return &Services{
		Criteria:      NewCriteriaService(d.CriteriaRepo),
		MedOrders:     NewMedOrdersService(d.MedOrdersRepo),
		MedOrderItems: NewMedOrderItemsService(d.MedOrderItemsRepo),
		MedUsers:      NewMedUsersService(d.MedUsersRepo),
	}
}
