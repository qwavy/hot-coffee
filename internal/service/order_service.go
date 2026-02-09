package service

import "hot-coffee/internal/dal"

type OrderService struct {
	OrderRepository *dal.OrderRepository
}

func NewOrderService(orderRepository *dal.OrderRepository) *OrderService {
	return &OrderService{OrderRepository: orderRepository}
}
