package services

type Services struct {
	Criteria CriteriaService
	// Orders  OrdersService
	// Items   OrderItemsService
	// Users   UsersService
}

type Reps struct {
	CriteriaRepo CriteriaRepoPort
	// OrdersRepo  OrdersRepoPort
	// UsersRepo   UsersRepoPort
}

func NewServices(d Reps) *Services {
	return &Services{
		Criteria: NewCriteriaService(d.CriteriaRepo),
		// Orders:  NewOrdersService(d.OrdersRepo, ...),
		// Items:   NewOrderItemsService(...),
		// Users:   NewUsersService(d.UsersRepo),
	}
}
