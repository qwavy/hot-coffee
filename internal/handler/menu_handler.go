package handler

import (
	"encoding/json"
	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
	"hot-coffee/models"
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
		utils.SendError(w, http.StatusInternalServerError, err.Error())
	}

	utils.SendJSON(w, http.StatusOK, menuItems)
}

func (h *MenuHandler) GetById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	menuItem, err := h.MenuService.GetById(productId)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
	}

	utils.SendJSON(w, http.StatusOK, menuItem)
}

func (h *MenuHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	err := h.MenuService.DeleteById(productId)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MenuHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var menuItem models.MenuItem

	err := json.NewDecoder(r.Body).Decode(&menuItem)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.MenuService.CreateItem(menuItem)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
