package handler


import (
	"hot-coffee/internal/service"
	"hot-coffee/internal/utils"
	"net/http"
	"fmt"
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

