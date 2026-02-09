package handler

import "hot-coffee/internal/service"

type InventoryHandler struct {
	InventoryService service.InventoryService
}

func NewInventoryHandler(inventoryService service.InventoryService) *InventoryHandler {
	return &InventoryHandler{InventoryService: inventoryService}
}
