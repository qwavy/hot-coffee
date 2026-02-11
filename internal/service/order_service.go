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

func (s *OrderService) GetAll() ([]models.Order, error) {
	orderItem, err := s.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return orderItem, nil
}

func (s *OrderService) GetById(productId string) (models.Order, error) {
	order, err := s.OrderRepository.GetById(productId)
	if err != nil {
		return models.Order{}, err
	}

	return order, err
}

func (s *OrderService) DeleteById(productId string) error {
	return s.OrderRepository.DeleteById(productId)
}
