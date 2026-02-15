package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
	"hot-coffee/models"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: orderService}
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	orderItem, err := h.OrderService.GetAll()
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusInternalServerError, "Error")
	}

	utils.SendJSON(w, http.StatusOK, orderItem)
}

func (h *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	orderItem, err := h.OrderService.GetById(productId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
	}

	utils.SendJSON(w, http.StatusOK, orderItem)
}

func (h *OrderHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	if err := h.OrderService.DeleteById(productId); err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *OrderHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	var newOrder models.Order

	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.OrderService.UpdateById(productId, newOrder); err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
