package server

import (
	"hot-coffee/internal/dal"
	"hot-coffee/internal/handler"
	"hot-coffee/internal/service"
	"net/http"
)

func Router() *http.ServeMux {
	r := http.NewServeMux()

	repositories := dal.NewRepositories("data/inventory.json", "data/menu_items.json", "data/orders.json")
	services := service.NewServices(repositories)
	handlers := handler.NewHandlers(services)

	r.HandleFunc("GET /menu", handlers.Menu.GetAll)
	r.HandleFunc("GET /menu/{id}", handlers.Menu.GetById)
	r.HandleFunc("DELETE /menu/{id}", handlers.Menu.DeleteById)
	r.HandleFunc("POST /menu", handlers.Menu.CreateItem)
	r.HandleFunc("PUT /menu/{id}", handlers.Menu.UpdateItem)

	r.HandleFunc("GET /inventory", handlers.Inventory.GetAll)

	return r
}
