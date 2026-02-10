package router

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

	r.HandleFunc("GET /menu", menuHandler.GetAll)
	r.HandleFunc("GET /menu/{id}", menuHandler.Get)

	return r
}