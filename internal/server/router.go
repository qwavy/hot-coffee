package server

import (
	"net/http"

	"hot-coffee/internal/dal"
	"hot-coffee/internal/handler"
	"hot-coffee/internal/service"
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
	// Order
	// POST /orders
	// PUT /orders/{id}
	// POST /orders/{id}/close
	r.HandleFunc("GET /order", orderHandler.GetAll)
	r.HandleFunc("GET /order/{id}", orderHandler.Get)
	r.HandleFunc("DELETE /order/{id}", orderHandler.DeleteById)
	r.HandleFunc("PUT /order/{id}", orderHandler.UpdateById)

	return r
}
