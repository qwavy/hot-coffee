package handler

import (
	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
	"net/http"
)

type InventoryHandler struct {
	InventoryService *service.InventoryService
}

func NewInventoryHandler(inventoryService *service.InventoryService) *InventoryHandler {
	return &InventoryHandler{InventoryService: inventoryService}
}

func (h *InventoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	inventoryItems, err := h.InventoryService.GetAll()

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
	}
	utils.SendJSON(w, http.StatusOK, inventoryItems)
}
