package server

import (
	"hot-coffee/internal/dal"
	"hot-coffee/internal/handler"
	"hot-coffee/internal/service"
	"net/http"
)

func Router() *http.ServeMux {
	r := http.NewServeMux()

	menuRepository := dal.NewMenuRepository("data/menu_items.json")

	menuService := service.NewMenuService(menuRepository)

	menuHandler := handler.NewMenuHandler(menuService)

	orderRepository := dal.NewOrderRepository("data/orders.json")

	orderService := service.NewOrderService(orderRepository)

	orderHandler := handler.NewOrderHandler(orderService)

	
	r.HandleFunc("GET /menu", menuHandler.GetAll)
	r.HandleFunc("GET /menu/{id}", menuHandler.Get)
	r.HandleFunc("GET /order", orderHandler.GetAll)
	r.HandleFunc("GET /order/{id}", orderHandler.Get)

	return r
}