package handler

import (
	"encoding/json"
	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
	"hot-coffee/models"
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

func (h *InventoryHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	inventoryItems, err := h.InventoryService.GetById(id)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
	}
	utils.SendJSON(w, http.StatusOK, inventoryItems)
}

func (h *InventoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var inventoryItem models.InventoryItem

	err := json.NewDecoder(r.Body).Decode(&inventoryItem)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.InventoryService.Create(inventoryItem)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *InventoryHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	defer r.Body.Close()

	var newInventoryItem models.InventoryItem

	json.NewDecoder(r.Body).Decode(&newInventoryItem)

	err := h.InventoryService.Update(id, newInventoryItem)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(http.StatusOK)
}

func (h *InventoryHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.InventoryService.DeleteById(id)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(http.StatusNoContent)
}
