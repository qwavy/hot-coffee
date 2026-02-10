package handler

import (
	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
	"net/http"
)

type MenuHandler struct {
	MenuService *service.MenuService
}

func NewMenuHandler(menuService *service.MenuService) *MenuHandler {
	return &MenuHandler{MenuService: menuService}
}

func (h *MenuHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	menuItems, err := h.MenuService.GetAll()

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Error")
	}

	utils.SendJSON(w, http.StatusOK, menuItems)
}

func (h *MenuHandler) Get(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	menuItem, err := h.MenuService.GetById(productId)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Error")
	}

	utils.SendJSON(w, http.StatusOK, menuItem)
}
