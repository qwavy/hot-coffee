package handler

import (
	"fmt"
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
		fmt.Println(err)
		utils.SendError(w, http.StatusInternalServerError, "Error")
	}

	utils.SendJSON(w, http.StatusOK, menuItems)
}
