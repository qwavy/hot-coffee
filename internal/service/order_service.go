package service

import (
	"hot-coffee/internal/dal"
	"hot-coffee/models"
)

type OrderService struct {
	OrderRepository     *dal.OrderRepository
	MenuRepository      *dal.MenuRepository
	InventoryRepository *dal.InventoryRepository
}

func NewOrderService(orderRepository *dal.OrderRepository, menuRepository *dal.MenuRepository, inventoryRepository *dal.InventoryRepository) *OrderService {
	return &OrderService{OrderRepository: orderRepository, MenuRepository: menuRepository, InventoryRepository: inventoryRepository}
}

func (s *OrderService) GetAll() ([]models.Order, error) {
	orderItem, err := s.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return orderItem, nil
}

func (s *OrderService) GetById(productId string) (*models.Order, error) {
	order, err := s.OrderRepository.GetById(productId)
	if err != nil {
		return nil, err
	}

	return &order, err
}

func (s *OrderService) DeleteById(productId string) error {
	return s.OrderRepository.DeleteById(productId)
}

func (s *OrderService) Create(order models.Order) error {
	orderItems := order.Items

	for _, value := range orderItems {
		menuItem, err := s.MenuRepository.GetById(value.ProductID)
		if err != nil {
			return err
		}

		for _, menuItemIngredient := range menuItem.Ingredients {
			inventoryItem, err := s.InventoryRepository.GetById(menuItemIngredient.IngredientID)
			if err != nil {
				return err
			}

			if inventoryItem.Quantity >= menuItemIngredient.Quantity {
				inventoryItem.Quantity -= menuItemIngredient.Quantity
			}

			s.InventoryRepository.UpdateById(inventoryItem.IngredientID, *inventoryItem)
		}

	}

	return s.OrderRepository.Create(order)
}

func (s *OrderService) UpdateById(productId string, newOrder models.Order) error {
	return s.OrderRepository.UpdateById(productId, newOrder)
}
