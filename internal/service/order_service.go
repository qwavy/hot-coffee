package service

import (
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type OrderService struct {
	OrderRepository *dal.OrderRepository
}

func NewOrderService(orderRepository *dal.OrderRepository) *OrderService {
	return &OrderService{OrderRepository: orderRepository}
}

func (s *OrderService) GetAll() ([]models.OrderItem, error) {
	orderItem, err := s.OrderRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return orderItem, nil
}