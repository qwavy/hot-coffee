package service

import "hot-coffee/internal/dal"

type Services struct {
	Inventory *InventoryService
	Menu      *MenuService
	Order     *OrderService
}

func NewServices(repositories *dal.Repositories) *Services {
	return &Services{
		Inventory: NewInventoryService(repositories.Inventory),
		Menu:      NewMenuService(repositories.Menu),
		Order:     NewOrderService(repositories.Order),
	}
}
